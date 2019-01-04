
# USAGE

	scripttool [OPTIONS] VERB [VERB PARAMETERS]

## SYNOPSIS

_scripttool_ a program for converting between screenplay formats (e.g. .fdx, .fadein, .fountain)

## DESCRIPTION

_scripttool_ converts screen play file formats. Supported formats include FileDraft's XML format, FadeIn's zipped XML format, Fountain formatted plain text as the Open Screenplay Format XML documents. The command line program is based on a Go package also called scripttool. The Go package can be compiled to a shared library and integrated with Python via the ctypes package.  

## OPTIONS

Below are a set of options available.

```
    -examples            display examples
    -generate-manpage    generate man page
    -generate-markdown   generate Markdown documentation
    -h, -help            display help
    -i, -input           set input filename
    -l, -license         display license
    -nl, -newline        add a trailing newline to output
    -o, -output          set output filename
    -quiet               suppress error messages
    -v, -version         display version
```


## EXAMPLES


Converting *screenplay.fdx* to *screenplay.fountain* (2 examples)

```
    scripttool fdx2fountain screenplay.fdx screenplay.fountain
    scripttool -i screenplay.fdx -o screenplay.fountain fdx2fountain
```

Converting *screenplay.fountain* to *screenplay.fdx* (2 examples)

```
    scripttool fountain2fdx screenplay.fountain screenplay.fdx
    scripttool -i screenplay.fountain -o screenplay.fdx fountain2fdx
```

Listing characters in *screenplay.fountain* or in *screenplay.fdx*.
(2 examples each)

```
    scripttool characters screenplay.fountain
    scripttool -i screenplay.fountain characters
    scripttool characters screenplay.fdx
    scripttool -i screenplay.fdx characters
```


scripttool v0.0.3
