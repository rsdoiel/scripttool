
# generated with CMTools 0.0.0 f50285c

#
# Simple Makefile for Golang based Projects built under POSIX.
#
PROJECT = scripttool

GIT_GROUP = rsdoiel

PROGRAMS = scripttool

RELEASE_DATE = $(shell date +%Y-%m-%d)

RELEASE_HASH=$(shell git log --pretty=format:'%h' -n 1)

MAN_PAGES_1 = $(shell ls -1 *.1.md | sed -E 's/.1.md/.1/g')

# None of these projects has a section 3 man page; uncomment to add one.
#MAN_PAGES_3 = $(shell ls -1 *.3.md | sed -E 's/.3.md/.3/g')

# One section 7 man page per help topic, written by the $(PROGRAMS) rule from
# "scripttool help TOPIC". The list is explicit, not a glob, so a new topic is
# not skipped before its page exists; TestMakefileHelpTopicsMatch keeps it equal
# to HelpTopics in help.go. Add a topic in both places.
HELP_TOPICS = fadein fdx osf fountain characters options examples

MAN_PAGES_7 = $(foreach T,$(HELP_TOPICS),$(PROJECT)-$(T).7)

HTML_PAGES = $(shell find . -type f -not -path './bin/*' -not -path './.git/*' | grep -E '\.html$$')

DOCS = $(shell ls -1 *.?.md)

# version.go is left out: it has a rule (cmt, under generate), and listing it
# here would make every build run cmt through that rule.
PACKAGE = $(filter-out version.go,$(shell ls -1 *.go))

VERSION = $(shell grep '"version":' codemeta.json | cut -d\"  -f 4)

BRANCH = $(shell git branch --show-current)

OS = $(shell uname)

#PREFIX = /usr/local/bin
PREFIX = $(HOME)

ifneq ($(prefix),)
	PREFIX = $(prefix)
endif

EXT =
ifeq ($(OS), Windows)
	EXT = .exe
endif

# build compiles the programs and regenerates the help text and man pages. It
# does not run cmt: version.go, CITATION.cff, about.md and the installers are
# generated from codemeta.json by `make generate`, so the version and release
# hash change only when that is asked for (the release target does).
build: $(PROGRAMS) man

# generate runs cmt over codemeta.json. Run it when the version, the release
# notes or the metadata change, and before a release; a plain `make` leaves
# the committed version.go, about.md and friends alone.
generate: version.go CITATION.cff about.md installer.sh installer.ps1

version.go: .FORCE
	cmt codemeta.json version.go

hash: .FORCE
	git log --pretty=format:'%h' -n 1

man: $(MAN_PAGES_1) $(MAN_PAGES_7) # $(MAN_PAGES_3)

$(MAN_PAGES_1): .FORCE
	mkdir -p man/man1
	pandoc $@.md --from markdown --to man -s >man/man1/$@

#$(MAN_PAGES_3): .FORCE
#	mkdir -p man/man3
#	pandoc $@.md --from markdown --to man -s >man/man3/$@

$(MAN_PAGES_7): .FORCE
	mkdir -p man/man7
	pandoc $@.md --from markdown --to man -s >man/man7/$@

$(PROGRAMS): $(PACKAGE)
	@mkdir -p bin
	go build -o "bin/$@$(EXT)" cmd/$@/*.go
	@./bin/$@ -help >$@.1.md
	@for T in $(HELP_TOPICS); do ./bin/$@ help $$T >$@-$$T.7.md; done

CITATION.cff: .FORCE
	cmt codemeta.json CITATION.cff

about.md: codemeta.json $(PROGRAMS)
	cmt codemeta.json about.md

installer.sh: .FORCE
	cmt codemeta.json installer.sh

installer.ps1: .FORCE
	cmt codemeta.json installer.ps1


# check reports files gofmt would change, and vets; it rewrites nothing.
check: .FORCE
	@UNFORMATTED="$$(gofmt -l $(filter-out version.go,$(shell ls -1 *.go)) cmd)"; if [ "$$UNFORMATTED" != "" ]; then echo "gofmt would change:"; echo "$$UNFORMATTED"; fi
	go vet ./...

test: $(PACKAGE)
	go test ./...

website: clean-website .FORCE
	make -f website.mak

status:
	git status

save:
	@if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

refresh:
	git fetch origin
	git pull origin $(BRANCH)

publish: build website .FORCE
	./publish.bash

clean:
	@if [ -f version.go ]; then rm version.go; fi
	@if [ -d bin ]; then rm -fR bin; fi
	@if [ -d dist ]; then rm -fR dist; fi
	@if [ -d man ]; then rm -fR man; fi
	@if [ -d testout ]; then rm -fR testout; fi

clean-website:
	@for FNAME in $(HTML_PAGES); do if [ -f "$${FNAME}" ]; then rm "$${FNAME}"; fi; done

install: build
	@echo "Installing programs in $(PREFIX)/bin"
	@for FNAME in $(PROGRAMS); do if [ -f "./bin/$${FNAME}$(EXT)" ]; then mv -v "./bin/$${FNAME}$(EXT)" "$(PREFIX)/bin/$${FNAME}$(EXT)"; fi; done
	@echo ""
	@echo "Make sure $(PREFIX)/bin is in your PATH"
	@echo "Installing man page in $(PREFIX)/man"
	@mkdir -p $(PREFIX)/man/man1
	@for FNAME in $(MAN_PAGES_1); do if [ -f "./man/man1/$${FNAME}" ]; then cp -v "./man/man1/$${FNAME}" "$(PREFIX)/man/man1/$${FNAME}"; fi; done
#	@mkdir -p $(PREFIX)/man/man3
#	@for FNAME in $(MAN_PAGES_3); do if [ -f "./man/man3/$${FNAME}" ]; then cp -v "./man/man3/$${FNAME}" "$(PREFIX)/man/man3/$${FNAME}"; fi; done
#	@mkdir -p $(PREFIX)/man/man7
#	@for FNAME in $(MAN_PAGES_7); do if [ -f "./man/man7/$${FNAME}" ]; then cp -v "./man/man7/$${FNAME}" "$(PREFIX)/man/man7/$${FNAME}"; fi; done
	@echo ""
	@echo "Make sure $(PREFIX)/man is in your MANPATH"

uninstall: .FORCE
	@echo "Removing programs in $(PREFIX)/bin"
	@for FNAME in $(PROGRAMS); do if [ -f "$(PREFIX)/bin/$${FNAME}$(EXT)" ]; then rm -v "$(PREFIX)/bin/$${FNAME}$(EXT)"; fi; done
	@echo "Removing man pages in $(PREFIX)/man"
	@for FNAME in $(MAN_PAGES_1); do if [ -f "$(PREFIX)/man/man1/$${FNAME}" ]; then rm -v "$(PREFIX)/man/man1/$${FNAME}"; fi; done
#	@for FNAME in $(MAN_PAGES_3); do if [ -f "$(PREFIX)/man/man3/$${FNAME}" ]; then rm -v "$(PREFIX)/man/man3/$${FNAME}"; fi; done
#	@for FNAME in $(MAN_PAGES_7); do if [ -f "$(PREFIX)/man/man7/$${FNAME}" ]; then rm -v "$(PREFIX)/man/man7/$${FNAME}"; fi; done

setup_dist: .FORCE
	@mkdir -p dist
	@rm -fR dist/

dist/Linux-x86_64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env  GOOS=linux GOARCH=amd64 go build -o "dist/bin/$${FNAME}" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-Linux-x86_64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/* $(DOCS)
	@rm -fR dist/bin

dist/Linux-aarch64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env  GOOS=linux GOARCH=arm64 go build -o "dist/bin/$${FNAME}" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-Linux-aarch64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/* $(DOCS)
	@rm -fR dist/bin

dist/macOS-x86_64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=darwin GOARCH=amd64 go build -o "dist/bin/$${FNAME}" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-macOS-x86_64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/* $(DOCS)
	@rm -fR dist/bin


dist/macOS-arm64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=darwin GOARCH=arm64 go build -o "dist/bin/$${FNAME}" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-macOS-arm64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/* $(DOCS)
	@rm -fR dist/bin


dist/Windows-x86_64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=windows GOARCH=amd64 go build -o "dist/bin/$${FNAME}.exe" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-Windows-x86_64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/* $(DOCS)
	@rm -fR dist/bin

dist/Windows-arm64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=windows GOARCH=arm64 go build -o "dist/bin/$${FNAME}.exe" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-Windows-arm64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/* $(DOCS)
	@rm -fR dist/bin

# Raspberry Pi OS (32 bit) based on report from Raspberry Pi Model 3B+
dist/Linux-armv7l: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=linux GOARCH=arm GOARM=7 go build -o "dist/bin/$${FNAME}" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-Linux-armv7l.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/* $(DOCS)
	@rm -fR dist/bin

distribute_docs:
	@mkdir -p dist/
	@cp -v codemeta.json dist/
	@cp -v CITATION.cff dist/
	@cp -v README.md dist/
	@cp -v LICENSE dist/
	@cp -v INSTALL.md dist/
	@cp -vR man dist/
	@for DNAME in $(DOCS); do cp -vR $$DNAME dist/; done

# check-release reads the files a release is cut from and fails when they
# disagree; it writes nothing. Run it before `make release` (the release target
# runs it too, after generate). It fails when
#   - version.go, CITATION.cff or an installer carries a version other than codemeta.json's,
#     or CITATION.cff's date-released is not datePublished (run `make generate`),
#   - datePublished is not today's date, the release day,
#   - v$(VERSION) is already tagged here or on origin (codemeta.json was not bumped),
#   - releaseNotes are the same as the last tag's (not written for this release),
#   - softwareRequirements names a github.com/rsdoiel module at a version go.mod does not require.
check-release: .FORCE
	@FAIL=0; V="$(VERSION)"; TODAY="$$(date +%Y-%m-%d)"; \
	bad() { echo "check-release: $$1"; FAIL=1; }; \
	[ -n "$$V" ] || bad "codemeta.json has no version"; \
	GOV="$$(sed -n 's/^ *Version = "\(.*\)".*/\1/p' version.go | head -1)"; \
	[ "$$GOV" = "$$V" ] || bad "version.go has '$$GOV', codemeta.json has '$$V' (run make generate)"; \
	CFF="$$(sed -n 's/^version: *//p' CITATION.cff | head -1)"; \
	[ "$$CFF" = "$$V" ] || bad "CITATION.cff has '$$CFF', codemeta.json has '$$V' (run make generate)"; \
	DR="$$(sed -n 's/^date-released: *//p' CITATION.cff | head -1)"; \
	PUB="$$(jq -r '.datePublished // empty' codemeta.json)"; \
	[ "$$DR" = "$$PUB" ] || bad "CITATION.cff date-released is '$$DR', datePublished is '$$PUB' (run make generate)"; \
	[ "$$PUB" = "$$TODAY" ] || bad "datePublished is '$$PUB', today is '$$TODAY': set datePublished and dateModified to the release day"; \
	for F in installer.sh installer.ps1; do \
		[ -f "$$F" ] || continue; \
		IV="$$(grep -m1 'VERSION *= *"' "$$F" | sed 's/.*VERSION *= *"\([^"]*\)".*/\1/')"; \
		[ "$$IV" = "$$V" ] || bad "$$F has '$$IV', codemeta.json has '$$V' (run make generate)"; \
	done; \
	if git rev-parse -q --verify "refs/tags/v$$V" >/dev/null; then bad "v$$V is already tagged here: codemeta.json was not bumped"; fi; \
	if git ls-remote --exit-code --tags origin "refs/tags/v$$V" >/dev/null 2>&1; then bad "v$$V is already tagged on origin: codemeta.json was not bumped"; fi; \
	NEW="$$(jq -r '.releaseNotes // empty' codemeta.json)"; \
	[ -n "$$NEW" ] || bad "codemeta.json has no releaseNotes"; \
	LAST="$$(git tag --list 'v*' --sort=-v:refname | head -1)"; \
	if [ -n "$$LAST" ] && [ "$$LAST" != "v$$V" ]; then \
		OLD="$$(git show "$$LAST:codemeta.json" 2>/dev/null | jq -r '.releaseNotes // empty')"; \
		[ "$$NEW" != "$$OLD" ] || bad "releaseNotes are the same as $$LAST's: write them for v$$V"; \
	fi; \
	for REQ in $$(jq -r '.softwareRequirements[]? | select(startswith("github.com/rsdoiel/")) | gsub(" "; "@")' codemeta.json); do \
		M="$${REQ%%@*}"; MIN="$${REQ##*@}"; \
		GM="$$(grep -o "$$M v[^ ]*" go.mod | head -1 | sed 's/.* v//')"; \
		[ -n "$$GM" ] || { bad "softwareRequirements names $$M, which go.mod does not require"; continue; }; \
		[ "$$GM" = "$$MIN" ] || bad "softwareRequirements says $$M >= $$MIN, go.mod requires v$$GM"; \
	done; \
	if [ $$FAIL -ne 0 ]; then echo "check-release: v$$V is NOT ready to release"; exit 1; fi; \
	echo "check-release: v$$V is consistent"

release: generate check-release build save setup_dist distribute_docs dist/Linux-x86_64 dist/Linux-aarch64 dist/macOS-x86_64 dist/macOS-arm64 dist/Windows-x86_64 dist/Windows-arm64 dist/Linux-armv7l
	@printf "\nready to run\n\n\trelease.bash\n\n"


.FORCE:
