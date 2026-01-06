module github.com/Azure/azure-service-operator/v2/tools/generator

go 1.24.4

// Needed to reference shared version numbering:
replace github.com/Azure/azure-service-operator/v2 => ../../

// Modified version that doesn’t panic on golang-invalid regexes:
replace github.com/xeipuuv/gojsonschema => github.com/devigned/gojsonschema v1.2.1-0.20191231010529-c593123f1e5d

require (
	github.com/Azure/azure-service-operator/v2 v2.12.0
	github.com/bmatcuk/doublestar v1.3.4
	github.com/dave/dst v0.27.3
	github.com/devigned/tab v0.1.1
	github.com/go-logr/logr v1.4.3
	github.com/go-logr/zerologr v1.2.3
	github.com/go-openapi/jsonpointer v0.22.3
	github.com/go-openapi/spec v0.22.1
	github.com/gobuffalo/flect v1.0.3
	github.com/google/go-cmp v0.7.0
	github.com/kr/pretty v0.3.1
	github.com/kylelemons/godebug v1.1.0
	github.com/leanovate/gopter v0.2.11
	github.com/onsi/gomega v1.38.2
	github.com/rotisserie/eris v0.5.4
	github.com/rs/zerolog v1.34.0
	github.com/sebdah/goldie/v2 v2.8.0
	github.com/spf13/cobra v1.10.2
	github.com/spf13/pflag v1.0.10
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394
	golang.org/x/mod v0.30.0
	golang.org/x/net v0.47.0
	golang.org/x/sync v0.18.0
	golang.org/x/text v0.31.0
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/apimachinery v0.34.2
)

require (
	github.com/go-openapi/jsonreference v0.21.3 // indirect
	github.com/go-openapi/swag/conv v0.25.1 // indirect
	github.com/go-openapi/swag/jsonname v0.25.1 // indirect
	github.com/go-openapi/swag/jsonutils v0.25.1 // indirect
	github.com/go-openapi/swag/loading v0.25.1 // indirect
	github.com/go-openapi/swag/stringutils v0.25.1 // indirect
	github.com/go-openapi/swag/typeutils v0.25.1 // indirect
	github.com/go-openapi/swag/yamlutils v0.25.1 // indirect
	github.com/hbollon/go-edlib v1.7.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/tools v0.38.0 // indirect
)
