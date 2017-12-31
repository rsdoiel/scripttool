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
	"fmt"

	// My packages
	"github.com/rsdoiel/fdx"
	"github.com/rsdoiel/fountain"
	"github.com/rsdoiel/osf"
)

//
// Below are conversions to other formats (e.g. .osf/.fadein, and .fdx)
//

// fountainToOSF takes a Fountain struct and returns a osf.OpenScreenplay struct
func fountainToOSF(document *fountain.Fountain) (*osf.OpenScreenplay, error) {
	screenplay := new(osf.OpenScreenplay)
	//FIXME: implement the translation from Fountain struct to OSF struct
	return screenplay, fmt.Errorf("FountainToOSF() not implemented")
}

// fountainToFdx takes a Fountain struct and returns a fdx.FinalDraft struct
func fountainToFdx(document *fountain.Fountain) (*fdx.FinalDraft, error) {
	screenplay := new(fdx.FinalDraft)
	//FIXME: implement the translation from Fountain struct to FinalDraft struct
	return screenplay, fmt.Errorf("FountainToFdx() not implemented")
}
