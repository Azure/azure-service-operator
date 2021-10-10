/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package main

import (
	"flag"
	"os"
	"time"

	"github.com/go-logr/logr"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"github.com/Azure/azure-service-operator/v2/internal/armclient"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/Azure/azure-service-operator/v2/internal/version"
)

func main() {
	exeName := os.Args[0] + " " + version.BuildVersion
	flagSet := flag.NewFlagSet(exeName, flag.ExitOnError)
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

	cfg, err := config.ReadAndValidate()
	if err != nil {
		setupLog.Error(err, "unable to get env configuration values")
		os.Exit(1)
	}

	var cacheFunc cache.NewCacheFunc
	if cfg.TargetNamespaces != nil {
		cacheFunc = cache.MultiNamespacedCacheBuilder(cfg.TargetNamespaces)
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		NewCache:           cacheFunc,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "controllers-leader-election-azinfra-generated",
		Port:               9443,
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

	armApplier, err := armclient.NewAzureTemplateClient(authorizer, cfg.SubscriptionID)
	if err != nil {
		setupLog.Error(err, "failed to create ARM applier")
		os.Exit(1)
	}

	log := ctrl.Log.WithName("controllers")
	if cfg.OperatorMode.IncludesWatchers() {
		if errs := controllers.RegisterAll(mgr, armApplier, controllers.GetKnownStorageTypes(), makeControllerOptions(log, cfg)); errs != nil {
			setupLog.Error(err, "failed to register gvks")
			os.Exit(1)
		}
	}

	if cfg.OperatorMode.IncludesWebhooks() {
		if errs := controllers.RegisterWebhooks(mgr, controllers.GetKnownTypes()); errs != nil {
			setupLog.Error(err, "failed to register webhook for gvks")
			os.Exit(1)
		}
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func makeControllerOptions(log logr.Logger, cfg config.Values) controllers.Options {
	return controllers.Options{
		Config: cfg,
		Options: controller.Options{
			MaxConcurrentReconciles: 1,
			Log:                     log,
			RateLimiter:             controllers.NewRateLimiter(1*time.Second, 1*time.Minute),
		},
	}
}
