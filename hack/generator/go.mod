module github.com/Azure/k8s-infra/hack/generator

go 1.13

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.0
	github.com/devigned/tab v0.1.1
	github.com/devigned/tab/opencensus v0.1.2
	github.com/golang/groupcache v0.0.0-20191227052852-215e87163ea7 // indirect
	github.com/onsi/gomega v1.8.1
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/sebdah/goldie/v2 v2.3.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.1
	github.com/uber/jaeger-client-go v2.21.1+incompatible // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonschema v1.2.0
	go.opencensus.io v0.22.2
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/sys v0.0.0-20191228213918-04cbcbbfeed8 // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543 // indirect
	google.golang.org/api v0.15.0 // indirect
	gopkg.in/ini.v1 v1.51.1 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace github.com/xeipuuv/gojsonschema => github.com/devigned/gojsonschema v1.2.1-0.20191231010529-c593123f1e5d
