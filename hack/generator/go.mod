module github.com/Azure/k8s-infra/hack/generator

go 1.13

require (
	github.com/bmatcuk/doublestar v1.3.1
	github.com/devigned/tab v0.1.1
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-openapi/jsonpointer v0.19.3
	github.com/go-openapi/jsonreference v0.19.3 // indirect
	github.com/go-openapi/spec v0.19.3
	github.com/gobuffalo/flect v0.2.1
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1
	github.com/sebdah/goldie/v2 v2.3.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	golang.org/x/text v0.3.3 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
	k8s.io/apimachinery v0.18.6
	k8s.io/klog/v2 v2.0.0
)

replace github.com/xeipuuv/gojsonschema => github.com/devigned/gojsonschema v1.2.1-0.20191231010529-c593123f1e5d
