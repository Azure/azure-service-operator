module github.com/Azure/azure-service-operator/v2/tools/generator

go 1.18

// Needed to reference shared version numbering:
replace github.com/Azure/azure-service-operator/v2 => ../../

// Modified version that doesn’t panic on golang-invalid regexes:
replace github.com/xeipuuv/gojsonschema => github.com/devigned/gojsonschema v1.2.1-0.20191231010529-c593123f1e5d

require (
	github.com/Azure/azure-service-operator/v2 v2.0.0-00010101000000-000000000000
	github.com/bmatcuk/doublestar v1.3.4
	github.com/dave/dst v0.26.2
	github.com/devigned/tab v0.1.1
	github.com/go-openapi/jsonpointer v0.19.5
	github.com/go-openapi/spec v0.20.4
	github.com/gobuffalo/flect v0.2.3
	github.com/google/go-cmp v0.5.9
	github.com/hbollon/go-edlib v1.6.0
	github.com/kylelemons/godebug v1.1.0
	github.com/leanovate/gopter v0.2.9
	github.com/onsi/gomega v1.20.1
	github.com/pkg/errors v0.9.1
	github.com/sebdah/goldie/v2 v2.5.3
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/exp v0.0.0-20220414153411-bcd21879b8fd
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4
	golang.org/x/net v0.0.0-20221004154528-8021a29435af
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/apimachinery v0.25.2
	k8s.io/klog/v2 v2.80.1
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.6.1 // indirect
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec // indirect
	golang.org/x/text v0.3.8 // indirect
	golang.org/x/tools v0.1.12 // indirect
)
