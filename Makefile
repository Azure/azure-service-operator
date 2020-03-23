SHELL := /bin/bash
.DEFAULT_GOAL:=build

timestamp := $(shell /bin/date "+%Y%m%d-%H%M%S")
REGISTRY ?= devigned
IMG ?= k8s-infra-contoller-dev:$(timestamp)
CRD_OPTIONS ?= "crd:crdVersions=v1"

KIND_CLUSTER_NAME ?= k8sinfra
KIND_CLUSTER_TOUCH := .$(KIND_CLUSTER_NAME).cluster
KIND_KUBECONFIG := $(HOME)/.kube/kind-$(KIND_CLUSTER_NAME)
TLS_CERT_PATH := pki/certs/tls.crt

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Directories.
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
TOOLS_DIR := hack/tools
TOOLS_BIN_DIR := $(TOOLS_DIR)/bin

# Binaries.
GOLANGCI_LINT := $(TOOLS_BIN_DIR)/golangci-lint
CONTROLLER_GEN := $(TOOLS_BIN_DIR)/controller-gen
CONVERSION_GEN := $(TOOLS_BIN_DIR)/conversion-gen
KUBECTL=$(TOOLS_BIN_DIR)/kubectl
KUBE_APISERVER=$(TOOLS_BIN_DIR)/kube-apiserver
ETCD=$(TOOLS_BIN_DIR)/etcd
KUBEBUILDER=$(TOOLS_BIN_DIR)/kubebuilder
CFSSL=$(TOOLS_BIN_DIR)/cfssl
CFSSLJSON=$(TOOLS_BIN_DIR)/cfssljson
MKBUNDLE=$(TOOLS_BIN_DIR)/mkbundle
KIND=$(TOOLS_BIN_DIR)/kind
KUSTOMIZE=$(TOOLS_BIN_DIR)/kustomize

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Testing
## --------------------------------------
.PHONY: test test-int test-covers
test test-int test-cover: export TEST_ASSET_KUBECTL = $(ROOT_DIR)/$(KUBECTL)
test test-int test-cover: export TEST_ASSET_KUBE_APISERVER = $(ROOT_DIR)/$(KUBE_APISERVER)
test test-int test-cover: export TEST_ASSET_ETCD = $(ROOT_DIR)/$(ETCD)

test: $(KUBECTL) $(KUBE_APISERVER) $(ETCD) fmt lint header-check manifests ## Run tests
	go test -v ./...

test-int: .env $(KUBECTL) $(KUBE_APISERVER) $(ETCD) fmt generate lint manifests ## Run integration tests
	# MUST be executed as single command, or env vars will not propagate to test execution
	. .env && go test -v ./... -tags integration

.env: ## create a service principal and save the identity to .env for use in integration tests (requries jq and az)
	./scripts/create_testing_creds.sh

test-cover: $(KUBECTL) $(KUBE_APISERVER) $(ETCD) generate lint manifests ## Run tests w/ code coverage (./cover.out)
	go test ./... -tags integration -coverprofile cover.out

$(KUBECTL) $(KUBE_APISERVER) $(ETCD) $(KUBEBUILDER): ## Install test asset kubectl, kube-apiserver, etcd
	. ./scripts/fetch_ext_bins.sh && fetch_tools

$(CFSSL): ## Install cfssl tool
	cd $(TOOLS_DIR); go build -tags=tools -o bin/cfssl github.com/cloudflare/cfssl/cmd/cfssl

$(CFSSLJSON): ## Install cfssljson tool
	cd $(TOOLS_DIR); go build -tags=tools -o bin/cfssljson github.com/cloudflare/cfssl/cmd/cfssljson

$(MKBUNDLE): ## Install mkbundle tool
	cd $(TOOLS_DIR); go build -tags=tools -o bin/mkbundle github.com/cloudflare/cfssl/cmd/mkbundle

$(KIND): ## Install kind tool
	cd $(TOOLS_DIR); GOBIN=$(ROOT_DIR)/$(TOOLS_BIN_DIR) go get -tags=tools sigs.k8s.io/kind@v0.7.0

$(KUSTOMIZE): ## Install kustomize
	cd $(TOOLS_DIR); GOBIN=$(ROOT_DIR)/$(TOOLS_BIN_DIR) go get -tags=tools sigs.k8s.io/kustomize/kustomize/v3

$(CONTROLLER_GEN): $(TOOLS_DIR)/go.mod ## Build controller-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o bin/controller-gen sigs.k8s.io/controller-tools/cmd/controller-gen

$(CONVERSION_GEN): $(TOOLS_DIR)/go.mod ## Build conversion-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o bin/conversion-gen k8s.io/code-generator/cmd/conversion-gen

$(GOLANGCI_LINT): $(TOOLS_DIR)/go.mod ## Build golangci-lint from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

## --------------------------------------
## Linting
## --------------------------------------

.PHONY: lint
lint: $(GOLANGCI_LINT) ## Lint codebase
	$(GOLANGCI_LINT) run -v --timeout 5m

.PHONY: lint-full
lint-full: $(GOLANGCI_LINT) ## Run slower linters to detect possible issues
	$(GOLANGCI_LINT) run -v --fast=false --timeout 5m

.PHONY: build
build: fmt ## Build manager binary
	go build -o bin/manager main.go

$(TLS_CERT_PATH): $(CFSSL) $(CFSSLJSON) $(MKBUNDLE) ## Generate local certificates so the webhooks will run
	./scripts/gen-certs.sh

.PHONY: run
run: $(KIND) $(KIND_CLUSTER_TOUCH)
run: export KUBECONFIG = $(KIND_KUBECONFIG)
run: export ENVIRONMENT = development
run: $(TLS_CERT_PATH) generate fmt manifests install ## Run a development cluster using kind
	go run ./main.go

$(KIND_CLUSTER_TOUCH): $(KIND) $(KUBECTL)
	$(KIND) create cluster --name=$(KIND_CLUSTER_NAME) --kubeconfig=$(KIND_KUBECONFIG) --image=kindest/node:v1.16.4
	touch $(KIND_CLUSTER_TOUCH)

.PHONY: apply-certs-and-secrets
apply-certs-and-secrets: $(KUBECTL)
	./scripts/apply_cert_and_secrets.sh

.PHONY: kind-reset
kind-reset: $(KIND) ## Destroys the "k8sinfra" kind cluster.
	$(KIND) delete cluster --name=$(KIND_CLUSTER_NAME) || true
	rm -f $(KIND_CLUSTER_TOUCH)

.PHONY: install
install: manifests $(KUBECTL) $(KUSTOMIZE) ## Install CRDs into a cluster
	$(KUSTOMIZE) build config/crd | $(KUBECTL) apply -f -

.PHONY: uninstall
uninstall: manifests $(KUBECTL) $(KUSTOMIZE) ## Uninstall CRDs from a cluster
	$(KUSTOMIZE) build config/crd | $(KUBECTL) delete -f -

.PHONY: deploy
deploy: manifests $(KUBECTL) $(KUSTOMIZE) docker-build docker-push ## Deploy controller in the configured Kubernetes cluster in ~/.kube/config
	cd config/manager && $(ROOT_DIR)/$(TOOLS_BIN_DIR)/kustomize edit set image controller=$(REGISTRY)/${IMG}
	$(KUSTOMIZE) build config/default | $(KUBECTL) apply -f -

.PHONY: deploy-kind
deploy-kind: $(KIND) $(KIND_CLUSTER_TOUCH)
deploy-kind: export KUBECONFIG = $(KIND_KUBECONFIG)
deploy-kind: apply-certs-and-secrets deploy

.PHONY: manifests
manifests: $(CONTROLLER_GEN) ## Generate manifests e.g. CRD, RBAC etc.
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

.PHONY: fmt
fmt: ## Run go fmt against code
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code
	go vet ./...

.PHONY: header-check
header-check:
	./scripts/verify_boilerplate.sh

.PHONY: generate
generate: $(CONTROLLER_GEN) $(CONVERSION_GEN) ## Generate code
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths="./..."

	$(CONVERSION_GEN) \
    		--input-dirs=./apis/microsoft.network/v20191101,./apis/microsoft.resources/v20191001,./apis/microsoft.resources/v20150101 \
    		--output-file-base=zz_generated.conversion \
    		--output-base=$(ROOT_DIR) \
    		--go-header-file=./hack/boilerplate.go.txt

.PHONY: docker-build
docker-build: test ## Build the docker image
	docker build . -t $(REGISTRY)/${IMG}

.PHONY: docker-push
docker-push: ## Push the docker image
	docker push $(REGISTRY)/${IMG}

.PHONY: dist
dist:
	mkdir -p dist
	cd config/manager && $(ROOT_DIR)/$(TOOLS_BIN_DIR)/kustomize edit set image controller=$(REGISTRY)/${IMG}
	$(KUSTOMIZE) build config/default > dist/release.yaml

.PHONY: release
release: dist docker-build docker-push $(KUSTOMIZE) ## Build, push, generate dist for release