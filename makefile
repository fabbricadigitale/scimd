SHELL := /bin/bash

# The name of the executable (default is the current directory name)
TARGET := $(shell echo $${PWD\#\#*/})
.DEFAULT_GOAL: $(TARGET)

ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

# All packages
LIST := $(shell go list ./... | grep -v /vendor/)
# Main package
REPO := $(word 1,$(LIST))

# Source files (ignoring vendor directory)
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Target variables
VERSION = $(shell git name-rev --name-only --tags --no-undefined HEAD 2>/dev/null | sed -n 's/^\([^^~]\{1,\}\)\(\^0\)\{0,1\}$$/\1/p')
COMMIT = $(shell git rev-parse HEAD)
BRANCH = $(shell git rev-parse --abbrev-ref HEAD) 
SUMMARY = $(shell git describe --tags --dirty --always)
ISO8601DATE = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS = "-X=$(REPO)/cmd.version=$(VERSION) -X=$(REPO)/cmd.commit=$(COMMIT) -X=$(REPO)/cmd.branch=$(BRANCH) -X=$(REPO)/cmd.summary=$(SUMMARY) -X=$(REPO)/cmd.date=$(ISO8601DATE)"

CLEANBRANCH := $(shell echo $(BRANCH) | sed -e "s/[^[:alnum:]]/-/g")
tag ?= $(CLEANBRANCH)

.PHONY: all build clean install uninstall format simplify check docs image

all: check install

$(TARGET): $(SRC)
	@go build -ldflags $(LDFLAGS) -o $(TARGET)

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

install:
	@go install -ldflags $(LDFLAGS)

uninstall: clean
	@rm -f $$(which ${TARGET})

format:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

check:
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make format'"
	@for d in $(LIST); do golint $${d}; done
	@go tool vet $(SRC)

docs/cli:
	@mkdir -p $(addprefix $(ROOT_DIR),$@)

docs: docs/cli
	@go run $(addprefix $(ROOT_DIR),cmd/docs/main.go) $(ROOT_DIR)

image:
	docker build --build-arg LDFLAGS=$(LDFLAGS) -t "$(REPO:github.com/%=%):$(tag)" .