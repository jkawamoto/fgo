#
# Makefile
#
# Copyright (c) 2016-2018 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
VERSION = snapshot
GHRFLAGS =
.PHONY: asset build release get-deps test

default: build

asset:
	go-bindata -pkg fgo -o fgo/assets.go -nometadata assets

build: asset
	goxc -d=pkg -pv=$(VERSION) -os="darwin linux"

release:
	ghr -u jkawamoto $(GHRFLAGS) $(VERSION) pkg/$(VERSION)

get-deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get -d -t -v .

test: asset
	go test -v ./...
