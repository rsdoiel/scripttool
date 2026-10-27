% scripttool-fountain(7) scripttool 0.0.12
% R. S. Doiel

# NAME

scripttool-fountain - converting, formatting and rendering Fountain

# SYNOPSIS

scripttool VERB [OPTIONS] [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

Without an input filename these verbs read standard input and write
standard output.

fountain2fadein
: to FadeIn. The output filename is required.

fountain2fdx
: to Final Draft XML. A trailing #12A# in a scene heading becomes the scene's
  Number attribute.

fountain2osf
: to Open Screenplay Format 2.0

fountain2json
: to JSON

fountain2yaml
: to YAML

fountain2fountain
: parse and write the document back out as Fountain

fountainfmt
: pretty print a Fountain document; see the options topic for -notes,
  -synopsis, -section and -width

fountain2html
: Fountain scrippets HTML; see the options topic for -html, -inline-css
  and -link-css

# EXAMPLES

~~~shell
    scripttool fountain2fdx screenplay.fountain screenplay.fdx
    scripttool fountainfmt -width 60 screenplay.fountain
~~~
