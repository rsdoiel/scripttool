package main

import (
	"bytes"
	"errors"
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
