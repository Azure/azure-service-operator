/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package app

import (
	"flag"
	"fmt"
)

type Flags struct {
	MetricsAddr          *string
	ProfilingMetrics     *bool
	SecureMetrics        *bool
	HealthAddr           *string
	WebhookPort          *int
	WebhookCertDir       *string
	EnableLeaderElection *bool
	CRDManagementMode    *string
	CRDPatterns          *string // This is a ';' delimited string containing a collection of patterns
}

func (f Flags) String() string {
	return fmt.Sprintf(
		"MetricsAddr: %s, SecureMetrics: %t, ProfilingMetrics: %t, HealthAddr: %s, WebhookPort: %d, WebhookCertDir: %s, EnableLeaderElection: %t, CRDManagementMode: %s, CRDPatterns: %s",
		*f.MetricsAddr,
		*f.SecureMetrics,
		*f.ProfilingMetrics,
		*f.HealthAddr,
		*f.WebhookPort,
		*f.WebhookCertDir,
		*f.EnableLeaderElection,
		*f.CRDManagementMode,
		*f.CRDPatterns)
}

func InitFlags(flagSet *flag.FlagSet) Flags {

	// default here for 'MetricsAddr' is set to "0", which sets metrics to be disabled if 'metrics-addr' flag is omitted.
	metricsAddr := flagSet.String("metrics-addr", "0", "The address the metric endpoint binds to.")
	secureMetrics := flagSet.Bool("secure-metrics", true, "Enable secure metrics. This secures the pprof and metrics endpoints via Kubernetes RBAC and HTTPS")
	profilingMetrics := flagSet.Bool("profiling-metrics", false, "Enable pprof metrics, only enabled in conjunction with secure-metrics. This will enable serving pprof metrics endpoints")

	healthAddr := flagSet.String("health-addr", "", "The address the healthz endpoint binds to.")
	webhookPort := flagSet.Int("webhook-port", 9443, "The port the webhook endpoint binds to.")
	webhookCertDir := flagSet.String("webhook-cert-dir", "", "The directory the webhook server's certs are stored.")

	enableLeaderElection := flagSet.Bool("enable-leader-election", false, "Enable leader election for controllers manager. Enabling this will ensure there is only one active controllers manager.")

	crdManagementMode := flagSet.String("crd-management", "auto",
		"Instructs the operator on how it should manage the Custom Resource Definitions. One of 'auto', 'none'")
	crdPatterns := flagSet.String("crd-pattern", "", "Install these CRDs. CRDs already in the cluster will also always be upgraded.")

	return Flags{
		MetricsAddr:      metricsAddr,
		SecureMetrics:    secureMetrics,
		ProfilingMetrics: profilingMetrics,

		HealthAddr:           healthAddr,
		WebhookPort:          webhookPort,
		WebhookCertDir:       webhookCertDir,
		EnableLeaderElection: enableLeaderElection,
		CRDManagementMode:    crdManagementMode,
		CRDPatterns:          crdPatterns,
	}
}
