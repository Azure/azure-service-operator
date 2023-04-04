/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package main

import (
	"os"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/azure-service-operator/v2/cmd/controller/app"
)

func main() {
	setupLog := ctrl.Log.WithName("setup")
	ctrl.SetLogger(klogr.New())
	ctx := ctrl.SetupSignalHandler()

	flgs, err := app.ParseFlags(os.Args)
	if err != nil {
		setupLog.Error(err, "failed to parse cmdline flags")
		os.Exit(1)
	}

	mgr := app.SetupControllerManager(setupLog, flgs, ctx)

	setupLog.Info("starting manager")
	if err = mgr.Start(ctx); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
