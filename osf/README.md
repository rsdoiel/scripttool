{
    "title": "OSF, A Go package support Open Screenplay Format"
}
[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)


osf
===

Experimental golang package for working with Open Screenplay Format 2.0.
Open Screenplay Format is an open XML format for screenplays and the
native format (when zipped) for [Fade In](https://www.fadeinpro.com).
Two package will include several demonstration command line programs 
[osf2txt](docs/osf2txt.html) which will read a osf file and render plain 
text in a [Fountain](https://fountain.io) like format, [txt2osf](docs/txt2osf.html) 
which takes a plain text file and attempts to render an OSF 2.0 document 
and finally [fadein2txt](docs/fadein2txt) which will read in a Fade In file 
and write out plain text in Fountain format.

