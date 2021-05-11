/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Azure/go-autorest/autorest/azure/auth"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"github.com/Azure/k8s-infra/hack/generated/controllers"
	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
)

func main() {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	klog.InitFlags(flagSet)
	setupLog := ctrl.Log.WithName("setup")

	var metricsAddr string
	var enableLeaderElection bool
	flagSet.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flagSet.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controllers manager. Enabling this will ensure there is only one active controllers manager.")
	flagSet.Parse(os.Args[1:]) //nolint:errcheck

	scheme := controllers.CreateScheme()

	ctrl.SetLogger(klogr.New())

	syncPeriod := 30 * time.Minute

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "controllers-leader-election-azinfra-generated",
		Port:               9443,
		SyncPeriod:         &syncPeriod, // TODO: Don't leave this set, just playing with it
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	authorizer, err := armclient.AuthorizerFromEnvironment()
	if err != nil {
		setupLog.Error(err, "unable to get authorization settings")
		os.Exit(1)
	}

	subID := os.Getenv(auth.SubscriptionID)
	if subID == "" {
		setupLog.Error(err, fmt.Sprintf("unable to get env var %q", auth.SubscriptionID))
		os.Exit(1)
	}

	armApplier, err := armclient.NewAzureTemplateClient(authorizer, subID)
	if err != nil {
		setupLog.Error(err, "failed to create ARM applier")
		os.Exit(1)
	}

	if errs := controllers.RegisterAll(mgr, armApplier, controllers.GetKnownStorageTypes(), ctrl.Log.WithName("controllers"), concurrency(1)); errs != nil {
		setupLog.Error(err, "failed to register gvks")
		os.Exit(1)
	}

	if errs := controllers.RegisterWebhooks(mgr, controllers.GetKnownTypes()); errs != nil {
		setupLog.Error(err, "failed to register webhook for gvks")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func concurrency(c int) controllers.Options {
	return controllers.Options{
		Options: controller.Options{
			MaxConcurrentReconciles: c,
		},
	}
}
