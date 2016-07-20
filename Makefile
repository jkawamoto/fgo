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
	go-bindata -pkg command -o command/assets.go assets

.PHONY: build
build:
	goxc -d=pkg -pv=$(VERSION) -os="darwin"

.PHONY: release
release:
	ghr  -u jkawamoto  v$(VERSION) pkg/$(VERSION)

.PHONY: get-deps
get-deps:
	go get github.com/tcnksm/go-gitconfig
	go get github.com/ttacon/chalk
	go get github.com/urfave/cli
