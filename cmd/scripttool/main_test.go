package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// buildScripttool compiles the command into a temp directory and returns the
// path of the binary, so the tests exercise main and its exit codes.
func buildScripttool(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "scripttool")
	if out, err := exec.Command("go", "build", "-o", bin, ".").CombinedOutput(); err != nil {
		t.Fatalf("go build: %s\n%s", err, out)
	}
	return bin
}

// runBin runs the binary and returns stdout, stderr and the exit code.
func runBin(t *testing.T, bin string, args ...string) (string, string, int) {
	t.Helper()
	cmd := exec.Command(bin, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	code := 0
	if err := cmd.Run(); err != nil {
		var ee *exec.ExitError
		if !errors.As(err, &ee) {
			t.Fatal(err)
		}
		code = ee.ExitCode()
	}
	return stdout.String(), stderr.String(), code
}

// Exit codes follow workspace DR-0003: 0 ok, 2 usage.
func TestExitCodes(t *testing.T) {
	bin := buildScripttool(t)
	tests := []struct {
		name   string
		args   []string
		code   int
		stdout string
		stderr string
	}{
		{"help index", []string{"help", "index"}, 0, "fdx", ""},
		{"help topic", []string{"help", "fdx"}, 0, "# NAME", ""},
		{"unknown topic", []string{"help", "nosuchtopic"}, 2, "", "help index"},
		{"surplus argument", []string{"help", "index", "extra"}, 2, "", "help"},
		{"bad flag", []string{"fdx2fountain", "-nosuchflag"}, 2, "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, stderr, code := runBin(t, bin, tt.args...)
			if code != tt.code {
				t.Errorf("exit code %d, want %d\nstdout: %s\nstderr: %s", code, tt.code, stdout, stderr)
			}
			if !strings.Contains(stdout, tt.stdout) {
				t.Errorf("stdout lacks %q:\n%s", tt.stdout, stdout)
			}
			if !strings.Contains(stderr, tt.stderr) {
				t.Errorf("stderr lacks %q:\n%s", tt.stderr, stderr)
			}
		})
	}
}

// writeFile writes content to dir/name and returns the path.
func writeFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	p := filepath.Join(dir, name)
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

// runBinIn is runBin with a working directory and standard input.
func runBinIn(t *testing.T, bin, dir, stdin string, args ...string) (string, string, int) {
	t.Helper()
	cmd := exec.Command(bin, args...)
	cmd.Dir = dir
	cmd.Stdin = strings.NewReader(stdin)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	code := 0
	if err := cmd.Run(); err != nil {
		var ee *exec.ExitError
		if !errors.As(err, &ee) {
			t.Fatal(err)
		}
		code = ee.ExitCode()
	}
	return stdout.String(), stderr.String(), code
}

// The rest of DR-0003: 65 for content that cannot be read, 66 for a missing
// input, 73 for an output that cannot be created, usage errors are 2, and
// every message is one line led by the program name.
func TestExitCodesDR0003(t *testing.T) {
	bin := buildScripttool(t)
	dir := t.TempDir()
	missing := filepath.Join(dir, "missing.in")
	garbage := writeFile(t, dir, "garbage.xml", "this is not xml <")
	good, _ := filepath.Abs("../../testdata/sample-01.fountain")
	noDir := filepath.Join(dir, "no", "such", "dir", "out.json")
	tests := []struct {
		name  string
		stdin string
		args  []string
		code  int
	}{
		{"no arguments", "", nil, 2},
		{"unknown verb", "", []string{"nosuchverb", good}, 2},
		{"unknown option", "", []string{"-nosuchoption"}, 2},
		{"surplus argument", "", []string{"fountain2json", good, filepath.Join(dir, "a"), filepath.Join(dir, "b")}, 2},
		{"fadein2yaml without input", "", []string{"fadein2yaml"}, 2},
		{"fadein2json without input", "", []string{"fadein2json"}, 2},
		{"fdx2fadein without output", "<x/>", []string{"fdx2fadein"}, 2},
		{"missing fountain input", "", []string{"fountain2json", missing}, 66},
		{"missing fdx input with -i", "", []string{"fdx2json", "-i", missing}, 66},
		{"missing fadein input", "", []string{"fadein2json", missing}, 66},
		{"fdx input is not xml", "", []string{"fdx2json", garbage}, 65},
		{"osf input is not xml", "this is not xml <", []string{"osf2json"}, 65},
		{"output cannot be created", "", []string{"fountain2json", good, noDir}, 73},
		{"help of a verb", "", []string{"fountain2json", "-h"}, 0},
		{"short -h", "", []string{"-h"}, 0},
		{"short -l", "", []string{"-l"}, 0},
		{"short -v", "", []string{"-v"}, 0},
		{"good conversion", "", []string{"fountain2json", good}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout, stderr, code := runBinIn(t, bin, dir, tt.stdin, tt.args...)
			if code != tt.code {
				t.Errorf("exit %d, want %d\nstdout: %.200s\nstderr: %.300s", code, tt.code, stdout, stderr)
			}
			if tt.code != 0 && !strings.HasPrefix(stderr, "scripttool: ") {
				t.Errorf("stderr does not start with the program name: %.200q", stderr)
			}
			if tt.code != 0 && tt.name != "no arguments" && tt.name != "unknown option" && strings.Count(strings.TrimRight(stderr, "\n"), "\n") > 0 && tt.code != 2 {
				t.Errorf("stderr is more than one line: %q", stderr)
			}
		})
	}
}

// A failing run does not create the output file, and does not empty one that
// is already there.
func TestFailedRunLeavesOutputAlone(t *testing.T) {
	bin := buildScripttool(t)
	dir := t.TempDir()
	garbage := writeFile(t, dir, "garbage.xml", "this is not xml <")
	keep := writeFile(t, dir, "keep.json", "keep")
	if _, _, code := runBinIn(t, bin, dir, "", "fdx2json", garbage, keep); code != 65 {
		t.Errorf("exit %d, want 65", code)
	}
	if b, _ := os.ReadFile(keep); string(b) != "keep" {
		t.Errorf("a bad input emptied the existing output: %q", b)
	}
	fresh := filepath.Join(dir, "fresh.json")
	runBinIn(t, bin, dir, "", "fdx2json", garbage, fresh)
	runBinIn(t, bin, dir, "", "nosuchverb", garbage, fresh)
	if _, err := os.Stat(fresh); err == nil {
		t.Errorf("a failed run created %s", fresh)
	}
}
