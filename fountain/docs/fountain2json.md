
# USAGE

	fountain2json [OPTIONS]

## DESCRIPTION

fountain2json is a command line program that reads an fountain document and returns a JSON representation of it.



## OPTIONS

Below are a set of options available.

```
    -generate-manpage    generate man page
    -generate-markdown   generate Markdown documentation
    -h, -help            display help
    -i, -input           set the input filename
    -l, -license         display license
    -nl, -newline        add a trailing newline
    -o, -output          set the output filename
    -p, -pretty          pretty print the JSON output
    -quiet               suppress error messages
    -v, -version         display version
    -w, -width           set the width for the text
```


## EXAMPLES

Render *screenplay.fountain* as *screenplay.json*.

    fountain2json -i screenplay.fountain -o screenplay.json

Or alternatively

    cat screenplay.fountain | fountain2json > screenplay.json


fountain2json v0.0.2
