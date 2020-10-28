SHELL             = /bin/bash
PACKAGE           = github.com/Azure/k8s-infra/hack/generated
APP               = k8sinfra-controller

CONFIG_REGISTRY = kind-registry:5000/fake/k8s-infra-controller:latest
IMG               ?= k8s-infra-generated-contoller:latest
KIND_CLUSTER_NAME = k8sinfra-generated

include ../../tools.mk

CRD_OPTIONS     ?= "crd:crdVersions=v1,allowDangerousTypes=true"
GO_DIRS         := $(shell $(GO) list -f '{{.Dir}}' ./...)
# We exclude the apis folder because it's really large and test discovery takes a good amount of time (>10s)
GO_DIRS_TO_TEST := $(shell $(GO) list -f '{{.Dir}}' ./... | grep -v /apis/)
CONTROLLER_DEBUG_LOGLEVEL := 4

V = 0
Q = $(if $(filter 1,$V),,@)

.PHONY: all
all: generate header-check fmt build test test-int  # TODO: Do we want to remove test-int from the all target so it's not run every time locally?

# There is a ci specific target because we want the CI pass to fail if
# the code has not been go fmt-ed, whereas locally we want "make all"
# to just format the code for you
.PHONY: ci
ci: generate build  # test-cover or test-cover-int will be called by the CI job directly

## --------------------------------------
## Build
## --------------------------------------

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

.PHONY: generate
generate: $(CONTROLLER_GEN) $(CONVERSION_GEN) ## Generate code
	# Force regeneration of all of the conversions
	@echo "Deleting old deepcopy files"
	$(Q) find "./apis" -type f -name "zz_generated.*" -delete

	@echo "Executing controller-gen to generate deepcopy functions"
	$(Q) $(CONTROLLER_GEN) object:headerFile=../boilerplate.go.txt paths="./..."

#	@echo "Executing conversion-gen"
#	$(Q) $(CONVERSION_GEN) \
#		--input-dirs="./apis/..." \
#		--output-file-base=zz_generated.conversion \
#		--output-base=./ \
#		--go-header-file=../boilerplate.go.txt

	# Force regeneration of all of the CRDs
	@echo "Deleting old CRD YAMLs"
	$(Q) if [ -d "./config/crd/bases" ]; then find "./config/crd/bases" -type f -name "*" -delete; fi

	@echo "Executing controller-gen to generate CRD and RBAC YAMLs"
	$(Q) $(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

.PHONY: build
build: tidy lint ; $(info $(M) building ./bin/$(APP))
	$(Q) CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on $(GO) build -o ./bin/$(APP)

## --------------------------------------
## Test
## --------------------------------------
.PHONY: test
test: ; $(info $(M) running go test…)
	$(Q) $(GO) test $(GO_DIRS_TO_TEST) -short -tags=noexit

.PHONY: test-cover
test-cover: $(GCOV2LCOV) ; $(info $(M) running go test…)
	# NOTE: if you update the 'test-cover' target, also update ./github/workflows/test.yml
	#       for the Windows part of the "test-generator" job.
	$(Q) $(GO) test -short -tags=noexit -race -covermode atomic -coverprofile=cover.out -coverpkg=./... $(GO_DIRS_TO_TEST)
	$(Q) $(GCOV2LCOV) -infile cover.out -outfile coverage.lcov

# Initially this target uses kind as there is no need to end to end test on Azure.
# Eventually when we start supporting features like MSI that only work in Azure,
# it's likely we'll need to make the integration tests (or at least some of them)
# run in an AKS cluster.
.PHONY: test-int-no-cleanup
test-int-no-cleanup: ; $(info $(M) running controller integration test…)
	# MUST be executed as single command, or env vars will not propagate to test execution
	$(Q) $(GO) test $(GO_DIRS_TO_TEST) # TODO: better way to just run integration tests?

.PHONY: test-int-no-cleanup-cover
test-int-no-cleanup-cover: $(GCOV2LCOV) ; $(info $(M) running controller integration test…)
	# MUST be executed as single command, or env vars will not propagate to test execution
	$(Q) $(GO) test -race -covermode atomic -coverprofile=cover.out -coverpkg=./... $(GO_DIRS_TO_TEST)
	$(Q) $(GCOV2LCOV) -infile cover.out -outfile coverage.lcov

.PHONY: test-int
test-int: kind-create deploy test-int-no-cleanup kind-delete cleanup-test-azure-resources

.PHONY: test-int-cover
test-int-cover: kind-create deploy test-int-no-cleanup-cover kind-delete cleanup-test-azure-resources

# Cleanup resource groups created by tests -- this isn't strictly required as the tests
# clean up after themselves, but doing it here anyway just to be doubly sure we don't leak
# resources in cases where the test pass is terminated, panics, etc
# This finds all resource groups which match the specified pattern (k8sinfratest) and are older than a day
# (86400 seconds). This is a bit horrible but it works...
.PHONY: cleanup-test-azure-resources
cleanup-test-azure-resources:
	$(Q) rgs=`az group list --query "[*].{Name: name, CreatedAt: tags.CreatedAt}" \
		| jq -r '.[] | select(.Name | test("^k8sinfratest")) | select(.CreatedAt == null or now-(.CreatedAt | fromdate) > 86400) | .Name'`; \
	for rgname in $${rgs[@]} ; do \
		echo "$$rgname will be deleted"; \
		az group delete --name $$rgname --no-wait --yes; \
	done

## --------------------------------------
## Development
## --------------------------------------

.PHONY: kind-delete
kind-delete: $(KIND) ## Destroys the "k8sinfra" kind cluster.
	$(KIND) delete cluster --name=$(KIND_CLUSTER_NAME) || true

.PHONY: kind-create
kind-create:
	export KIND_CLUSTER_NAME=$(KIND_CLUSTER_NAME) && $(SCRIPTS_DIR)/kind-with-registry.sh

# TODO: We may want this later
#.PHONY: apply-certs-and-secrets
#apply-certs-and-secrets: $(KUBECTL) ## Apply certificate manager and manager secrets into cluster
#	$(SCRIPTS_DIR)/apply_cert_and_secrets.sh

.PHONY: run
run: export ENVIRONMENT = development
run: kind-create install ## Run a development cluster using kind
	$(Q) $(GO) run ./main.go -v $(CONTROLLER_DEBUG_LOGLEVEL)

## --------------------------------------
## Deploy
## --------------------------------------

.PHONY: install
install: generate $(KUBECTL) $(KUSTOMIZE) ## Install CRDs into a cluster
	$(KUSTOMIZE) build config/crd | $(KUBECTL) apply -f -

.PHONY: uninstall
uninstall: generate $(KUBECTL) $(KUSTOMIZE) ## Uninstall CRDs from a cluster
	$(KUSTOMIZE) build config/crd | $(KUBECTL) delete -f -

.PHONY: deploy
deploy: generate $(KUBECTL) $(KUSTOMIZE) docker-build docker-push ## Deploy controller in the configured Kubernetes cluster in ~/.kube/config
	# TODO: Consider patching in CONTROLLER_DEBUG_LOGLEVEL?
	$(KUSTOMIZE) build config/default | sed "s_${CONFIG_REGISTRY}_${REGISTRY}/${IMG}_" | $(KUBECTL) apply -f -
	$(SCRIPTS_DIR)/deploy_testing_secret.sh

.PHONY: docker-build
docker-build: ## Build the docker image
	docker build . -t $(REGISTRY)/${IMG}

.PHONY: docker-push
docker-push: ## Push the docker image
	docker push $(REGISTRY)/${IMG}
