% scripttool-osf(7) scripttool 0.0.12
% R. S. Doiel

# NAME

scripttool-osf - converting from Open Screenplay Format 2.0

# SYNOPSIS

scripttool VERB [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

Without an input filename these verbs read standard input and write
standard output. Open Screenplay Format keeps no per-scene number.

osf2fadein
: to FadeIn. The output filename is required.

osf2fdx
: to Final Draft XML

osf2fountain
: to Fountain

osf2json
: to JSON

osf2yaml
: to YAML

# EXAMPLES

~~~shell
    scripttool osf2fountain screenplay.osf screenplay.fountain
~~~
