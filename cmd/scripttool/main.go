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
	"os"
	"path"

	// My packages
	"github.com/rsdoiel/scripttool"
)

func main() {
	appName := path.Base(os.Args[0])
	var showHelp, showLicense, showVersion bool

	// Only the options before the verb belong to main; the verb parses its own.
	fset := flag.NewFlagSet(appName, flag.ContinueOnError)
	fset.SetOutput(io.Discard)
	fset.BoolVar(&showHelp, "h", false, "display help")
	fset.BoolVar(&showHelp, "help", false, "display help")
	fset.BoolVar(&showLicense, "l", false, "display license")
	fset.BoolVar(&showLicense, "license", false, "display license")
	fset.BoolVar(&showVersion, "v", false, "display version")
	fset.BoolVar(&showVersion, "version", false, "display version")
	if err := fset.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v, run '%s -help'\n", appName, err, appName)
		os.Exit(scripttool.ExitUsage)
	}
	if showHelp {
		fmt.Println(scripttool.FmtCliText(scripttool.HelpText, appName, "", scripttool.Version))
		os.Exit(0)
	}
	if showLicense {
		fmt.Println(scripttool.FmtCliText(scripttool.LicenseText, appName, "", scripttool.Version))
		os.Exit(0)
	}
	if showVersion {
		fmt.Printf("%s %s\n", appName, scripttool.Version)
		os.Exit(0)
	}

	// The verb and everything after it, as RunScripttool takes them.
	args := append([]string{os.Args[0]}, fset.Args()...)
	if err := scripttool.RunScripttool(os.Stdin, os.Stdout, os.Stderr, args); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", appName, err)
		os.Exit(scripttool.ExitCode(err))
	}
}
