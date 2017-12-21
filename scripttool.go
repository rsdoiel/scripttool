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
	"encoding/xml"
)

const (
	DocString = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>`
)

type FinalDraft struct {
	XMLName         xml.Name          `json:"-"`
	DocumentType    string            `xml:",attr" json:"type"`
	Template        string            `xml:",attr" json:"template"`
	Version         string            `xml:",attr" json:"version"`
	Content         Content           `xml:",omitempty"`
	TitlePage       TitlePage         `xml:",omitempty"`
	Revisions       Revisions         `xml:",omitempty"`
	ElementSettings []ElementSettings `xml:",omitempty"`

	/*
		SpellCheckIgnoreLists *SpellCheckIgnoreLists `xml:"SpellCheckIgnoreLists"`
		PageLayout            *PageLayout
		WindowState           *WindowState
		TextState             *TextState
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
	XMLName   xml.Name    `json:"-"`
	Paragraph []Paragraph `xml:",omitempty"`
}

type Paragraph struct {
	XMLName         xml.Name
	Type            string            `xml:",attr,omitempty"`
	Number          string            `xml:",attr,omitempty"`
	Alignment       string            `xml:",attr,omitempty"`
	FirstIndent     string            `xml:",attr,omitempty"`
	Leading         string            `xml:",attr,omitempty"`
	LeftIndent      string            `xml:",attr,omitempty"`
	RightIndent     string            `xml:",attr,omitempty"`
	SpaceBefore     string            `xml:",attr,omitempty"`
	Spacing         string            `xml:",attr,omitempty"`
	StartsNewPage   string            `xml:",attr,omitempty"`
	SceneProperties []SceneProperties `xml:",omitempty"`
	DynamicLabel    []DynamicLabel    `xml:",omitempty"`
	Text            []Text            `xml:",omitempty"`
}

type SceneProperties struct {
	XMLName xml.Name
	Length  string `xml:",attr,omitempty"`
	Page    string `xml:",attr,omitempty"`
	Title   string `xml:",attr,omitempty"`
}

type HeaderAndFooter struct {
	XMLName         xml.Name
	FooterFirstPage string `xml:",attr,omitempty"`
	FooterVisible   string `xml:",attr,omitempty"`
	HeaderFirstPage string `xml:",attr,omitempty"`
	HeaderVisible   string `xml:",attr,omitempty"`
	StartingPage    string `xml:",attr,omitempty"`
	Header          Header `xml:",omitempty"`
	Footer          Footer `xml:",omitempty"`
}

type Header struct {
	XMLName   xml.Name
	Paragraph []Paragraph `xml:",omitempty"`
}

type DynamicLabel struct {
	XMLName xml.Name
	Type    string `xml:",attr,omitempty"`
}

type Footer struct {
	XMLName   xml.Name    `json:"-"`
	Paragraph []Paragraph `xml:",omitempty"`
}

type Text struct {
	XMLName        xml.Name `json:"-"`
	AdornmentStyle string   `xml:",attr,omitempty"`
	Background     string   `xml:",attr,omitempty"`
	Color          string   `xml:",attr,omitempty"`
	Font           string   `xml:",attr,omitempty"`
	RevisionID     string   `xml:",attr,omitempty"`
	Size           string   `xml:",attr,omitempty"`
	Style          string   `xml:",attr,omitempty"`
	InnerText      string   `xml:",chardata"`
}

type TitlePage struct {
	XMLName         xml.Name        `json:"-"`
	HeaderAndFooter HeaderAndFooter `xml:",omitempty"`
}

type Revisions struct {
	XMLName        xml.Name   `json:"-"`
	ActiveSet      string     `xml:",attr,omitempty"`
	Location       string     `xml:",attr,omitempty"`
	RevisionMode   string     `xml:",attr,omitempty"`
	RevisionsShown string     `xml:",attr,omitempty"`
	ShowAllMarks   string     `xml:",attr,omitempty"`
	ShowAllSets    string     `xml:",attr,omitempty"`
	Revision       []Revision `xml:",omitempty"`
}

type Revision struct {
	Color        string `xml:",attr,omitempty"`
	FullRevision string `xml:",attr,omitempty"`
	ID           string `xml:",attr,omitempty"`
	Mark         string `xml:",attr,omitempty"`
	Name         string `xml:",attr,omitempty"`
	Style        string `xml:",attr,omitempty"`
}

type ElementSettings struct {
	XMLName       xml.Name      `json:"-"`
	Type          string        `xml:",attr"`
	FontSpec      FontSpec      `xml:",omitempty"`
	ParagraphSpec ParagraphSpec `xml:",omitempty"`
	Behavior      Behavior      `xml:",omitempty"`
}

type FontSpec struct {
	XMLName        xml.Name `json:"-"`
	AdornmentStyle string   `xml:",attr,omitempty"`
	Background     string   `xml:",attr,omitempty"`
	Color          string   `xml:",attr,omitempty"`
	Font           string   `xml:",attr,omitempty"`
	RevisionID     string   `xml:",attr,omitempty"`
	Size           string   `xml:",attr,omitempty"`
	Style          string   `xml:",attr,omitempty"`
}

type ParagraphSpec struct {
	XMLName       xml.Name `json:"-"`
	Alignment     string   `xml:",attr,omitemty"`
	FirstIndent   string   `xml:",attr,omitempty"`
	Leading       string   `xml:",attr,omitempty"`
	LeftIndent    string   `xml:",attr,omitempty"`
	RightIndent   string   `xml:",attr,omitempty"`
	SpaceBefore   string   `xml:",attr,omitempty"`
	Spacing       string   `xml:",attr,omitempty"`
	StartsNewPage string   `xml:",attr,omitempty"`
}

type Behavior struct {
	XMLName    xml.Name `json:"-"`
	PaginateAs string   `xml:",attr,omitempty"`
	ReturnKey  string   `xml:",attr,omitempty"`
	Shortcut   string   `xml:",attr,omitempty"`
}
