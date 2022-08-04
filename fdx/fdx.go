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
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"strings"
)

const (
	Version = `v0.0.3`

	DocString = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>`

	// Style
	UnderlineStyle = "Underline"
	ItalicStyle    = "Italic"
	BoldStyle      = "Bold"
	AllCapsStyle   = "AllCaps"
	Strikethrough  = "Strikethrough"

	// Alignments
	CenterAlignment = "Center"
	LeftAlignment   = "Left"
	RightAlignment  = "Right"

	// Types used in ElementSettings and Paragraph elements
	GeneralType       = "General"
	SceneHeadingType  = "Scene Heading"
	ActionType        = "Action"
	CharacterType     = "Character"
	DialogueType      = "Dialogue"
	ParentheticalType = "Parenthetical"
	TransitionType    = "Transition"
	CastListType      = "Cast List"
	ShotType          = "Shot"
	SingingType       = "Singing"

	// DynamicLabel types
	PageNoType      = "Page #"
	LastRevisedType = "Last Revised"

	// Tabstop types
	RightType = "Right"
	LeftType  = "Left"
)

var (
	// MaxLineWidth is the number of characters wide a line can be
	// based on a monospace font.
	MaxLineWidth = 80
)

type FinalDraft struct {
	XMLName               xml.Name `json:"-"`
	DocumentType          string   `xml:",attr,omitempty" json:"type,omitempty"`
	Template              string   `xml:",attr" json:"template,omitempty"`
	Version               string   `xml:",attr" json:"version,omitempty"`
	Content               *Content
	TitlePage             *TitlePage
	ElementSettings       []*ElementSettings
	HeaderAndFooter       *HeaderAndFooter
	SpellCheckIgnoreLists *SpellCheckIgnoreLists
	PageLayout            *PageLayout
	WindowState           *WindowState
	TextState             *TextState
	ScriptNoteDefinitions *ScriptNoteDefinitions
	SmartType             *SmartType
	MoresAndContinueds    *MoresAndContinueds
	LockedPages           *LockedPages
	Revisions             *Revisions
	SplitState            *SplitState
	Macros                *Macros
	Actors                *Actors
	Cast                  *Cast `xml:"Cast,omitempty"`
	SceneNumberOptions    *SceneNumberOptions
}

type Content struct {
	XMLName   xml.Name     `json:"-"`
	Paragraph []*Paragraph `json:"paragraphs,omitempty"`
}

type Paragraph struct {
	XMLName         xml.Name `json:"-"`
	Type            string   `xml:",attr,omitempty" json:"type,omitempty"`
	Number          string   `xml:",attr,omitempty" json:"number,omitempty"`
	Alignment       string   `xml:",attr,omitempty" json:"alignment,omitempty"`
	FirstIndent     string   `xml:",attr,omitempty" json:"first_indent,omitempty"`
	Leading         string   `xml:",attr,omitempty" json:"leading,omitempty"`
	LeftIndent      string   `xml:",attr,omitempty" json:"left_indent,omitempty"`
	RightIndent     string   `xml:",attr,omitempty" json:"right_indent,omitempty"`
	SpaceBefore     string   `xml:",attr,omitempty" json:"space_before,omitempty"`
	Spacing         string   `xml:",attr,omitempty" json:"spacing,omitempty"`
	StartsNewPage   string   `xml:",attr,omitempty" json:"starts_new_page,omitempty"`
	SceneProperties []*SceneProperties
	DynamicLabel    []*DynamicLabel
	Text            []*Text
}

type SceneProperties struct {
	XMLName xml.Name `json:"-"`
	Length  string   `xml:",attr,omitempty" json:"length,omitempty"`
	Page    string   `xml:",attr,omitempty" json:"page,omitempty"`
	Title   string   `xml:",attr,omitempty" json:"title,omitempty"`
}

type HeaderAndFooter struct {
	XMLName         xml.Name `json:"-"`
	FooterFirstPage string   `xml:",attr,omitempty" json:"footer_first_page,omitempty"`
	FooterVisible   string   `xml:",attr,omitempty" json:"footer_visible,omitempty"`
	HeaderFirstPage string   `xml:",attr,omitempty" json:"header_first_page,omitempty"`
	HeaderVisible   string   `xml:",attr,omitempty" json:"header_visible,omitempty"`
	StartingPage    string   `xml:",attr,omitempty" json:"starting_page,omitempty"`
	Header          Header
	Footer          Footer
}

type Header struct {
	XMLName   xml.Name    `json:"-"`
	Paragraph []Paragraph `json:"paragraphs,omitempty"`
}

type DynamicLabel struct {
	XMLName xml.Name `json:"-"`
	Type    string   `xml:",attr,omitempty"`
}

type Footer struct {
	XMLName   xml.Name    `json:"-"`
	Paragraph []Paragraph `json:"paragraphs,omitempty"`
}

type Text struct {
	XMLName        xml.Name `json:"-"`
	AdornmentStyle string   `xml:",attr,omitempty" json:"adornment_type,omitempty"`
	Background     string   `xml:",attr,omitempty" json:"background,omitempty"`
	Color          string   `xml:",attr,omitempty" json:"color,omitempty"`
	Font           string   `xml:",attr,omitempty" json:"font,omitempty"`
	RevisionID     string   `xml:",attr,omitempty" json:"revision_id,omitempty"`
	Size           string   `xml:",attr,omitempty" json:"size,omitempty"`
	Style          string   `xml:",attr,omitempty" json:"style,omitempty"`
	InnerText      string   `xml:",chardata" json:"text,omitempty"`
}

type TitlePage struct {
	XMLName         xml.Name `json:"-"`
	HeaderAndFooter *HeaderAndFooter
	Content         *Content
}

type Revisions struct {
	XMLName        xml.Name `json:"-"`
	ActiveSet      string   `xml:",attr,omitempty" json:"active_set,omitempty"`
	Location       string   `xml:",attr,omitempty" json:"location,omitempty"`
	RevisionMode   string   `xml:",attr,omitempty" json:"revision_mode,omitempty"`
	RevisionsShown string   `xml:",attr,omitempty" json:"revisions_shown,omitempty"`
	ShowAllMarks   string   `xml:",attr,omitempty" json:"show_all_marks,omitempty"`
	ShowAllSets    string   `xml:",attr,omitempty" json:"show_all_sets,omitempty"`
	Revision       []Revision
}

type Revision struct {
	XMLName      xml.Name `json:"-"`
	Color        string   `xml:",attr,omitempty" json:"color,omitempty"`
	FullRevision string   `xml:",attr,omitempty" json:"full_revision,omitempty"`
	ID           string   `xml:",attr,omitempty" json:"id,omitempty"`
	Mark         string   `xml:",attr,omitempty" json:"mark,omitempty"`
	Name         string   `xml:",attr,omitempty" json:"name,omitempty"`
	Style        string   `xml:",attr,omitempty" json:"style,omitempty"`
}

type ElementSettings struct {
	XMLName       xml.Name `json:"-"`
	Type          string   `xml:",attr,omitempty" json:"type,omitempty"`
	FontSpec      *FontSpec
	ParagraphSpec *ParagraphSpec
	Behavior      *Behavior
}

type FontSpec struct {
	XMLName        xml.Name `json:"-"`
	AdornmentStyle string   `xml:",attr,omitempty" json:"adornment_style,omitempty"`
	Background     string   `xml:",attr,omitempty" json:"background,omitempty"`
	Color          string   `xml:",attr,omitempty" json:"color,omitempty"`
	Font           string   `xml:",attr,omitempty" json:"font,omitempty"`
	RevisionID     string   `xml:",attr,omitempty" json:"revision_id,omitempty"`
	Size           string   `xml:",attr,omitempty" json:"size,omitempty"`
	Style          string   `xml:",attr,omitempty" json:"style,omitempty"`
}

type ParagraphSpec struct {
	XMLName       xml.Name `json:"-"`
	Alignment     string   `xml:",attr,omitempty" json:"alignment,omitempty"`
	FirstIndent   string   `xml:",attr,omitempty" json:"first_indent,omitempty"`
	Leading       string   `xml:",attr,omitempty" json:"leading,omitempty"`
	LeftIndent    string   `xml:",attr,omitempty" json:"left_indent,omitempty"`
	RightIndent   string   `xml:",attr,omitempty" json:"right_indent,omitempty"`
	SpaceBefore   string   `xml:",attr,omitempty" json:"space_before,omitempty"`
	Spacing       string   `xml:",attr,omitempty" json:"spacing,omitempty"`
	StartsNewPage string   `xml:",attr,omitempty" json:"starts_new_page,omitempty"`
}

type Behavior struct {
	XMLName    xml.Name `json:"-"`
	PaginateAs string   `xml:",attr,omitempty" json:"paginate_as,omitempty"`
	ReturnKey  string   `xml:",attr,omitempty" json:"return_key,omitempty"`
	Shortcut   string   `xml:",attr,omitempty" json:"shortcut,omitempty"`
}

type SpellCheckIgnoreLists struct {
	XMLName       xml.Name `json:"-"`
	IgnoredRanges *IgnoredRanges
	IgnoredWords  []*IgnoredWords
}

type IgnoredRanges struct {
	XMLName xml.Name `json:"-"`
}

type IgnoredWords struct {
	XMLName xml.Name `json:"-"`
	Word    []*Word
}

type Word struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty"`
}

type PageLayout struct {
	XMLName                           xml.Name `json:"-"`
	BackgroundColor                   string   `xml:",attr,omitempty" json:"background_color,omitempty"`
	BottomMargin                      string   `xml:",attr,omitempty" json:"bottom_margin,omitempty"`
	BreakDialogueAndActionAtSentences string   `xml:",attr,omitempty" json:"break_dialogue_and_action_at_sentences,omitempty"`
	DocumentLeading                   string   `xml:",attr,omitempty" json:"document_leading,omitempty"`
	FooterMargin                      string   `xml:",attr,omitempty" json:"footer_margin,omitempty"`
	ForegroundColor                   string   `xml:",attr,omitempty" json:"foreground_color,omitempty"`
	HeaderMargin                      string   `xml:",attr,omitempty" json:"header_margin,omitempty"`
	InvisiblesColor                   string   `xml:",attr,omitempty" json:"invisible_colors,omitempty"`
	TopMargin                         string   `xml:",attr,omitempty" json:"top_margin,omitempty"`
	UsesSmartQuotes                   string   `xml:",attr,omitempty" json:"uses_smart_quotes,omitempty"`
	AutoCastList                      *AutoCastList
}

type AutoCastList struct {
	XMLName               xml.Name `json:"-"`
	AddParentheses        string   `xml:",attr,omitempty" json:"add_parentheses,omitempty"`
	AutomaticallyGenerate string   `xml:",attr,omitempty" json:"automatically_generate,omitempty"`
	CastListElement       string   `xml:",attr,omitempty" json:"cast_list_element,omitempty"`
}

type WindowState struct {
	XMLName xml.Name `json:"-"`
	Height  string   `xml:",attr,omitempty" json:"height,omitempty"`
	Left    string   `xml:",attr,omitempty" json:"left,omitempty"`
	Mode    string   `xml:",attr,omitempty" json:"mode,omitempty"`
	Top     string   `xml:",attr,omitempty" json:"top,omitempty"`
	Width   string   `xml:",attr,omitempty" json:"width,omitempty"`
}

type TextState struct {
	XMLName        xml.Name `json:"-"`
	Scaling        string   `xml:",attr,omitempty" json:"scaling,omitempty"`
	Selection      string   `xml:",attr,omitempty" json:"selection,omitempty"`
	ShowInvisibles string   `xml:",attr,omitempty" json:"show_invisibles,omitempty"`
}

type ScriptNoteDefinitions struct {
	XMLName              xml.Name `json:"-"`
	Active               string   `xml:",attr,omitempty" json:"active,omitempty"`
	ScriptNoteDefinition []*ScriptNoteDefinition
}

type ScriptNoteDefinition struct {
	XMLName xml.Name `json:"-"`
	Color   string   `xml:",attr,omitempty" json:"color,omitempty"`
	ID      string   `xml:",attr,omitempty" json:"id,omitempty"`
	Marker  string   `xml:",attr,omitempty" json:"marker,omitempty"`
	Name    string   `xml:",attr,omitempty" json:"name,omitempty"`
}

type SmartType struct {
	XMLName     xml.Name `json:"-"`
	Characters  *Characters
	Extensions  *Extensions
	SceneIntros *SceneIntros
	Locations   *Locations
	TimesOfDay  *TimesOfDay
	Transitions *Transitions
}

type Characters struct {
	XMLName   xml.Name `json:"-"`
	Character []*Character
}

type Character struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty"`
}

type Extensions struct {
	XMLName   xml.Name `json:"-"`
	Extension []*Extension
}

type Extension struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty"`
}

type SceneIntros struct {
	XMLName    xml.Name `json:"-"`
	SceneIntro []*SceneIntro
}

type SceneIntro struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty"`
}

type Locations struct {
	XMLName  xml.Name `json:"-"`
	Location []*Location
}

type Location struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty"`
}

type TimesOfDay struct {
	XMLName   xml.Name `json:"-"`
	Separator string   `xml:",attr,omitempty" json:"times_of_day,omitempty"`
	TimeOfDay []*TimeOfDay
}

type TimeOfDay struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty"`
}

type Transitions struct {
	XMLName    xml.Name `json:"-"`
	Transition []*Transition
}

type Transition struct {
	XMLName   xml.Name `json:"-"`
	InnerText string   `xml:",chardata" json:"text,omitempty"`
}

type MoresAndContinueds struct {
	XMLName        xml.Name `json:"-"`
	FontSpec       *FontSpec
	DialogueBreaks *DialogueBreaks
	SceneBreaks    *SceneBreaks
}

type DialogueBreaks struct {
	XMLNAme        xml.Name `json:"-"`
	BottomOfPage   string   `xml:",attr,omitempty" json:"bottom_of_page,omitempty"`
	DialogueBottom string   `xml:",attr,omitempty" json:"dialogue_bottom,omitempty"`
	DialogueTop    string   `xml:",attr,omitempty" json:"dialogue_top,omitempty"`
	TopOfNext      string   `xml:",attr,omitempty" json:"top_of_next,omitempty"`
}

type SceneBreaks struct {
	XMLName           xml.Name `json:"-"`
	ContinuedNumber   string   `xml:",attr,omitempty" json:"continued_number,omitempty"`
	SceneBottom       string   `xml:",attr,omitempty" json:"scene_bottom,omitempty"`
	SceneBottomOfPage string   `xml:",attr,omitempty" json:"scene_bottom_of_page,omitempty"`
	SceneTop          string   `xml:",attr,omitempty" json:"scene_top,omitempty"`
	SceneTopOfNext    string   `xml:",attr,omitempty" json:"scene_top_of_next,omitempty"`
}

type LockedPages struct {
	XMLName xml.Name `json:"-"`
}

type Macros struct {
	XMLName xml.Name `json:"-"`
	Macro   []*Macro
}

type Macro struct {
	XMLName    xml.Name `json:"-"`
	Element    string   `xml:",attr,omitempty" json:"element,omitempty"`
	Name       string   `xml:",attr,omitempty" json:"name,omitempty"`
	Shortcut   string   `xml:",attr,omitempty" json:"shortcut,omitempty"`
	Text       string   `xml:",attr,omitempty" json:"text,omitempty"`
	Transition string   `xml:",attr,omitempty" json:"transition,omitempty"`
	Alias      []*Alias
}

type Alias struct {
	XMLName      xml.Name `json:"-"`
	Confirm      string   `xml:",attr,omitempty" json:"confirm,omitempty"`
	MatchCase    string   `xml:",attr,omitempty" json:"match_case,omitempty"`
	SmartReplace string   `xml:",attr,omitempty" json:"smart_replace,omitempty"`
	Text         string   `xml:",attr,omitempty" json:"text,omitempty"`
	WordOnly     string   `xml:",attr,omitempty" json:"word_only,omitempty"`
	ActivateIn   []*ActivateIn
}

type ActivateIn struct {
	XMLName xml.Name `json:"-"`
	Element string   `xml:",attr,omitempty" json:"element,omitempty"`
}

type Actors struct {
	XMLName xml.Name `json:"-"`
	Actor   []*Actor
}

type Actor struct {
	XMLName  xml.Name `json:"-"`
	MacVoice string   `xml:",attr,omitempty" json:"mac_voice,omitempty"`
	Name     string   `xml:",attr,omitempty" json:"name,omitempty"`
	Pitch    string   `xml:",attr,omitempty" json:"pitch,omitempty"`
	Speed    string   `xml:",attr,omitempty" json:"speed,omitempty"`
	WinVoice string   `xml:",attr,omitempty" json:"win_voice,omitempty"`
}

type Cast struct {
	XMLName  xml.Name `json:"-"`
	Narrator *Narrator
	Member   []*Member
}

type Narrator struct {
	XMLName xml.Name `json:"-"`
	Element []*Element
}

type Element struct {
	XMLName xml.Name `json:"-"`
	Type    string   `xml:",attr,omitempty" json:"type,omitempty"`
}

type Member struct {
	XMLName   xml.Name `json:"-"`
	Actor     string   `xml:",attr,omitempty" json:"actor,omitempty"`
	Character string   `xml:",attr,omitempty" json:"character,omitempty"`
}

type SplitState struct {
	XMLName          xml.Name `json:"-"`
	ActivePanel      string   `xml:",attr,omitempty" json:"active_panel,omitempty"`
	SplitMode        string   `xml:",attr,omitempty" json:"split_mode,omitempty"`
	SplitterPosition string   `xml:",attr,omitempty" json:"splitter_position,omitempty"`
	ScriptPanel      *ScriptPanel
}

type ScriptPanel struct {
	XMLName     xml.Name `json:"-"`
	DisplayMode string   `xml:",attr,omitempty" json:"display_mode,omitempty"`
	FontSpec    *FontSpec
}

type SceneNumberOptions struct {
	XMLName            xml.Name `json:"-"`
	LeftLocation       string   `xml:",attr,omitempty" json:"left_location,omitempty"`
	RightLocation      string   `xml:",attr,omitempty" json:"right_location,omitempty"`
	ShowNumbersOnLeft  string   `xml:",attr,omitempty" json:"show_numbers_on_left,omitempty"`
	ShowNumbersOnRight string   `xml:",attr,omitempty" json:"show_numbers_on_right,omitempty"`
	FontSpec           *FontSpec
}

// NewFinalDraft returns a new FinalDraft struct
func NewFinalDraft() *FinalDraft {
	document := new(FinalDraft)
	document.DocumentType = "Script"
	document.Version = "1"
	return document
}

// String (of Text) returns plain text in the Fountain format for a single text element
func (text *Text) String() string {
	if text != nil {
		src := text.InnerText
		if strings.Contains(text.Style, AllCapsStyle) || strings.Contains(text.Font, "Capitals") {
			src = strings.ToUpper(src)
		}
		if strings.Contains(text.Style, ItalicStyle) {
			src = "*" + src + "*"
		}
		if strings.Contains(text.Style, BoldStyle) {
			src = "**" + src + "**"
		}
		if strings.Contains(text.Style, UnderlineStyle) {
			src = "_" + src + "_"
		}
		if strings.Contains(text.Style, Strikethrough) {
			src = "~~" + src + "~~"
		}
		return src
	}
	return ""
}

// String (of Paragraph) returns plain text in Fountain format for a single paragraph
func (paragraph *Paragraph) String() string {
	if paragraph != nil {
		src := []string{}
		if paragraph.StartsNewPage == "Yes" {
			src = append(src, "===\n\n")
		}
		for _, text := range paragraph.Text {
			s := text.String()
			switch paragraph.Type {
			case SceneHeadingType:
				s = strings.ToUpper(s)
			case CharacterType:
				s = strings.ToUpper(s)
			case TransitionType:
				s = strings.ToUpper(s)
			case SingingType:
				s = "~" + s
			case ParentheticalType:
				if strings.HasPrefix(s, "(") == false &&
					strings.HasSuffix(s, ")") == false {
					s = "(" + s + ")"
				}
			}
			if len(s) > 0 {
				switch paragraph.Alignment {
				case CenterAlignment:
					s = ">" + s + "<"
				}
			}
			src = append(src, s, "\n")
		}
		switch paragraph.Type {
		/*
			case GeneralType:
				//src = append(src, "\n")
		*/
		case SceneHeadingType:
			src = append(src, "\n")
		case ActionType:
			src = append(src, "\n")
			/*
				case CharacterType:
					//src = append(src, "\n")
				case ParentheticalType:
					//src = append(src, "\n")
			*/
		case DialogueType:
			src = append(src, "\n")
		case TransitionType:
			src = append(src, "\n")
			/*
				case ShotType:
					//src = append(src, "\n")
				default:
					//src = append(src, "\n")
			*/
		}
		return strings.Join(src, "")
	}
	return ""
}

// String (of Content) returns plain text in Fountain format for Content
func (c *Content) String() string {
	if c != nil && c.Paragraph != nil && len(c.Paragraph) > 0 {
		src := []string{}
		for _, p := range c.Paragraph {
			s := p.String()
			src = append(src, s)
		}
		return strings.Join(src, "")
	}
	return ""
}

// String (of TitlePage) returns a plain text in Fountain format (unfielded) for TitlePage
func (tp *TitlePage) String() string {
	if tp != nil && tp.Content != nil {
		// Move through Title Page content and render the plain text.
		return tp.Content.String()
	}
	return ""
}

// String (of FinalDraft) returns a plan text in Fountain format for FinalDraft
func (doc *FinalDraft) String() string {
	if doc != nil {
		src := []string{}
		if doc.TitlePage != nil {
			s := doc.TitlePage.String()
			// FIXME: Apply screen playwide settings (e.g. PageLayout, HeaderAndFooter, etc)
			src = append(src, s, "\n")
		}
		if doc.Content != nil {
			s := doc.Content.String()
			// FIXME: Apply screen playwide settings (e.g. PageLayout, HeaderAndFooter, etc)
			src = append(src, s)
		}
		return strings.Join(src, "")
	}
	return ""
}

/*
// String (of FinalDraft) returns a plain text in Fountain format
func (doc *FinalDraft) String() string {
	if doc != nil && doc.TitlePage != nil && doc.Content != nil {
		return doc.TitlePage.String() + "\n\n" + doc.Content.String()
	}
	if doc != nil && doc.TitlePage != nil {
		return doc.TitlePage.String() + "\n\n"
	}
	if doc != nil && doc.Content != nil {
		return doc.Content.String()
	}
	return ""
}
*/

// Parse takes []byte and returns a FinalDraft struct and error
func Parse(src []byte) (*FinalDraft, error) {
	document := new(FinalDraft)
	err := xml.Unmarshal(src, &document)
	return document, err
}

// ParseFile takes a filename and returns a FinalDraft struct and error
func ParseFile(fname string) (*FinalDraft, error) {
	src, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	return Parse(src)
}

// CleanupSelfClosingElements changes something like <styles></styles> to <styles/>
func CleanupSelfClosingElements(src []byte) []byte {
	selfClosing := []string{
		"SceneProperties",
		"Member",
		"FontSpec",
		"DynamicLabel",
		"IgnoredRanges",
		"AutoCastList",
		"WindowState",
		"TextState",
		"ParagraphSpec",
		"Behavior",
		"ScriptNoteDefinition",
		"DialogueBreaks",
		"SceneBreaks",
		"LockedPages",
		"Revision",
		"ActivateIn",
		"Actor",
		"Element",
	}
	for _, elem := range selfClosing {
		src = bytes.Replace(src, []byte("></"+elem+">"), []byte("/>"), -1)
	}
	return src
}

// ToXML takes an FinalDraft struct and renders the XML
func (document *FinalDraft) ToXML() ([]byte, error) {
	src, err := xml.MarshalIndent(document, "", "    ")
	if err != nil {
		return nil, err
	}
	return CleanupSelfClosingElements(src), nil
}
