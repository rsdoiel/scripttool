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
	fname := "testdata/testplay-01.fountain"
	src, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	fdx := new(FinalDraft)
	if err := xml.Unmarshal(src, &fdx); err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	} else {
		fmt.Fprintf("DEBUG fdx: %+v\n", fdx)
	}
}

func TestMain(m *testing.M) {
	// Setup everything, process flags, etc.
	os.Exit(m.Run())
}
