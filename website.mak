#
# Makefile for running pandoc on all Markdown docs ending in .md
#
PROJECT = scripttool

MD_PAGES = $(shell ls -1 *.md | grep -v "nav.md")

HTML_PAGES = $(shell ls -1 *.md | grep -v "nav.md" | sed -E 's/.md/.html/g')

MD_DOCS = docs/scripttool.md

HTML_DOCS = scripttool.html

build: $(HTML_PAGES) $(MD_PAGES) $(HTML_DOCS) $(MD_DOCS) license.html 

$(HTML_PAGES): $(MD_PAGES) .FORCE
	pandoc --metadata title=$(basename $@) -s --to html5 $(basename $@).md -o $(basename $@).html \
	    --template=page.tmpl
	@if [ $@ = "README.html" ]; then mv README.html index.html; fi

$(HTML_DOCS): $(MD_DOCS) .FORCE
	pandoc --metadata title=$(basename $@) -s --to html5 docs/$(basename $@).md -o $(basename $@).html \
	    --template=page.tmpl

license.html: LICENSE
	pandoc --metadata title="$(PROJECT): License" -s --from Markdown --to html5 LICENSE -o license.html \
	    --template=page.tmpl

clean:
	@if [ -f index.html ]; then rm *.html; fi

.FORCE:
