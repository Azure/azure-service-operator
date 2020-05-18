APP             = k8sinfra-gen
PACKAGE         = github.com/Azure/k8s-infra/hack/generator
DATE            ?= $(shell date +%FT%T%z)
VERSION         ?= $(shell git rev-list -1 HEAD)
SHORT_VERSION   ?= $(shell git rev-parse --short HEAD)
GOBIN           ?= $(HOME)/go/bin
GOFMT           = gofmt
GO              = go
PKGS            = $(or $(PKG),$(shell $(GO) list ./... | grep -vE "^$(PACKAGE)/templates/"))
ROOT_DIR        :=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
TOOLS_DIR       = $(abspath $(ROOT_DIR)/../tools)
TOOLS_BIN       = $(TOOLS_DIR)/bin
GOLINT          = $(TOOLS_BIN)/golint
GOX             = $(TOOLS_BIN)/gox
GO_INSTALL      = $(abspath $(ROOT_DIR)/../../scripts/go_install.sh)

V = 0
Q = $(if $(filter 1,$V),,@)

.PHONY: all
all: build test

$(GOLINT): ## Build golint
	GOBIN=$(TOOLS_BIN) $(GO_INSTALL) golang.org/x/lint/golint

$(GOX): ## Build gox
	GOBIN=$(TOOLS_BIN) $(GO_INSTALL)  github.com/mitchellh/gox@v1.0.1

build: tidy fmt vet lint ; $(info $(M) buiding ./bin/$(APP))
	$Q $(GO)  build -ldflags "-X $(PACKAGE)/cmd.GitCommit=$(VERSION)" -o ./bin/$(APP)

.PHONY: lint
lint: $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$(Q) $(GOLINT) -set_exit_status `go list ./... | grep -v /internal/`

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./...); do \
		$(GOFMT) -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

.PHONY: vet
vet: $(GOLINT) ; $(info $(M) running vet…) @ ## Run vet
	$Q $(GO) vet ./...

.PHONY: tidy
tidy: ; $(info $(M) running tidy…) @ ## Run tidy
	$Q $(GO) mod tidy

.PHONY: build-debug
build-debug: ; $(info $(M) buiding debug...)
	$Q $(GO)  build -o ./bin/$(APP) -tags debug

.PHONY: test
test: ; $(info $(M) running go test…)
	$(Q) $(GO) test ./... -tags=noexit

.PHONY: test-cover
test-cover: ; $(info $(M) running go test…)
	$(Q) $(GO) test -tags=noexit -race -covermode atomic -coverprofile=profile.cov ./...
	$(Q) $(TOOLS_BIN)/goveralls -coverprofile=profile.cov -service=github

.PHONY: gox
gox: install-tools
	$(Q) $(TOOLS_BIN)/gox -osarch="darwin/amd64 windows/amd64 linux/amd64" -ldflags "-X $(PACKAGE)/cmd.GitCommit=$(VERSION)" -output "./bin/$(SHORT_VERSION)/{{.Dir}}_{{.OS}}_{{.Arch}}"
	$(Q) tar -czvf ./bin/$(SHORT_VERSION)/$(APP)_darwin_amd64.tar.gz -C ./bin/$(SHORT_VERSION)/ $(APP)_darwin_amd64
	$(Q) tar -czvf ./bin/$(SHORT_VERSION)/$(APP)_linux_amd64.tar.gz -C ./bin/$(SHORT_VERSION)/ $(APP)_linux_amd64
	$(Q) tar -czvf ./bin/$(SHORT_VERSION)/$(APP)_windows_amd64.tar.gz -C ./bin/$(SHORT_VERSION)/ $(APP)_windows_amd64.exe

.PHONY: ci
ci: all