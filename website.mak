
# generated with CMTools 0.0.1 844c936

#
# Makefile for running pandoc on all Markdown docs ending in .md
#
PROJECT = scripttool

PANDOC = $(shell which pandoc)

# Internal notes, not website pages. Kept out of the build so they stay unpublished.
NO_PUBLISH = TODO.md

MD_PAGES = $(filter-out $(NO_PUBLISH),$(shell ls -1 *.md))

HTML_PAGES = $(patsubst %.md,%.html,$(MD_PAGES))

build: $(HTML_PAGES) $(MD_PAGES) # pagefind

$(HTML_PAGES): $(MD_PAGES) .FORCE
	if [ -f $(PANDOC) ]; then $(PANDOC) --metadata title=$(basename $@) -s --to html5 $(basename $@).md -o $(basename $@).html \
		--lua-filter=links-to-html.lua \
		--lua-filter=add-col-scope.lua \
	    --template=page.tmpl; fi
	@if [ $@ = "README.html" ]; then mv README.html index.html; fi

pagefind: .FORCE
	# NOTE: I am not including most of the archive in PageFind index since it doesn't make sense in this case.
	pagefind --verbose --glob="{*.html,docs/*.html}" --force-language en-US --exclude-selectors="nav,header,footer" --output-path ./pagefind --site .
	git add pagefind

clean:
	@rm *.html

.FORCE:
