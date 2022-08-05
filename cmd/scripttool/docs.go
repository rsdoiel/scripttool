package main

const (
	helpText = `
% scripttool(1) scriptool user manual
% R. S. Doiel
% August 4, 2022

# NAME

{app_name}

# SYNOPSIS

{app_name} VERB [OPTIONS] [PARAMETERS]

# DESCRIPTION

{app_name} is a program for converting between screenplay formats (e.g. .fdx, .fadein, .fountain). Supported formats include FileDraft's XML format, FadeIn's zipped XML format, Fountain formatted plain text as the Open Screenplay Format XML documents. The command line program is based on a Go package also called {app_name}. The Go package can be compiled to a shared library and integrated with Python via the ctypes package.  

# EXAMPLES
Converting *screenplay.fdx* to *screenplay.fountain* (2 examples)

` + "```" + `
    {app_name} fdx2fountain screenplay.fdx screenplay.fountain
    {app_name} fdx2fountain -i screenplay.fdx -o screenplay.fountain 
` + "```" + `

Converting *screenplay.fountain* to *screenplay.fdx* (2 examples)

` + "```" + `
    {app_name} fountain2fdx screenplay.fountain screenplay.fdx
    {app_name} fountain2fdx -i screenplay.fountain -o screenplay.fdx 
` + "```" + `

Converting *screenplay.fountain* to *screenplay.html* to produce
an HTML fragment suitable for including in a webpage. (2 examples)

` + "```" + `
    {app_name} fountain2html screenplay.fountain screenplay.html
    {app_name} -i screenplay.fountain -o screenplay.html fountain2html
` + "```" + `

Converting *screenplay.fountain* to *screenplay.html* to produce
a full HTML page with inline CSS (2 examples)

` + "```" + `
    {app_name} fountain2html -page -inline-css screenplay.fountain \
	    screenplay.html
    {app_name} -i screenplay.fountain -o screenplay.html \
	    fountain2html -page -inline-css
` + "```" + `

Converting *screenplay.fountain* to *screenplay.json* 

` + "```" + `
    {app_name} fountain2json screenplay.fountain screenplay.html
    {app_name} -i screenplay.fountain -o screenplay.html fountain2json
` + "```" + `

Pretty print *screenplay.fountain* to *draft-01.foutain*, then 
pretty print to *draft-02.fountain include notes, sections and synopsis. 

` + "```" + `
    {app_name} fountainfmt screenplay.fountain draft-01.fountain
	{app_name} fountainfmt -notes -sections -synopsis \
	           screenplay.fountain draft-02.fountain
` + "```" + `

Listing characters in *screenplay.fountain* or in *screenplay.fdx*.
(2 examples each)

` + "```" + `
    {app_name} characters screenplay.fountain
    {app_name} -i screenplay.fountain characters
    {app_name} characters screenplay.fdx
    {app_name} -i screenplay.fdx characters
` + "```" + `

# ALSO SEE

- Website: https://rsdoiel.github.io/{app_name}
`

	licenseText = `
{app_name} {version}

BSD 2-Clause License

Copyright (c) 2021, R. S. Doiel
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`
)
