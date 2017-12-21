#
# Simple Makefile
#
PROJECT = scripttools

VERSION = $(shell grep -m1 'Version = ' $(PROJECT).go | cut -d\"  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif

build:
	go build -o bin/fdx2fountain$(EXT) cmd/fdx2fountain/fdx2fountain.go
	go build -o bin/fountain2fdx$(EXT) cmd/fountain2fdx/fountain2fdx.go

test:
	go test

clean: 
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

install:
	env GOBIN=$(HOME)/bin go install cmd/fdx2fountain/fdx2fountain.go
	env GOBIN=$(HOME)/bin go install cmd/fountain2fdx/fountain2fdx.go

uninstall:
	if [ -f $(GOBIN)/fdx2fountain$(EXT) ]; then rm $(GOBIN)/fdx2fountain$(EXT); fi
	if [ -f $(GOBIN)/fountain2fdx$(EXT) ]; then rm $(GOBIN)/fountain2fdx$(EXT); fi

website: build
	./mk-website.bash

status:
	git status

save:
	git commit -am "Quick Save"
	git push origin $(BRANCH)

dist/linux-amd64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/fdx2fountain cmd/fdx2fountain/fdx2fountain.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/fountain2fdx cmd/fountain2fdx/fountain2fdx.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/fountain2fdx.exe cmd/fountain2fdx/fountain2fdx.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/fdx2fountain.exe cmd/fdx2fountain/fdx2fountain.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env GOOS=darwin	GOARCH=amd64 go build -o dist/bin/fdx2fountain cmd/fdx2fountain/fdx2fountain.go
	env GOOS=darwin	GOARCH=amd64 go build -o dist/bin/fountain2fdx cmd/fountain2fdx/fountain2fdx.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/fdx2fountain cmd/fdx2fountain/fdx2fountain.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/fountain2fdx cmd/fountain2fdx/fountain2fdx.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

dist/linux-arm64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/fdx2fountain cmd/fdx2fountain/fdx2fountain.go
	env GOOS=linux GOARCH=arm64 go build -o dist/bin/fountain2fdx cmd/fountain2fdx/fountain2fdx.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-arm64.zip README.md LICENSE INSTSALL.md bin/*
	rm -fR dist/bin

distribute_docs:
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	cp -v docs/fdx2fountain.md dist/
	cp -v docs/fountain2fdx.md dist/

release: distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7 dist/linux-arm64

publish:
	./mk-website.bash
	./publish.bash

