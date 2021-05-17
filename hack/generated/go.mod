module github.com/Azure/azure-service-operator/hack/generated

go 1.13

require (
	github.com/Azure/azure-service-operator v0.0.0 // version here doesn't matter
	github.com/Azure/go-autorest/autorest v0.11.0
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.0
	github.com/Azure/go-autorest/autorest/date v0.3.0
	github.com/devigned/tab v0.1.1
	github.com/dnaeon/go-vcr v1.1.0
	github.com/go-logr/logr v0.1.0
	github.com/google/go-cmp v0.5.2
	github.com/google/uuid v1.1.1
	github.com/kr/pretty v0.2.0
	github.com/kylelemons/godebug v1.1.0
	github.com/leanovate/gopter v0.2.8
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1
	k8s.io/api v0.18.6
	k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	k8s.io/klog/v2 v2.0.0
	sigs.k8s.io/controller-runtime v0.6.5
)

replace github.com/Azure/azure-service-operator => ../../
