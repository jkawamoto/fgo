#
# Makefile
#
# Copyright (c) 2016 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
VERSION = snapshot

default: build

.PHONY: asset
asset:
	go-bindata -pkg command -o command/assets.go -nometadata assets

.PHONY: build
build: asset
	goxc -d=pkg -pv=$(VERSION) -os="darwin"

.PHONY: release
release:
	ghr  -u jkawamoto  v$(VERSION) pkg/$(VERSION)

.PHONY: get-deps
get-deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get -d -t -v .

.PHONY: test
test: asset
	go test -v ./...
