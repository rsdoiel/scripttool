
# USAGE

	txt2fdx [OPTIONS]

## SYNOPSIS

txt2fdx is a command line program that reads a plain text file file
and returns a fdx file.


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

Convert *screenplay.txt* into *screenplay.fdx*.

    txt2fdx -i screenplay.txt -o screenplay.fdx

Or alternatively

    cat screenplay.txt | txt2fdx > screenplay.fdx


txt2fdx v0.0.0-dev
