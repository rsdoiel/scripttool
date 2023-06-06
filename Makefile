#
# Simple Makefile for Golang based Projects.
#
PROJECT = scripttool

GIT_GROUP = rsdoiel

RELEASE_DATE=$(shell date +'%Y-%m-%d')

RELEASE_HASH=$(shell git log --pretty=format:'%h' -n 1)

PROGRAMS = $(shell ls -1 cmd)

MAN_PAGES = $(shell ls -1 *.1.md | sed -E 's/\.1.md/.1/g')

HTML_PAGES = $(shell find . -type f | grep -E '\.html')

VERSION = $(shell grep '"version":' codemeta.json | cut -d\"  -f 4)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

PACKAGE = $(shell ls -1 *.go | grep -v 'version.go')

SUBPACKAGES = $(shell ls -1 */*.go)

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

DIST_FOLDERS = bin/*

build: version.go $(PROGRAMS) man CITATION.cff about.md installer.sh

version.go: .FORCE
	echo '' | pandoc --from t2t --to plain \
		--metadata-file codemeta.json \
		--metadata package=$(PROJECT) \
		--metadata version=$(VERSION) \
		--metadata release_date=$(RELEASE_DATE) \
		--metadata release_hash=$(RELEASE_HASH) \
		--template codemeta-version-go.tmpl \
		LICENSE >version.go
##	@echo "package $(PROJECT)" >version.go
##	@echo '' >>version.go
##	@echo 'const (' >>version.go
##	@echo  '    // Version number of release'>>version.go
##	@echo '    Version = "$(VERSION)"' >>version.go
##	@echo '' >>version.go
##	@echo  '    // ReleaseDate, the date version.go was generated'>>version.go
##	@echo '    ReleaseDate = "$(RELEASE_DATE)"' >>version.go
##	@echo '' >>version.go
##	@echo  '    // ReleaseHash, the Git hash when version.go was generated'>>version.go
##	@echo '    ReleaseHash = "$(RELEASE_HASH)"' >>version.go
##	@echo '' >>version.go
##	@echo '    LicenseText = `' >>version.go
##	@cat LICENSE >>version.go
##	@echo '`' >>version.go
##	@echo ')' >>version.go
##	@echo '' >>version.go
##	@git add version.go
##	@if [ -f bin/codemeta ]; then ./bin/codemeta; fi


$(PROGRAMS): $(PACKAGE)
	@mkdir -p bin
	go build -o "bin/$@$(EXT)" cmd/$@/*.go
	@./bin/$@ -help >$@.1.md

man: $(MAN_PAGES)

$(MAN_PAGES): .FORCE
	mkdir -p man/man1
	pandoc $@.md --from markdown --to man -s >man/man1/$@

CITATION.cff: .FORCE
	@cat codemeta.json | sed -E   's/"@context"/"at__context"/g;s/"@type"/"at__type"/g;s/"@id"/"at__id"/g' >_codemeta.json
	@echo '' | pandoc --metadata title="Cite $(PROJECT)" --metadata-file=_codemeta.json --template=codemeta-cff.tmpl >CITATION.cff

about.md: .FORCE 
	@cat codemeta.json | sed -E 's/"@context"/"at__context"/g;s/"@type"/"at__type"/g;s/"@id"/"at__id"/g' >_codemeta.json
	@echo "" | pandoc --metadata-file=_codemeta.json --template codemeta-about.tmpl >about.md 2>/dev/null;
	@if [ -f _codemeta.json ]; then rm _codemeta.json; fi

installer.sh: .FORCE
	@echo '' | pandoc --metadata title="Installer" --metadata git_org_or_person="$(GIT_GROUP)" --metadata-file codemeta.json --template codemeta-installer.tmpl >installer.sh
	@chmod 775 installer.sh
	@git add -f installer.sh


clean-website:
	make -f website.mak clean

website: clean-website .FORCE
	make -f website.mak


# NOTE: on macOS you must use "mv" instead of "cp" to avoid problems
install: build man .FORCE
	@if [ ! -d $(PREFIX)/bin ]; then mkdir -p $(PREFIX)/bin; fi
	@echo "Installing programs in $(PREFIX)/bin"
	@for FNAME in $(PROGRAMS); do if [ -f ./bin/$$FNAME ]; then mv -v ./bin/$$FNAME $(PREFIX)/bin/$$FNAME; fi; done
	@echo ""
	@echo "Make sure $(PREFIX)/bin is in your PATH"
	@echo ""
	@if [ ! -d $(PREFIX)/man/man1 ]; then mkdir -p $(PREFIX)/man/man1; fi
	@for MAN_PAGE in $(MAN_PAGES); do cp -v man/man1/$$MAN_PAGE $(PREFIX)/man/man1/;done
	@echo ""
	@echo "Make sure $(PREFIX)/man is in your MANPATH"
	@echo ""

uninstall: .FORCE
	@echo "Removing programs in $(PREFIX)/bin"
	-for FNAME in $(PROGRAMS); do if [ -f $(PREFIX)/bin/$$FNAME ]; then rm -v $(PREFIX)/bin/$$FNAME; fi; done
	-for MAN_PAGE in $(MAN_PAGES); do if [ -f "$(PREFIX)/man/man1/$$MAN_PAGE" ]; then rm "$(PREFIX)/man/man1/$$MAN_PAGE"; fi; done


hash: .FORCE
	git log --pretty=format:'%h' -n 1

check: .FORCE
	for FNAME in $(shell ls -1 *.go); do go fmt $$FNAME; done
	go vet *.go

test: clean build
	go test

clean: 
	-if [ -d bin ]; then rm -fR bin; fi
	-if [ -d dist ]; then rm -fR dist; fi
	-if [ -d testout ]; then rm -fR testout; fi
	-for MAN_PAGE in $(MAN_PAGES); do if [ -f man/man1/$$MAN_PAGE.1 ]; then rm man/man1/$$MAN_PAGE.1; fi;done

status:
	git status

save:
	@if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

refresh:
	git fetch origin
	git pull origin $(BRANCH)

publish: build website save .FORCE
	./publish.bash


dist/Linux-x86_64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env  GOOS=linux GOARCH=amd64 go build -o "dist/bin/$${FNAME}" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-Linux-x86_64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/*
	@rm -fR dist/bin

dist/Linux-aarch64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env  GOOS=linux GOARCH=arm64 go build -o "dist/bin/$${FNAME}" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-Linux-aarch64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/*
	@rm -fR dist/bin

dist/macOS-x86_64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=darwin GOARCH=amd64 go build -o "dist/bin/$${FNAME}" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-macOS-x86_64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/*
	@rm -fR dist/bin


dist/macOS-arm64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=darwin GOARCH=arm64 go build -o "dist/bin/$${FNAME}" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-macOS-arm64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/*
	@rm -fR dist/bin


dist/Windows-x86_64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=windows GOARCH=amd64 go build -o "dist/bin/$${FNAME}.exe" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-Windows-x86_64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/*
	@rm -fR dist/bin

dist/Windows-arm64: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=windows GOARCH=arm64 go build -o "dist/bin/$${FNAME}.exe" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-Windows-arm64.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/*
	@rm -fR dist/bin

# Raspberry PI OS 32 bit, as reported by Raspberry Pi 3B+
dist/Linux-armv7l: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=linux GOARCH=arm GOARM=7 go build -o "dist/bin/$${FNAME}" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-Linux-armv7l.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/*
	@rm -fR dist/bin


dist/RaspberryPiOS-arm7: $(PROGRAMS)
	@mkdir -p dist/bin
	@for FNAME in $(PROGRAMS); do env GOOS=linux GOARCH=arm GOARM=7 go build -o "dist/bin/$${FNAME}" cmd/$${FNAME}/*.go; done
	@cd dist && zip -r $(PROJECT)-v$(VERSION)-RaspberryPiOS-arm7.zip LICENSE codemeta.json CITATION.cff *.md bin/* man/*
	@rm -fR dist/bin

distribute_docs:
	@mkdir -p dist/
	@cp -v codemeta.json dist/
	@cp -v CITATION.cff dist/
	@cp -v README.md dist/
	@cp -v LICENSE dist/
	@cp -v INSTALL.md dist/
	@cp -vR man dist/

release: .FORCE installer.sh save build save distribute_docs dist/Linux-x86_64 dist/Linux-aarch64 dist/macOS-x86_64 dist/macOS-arm64 dist/Windows-x86_64 dist/Windows-arm64 dist/RaspberryPiOS-arm7 dist/Linux-armv7l

.FORCE:
