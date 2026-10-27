package scripttool

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

/** ErrUsage marks an error as a mistake in the command line: nothing was
 * attempted. cmd/scripttool exits with status 2 for it (workspace DR-0003).
 *
 * Example:
 *   err := fmt.Errorf("%w: unknown help topic", ErrUsage)
 *   errors.Is(err, ErrUsage) // true
 */
var ErrUsage = errors.New("usage")

/** HelpTopic is one page of the help system. Text is Pandoc man page source,
 * so "scripttool help TOPIC" can be redirected to scripttool-TOPIC.7.md.
 *
 * Fields:
 *   Name (string) — the word typed after "help", lower case
 *   Summary (string) — one line shown by "help index"
 *   Text (string) — Markdown; {app_name} and {version} are substituted
 *
 * Example:
 *   topic := HelpTopic{Name: "demo", Summary: "a demo", Text: "# NAME\n\ndemo"}
 */
type HelpTopic struct {
	Name    string
	Summary string
	Text    string
}

/** HelpTopics lists every topic "scripttool help index" shows, in the order
 * shown. The Makefile generates one scripttool-TOPIC.7.md from each.
 *
 * Example:
 *   for _, t := range HelpTopics { fmt.Println(t.Name) }
 */
var HelpTopics = []HelpTopic{
	{"fadein", "convert from FadeIn files", `% {app_name}-fadein(7) {app_name} {version}
% R. S. Doiel

# NAME

{app_name}-fadein - converting from FadeIn files

# SYNOPSIS

{app_name} VERB FILENAME.fadein [OUTPUT_FILENAME]

# DESCRIPTION

A FadeIn file is a zipped XML document, so it cannot be read from standard
input. The input filename is required for these verbs.

fadein2fdx
: to Final Draft XML

fadein2fountain
: to Fountain

fadein2json
: to JSON

fadein2yaml
: to YAML

fadein2osf
: to Open Screenplay Format 2.0

# EXAMPLES

~~~shell
    {app_name} fadein2fountain screenplay.fadein screenplay.fountain
~~~
`},
	{"fdx", "convert from Final Draft XML", `% {app_name}-fdx(7) {app_name} {version}
% R. S. Doiel

# NAME

{app_name}-fdx - converting from Final Draft XML

# SYNOPSIS

{app_name} VERB [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

Without an input filename these verbs read standard input and write
standard output.

fdx2fadein
: to FadeIn. The output filename is required.

fdx2fountain
: to Fountain. A scene number becomes a trailing #12A# in the scene heading.

fdx2json
: to JSON

fdx2yaml
: to YAML

fdx2osf
: to Open Screenplay Format 2.0

To write Final Draft XML use fountain2fdx, osf2fdx or fadein2fdx.

# EXAMPLES

~~~shell
    {app_name} fdx2fountain screenplay.fdx screenplay.fountain
~~~
`},
	{"osf", "convert from Open Screenplay Format", `% {app_name}-osf(7) {app_name} {version}
% R. S. Doiel

# NAME

{app_name}-osf - converting from Open Screenplay Format 2.0

# SYNOPSIS

{app_name} VERB [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

Without an input filename these verbs read standard input and write
standard output. Open Screenplay Format keeps no per-scene number.

osf2fadein
: to FadeIn. The output filename is required.

osf2fdx
: to Final Draft XML

osf2fountain
: to Fountain

osf2json
: to JSON

osf2yaml
: to YAML

# EXAMPLES

~~~shell
    {app_name} osf2fountain screenplay.osf screenplay.fountain
~~~
`},
	{"fountain", "convert, format and render Fountain text", `% {app_name}-fountain(7) {app_name} {version}
% R. S. Doiel

# NAME

{app_name}-fountain - converting, formatting and rendering Fountain

# SYNOPSIS

{app_name} VERB [OPTIONS] [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

Without an input filename these verbs read standard input and write
standard output.

fountain2fadein
: to FadeIn. The output filename is required.

fountain2fdx
: to Final Draft XML. A trailing #12A# in a scene heading becomes the scene's
  Number attribute.

fountain2osf
: to Open Screenplay Format 2.0

fountain2json
: to JSON

fountain2yaml
: to YAML

fountain2fountain
: parse and write the document back out as Fountain

fountainfmt
: pretty print a Fountain document; see the options topic for -notes,
  -synopsis, -section and -width

fountain2html
: Fountain scrippets HTML; see the options topic for -html, -inline-css
  and -link-css

# EXAMPLES

~~~shell
    {app_name} fountain2fdx screenplay.fountain screenplay.fdx
    {app_name} fountainfmt -width 60 screenplay.fountain
~~~
`},
	{"characters", "list the characters in a Fountain file", `% {app_name}-characters(7) {app_name} {version}
% R. S. Doiel

# NAME

{app_name}-characters - list the characters in a Fountain file

# SYNOPSIS

{app_name} characters [-alpha] [INPUT_FILENAME]

# DESCRIPTION

Parses the Fountain document and counts the references to each character
in its elements. The list is in order of appearance; -alpha sorts it
alphabetically instead. The input must be Fountain.

# EXAMPLES

~~~shell
    {app_name} characters screenplay.fountain
    {app_name} characters -alpha screenplay.fountain
~~~
`},
	{"options", "options that follow a verb", `% {app_name}-options(7) {app_name} {version}
% R. S. Doiel

# NAME

{app_name}-options - options that follow a verb

# SYNOPSIS

{app_name} VERB [OPTIONS] [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

-i FILENAME
: input filename

-o FILENAME
: output filename

-notes
: include notes in output

-synopsis
: include synopsis in output

-section
: include section headings in output

-width N
: maximum line width

-html
: output a full HTML page

-inline-css
: include inline CSS (with -html)

-link-css
: include a CSS link (with -html)

-pretty
: pretty print output

-alpha
: sort a character list alphabetically

A filename of - means standard input or standard output.
`},
	{"examples", "worked examples", `% {app_name}-examples(7) {app_name} {version}
% R. S. Doiel

# NAME

{app_name}-examples - worked examples

# EXAMPLES

Convert screenplay.fdx to screenplay.fountain, two ways.

~~~shell
    {app_name} fdx2fountain screenplay.fdx screenplay.fountain
    {app_name} fdx2fountain -i screenplay.fdx -o screenplay.fountain
~~~

Read standard input and write standard output.

~~~shell
    cat screenplay.fountain | {app_name} fountain2fdx >screenplay.fdx
~~~

List the characters in order of appearance, then alphabetically.

~~~shell
    {app_name} characters screenplay.fountain
    {app_name} characters -alpha screenplay.fountain
~~~
`},
}

/** HelpIndex returns the text "help index" prints: one line per topic, the
 * topic name first, then its summary, so a script can read the first word.
 *
 * Parameters:
 *   appName (string) — the program name (kept for symmetry with the other
 *   help functions; the index itself does not print it)
 *
 * Returns:
 *   string — the index, ending in a newline
 *
 * Example:
 *   fmt.Print(HelpIndex("scripttool"))
 */
func HelpIndex(appName string) string {
	_ = appName
	var b strings.Builder
	for _, t := range HelpTopics {
		fmt.Fprintf(&b, "%-12s %s\n", t.Name, t.Summary)
	}
	return b.String()
}

/** RunHelp writes the manual, the index or one topic to out.
 *
 * Parameters:
 *   out (io.Writer) — where the help text goes
 *   appName (string) — the program name substituted into the text
 *   args ([]string) — what followed "help": nothing, "index" or a topic name
 *
 * Returns:
 *   error — nil on success; wraps ErrUsage for an unknown topic or more than
 *   one argument
 *
 * Example:
 *   err := RunHelp(os.Stdout, "scripttool", []string{"index"})
 */
func RunHelp(out io.Writer, appName string, args []string) error {
	switch len(args) {
	case 0:
		_, err := fmt.Fprintln(out, FmtCliText(HelpText, appName, "help", Version))
		return err
	case 1:
		name := strings.ToLower(strings.TrimSpace(args[0]))
		if name == "index" {
			_, err := fmt.Fprint(out, HelpIndex(appName))
			return err
		}
		for _, t := range HelpTopics {
			if t.Name == name {
				_, err := fmt.Fprint(out, FmtCliText(t.Text, appName, "help", Version))
				return err
			}
		}
		return fmt.Errorf("%w: unknown help topic %q, run '%s help index' for the list", ErrUsage, args[0], appName)
	default:
		return fmt.Errorf("%w: help takes at most one topic, got %d; run '%s help index' for the list", ErrUsage, len(args), appName)
	}
}
