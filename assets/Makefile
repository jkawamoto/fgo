#
# Makefile
#
VERSION = snapshot
GHRFLAGS =
.PHONY: build release

default: build

build:
	goxc -d={{.Dest}} -pv=$(VERSION)

release:
	ghr {{if .UserName}} -u {{.UserName}} {{end}} $(GHRFLAGS) v$(VERSION) {{.Dest}}/$(VERSION)
