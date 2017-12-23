
# USAGE

	scripttool [OPTIONS] [ACTION] [ACTION PARAMETERS...]

## SYNOPSIS


scripttool a command line program for working with screenplay file formats
(e.g. fdx, fountain).


## OPTIONS

Options are shared between all actions and must precede the action on the command line.

```
    -examples                 display examples
    -generate-markdown-docs   generate Markdown documentation
    -h, -help                 display help
    -i, -input                set input filename
    -l, -license              display license
    -o, -output               set output filename
    -quiet                    suppress error messages
    -v, -version              display version
```


## ACTIONS

```
    characters     list the characters in screenplay
    fdx2fountain   converts fdx to fountain
    fountain2fdx   converts fountain to fdx
```


## EXAMPLES


Convert a *screenplay.fdx* to *screenplay.fountain*

     scripttool fdx2fountain screenplay.fdx screenplay.fountain


Related: [characters](characters.html), [fdx2fountain](fdx2fountain.html), [fountain2fdx](fountain2fdx.html)

scripttool v0.0.0-dev
