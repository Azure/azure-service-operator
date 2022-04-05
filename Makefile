PUBLIC_REPO=mcr.microsoft.com/k8s/azureserviceoperator
PLACEHOLDER_IMAGE=controller:latest

# Image URL to use all building/pushing image targets
IMG ?= $(PLACEHOLDER_IMAGE)

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

CRD_OPTIONS ?= "crd:crdVersions=v1"

BUILD_ID ?= $(shell git rev-parse --short HEAD)
timestamp := $(shell /bin/date "+%Y%m%d-%H%M%S")

# best to keep the prefix as short as possible to not exceed naming limits for things like keyvault (24 chars)
TEST_RESOURCE_PREFIX ?= aso-$(BUILD_ID)
KIND_CLUSTER_NAME = aso-kind

# Go compiler builds tags: some parts of the test suite use these to selectively compile tests.
BUILD_TAGS ?= all

ifdef TMPDIR
TMPDIR := $(realpath ${TMPDIR})
else
TMPDIR := /tmp
endif

.PHONY: all
all: manager

# Generate test certs for development
.PHONY: generate-test-certs
generate-test-certs: CONFIGTXT := $(shell mktemp)
generate-test-certs: WEBHOOK_DIR := $(TMPDIR)/k8s-webhook-server
generate-test-certs: WEBHOOK_CERT_DIR := $(TMPDIR)/k8s-webhook-server/serving-certs
generate-test-certs:
	rm -rf $(WEBHOOK_DIR)
	mkdir -p $(WEBHOOK_CERT_DIR)

	@echo "[req]" > $(CONFIGTXT)
	@echo "distinguished_name = req_distinguished_name" >> $(CONFIGTXT)
	@echo "[req_distinguished_name]" >> $(CONFIGTXT)
	@echo "[SAN]" >> $(CONFIGTXT)
	@echo "subjectAltName=DNS:azureoperator-webhook-service.azureoperator-system.svc.cluster.local" >> $(CONFIGTXT)

	@echo "OpenSSL Config:"
	@cat $(CONFIGTXT)
	@echo

	openssl req -x509 -days 730 -out $(WEBHOOK_CERT_DIR)/tls.crt -keyout $(WEBHOOK_CERT_DIR)/tls.key -newkey rsa:4096 -subj "/CN=azureoperator-webhook-service.azureoperator-system" -config $(CONFIGTXT) -nodes

# Run Controller tests against the configured cluster
.PHONY: test-integration-controllers
test-integration-controllers: generate fmt vet manifests
	TEST_RESOURCE_PREFIX=$(TEST_RESOURCE_PREFIX) TEST_USE_EXISTING_CLUSTER=false REQUEUE_AFTER=20 \
	AZURE_TARGET_NAMESPACES=default,watched \
	go test -v -tags "$(BUILD_TAGS)" -coverprofile=reports/integration-controllers-coverage-output.txt -coverpkg=./... -covermode count -parallel 4 -timeout 45m \
		./controllers/... \
		./pkg/secrets/...
		# TODO: Note that the above test (secrets/keyvault) is not an integration-controller test... but it's not a unit test either and unfortunately the test-integration-managers target isn't run in CI either?

# Check that when there are no target namespaces all namespaces are watched
.PHONY: test-no-target-namespaces
test-no-target-namespaces: generate fmt vet manifests
	TEST_RESOURCE_PREFIX=$(TEST_RESOURCE_PREFIX) TEST_USE_EXISTING_CLUSTER=false REQUEUE_AFTER=20 \
	AZURE_TARGET_NAMESPACES= \
	go test -v -tags "$(BUILD_TAGS)" -coverprofile=reports/no-target-namespaces-coverage-output.txt -coverpkg=./... -covermode count -parallel 4 -timeout 45m \
		-run TestTargetNamespaces \
		./controllers/...

# Check that we do the right thing in webhooks-only mode.
.PHONY: test-webhooks-only-mode
test-webhooks-only-mode: generate fmt vet manifests
	TEST_RESOURCE_PREFIX=$(TEST_RESOURCE_PREFIX) TEST_USE_EXISTING_CLUSTER=false REQUEUE_AFTER=20 \
	AZURE_OPERATOR_MODE=webhooks \
	go test -v -tags "$(BUILD_TAGS)" -coverprofile=reports/webhooks-only-coverage-output.txt -coverpkg=./... -covermode count -parallel 4 -timeout 45m \
		-run TestOperatorMode \
		./controllers

# Check that when there are no target namespaces all namespaces are watched
.PHONY: test-watchers-only-mode
test-watchers-only-mode: generate fmt vet manifests
	TEST_RESOURCE_PREFIX=$(TEST_RESOURCE_PREFIX) TEST_USE_EXISTING_CLUSTER=false REQUEUE_AFTER=20 \
	AZURE_OPERATOR_MODE=watchers \
	go test -v -tags "$(BUILD_TAGS)" -coverprofile=reports/watchers-only-coverage-output.txt -coverpkg=./... -covermode count -parallel 4 -timeout 45m \
		-run TestOperatorMode \
		./controllers

# Run subset of tests with v1 secret naming enabled to ensure no regression in old secret naming
.PHONY: test-v1-secret-naming
test-v1-secret-naming: generate fmt vet manifests
	TEST_RESOURCE_PREFIX=$(TEST_RESOURCE_PREFIX) TEST_USE_EXISTING_CLUSTER=false REQUEUE_AFTER=20 AZURE_SECRET_NAMING_VERSION=1 \
	go test -v -run "^.*_SecretNamedCorrectly$$" -tags "$(BUILD_TAGS)" -coverprofile=reports/v1-secret-naming-coverage-output.txt -coverpkg=./... -covermode count -parallel 4 -timeout 15m \
		./controllers/...

# Run Resource Manager tests against the configured cluster
.PHONY: test-integration-managers
test-integration-managers: generate fmt vet manifests
	TEST_USE_EXISTING_CLUSTER=true TEST_CONTROLLER_WITH_MOCKS=false REQUEUE_AFTER=20 \
	go test -v -coverprofile=reports/integration-managers-coverage-ouput.txt -coverpkg=./... -covermode count -parallel 4 -timeout 45m \
  ./api/... \
	./pkg/resourcemanager/eventhubs/...  \
	./pkg/resourcemanager/resourcegroups/...  \
	./pkg/resourcemanager/storages/... \
	./pkg/resourcemanager/psql/server/... \
	./pkg/resourcemanager/psql/database/... \
	./pkg/resourcemanager/psql/firewallrule/... \
	./pkg/resourcemanager/appinsights/... \
	./pkg/resourcemanager/vnet/...

# Run all available unit tests.
.PHONY: test-unit
test-unit: generate fmt vet manifests 
	TEST_USE_EXISTING_CLUSTER=false REQUEUE_AFTER=20 \
	go test -v -tags "$(BUILD_TAGS)" -coverprofile=reports/unittest-coverage-ouput.txt -covermode count -parallel 4 -timeout 10m \
	./pkg/resourcemanager/keyvaults/unittest/ \
	./pkg/resourcemanager/azuresql/azuresqlfailovergroup \
	./pkg/resourcemanager/cosmosdb/sqldatabase
	# The below folders are commented out because the tests in them fail...
	# ./api/... \

# Merge all the available test coverage results and publish a single report
.PHONY: test-process-coverage
test-process-coverage:
	find reports -name "*-coverage-output.txt" -type f -print | xargs gocovmerge > reports/merged-coverage-output.txt
	gocov convert reports/merged-coverage-output.txt > reports/merged-coverage-output.json
	gocov-xml < reports/merged-coverage-output.json > reports/merged-coverage.xml
	go tool cover -html=reports/merged-coverage-output.txt -o reports/merged-coverage.html

# Cleanup resource groups azure created by tests using pattern matching 't-rg-'
.PHONY: test-cleanup-azure-resources
test-cleanup-azure-resources: 	
	# Delete the resource groups that match the pattern
	for rgname in `az group list --query "[*].[name]" -o table | grep '^${TEST_RESOURCE_PREFIX}' `; do \
		echo "$$rgname will be deleted"; \
		az group delete --name $$rgname --no-wait --yes; \
	done

	for rgname in `az group list --query "[*].[name]" -o table | grep 'rg-prime$$' `; do \
		echo "$$rgname will be deleted"; \
		az group delete --name $$rgname --no-wait --yes; \
	done

# Build the docker image
.PHONY: docker-build
docker-build:
	docker build . -t ${IMG} ${ARGS}
	@echo "updating kustomize image patch file for manager resource"
	sed -i'' -e 's@image: .*@image: '"${IMG}"'@' ./config/default/manager_image_patch.yaml

# Push the docker image
.PHONY: docker-push
docker-push:
	docker push ${IMG}

# Build and Push the docker image
.PHONY: docker-build-and-push
docker-build-and-push: docker-build docker-push

# Build manager binary
.PHONY: manager
manager: generate fmt vet
	go build -o bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
.PHONY: run
run: generate fmt vet
	go run ./main.go

# Install CRDs into a cluster
.PHONY: generate
install: generate
	kubectl apply -f config/crd/bases

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
.PHONY: deploy
deploy: manifests
	kustomize build config/default | kubectl apply -f -

.PHONY: update
update:
	IMG="docker.io/controllertest:$(timestamp)" make ARGS="${ARGS}" docker-build
	kind load docker-image docker.io/controllertest:$(timestamp) --loglevel "trace"
	make install
	make deploy
	sed -i'' -e 's@image: .*@image: '"IMAGE_URL"'@' ./config/default/manager_image_patch.yaml

.PHONY: delete
delete:
	kubectl delete -f config/crd/bases
	kustomize build config/default | kubectl delete -f -

# Validate copyright headers
.PHONY: validate-copyright-headers
validate-copyright-headers:
	@./scripts/validate-copyright-headers.sh

# Validate cainjection files:
.PHONY: validate-cainjection-files
validate-cainjection-files:
	@./scripts/validate-cainjection-files.sh

# Generate manifests for helm and package them up
.PHONY: helm-chart-manifests
helm-chart-manifests: LATEST_TAG := $(shell curl -sL https://api.github.com/repos/Azure/azure-service-operator/releases/latest  | jq '.tag_name' --raw-output )
helm-chart-manifests: KUBE_RBAC_PROXY := gcr.io/kubebuilder/kube-rbac-proxy:v0.5.0
helm-chart-manifests: generate
	@echo "Latest released tag is $(LATEST_TAG)"
	# substitute released tag into values file.
	perl -pi -e 's,repository: $(PUBLIC_REPO):\K.*,$(LATEST_TAG),' ./charts/azure-service-operator/values.yaml
	# remove generated files
	rm -rf charts/azure-service-operator/templates/generated/
	rm -rf charts/azure-service-operator/crds
	# create directory for generated files
	mkdir charts/azure-service-operator/templates/generated
	mkdir charts/azure-service-operator/crds
	# generate files using kustomize
	kustomize build ./config/default -o ./charts/azure-service-operator/templates/generated
	# move CRD definitions to crd folder
	find ./charts/azure-service-operator/templates/generated/*_customresourcedefinition_* -exec mv '{}' ./charts/azure-service-operator/crds \;
	# remove namespace as we will let Helm manage it
	rm charts/azure-service-operator/templates/generated/*_namespace_*
	# replace hard coded ASO and kube-rbac images with Helm templating
	perl -pi -e s,controller:latest,"{{ .Values.image.repository }}",g ./charts/azure-service-operator/templates/generated/*_deployment_*
	# Ensure that what we're about to try to replace actually exists (if it doesn't we want to fail)
	grep -E $(KUBE_RBAC_PROXY) ./charts/azure-service-operator/templates/generated/*_deployment_*
	perl -pi -e s,$(KUBE_RBAC_PROXY),"{{ .Values.image.kubeRBACProxy }}",g ./charts/azure-service-operator/templates/generated/*_deployment_*
	# replace hard coded cert-manager version with templating
	sed -i "s@apiVersion: cert-manager.io/.*@apiVersion: {{ .Values.certManagerResourcesAPIVersion }}@g" ./charts/azure-service-operator/templates/generated/*cert-manager.io*
	# replace hard coded namespace with Helm templating
	find ./charts/azure-service-operator/templates/generated/ -type f -exec perl -pi -e s,azureoperator-system,"{{ .Release.Namespace }}",g {} \;
	# create unique names so each instance of the operator has its own role binding 
	find ./charts/azure-service-operator/templates/generated/ -name *clusterrole* -exec perl -pi -e 's/$$/-{{ .Release.Namespace }}/ if /name: azure/' {} \;
	# package the necessary files into a tar file
	helm package ./charts/azure-service-operator -d ./charts
	# update Chart.yaml for Helm Repository
	helm repo index ./charts

# Generate manifests e.g. CRD, RBAC etc.
.PHONY: manifests
manifests: install-tools
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases
	# update manifests to force preserveUnknownFields to false. We can't use controller-gen to set this to false because it has a bug...
	# see: https://github.com/kubernetes-sigs/controller-tools/issues/476
	# TODO: After this has been in the release for "a while" we can remove it since the default is also false and we just
	# TODO: need it for the upgrade scenario between v1beta1 and v1 CRD types
	ls config/crd/bases | xargs -I % yq eval -i ".spec.preserveUnknownFields = false" config/crd/bases/%

# Run go fmt against code
.PHONY: fmt
fmt:
	go fmt ./...

# Run go vet against code
.PHONY: vet
vet: 
	go vet ./...

# Generate code
.PHONY: generate
generate: manifests
	$(CONTROLLER_GEN) object:headerFile=./v2/boilerplate.go.txt paths=./api/...

.PHONY: install-bindata
install-bindata:
	go get -u github.com/jteeuwen/go-bindata/...

.PHONY: generate-template
generate-template:
	go-bindata -pkg template -prefix pkg/template/assets/ -o pkg/template/templates.go pkg/template/assets/

# TODO: These kind-delete / kind-create targets were stolen from k8s-infra and
# TODO: should be merged back together when the projects more closely align
.PHONY: kind-delete
kind-delete: install-test-tools
	kind delete cluster --name=$(KIND_CLUSTER_NAME) || true

.PHONY: kind-create
kind-create: install-test-tools
	kind get clusters | grep -E $(KIND_CLUSTER_NAME) > /dev/null;\
	EXISTS=$$?;\
	if [ $$EXISTS -eq 0 ]; then \
		echo "$(KIND_CLUSTER_NAME) already exists"; \
	else \
		kind create cluster --name=$(KIND_CLUSTER_NAME); \
	fi; \

.PHONY: set-kindcluster
set-kindcluster: kind-create
ifeq (${shell kind get kubeconfig-path --name=$(KIND_CLUSTER_NAME)},${KUBECONFIG})
	@echo "kubeconfig-path points to kind path"
else
	@echo "please run below command in your shell and then re-run make set-kindcluster"
	@echo  "\e[31mexport KUBECONFIG=$(shell kind get kubeconfig-path --name="$(KIND_CLUSTER_NAME)")\e[0m"
	@exit 111
endif
	@echo "getting value of KUBECONFIG"
	@echo ${KUBECONFIG}
	@echo "getting value of kind kubeconfig-path"

	kubectl cluster-info
	kubectl create namespace azureoperator-system
	kubectl --namespace azureoperator-system \
		create secret generic azureoperatorsettings \
		--from-literal=AZURE_CLIENT_ID=${AZURE_CLIENT_ID} \
		--from-literal=AZURE_CLIENT_SECRET=${AZURE_CLIENT_SECRET} \
		--from-literal=AZURE_SUBSCRIPTION_ID=${AZURE_SUBSCRIPTION_ID} \
		--from-literal=AZURE_TENANT_ID=${AZURE_TENANT_ID}

	make install-cert-manager

	#create image and load it into cluster
	make install
	IMG="docker.io/controllertest:1" make docker-build
	kind load docker-image docker.io/controllertest:1 --loglevel "trace" --name=$(KIND_CLUSTER_NAME)

	kubectl get namespaces
	kubectl get pods --namespace cert-manager
	@echo "Waiting for cert-manager to be ready"
	kubectl wait pod -n cert-manager --for condition=ready --timeout=60s --all
	@echo "all the pods should be running"
	make deploy
	sed -i'' -e 's@image: .*@image: '"IMAGE_URL"'@' ./config/default/manager_image_patch.yaml

.PHONY: install-kubebuilder
install-kubebuilder: OS := $(shell go env GOOS)
install-kubebuilder: ARCH := $(shell go env GOARCH)
install-kubebuilder: KUBEBUILDER_VERSION := 2.3.1
install-kubebuilder: KUBEBUILDER_DEST := $(shell go env GOPATH)/kubebuilder
install-kubebuilder:
ifeq (,$(shell which kubebuilder))
	@echo "installing kubebuilder"

	@echo "Installing kubebuilder ${KUBEBUILDER_VERSION} (${OS} ${ARCH})..."
	curl -L "https://github.com/kubernetes-sigs/kubebuilder/releases/download/v${KUBEBUILDER_VERSION}/kubebuilder_${KUBEBUILDER_VERSION}_${OS}_${ARCH}.tar.gz" | tar -xz -C $(TMPDIR)/
	mv "$(TMPDIR)/kubebuilder_${KUBEBUILDER_VERSION}_${OS}_${ARCH}" "${KUBEBUILDER_DEST}"
	export PATH=$$PATH:${KUBEBUILDER_DEST}/bin
else
	@echo "kubebuilder has been installed"
endif

.PHONY: install-cert-manager
install-cert-manager:
	kubectl create namespace cert-manager
	kubectl label namespace cert-manager cert-manager.io/disable-validation=true
	kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.7.2/cert-manager.yaml

.PHONY: install-aad-pod-identity
install-aad-pod-identity:
	kubectl apply -f https://raw.githubusercontent.com/Azure/aad-pod-identity/master/deploy/infra/deployment-rbac.yaml

.PHONY: install-test-tools
install-test-tools: TEST_TOOLS_MOD_DIR := $(shell mktemp -d -t goinstall_XXXXXXXXXX)
install-test-tools: install-tools
	cd $(TEST_TOOLS_MOD_DIR) \
	&& go mod init fake/mod \
	&& go get github.com/jstemmer/go-junit-report \
	&& go get github.com/axw/gocov/gocov \
	&& go get github.com/AlekSi/gocov-xml \
	&& go get github.com/wadey/gocovmerge \
	&& go get sigs.k8s.io/kind@v0.11.1
	rm -r $(TEST_TOOLS_MOD_DIR)

.PHONY: install-tools
install-tools: TEMP_DIR := $(shell mktemp -d -t goinstall_XXXXXXXXXX)
install-tools:
	go install github.com/mikefarah/yq/v4@v4.23.1
	go install k8s.io/code-generator/cmd/conversion-gen@v0.23.5
	go install sigs.k8s.io/kustomize/kustomize/v4@v4.5.4 
	go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.8.0
    CONTROLLER_GEN=$(shell go env GOPATH)/bin/controller-gen

# Operator-sdk release version
RELEASE_VERSION ?= v1.11.0

.PHONY: install-operator-sdk
install-operator-sdk:
ifeq ($(shell uname -s), Darwin)
	curl -LO https://github.com/operator-framework/operator-sdk/releases/download/${RELEASE_VERSION}/operator-sdk_darwin_amd64
	chmod +x operator-sdk_darwin_amd64 && sudo mkdir -p /usr/local/bin/ && sudo cp operator-sdk_darwin_amd64 /usr/local/bin/operator-sdk && rm operator-sdk_darwin_amd64
else
	curl -LO https://github.com/operator-framework/operator-sdk/releases/download/${RELEASE_VERSION}/operator-sdk_linux_amd64
	chmod +x operator-sdk_linux_amd64 && sudo mkdir -p /usr/local/bin/ && sudo cp operator-sdk_linux_amd64 /usr/local/bin/operator-sdk && rm operator-sdk_linux_amd64
endif

PREVIOUS_BUNDLE_VERSION ?= 1.0.28631

.PHONY: generate-operator-bundle
generate-operator-bundle: LATEST_TAG := $(shell curl -sL https://api.github.com/repos/Azure/azure-service-operator/releases/latest  | jq '.tag_name' --raw-output )
generate-operator-bundle: manifests
	@echo "Latest released tag is $(LATEST_TAG)"
	@echo "Previous bundle version is $(PREVIOUS_BUNDLE_VERSION)"
	rm -rf "bundle/manifests"
	kustomize build config/operator-bundle | operator-sdk generate bundle --version $(LATEST_TAG) --channels stable --default-channel stable --overwrite --kustomize-dir config/operator-bundle
	# Building the docker bundle requires a tests/scorecard directory.
	mkdir -p bundle/tests/scorecard
	# Remove the webhook service - OLM will create one when installing
	# the bundle.
	rm bundle/manifests/azureoperator-webhook-service_v1_service.yaml
	# Inject the container reference into the bundle.
	scripts/inject-container-reference.sh "$(PUBLIC_REPO):$(LATEST_TAG)"
	# Update webhooks to use operator namespace and remove
	# cert-manager annotations.
	scripts/update-webhook-references-in-operator-bundle.sh
	# Include the replaces field with the old version.
	yq eval -i ".spec.replaces = \"azure-service-operator.v$(PREVIOUS_BUNDLE_VERSION)\"" bundle/manifests/azure-service-operator.clusterserviceversion.yaml
	# Rename the csv to simplify adding to the community-operators repo for a PR
	mv bundle/manifests/azure-service-operator.clusterserviceversion.yaml bundle/manifests/azure-service-operator.v$(LATEST_TAG).clusterserviceversion.yaml
