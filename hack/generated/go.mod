module github.com/Azure/azure-service-operator/hack/generated

go 1.16

require (
	github.com/Azure/go-autorest/autorest v0.11.19
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.0
	github.com/Azure/go-autorest/autorest/date v0.3.0
	github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/benbjohnson/clock v1.1.0
	github.com/devigned/tab v0.1.1
	github.com/dnaeon/go-vcr v1.1.0
	github.com/go-logr/logr v0.4.0
	github.com/google/go-cmp v0.5.5
	github.com/google/uuid v1.1.2
	github.com/kr/pretty v0.2.0
	github.com/kylelemons/godebug v1.1.0
	github.com/leanovate/gopter v0.2.8
	github.com/onsi/gomega v1.14.0
	github.com/pkg/errors v0.9.1
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/time v0.0.0-20210611083556-38a9dc6acbc6 // indirect
	k8s.io/api v0.21.2
	k8s.io/apiextensions-apiserver v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
	k8s.io/klog/v2 v2.8.0
	sigs.k8s.io/controller-runtime v0.9.1
)

replace github.com/Azure/azure-service-operator => ../../
