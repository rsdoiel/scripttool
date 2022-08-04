
# USAGE

	osf2txt [OPTIONS]

## SYNOPSIS

osf2txt is a command line program that reads an osf file
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

Cervert *screenplay.osf* into *screenplay.txt*.

    osf2txt -i screenplay.osf -o screenplay.txt

Or alternatively

    cat screenplay.osf | osf2txt > screenplay.txt


