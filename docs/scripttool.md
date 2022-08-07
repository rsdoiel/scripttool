% scripttool(1) scripttool user manual
% R. S. Doiel
% Augest 4, 2022

# NAME

scripttool

# SYNOPSIS

scripttool [-help,-version,-license] VERB [OPTIONS] [VERB PARAMETERS]


# DESCRIPTION

_scripttool_ a program for converting between screenplay formats (e.g. .fdx, .fadein, .fountain)

_scripttool_ converts screen play file formats. Supported formats include FileDraft's XML format, FadeIn's zipped XML format, Fountain formatted plain text as the Open Screenplay Format XML documents. The command line program is based on a Go package also called scripttool. The Go package can be compiled to a shared library and integrated with Python via the ctypes package.  

Internally _scripttool_ works with Fountain formatter files for reporting purposes. You can use the `fountain2json` verb to see the data structure represented as JSON for reports.

# verbs

Like many recent command line tools running under POSIX _scripttool_ uses
a `CMD VERB MODIFIERS` scheme. The follownig verbs are provided. Each
"verb" may also have related options.

fdx2fountain
: Convert from Final Draft XML ("*.fdx" files) to Fountain screenplay format)

osf2fountain
: Convert from Open Screenplay Format 2.0 to Fountain screenplay format

fountain2fdx
: Convert a Fountain screenplay formatted file to Final Draft XML

fountain2osf
: Convert a Fountain screenplay formatted file to Open Screenplay Format 2.0 

fadein2fountain
: Converts a FadeIn file to fountain screenplay format

fountain2fadein
: Converts a fountain screenplay formatted file to FadeIn formatted file.

fountainfmt
: Pretty print a fountain screenplay format

fountain2json
: Convert a fountain screenplay to JSON for machine or data processing

characters
: Provide a character list from a fountain formatted file. Internally the fountain file is parse and resulting JSON structure is analyzed to produce the count of character references in the elements of the file.



# OPTIONS

Below are a set of options available.

-help
: display help

-license
: display license

-version
: display version number

# VERB OPTIONS

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

-html
: output full HTML page

-inline-css
: include inline CSS (works with -html option)

-link-css
: include CSS link (works with -html option)

-pretty
: pretty print output

-alpha
: sort characters alphabetically instead of appearence order


# EXAMPLES

Converting *screenplay.fdx* to *screenplay.fountain* (2 examples)

~~~shell
    scripttool fdx2fountain screenplay.fdx screenplay.fountain
    scripttool fdx2fountain -i screenplay.fdx -o screenplay.fountain
~~~

Converting *screenplay.fountain* to *screenplay.fdx* (2 examples)

~~~shell
    scripttool fountain2fdx screenplay.fountain screenplay.fdx
    scripttool fountain2fdx -i screenplay.fountain -o screenplay.fdx
~~~

Listing characters in *screenplay.fountain*. First list is order of appearence and the second set is characters sorted alphabetically using the "-alpha" option. (2 examples each)

~~~shell
    scripttool characters screenplay.fountain
    scripttool characters -i screenplay.fountain
    scripttool characters -alpha screenplay.fountain
    scripttool characters -alpha -i screenplay.fountain
~~~

