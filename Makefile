BINARY := aws-assume-run
VERSION ?= v1.0.0
PLATFORMS := linux darwin

os = $(word 1, $@)

.PHONY: deps
deps:
	dep ensure

.PHONY: $(PLATFORMS)
$(PLATFORMS): deps
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64
	sha256sum release/$(BINARY)-$(VERSION)-$(os)-amd64 > release/checksum-$(VERSION)-$(os).txt

.PHONY: release
release: linux darwin
