// scripttool is a tool for working with screenplay file formats.
//
// # BSD 2-Clause License
//
// Copyright (c) 2021, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
//   - Redistributions of source code must retain the above copyright notice, this
//     list of conditions and the following disclaimer.
//
//   - Redistributions in binary form must reproduce the above copyright notice,
//     this list of conditions and the following disclaimer in the documentation
//     and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	// My packages
	"github.com/rsdoiel/scripttool"
	"github.com/rsdoiel/scripttool/fountain"
)

var (
	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool
	inputFName  string
	outputFName string
	quiet       bool
	newLine     bool

	// App options
	showNotes     bool
	showSynopsis  bool
	showSections  bool
	width         int
	sectionHeight string
	sectionWidth  string
	asHTMLPage    bool
	inlineCSS     bool
	linkCSS       bool
	includeCSS    string
	prettyPrint   bool
)

func displayText(s string, appName string, verb string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "{app_name}", appName), "{verb}", verb), "{version}", version)
}
func exitOnError(eout io.Writer, err error, exitCode int) {
	if err != nil {
		fmt.Fprintf(eout, "%s\n", err)
		os.Exit(exitCode)
	}
}

func onError(eout io.Writer, err error) int {
	if err != nil {
		if quiet == false {
			fmt.Fprintf(eout, "%s\n", err)
		}
		return 1
	}
	return 0
}

func doFdxToFountain(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	return onError(eout, scripttool.FdxToFountain(in, out))
}

func doFountainToFdx(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	return onError(eout, scripttool.FountainToFdx(in, out))
}

func doOSFToFountain(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	return onError(eout, scripttool.OSFToFountain(in, out))
}

func doFountainToOSF(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	return onError(eout, scripttool.FountainToOSF(in, out))
}

func doCharacters(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	return onError(eout, scripttool.CharacterList(in, out))
}

func doFadeInToFountain(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	if inputFName == "" && len(args) > 0 {
		inputFName = args[0]
	}
	if inputFName == "" {
		fmt.Fprintf(eout, "Missing input FadeIn filename\n")
		return 1
	}
	return onError(eout, scripttool.FadeInToFountain(inputFName, out))
}

func doFountainToFadeIn(in *os.File, out *os.File, eout *os.File, args []string, flagSet *flag.FlagSet) int {
	var err error

	if inputFName == "" && len(args) > 0 {
		inputFName = args[0]
		in, err = os.Open(inputFName)
		exitOnError(eout, err, 1)
		defer in.Close()
	}
	if outputFName == "" && len(args) > 1 {
		outputFName = args[1]
	}
	if inputFName == "" {
		fmt.Fprintf(eout, "Missing input fountain filename\n")
		return 1
	}
	if outputFName == "" {
		fmt.Fprintf(eout, "Missing output FadeIn filename\n")
		return 1
	}
	err = scripttool.FountainToFadeIn(in, outputFName)
	exitOnError(eout, err, 1)
	return 0
}

func doFountainToHTML(in *os.File, out *os.File, eout *os.File, args []string, flagSet *flag.FlagSet) int {
	var err error
	flagSet.Parse(args)
	args = flagSet.Args()
	if inputFName == "" && len(args) > 0 {
		inputFName = args[0]
		in, err = os.Open(inputFName)
		exitOnError(eout, err, 1)
		defer in.Close()
	}
	if outputFName == "" && len(args) > 1 {
		outputFName = args[1]
		out, err = os.Create(outputFName)
		exitOnError(eout, err, 1)
		defer out.Close()
	}
	// Override defaults
	fountain.AsHTMLPage = asHTMLPage
	fountain.InlineCSS = inlineCSS
	fountain.LinkCSS = linkCSS
	if includeCSS != "" {
		fountain.CSS = includeCSS
		fmt.Fprintf(eout, "DEBUG include CSS is now %q\n", fountain.CSS)
	}
	err = scripttool.FountainToHTML(in, out)
	exitOnError(eout, err, 1)
	return 0
}

func doFountainToJSON(in *os.File, out *os.File, eout *os.File, args []string, flagSet *flag.FlagSet) int {
	var err error
	flagSet.Parse(args)
	args = flagSet.Args()
	if inputFName == "" && len(args) > 0 {
		inputFName = args[0]
		in, err = os.Open(inputFName)
		exitOnError(os.Stderr, err, 1)
		defer in.Close()
	}
	if outputFName == "" && len(args) > 1 {
		outputFName = args[1]
		out, err = os.Create(outputFName)
		exitOnError(os.Stderr, err, 1)
		defer out.Close()
	}
	// Override defaults
	fountain.PrettyPrint = prettyPrint
	err = scripttool.FountainToJSON(in, out)
	exitOnError(eout, err, 1)
	return 0
}

func doFountainFmt(in *os.File, out *os.File, eout *os.File, args []string, flagSet *flag.FlagSet) int {
	var err error
	flagSet.Parse(args)
	args = flagSet.Args()
	if inputFName == "" && len(args) > 0 {
		inputFName = args[0]
		in, err = os.Open(inputFName)
		exitOnError(eout, err, 1)
		defer in.Close()
	}
	if outputFName == "" && len(args) > 1 {
		outputFName = args[1]
		out, err = os.Create(outputFName)
		exitOnError(eout, err, 1)
		defer in.Close()
	}
	// Override defaults
	fountain.MaxWidth = width
	fountain.ShowNotes = showNotes
	fountain.ShowSection = showSections
	fountain.ShowSynopsis = showSynopsis
	err = scripttool.FountainFmt(in, out)
	exitOnError(eout, err, 1)
	return 0
}

func main() {
	appName := path.Base(os.Args[0])
	if len(os.Args) == 1 {
		displayText(helpText, appName, "", scripttool.Version)
		os.Exit(1)
	}

	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.Parse()
	if showHelp {
		displayText(helpText, appName, "", scripttool.Version)
		os.Exit(0)
	}
	if showLicense {
		displayText(licenseText, appName, "", scripttool.Version)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(os.Stdout, "%s %s\n", appName, scripttool.Version)
		os.Exit(0)
	}

	in := os.Stdin
	out := os.Stdout
	eout := os.Stderr
	exitCode := 0

	verb := strings.ToLower(strings.TrimSpace(os.Args[1]))
	// Standard flags
	flagSet := flag.NewFlagSet(fmt.Sprintf("%s:%s", appName, verb), flag.ExitOnError)
	flagSet.StringVar(&inputFName, "i", "", "set input filename")
	flagSet.StringVar(&outputFName, "o", "", "set output filename")
	flagSet.BoolVar(&showNotes, "notes", false, "include notes in output")
	flagSet.BoolVar(&showSynopsis, "synopsis", false, "include synopsis in output")
	flagSet.BoolVar(&showSections, "section", false, "include section headings in output")
	flagSet.IntVar(&width, "width", fountain.MaxWidth, "set width in integers")
	flagSet.StringVar(&sectionHeight, "height", "", "section height")
	flagSet.StringVar(&sectionWidth, "width", "", "section width")
	flagSet.BoolVar(&asHTMLPage, "page", false, "output full HTML page")
	flagSet.BoolVar(&inlineCSS, "inline-css", false, "include inline CSS")
	flagSet.BoolVar(&linkCSS, "link-css", false, "include CSS link")
	flagSet.StringVar(&includeCSS, "css", "", "include custom CSS")
	flagSet.BoolVar(&prettyPrint, "pretty", false, "prety print output")
	flagSet.Parse(os.Args[1:])

	switch verb {
	case "fdx2fountain":
		exitCode = doFdxToFountain(in, out, eout, os.Args[1:], flagSet)
	case "osf2fountain":
		exitCode = doOSFToFountain(in, out, eout, os.Args[1:], flagSet)
	case "fountain2fdx":
		exitCode = doFountainToFdx(in, out, eout, os.Args[1:], flagSet)
	case "fountain2osf":
		exitCode = doFountainToOSF(in, out, eout, os.Args[1:], flagSet)
	case "characters":
		exitCode = doCharacters(in, out, eout, os.Args[1:], flagSet)
	case "fadein2fountain":
		exitCode = doFadeInToFountain(in, out, eout, os.Args[1:], flagSet)
	case "fountain2fadein":
		exitCode = doFountainToFadeIn(in, out, eout, os.Args[1:], flagSet)
	case "fountainfmt":
		exitCode = doFountainFmt(in, out, eout, os.Args[1:], flagSet)
	case "fountain2html":
		exitCode = doFountainToHTML(in, out, eout, os.Args[1:], flagSet)
	case "fountain2json":
		exitCode = doFountainToJSON(in, out, eout, os.Args[1:], flagSet)
	case "help":
		displayText(helpText, appName, verb, scripttool.Version)
	default:
		fmt.Fprintf(os.Stderr, "Donnot under standand %q\n", strings.Join(os.Args, " "))
		os.Exit(1)
	}
	os.Exit(exitCode)
}
