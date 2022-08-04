#
# Simple Makefile
#
PROJECT = scripttool

VERSION = $(shell grep -m1 'Version = `' $(PROJECT).go | cut -d\`  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

OS = $(shell uname)

#PREFIX = /usr/local
PREFIX = $(HOME)


EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif

build:
	go build -o bin/scripttool$(EXT) cmd/scripttool/scripttool.go

test:
	go test

man: build
	bin/scripttool$(EXT) -generate-manpage > man/man1/scripttool.man
	cat man/man1/scripttool.man | nroff -Tutf8 -man > man/man1/scripttool.1

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d man ]; then rm -fR man; fi

install:
	env GOBIN=$(PREFIX)/bin go install cmd/scripttool/scripttool.go

uninstall:
	if [ -f $(PREFIX)/bin/scripttool$(EXT) ]; then rm $(PREFIX)/bin/scripttool$(EXT); fi
	if [ -f $(PREFIX)/man/man1/scripttool.man ]; then rm $(PREFIX)/man/man1/scripttool.man; fi
	if [ -f $(PREFIX)/man/man1/scripttool.1 ]; then rm $(PREFIX)/man/man1/scripttool.1; fi

website: build
	./mk_website.py

status:
	git status

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

dist/linux-amd64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/scripttool cmd/scripttool/scripttool.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/scripttool.exe cmd/scripttool/scripttool.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/macos-amd64:
	mkdir -p dist/bin
	env GOOS=darwin	GOARCH=amd64 go build -o dist/bin/scripttool cmd/scripttool/scripttool.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/macos-arm64:
	mkdir -p dist/bin
	env GOOS=darwin	GOARCH=arm64 go build -o dist/bin/scripttool cmd/scripttool/scripttool.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-arm64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/scripttool cmd/scripttool/scripttool.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/linux-arm64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/scripttool cmd/scripttool/scripttool.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-arm64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

distribute_docs:
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	cp -v docs/scripttool.md dist/

release: distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macos-amd64 dist/macos-arm64 dist/raspbian-arm7 dist/linux-arm64

publish:
	./mk_website.py
	./publish.bash

