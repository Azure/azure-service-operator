# Image URL to use all building/pushing image targets
IMG ?= controller:latest

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd"

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
	TEST_RESOURCE_PREFIX=$(TEST_RESOURCE_PREFIX) TEST_USE_EXISTING_CLUSTER=true REQUEUE_AFTER=20 \
	go test -v -tags "$(BUILD_TAGS)" -coverprofile=reports/integration-controllers-coverage-output.txt -coverpkg=./... -covermode count -parallel 4 -timeout 45m \
	./controllers/... 
	#2>&1 | tee reports/integration-controllers-output.txt
	#go-junit-report < reports/integration-controllers-output.txt > reports/integration-controllers-report.xml

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
	#2>&1 | tee reports/integration-managers-output.txt
	#go-junit-report < reports/integration-managers-output.txt > reports/integration-managers-report.xml

# Run all available tests. Note that Controllers are not unit-testable.
.PHONY: test-unit
test-unit: generate fmt vet manifests 
	TEST_USE_EXISTING_CLUSTER=false REQUEUE_AFTER=20 \
	go test -v -tags "$(BUILD_TAGS)" -coverprofile=coverage-unit.txt -covermode count -parallel 4 -timeout 10m \
	./api/... \
	./pkg/secrets/...
	./pkg/resourcemanager/keyvaults/unittest/ \
	#2>&1 | tee testlogs.txt
	#go-junit-report < testlogs.txt > report-unit.xml
	go tool cover -html=coverage/coverage.txt -o cover-unit.html

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
.PHONY: build-and-push
build-and-push: docker-build docker-push

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
helm-chart-manifests: generate
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
	# replace hard coded ASO image with Helm templating
	perl -pi -e s,controller:latest,"{{ .Values.image.repository }}",g ./charts/azure-service-operator/templates/generated/*_deployment_*
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
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

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
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths=./api/...

# find or download controller-gen
# download controller-gen if necessary
.PHONY: controller-gen
controller-gen:
ifeq (, $(shell which controller-gen))
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.5
CONTROLLER_GEN=$(shell go env GOPATH)/bin/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

.PHONY: install-bindata
install-bindata:
	go get -u github.com/jteeuwen/go-bindata/...

.PHONY: generate-template
generate-template:
	go-bindata -pkg template -prefix pkg/template/assets/ -o pkg/template/templates.go pkg/template/assets/

# TODO: These kind-delete / kind-create targets were stolen from k8s-infra and
# TODO: should be merged back together when the projects more closely align
.PHONY: kind-delete
kind-delete: install-kind
	kind delete cluster --name=$(KIND_CLUSTER_NAME) || true

.PHONY: kind-create
kind-create: install-kind
	kind get clusters | grep -E $(KIND_CLUSTER_NAME) > /dev/null;\
	EXISTS=$$?;\
	if [ $$EXISTS -eq 0 ]; then \
		echo "$(KIND_CLUSTER_NAME) already exists"; \
	else \
		kind create cluster --name=$(KIND_CLUSTER_NAME); \
	fi; \

.PHONY: set-kindcluster
set-kindcluster: install-kind kind-create
ifeq (${shell kind get kubeconfig-path --name="kind"},${KUBECONFIG})
	@echo "kubeconfig-path points to kind path"
else
	@echo "please run below command in your shell and then re-run make set-kindcluster"
	@echo  "\e[31mexport KUBECONFIG=$(shell kind get kubeconfig-path --name="kind")\e[0m"
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
	kind load docker-image docker.io/controllertest:1 --loglevel "trace"

	kubectl get namespaces
	kubectl get pods --namespace cert-manager
	@echo "Waiting for cert-manager to be ready"
	kubectl wait pod -n cert-manager --for condition=ready --timeout=60s --all
	@echo "all the pods should be running"
	make deploy
	sed -i'' -e 's@image: .*@image: '"IMAGE_URL"'@' ./config/default/manager_image_patch.yaml

.PHONY: install-kind
install-kind:
ifeq (,$(shell which kind))
	@echo "installing kind"
	GO111MODULE="on" go get sigs.k8s.io/kind@v0.8.1
else
	@echo "kind has been installed"
endif

.PHONY: install-kubebuilder
install-kubebuilder:
ifeq (,$(shell which kubebuilder))
	@echo "installing kubebuilder"
	# download kubebuilder and extract it to tmp
	curl -sL https://go.kubebuilder.io/dl/2.0.0/$(shell go env GOOS)/$(shell go env GOARCH) | tar -xz -C $(TMPDIR)
	# move to a long-term location and put it on your path
	# (you'll need to set the KUBEBUILDER_ASSETS env var if you put it somewhere else)
	mv $(TMPDIR)/kubebuilder_2.0.0_$(shell go env GOOS)_$(shell go env GOARCH) /usr/local/kubebuilder
	export PATH=$$PATH:/usr/local/kubebuilder/bin
else
	@echo "kubebuilder has been installed"
endif

.PHONY: install-kustomize
install-kustomize:
ifeq (,$(shell which kustomize))
	@echo "installing kustomize"
	mkdir -p /usr/local/kubebuilder/bin
	# download kustomize
	curl -o /usr/local/kubebuilder/bin/kustomize -sL "https://go.kubebuilder.io/kustomize/$(shell go env GOOS)/$(shell go env GOARCH)"
	# set permission
	chmod a+x /usr/local/kubebuilder/bin/kustomize
	$(shell which kustomize)
else
	@echo "kustomize has been installed"
endif

.PHONY: install-cert-manager
install-cert-manager:
	kubectl create namespace cert-manager
	kubectl label namespace cert-manager cert-manager.io/disable-validation=true
	kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v0.12.0/cert-manager.yaml

.PHONY: install-aad-pod-identity
install-aad-pod-identity:
	kubectl apply -f https://raw.githubusercontent.com/Azure/aad-pod-identity/master/deploy/infra/deployment-rbac.yaml

.PHONY: install-test-dependencies
install-test-dependencies:
	go get github.com/jstemmer/go-junit-report \
	&& go get github.com/axw/gocov/gocov \
	&& go get github.com/AlekSi/gocov-xml \
	&& go get github.com/wadey/gocovmerge

# Operator-sdk release version
RELEASE_VERSION ?= v1.0.1

.PHONY: install-operator-sdk
install-operator-sdk:
ifeq ($(shell uname -s), Darwin)
	curl -LO https://github.com/operator-framework/operator-sdk/releases/download/${RELEASE_VERSION}/operator-sdk-${RELEASE_VERSION}-x86_64-apple-darwin
	chmod +x operator-sdk-${RELEASE_VERSION}-x86_64-apple-darwin && sudo mkdir -p /usr/local/bin/ && sudo cp operator-sdk-${RELEASE_VERSION}-x86_64-apple-darwin /usr/local/bin/operator-sdk && rm operator-sdk-${RELEASE_VERSION}-x86_64-apple-darwin
else
	curl -LO https://github.com/operator-framework/operator-sdk/releases/download/${RELEASE_VERSION}/operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu
	chmod +x operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu && sudo mkdir -p /usr/local/bin/ && sudo cp operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu /usr/local/bin/operator-sdk && rm operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu
endif

# Current operator version
VERSION ?= 0.37.0

generate-operator-bundle: manifests
	kustomize build config/manifests | operator-sdk generate bundle --version $(VERSION) --channels stable --default-channel stable --overwrite
	# This is only needed until CRD conversion support is released in OpenShift 4.6.x/Operator Lifecycle Manager 0.16.x
	scripts/add-openshift-cert-handling.sh
	# Rather than modify config/rbac manifests, replace CSV's default serviceAccount with azure-service-operator
	sed -i 's/serviceAccountName: default/serviceAccountName: azure-service-operator/g' bundle/manifests/azure-service-operator.clusterserviceversion.yaml
