## Windows
# GNU Make 3.81
# Copyright (C) 2006  Free Software Foundation, Inc.
# This is free software; see the source for copying conditions.
# There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A
# PARTICULAR PURPOSE.
#
# This program built for i386-pc-mingw32

## macOS
# GNU Make 3.81
# Copyright (C) 2006  Free Software Foundation, Inc.
# This is free software; see the source for copying conditions.
# There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A
# PARTICULAR PURPOSE.
#
# This program built for i386-apple-darwin11.3.0

## Linux
# GNU Make 4.1
# Built for x86_64-pc-linux-gnu
# Copyright (C) 1988-2014 Free Software Foundation, Inc.
# License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
# This is free software: you are free to change and redistribute it.
# There is NO WARRANTY, to the extent permitted by law.

# https://stackoverflow.com/questions/714100
ifeq ($(OS),Windows_NT)
    GOPATH=C:\Golang\gopath
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        GOPATH=$(shell env | grep GOPATH | cut -d'=' -f 2)
    endif
    ifeq ($(UNAME_S),Darwin)
        GOPATH=$(shell env | grep GOPATH | cut -d'=' -f 2)
    endif
endif

# constant
APP=01
SOURCE=./...

# change env variables
export GO111MODULE=on

# common targets
test:
	@echo testing
	go vet $(SOURCE)
	go test -race -cover -timeout 90s $(SOURCE)

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out > coverage.ana

clean:
	rm -rf bin pkg
	go clean --modcache
	docker system prune -f

vendor:
	go build -v $(SOURCE) 
	go mod tidy
	go mod vendor

build:
	go install -mod=vendor -i -v $(SOURCE)

## cobra command - parse
parseg: build
	$(GOPATH)/bin/$(APP) parse --indent false --command generate

parsev: build
	$(GOPATH)/bin/$(APP) parse --indent false --command validate

## cobra command - selenium
selenium: build
	$(GOPATH)/bin/$(APP) selenium