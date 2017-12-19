//
// scripttool is a package focused on converting to/from different
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
package scripttool

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var (
	expectedDocs map[string][]byte
)

func TestConversion(t *testing.T) {
	for fname, buf := range expectedDocs {
		fmt.Printf("%s -> %s\n", fname, buf)
	}
}

func TestMain(m *testing.M) {
	expectedDocs = make(map[string][]byte)
	fname := "testdata/testplay-01.fountain"
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	expectedDocs[fname] = buf

	fname = "testdata/testplay-01.fdx"
	buf, err = ioutil.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	expectedDocs[fname] = buf

	// Setup everything, process flags, etc.
	os.Exit(m.Run())
}
