/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var cfg *rest.Config
var k8sClient client.Client
var k8sManager ctrl.Manager
var testEnv *envtest.Environment

// func TestAPIs(t *testing.T) {
// 	resourceGroupName = "t-rg-dev-controller"
// 	resourcegroupLocation = "westus"

// 	eventhubNamespaceName = "t-ns-dev-eh-ns"
// 	eventhubName = "t-eh-dev-sample"
// 	namespaceLocation = "westus"
// 	RunSpecsWithDefaultAndCustomReporters(t,
// 		"Controller Suite",
// 		[]Reporter{envtest.NewlineReporter{}})
// }

func setup() error {
	logf.SetLogger(zap.Logger(true))
	resourcemanagerconfig.LoadSettings()

	if os.Getenv("TEST_USE_EXISTING_CLUSTER") == "true" {
		t := true
		testEnv = &envtest.Environment{
			UseExistingCluster: &t,
		}
	} else {
		testEnv = &envtest.Environment{
			CRDDirectoryPaths: []string{filepath.Join("..", "config", "crd", "bases")},
		}
	}

	cfg, err := testEnv.Start()
	if err != nil {
		return err
	}
	if cfg == nil {
		return fmt.Errorf("testenv config is nil")
	}

	err = azurev1.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	// +kubebuilder:scaffold:scheme
	k8sManager, err = ctrl.NewManager(cfg, ctrl.Options{
		Scheme: scheme.Scheme,
	})
	if err != nil {
		return err
	}

	err = (&EventhubReconciler{
		Client:   k8sManager.GetClient(),
		Log:      ctrl.Log.WithName("controllers").WithName("EventHub"),
		Recorder: k8sManager.GetEventRecorderFor("Eventhub-controller"),
		Scheme:   scheme.Scheme,
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&ResourceGroupReconciler{
		Client:   k8sManager.GetClient(),
		Log:      ctrl.Log.WithName("controllers").WithName("ResourceGroup"),
		Recorder: k8sManager.GetEventRecorderFor("ResourceGroup-controller"),
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&EventhubNamespaceReconciler{
		Client:   k8sManager.GetClient(),
		Log:      ctrl.Log.WithName("controllers").WithName("EventhubNamespace"),
		Recorder: k8sManager.GetEventRecorderFor("EventhubNamespace-controller"),
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&ConsumerGroupReconciler{
		Client:   k8sManager.GetClient(),
		Log:      ctrl.Log.WithName("controllers").WithName("ConsumerGroup"),
		Recorder: k8sManager.GetEventRecorderFor("ConsumerGroup-controller"),
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	go func() {
		err = k8sManager.Start(ctrl.SetupSignalHandler())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	k8sClient = k8sManager.GetClient()
	if k8sClient == nil {
		return fmt.Errorf("k8sClient is nil")
	}

	return nil
}

func teardown() error {
	err := testEnv.Stop()
	return err
}

// TestMain is the main entry point for tests
func TestMain(m *testing.M) {
	var err error
	var code int

	err = setup()
	if err != nil {
		log.Println(fmt.Sprintf("could not set up environment: %v\n", err))
	}

	code = m.Run()

	err = teardown()
	if err != nil {
		log.Println(fmt.Sprintf("could not tear down environment: %v\n; original exit code: %v\n", err, code))
	}

	os.Exit(code)
}
