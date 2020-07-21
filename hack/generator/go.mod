module github.com/Azure/k8s-infra/hack/generator

go 1.13

require (
	github.com/bmatcuk/doublestar v1.3.1
	github.com/devigned/tab v0.1.1
	github.com/gobuffalo/flect v0.2.1
	github.com/onsi/gomega v1.10.1
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sebdah/goldie/v2 v2.3.0
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.1
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415
	github.com/xeipuuv/gojsonschema v1.2.0
	gopkg.in/ini.v1 v1.51.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
	k8s.io/apimachinery v0.18.4
	k8s.io/klog/v2 v2.0.0
	sigs.k8s.io/controller-runtime v0.6.1
)

replace github.com/xeipuuv/gojsonschema => github.com/devigned/gojsonschema v1.2.1-0.20191231010529-c593123f1e5d
