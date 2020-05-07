NAME := mf-vault
BINARY ?= out/$(NAME)

export GOPRIVATE=git.missionfocus.com
export GOFLAGS=-mod=vendor

.PHONY: gitlab-release
gitlab-release:
	@scripts/gitlab-release.sh

.PHONY: build
build:
	CGO_ENABLED=0 go build -v -ldflags "-X git.missionfocus.com/ours/code/tools/mfc/cmd/mfc.version=$(CI_COMMIT_TAG)" -o $(BINARY)
	openssl sha1 $(BINARY) > $(BINARY).checksum
