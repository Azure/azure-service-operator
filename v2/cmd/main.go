/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/go-logr/logr"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	"github.com/Azure/azure-service-operator/v2/internal/controller/armclient"
	"github.com/Azure/azure-service-operator/v2/internal/controller/config"
	"github.com/Azure/azure-service-operator/v2/internal/controller/controllers"
	"github.com/Azure/azure-service-operator/v2/internal/controller/version"
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

	targetNamespaces := parseTargetNamespaces(os.Getenv(targetNamespacesVar))
	var cacheFunc cache.NewCacheFunc
	if targetNamespaces != nil {
		cacheFunc = cache.MultiNamespacedCacheBuilder(targetNamespaces)
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

	subID := os.Getenv(auth.SubscriptionID)
	if subID == "" {
		setupLog.Error(err, fmt.Sprintf("unable to get env var %q", auth.SubscriptionID))
		os.Exit(1)
	}

	var selectedMode config.OperatorMode
	if modeValue := os.Getenv(operatorModeVar); modeValue == "" {
		selectedMode = config.OperatorModeBoth
	} else {
		selectedMode, err = config.ParseOperatorMode(modeValue)
		if err != nil {
			setupLog.Error(err, "getting operator mode")
		}
	}

	armApplier, err := armclient.NewAzureTemplateClient(authorizer, subID)
	if err != nil {
		setupLog.Error(err, "failed to create ARM applier")
		os.Exit(1)
	}

	log := ctrl.Log.WithName("controllers")
	if selectedMode.IncludesWatchers() {
		if errs := controllers.RegisterAll(mgr, armApplier, controllers.GetKnownStorageTypes(), makeControllerOptions(log)); errs != nil {
			setupLog.Error(err, "failed to register gvks")
			os.Exit(1)
		}
	}

	if selectedMode.IncludesWebhooks() {
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

func makeControllerOptions(log logr.Logger) controllers.Options {
	return controllers.Options{
		Options: controller.Options{
			MaxConcurrentReconciles: 1,
			Log:                     log,
			RateLimiter:             controllers.NewRateLimiter(1*time.Second, 1*time.Minute),
		},
	}
}

const (
	targetNamespacesVar = "AZURE_TARGET_NAMESPACES"
	operatorModeVar     = "AZURE_OPERATOR_MODE"
)

func parseTargetNamespaces(fromEnv string) []string {
	if len(strings.TrimSpace(fromEnv)) == 0 {
		return nil
	}
	items := strings.Split(fromEnv, ",")
	// Remove any whitespace used to separate items.
	for i, item := range items {
		items[i] = strings.TrimSpace(item)
	}
	return items
}
