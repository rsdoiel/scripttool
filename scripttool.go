// scripttool is a package focused on converting to/from different
// file formats used in working with scripts,screenplays and other
// creative writing.
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
package scripttool

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"

	// My packages
	"github.com/rsdoiel/fdx"
	"github.com/rsdoiel/fountain"
	"github.com/rsdoiel/osf"
)

var (
	reParens = regexp.MustCompile(`\([[:alnum:]|[:space:]|\.|,|\?]+\)`)
	reConcat = regexp.MustCompile(`&|,| AND `)
)

// FdxToFadeIn converts an input buffer from .fdx to FadeIn file
func FdxToFadeIn(in io.Reader, outputFName string) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := fdx.Parse(src)
	if err != nil {
		return err
	}
	// Convert fdx file to fountain
	screenplay, err := fountain.Parse([]byte(document.String()))
	if err != nil {
		return err
	}
	// Convert fountain to OSF 2.0
	osfDoc := osf.NewOpenScreenplay20()
	document.FromFountain(screenplay)
	src, err = osfDoc.ToXML()
	if err != nil {
		return err
	}

	// Write Zip of OSF 2.0 to produce FadeIn file
	outExt := path.Ext(outputFName)
	outOSFName := strings.TrimSuffix(outputFName, outExt) + ".osf"
	out, err := os.Create(outputFName)
	if err != nil {
		return err
	}
	defer out.Close()
	w := zip.NewWriter(out)
	var files = []struct {
		Name string
		Body []byte
	}{
		{outOSFName, src},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			return err
		}
		_, err = f.Write(file.Body)
		if err != nil {
			return err
		}
	}
	if err := w.Close(); err != nil {
		return err
	}
	return nil
}

// FdxToFountain converts an input buffer from .fdx to a .fountain format.
func FdxToFountain(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	document, err := fdx.Parse(src)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", document.String())
	return nil
}

// FdxToJSON converts an input buffer from a .fdx file to a .json format
func FdxToJSON(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := fdx.Parse(src)
	if err != nil {
		return err
	}
	src, err = json.Marshal(document)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", src)
	return nil
}

// FdxToOSF concerts an input buffer from .fdx to Open Screenplay Format 2.0
func FdxToOSF(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := fdx.Parse(src)
	if err != nil {
		return err
	}
	// Convert fdx file to fountain
	screenplay, err := fountain.Parse([]byte(document.String()))
	if err != nil {
		return err
	}
	// Convert fountain to OSF 2.0
	osfDoc := osf.NewOpenScreenplay20()
	document.FromFountain(screenplay)
	src, err = osfDoc.ToXML()
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", src)
	return nil
}

// OSFToFadeIn converts the input buffer from .osf to FadeIn file
func OSFToFadeIn(in io.Reader, outputFName string) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := osf.Parse(src)
	if err != nil {
		return err
	}
	src, err = document.ToXML()
	if err != nil {
		return err
	}

	// Write Zip of OSF 2.0 to produce FadeIn file
	outExt := path.Ext(outputFName)
	outOSFName := strings.TrimSuffix(outputFName, outExt) + ".osf"
	out, err := os.Create(outputFName)
	if err != nil {
		return err
	}
	defer out.Close()
	w := zip.NewWriter(out)
	var files = []struct {
		Name string
		Body []byte
	}{
		{outOSFName, src},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			return err
		}
		_, err = f.Write(file.Body)
		if err != nil {
			return err
		}
	}
	if err := w.Close(); err != nil {
		return err
	}
	return nil
}

// OSFToFdx converts the input buffer from .osf to .fdx format.
func OSFToFdx(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := osf.Parse(src)
	if err != nil {
		return err
	}
	// Now convert to Fountain format
	screenplay, err := fountain.Parse([]byte(document.String()))
	if err != nil {
		return err
	}
	fdxDoc := fdx.NewFinalDraft()
	fdxDoc.FromFountain(screenplay)
	src, err = fdxDoc.ToXML()
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", src)
	return nil
}

// OSFToFountain converts the input buffer from .osf to .fountain format.
func OSFToFountain(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := osf.Parse(src)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", document.String())
	return nil
}

// OSFToJSON converts the input buffer from .osf to .json format.
func OSFToJSON(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := osf.Parse(src)
	if err != nil {
		return err
	}
	src, err = json.Marshal(document)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", src)
	return nil
}

// Fountain2Fountain reads a input buffer as .fountain and pretty prints output as .fountain
func FountainToFountain(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := fountain.Parse(src)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", document.String())
	return nil
}

// Fountain2Fdx converts an input buffer in .fountain format to output buffer in .fdx
func FountainToFdx(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := fountain.Parse(src)
	if err != nil {
		return err
	}
	newDoc := fountainToFdx(document)
	src, err = newDoc.ToXML()
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", src)
	return nil
}

// FountainToOSF converts an input buffer in .fountain format to output buffer in .osf
func FountainToOSF(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := fountain.Parse(src)
	if err != nil {
		return err
	}
	newDoc := fountainToOSF(document)
	src, err = newDoc.ToXML()
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", src)
	return nil
}

// FadeInToFDX converts an input file to .fdx format
func FadeInToFDX(inputFName string, out io.Writer) error {
	// NOTE: Need to unzip, extract document.xml then pass the source
	// of document.xml to osf.Parse()
	r, err := zip.OpenReader(inputFName)
	if err != nil {
		return err
	}
	defer r.Close()
	src := []byte{}
	for _, f := range r.File {
		if f.Name == "document.xml" {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			src, err = ioutil.ReadAll(rc)
			if err != nil {
				return err
			}
			rc.Close()
			break
		}
	}
	document, err := osf.Parse(src)
	if err != nil {
		return err
	}
	// Convert OSF to Fountain
	screenplay, err := fountain.Parse([]byte(document.String()))
	if err != nil {
		return err
	}
	// Now convert the fountain format to .fdx
	fdxDoc := fdx.NewFinalDraft()
	fdxDoc.FromFountain(screenplay)
	src, err = fdxDoc.ToXML()
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", src)
	return nil
}

// FadeInToFountain converts an input file to .fountain format.
func FadeInToFountain(inputFName string, out io.Writer) error {
	// NOTE: Need to unzip, extract document.xml then pass the source
	// of document.xml to osf.Parse()
	r, err := zip.OpenReader(inputFName)
	if err != nil {
		return err
	}
	defer r.Close()
	src := []byte{}
	for _, f := range r.File {
		if f.Name == "document.xml" {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			src, err = ioutil.ReadAll(rc)
			if err != nil {
				return err
			}
			rc.Close()
			break
		}
	}
	document, err := osf.Parse(src)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", document.String())
	return nil
}

// FadeInToJSON converts an input file to JSON format.
func FadeInToJSON(inputFName string, out io.Writer) error {
	// NOTE: Need to unzip, extract document.xml then pass the source
	// of document.xml to osf.Parse()
	r, err := zip.OpenReader(inputFName)
	if err != nil {
		return err
	}
	defer r.Close()
	src := []byte{}
	for _, f := range r.File {
		if f.Name == "document.xml" {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			src, err = ioutil.ReadAll(rc)
			if err != nil {
				return err
			}
			rc.Close()
			break
		}
	}
	document, err := osf.Parse(src)
	if err != nil {
		return err
	}
	src, err = json.Marshal(document)
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", src)
	return nil
}

// FadeInToOSF converts an input file to .fountain format.
func FadeInToOSF(inputFName string, out io.Writer) error {
	// NOTE: Need to unzip, extract document.xml then pass the source
	// of document.xml to osf.Parse()
	r, err := zip.OpenReader(inputFName)
	if err != nil {
		return err
	}
	defer r.Close()
	src := []byte{}
	for _, f := range r.File {
		if f.Name == "document.xml" {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			src, err = ioutil.ReadAll(rc)
			if err != nil {
				return err
			}
			rc.Close()
			break
		}
	}
	document, err := osf.Parse(src)
	if err != nil {
		return err
	}
	src, err = document.ToXML()
	if err != nil {
		return err
	}
	fmt.Fprintf(out, "%s", src)
	return nil
}

// FountainToFadeIn converts an input buffer in .fountain format to output buffer in .fadein
func FountainToFadeIn(in io.Reader, outFName string) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	document, err := fountain.Parse(src)
	if err != nil {
		return err
	}
	newDoc := fountainToOSF(document)
	src, err = newDoc.ToXML()
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	f, err := w.Create("document.xml")
	if err != nil {
		return err
	}
	_, err = f.Write(src)
	if err != nil {
		return err
	}
	w.Close()
	//NOTE: How do we set the Zipfile's name by write our
	// buf out to disc.
	err = ioutil.WriteFile(outFName, buf.Bytes(), 0664)
	if err != nil {
		return err
	}
	return nil
}

func inList(list []string, target string) bool {
	for _, val := range list {
		if strings.Compare(val, target) == 0 {
			return true
		}
	}
	return false
}

// CharacterList lists character in a screenplay (in should
// be fountain formated text).
func CharacterList(in io.Reader, out io.Writer, alphaSort bool) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	// text must be in Fountain format
	screenplay, err := fountain.Parse(src)
	if err != nil {
		return err
	}
	characters := []string{}
	for _, element := range screenplay.Elements {
		if element.Type == fountain.CharacterType {
			name := strings.TrimSpace(element.Content)
			if name != "" {
				// Remove parentheticals
				if reParens.MatchString(name) {
					name = strings.TrimSpace(reParens.ReplaceAllString(name, ""))
				}
				// handle multiple names
				if strings.Contains(name, ` `) {
					parts := []string{}
					for _, s := range strings.Split(name, ` `) {
						if strings.HasSuffix(s, "'S") {
							s = strings.TrimSuffix(s, "'S")
						}
						if strings.TrimSpace(s) != "" {
							parts = append(parts, s)
						}
					}
					name = strings.TrimSpace(strings.Join(parts, " "))
					// Now handle possible & or "AND" joins
					if reConcat.MatchString(name) {
						name = reConcat.ReplaceAllString(name, ",")
						for _, s := range strings.Split(name, ",") {
							s = strings.TrimSpace(s)
							if !inList(characters, s) {
								characters = append(characters, s)
							}
						}
					}
				} else {
					if name != "" {
						if !inList(characters, name) {
							characters = append(characters, name)
						}
					}
				}
			}
		}
	}
	if alphaSort {
		sort.Strings(characters)
	}
	fmt.Fprintf(out, "%s\n", strings.Join(characters, "\n"))
	return nil
}

// FountainToHTML takes a fountain script and formats it in
// the fountain scrippets HTML markup that can be used
// with the scrippets CSS, see https://fountain.io/scrippets
// and https://johnaugust.com/2004/screenbox.
// FIXME: Write wrapping functioHTMLon and add appropriate writer
// method to fountain package.
func FountainToHTML(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	src, err = fountain.Run(src)
	if err != nil {
		return err
	}
	_, err = out.Write(src)
	return err
}

// FountainFmt pretty prints a fountain document, optionally
// passing on sections, synopsis and notes.
func FountainFmt(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	screenplay, err := fountain.Parse(src)
	if err != nil {
		return err
	}
	_, err = out.Write([]byte(screenplay.String()))
	return err
}

// FountainToJSON convert .fountain file to JSON
func FountainToJSON(in io.Reader, out io.Writer) error {
	src, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	screenplay, err := fountain.Parse(src)
	if err != nil {
		return err
	}
	src, err = screenplay.ToJSON()
	if err != nil {
		return err
	}
	_, err = out.Write(src)
	return err
}
