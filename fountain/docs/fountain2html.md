
# USAGE

	fountain2html [OPTIONS]

## DESCRIPTION

fountain2html is a command line program that reads an fountain document and writes out HTML.


## OPTIONS

Below are a set of options available.

```
    -css                 Add link for CSS
    -generate-manpage    generate man page
    -generate-markdown   generate Markdown documentation
    -h, -help            display help
    -html-page           If true output an HTML page otherwise a fragement
    -i, -input           set the input filename
    -inline-css          Add inline CSS
    -l, -license         display license
    -nl, -newline        add a trailing newline
    -o, -output          set the output filename
    -quiet               suppress error messages
    -v, -version         display version
    -w, -width           set the width for the text
```


## EXAMPLES

Convert a *screenplay.fountain* to *screenplay.html*.

    fountain2html -i screenplay.foutnain -o screenplay.html

Or alternatively

    cat screenplay.fountain | foutnain2html > screenplay.html


fountain2html v0.0.2
