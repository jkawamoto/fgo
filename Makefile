#
# Makefile
#
# Copyright (c) 2016-2019 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
VERSION = snapshot
GHRFLAGS =
.PHONY: assets build release get-deps test

default: build

fgo/assets/bindata.go:
	go generate fgo/assets.go


build: fgo/assets/bindata.go
	goxc -d=pkg -pv=$(VERSION) -os="darwin linux"

release:
	ghr -u jkawamoto $(GHRFLAGS) $(VERSION) pkg/$(VERSION)

get-deps:
	go get -u github.com/jessevdk/go-assets-builder
	go get -d -t -v .

test: fgo/assets/bindata.go
	go test -v ./...
