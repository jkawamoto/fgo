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

.PHONY: build
build:
	goxc -d={{.Dest}} -pv=$(VERSION)

.PHONY: release
release:
	ghr {{if .UserName}} -u {{.UserName}} {{end}} v$(VERSION) pkg/$(VERSION)
