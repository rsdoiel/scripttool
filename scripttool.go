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
	"fmt"
	"io"
	"io/ioutil"

	// My packages
	"github.com/rsdoiel/fdx"
	"github.com/rsdoiel/fountain"
	"github.com/rsdoiel/osf"
)

const (
	Version = `v0.0.1-dev`
)

// FdxToFountain converts the an input buffer from .fdx to a .fountain format.
func FdxToFountain(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	document, err := fdx.Parse(src)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", document.String())
	return nil
}

// OSFToFountain converts the input buffer from .osf or .fadein to .fountain format.
func OSFToFountain(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := osf.Parse(src)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", document.String())
	return nil
}

// Fountain2Fountain reads a input buffer as .fountain and pretty prints output as .fountain
func FountainToFountain(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := fountain.Parse(src)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", document.String())
	return nil
}

// Fountain2Fdx converts an input buffer in .fountain format to output buffer in .fdx
func FountainToFdx(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := fountain.Parse(src)
	if err != nil {
		return err
	}
	newDoc, err := fountainToFdx(document)
	if err != nil {
	}
	fmt.Fprintf(out, "%s", newDoc.String())
	return nil
}

// FountainToOSF converts an input buffer in .fountain format to output buffer in .osf
func FountainToOSF(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := fountain.Parse(src)
	if err != nil {
		return err
	}
	newDoc, err := fountainToOSF(document)
	if err != nil {
	}
	fmt.Fprintf(out, "%s", newDoc.String())
	return nil
}

// CharacterList lists character in a screenplay
func CharacterList(in io.Reader, out io.Writer) error {
	// What format do we have?
	// Convert to Fountain
	// collect character names
	return fmt.Errorf("CharacterList(in, out) error, not implemented")
}
