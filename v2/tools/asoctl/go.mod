module github.com/Azure/azure-service-operator/v2/tools/asoctl

go 1.18

// Needed to reference shared version numbering:
replace github.com/Azure/azure-service-operator/v2 => ../../

require (
	github.com/spf13/cobra v1.6.1
	github.com/spf13/pflag v1.0.5
	k8s.io/klog v1.0.0
	k8s.io/klog/v2 v2.80.1
)

require (
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
)
