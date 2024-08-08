GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
SOURCES := $(shell find . -type f  -name '*.go')

GIT_VERSION ?= $(shell git describe --tags --dirty)
GIT_COMMIT_HASH ?= $(shell git rev-parse HEAD)
GIT_TREESTATE = "clean"
GIT_DIFF = $(shell git diff --quiet >/dev/null 2>&1; if [ $$? -eq 1 ]; then echo "1"; fi)
ifeq ($(GIT_DIFF), 1)
    GIT_TREESTATE = "dirty"
endif

BUILD_DATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
OUT_PATH = out/$(GOOS)-$(GOARCH)

LDFLAGS := "-X github.com/zirain/ubrain/pkg/version.gitVersion=$(GIT_VERSION) \
			-X github.com/zirain/ubrain/pkg/version.gitCommit=$(GIT_COMMIT_HASH) \
			-X github.com/zirain/ubrain/pkg/version.gitTreeState=$(GIT_TREESTATE) \
			-X github.com/zirain/ubrain/pkg/version.buildDate=$(BUILD_DATE)"

.PHONY: ubrain
ubrain: clean
	CGO_ENABLED=0 GOOS=$(GOOS) go build \
		-ldflags $(LDFLAGS) \
		-o $(OUT_PATH)/ubrainctl \
		cmd/ubrain/main.go

.PHONY: clean
clean:
	rm -rf $(OUT_PATH)