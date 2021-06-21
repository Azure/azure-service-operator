module github.com/Azure/azure-service-operator/hack/generator

go 1.16

require (
	github.com/bmatcuk/doublestar v1.3.4
	github.com/dave/dst v0.26.2
	github.com/devigned/tab v0.1.1
	github.com/go-openapi/jsonpointer v0.19.5
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.3
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/gobuffalo/flect v0.2.3
	github.com/leanovate/gopter v0.2.9
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/onsi/gomega v1.13.0
	github.com/pkg/errors v0.9.1
	github.com/sebdah/goldie/v2 v2.5.3
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/mod v0.4.2
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	golang.org/x/tools v0.1.3 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/apimachinery v0.21.2
	k8s.io/klog/v2 v2.9.0
)

replace github.com/xeipuuv/gojsonschema => github.com/devigned/gojsonschema v1.2.1-0.20191231010529-c593123f1e5d
