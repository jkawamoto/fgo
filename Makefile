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
	goxc -d=pkg -pv=$(VERSION) -os="darwin"

.PHONY: release
release:
	ghr  -u jkawamoto  v$(VERSIOM) pkg/$(VERSIOM)
