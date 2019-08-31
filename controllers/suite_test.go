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
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"
	"k8s.io/client-go/rest"
	"log"
	"os"
	"path/filepath"
	"testing"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	resoucegroupsconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resoucegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var testEnv *envtest.Environment

type TestContext struct {
	Cfg                   rest.Config
	K8sClient             client.Client
	ResourceGroupName     string
	ResourcegroupLocation string
	EventhubNamespaceName string
	EventhubName          string
	NamespaceLocation     string
	StorageAccountName    string
	BlobContainerName     string
}

var tc TestContext

func TestAPIs(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{envtest.NewlineReporter{}})
}

var _ = SynchronizedBeforeSuite(func() []byte {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	resoucegroupsconfig.ParseEnvironment()
	resourceGroupName := "t-rg-dev-controller-" + helpers.RandomString(10)
	resourcegroupLocation := resoucegroupsconfig.DefaultLocation()

	eventhubNamespaceName := "t-ns-dev-eh-ns-" + helpers.RandomString(10)
	eventhubName := "t-eh-dev-sample-" + helpers.RandomString(10)
	namespaceLocation := resoucegroupsconfig.DefaultLocation()

	storageAccountName := "tsadeveh" + helpers.RandomString(10)
	blobContainerName := "t-bc-dev-eh-" + helpers.RandomString(10)

	log.Println(fmt.Sprintf("common test setup"))
	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{filepath.Join("..", "config", "crd", "bases")},
	}

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

	var k8sManager ctrl.Manager
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

	err = (&ConsumerGroupReconciler{
		Client:   k8sManager.GetClient(),
		Log:      ctrl.Log.WithName("controllers").WithName("ConsumerGroup"),
		Recorder: k8sManager.GetEventRecorderFor("ConsumerGroup-controller"),
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	go func() {
		err = k8sManager.Start(ctrl.SetupSignalHandler())
		Expect(err).ToNot(HaveOccurred())
	}()

	//k8sClient = k8sManager.GetClient()
	k8sClient, _ := client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).ToNot(HaveOccurred())
	Expect(k8sClient).ToNot(BeNil())
	log.Println("k8sClient:\n ", k8sClient)

	// Create the Resourcegroup resource
	result, _ := resoucegroupsresourcemanager.CheckExistence(context.Background(), resourceGroupName)
	if result.Response.StatusCode != 204 {
		_, _ = resoucegroupsresourcemanager.CreateGroup(context.Background(), resourceGroupName, resourcegroupLocation)
	}

	// Create the Eventhub namespace resource
	_, err = eventhubs.CreateNamespaceAndWait(context.Background(), resourceGroupName, eventhubNamespaceName, namespaceLocation)

	// Create the Eventhub resource
	_, err = eventhubs.CreateHub(context.Background(), resourceGroupName, eventhubNamespaceName, eventhubName, int32(7), int32(1), nil)

	// Create the Storage Account and Container
	_, err = storages.CreateStorage(context.Background(), resourceGroupName, storageAccountName, resourcegroupLocation, azurev1.StorageSku{
		Name: "Standard_LRS",
	}, "Storage", map[string]*string{}, "", nil)

	_, err = storages.CreateBlobContainer(context.Background(), resourceGroupName, storageAccountName, blobContainerName)

	tc := TestContext{
		Cfg:                   *cfg,
		ResourceGroupName:     resourceGroupName,
		ResourcegroupLocation: resourcegroupLocation,
		EventhubNamespaceName: eventhubNamespaceName,
		EventhubName:          eventhubName,
		NamespaceLocation:     namespaceLocation,
		StorageAccountName:    storageAccountName,
		BlobContainerName:     blobContainerName,
	}
	bytes, _ := toByteArray(&tc)

	return bytes
}, func(r []byte) {
	log.Println(fmt.Sprintf("per parallel process test setup"))
	_ = fromByteArray(r, &tc)

	//k8sClient = k8sManager.GetClient()
	k8sClient, err := client.New(&tc.Cfg, client.Options{Scheme: scheme.Scheme})
	Expect(err).ToNot(HaveOccurred())
	log.Println("k8sClient:\n ", k8sClient)

	tc.K8sClient = k8sClient

}, 120)

var _ = SynchronizedAfterSuite(func() {
	log.Println(fmt.Sprintf("per parallel process test teardown"))
}, func() {
	//clean up the resources created for test
	log.Println(fmt.Sprintf("common test teardown"))

	//clean up the resources created for test
	By("tearing down the test environment")

	// delete the resource group and contained resources
	_, _ = resoucegroupsresourcemanager.DeleteGroup(context.Background(), tc.ResourceGroupName)

	err := testEnv.Stop()
	Expect(err).ToNot(HaveOccurred())
}, 60)

func toByteArray(x interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	// Stand-in for a buffer connection
	enc := gob.NewEncoder(&buffer)
	// Will write to buffer.
	err := enc.Encode(x)
	if err != nil {
		return []byte{}, err
	}
	return buffer.Bytes(), err
}

func fromByteArray(r []byte, x interface{}) error {
	var buffer bytes.Buffer
	// Stand-in for a buffer connection
	dec := gob.NewDecoder(&buffer)
	// Will write to buffer.
	buffer.Write(r)
	return dec.Decode(x)
}
