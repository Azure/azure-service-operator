SHELL           = /bin/bash
APP             = k8sinfra-gen
PACKAGE         = github.com/Azure/k8s-infra/hack/generator

include ../../tools.mk

DATE            ?= $(shell date +%FT%T%z)
VERSION         ?= $(shell git rev-list -1 HEAD)
SHORT_VERSION   ?= $(shell git rev-parse --short HEAD)
TREE_STATE      ?= $(if $(shell git status --porcelain),dirty,clean)
CRD_OPTIONS     ?= "crd:crdVersions=v1,allowDangerousTypes=true"
GO_DIRS         := $(shell $(GO) list -f '{{.Dir}}' ./...)

V = 0
Q = $(if $(filter 1,$V),,@)

.PHONY: all
all: header-check fmt build test

.PHONY: manifests
manifests: $(CONTROLLER_GEN)
	$(Q) $(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./apis/..." output:crd:artifacts:config=config/crd/bases

# TODO: This doesn't work right now because the manifests target fails
.PHONY: generate
generate: manifests $(CONTROLLER_GEN) $(CONVERSION_GEN) ## Generate code
	$(Q) $(CONTROLLER_GEN) object:headerFile=../boilerplate.go.txt paths="./..."

	$(Q) $(CONVERSION_GEN) \
		--input-dirs="./apis/..." \
		--output-file-base=zz_generated.conversion \
		--output-base=./ \
		--go-header-file=../boilerplate.go.txt

.PHONY: gen-arm
gen-arm: rm-apis build ; $(info $(M) running k8sinfra-gen gen…) @ ## Running gen
	$(Q) ./bin/$(APP) gen azure-arm.yaml
	# Run with verbosity 4 to see a stack trace for errors.
	# $(Q) ./bin/$(APP) gen azure-arm.yaml -v 4

	$(Q) $(CONTROLLER_GEN) object:headerFile=../boilerplate.go.txt paths="./..."

# TODO: We should remove this target once we place the generated code outside of the generator.
# TODO: The generator already cleans up the APIs directory but since right now the generated code
# TODO: doesn't build we need to rm-apis before build.
.PHONY: rm-apis
rm-apis: ; $(info $(M) removing previously generated apis…)
	$(Q) rm -rf apis

.PHONY: build
build: tidy lint ; $(info $(M) building ./bin/$(APP))
	$(Q) $(GO) build -ldflags "-X $(PACKAGE)/pkg/codegen.GitCommit=$(SHORT_VERSION) -X $(PACKAGE)/pkg/codegen.GitTreeState=$(TREE_STATE)" -o ./bin/$(APP)

.PHONY: lint
lint: $(GOLANGCI_LINT) ; $(info $(M) running golangci configured linters…) ## Lint codebase
	$(Q) $(GOLANGCI_LINT) run -v

.PHONY: lint-full
lint-full: $(GOLANGCI_LINT) ; $(info $(M) running golangci configured linters…) ## Run slower linters to detect possible issues
	$(Q) $(GOLANGCI_LINT) run -v --fast=false

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	$(Q) $(GOFMT) -l -s -w $(GO_DIRS)

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
test-cover: $(GCOV2LCOV) ; $(info $(M) running go test…)
	$(Q) $(GO) test -tags=noexit -race -covermode atomic -coverprofile=cover.out ./...
	$(Q) $(GCOV2LCOV) -infile cover.out -outfile coverage.lcov

.PHONY: gox
gox: install-tools
	$(Q) $(TOOLS_BIN)/gox -osarch="darwin/amd64 windows/amd64 linux/amd64" -ldflags "-X $(PACKAGE)/cmd.GitCommit=$(VERSION)" -output "./bin/$(SHORT_VERSION)/{{.Dir}}_{{.OS}}_{{.Arch}}"
	$(Q) tar -czvf ./bin/$(SHORT_VERSION)/$(APP)_darwin_amd64.tar.gz -C ./bin/$(SHORT_VERSION)/ $(APP)_darwin_amd64
	$(Q) tar -czvf ./bin/$(SHORT_VERSION)/$(APP)_linux_amd64.tar.gz -C ./bin/$(SHORT_VERSION)/ $(APP)_linux_amd64
	$(Q) tar -czvf ./bin/$(SHORT_VERSION)/$(APP)_windows_amd64.tar.gz -C ./bin/$(SHORT_VERSION)/ $(APP)_windows_amd64.exe

.PHONY: ci
ci: build test-cover

.PHONY: diagrams
diagrams: $(patsubst %.dot,%.png,$(wildcard spec/images/*.dot))

%.png: %.dot 
	$(info $(M) Generating diagram $@ from $< …)
	$(Q) dot $< -Tpng -o $@

