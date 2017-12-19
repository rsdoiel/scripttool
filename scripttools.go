//
// scripttools is a package focused on converting to/from different
// file formats used in working with scripts,screenplays and other
// creative writing work.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
//BSD 3-Clause License
//
//Copyright (c) 2017, R. S. Doiel
//All rights reserved.
//
//Redistribution and use in source and binary forms, with or without
//modification, are permitted provided that the following conditions are met:
//
//* Redistributions of source code must retain the above copyright notice, this
//  list of conditions and the following disclaimer.
//
//* Redistributions in binary form must reproduce the above copyright notice,
//  this list of conditions and the following disclaimer in the documentation
//  and/or other materials provided with the distribution.
//
//* Neither the name of the copyright holder nor the names of its
//  contributors may be used to endorse or promote products derived from
//  this software without specific prior written permission.
//
//THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
//AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
//IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
//DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
//FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
//DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
//SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
//CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
//OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
//OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package scripttools

import (
	"encoding/xml"
	"io"
)

// Fountain2fdx reads a fountain file and writes a fdx file
func Fountain2fdx(fountain io.Reader, fdx io.Writer) error {
}

// Fdx2fountain reads an fdx file and writes a fountain file
func Fdx2fountain(fdx io.Reader, fountain io.Writer) error {
}

// Trelby2fountain reads a trelby file and writes a fountain file
func Trelby2fountain(trelby io.Reader, fountain io.Writer) error {
}

// Trelby2fdx reads a trelby file and writes an fdx file
func Trelby2fdx(trelby io.Reader, fdx io.Writer) error {
}
