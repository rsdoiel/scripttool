#
# Simple Makefile
#
PROJECT = fountain

VERSION = $(shell grep -m1 'Version = `' $(PROJECT).go | cut -d\`  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif

build:  fountain.go cmd/fountainfmt/fountainfmt.go cmd/fountain2html/fountain2html.go
	go build -o bin/fountainfmt$(EXT) cmd/fountainfmt/fountainfmt.go
	go build -o bin/fountain2json$(EXT) cmd/fountain2json/fountain2json.go
	go build -o bin/fountain2html$(EXT) cmd/fountain2html/fountain2html.go

test:
	go test

man: build
	mkdir -p man/man1
	bin/fountainfmt -generate-manpage | nroff -Tutf8 -man > man/man1/fountainfmt.1
	bin/fountain2json -generate-manpage | nroff -Tutf8 -man > man/man1/fountain2json.1
	bin/fountain2html -generate-manpage | nroff -Tutf8 -man > man/man1/fountain2html.1

install:
	env GOBIN=$(HOME)/bin go install cmd/fountainfmt/fountainfmt.go
	env GOBIN=$(HOME)/bin go install cmd/fountain2json/fountain2json.go
	env GOBIN=$(HOME)/bin go install cmd/fountain2html/fountain2html.go

fetch_scrippets_css:
	if [ ! -f css/scrippets.css ]; then \
		curl -L -o css/scrippets.css \
		"http://johnaugust.com/wp-content/plugins/wp-scrippets/scrippets.css?v2.0"; fi


dist/linux-amd64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/fountainfmt cmd/fountainfmt/fountainfmt.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/fountain2json cmd/fountain2json/fountain2json.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/fountain2html cmd/fountain2html/fountain2html.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/fountainfmt.exe cmd/fountainfmt/fountainfmt.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/fountain2json.exe cmd/fountain2json/fountain2json.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/fountain2html.exe cmd/fountain2html/fountain2html.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/macos-amd64:
	mkdir -p dist/bin
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/fountainfmt cmd/fountainfmt/fountainfmt.go
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/fountain2json cmd/fountain2json/fountain2json.go
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/fountain2html cmd/fountain2html/fountain2html.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/macos-arm64:
	mkdir -p dist/bin
	env GOOS=darwin GOARCH=arm64 go build -o dist/bin/fountainfmt cmd/fountainfmt/fountainfmt.go
	env GOOS=darwin GOARCH=arm64 go build -o dist/bin/fountain2json cmd/fountain2json/fountain2json.go
	env GOOS=darwin GOARCH=arm64 go build -o dist/bin/fountain2html cmd/fountain2html/fountain2html.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-arm64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/fountainfmt cmd/fountainfmt/fountainfmt.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/fountain2json cmd/fountain2json/fountain2json.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/fountain2html cmd/fountain2html/fountain2html.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/linux-arm64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/fountainfmt cmd/fountainfmt/fountainfmt.go
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/fountain2json cmd/fountain2json/fountain2json.go
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/fountain2html cmd/fountain2html/fountain2html.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-arm64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

distribute_docs:
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/

release: distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macos-amd64 dist/macos-arm64 dist/raspbian-arm7 dist/linux-arm64

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d man ]; then rm -fR man; fi

website:
	./mk_website.py

status:
	git status

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

publish:
	./mk_website.py
	./publish.bash

