VERSION = snapshot

default: build

.PHONY: build
build:
	goxc -d={{.Dest}} -pv=$(VERSION)

.PHONY: release
release:
	ghr {{if .UserName}} -u {{.UserName}} {{end}} v$(VERSIOM) pkg/$(VERSIOM)
