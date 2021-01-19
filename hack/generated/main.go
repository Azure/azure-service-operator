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
	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"

	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	batch "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.batch/v20170901"
	resources "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.resources/v20200601"
	storage "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.storage/v20190401"
	"github.com/Azure/k8s-infra/hack/generated/controllers"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	klog.InitFlags(nil)

	// TODO: Need to generate this
	_ = clientgoscheme.AddToScheme(scheme)
	_ = batch.AddToScheme(scheme)
	_ = storage.AddToScheme(scheme)
	_ = resources.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controllers manager. Enabling this will ensure there is only one active controllers manager.")
	flag.Parse()

	ctrl.SetLogger(klogr.New())

	// syncPeriod := 30 * time.Second
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

	if errs := controllers.RegisterAll(mgr, armApplier, controllers.KnownTypes, ctrl.Log.WithName("controllers"), concurrency(1)); errs != nil {
		for _, err := range errs {
			setupLog.Error(err, "failed to register gvk: %v")
		}
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
