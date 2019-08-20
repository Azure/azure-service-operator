
# Image URL to use all building/pushing image targets
IMG ?= controller:latest
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"

all: manager

# Run tests
test: generate fmt vet manifests
	TEST_USE_EXISTING_CLUSTER=false go test -v -coverprofile=coverage.txt -covermode count ./api/... ./controllers/... ./resourcemanager/eventhubs/...  ./resourcemanager/resourcegroups/... 2>&1 | tee testlogs.txt
	go-junit-report < testlogs.txt  > report.xml
	go tool cover -html=coverage.txt -o cover.html
# Run tests with existing cluster
test-existing: generate fmt vet manifests
	TEST_USE_EXISTING_CLUSTER=true go test -v -coverprofile=coverage-existing.txt -covermode count ./api/... ./controllers/... ./resourcemanager/eventhubs/...  ./resourcemanager/resourcegroups/... 2>&1 | tee testlogs-existing.txt
	go-junit-report < testlogs-existing.txt  > report-existing.xml
	go tool cover -html=coverage-existing.txt -o cover-existing.html

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./main.go

# Install CRDs into a cluster
install: manifests
	kubectl apply -f config/crd/bases

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	kubectl apply -f config/crd/bases
	kustomize build config/default | kubectl apply -f -

delete:
	kubectl delete -f config/crd/bases
	kustomize build config/default | kubectl delete -f -

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
generate: controller-gen
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths=./api/...

# Build the docker image
docker-build: 
	docker build . -t ${IMG}
	@echo "updating kustomize image patch file for manager resource"
	sed -i'' -e 's@image: .*@image: '"${IMG}"'@' ./config/default/manager_image_patch.yaml

# Push the docker image
docker-push:
	docker push ${IMG}

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.0-beta.3
CONTROLLER_GEN=$(shell go env GOPATH)/bin/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

install-kind:
ifeq (,$(shell which kind))
	@echo "installing kind"
	GO111MODULE="on" go get sigs.k8s.io/kind@v0.4.0
else
	@echo "kind has been installed"
endif

create-kindcluster:
ifeq (,$(shell kind get clusters))
	@echo "no kind cluster"
else
	@echo "kind cluster is running, deleteing the current cluster"
	kind delete cluster 
endif
	@echo "creating kind cluster"
	kind create cluster

install-cert-manager:
	kubectl create namespace cert-manager
	kubectl label namespace cert-manager certmanager.k8s.io/disable-validation=true
	kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v0.9.0/cert-manager.yaml

set-kindcluster: install-kind
	#KUBECONFIG=$(shell kind get kubeconfig-path --name="kind")
	#$(shell export KUBECONFIG="$(kind get kubeconfig-path --name="kind")")
ifeq ($(shell kind get kubeconfig-path --name="kind"),$(KUBECONFIG))
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
	kind get kubeconfig-path --name="kind"
	
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
	IMG="docker.io/controllertest:1" make docker-build
	kind load docker-image docker.io/controllertest:1 --loglevel "trace"
	make install
	kubectl get namespaces
	@echo "sleep 80 seconds to get the cert pods running"
	sleep 80
	@echo "end of sleep"
	kubectl get pods --namespace cert-manager
	@echo "all the pods should be running"
	make deploy

install-kustomize:
ifeq (,$(shell which kustomize))
	@echo "installing kustomize"
	# download kustomize
	sudo mkdir -p /usr/local/kustomize/
	sudo curl -o /usr/local/kubebuilder/bin/kustomize -sL "https://go.kubebuilder.io/kustomize/$(shell go env GOOS)/$(shell go env GOARCH)"
	# set permission
	sudo chmod a+x /usr/local/kubebuilder/bin/kustomize
	# export path
	$(shell which kustomize)
	
else
	@echo "kustomize has been installed"
endif
