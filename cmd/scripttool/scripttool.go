// scripttool is a tool for working with screenplay file formats.
//
// BSD 2-Clause License
//
// Copyright (c) 2017, R. S. Doiel
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
	"fmt"
	"io"
	"log"
	"os"

	// Caltech Library packages
	"github.com/caltechlibrary/cli"

	// My packages
	//"github.com/rsdoiel/fdx"
	"github.com/rsdoiel/scripttool"
)

var (
	description = `
scripttool a command line program for working with screenplay file formats
(e.g. fdx, fountain).
`

	examples = `
Converting *screenplay.fdx* to *screenplay.fountain* (2 examples) 

    scripttool fdx2fountain screenplay.fdx screenplay.fountain
	scripttool -i screenplay.fdx -o screenplay.fountain fdx2fountain

Converting *screenplay.fountain* to *screenplay.fdx* (2 examples)

    scripttool fountain2fdx screenplay.fountain screenplay.fdx
	scripttool -i screenplay.fountain -o screenplay.fdx fountain2fdx

Listing characters in *screenplay.fountain* or in *screenplay.fdx*.
(2 examples each)

	scripttool characters screenplay.fountain
	scripttool -i screenplay.fountain characters
	scripttool characters screenplay.fdx
	scripttool -i screenplay.fdx characters
`

	license = `
%s %s

BSD 2-Clause License

Copyright (c) 2017, R. S. Doiel
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
	showHelp             bool
	showLicense          bool
	showVersion          bool
	showExamples         bool
	generateMarkdownDocs bool
	inputFName           string
	outputFName          string
	quiet                bool
	newLine              bool
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

func doFdxToFountain(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	return onError(eout, scripttool.FdxToFountain(in, out))
}

func doFountainToFdx(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	return onError(eout, scripttool.FountainToFdx(in, out))
}

func doOSFToFountain(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	return onError(eout, scripttool.OSFToFountain(in, out))
}

func doFountainToOSF(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	return onError(eout, scripttool.FountainToOSF(in, out))
}

func doCharacters(in io.Reader, out io.Writer, eout io.Writer, args []string) int {
	return onError(eout, scripttool.CharacterList(in, out))
}

func main() {
	app := cli.NewCli(scripttool.Version)
	appName := app.AppName()
	app.AddParams("ACTION", "[ACTION PARAMETERS]")
	app.AddHelp("description", []byte(description))
	app.AddHelp("examples", []byte(examples))
	app.AddHelp("license", []byte(fmt.Sprintf(license, appName, scripttool.Version)))
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showExamples, "examples", false, "display examples")
	app.BoolVar(&generateMarkdownDocs, "generate-markdown-docs", false, "generate Markdown documentation")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.StringVar(&inputFName, "i,input", "", "set input filename")
	app.StringVar(&outputFName, "o,output", "", "set output filename")
	app.BoolVar(&newLine, "nl,newline", true, "add a trailing newline to output")

	app.AddAction("fdx2fountain", doFdxToFountain, "convert .fdx to .fountain")
	app.AddAction("osf2fountain", doOSFToFountain, "convert .osf to .fountain")
	app.AddAction("fountain2fdx", doFountainToFdx, "convert .fountain to .fdx")
	app.AddAction("fountain2osf", doFountainToOSF, "convert .fountain to .osf")
	app.AddAction("characters", doCharacters, "list characters in a screenplay")

	app.AddVerb("fadein2fountain", "convert .fadein to .fountain")
	app.AddVerb("fountain2fadein", "convert .fountain to .fadein")

	// Parse environment and command line
	if err := app.Parse(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	args := app.Args()

	if generateMarkdownDocs {
		app.GenerateMarkdownDocs(os.Stdout)
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

	// Map verb phrases to standard options
	var (
		verb             string
		fadeInToFountain bool
		fountainToFadeIn bool
	)

	switch app.Verb(args) {
	case "fdx2fountain":
		if inputFName == "" && len(args) > 1 {
			verb, args = cli.ShiftArg(args)
			inputFName, args = cli.ShiftArg(args)
			args = cli.UnshiftArg(verb, args)
		}
		if outputFName == "" && len(args) > 1 {
			verb, args = cli.ShiftArg(args)
			outputFName, args = cli.ShiftArg(args)
			args = cli.UnshiftArg(verb, args)
		}
	case "fountain2fdx":
		if inputFName == "" && len(args) > 1 {
			verb, args = cli.ShiftArg(args)
			inputFName, args = cli.ShiftArg(args)
			args = cli.UnshiftArg(verb, args)
		}
		if outputFName == "" && len(args) > 1 {
			verb, args = cli.ShiftArg(args)
			outputFName, args = cli.ShiftArg(args)
			args = cli.UnshiftArg(verb, args)
		}
	case "fadein2fountain":
		fadeInToFountain = true
		if inputFName == "" && len(args) > 1 {
			verb, args = cli.ShiftArg(args)
			inputFName, args = cli.ShiftArg(args)
			args = cli.UnshiftArg(verb, args)
		}
		if outputFName == "" && len(args) > 1 {
			verb, args = cli.ShiftArg(args)
			outputFName, args = cli.ShiftArg(args)
			args = cli.UnshiftArg(verb, args)
		}
		if inputFName == "" {
			log.Fatal("Missing FadeIn input filename")
		}
	case "fountain2fadein":
		fountainToFadeIn = true
		if inputFName == "" && len(args) > 1 {
			verb, args = cli.ShiftArg(args)
			inputFName, args = cli.ShiftArg(args)
			args = cli.UnshiftArg(verb, args)
		}
		if outputFName == "" && len(args) > 1 {
			verb, args = cli.ShiftArg(args)
			outputFName, args = cli.ShiftArg(args)
			args = cli.UnshiftArg(verb, args)
		}
		if outputFName == "" {
			log.Fatal("Missing FadeIn output filename")
		}
	case "osf2fountain":
		if inputFName == "" && len(args) > 1 {
			verb, args = cli.ShiftArg(args)
			inputFName, args = cli.ShiftArg(args)
			args = cli.UnshiftArg(verb, args)
		}
		if outputFName == "" && len(args) > 1 {
			verb, args = cli.ShiftArg(args)
			outputFName, args = cli.ShiftArg(args)
			args = cli.UnshiftArg(verb, args)
		}
	case "fountain2osf":
		if inputFName == "" && len(args) > 1 {
			verb, args = cli.ShiftArg(args)
			inputFName, args = cli.ShiftArg(args)
			args = cli.UnshiftArg(verb, args)
		}
		if outputFName == "" && len(args) > 1 {
			verb, args = cli.ShiftArg(args)
			outputFName, args = cli.ShiftArg(args)
			args = cli.UnshiftArg(verb, args)
		}
	case "characters":
		if inputFName == "" && len(args) > 1 {
			verb, args = cli.ShiftArg(args)
			inputFName, args = cli.ShiftArg(args)
			args = cli.UnshiftArg(verb, args)
		}
	}

	// Setup IO
	var err error

	app.Eout = os.Stderr

	if fadeInToFountain {
		app.Out, err = cli.Create(outputFName, os.Stdout)
		cli.ExitOnError(app.Eout, err, quiet)
		err = scripttool.FadeInToFountain(inputFName, app.Out)
		cli.ExitOnError(app.Eout, err, quiet)
		if newLine {
			fmt.Fprintln(app.Out, "")
		}
		os.Exit(0)
	}
	if fountainToFadeIn {
		app.In, err = cli.Open(inputFName, os.Stdin)
		cli.ExitOnError(app.Eout, err, quiet)
		err = scripttool.FountainToFadeIn(app.In, outputFName)
		cli.ExitOnError(app.Eout, err, quiet)
		os.Exit(0)
	}

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
