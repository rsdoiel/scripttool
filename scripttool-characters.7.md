% scripttool-characters(7) scripttool 0.0.12
% R. S. Doiel

# NAME

scripttool-characters - list the characters in a Fountain file

# SYNOPSIS

scripttool characters [-alpha] [INPUT_FILENAME]

# DESCRIPTION

Parses the Fountain document and counts the references to each character
in its elements. The list is in order of appearance; -alpha sorts it
alphabetically instead. The input must be Fountain.

# EXAMPLES

~~~shell
    scripttool characters screenplay.fountain
    scripttool characters -alpha screenplay.fountain
~~~
