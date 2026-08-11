/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package app

import (
	"crypto/tls"
	"flag"
	"fmt"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/internal/crdmanagement"
	"github.com/Azure/azure-service-operator/v2/internal/labels"
)

type Flags struct {
	MetricsAddr          string
	ProfilingMetrics     bool
	SecureMetrics        bool
	MetricsCertDir       string
	HealthAddr           string
	WebhookPort          int
	WebhookCertDir       string
	EnableLeaderElection bool
	CRDManagementMode    string
	CRDPatterns          string // This is a ';' delimited string containing a collection of patterns
	CRDLabels            string // This is a ',' or ';' delimited string containing labels to apply to managed CRDs
	TLSMinVersion        string
}

// parseCRDLabels parses the --crd-labels flag, rejecting any label reserved for ASO's own use.
func parseCRDLabels(value string) (map[string]string, error) {
	result, err := labels.ParseMap(value)
	if err != nil {
		return nil, err
	}

	// Done as a separate pass so that the general purpose label parsing stays free of CRD specific rules.
	for key := range result {
		if crdmanagement.IsReservedLabel(key) {
			return nil, eris.Errorf("label %q is reserved for use by Azure Service Operator and cannot be overridden", key)
		}
	}

	return result, nil
}

func (f Flags) String() string {
	return fmt.Sprintf(
		"MetricsAddr: %s, SecureMetrics: %t, ProfilingMetrics: %t, MetricsCertDir: %s, HealthAddr: %s, WebhookPort: %d, WebhookCertDir: %s, EnableLeaderElection: %t, CRDManagementMode: %s, CRDPatterns: %s, CRDLabels: %s, TLSMinVersion: %s",
		f.MetricsAddr,
		f.SecureMetrics,
		f.ProfilingMetrics,
		f.MetricsCertDir,
		f.HealthAddr,
		f.WebhookPort,
		f.WebhookCertDir,
		f.EnableLeaderElection,
		f.CRDManagementMode,
		f.CRDPatterns,
		f.CRDLabels,
		f.TLSMinVersion,
	)
}

func InitFlags(flagSet *flag.FlagSet) *Flags {
	result := &Flags{}

	// default here for 'MetricsAddr' is set to "0", which sets metrics to be disabled if 'metrics-addr' flag is omitted.
	flagSet.StringVar(&result.MetricsAddr, "metrics-addr", "0", "The address the metric endpoint binds to.")
	flagSet.BoolVar(&result.SecureMetrics, "secure-metrics", true, "Enable secure metrics. This secures the pprof and metrics endpoints via Kubernetes RBAC and HTTPS")
	flagSet.BoolVar(&result.ProfilingMetrics, "profiling-metrics", false, "Enable pprof metrics, only enabled in conjunction with secure-metrics. This will enable serving pprof metrics endpoints")
	flagSet.StringVar(&result.MetricsCertDir, "metrics-cert-dir", "", "The directory the metrics server's certs are stored.")
	flagSet.StringVar(&result.HealthAddr, "health-addr", "", "The address the healthz endpoint binds to.")
	flagSet.IntVar(&result.WebhookPort, "webhook-port", 9443, "The port the webhook endpoint binds to.")
	flagSet.StringVar(&result.WebhookCertDir, "webhook-cert-dir", "", "The directory the webhook server's certs are stored.")
	flagSet.BoolVar(&result.EnableLeaderElection, "enable-leader-election", false, "Enable leader election for controllers manager. Enabling this will ensure there is only one active controllers manager.")

	flagSet.StringVar(&result.CRDManagementMode, "crd-management", "auto",
		"Instructs the operator on how it should manage the Custom Resource Definitions. One of 'auto', 'none'")
	flagSet.StringVar(&result.CRDPatterns, "crd-pattern", "", "Install these CRDs. CRDs already in the cluster will also always be upgraded.")
	flagSet.StringVar(&result.CRDLabels, "crd-labels", "", "Comma-separated (or semicolon-separated) labels to apply to all managed CRDs (for example, example.com/owner=aso,environment=production). Labels reserved by the operator (app.kubernetes.io/name, app.kubernetes.io/version and the serviceoperator.azure.com/ prefix) cannot be set.")
	flagSet.StringVar(&result.TLSMinVersion, "tls-min-version", "VersionTLS12", "The minimum TLS version in use by the webhook and metrics servers. Possible values: VersionTLS12, VersionTLS13.")

	return result
}

var tlsVersionMap = map[string]uint16{
	"VersionTLS12": tls.VersionTLS12,
	"VersionTLS13": tls.VersionTLS13,
}

func (f Flags) TLSVersion() (uint16, error) {
	v, ok := tlsVersionMap[f.TLSMinVersion]
	if !ok {
		return 0, fmt.Errorf("invalid TLS version %q, must be one of: VersionTLS12, VersionTLS13", f.TLSMinVersion)
	}
	return v, nil
}
