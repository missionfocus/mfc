NAME := mfc
BINARY ?= bin/$(NAME)

export GOPRIVATE=git.missionfocus.com
export GOFLAGS=-mod=vendor

.PHONY: gitlab-release
gitlab-release:
	@scripts/gitlab-release.sh

.PHONY: build
build:
	CGO_ENABLED=0 go build -v -ldflags "-X main.version=$(CI_COMMIT_TAG)" -o $(BINARY) github.com/missionfocus.com/mfc/cmd/mfc
	openssl sha1 $(BINARY) > $(BINARY).checksum
