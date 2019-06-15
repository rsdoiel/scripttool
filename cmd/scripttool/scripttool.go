// scripttool is a tool for working with screenplay file formats.
//
// BSD 2-Clause License
//
// Copyright (c) 2019, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
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
//
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"

	// My packages
	"github.com/rsdoiel/scripttool"
)

var (
	synopsis = `_scripttool_ a program for converting between screenplay formats (e.g. .fdx, .fadein, .fountain)`

	description = `_scripttool_ converts screen play file formats. Supported formats include FileDraft's XML format, FadeIn's zipped XML format, Fountain formatted plain text as the Open Screenplay Format XML documents. The command line program is based on a Go package also called scripttool. The Go package can be compiled to a shared library and integrated with Python via the ctypes package.  `

	examples = `
Converting *screenplay.fdx* to *screenplay.fountain* (2 examples)

` + "```" + `
    scripttool fdx2fountain screenplay.fdx screenplay.fountain
    scripttool -i screenplay.fdx -o screenplay.fountain fdx2fountain
` + "```" + `

Converting *screenplay.fountain* to *screenplay.fdx* (2 examples)

` + "```" + `
    scripttool fountain2fdx screenplay.fountain screenplay.fdx
    scripttool -i screenplay.fountain -o screenplay.fdx fountain2fdx
` + "```" + `

Listing characters in *screenplay.fountain* or in *screenplay.fdx*.
(2 examples each)

` + "```" + `
    scripttool characters screenplay.fountain
    scripttool -i screenplay.fountain characters
    scripttool characters screenplay.fdx
    scripttool -i screenplay.fdx characters
` + "```" + `
`

	license = `
%s %s

BSD 2-Clause License

Copyright (c) 2019, R. S. Doiel
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	generateMarkdown bool
	generateManPage  bool
	inputFName       string
	outputFName      string
	quiet            bool
	newLine          bool
)

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

func doFountainToFadeIn(in io.Reader, out io.Writer, eout io.Writer, args []string, flagSet *flag.FlagSet) int {
	var err error

	if inputFName == "" && len(args) > 0 {
		inputFName = args[0]
		in, err = cli.Open(inputFName, os.Stdin)
		cli.ExitOnError(os.Stderr, err, quiet)
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
	return onError(eout, scripttool.FountainToFadeIn(in, outputFName))
}

func main() {
	app := cli.NewCli(scripttool.Version)
	appName := app.AppName()
	app.SetParams("VERB", "[VERB PARAMETERS]")

	// Standard flags
	app.AddHelp("synopsis", []byte(synopsis))
	app.AddHelp("description", []byte(description))
	app.AddHelp("examples", []byte(examples))
	app.AddHelp("license", []byte(fmt.Sprintf(license, appName, scripttool.Version)))
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showExamples, "examples", false, "display examples")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate Markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.StringVar(&inputFName, "i,input", "", "set input filename")
	app.StringVar(&outputFName, "o,output", "", "set output filename")
	app.BoolVar(&newLine, "nl,newline", true, "add a trailing newline to output")

	// Add Verbs
	app.VerbsRequired = true
	vFdxToFountain := app.NewVerb("fdx2fountain", "convert .fdx to .fountain", doFdxToFountain)
	vFdxToFountain.SetParams("[INPUT_FILENAME]", "[OUTPUT_FILENAME]")

	vOSFToFountain := app.NewVerb("osf2fountain", "convert .osf to .fountain", doOSFToFountain)
	vOSFToFountain.SetParams("[INPUT_FILENAME]", "[OUTPUT_FILENAME]")

	vFountainToFdx := app.NewVerb("fountain2fdx", "convert .fountain to .fdx", doFountainToFdx)
	vFountainToFdx.SetParams("[INPUT_FILENAME]", "[OUTPUT_FILENAME]")

	vFountainToOSF := app.NewVerb("fountain2osf", "convert .fountain to .osf", doFountainToOSF)
	vFountainToOSF.SetParams("[INPUT_FILENAME]", "[OUTPUT_FILENAME]")

	vCharacters := app.NewVerb("characters", "list characters in a screenplay", doCharacters)
	vCharacters.SetParams("[INPUT_FILENAME]")

	vFadeInToFountain := app.NewVerb("fadein2fountain", "convert .fadein to .fountain", doFadeInToFountain)
	vFadeInToFountain.SetParams("INPUT_FILENAME", "[OUTPUT_FILENAME]")

	vFountainToFadeIn := app.NewVerb("fountain2fadein", "convert .fountain to .fadein", doFountainToFadeIn)
	vFountainToFadeIn.SetParams("INPUT_FILENAME", "OUTPUT_FILENAME")

	// Parse environment and command line
	if err := app.Parse(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	args := app.Args()

	if generateMarkdown {
		app.GenerateMarkdown(os.Stdout)
		os.Exit(0)
	}
	if generateManPage {
		app.GenerateManPage(os.Stdout)
		os.Exit(0)
	}
	if showHelp || showExamples {
		if len(args) > 0 {
			fmt.Fprintln(app.Out, app.Help(args...))
		} else {
			app.Usage(app.Out)
		}
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintln(app.Out, app.License())
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintln(app.Out, app.Version())
		os.Exit(0)
	}

	// Setup IO
	var err error

	app.Eout = os.Stderr

	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)

	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)

	// Run our program
	exitCode := app.Run(args)

	// Add a trailing newLine if set
	if newLine {
		fmt.Fprintln(app.Out, "")
	}
	os.Exit(exitCode)
}
