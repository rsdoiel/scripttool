{
    "title": "OSF, A Go package support Open Screenplay Format"
}

# Action Items

## Bugs


## Next

+ [ ] implement a txt2osf demonstration
+ [ ] review Text element, make sure I am mapping embedded newlines and formatting correctly
+ [ ] validate ToXML() after FromFountain() can be read by FadeIn

## Someday, Maybe

+ [ ] write and osf2html using [scrippets](https://fountain.io/scrippets) approach
+ [ ] add support for Ron Severdia's Open Screenplay Format 2.1 spec
    + rename divergent structs' xml defs with 20 and 21 suffix
    + make sure they are all 20/21 structs are treaded as tag ",omitempty" 
    + duplicate String methods as needed
    + Parse, ParseFile should work without sniffing using a single struct tree for 1.2, 2.0 or 2.1 

## Completed

+ [x] String (Fountain style plain text) needs to be formatted correctly...
+ [x] Write osf.go, osf_test.go based on [Open Screenplay Format 2.0](https://sourceforge.net/projects/openscrfmt/) and in the mode of [fdx](https://github.com/rsdoiel/fdx) package
+ [x] Write osf2txt
+ [x] Write fadein2osf
+ [x] Write fadein2txt
+ [x] self closing tags should be self closing
+ [x] Support parsing .fadein files (i.e. unzip the Fade In file, then parse document.xml)
+ [x] Add ParseFile() to osf.go, if file extension is ".fadein" then it should handle the unzipping and and parsing of document.xml as OSF

### Reference links

+ [Fountain](https://fountain.io)
+ [Open Screenplay Format 2.0](https://sourceforge.net/projects/openscrfmt/) (the one targetted by osf.go)
+ [Open Screenplay Format 2.1](https://github.com/severdia/Open-Screenplay-Format)
+ [Fade In](https://www.fadeinpro.com)
+ [Open Screenplay Format by Kent Tessman](http://www.kenttessman.com/2012/02/open-screenplay-format/)
+ [screenplay-parser](https://github.com/azcoppen/screenplay-parser) - a PHP repo with a really nice README.md discussing format and convertion issues and challenges

