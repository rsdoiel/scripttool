// Copyright (C) 2017-2026 R. S. Doiel
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package scripttool

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// These tests run release.bash (workspace DR-0013) in a scratch repository
// with a bare origin and a stand-in for the gh command. The stand-in serves
// the release JSON it finds in $FAKE_GH_VIEW (exit 1 when there is none) and
// appends every call to $FAKE_GH_LOG, so nothing reaches GitHub.

const fakeGh = `#!/bin/sh
echo "$*" >> "$FAKE_GH_LOG"
case "$1 $2" in
"release view")
	if [ -f "$FAKE_GH_VIEW" ]; then cat "$FAKE_GH_VIEW"; exit 0; fi
	exit 1 ;;
esac
exit 0
`

type releaseRepo struct {
	dir, log, view, head string
	t                    *testing.T
}

func sha(b []byte) string {
	h := sha256.Sum256(b)
	return "sha256:" + hex.EncodeToString(h[:])
}

// newReleaseRepo makes a pushed repository holding release.bash, the Makefile
// (for check-committed), a codemeta.json for v9.9.9 and a dist/ with two zips.
func newReleaseRepo(t *testing.T) *releaseRepo {
	t.Helper()
	for _, tool := range []string{"git", "make", "jq", "bash", "sha256sum"} {
		if _, err := exec.LookPath(tool); err != nil {
			t.Skipf("%s not available", tool)
		}
	}
	root := t.TempDir()
	r := &releaseRepo{dir: filepath.Join(root, "x"), log: filepath.Join(root, "gh.log"), view: filepath.Join(root, "view.json"), t: t}
	origin := filepath.Join(root, "origin.git")
	bin := filepath.Join(root, "bin")
	for _, d := range []string{bin, filepath.Join(r.dir, "dist")} {
		if err := os.MkdirAll(d, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	r.git(root, "init", "-q", "--bare", origin)
	r.git(root, "init", "-q", "-b", "main", r.dir)
	r.git(r.dir, "remote", "add", "origin", origin)
	for _, f := range []string{"release.bash", "Makefile"} {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		r.write(f, string(b), 0o755)
	}
	r.write("codemeta.json", `{"version": "9.9.9", "releaseNotes": "Notes."}`, 0o644)
	r.write(".gitignore", "dist/\n", 0o644)
	r.git(r.dir, "add", "-A")
	r.git(r.dir, "commit", "-q", "-m", "one")
	r.git(r.dir, "push", "-q", "origin", "main")
	r.head = strings.TrimSpace(r.git(r.dir, "rev-parse", "HEAD"))
	for name, body := range map[string]string{"x-v9.9.9-Linux.zip": "linux", "x-v9.9.9-macOS.zip": "macos"} {
		r.write(filepath.Join("dist", name), body, 0o644)
	}
	if err := os.WriteFile(filepath.Join(bin, "gh"), []byte(fakeGh), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", bin+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("FAKE_GH_LOG", r.log)
	t.Setenv("FAKE_GH_VIEW", r.view)
	return r
}

func (r *releaseRepo) write(name, body string, mode os.FileMode) {
	r.t.Helper()
	if err := os.WriteFile(filepath.Join(r.dir, name), []byte(body), mode); err != nil {
		r.t.Fatal(err)
	}
}

func (r *releaseRepo) git(dir string, args ...string) string {
	r.t.Helper()
	c := exec.Command("git", args...)
	c.Dir = dir
	c.Env = append(os.Environ(), "GIT_AUTHOR_NAME=t", "GIT_AUTHOR_EMAIL=t@example.org", "GIT_COMMITTER_NAME=t", "GIT_COMMITTER_EMAIL=t@example.org")
	out, err := c.CombinedOutput()
	if err != nil {
		r.t.Fatalf("git %v: %v\n%s", args, err, out)
	}
	return string(out)
}

// asset is one entry of the stand-in's `gh release view` answer.
type asset struct{ name, body, state string }

// draft writes the stand-in's release JSON: a release for v9.9.9 that is a
// draft or not, aimed at target, holding assets whose digest is the digest of body.
func (r *releaseRepo) draft(isDraft bool, target string, assets ...asset) {
	var list []string
	for _, a := range assets {
		state := a.state
		if state == "" {
			state = "uploaded"
		}
		list = append(list, fmt.Sprintf(`{"name":%q,"size":%d,"state":%q,"digest":%q}`, a.name, len(a.body), state, sha([]byte(a.body))))
	}
	j := fmt.Sprintf(`{"isDraft":%t,"tagName":"v9.9.9","targetCommitish":%q,"assets":[%s]}`, isDraft, target, strings.Join(list, ","))
	if err := os.WriteFile(r.view, []byte(j), 0o644); err != nil {
		r.t.Fatal(err)
	}
}

// release runs release.bash with answer on its input and returns the combined
// output, the exit code and the lines of the gh log.
func (r *releaseRepo) release(answer string, args ...string) (string, int, []string) {
	r.t.Helper()
	os.Remove(r.log)
	c := exec.Command("bash", append([]string{"release.bash"}, args...)...)
	c.Dir = r.dir
	c.Stdin = strings.NewReader(answer)
	out, err := c.CombinedOutput()
	code := 0
	if err != nil {
		ee, ok := err.(*exec.ExitError)
		if !ok {
			r.t.Fatal(err)
		}
		code = ee.ExitCode()
	}
	var calls []string
	if b, err := os.ReadFile(r.log); err == nil {
		calls = strings.Split(strings.TrimSpace(string(b)), "\n")
	}
	return string(out), code, calls
}

// uploads picks the `release upload` calls out of the gh log.
func uploads(calls []string) (names []string, clobbered map[string]bool) {
	clobbered = map[string]bool{}
	for _, c := range calls {
		if !strings.HasPrefix(c, "release upload ") {
			continue
		}
		f := strings.Fields(c)
		name := filepath.Base(f[3])
		names = append(names, name)
		if strings.Contains(c, "--clobber") {
			clobbered[name] = true
		}
	}
	return names, clobbered
}

func created(calls []string) bool {
	for _, c := range calls {
		if strings.HasPrefix(c, "release create ") {
			return true
		}
	}
	return false
}

/** TestReleaseNewDraft is the unchanged path: no release exists, so the draft is
 * created at HEAD and the checksums file and both zips are uploaded.
 */
func TestReleaseNewDraft(t *testing.T) {
	r := newReleaseRepo(t)
	out, code, calls := r.release("y\n")
	names, _ := uploads(calls)
	if code != 0 || !created(calls) || len(names) != 3 {
		t.Fatalf("exit %d, created %v, uploads %v\n%s", code, created(calls), names, out)
	}
	if !strings.Contains(strings.Join(calls, "\n"), "--target "+r.head) {
		t.Errorf("the draft is not created at HEAD %s:\n%s", r.head, strings.Join(calls, "\n"))
	}
	if _, err := os.Stat(filepath.Join(r.dir, "release_notes.tmp")); err == nil {
		t.Error("release_notes.tmp was left behind")
	}
}

/** TestReleaseResumesInterruptedDraft is the fdx v1.0.5 failure: a draft for HEAD
 * holds all but one of the files. The run uploads only the missing one and does
 * not create a second release.
 */
func TestReleaseResumesInterruptedDraft(t *testing.T) {
	r := newReleaseRepo(t)
	// release.bash writes the checksums file; a dry run leaves it in dist/ so the draft can hold a copy.
	if _, code, _ := r.release("", "--dry-run"); code != 0 {
		t.Fatal("dry run failed")
	}
	sumBody, err := os.ReadFile(filepath.Join(r.dir, "dist", "x-v9.9.9-checksums.txt"))
	if err != nil {
		t.Fatal(err)
	}
	r.draft(true, r.head,
		asset{"x-v9.9.9-checksums.txt", string(sumBody), ""},
		asset{"x-v9.9.9-Linux.zip", "linux", ""})
	out, code, calls := r.release("y\n")
	names, clobbered := uploads(calls)
	if code != 0 || created(calls) || len(names) != 1 || names[0] != "x-v9.9.9-macOS.zip" || len(clobbered) != 0 {
		t.Fatalf("exit %d, created %v, uploads %v, clobbered %v\n%s", code, created(calls), names, clobbered, out)
	}
	if !strings.Contains(out, "Skipped (already uploaded): dist/x-v9.9.9-Linux.zip") || !strings.Contains(out, "Resume:") {
		t.Errorf("the output does not say it is resuming and what it skipped:\n%s", out)
	}
}

/** TestReleaseReplacesWrongAssets: an asset that is on the draft but not the file
 * in dist/ (a partial upload, or a zip rebuilt since) is uploaded again with
 * --clobber, even when its size is the same; a correct one is left alone.
 */
func TestReleaseReplacesWrongAssets(t *testing.T) {
	r := newReleaseRepo(t)
	if _, code, _ := r.release("", "--dry-run"); code != 0 {
		t.Fatal("dry run failed")
	}
	sumBody, _ := os.ReadFile(filepath.Join(r.dir, "dist", "x-v9.9.9-checksums.txt"))
	r.draft(true, r.head,
		asset{"x-v9.9.9-checksums.txt", string(sumBody), ""},
		asset{"x-v9.9.9-Linux.zip", "LINUX", ""},        // same size, different bytes
		asset{"x-v9.9.9-macOS.zip", "macos", "starter"}) // the right bytes but never finished
	out, code, calls := r.release("y\n")
	names, clobbered := uploads(calls)
	if code != 0 || len(names) != 2 || !clobbered["x-v9.9.9-Linux.zip"] || !clobbered["x-v9.9.9-macOS.zip"] || clobbered["x-v9.9.9-checksums.txt"] {
		t.Fatalf("exit %d, uploads %v, clobbered %v\n%s", code, names, clobbered, out)
	}
}

/** TestReleaseRefusesWhatItMustNotResume: a published release, a draft for
 * another commit and a draft whose target is a branch name stop the script with
 * one line and no upload.
 */
func TestReleaseRefusesWhatItMustNotResume(t *testing.T) {
	r := newReleaseRepo(t)
	for name, c := range map[string]struct {
		draft  bool
		target string
		want   string
	}{
		"published":        {false, r.head, "not a draft"},
		"other commit":     {true, strings.Repeat("a", 40), "targets"},
		"branch as target": {true, "main", "targets"},
	} {
		r.draft(c.draft, c.target)
		out, code, calls := r.release("y\n")
		names, _ := uploads(calls)
		if code == 0 || created(calls) || len(names) != 0 || !strings.Contains(out, c.want) {
			t.Errorf("%s: exit %d, uploads %v, output:\n%s", name, code, names, out)
		}
	}
}

/** TestReleaseDryRun reports skip, upload and replace for a draft and changes
 * nothing: no create and no upload.
 */
func TestReleaseDryRun(t *testing.T) {
	r := newReleaseRepo(t)
	r.draft(true, r.head, asset{"x-v9.9.9-Linux.zip", "linux", ""}, asset{"x-v9.9.9-macOS.zip", "MACOS", ""})
	out, code, calls := r.release("", "--dry-run")
	names, _ := uploads(calls)
	if code != 0 || created(calls) || len(names) != 0 {
		t.Fatalf("exit %d, created %v, uploads %v\n%s", code, created(calls), names, out)
	}
	for _, want := range []string{"skip", "upload", "replace"} {
		if !strings.Contains(out, want) {
			t.Errorf("the dry run does not name a %q:\n%s", want, out)
		}
	}
}

/** TestReleaseNeedsCommittedTree: an uncommitted change stops a resume as it
 * stops a new release.
 */
func TestReleaseNeedsCommittedTree(t *testing.T) {
	r := newReleaseRepo(t)
	r.write("codemeta.json", `{"version": "9.9.9", "releaseNotes": "Changed."}`, 0o644)
	r.draft(true, r.head)
	_, code, calls := r.release("y\n")
	if names, _ := uploads(calls); code == 0 || len(names) != 0 {
		t.Errorf("exit %d, uploads %v; want a refusal", code, names)
	}
}
