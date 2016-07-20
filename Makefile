VERSION = snapshot

default: build

.PHONY: build
build:
	goxc -d=pkg -pv=$(VERSION) -os="darwin"

.PHONY: release
release:
	ghr  -u jkawamoto  v$(VERSIOM) pkg/$(VERSIOM)
