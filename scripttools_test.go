//
// scripttools is a package focused on converting to/from different
// file formats used in creative writing (e.g. screenplays) and
// providing a common tool for analysis for supported formats.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// BSD 3-Clause License
//
// Copyright (c) 2017, R. S. Doiel
// All rights reserved.
//
//
package scripttools

import (
	"os"
	"testing"
)

var (
	fountainDocs map[string][]byte
	fdxDocs      map[string][]byte
	pdfDocs      map[string][]byte
)

func TestFountain2FDX(t *testing.T) {
}

func TestFDX2Fountain(t *testing.T) {
}

func TestMain(m *testing.M) {
	//NOTE: fountain.io website has fountian, fdx and pdf examples
	// I can also create some of my own from using Trelby as a reference
	// and generating some test cases that way.

	//FIXME: get the filenames in testdata
	//FIXME: read the files to fountainTest, fdxTest, pdfTest

	// Setup everything, process flags, etc.
	os.Exit(m.Run())
}
