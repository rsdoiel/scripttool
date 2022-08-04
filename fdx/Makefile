#
# Simple Makefile
#
PROJECT = fdx

VERSION = $(shell grep -m1 'Version = `' $(PROJECT).go | cut -d\`  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif


build: bin/fdx2txt$(EXT) bin/txt2fdx$(EXT)

bin/fdx2txt$(EXT): fdx.go fromfountain.go cmd/fdx2txt/fdx2txt.go
	go build -o bin/fdx2txt$(EXT) cmd/fdx2txt/fdx2txt.go

bin/txt2fdx$(EXT): fdx.go fromfountain.go cmd/txt2fdx/txt2fdx.go
	go build -o bin/txt2fdx$(EXT) cmd/txt2fdx/txt2fdx.go

install:
	env GOBIN=$(HOME)/bin go install cmd/fdx2txt/fdx2txt.go
	env GOBIN=$(HOME)/bin go install cmd/txt2fdx/txt2fdx.go

man: build
	mkdir -p man/man1
	bin/fdx2txt$(EXT) -generate-manpage | nroff -Tutf8 -man > man/man1/fdx2txt.1
	bin/txt2fdx$(EXT) -generate-manpage | nroff -Tutf8 -man > man/man1/txt2fdx.1
	
test:
	go test

dist/linux-amd64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/fdx2txt cmd/fdx2txt/fdx2txt.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/txt2fdx cmd/txt2fdx/txt2fdx.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/fdx2txt.exe cmd/fdx2txt/fdx2txt.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/txt2fdx.exe cmd/txt2fdx/txt2fdx.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/macos-amd64:
	mkdir -p dist/bin
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/fdx2txt cmd/fdx2txt/fdx2txt.go
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/txt2fdx cmd/txt2fdx/txt2fdx.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/macos-arm64:
	mkdir -p dist/bin
	env GOOS=darwin GOARCH=arm64 go build -o dist/bin/fdx2txt cmd/fdx2txt/fdx2txt.go
	env GOOS=darwin GOARCH=arm64 go build -o dist/bin/txt2fdx cmd/txt2fdx/txt2fdx.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macos-arm64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/fdx2txt cmd/fdx2txt/fdx2txt.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/txt2fdx cmd/txt2fdx/txt2fdx.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/linux-arm64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/fdx2txt cmd/fdx2txt/fdx2txt.go
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/txt2fdx cmd/txt2fdx/txt2fdx.go
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

