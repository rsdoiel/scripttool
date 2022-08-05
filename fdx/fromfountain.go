// fdx is a package encoding/decoding fdx formatted XML files.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// # BSD 2-Clause License
//
// Copyright (c) 2019, R. S. Doiel
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
package fdx

import (
	"strings"

	// My Packages
	"github.com/rsdoiel/scriptool/fountain"
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
func (document *FinalDraft) FromFountain(screenplay *fountain.Fountain) {
	if screenplay.TitlePage != nil {
		// Build the Info section
		if document.TitlePage == nil {
			document.TitlePage = new(TitlePage)
		}
		if document.TitlePage.Content == nil {
			document.TitlePage.Content = new(Content)
		}

		// NOTE: build a map of elements that will belong to Info.
		for _, elem := range screenplay.TitlePage {
			switch strings.ToLower(elem.Name) {
			case "title":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "credit":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "author":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "copyright":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "draft date":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "contact":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "source":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			case "story by":
				paragraph := new(Paragraph)
				text := new(Text)
				text.InnerText = elem.Content
				paragraph.Text = append(paragraph.Text, text)
				document.TitlePage.Content.Paragraph = append(document.TitlePage.Content.Paragraph, paragraph)
			}
		}
	}

	if screenplay.Elements != nil {
		// Populate the Paragraphs array
		if document.Content == nil {
			document.Content = new(Content)
		}
		for _, elem := range screenplay.Elements {
			text := new(Text)
			text.InnerText = elem.Content
			paragraph := new(Paragraph)
			paragraph.Type = elem.TypeName()
			paragraph.Text = append(paragraph.Text, text)
			document.Content.Paragraph = append(document.Content.Paragraph, paragraph)
		}
	}
}
