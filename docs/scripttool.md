% scripttool(1) scripttool user name
% R. S. Doiel
% Augest 4, 2022

# NAME

scripttool

# SYNOPSIS

scripttool [-help,-version,-license] VERB [OPTIONS] [VERB PARAMETERS]


# DESCRIPTION

_scripttool_ a program for converting between screenplay formats (e.g. .fdx, .fadein, .fountain)

_scripttool_ converts screen play file formats. Supported formats include FileDraft's XML format, FadeIn's zipped XML format, Fountain formatted plain text as the Open Screenplay Format XML documents. The command line program is based on a Go package also called scripttool. The Go package can be compiled to a shared library and integrated with Python via the ctypes package.  

# OPTIONS

Below are a set of options available.

-help
: display help

-license
: display license

-version
: display version number

## VERB OPTIONS

-i
: set input filename

-o
: set output filename

-notes
: include notes in output

-synopsis
: include synopsis in output

-section
: include section headings in output

-width
: set max width in integers

-height
: section height

-width
: section width

-page
: output full HTML page

-inline-css
: include inline CSS

-link-css
: include CSS link

-css
: include custom CSS

-pretty
: pretty print output


# EXAMPLES

Converting *screenplay.fdx* to *screenplay.fountain* (2 examples)

~~~shell
    scripttool fdx2fountain screenplay.fdx screenplay.fountain
    scripttool -i screenplay.fdx -o screenplay.fountain fdx2fountain
~~~

Converting *screenplay.fountain* to *screenplay.fdx* (2 examples)

~~~shell
    scripttool fountain2fdx screenplay.fountain screenplay.fdx
    scripttool -i screenplay.fountain -o screenplay.fdx fountain2fdx
~~~

Listing characters in *screenplay.fountain* or in *screenplay.fdx*.
(2 examples each)

~~~shell
    scripttool characters screenplay.fountain
    scripttool -i screenplay.fountain characters
    scripttool characters screenplay.fdx
    scripttool -i screenplay.fdx characters
~~~

