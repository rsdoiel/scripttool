//
// fountain2json parses a fountain file and returns a JSON representation.
//
// fountain is a package encoding/decoding fountain formatted screenplays.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
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
	"fmt"
	"io/ioutil"
	"os"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"

	// My packages
	"github.com/rsdoiel/fountain"
)

var (
	description = `%s is a command line program that reads an fountain document and returns a JSON representation of it.

`

	examples = `Render *screenplay.fountain* as *screenplay.json*.

    %s -i screenplay.fountain -o screenplay.json

Or alternatively

    cat screenplay.fountain | %s > screenplay.json
`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	generateMarkdown bool
	generateManPage  bool
	newLine          bool
	quiet            bool
	inputFName       string
	outputFName      string

	// App Option
	width       int
	prettyPrint bool
)

func main() {
	app := cli.NewCli(fountain.Version)
	appName := app.AppName()

	// Add Help
	app.AddHelp("description", []byte(fmt.Sprintf(description, appName)))
	app.AddHelp("examples", []byte(fmt.Sprintf(examples, appName, appName)))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate Markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")
	app.BoolVar(&newLine, "nl,newline", true, "add a trailing newline")
	app.BoolVar(&quiet, "quiet", false, "suppress error messages")
	app.StringVar(&inputFName, "i,input", "", "set the input filename")
	app.StringVar(&outputFName, "o,output", "", "set the output filename")

	// App Option
	app.BoolVar(&prettyPrint, "p,pretty", false, "pretty print the JSON output")
	app.IntVar(&width, "w,width", 65, "set the width for the text")

	// Parse environment and options
	app.Parse()
	args := app.Args()

	// Setup IO
	var err error
	app.Eout = os.Stderr
	app.In, err = cli.Open(inputFName, os.Stdin)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(inputFName, app.In)
	app.Out, err = cli.Create(outputFName, os.Stdout)
	cli.ExitOnError(app.Eout, err, quiet)
	defer cli.CloseFile(outputFName, app.Out)

	// Process options
	if generateMarkdown {
		app.GenerateMarkdown(app.Out)
		os.Exit(0)
	}
	if generateManPage {
		app.GenerateManPage(app.Out)
		os.Exit(0)
	}
	if showHelp {
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

	// ReadAll of input
	src, err := ioutil.ReadAll(app.In)
	cli.ExitOnError(app.Eout, err, quiet)
	// Parse input
	screenplay, err := fountain.Parse(src)
	cli.OnError(app.Eout, err, quiet)

	fountain.PrettyPrint = prettyPrint
	src, err = screenplay.ToJSON()
	cli.OnError(app.Eout, err, quiet)

	fmt.Fprintf(app.Out, "%s", src)
	if newLine {
		fmt.Fprintln(app.Out)
	}
}
