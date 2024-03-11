package scripttool

import (
	"flag"
	"fmt"
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

func RunScripttool(in *os.File, out *os.File, eout *os.File, args []string) error {
	var (
		// Options
		showHelp    bool
		inputFName  string
		outputFName string
		alphaSort   bool

		err error
	)

	appName := path.Base(args[0])
	if len(args) == 1 {
		return fmt.Errorf("%s", FmtCliText(HelpText, appName, "", Version))
	}

	verb := strings.ToLower(strings.TrimSpace(args[1]))

	// Standard flags
	flagSet := flag.NewFlagSet(fmt.Sprintf("%s.%s", appName, verb), flag.ExitOnError)
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
	flagSet.Parse(args[2:])
	args = flagSet.Args()

	if showHelp {
		fmt.Fprintln(out, FmtCliHelp(appName, verb))
		return nil
	}

	if inputFName == "" && len(args) >= 1 {
		inputFName = args[0]
	}
	if outputFName == "" && len(args) >= 2 {
		outputFName = args[1]
	}

	// Figure out if we need to open files or use in, out, eout
	if inputFName != "" && inputFName != "-" {
		in, err = os.Open(inputFName)
		if err != nil {
			return err
		}
		defer in.Close()
	}

	if outputFName != "" && outputFName != "-" {
		out, err = os.Create(outputFName)
		if err != nil {
			return err
		}
		defer out.Close()
	}

	switch verb {
	// fadein to something
	case "fadein2fdx":
		if inputFName != "" {
			return FadeInToFDX(inputFName, out)
		}
		return fmt.Errorf("A FadeIn filename required for input")
	case "fadein2fountain":
		if inputFName != "" {
			return FadeInToFountain(inputFName, out)
		}
		return fmt.Errorf("A FadeIn filename required for input")
	case "fadein2json":
		if inputFName != "" {
			return FadeInToJSON(inputFName, out)

		}
		return fmt.Errorf("A FadeIn filename required for input")
	case "fadein2yaml":
		if inputFName != "" {
			return FadeInToYAML(inputFName, out)
		}
	case "fadein2osf":
		if inputFName != "" {
			return FadeInToOSF(inputFName, out)

		}
		return fmt.Errorf("A FadeIn filename required for input")

	// fdx to something
	case "fdx2fadein":
		if outputFName != "" {
			return FdxToFadeIn(in, outputFName)
		}
		return fmt.Errorf("A FadeIn filename required for output")
	case "fdx2fountain":
		return FdxToFountain(in, out)
	case "fdx2json":
		return FdxToJSON(in, out)
	case "fdx2yaml":
		return FdxToYAML(in, out)
	case "fdx2osf":
		return FdxToOSF(in, out)

	// fountain to something
	case "fountain2fadein":
		if outputFName != "" {
			return FountainToFadeIn(in, outputFName)
		}
		return fmt.Errorf("A FadeIn filename required for output")
	case "fountain2fdx":
		return FountainToFdx(in, out)
	case "fountain2json":
		return FountainToJSON(in, out)
	case "fountain2yaml":
		return FountainToYAML(in, out)
	case "fountain2osf":
		return FountainToOSF(in, out)

	// osf to something
	case "osf2fadein":
		if outputFName != "" {
			return OSFToFadeIn(in, outputFName)
		}
		return fmt.Errorf("A FadeIn filename required for output")
	case "osf2fdx":
		return OSFToFdx(in, out)
	case "osf2fountain":
		return OSFToFountain(in, out)
	case "osf2json":
		return OSFToJSON(in, out)
	case "osf2yaml":
		return OSFToYAML(in, out)

	// Utility fountain functions
	case "fountain2fountain":
		return FountainToFountain(in, out)
	case "fountainfmt":
		return FountainFmt(in, out)
	case "fountain2html":
		return FountainToHTML(in, out)
	case "characters":
		//FIXME: add detection of file type, convert to fountain then
		// run report.
		return CharacterList(in, out, alphaSort)

	// Help system
	case "help":
		fmt.Fprintln(out, FmtCliText(HelpText, appName, verb, Version))
		return nil

	// If we got this far we have an unknown verb
	default:
		return fmt.Errorf("do not understand %q in %q", verb, strings.Join(args, " "))
	}
	return nil
}
