SHELL := /bin/bash
.DEFAULT_GOAL:=build

timestamp := $(shell /bin/date "+%Y%m%d-%H%M%S")
# Image URL to use all building/pushing image targets
IMG ?= k8s-infra-contoller:$(timestamp)
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"

KIND_CLUSTER_NAME ?= k8s-infra

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
KUBECTL=$(TOOLS_BIN_DIR)/kubectl
KUBE_APISERVER=$(TOOLS_BIN_DIR)/kube-apiserver
ETCD=$(TOOLS_BIN_DIR)/etcd

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Testing
## --------------------------------------
.PHONY: test test-int test-covers
test test-int test-cover: export TEST_ASSET_KUBECTL = $(ROOT_DIR)/$(KUBECTL)
test test-int test-cover: export TEST_ASSET_KUBE_APISERVER = $(ROOT_DIR)/$(KUBE_APISERVER)
test test-int test-cover: export TEST_ASSET_ETCD = $(ROOT_DIR)/$(ETCD)

test: $(KUBECTL) $(KUBE_APISERVER) $(ETCD) fmt generate lint manifests ## Run tests
	go test -v ./...

test-int: .env $(KUBECTL) $(KUBE_APISERVER) $(ETCD) fmt generate lint manifests ## Run integration tests
	# MUST be executed as single command, or env vars will not propagate to test execution
	. .env && go test -v ./... -tags integration

.env: ## create a service principal and save the identity to .env for use in integration tests (requries jq and az)
	./scripts/create_testing_creds.sh

test-cover: $(KUBECTL) $(KUBE_APISERVER) $(ETCD) generate lint manifests ## Run tests w/ code coverage (./cover.out)
	go test ./... -tags integration -coverprofile cover.out

$(KUBECTL) $(KUBE_APISERVER) $(ETCD): ## Install test asset kubectl, kube-apiserver, etcd
	. ./scripts/fetch_ext_bins.sh && fetch_tools

$(CONTROLLER_GEN): $(TOOLS_DIR)/go.mod ## Build controller-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o bin/controller-gen sigs.k8s.io/controller-tools/cmd/controller-gen

$(GOLANGCI_LINT): $(TOOLS_DIR)/go.mod ## Build golangci-lint from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

## --------------------------------------
## Linting
## --------------------------------------

.PHONY: lint
lint: $(GOLANGCI_LINT) ## Lint codebase
	$(GOLANGCI_LINT) run -v --timeout 5m

lint-full: $(GOLANGCI_LINT) ## Run slower linters to detect possible issues
	$(GOLANGCI_LINT) run -v --fast=false --timeout 5m

build: generate fmt ## Build manager binary
	go build -o bin/manager main.go

run: export KUBECONFIG = $(shell kind get kubeconfig-path --name="k8s-infra")
run: .k8s-infra.cluster generate fmt manifests install ## Run a development cluster using kind
	go run ./main.go

.k8s-infra.cluster:
	kind create cluster --name=$(KIND_CLUSTER_NAME) --image=kindest/node:v1.16.2
	touch .$(KIND_CLUSTER_NAME).cluster

.PHONY: kind-reset
kind-reset: ## Destroys the "k8s-infra" kind cluster.
	kind delete cluster --name=$(KIND_CLUSTER_NAME) || true
	rm .$(KIND_CLUSTER_NAME).cluster

install: manifests ## Install CRDs into a cluster
	kustomize build config/crd | kubectl apply -f -

uninstall: manifests ## Uninstall CRDs from a cluster
	kustomize build config/crd | kubectl delete -f -

deploy: manifests ## Deploy controller in the configured Kubernetes cluster in ~/.kube/config
	cd config/manager && kustomize edit set image controller=${IMG}
	kustomize build config/default | kubectl apply -f -

manifests: $(CONTROLLER_GEN) ## Generate manifests e.g. CRD, RBAC etc.
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

fmt: ## Run go fmt against code
	go fmt ./...

vet: ## Run go vet against code
	go vet ./...

generate: $(CONTROLLER_GEN) ## Generate code
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths="./..."

docker-build: test ## Build the docker image
	docker build . -t ${IMG}

docker-push: ## Push the docker image
	docker push ${IMG}
