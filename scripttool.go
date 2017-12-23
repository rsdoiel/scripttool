//
// scripttool is a package focused on converting to/from different
// file formats used in working with scripts,screenplays and other
// creative writing.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
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
package scripttool

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	// My packages
	"github.com/rsdoiel/fdx"
)

const (
	Version = `v0.0.0-dev`
)

// Fdx2Fountain converts the an input buffer from .fdx to
// a .fountain formatted output buffer.
func Fdx2Fountain(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	screenplay := new(fdx.FinalDraft)
	err = xml.Unmarshal(src, &screenplay)
	if err != nil {
		return err
	}
	// FIXME: See if we have a title page
	if screenplay.TitlePage != nil {
		fmt.Println("DEBUG we have a title page!")
		m := screenplay.TitlePageAsMap()
		if val, ok := m["Title"]; ok == true {
			fmt.Fprintf(out, "Title:\n")
			for _, line := range strings.Split(val, "\n") {
				fmt.Fprintln(out, "  %s\n", line)
			}
		}
		if val, ok := m["Credit"]; ok == true {
			fmt.Fprintf(out, "Credit: ")
			for _, line := range strings.Split(val, "\n") {
				fmt.Fprintln(out, "  %s\n", line)
			}
		}
		if val, ok := m["Author"]; ok == true {
			fmt.Fprintf(out, "Author: ")
			for _, line := range strings.Split(val, "\n") {
				fmt.Fprintln(out, "  %s\n", line)
			}
		}
		if val, ok := m["Source"]; ok == true {
			fmt.Fprintf(out, "Source: ")
			for _, line := range strings.Split(val, "\n") {
				fmt.Fprintln(out, "  %s\n", line)
			}
		}
		if val, ok := m["Draft date"]; ok == true {
			fmt.Fprintf(out, "Draft date: ")
			for _, line := range strings.Split(val, "\n") {
				fmt.Fprintln(out, "  %s\n", line)
			}
		}
		if val, ok := m["Contact"]; ok == true {
			fmt.Fprintf(out, "Contact:\n")
			for _, line := range strings.Split(val, "\n") {
				fmt.Fprintln(out, "  %s\n", line)
			}
		}
		if len(m) > 0 {
			// Add the implicit page marker as two blank lines.
			fmt.Fprintln(out, "\n\n")
		}
	}
	if screenplay.Content != nil {
		fmt.Println("DEBUG we have screenplay content!")
	}
	return nil
}

func Fountain2Fdx(in io.Reader, out io.Writer) error {
	return fmt.Errorf("Fountain2Fdx() not implemented")
}

func Characters(in io.Reader, out io.Writer) error {
	fmt.Println("DEBUG in Characters, returning error!")
	return fmt.Errorf("Characters() not implemented!")
}
