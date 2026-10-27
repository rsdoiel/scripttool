% scripttool-fdx(7) scripttool 0.0.13
% R. S. Doiel

# NAME

scripttool-fdx - converting from Final Draft XML

# SYNOPSIS

scripttool VERB [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

Without an input filename these verbs read standard input and write
standard output.

fdx2fadein
: to FadeIn. The output filename is required.

fdx2fountain
: to Fountain. A scene number becomes a trailing #12A# in the scene heading.

fdx2json
: to JSON

fdx2yaml
: to YAML

fdx2osf
: to Open Screenplay Format 2.0

To write Final Draft XML use fountain2fdx, osf2fdx or fadein2fdx.

# EXAMPLES

~~~shell
    scripttool fdx2fountain screenplay.fdx screenplay.fountain
~~~
