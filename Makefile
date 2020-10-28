SHELL := /bin/bash
.DEFAULT_GOAL:=build

timestamp := $(shell /bin/date "+%Y%m%d-%H%M%S")
CONFIG_REGISTRY = kind-registry:5000/fake/k8s-infra-controller:latest
IMG ?= k8s-infra-contoller:$(timestamp)
CRD_OPTIONS ?= "crd:crdVersions=v1"

KIND_CLUSTER_NAME ?= k8sinfra
KIND_KUBECONFIG := $(HOME)/.kube/kind-$(KIND_CLUSTER_NAME)
TLS_CERT_PATH := pki/certs/tls.crt

# Include common variables and targets
include tools.mk

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Testing / Toooling
## --------------------------------------
.PHONY: test test-int test-cover test-cover-int
test test-int test-cover test-cover-int: export TEST_ASSET_KUBECTL = $(KUBECTL)
test test-int test-cover test-cover-int: export TEST_ASSET_KUBE_APISERVER = $(KUBE_APISERVER)
test test-int test-cover test-cover-int: export TEST_ASSET_ETCD = $(ETCD)

test: $(KUBECTL) $(KUBE_APISERVER) $(ETCD) lint header-check ## Run tests
	$(GO) test -v ./...

test-int: $(ROOT_DIR)/.env $(KUBECTL) $(KUBE_APISERVER) $(ETCD) header-check lint ## Run integration tests
	# MUST be executed as single command, or env vars will not propagate to test execution
	. .env && $(GO) test -v ./... -tags integration

.env: ## create a service principal and save the identity to .env for use in integration tests (requries jq and az)
	$(SCRIPTS_DIR)/create_testing_creds.sh

test-cover: $(KUBECTL) $(KUBE_APISERVER) $(ETCD) header-check lint ## Run tests w/ code coverage (./cover.out)
	$(GO) test ./... -coverprofile=cover.out -coverpkg=./...

test-cover-int: $(KUBECTL) $(KUBE_APISERVER) $(ETCD) header-check lint ## Run tests w/ code coverage (./cover.out)
	$(GO) test ./... -tags integration -coverprofile=cover.out -coverpkg=./...

## --------------------------------------
## Linting
## --------------------------------------

.PHONY: lint
lint: $(GOLANGCI_LINT) ## Lint codebase
	$(GOLANGCI_LINT) run -v

.PHONY: lint-full
lint-full: $(GOLANGCI_LINT) ## Run slower linters to detect possible issues
	$(GOLANGCI_LINT) run -v --fast=false

## --------------------------------------
## Build
## --------------------------------------

.PHONY: build
build: fmt vet lint ## Build manager binary
	$(GO) build -o bin/manager main.go

.PHONY: fmt
fmt: ## Run go fmt against code
	$(GO) fmt ./...

.PHONY: vet
vet: ## Run go vet against code
	$(GO) vet ./...

.PHONY: tidy
tidy: ## Runs go mod to ensure tidy.
	$(GO) mod tidy

## --------------------------------------
## Generate
## --------------------------------------

.PHONY: manifests
manifests: $(CONTROLLER_GEN) ## Generate manifests e.g. CRD, RBAC etc.
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

.PHONY: generate
generate: manifests $(CONTROLLER_GEN) $(CONVERSION_GEN) ## Generate code
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths="./..."

	$(CONVERSION_GEN) \
    		--input-dirs=./apis/microsoft.network/v20191101,./apis/microsoft.resources/v20191001,./apis/microsoft.compute/v20191201,./apis/microsoft.resources/v20150101 \
    		--output-file-base=zz_generated.conversion \
    		--output-base=$(ROOT_DIR) \
    		--go-header-file=./hack/boilerplate.go.txt

## --------------------------------------
## Development
## --------------------------------------

.PHONY: tilt-up
tilt-up: kind-create $(ROOT_DIR)/.env ## start tilt and build kind cluster if needed
	tilt up

.PHONY: kind-reset
kind-reset: $(KIND) ## Destroys the "k8sinfra" kind cluster.
	$(KIND) delete cluster --name=$(KIND_CLUSTER_NAME) || true

.PHONY: kind-create
kind-create: $(KIND) ## Destroys the "k8sinfra" kind cluster.
	./scripts/kind-with-registry.sh

.PHONY: run
run: $(KIND) kind-create
run: export KUBECONFIG = $(KIND_KUBECONFIG)
run: export ENVIRONMENT = development
run: $(TLS_CERT_PATH) generate fmt manifests install ## Run a development cluster using kind
	$(GO) run ./main.go

$(KIND_CLUSTER_TOUCH): $(KIND) $(KUBECTL)
	$(KIND) create cluster --name=$(KIND_CLUSTER_NAME) --kubeconfig=$(KIND_KUBECONFIG) --image=kindest/node:v1.17.4
	touch $(KIND_CLUSTER_TOUCH)

.PHONY: apply-certs-and-secrets
apply-certs-and-secrets: $(KUBECTL) ## Apply certificate manager and manager secrets into cluster
	./scripts/apply_cert_and_secrets.sh

.PHONY: deploy-kind
deploy-kind: $(KIND) kind-create
deploy-kind: export KUBECONFIG = $(KIND_KUBECONFIG)
deploy-kind: apply-certs-and-secrets deploy ## Deploy manager and secrets into kind cluster

.PHONY: install
install: manifests $(KUBECTL) $(KUSTOMIZE) ## Install CRDs into a cluster
	$(KUSTOMIZE) build config/crd | $(KUBECTL) apply -f -

.PHONY: uninstall
uninstall: manifests $(KUBECTL) $(KUSTOMIZE) ## Uninstall CRDs from a cluster
	$(KUSTOMIZE) build config/crd | $(KUBECTL) delete -f -

## --------------------------------------
## Deploy / Release
## --------------------------------------

.PHONY: deploy
deploy: manifests $(KUBECTL) $(KUSTOMIZE) docker-build docker-push ## Deploy controller in the configured Kubernetes cluster in ~/.kube/config
	$(KUSTOMIZE) build config/default | sed "s_${CONFIG_REGISTRY}_${REGISTRY}/${IMG}_" | $(KUBECTL) apply -f -

.PHONY: docker-build
docker-build: ## Build the docker image
	docker build . -t $(REGISTRY)/${IMG}

.PHONY: docker-push
docker-push: ## Push the docker image
	docker push $(REGISTRY)/${IMG}

.PHONY: dist
dist: $(KUSTOMIZE)
	mkdir -p dist
	$(KUSTOMIZE) build config/default | sed "s_${CONFIG_REGISTRY}_${REGISTRY}/${IMG}_" > dist/release.yaml

.PHONY: release
release: dist docker-build docker-push ## Build, push, generate dist for release
