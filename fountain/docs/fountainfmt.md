
# USAGE

	fountainfmt [OPTIONS]

## DESCRIPTION

fountainfmt is a command line program that reads an fountain document and pretty prints it.


## OPTIONS

Below are a set of options available.

```
    -debug               display type and element content
    -generate-manpage    generate man page
    -generate-markdown   generate Markdown documentation
    -h, -help            display help
    -i, -input           set the input filename
    -l, -license         display license
    -nl, -newline        add a trailing newline
    -o, -output          set the output filename
    -quiet               suppress error messages
    -v, -version         display version
    -w, -width           set the width for the text
```


## EXAMPLES

Pretty print *screenplay.txt* saving it as *screenplay.fountain*.

    fountainfmt -i screenplay.txt -o screenplay.fountain

Or alternatively

    cat screenplay.txt | foutnainfmt > screenplay.fountain


fountainfmt v0.0.2
