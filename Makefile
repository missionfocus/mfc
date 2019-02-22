NAME := mf-vault
BINARY ?= out/$(NAME)

.PHONY: gitlab-release
gitlab-release:
	@scripts/gitlab-release.sh

.PHONY: build
build:
	CGO_ENABLED=0 go build -v -ldflags "-X git.missionfocus.com/open-source/mf-vault/cmd/mf-vault.version=$(CI_COMMIT_TAG)" -o $(BINARY)
	openssl sha1 $(BINARY) > $(BINARY).checksum
