/*
MIT License

Copyright (c) Microsoft Corporation. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE
*/

package main

import (
	"os"
	"strings"

	"github.com/spf13/pflag"

	servicev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/controllers"
	"github.com/Azure/azure-service-operator/pkg/config"
	"k8s.io/apimachinery/pkg/runtime"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

var (
	scheme                                                      = runtime.NewScheme()
	setupLog                                                    = ctrl.Log.WithName("setup")
	masterURL, kubeconfig, resources, clusterName               string
	cloudName, tenantID, subscriptionID, clientID, clientSecret string
	useAADPodIdentity                                           bool
)

func init() {

	servicev1alpha1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	pflag.StringVarP(&metricsAddr, "metrics-addr", "", ":8080", "The address the metric endpoint binds to.")
	pflag.BoolVarP(&enableLeaderElection, "enable-leader-election", "", false, "Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	pflag.StringVarP(&masterURL, "master-url", "", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig.")
	pflag.StringVarP(&kubeconfig, "kubeconfig", "k", "", "Path to local kubeconfig file (mainly used for development)")
	pflag.StringVarP(&resources, "resources", "", "storage,cosmosdb", "Comma delimited list of CRDs to deploy")
	pflag.StringVarP(&clusterName, "cluster-name", "i", "azure-operator", "Cluster name for the Application to run as, used to avoid conflict")
	pflag.StringVarP(&cloudName, "cloud-name", "c", "AzurePublicCloud", "The cloud name")
	pflag.StringVarP(&tenantID, "tenant-id", "t", "", "The AAD tenant, must provide when using service principals")
	pflag.StringVarP(&subscriptionID, "subscription-id", "s", "", "The subscription ID")
	pflag.StringVarP(&clientID, "client-id", "u", "", "The service principal client ID")
	pflag.StringVarP(&clientSecret, "client-secret", "p", "", "The service principal client secret")
	pflag.BoolVarP(&useAADPodIdentity, "use-aad-pod-identity", "", false, "whether use AAD pod identity")
	pflag.Parse()

	ctrl.SetLogger(zap.Logger(true))

	cfg := config.Config{}
	cfg, err := getConfig()
	if err != nil {
		setupLog.Error(err, "unable to get config")
		os.Exit(1)
	}
	config.Instance = &cfg

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	err = (&controllers.StorageReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("Storage"),
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Storage")
		os.Exit(1)
	}
	err = (&controllers.CosmosDBReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("CosmosDB"),
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "CosmosDB")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func getConfig() (c config.Config, err error) {
	resourcesMap := map[string]bool{}
	for _, r := range strings.Split(resources, ",") {
		resourcesMap[r] = true
	}

	kubeclientset, err := config.CreateKubeClientset(masterURL, kubeconfig)
	if err != nil {
		return c, err
	}

	c = config.Config{
		KubeClientset:     kubeclientset,
		Resources:         resourcesMap,
		ClusterName:       clusterName,
		CloudName:         cloudName,
		TenantID:          tenantID,
		SubscriptionID:    subscriptionID,
		ClientID:          clientID,
		ClientSecret:      clientSecret,
		UseAADPodIdentity: useAADPodIdentity,
	}

	return c, nil
}
