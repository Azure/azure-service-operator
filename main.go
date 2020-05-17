/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package main

import (
	"flag"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	microsoftnetworkv1 "github.com/Azure/k8s-infra/apis/microsoft.network/v1"
	microsoftnetworkv20191101 "github.com/Azure/k8s-infra/apis/microsoft.network/v20191101"
	microsoftresourcesv1 "github.com/Azure/k8s-infra/apis/microsoft.resources/v1"
	microsoftresourcesv20150101 "github.com/Azure/k8s-infra/apis/microsoft.resources/v20150101"
	microsoftresourcesv20191001 "github.com/Azure/k8s-infra/apis/microsoft.resources/v20191001"
	"github.com/Azure/k8s-infra/controllers"
	"github.com/Azure/k8s-infra/pkg/zips"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	klog.InitFlags(nil)

	_ = clientgoscheme.AddToScheme(scheme)
	_ = microsoftresourcesv20191001.AddToScheme(scheme)
	_ = microsoftresourcesv20150101.AddToScheme(scheme)
	_ = microsoftresourcesv1.AddToScheme(scheme)
	_ = microsoftnetworkv20191101.AddToScheme(scheme)
	_ = microsoftnetworkv1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	ctrl.SetLogger(klogr.New())

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "controller-leader-election-azinfra",
		Port:               9443,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	applier, err := zips.NewAzureTemplateClient()
	if err != nil {
		setupLog.Error(err, "failed to create zips Applier.")
		os.Exit(1)
	}

	if errs := controllers.RegisterAll(mgr, applier, controllers.KnownTypes, ctrl.Log.WithName("controllers"), concurrency(1)); errs != nil {
		for _, err := range errs {
			setupLog.Error(err, "failed to register gvk: %v")
		}
		os.Exit(1)
	}

	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func concurrency(c int) controller.Options {
	return controller.Options{MaxConcurrentReconciles: c}
}
