module github.com/Azure/azure-service-operator/v2/tools/generator

go 1.20

// Needed to reference shared version numbering:
replace github.com/Azure/azure-service-operator/v2 => ../../

// Modified version that doesn’t panic on golang-invalid regexes:
replace github.com/xeipuuv/gojsonschema => github.com/devigned/gojsonschema v1.2.1-0.20191231010529-c593123f1e5d

require (
	github.com/Azure/azure-service-operator/v2 v2.0.0
	github.com/bmatcuk/doublestar v1.3.4
	github.com/dave/dst v0.27.2
	github.com/devigned/tab v0.1.1
	github.com/go-logr/logr v1.2.4
	github.com/go-logr/zerologr v1.2.3
	github.com/go-openapi/jsonpointer v0.19.6
	github.com/go-openapi/spec v0.20.9
	github.com/gobuffalo/flect v1.0.2
	github.com/google/go-cmp v0.5.9
	github.com/kr/pretty v0.3.1
	github.com/kylelemons/godebug v1.1.0
	github.com/leanovate/gopter v0.2.9
	github.com/onsi/gomega v1.27.7
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.29.1
	github.com/sebdah/goldie/v2 v2.5.3
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/exp v0.0.0-20230510235704-dd950f8aeaea
	golang.org/x/mod v0.10.0
	golang.org/x/net v0.10.0
	golang.org/x/sync v0.2.0
	golang.org/x/text v0.9.0
	gopkg.in/yaml.v3 v3.0.1
	k8s.io/apimachinery v0.27.2
)

require (
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/hbollon/go-edlib v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/sergi/go-diff v1.3.1 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/tools v0.9.1 // indirect
)
