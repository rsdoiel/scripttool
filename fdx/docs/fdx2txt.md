
# USAGE

	fdx2txt [OPTIONS]

## SYNOPSIS

fdx2txt is a command line program that reads an fdx file
and returns plain text


## OPTIONS

```
    -generate-markdown-docs   generate Markdown documentation
    -h, -help                 display help
    -i, -input                set the input filename
    -l, -license              display license
    -nl, -newline             add a trailing newline
    -o, -output               set the output filename
    -quiet                    suppress error messages
    -v, -version              display version
```


## EXAMPLES

Cervert *screenplay.fdx* into *screenplay.txt*.

    fdx2txt -i screenplay.fdx -o screenplay.txt

Or alternatively

    cat screenplay.fdx | fdx2txt > screenplay.txt


fdx2txt v0.0.0-dev
