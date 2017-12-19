//
// scripttool is a package focused on converting to/from different
// file formats used in working with scripts,screenplays and other
// creative writing work.
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
)

const (
	DocString = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>`
)

type FinalDraft struct {
	XMLName      xml.Name   `json:"-"`
	DocumentType string     `xml:"DocumentType,attr" json:"type"`
	Template     string     `xml:"Template,attr" json:"template"`
	Version      string     `xml:"Version,attr" json:"version"`
	Content      *Content   `xml:"Content,omitempty"`
	TitlePage    *TitlePage `xml:"TitlePage,omitempty"`
	Revisions    *Revisions `xml:"Revisions,omitempty"`

	/*
		HanderAndFooter       *HeaderAndFooter       `xml:"HeaderAndFooter"`
		SpellCheckIgnoreLists *SpellCheckIgnoreLists `xml:"SpellCheckIgnoreLists"`
		PageLayout            *PageLayout
		WindowState           *WindowState
		TextState             *TextState
		ElementSettings       []*ElementSettings
		ScriptNoteDefinitions *ScriptNoteDefinitions
		SmartType             *SmartType
		MoresAndContinues     *MoresAndContinies
		LockedPages *LockedPages
		Macros *Macros
		Actors    []*Actor   `xml:"Actors"`
		Cast      *Cast      `xml:"Cast"`
		SplitState *SplitState
		SceneOptions *SceneOptions
	*/
}

type Content struct {
	XMLName xml.Name
	Nodes   []*interface{} `xml:"Paragraph"`
}

type Paragraph struct {
	XMLName       xml.Name
	Type          string `xml:"Type,attr,omitempty"`
	Number        string `xml:"Number,attr,omitempty"`
	Alignment     string `xml:"Alignment,attr,omitempty"`
	FirstIndent   string `xml:"FirstIndent,attr,omitempty"`
	Leading       string `xml:"Leading,attr,omitempty"`
	LeftIndent    string `xml:"LeftIndent,attr,omitempty"`
	RightIndent   string `xml:"RightIndent,attr,omitempty"`
	SpaceBefore   string `xml:"SpaceBefore,attr,omitempty"`
	Spacing       string `xml:"Spacing,attr,omitempty"`
	StartsNewPage string `xml:"StartsNewPage,attr,omitempty"`
	Nodes         []*interface{}
}

type SceneProperties struct {
	XMLName xml.Name
	Length  string `xml:"Length,attr,omitempty"`
	Page    string `xml:"Page,attr,omitempty"`
	Title   string `xml:"Title,attr,omitempty"`
}

type HeaderAndFooter struct {
	XMLName         xml.Name
	FooterFirstPage string `xml:"FooterFirstPage,attr,omitempty"`
	FooterVisible   string `xml:"FooterVisible,attr,omitempty"`
	HeaderFirstPage string `xml:"HeaderFirstPage,attr,omitempty"`
	HeaderVisible   string `xml:"HeaderVisible,attr,omitempty"`
	StartingPage    string `xml:"StartingPage,attr,omitempty"`
}

type Header struct {
	XMLName xml.Name
}

type Footer struct {
	XMLName xml.Name
}

type Text struct {
	XMLName        xml.Name `json:"-"`
	AdornmentStyle string   `xml:"AdornmentStyle,attr,omitempty"`
	Background     string   `xml:"Background,attr,omitempty"`
	Color          string   `xml:"Color,attr,omitempty"`
	Font           string   `xml:"Font,attr,omitempty"`
	RevisionID     string   `xml:"RevisionID,attr,omitempty"`
	Size           string   `xml:"Size,attr,omitempty"`
	Style          string   `xml:"Style,attr,omitempty"`
	InnerText      string   `xml:",chardata"`
}

// String prints the raw translation of the XML to struct
func (doc *DocumentRoot) String() string {
	return fmt.Sprintf("%+v", doc)
}

// Marshal takes a DocumentRoot struct and returns a XML version
func Marshal(data *DocumentRoot) ([]byte, error) {
}

// Unmarshal takes a FDX byte array and turns it into a DcoumentRoot struct
func Unmarshal(src []byte, data *DocumentRoot) error {
}
