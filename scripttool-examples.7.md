% scripttool-examples(7) scripttool 0.0.13
% R. S. Doiel

# NAME

scripttool-examples - worked examples

# EXAMPLES

Convert screenplay.fdx to screenplay.fountain, two ways.

~~~shell
    scripttool fdx2fountain screenplay.fdx screenplay.fountain
    scripttool fdx2fountain -i screenplay.fdx -o screenplay.fountain
~~~

Read standard input and write standard output.

~~~shell
    cat screenplay.fountain | scripttool fountain2fdx >screenplay.fdx
~~~

List the characters in order of appearance, then alphabetically.

~~~shell
    scripttool characters screenplay.fountain
    scripttool characters -alpha screenplay.fountain
~~~
