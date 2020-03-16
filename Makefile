# Image URL to use all building/pushing image targets


# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

IMG ?= controller:latest
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"

BUILD_ID ?= $(shell git rev-parse --short HEAD)

# best to keep the prefix as short as possible to not exceed naming limits for things like keyvault (24 chars)
TEST_RESOURCE_PREFIX ?= aso-$(BUILD_ID)

all: manager

# Generate test certs for development
generate-test-certs:
	echo "[req]" > config.txt
	echo "distinguished_name = req_distinguished_name" >> config.txt
	echo "[req_distinguished_name]" >> config.txt
	echo "[SAN]" >> config.txt
	echo "subjectAltName=DNS:azureoperator-webhook-service.azureoperator-system.svc.cluster.local" >> config.txt
	openssl req -x509 -days 730 -out tls.crt -keyout tls.key -newkey rsa:4096 -subj "/CN=azureoperator-webhook-service.azureoperator-system" -config config.txt -nodes
	rm -rf /tmp/k8s-webhook-server
	mkdir -p /tmp/k8s-webhook-server/serving-certs
	mv tls.* /tmp/k8s-webhook-server/serving-certs/

# Run API unittests
api-test: generate fmt vet manifests
	TEST_USE_EXISTING_CLUSTER=false go test -v -coverprofile=coverage.txt -covermode count ./api/...  2>&1 | tee testlogs.txt
	go-junit-report < testlogs.txt  > report.xml
	go tool cover -html=coverage.txt -o cover.html

# Run tests
test: generate fmt vet manifests 
	TEST_USE_EXISTING_CLUSTER=false TEST_CONTROLLER_WITH_MOCKS=true REQUEUE_AFTER=20 \
	go test -tags all -parallel 3 -v -coverprofile=coverage.txt -covermode count \
	./api/... \
	./controllers/... \
	-timeout 10m 2>&1 | tee testlogs.txt
	go-junit-report < testlogs.txt > report.xml
	go tool cover -html=coverage.txt -o cover.html

# Run tests with existing cluster
test-existing-controllers: generate fmt vet manifests
	TEST_RESOURCE_PREFIX=$(TEST_RESOURCE_PREFIX) TEST_USE_EXISTING_CLUSTER=true REQUEUE_AFTER=20 go test -tags all -parallel 4 -v ./controllers/... -timeout 45m

unit-tests:
	go test ./pkg/resourcemanager/keyvaults/unittest/


# Run tests with existing cluster
test-existing-managers: generate fmt vet manifests
	TEST_USE_EXISTING_CLUSTER=true TEST_CONTROLLER_WITH_MOCKS=false REQUEUE_AFTER=20 \
	go test -v -coverprofile=coverage-existing.txt -covermode count \
	./api/... \
	./pkg/resourcemanager/eventhubs/...  \
	./pkg/resourcemanager/resourcegroups/...  \
	./pkg/resourcemanager/storages/... \
	./pkg/resourcemanager/psql/server/... \
	./pkg/resourcemanager/psql/database/... \
	./pkg/resourcemanager/psql/firewallrule/... \
	./pkg/resourcemanager/appinsights/... \
	./pkg/resourcemanager/vnet/... \
	./pkg/resourcemanager/apim/apimgmt...

# Cleanup resource groups azure created by tests using pattern matching 't-rg-'
test-cleanup-azure-resources: 	
	# Delete the resource groups that match the pattern
	for rgname in `az group list --query "[*].[name]" -o table | grep '^${TEST_RESOURCE_PREFIX}' `; do \
	    echo "$$rgname will be deleted"; \
	    az group delete --name $$rgname --no-wait --yes; \
    done

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./main.go

# Install CRDs into a cluster
install: generate
	kubectl apply -f config/crd/bases

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	kubectl apply -f config/crd/bases
	kustomize build config/default | kubectl apply -f -

timestamp := $(shell /bin/date "+%Y%m%d-%H%M%S")

update:
	IMG="docker.io/controllertest:$(timestamp)" make ARGS="${ARGS}" docker-build
	kind load docker-image docker.io/controllertest:$(timestamp) --loglevel "trace"
	make install
	make deploy
	sed -i'' -e 's@image: .*@image: '"IMAGE_URL"'@' ./config/default/manager_image_patch.yaml

delete:
	kubectl delete -f config/crd/bases
	kustomize build config/default | kubectl delete -f -

# Generate manifests for helm and package them up
helm-chart-manifests: manifests
	kustomize build ./config/default -o ./charts/azure-service-operator/templates
	rm charts/azure-service-operator/templates/~g_v1_namespace_azureoperator-system.yaml
	sed -i '' -e 's@controller:latest@{{ .Values.image.repository }}@' ./charts/azure-service-operator/templates/apps_v1_deployment_azureoperator-controller-manager.yaml
	helm package ./charts/azure-service-operator -d ./charts
	helm repo index ./charts

# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate: manifests
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths=./api/...

# Build the docker image
docker-build:
	docker build . -t ${IMG} ${ARGS}
	@echo "updating kustomize image patch file for manager resource"
	sed -i'' -e 's@image: .*@image: '"${IMG}"'@' ./config/default/manager_image_patch.yaml

# Push the docker image
docker-push:
	docker push ${IMG}

# Build and Push the docker image
build-and-push: docker-build docker-push

# Deploy operator infrastructure
terraform:
	terraform init devops/terraform
	terraform apply devops/terraform

terraform-and-deploy: terraform generate install-cert-manager build-and-push deploy

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.0
CONTROLLER_GEN=$(shell go env GOPATH)/bin/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

.PHONY: install-bindata
install-bindata:
	go get -u github.com/jteeuwen/go-bindata/...

.PHONE:
generate-template:
	go-bindata -pkg template -prefix pkg/template/assets/ -o pkg/template/templates.go pkg/template/assets/

create-kindcluster:
ifeq (,$(shell kind get clusters))
	@echo "no kind cluster"
else
	@echo "kind cluster is running, deleteing the current cluster"
	kind delete cluster
endif
	@echo "creating kind cluster"
	kind create cluster

set-kindcluster: install-kind
ifeq (${shell kind get kubeconfig-path --name="kind"},${KUBECONFIG})
	@echo "kubeconfig-path points to kind path"
else
	@echo "please run below command in your shell and then re-run make set-kindcluster"
	@echo  "\e[31mexport KUBECONFIG=$(shell kind get kubeconfig-path --name="kind")\e[0m"
	@exit 111
endif
	make create-kindcluster

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

install-kind:
ifeq (,$(shell which kind))
	@echo "installing kind"
	GO111MODULE="on" go get sigs.k8s.io/kind@v0.4.0
else
	@echo "kind has been installed"
endif

install-kubebuilder:
ifeq (,$(shell which kubebuilder))
	@echo "installing kubebuilder"
	# download kubebuilder and extract it to tmp
	curl -sL https://go.kubebuilder.io/dl/2.0.0/$(shell go env GOOS)/$(shell go env GOARCH) | tar -xz -C /tmp/
	# move to a long-term location and put it on your path
	# (you'll need to set the KUBEBUILDER_ASSETS env var if you put it somewhere else)
	mv /tmp/kubebuilder_2.0.0_$(shell go env GOOS)_$(shell go env GOARCH) /usr/local/kubebuilder
	export PATH=$$PATH:/usr/local/kubebuilder/bin
else
	@echo "kubebuilder has been installed"
endif

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

install-cert-manager:
	kubectl create namespace cert-manager
	kubectl label namespace cert-manager cert-manager.io/disable-validation=true
	kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v0.12.0/cert-manager.yaml

install-aad-pod-identity:
	kubectl apply -f https://raw.githubusercontent.com/Azure/aad-pod-identity/master/deploy/infra/deployment-rbac.yaml

install-test-dependency:
	go get -u github.com/jstemmer/go-junit-report \
	&& go get github.com/axw/gocov/gocov \
	&& go get github.com/AlekSi/gocov-xml \
	&& go get github.com/onsi/ginkgo/ginkgo \
	&& go get golang.org/x/tools/cmd/cover
