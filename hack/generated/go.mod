module github.com/Azure/azure-service-operator/hack/generated

go 1.16

require (
	github.com/Azure/go-autorest/autorest v0.11.19
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.0
	github.com/Azure/go-autorest/autorest/date v0.3.0
	github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/armon/consul-api v0.0.0-20180202201655-eb2c6b5be1b6 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/coreos/go-etcd v2.0.0+incompatible // indirect
	github.com/cpuguy83/go-md2man v1.0.10 // indirect
	github.com/devigned/tab v0.1.1
	github.com/dnaeon/go-vcr v1.1.0
	github.com/docker/docker v0.7.3-0.20190327010347-be7ac8be2ae0 // indirect
	github.com/docker/spdystream v0.0.0-20160310174837-449fdfce4d96 // indirect
	github.com/go-logr/logr v0.4.0
	github.com/go-openapi/validate v0.19.5 // indirect
	github.com/google/go-cmp v0.5.5
	github.com/google/uuid v1.1.2
	github.com/gophercloud/gophercloud v0.1.0 // indirect
	github.com/kr/pretty v0.2.0
	github.com/kr/text v0.2.0 // indirect
	github.com/kylelemons/godebug v1.1.0
	github.com/leanovate/gopter v0.2.8
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/gomega v1.14.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/ugorji/go/codec v0.0.0-20181204163529-d75b2dcb6bc8 // indirect
	github.com/xordataexchange/crypt v0.0.3-0.20170626215501-b2862e3d0a77 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gotest.tools v2.2.0+incompatible // indirect
	k8s.io/api v0.21.2
	k8s.io/apiextensions-apiserver v0.21.2 // indirect
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
	k8s.io/klog v1.0.0 // indirect
	k8s.io/klog/v2 v2.8.0
	sigs.k8s.io/controller-runtime v0.9.1
	sigs.k8s.io/structured-merge-diff/v3 v3.0.0 // indirect
)

replace github.com/Azure/azure-service-operator => ../../
