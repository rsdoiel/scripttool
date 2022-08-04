//
// osf is a package for working with Open Screenplay Format 1.2 and 2.0 XML documents.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// BSD 2-Clause License
//
// Copyright (c) 2021, R. S. Doiel
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
package osf

import (
	"fmt"
	"os"
	"testing"
)

var (
	expectedDocs map[string][]byte
)

func TestToString(t *testing.T) {
	expected := "Hello World!"
	text := new(Text)
	text.InnerText = expected
	result := text.String()
	if expected != result {
		t.Errorf("expected (*Text.InnerText) %q, got %q for %T", expected, result, text)
	}

	expected = expected + "\n"
	para := new(Para)
	para.Text = append(para.Text, text)
	result = para.String()
	if (expected) != result {
		t.Errorf("expected %q, got %q for %T", expected, result, para)
	}
	paragraphs := new(Paragraphs)
	paragraphs.Para = append(paragraphs.Para, para)
	result = paragraphs.String()
	if (expected) != result {
		t.Errorf("expected %q, got %q for %T", expected, result, paragraphs)
	}
	titlePage := new(TitlePage)
	titlePage.Para = paragraphs.Para
	result = titlePage.String()
	if (expected) != result {
		t.Errorf("expected %q, got %q for %T", expected, result, titlePage)
	}
	doc := new(OpenScreenplay)
	doc.TitlePage = titlePage
	doc.Paragraphs = paragraphs
	expected = fmt.Sprintf("%s%s", expected, expected)
	result = doc.String()
	if expected != result {
		t.Errorf("expected %q, got %q for %T", expected, result, doc)
	}
}

func TestMain(m *testing.M) {
	// Setup everything, process flags, etc.
	os.Exit(m.Run())
}
