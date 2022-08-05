// osf is a package for working with Open Screenplay Format 1.2 and 2.0 XML documents.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
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
package osf

import (
	"strings"

	// My Packages
	"github.com/rsdoiel/scripttool/fountain"
)

func StringToTextArray(s string) []*Text {
	var a []*Text

	//FIXME: if string has embedded formatting split into parts and set format attributes
	// (e.g. text.Bold, text.Italic) as needed.
	text := new(Text)
	text.InnerText = s
	a = append(a, text)
	return a
}

// AddInfo parses a map[string]string and updates/adds Info struct to document as needed.
func (document *OpenScreenplay) FromFountain(screenplay *fountain.Fountain) {
	if screenplay.TitlePage != nil {
		// Build the Info section
		if document.Info == nil {
			document.Info = new(Info)
		}
		// Populate the TitlePage
		if document.TitlePage == nil {
			document.TitlePage = new(TitlePage)
		}

		// NOTE: build a map of elements that will belong to Info.
		for _, elem := range screenplay.TitlePage {
			switch strings.ToLower(elem.Name) {
			case "title":
				document.Info.Title = elem.Content
				para := new(Para)
				para.Bookmark = "Title"
				para.Text = StringToTextArray(elem.Content)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "credit":
				document.Info.WrittenBy = elem.Content
				para := new(Para)
				para.Bookmark = "Credits"
				para.Text = StringToTextArray(elem.Content)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "author":
				document.Info.WrittenBy = elem.Content
				para := new(Para)
				para.Bookmark = "Author"
				para.Text = StringToTextArray(elem.Content)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "copyright":
				document.Info.Copyright = elem.Content
				para := new(Para)
				para.Bookmark = "Copyright"
				para.Text = StringToTextArray(elem.Content)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "draft date":
				document.Info.Drafts = elem.Content
				para := new(Para)
				para.Bookmark = "Drafts"
				para.Text = StringToTextArray(elem.Content)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "contact":
				document.Info.Contact = elem.Content
				para := new(Para)
				para.Bookmark = "Contact"
				para.Text = StringToTextArray(elem.Content)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "source":
				para := new(Para)
				para.Bookmark = "Source"
				para.Text = StringToTextArray(elem.Content)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "story by":
				para := new(Para)
				para.Bookmark = "Story By"
				para.Text = StringToTextArray(elem.Content)
				document.TitlePage.Para = append(document.TitlePage.Para, para)
			case "uuid":
				document.Info.UUID = elem.Content
			case "page_count":
				document.Info.PageCount = elem.Content
			case "title_format":
				document.Info.TitleFormat = elem.Content
			}
		}
	}

	if screenplay.Elements != nil {
		// Populate the Paragraphs array
		if document.Paragraphs == nil {
			document.Paragraphs = new(Paragraphs)
		}
		for _, elem := range screenplay.Elements {
			para := new(Para)
			para.Bookmark = elem.TypeName()
			para.Text = StringToTextArray(elem.Content)
			document.Paragraphs.Para = append(document.Paragraphs.Para, para)
		}
	}
}
