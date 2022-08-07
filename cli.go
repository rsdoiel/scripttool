package scripttool

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	// My packages
	"github.com/rsdoiel/scripttool/fountain"
)

func FmtCliText(s string, appName string, verb string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "{app_name}", appName), "{verb}", verb), "{version}", version)
}

func FmtCliHelp(appName string, verb string) string {
	return fmt.Sprintf("%s %s is not documented", appName, verb)
}

func RunScripttool(in *os.File, out *os.File, eout *os.File, args []string) error {
	var (
		// Standard Options
		showHelp    bool
		inputFName  string
		outputFName string

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
	case "fdx2fountain":
		return FdxToFountain(in, out)
	case "osf2fountain":
		return OSFToFountain(in, out)
	case "fountain2fdx":
		return FountainToFdx(in, out)
	case "fountain2osf":
		return FountainToOSF(in, out)
	case "characters":
		return CharacterList(in, out)
	case "fadein2fountain":
		if inputFName != "" {
			return FadeInToFountain(inputFName, out)

		}
		return fmt.Errorf("A FadeIn filename must for input")
	case "fountain2fadein":
		if outputFName != "" {
			return FountainToFadeIn(in, outputFName)
		}
		return fmt.Errorf("A FadeIn filename must for output")
	case "fountainfmt":
		return FountainFmt(in, out)
	case "fountain2html":
		return FountainToHTML(in, out)
	case "fountain2json":
		return FountainToJSON(in, out)
	case "help":
		fmt.Fprintln(out, FmtCliText(HelpText, appName, verb, Version))
		return nil
	default:
		return fmt.Errorf("do not understand %q in %q", verb, strings.Join(args, " "))

	}
	return nil
}
