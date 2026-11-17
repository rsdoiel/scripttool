package scripttool

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"

	// My packages
	"github.com/rsdoiel/fountain"
)

func FmtCliText(s string, appName string, verb string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "{app_name}", appName), "{verb}", verb), "{version}", version)
}

func FmtCliHelp(appName string, verb string) string {
	return fmt.Sprintf("%s %s is not documented", appName, verb)
}

// The exit codes of every command line tool in the workspace (DR-0003).
const (
	ExitOK         = 0
	ExitUsage      = 2
	ExitData       = 65
	ExitNoInput    = 66
	ExitInternal   = 70
	ExitCantCreate = 73
	ExitIO         = 74
	ExitNoPerm     = 77
)

/** ExitError is an error that knows its exit code (DR-0003). RunScripttool
 * returns one for every failure it classifies.
 *
 * Example:
 *   err := &ExitError{Code: ExitNoInput, Err: os.ErrNotExist}
 *   ExitCode(err) // 66
 */
type ExitError struct {
	Code int
	Err  error
}

// Error returns the message of the wrapped error.
func (e *ExitError) Error() string { return e.Err.Error() }

// Unwrap returns the wrapped error.
func (e *ExitError) Unwrap() error { return e.Err }

/** ExitCode returns the DR-0003 exit code for an error from RunScripttool: 0
 * for nil, the code of an ExitError, 2 for an error that wraps ErrUsage, and
 * 70 (internal) for anything not classified.
 *
 * Parameters:
 *   err (error) — the error returned by RunScripttool
 *
 * Returns:
 *   int — the exit code
 *
 * Example:
 *   os.Exit(scripttool.ExitCode(err))
 */
func ExitCode(err error) int {
	var ee *ExitError
	switch {
	case err == nil:
		return ExitOK
	case errors.As(err, &ee):
		return ee.Code
	case errors.Is(err, ErrUsage):
		return ExitUsage
	default:
		return ExitInternal
	}
}

// failure wraps err with an exit code.
func failure(code int, format string, a ...any) error {
	return &ExitError{Code: code, Err: fmt.Errorf(format, a...)}
}

// classifyOpen turns an error from opening a named input or output into a
// code: a missing file is 66, a refused one 77, anything else 74.
func classifyOpen(err error) int {
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return ExitNoInput
	case errors.Is(err, fs.ErrPermission):
		return ExitNoPerm
	default:
		return ExitIO
	}
}

// recordingReader remembers the first read error that is not io.EOF, so a
// conversion that fails can be told apart: a failed read is an I/O error (74),
// content that does not parse is a data error (65).
type recordingReader struct {
	r   io.Reader
	err error
}

func (rr *recordingReader) Read(p []byte) (int, error) {
	n, err := rr.r.Read(p)
	if err != nil && err != io.EOF && rr.err == nil {
		rr.err = err
	}
	return n, err
}

// converter runs one verb. in and out are the streams; inName and outName are
// the file names for the verbs that take or give a FadeIn file by name.
type converter func(in io.Reader, out io.Writer, inName, outName string, alpha bool) error

// fromIn and fromName adapt the conversion functions to converter.
func fromIn(fn func(io.Reader, io.Writer) error) converter {
	return func(in io.Reader, out io.Writer, _, _ string, _ bool) error { return fn(in, out) }
}

func fromName(fn func(string, io.Writer) error) converter {
	return func(_ io.Reader, out io.Writer, inName, _ string, _ bool) error { return fn(inName, out) }
}

func toFadeIn(fn func(io.Reader, string) error) converter {
	return func(in io.Reader, _ io.Writer, _, outName string, _ bool) error { return fn(in, outName) }
}

// verbSpec says what a verb needs: a FadeIn input file by name, or a FadeIn
// output file by name, and what runs it.
type verbSpec struct {
	needsInputName, needsOutputName bool
	run                             converter
}

var verbs = map[string]verbSpec{
	"fadein2fdx":        {needsInputName: true, run: fromName(FadeInToFDX)},
	"fadein2fountain":   {needsInputName: true, run: fromName(FadeInToFountain)},
	"fadein2json":       {needsInputName: true, run: fromName(FadeInToJSON)},
	"fadein2yaml":       {needsInputName: true, run: fromName(FadeInToYAML)},
	"fadein2osf":        {needsInputName: true, run: fromName(FadeInToOSF)},
	"fdx2fadein":        {needsOutputName: true, run: toFadeIn(FdxToFadeIn)},
	"fdx2fountain":      {run: fromIn(FdxToFountain)},
	"fdx2json":          {run: fromIn(FdxToJSON)},
	"fdx2yaml":          {run: fromIn(FdxToYAML)},
	"fdx2osf":           {run: fromIn(FdxToOSF)},
	"fountain2fadein":   {needsOutputName: true, run: toFadeIn(FountainToFadeIn)},
	"fountain2fdx":      {run: fromIn(FountainToFdx)},
	"fountain2json":     {run: fromIn(FountainToJSON)},
	"fountain2yaml":     {run: fromIn(FountainToYAML)},
	"fountain2osf":      {run: fromIn(FountainToOSF)},
	"osf2fadein":        {needsOutputName: true, run: toFadeIn(OSFToFadeIn)},
	"osf2fdx":           {run: fromIn(OSFToFdx)},
	"osf2fountain":      {run: fromIn(OSFToFountain)},
	"osf2json":          {run: fromIn(OSFToJSON)},
	"osf2yaml":          {run: fromIn(OSFToYAML)},
	"fountain2fountain": {run: fromIn(FountainToFountain)},
	"fountainfmt":       {run: fromIn(FountainFmt)},
	"fountain2html":     {run: fromIn(FountainToHTML)},
	"characters": {run: func(in io.Reader, out io.Writer, _, _ string, alpha bool) error {
		//FIXME: add detection of file type, convert to fountain then
		// run report.
		return CharacterList(in, out, alpha)
	}},
}

/** RunScripttool runs one scripttool verb. args is the whole command line,
 * program name first. Every failure comes back as an *ExitError (or an error
 * wrapping ErrUsage) so ExitCode gives the DR-0003 exit code: 2 for a bad
 * command line, 65 for input that cannot be read as the format, 66 for a
 * missing input file, 73 for an output file that cannot be created, 74 for a
 * read or write that failed part way, 77 for a refused file. The output file
 * is created only after the conversion has succeeded.
 *
 * Parameters:
 *   in (*os.File) — standard input
 *   out (*os.File) — standard output
 *   eout (*os.File) — standard error
 *   args ([]string) — the command line, program name first
 *
 * Returns:
 *   error — nil on success
 *
 * Example:
 *   err := scripttool.RunScripttool(os.Stdin, os.Stdout, os.Stderr, os.Args)
 *   os.Exit(scripttool.ExitCode(err))
 */
func RunScripttool(in *os.File, out *os.File, eout *os.File, args []string) error {
	var (
		// Options
		showHelp    bool
		inputFName  string
		outputFName string
		alphaSort   bool
	)

	appName := path.Base(args[0])
	if len(args) == 1 {
		return fmt.Errorf("%w: no verb given, run '%s -help'", ErrUsage, appName)
	}

	verb := strings.ToLower(strings.TrimSpace(args[1]))

	// Standard flags
	flagSet := flag.NewFlagSet(fmt.Sprintf("%s.%s", appName, verb), flag.ContinueOnError)
	flagSet.SetOutput(io.Discard)
	flagSet.BoolVar(&showHelp, "h", false, "display verb help")
	flagSet.BoolVar(&showHelp, "help", false, "display verb help")
	flagSet.StringVar(&inputFName, "i", "", "set input filename")
	flagSet.StringVar(&outputFName, "o", "", "set output filename")
	flagSet.BoolVar(&fountain.ShowNotes, "notes", false, "include notes in output")
	flagSet.BoolVar(&fountain.ShowSynopsis, "synopsis", false, "include synopsis in output")
	flagSet.BoolVar(&fountain.ShowSection, "section", false, "include section headings in output")
	flagSet.IntVar(&fountain.MaxWidth, "width", fountain.MaxWidth, "set width in integers")
	flagSet.BoolVar(&fountain.AsHTMLPage, "html", false, "output as HTML")
	flagSet.BoolVar(&fountain.InlineCSS, "inline-css", false, "include inline CSS")
	flagSet.BoolVar(&fountain.LinkCSS, "link-css", false, "include CSS link")
	flagSet.BoolVar(&fountain.PrettyPrint, "pretty", false, "prety print output")
	flagSet.BoolVar(&alphaSort, "alpha", false, "sort a character list alpha betically versus order of appearence")
	if err := flagSet.Parse(args[2:]); err != nil {
		return fmt.Errorf("%w: %s: %v, run '%s help %s'", ErrUsage, verb, err, appName, verb)
	}
	args = flagSet.Args()

	if verb == "help" {
		return RunHelp(out, appName, args)
	}

	spec, known := verbs[verb]
	if !known {
		return fmt.Errorf("%w: do not understand %q, run '%s help index' for the verbs", ErrUsage, verb, appName)
	}

	if showHelp {
		fmt.Fprintln(out, FmtCliHelp(appName, verb))
		return nil
	}

	if len(args) > 2 {
		return fmt.Errorf("%w: %s takes at most an input and an output name, got surplus %q", ErrUsage, verb, args[2])
	}
	if inputFName == "" && len(args) >= 1 {
		inputFName = args[0]
	}
	if outputFName == "" && len(args) >= 2 {
		outputFName = args[1]
	}
	namedIn := inputFName != "" && inputFName != "-"
	namedOut := outputFName != "" && outputFName != "-"

	if spec.needsInputName && !namedIn {
		return fmt.Errorf("%w: %s needs a FadeIn input file name", ErrUsage, verb)
	}
	if spec.needsOutputName && !namedOut {
		return fmt.Errorf("%w: %s needs a FadeIn output file name", ErrUsage, verb)
	}

	// Open the input, a named file or standard input.
	var src io.Reader = in
	if namedIn {
		f, err := os.Open(inputFName)
		if err != nil {
			return &ExitError{Code: classifyOpen(err), Err: err}
		}
		defer f.Close()
		src = f
	}
	rr := &recordingReader{r: src}

	// Convert into memory; the output file is created only when that worked.
	var buf bytes.Buffer
	if err := spec.run(rr, &buf, inputFName, outputFName, alphaSort); err != nil {
		var pe *fs.PathError
		switch {
		case rr.err != nil:
			return &ExitError{Code: ExitIO, Err: rr.err}
		case spec.needsOutputName && errors.As(err, &pe) && pe.Path == outputFName:
			return &ExitError{Code: classifyOpen(err), Err: err}
		case spec.needsInputName && errors.Is(err, fs.ErrNotExist):
			return &ExitError{Code: ExitNoInput, Err: err}
		default:
			if namedIn {
				return &ExitError{Code: ExitData, Err: fmt.Errorf("%s: %w", inputFName, err)}
			}
			return &ExitError{Code: ExitData, Err: err}
		}
	}
	if spec.needsOutputName {
		return nil // the conversion wrote the named file itself
	}

	if !namedOut {
		if _, err := out.Write(buf.Bytes()); err != nil {
			return &ExitError{Code: ExitIO, Err: err}
		}
		return nil
	}
	f, err := os.Create(outputFName)
	if err != nil {
		code := ExitCantCreate
		if errors.Is(err, fs.ErrPermission) {
			code = ExitNoPerm
		}
		return &ExitError{Code: code, Err: err}
	}
	if _, err := f.Write(buf.Bytes()); err != nil {
		f.Close()
		return &ExitError{Code: ExitIO, Err: err}
	}
	if err := f.Close(); err != nil {
		return &ExitError{Code: ExitIO, Err: err}
	}
	return nil
}
