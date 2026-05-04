# Build flags aligned with .goreleaser.yml (-s -w + prometheus/version stamping).
VERSION  ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo unknown)
REVISION ?= $(shell git rev-parse HEAD 2>/dev/null || echo unknown)
BRANCH   ?= $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null || echo unknown)
DATE     ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
BUILDUSER ?= $(shell id -un)@$(shell hostname 2>/dev/null || echo localhost)

LDFLAGS := -s -w \
	-X github.com/prometheus/common/version.Version=$(VERSION) \
	-X github.com/prometheus/common/version.Revision=$(REVISION) \
	-X github.com/prometheus/common/version.Branch=$(BRANCH) \
	-X github.com/prometheus/common/version.BuildUser=$(BUILDUSER) \
	-X github.com/prometheus/common/version.BuildDate=$(DATE)

.PHONY: build
build:
	CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o prometheus-lvm-exporter .

.PHONY: clean
clean:
	rm -f prometheus-lvm-exporter
