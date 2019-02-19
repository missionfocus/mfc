NAME := mf-vault
VERSION := $(shell git describe --always --tags --abbrev=0)
BINARY ?= out/$(NAME)

.PHONY: gitlab-release
gitlab-release:
	@scripts/gitlab-release.sh

.PHONY: build
build:
	CGO_ENABLED=0 go build -o $(BINARY)
	openssl sha1 $(BINARY) > $(BINARY).checksum
