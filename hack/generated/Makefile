SHELL           = /bin/bash
PACKAGE         = github.com/Azure/k8s-infra/hack/generated
APP             = k8sinfra-controller

include ../../tools.mk

CRD_OPTIONS     ?= "crd:crdVersions=v1,allowDangerousTypes=true"

V = 0
Q = $(if $(filter 1,$V),,@)

.PHONY: all
all: generate build

.PHONY: generate
generate: $(CONTROLLER_GEN) $(CONVERSION_GEN) ## Generate code
	# Force regeneration of all of the conversions
	@echo "Deleting old controller gen files"
	$(Q) find "./apis" -type f -name "zz_generated.*" -delete

	@echo "Executing controller-gen"
	$(Q) $(CONTROLLER_GEN) object:headerFile=../boilerplate.go.txt paths="./..."

#	@echo "Executing conversion-gen"
#	$(Q) $(CONVERSION_GEN) \
#		--input-dirs="./apis/..." \
#		--output-file-base=zz_generated.conversion \
#		--output-base=./ \
#		--go-header-file=../boilerplate.go.txt

	@echo "Executing controller-gen to generate CRDs"
	$(Q) $(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./apis/..." output:crd:artifacts:config=config/crd/bases

# TODO: Right now this is just proving that the packages can be built - we don't have any executable consuming them yet so no binary is produced
.PHONY: build
build:
	$(Q) $(GO) build ./apis/...
