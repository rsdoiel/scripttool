% scripttool-fadein(7) scripttool 0.0.12
% R. S. Doiel

# NAME

scripttool-fadein - converting from FadeIn files

# SYNOPSIS

scripttool VERB FILENAME.fadein [OUTPUT_FILENAME]

# DESCRIPTION

A FadeIn file is a zipped XML document, so it cannot be read from standard
input. The input filename is required for these verbs.

fadein2fdx
: to Final Draft XML

fadein2fountain
: to Fountain

fadein2json
: to JSON

fadein2yaml
: to YAML

fadein2osf
: to Open Screenplay Format 2.0

# EXAMPLES

~~~shell
    scripttool fadein2fountain screenplay.fadein screenplay.fountain
~~~
