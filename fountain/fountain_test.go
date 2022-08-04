//
// fountain is a package encoding/decoding fountain formatted screenplays
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
package fountain

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func assertOK(t *testing.T, err error, msg string) {
	if err != nil {
		t.Errorf("%s, %s", err, msg)
	}
}

func TestDialogue(t *testing.T) {
	src := []byte(`
INT. LAB - DAY

CHARLIE
Bring that to me.

(turns to Wilma)
And you fetch the tongs.

(Wilma turns and leaves the room)
`)

	expected := []int{
		SceneHeadingType,  // INT. LAB - DAY
		EmptyType,         //
		CharacterType,     // CHARLIE
		DialogueType,      // Bring that to me.
		EmptyType,         //
		ParentheticalType, // (turns to Wilma)
		DialogueType,      // And you fetch the tongs.
		EmptyType,         //
		ParentheticalType, // (Wilma turns and leaves the room)
	}

	doc, err := Parse(src)
	assertOK(t, err, "Parse(src)")

	if doc == nil || doc.Elements == nil {
		t.Errorf("Couldn't create doc.Elements from Parse()")
		t.FailNow()
	}
	for i := 0; i < len(doc.Elements); i++ {
		if doc.Elements[i].Type != expected[i] {
			t.Errorf("expected %q, got %q for %q", typeName(expected[i]), typeName(doc.Elements[i].Type), doc.Elements[i].Content)
			t.FailNow()
		}
	}

}

func TestTypes(t *testing.T) {
	src, err := ioutil.ReadFile(path.Join("testdata", "sample-01.fountain"))
	assertOK(t, err, "ReadFile(testdata/sample-01.fountain)")

	doc, err := Parse(src)
	assertOK(t, err, "Parse(src)")

	/*
	   !FADE IN:

	   EXT. LIBRARY - DAY

	   A PROGRAMMER typing at an old laptop

	   PROGRAMMER
	   (excited)
	   Eureka!

	   > FADE TO BLACK.
	*/
	expected := []int{
		//NOTE: This "FADE IN:" should be action type because of the "!"
		ActionType,        // "!FADE IN:"
		EmptyType,         // ""
		SceneHeadingType,  // "EXT. LIBRARY - DAY"
		EmptyType,         // ""
		ActionType,        // "A PROGRAMMER typing at an old laptop"
		EmptyType,         // ""
		CharacterType,     // "PROGRAMMER"
		ParentheticalType, // "(excited)"
		DialogueType,      // "Eureka!"
		EmptyType,         // ""
		TransitionType,    // "> FADE TO BLACK."
		EmptyType,         // ""
	}

	if len(doc.Elements) > len(expected) {
		t.Errorf("Got more elements than expected, %+v", doc.Elements)
		for _, elem := range doc.Elements {
			fmt.Fprintf(os.Stderr, "%s %q\n", typeName(elem.Type), elem.Content)
		}
		t.FailNow()
	}
	for i := 0; i < len(doc.Elements); i++ {
		if doc.Elements[i].Type != expected[i] {
			t.Errorf("(%d) expected %q, got %q for %q", i, typeName(expected[i]), typeName(doc.Elements[i].Type), doc.Elements[i].Content)
			t.FailNow()
		}
	}

}

func TestSamples(t *testing.T) {
	files := []string{
		"sample-01.fountain",
		"sample-02.fountain",
		"sample-03.fountain",
		"sample-04.fountain",
		"sample-05.fountain",
		"sample-06.fountain",
		"sample-07.fountain",
	}

	for i, fName := range files {
		screenplay, err := ParseFile(path.Join("testdata", fName))
		if err != nil {
			t.Errorf("(%d) Should be able to read and parse %s, %s", i, fName, err)
			t.FailNow()
		}
		if screenplay.Elements == nil {
			t.Errorf("(%d) expected elements, got nil for %s", i, fName)
		}
		//FIXME: Check to see if the parse sequence is correct
	}
}

func TestMain(m *testing.M) {
	// Setup everything, process flags, etc.
	os.Exit(m.Run())
}
