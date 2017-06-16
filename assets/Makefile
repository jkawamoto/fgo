#
# Makefile
#
VERSION = snapshot
GHRFLAGS =

default: build

.PHONY: build
build:
	goxc -d={{.Dest}} -pv=$(VERSION)

.PHONY: release
release:
	ghr {{if .UserName}} -u {{.UserName}} {{end}} $(GHRFLAGS) v$(VERSION) {{.Dest}}/$(VERSION)
