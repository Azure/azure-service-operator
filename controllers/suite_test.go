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
	"context"
	"os"
	"path/filepath"
	"testing"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

	eventhubs "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resoucegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

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
var resourceGroupName string
var resourcegroupLocation string
var eventhubNamespaceName string
var namespaceLocation string

func TestAPIs(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)
	resourceGroupName = "t-rg-dev-controller"
	resourcegroupLocation = "westus"

	eventhubNamespaceName = "t-ns-dev-eh-ns"
	namespaceLocation = "westus"
	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{envtest.NewlineReporter{}})
}

var _ = BeforeSuite(func(done Done) {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	By("bootstrapping test environment")

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

	resourcemanagerconfig.LoadSettings()

	cfg, err := testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(cfg).ToNot(BeNil())

	err = azurev1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = azurev1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = azurev1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = azurev1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	// +kubebuilder:scaffold:scheme
	k8sManager, err = ctrl.NewManager(cfg, ctrl.Options{
		Scheme: scheme.Scheme,
	})
	Expect(err).ToNot(HaveOccurred())

	err = (&EventhubReconciler{
		Client:   k8sManager.GetClient(),
		Log:      ctrl.Log.WithName("controllers").WithName("EventHub"),
		Recorder: k8sManager.GetEventRecorderFor("Eventhub-controller"),
		Scheme:   scheme.Scheme,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&ResourceGroupReconciler{
		Client:   k8sManager.GetClient(),
		Log:      ctrl.Log.WithName("controllers").WithName("ResourceGroup"),
		Recorder: k8sManager.GetEventRecorderFor("ResourceGroup-controller"),
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&EventhubNamespaceReconciler{
		Client:   k8sManager.GetClient(),
		Log:      ctrl.Log.WithName("controllers").WithName("EventhubNamespace"),
		Recorder: k8sManager.GetEventRecorderFor("EventhubNamespace-controller"),
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	go func() {
		err = k8sManager.Start(ctrl.SetupSignalHandler())
		Expect(err).ToNot(HaveOccurred())
	}()

	k8sClient = k8sManager.GetClient()
	Expect(k8sClient).ToNot(BeNil())

	// Create the Resourcegroup resource
	result, _ := resoucegroupsresourcemanager.CheckExistence(context.Background(), resourceGroupName)
	if result.Response.StatusCode != 204 {
		_, _ = resoucegroupsresourcemanager.CreateGroup(context.Background(), resourceGroupName, resourcegroupLocation)
	}

	// Create the Eventhub namespace resource
	_, err = eventhubs.CreateNamespaceAndWait(context.Background(), resourceGroupName, eventhubNamespaceName, namespaceLocation)

	close(done)
}, 120)

var _ = AfterSuite(func(done Done) {
	//clean up the resources created for test

	By("tearing down the test environment")

	_, _ = resoucegroupsresourcemanager.DeleteGroup(context.Background(), resourceGroupName)

	err := testEnv.Stop()
	Expect(err).ToNot(HaveOccurred())
	close(done)

}, 60)
