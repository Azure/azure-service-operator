/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package app

import (
	"flag"
	"fmt"

	"k8s.io/klog/v2"

	"github.com/Azure/azure-service-operator/v2/internal/version"
)

type Flags struct {
	MetricsAddr          string
	ProfilingMetrics     bool
	SecureMetrics        bool
	HealthAddr           string
	WebhookPort          int
	WebhookCertDir       string
	EnableLeaderElection bool
	CRDManagementMode    string
	CRDPatterns          string // This is a ';' delimited string containing a collection of patterns
}

func (f Flags) String() string {
	return fmt.Sprintf(
		"MetricsAddr: %s, SecureMetrics: %t, ProfilingMetrics: %t, HealthAddr: %s, WebhookPort: %d, WebhookCertDir: %s, EnableLeaderElection: %t, CRDManagementMode: %s, CRDPatterns: %s",
		f.MetricsAddr,
		f.SecureMetrics,
		f.ProfilingMetrics,
		f.HealthAddr,
		f.WebhookPort,
		f.WebhookCertDir,
		f.EnableLeaderElection,
		f.CRDManagementMode,
		f.CRDPatterns)
}

func ParseFlags(args []string) (Flags, error) {
	exeName := args[0] + " " + version.BuildVersion
	flagSet := flag.NewFlagSet(exeName, flag.ExitOnError)
	klog.InitFlags(flagSet)

	var metricsAddr string
	var profilingMetrics bool
	var secureMetrics bool
	var healthAddr string
	var webhookPort int
	var webhookCertDir string
	var enableLeaderElection bool
	var crdManagementMode string
	var crdPatterns string

	// default here for 'MetricsAddr' is set to "0", which sets metrics to be disabled if 'metrics-addr' flag is omitted.
	flagSet.StringVar(&metricsAddr, "metrics-addr", "0", "The address the metric endpoint binds to.")
	flagSet.BoolVar(&secureMetrics, "secure-metrics", true, "Enable secure metrics. This secures the pprof and metrics endpoints via Kubernetes RBAC and HTTPS")
	flagSet.BoolVar(&profilingMetrics, "profiling-metrics", false, "Enable pprof metrics, only enabled in conjunction with secure-metrics. This will enable serving pprof metrics endpoints")

	flagSet.StringVar(&healthAddr, "health-addr", "", "The address the healthz endpoint binds to.")
	flagSet.IntVar(&webhookPort, "webhook-port", 9443, "The port the webhook endpoint binds to.")
	flagSet.StringVar(&webhookCertDir, "webhook-cert-dir", "", "The directory the webhook server's certs are stored.")
	flagSet.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controllers manager. Enabling this will ensure there is only one active controllers manager.")
	flagSet.StringVar(&crdManagementMode, "crd-management", "auto",
		"Instructs the operator on how it should manage the Custom Resource Definitions. One of 'auto', 'none'")
	flagSet.StringVar(&crdPatterns, "crd-pattern", "", "Install these CRDs. CRDs already in the cluster will also always be upgraded.")

	flagSet.Parse(args[1:]) //nolint:errcheck

	return Flags{
		MetricsAddr:          metricsAddr,
		SecureMetrics:        secureMetrics,
		HealthAddr:           healthAddr,
		WebhookPort:          webhookPort,
		WebhookCertDir:       webhookCertDir,
		EnableLeaderElection: enableLeaderElection,
		CRDManagementMode:    crdManagementMode,
		CRDPatterns:          crdPatterns,
	}, nil
}
