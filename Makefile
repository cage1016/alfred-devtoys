ENABLED_AUTO_UPDATE ?= "false"
LDFLAGS ?= -X github.com/cage1016/alfred-devtoys/cmd.EnabledAutoUpdate=$(ENABLED_AUTO_UPDATE)

HAVE_GO_BINDATA := $(shell command -v go-bindata 2> /dev/null)
generate: ## go generate
ifndef HAVE_GO_BINDATA
	@echo "requires 'go-bindata' (go get -u github.com/kevinburke/go-bindata/go-bindata)"
	@exit 1 # fail	
else
	go generate ./...
endif

.PHONY: test
test: ## run tests
	go test -v -race -cover -coverprofile coverage.txt -covermode=atomic ./...

.PHONY: build
build: generate ## build the binary
	ak alfred build -l "$(LDFLAGS)"

.PHONY: info
info: ## show alfred workflow info
	ak alfred info

.PHONY: link
link: ## link alfred workflow
	ak alfred link

.PHONY: unlink
unlink: ## unlink alfred workflow
	ak alfred unlink

.PHONY: pack
pack: ## pack alfred workflow
	ak alfred pack

.PHONY: help
help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_0-9-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help