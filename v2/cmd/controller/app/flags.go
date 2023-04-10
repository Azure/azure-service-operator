/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package app

import (
	"flag"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	internalflags "github.com/Azure/azure-service-operator/v2/internal/util/flags"
	"github.com/Azure/azure-service-operator/v2/internal/version"
)

type Flags struct {
	metricsAddr          string
	healthAddr           string
	enableLeaderElection bool
	crdPatterns          []string
}

func ParseFlags(args []string) (Flags, error) {
	exeName := args[0] + " " + version.BuildVersion
	flagSet := flag.NewFlagSet(exeName, flag.ExitOnError)
	klog.InitFlags(flagSet)

	var metricsAddr string
	var healthAddr string
	var enableLeaderElection bool
	var crdPatterns internalflags.SliceFlags
	var preUpgradeCheck bool

	// default here for 'metricsAddr' is set to "0", which sets metrics to be disabled if 'metrics-addr' flag is omitted.
	flagSet.StringVar(&metricsAddr, "metrics-addr", "0", "The address the metric endpoint binds to.")
	flagSet.StringVar(&healthAddr, "health-addr", "", "The address the healthz endpoint binds to.")
	flagSet.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controllers manager. Enabling this will ensure there is only one active controllers manager.")
	flagSet.Var(&crdPatterns, "crd-pattern", "Install these CRDs. Currently the only value supported is '*'")
	flagSet.BoolVar(&preUpgradeCheck, "pre-upgrade-check", false,
		"Enable pre upgrade check to check if existing crds contain helm 'keep' policy.")

	flagSet.Parse(args[1:]) //nolint:errcheck

	if len(crdPatterns) >= 2 {
		return Flags{}, errors.Errorf("crd-pattern flag can have at most 1 argument, had %d", len(crdPatterns))
	}
	if len(crdPatterns) == 1 && crdPatterns[0] != "*" {
		return Flags{}, errors.Errorf("crd-pattern flag must be '*', was %q", crdPatterns[0])
	}

	return Flags{
		metricsAddr:          metricsAddr,
		healthAddr:           healthAddr,
		enableLeaderElection: enableLeaderElection,
		crdPatterns:          crdPatterns,
	}, nil
}
