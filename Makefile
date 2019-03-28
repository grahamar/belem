COMMIT := $(shell git rev-parse --verify HEAD)
VERSION := $(shell git describe --tags --dirty --always)
BUILD_TIME := $(shell date +%FT%T%z)

ROOT_DIR := github.com/grahamar/belem/root
LDFLAGS := "-X ${ROOT_DIR}.Version=${VERSION} -X ${ROOT_DIR}.Commit=${COMMIT} -X ${ROOT_DIR}.BuildTime=${BUILD_TIME}"

# Test all packages.
.PHONY: test
test:
	@go test -cover ./...

# Clean build artifacts.
.PHONY: clean
clean:
	@git clean -f

.PHONY: deps
deps:
	@dep ensure

.PHONY: install
install:
	@echo "Installing version ${VERSION} of belem..."
	@go install -a -ldflags $(LDFLAGS) .
