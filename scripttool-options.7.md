% scripttool-options(7) scripttool 0.0.13
% R. S. Doiel

# NAME

scripttool-options - options that follow a verb

# SYNOPSIS

scripttool VERB [OPTIONS] [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

-i FILENAME
: input filename

-o FILENAME
: output filename

-notes
: include notes in output

-synopsis
: include synopsis in output

-section
: include section headings in output

-width N
: maximum line width

-html
: output a full HTML page

-inline-css
: include inline CSS (with -html)

-link-css
: include a CSS link (with -html)

-pretty
: pretty print output

-alpha
: sort a character list alphabetically

A filename of - means standard input or standard output.
