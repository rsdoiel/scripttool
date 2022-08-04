//
// fdx is a package encoding/decoding fdx formatted XML files.
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
package fdx

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

var (
	expectedDocs map[string][]byte
)

func testFdxFile(t *testing.T, fname string) {
	src, err := ioutil.ReadFile(fname)
	if err != nil {
		if strings.HasPrefix(fname, "sample") == true {
			t.Errorf("%s, %s", fname, err)
		} else {
			fmt.Fprintf(os.Stderr, "Skipping %s\n", fname)
		}
		return
	}
	fdx := new(FinalDraft)
	if err := xml.Unmarshal(src, &fdx); err != nil {
		t.Errorf("%s", err)
	} else {
		os.RemoveAll(path.Join("testout", path.Base(fname)))
		if src2, err := xml.MarshalIndent(fdx, " ", "    "); err != nil {
			t.Errorf("%s", err)
		} else {
			if err := ioutil.WriteFile(path.Join("testout", path.Base(fname)), src2, 0666); err != nil {
				t.Errorf("%s", err)
			}
		}
	}
}

func TestConversion(t *testing.T) {
	fileList := []string{
		"Big%20Fish.fdx",
		"Brick%20&%20Steel.fdx",
		"The%20Last%20Birthday%20Card.fdx",
		"sample-01.fdx",
		"sample-02.fdx",
		"sample-03.fdx",
	}
	for _, fname := range fileList {
		testFdxFile(t, path.Join("testdata", fname))
	}
}

func TestTitlePageToString(t *testing.T) {
	// The following shouldn't return populated maps
	noTitlePages := []string{
		"sample-01.fdx",
		"sample-02.fdx",
	}
	for _, fname := range noTitlePages {
		fullName := path.Join("testdata", fname)
		src, err := ioutil.ReadFile(fullName)
		if err != nil {
			t.Errorf("%s", err)
		} else {
			screenplay := &FinalDraft{}
			if err := xml.Unmarshal(src, &screenplay); err != nil {
				t.Errorf("Can't Unmarshal %s, %s", fullName, err)
			} else {
				if screenplay.TitlePage != nil {
					page := screenplay.TitlePage.String()
					if len(page) > 0 {
						t.Errorf("was expecting an nil TitlePage in %q, got %q\n", fullName, page)
						//fmt.Printf("DEBUG src %q\n\n%s\n", fullName, src)
					}
				}
			}
		}
	}

	haveTitlePages := map[string][]string{
		"Big%20Fish.fdx":                   []string{"BIG FISH", "This is a Southern story, full of lies and fabrications, "},
		"Brick%20&%20Steel.fdx":            []string{"BRICK AND STEEL"},
		"The%20Last%20Birthday%20Card.fdx": []string{"THE LAST BIRTHDAY CARD"},
		"sample-03.fdx":                    []string{"SAMPLE 03"},
	}
	for fname, textTerms := range haveTitlePages {
		src, err := ioutil.ReadFile(path.Join("testdata", fname))
		if err != nil {
			if strings.HasPrefix(fname, "sample") == true {
				t.Errorf("%s", err)
			} else {
				fmt.Printf("Skipping %s\n", fname)
			}
		} else {
			screenplay := new(FinalDraft)
			if err := xml.Unmarshal(src, &screenplay); err != nil {
				t.Errorf("Can't Unmarshal %s, %s", fname, err)
			} else {
				if screenplay.TitlePage == nil {
					t.Errorf("Missing title page for %s", fname)
				} else {
					page := screenplay.TitlePage.String()
					for _, knownText := range textTerms {
						if strings.Contains(page, knownText) == false {
							t.Errorf("%s is missing %q", fname, knownText)
						}
					}
				}
			}
		}
	}
}

func TestToString(t *testing.T) {
	expected := "Hello World!"
	text := new(Text)
	text.InnerText = expected
	result := text.String()
	if expected != result {
		t.Errorf("expected (font %q, style %q) %q, got %q for %T", text.Font, text.Style, expected, result, text)
	}
	text.Style = "AllCaps"
	expected = "HELLO WORLD!"
	result = text.String()
	if expected != result {
		t.Errorf("expected (font %q, style %q) %q, got %q for %T", text.Font, text.Style, expected, result, text)
	}
	text.Style = "Underline"
	expected = "_Hello World!_"
	result = text.String()
	if expected != result {
		t.Errorf("expected (font %q, style %q) %q, got %q for %T", text.Font, text.Style, expected, result, text)
	}
	text.Style = "Italic"
	expected = "*Hello World!*"
	result = text.String()
	if expected != result {
		t.Errorf("expected (font %q, style %q) %q, got %q for %T", text.Font, text.Style, expected, result, text)
	}
	text.Style = "Bold"
	expected = "**Hello World!**"
	result = text.String()
	if expected != result {
		t.Errorf("expected (font %q, style %q) %q, got %q for %T", text.Font, text.Style, expected, result, text)
	}
	text.Style = "Bold+Italic"
	expected = "***Hello World!***"
	result = text.String()
	if expected != result {
		t.Errorf("expected (font %q, style %q) %q, got %q for %T", text.Font, text.Style, expected, result, text)
	}
	text.Style = "Underline+Bold"
	expected = "_**Hello World!**_"
	result = text.String()
	if expected != result {
		t.Errorf("expected (font %q, style %q) %q, got %q for %T", text.Font, text.Style, expected, result, text)
	}
	text.Style = "Bold+Underline+AllCaps"
	expected = "_**HELLO WORLD!**_"
	result = text.String()
	if expected != result {
		t.Errorf("expected (font %q, style %q) %q, got %q for %T", text.Font, text.Style, expected, result, text)
	}
	text.Font = "Capitals"
	text.Style = "Underline"
	expected = "_HELLO WORLD!_"
	result = text.String()
	if expected != result {
		t.Errorf("expected (font %q, style %q) %q, got %q for %T", text.Font, text.Style, expected, result, text)
	}

	paragraph := new(Paragraph)
	paragraph.Text = append(paragraph.Text, text)
	expected = expected + "\n"
	result = paragraph.String()
	if expected != result {
		t.Errorf("expected %q, got %q for %T", expected, result, paragraph)
	}
	content := new(Content)
	content.Paragraph = append(content.Paragraph, paragraph)
	result = content.String()
	if expected != result {
		t.Errorf("expected %q, got %q for %T", expected, result, content)
	}
	titlePage := new(TitlePage)
	titlePage.Content = content
	result = titlePage.String()
	if expected != result {
		t.Errorf("expected %q, got %q for %T", expected, result, titlePage)
	}
	doc := new(FinalDraft)
	doc.TitlePage = titlePage
	doc.Content = content
	expected = fmt.Sprintf("%s\n%s", expected, expected)
	result = doc.String()
	if expected != result {
		t.Errorf("expected %q, got %q for %T", expected, result, doc)
	}
}

func TestMain(m *testing.M) {
	// Setup everything, process flags, etc.
	os.Exit(m.Run())
}
