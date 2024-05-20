{
    "scripttool": "A program for converting screen play formats"
}

Action Items
============

Bugs
----

- [ ] fountain parser is miss-identifying some elements as character when they are things like scene headings. This should up in JSON output and characters report, see testdata/sample-07.fountain for example of bug
    - `scripttool fountain2json testdata/sample-07.fountain`
    - `scripttool characters testdata/sample-07.fountain`


Next
----

+ [x] fountain2html
+ [x] fountain2json
+ [x] fountain2fdx (Fountain to Final Draft XML)
+ [x] fountain2osf (Fountain to Open Screenplay Format 1.2,2.0)
+ [x] fdx2fountain (build on fdx2txt in [fdx](https://github.com/rsdoiel/fdx), handle TitlePage better)
+ [x] fdx2osf (Final Draft XML to Open Screenplay Format 1.2,2.0 XML)
+ [x] osf2fdx (Open Screenplay Format 1.2,2.0 XML to Final Draft XML)
+ [x] osf2fountain (Open Screenplay Format 1.2,2.0 XML to Fountain) 
+ [x] fountain2fadein, fadein2fountain
+ [ ] edit (simple line oriented editor with colorization for Fountain/Markdown)
    + auto-convert to/fountain on open
    + auto-convert from fountain to origin version on save
    + timed-autosave
    + backup original files
+ [ ] headings - report the headings as a outline
+ [ ] scenes  - report the number of scenes, order and estimate time
+ [ ] lines - report line count
+ [ ] words - report word count
+ [ ] character - report character count
+ [ ] outline (OPML) to scenes/chapters
+ [ ] notes, sections and synopsis
+ [ ] eprints like metadata for the script
+ [ ] Story timeline (using stn) for temporal story outlines
+ [ ] Scene beat report, summaries per scene happenings (who speaks, description, estimated running time)
+ [ ] Search/indexing for script and related assets


Someday, Maybe
--------------

+ [ ] GUI for managing a screenplay project (like Scrivener but GTK based so as can running under Linux)
+ [ ] PDF generation from various formats (possibly via Pandoc)
+ [ ] write and fountain2html, fdx2html, osf2html, fadein2html using [scrippets](https://fountain.io/scrippets) approach
+ [ ] Create a fountain text to speech script reader
    + [ ] Should support configuration for assigning voices to different characters
+ [ ] fountain2trelby
+ [ ] osf2trelby
+ [ ] fdx2trelby
+ [ ] trelby2fdx
+ [ ] trelby2fountain
+ [ ] trelby2osf

Reference Links
---------------

+ [gofpdf](https://github.com/jung-kurt/gofpdf) - Kurt Jung's Go implementation of fpdf
+ [Fountain](https://fountain.io)
+ [Open Screenplay Format 2.0](https://sourceforge.net/projects/openscrfmt/)
+ [Open Screenplay Format 2.1](https://github.com/severdia/Open-Screenplay-Format)
+ [Screenbox HTML and CSS](https://johnaugust.com/2004/screenbox) is some CSS for marking up HTML classes hand presenting a script section in an HTML page
+ [scrippets Wordpress Plugin](https://wordpress.org/plugins/wp-scrippets/), scrippets.org website appears gone.

