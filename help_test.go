package scripttool

import (
	"os"
	"strings"
	"testing"
)

// runHelp runs RunScripttool with the given arguments after the program
// name and returns what it wrote to standard output.
func runHelp(t *testing.T, args ...string) (string, error) {
	t.Helper()
	outFile, err := os.CreateTemp(t.TempDir(), "out-*")
	if err != nil {
		t.Fatal(err)
	}
	defer outFile.Close()
	runErr := RunScripttool(os.Stdin, outFile, os.Stderr, append([]string{"scripttool"}, args...))
	b, err := os.ReadFile(outFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	return string(b), runErr
}

// topicNames returns the first word of every non-blank line of a help index.
func topicNames(index string) []string {
	var names []string
	for _, line := range strings.Split(index, "\n") {
		if f := strings.Fields(line); len(f) > 0 {
			names = append(names, f[0])
		}
	}
	return names
}

func TestHelpIndexListsTopics(t *testing.T) {
	got, err := runHelp(t, "help", "index")
	if err != nil {
		t.Fatalf("help index: %s", err)
	}
	names := topicNames(got)
	for _, want := range []string{"fadein", "fdx", "osf", "fountain", "characters", "options", "examples"} {
		found := false
		for _, name := range names {
			if name == want {
				found = true
			}
		}
		if !found {
			t.Errorf("help index does not list topic %q, got:\n%s", want, got)
		}
	}
}

func TestHelpTopicForEveryIndexEntry(t *testing.T) {
	index, err := runHelp(t, "help", "index")
	if err != nil {
		t.Fatalf("help index: %s", err)
	}
	manual, _ := runHelp(t, "help")
	for _, name := range topicNames(index) {
		got, err := runHelp(t, "help", name)
		if err != nil {
			t.Errorf("help %s: %s", name, err)
			continue
		}
		if strings.TrimSpace(got) == "" {
			t.Errorf("help %s is empty", name)
		}
		if got == manual {
			t.Errorf("help %s printed the whole manual instead of the topic", name)
		}
		if !strings.Contains(got, "# NAME") || !strings.Contains(got, name) {
			t.Errorf("help %s is not a man page source naming the topic:\n%s", name, got)
		}
	}
}

func TestHelpUnknownTopic(t *testing.T) {
	_, err := runHelp(t, "help", "nosuchtopic")
	if err == nil {
		t.Fatal("help nosuchtopic succeeded, want an error")
	}
	if !strings.Contains(err.Error(), "nosuchtopic") || !strings.Contains(err.Error(), "help index") {
		t.Errorf("error should name the topic and point at help index, got %q", err)
	}
}

func TestHelpWithoutTopicStillPrintsManual(t *testing.T) {
	got, err := runHelp(t, "help")
	if err != nil {
		t.Fatalf("help: %s", err)
	}
	if !strings.Contains(got, "# SYNOPSIS") {
		t.Errorf("help should print the manual, got:\n%s", got)
	}
}

// The manual must name every verb the dispatcher accepts and point at the
// help index, so a reader of scripttool(1) can find the topic pages.
func TestManualDocumentsEveryVerb(t *testing.T) {
	for _, verb := range []string{
		"fadein2yaml", "fdx2yaml", "fountain2yaml", "osf2yaml",
		"fountain2fountain", "fountain2html", "fountainfmt", "characters",
	} {
		if !strings.Contains(HelpText, verb) {
			t.Errorf("HelpText does not document verb %q", verb)
		}
	}
	if !strings.Contains(HelpText, "help index") {
		t.Error("HelpText does not mention `help index`")
	}
}

// The Makefile lists the help topics so it can build one man page for each.
// A glob over *.7.md would quietly miss a topic whose page does not exist yet,
// so the list is explicit and this test keeps it equal to HelpTopics.
func TestMakefileHelpTopicsMatch(t *testing.T) {
	src, err := os.ReadFile("Makefile")
	if err != nil {
		t.Fatal(err)
	}
	var listed []string
	for _, line := range strings.Split(string(src), "\n") {
		if rest, ok := strings.CutPrefix(line, "HELP_TOPICS"); ok {
			_, list, _ := strings.Cut(rest, "=")
			listed = strings.Fields(list)
		}
	}
	var want []string
	for _, topic := range HelpTopics {
		want = append(want, topic.Name)
	}
	if strings.Join(listed, " ") != strings.Join(want, " ") {
		t.Errorf("Makefile HELP_TOPICS = %q, HelpTopics = %q", listed, want)
	}
}
