#
# Simple Makefile
#
PROJECT = osf

VERSION = $(shell grep -m1 'Version = `' $(PROJECT).go | cut -d\`  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif

build: bin/osf2txt$(EXT) bin/fadein2osf$(EXT) bin/txt2osf$(EXT)

bin/osf2txt$(EXT): osf.go cmd/osf2txt/osf2txt.go 
	go build -o bin/osf2txt$(EXT) cmd/osf2txt/osf2txt.go

bin/fadein2osf$(EXT): osf.go cmd/fadein2osf/fadein2osf.go
	go build -o bin/fadein2osf$(EXT) cmd/fadein2osf/fadein2osf.go

bin/txt2osf$(EXT): osf.go cmd/txt2osf/txt2osf.go
	go build -o bin/txt2osf$(EXT) cmd/txt2osf/txt2osf.go

test:
	go test

man: build
	mkdir -p man/man1
	bin/osf2txt -generate-manpage | nroff -Tutf8 -man > man/man1/osf2txt.1
	bin/fadein2osf -generate-manpage | nroff -Tutf8 -man > man/man1/fadein2osf.1
	bin/txt2osf -generate-manpage | nroff -Tutf8 -man > man/man1/txt2osf.1

install:
	env GOBIN=$(HOME)/bin go install cmd/osf2txt/osf2txt.go
	env GOBIN=$(HOME)/bin go install cmd/fadein2osf/fadein2osf.go
	env GOBIN=$(HOME)/bin go install cmd/txt2osf/txt2osf.go

dist/linux-amd64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/osf2txt cmd/osf2txt/osf2txt.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/fadein2osf cmd/fadein2osf/fadein2osf.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/txt2osf cmd/txt2osf/txt2osf.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/osf2txt.exe cmd/osf2txt/osf2txt.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/fadein2osf.exe cmd/fadein2osf/fadein2osf.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/txt2osf.exe cmd/txt2osf/txt2osf.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/macos-amd64:
	mkdir -p dist/bin
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/osf2txt cmd/osf2txt/osf2txt.go
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/fadein2osf cmd/fadein2osf/fadein2osf.go
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/txt2osf cmd/txt2osf/txt2osf.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/macos-arm64:
	mkdir -p dist/bin
	env GOOS=darwin GOARCH=arm64 go build -o dist/bin/osf2txt cmd/osf2txt/osf2txt.go
	env GOOS=darwin GOARCH=arm64 go build -o dist/bin/fadein2osf cmd/fadein2osf/fadein2osf.go
	env GOOS=darwin GOARCH=arm64 go build -o dist/bin/txt2osf cmd/txt2osf/txt2osf.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-arm64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/osf2txt cmd/osf2txt/osf2txt.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/fadein2osf cmd/fadein2osf/fadein2osf.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/txt2osf cmd/txt2osf/txt2osf.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/linux-arm64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/osf2txt cmd/osf2txt/osf2txt.go
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/fadein2osf cmd/fadein2osf/fadein2osf.go
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/txt2osf cmd/txt2osf/txt2osf.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-arm64.zip README.md LICENSE INSTSALL.md bin/*
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

