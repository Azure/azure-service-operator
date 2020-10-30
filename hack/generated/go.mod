module github.com/Azure/k8s-infra/hack/generated

go 1.13

require (
	github.com/Azure/go-autorest/autorest v0.10.2
	github.com/Azure/go-autorest/autorest/azure/auth v0.4.2
	github.com/Azure/go-autorest/autorest/date v0.2.0
	github.com/Azure/k8s-infra v0.2.0
	github.com/devigned/tab v0.1.1
	github.com/go-logr/logr v0.1.0
	github.com/google/uuid v1.1.1
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1
	k8s.io/api v0.18.6
	k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	k8s.io/klog/v2 v2.0.0
	sigs.k8s.io/controller-runtime v0.6.2
)
