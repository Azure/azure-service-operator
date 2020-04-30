APP				= k8sinfra
PACKAGE  		= github.com/Azure/k8s-infra/hack/generator
DATE    		?= $(shell date +%FT%T%z)
VERSION 		?= $(shell git rev-list -1 HEAD)
SHORT_VERSION 	?= $(shell git rev-parse --short HEAD)
GOBIN      		?= $(HOME)/go/bin
GOFMT   		= gofmt
GO      		= go
PKGS     		= $(or $(PKG),$(shell $(GO) list ./... | grep -vE "^$(PACKAGE)/templates/"))
TOOLS_DIR		= ../tools
TOOLSBIN		= $(TOOLS_DIR)/bin
GOLINT			= $(TOOLSBIN)/golint
GOX				= $(TOOLSBIN)/gox

V = 0
Q = $(if $(filter 1,$V),,@)

.PHONY: all
all: install-tools fmt lint vet tidy build

install-tools: $(GOLINT) $(GOX); $(info $(M) installing tools…)

$(GOLINT): $(TOOLS_DIR)/go.mod ## Build golint
	cd $(TOOLS_DIR); go build -tags=tools -o bin/golint golang.org/x/lint/golint

$(GOX): $(TOOLS_DIR)/go.mod ## Build gox
	cd $(TOOLS_DIR); go build -tags=tools -o bin/gox github.com/mitchellh/gox

build: lint tidy ; $(info $(M) buiding ./bin/$(APP))
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
	$(Q) $(TOOLSBIN)/goveralls -coverprofile=profile.cov -service=github

.PHONY: gox
gox: install-tools
	$(Q) $(TOOLSBIN)/gox -osarch="darwin/amd64 windows/amd64 linux/amd64" -ldflags "-X $(PACKAGE)/cmd.GitCommit=$(VERSION)" -output "./bin/$(SHORT_VERSION)/{{.Dir}}_{{.OS}}_{{.Arch}}"
	$(Q) tar -czvf ./bin/$(SHORT_VERSION)/$(APP)_darwin_amd64.tar.gz -C ./bin/$(SHORT_VERSION)/ $(APP)_darwin_amd64
	$(Q) tar -czvf ./bin/$(SHORT_VERSION)/$(APP)_linux_amd64.tar.gz -C ./bin/$(SHORT_VERSION)/ $(APP)_linux_amd64
	$(Q) tar -czvf ./bin/$(SHORT_VERSION)/$(APP)_windows_amd64.tar.gz -C ./bin/$(SHORT_VERSION)/ $(APP)_windows_amd64.exe

.PHONY: ci
ci: install-tools fmt lint vet tidy build test-cover