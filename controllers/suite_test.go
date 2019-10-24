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
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	"k8s.io/client-go/rest"

	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagereventhub "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resourcemanagerkeyvaults "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	resourcemanagereventhubmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/eventhubs"
	resourcemanagerkeyvaultsmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/keyvaults"
	resourcegroupsresourcemanagermock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/resourcegroups"
	resourcemanagerstoragesmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/storages"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	resourcemanagerstorages "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"

	resourcemanagersql "github.com/Azure/azure-service-operator/pkg/resourcemanager/sqlclient"

	resourcemanagersqlmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/sqlclient"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
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

type testContext struct {
	k8sClient             client.Client
	resourceGroupName     string
	resourceGroupLocation string
	eventhubNamespaceName string
	eventhubName          string
	namespaceLocation     string
	storageAccountName    string
	blobContainerName     string
	resourceGroupManager  resourcegroupsresourcemanager.ResourceGroupManager
	eventHubManagers      resourcemanagereventhub.EventHubManagers
	storageManagers       resourcemanagerstorages.StorageManagers
	keyVaultManager       resourcemanagerkeyvaults.KeyVaultManager
	timeout               time.Duration
}

var tc testContext

func TestAPIs(t *testing.T) {
	t.Parallel()
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{envtest.NewlineReporter{}})
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))
	log.Println(fmt.Sprintf("Starting common controller test setup"))

	err := resourcemanagerconfig.ParseEnvironment()
	if err != nil {
		Fail(err.Error())
	}

	resourceGroupName := "t-rg-dev-controller-" + helpers.RandomString(10)
	resourcegroupLocation := resourcemanagerconfig.DefaultLocation()

	eventhubNamespaceName := "t-ns-dev-eh-ns-" + helpers.RandomString(10)
	eventhubName := "t-eh-dev-sample-" + helpers.RandomString(10)
	namespaceLocation := resourcemanagerconfig.DefaultLocation()

	storageAccountName := "tsadeveh" + helpers.RandomString(10)
	blobContainerName := "t-bc-dev-eh-" + helpers.RandomString(10)

	var timeout time.Duration

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{filepath.Join("..", "config", "crd", "bases")},
	}

	var cfg *rest.Config
	if os.Getenv("TEST_USE_EXISTING_CLUSTER") == "true" {
		t := true
		testEnv = &envtest.Environment{
			UseExistingCluster: &t,
		}
		cfg, err = ctrl.GetConfig()
		Expect(err).ToNot(HaveOccurred())
	} else {
		testEnv = &envtest.Environment{
			CRDDirectoryPaths: []string{filepath.Join("..", "config", "crd", "bases")},
		}
		cfg, err = testEnv.Start()
		Expect(err).ToNot(HaveOccurred())
	}

	Expect(cfg).ToNot(BeNil())

	err = azurev1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	var k8sManager ctrl.Manager

	// +kubebuilder:scaffold:scheme
	k8sManager, err = ctrl.NewManager(cfg, ctrl.Options{
		Scheme: scheme.Scheme,
	})
	Expect(err).ToNot(HaveOccurred())

	var resourceGroupManager resourcegroupsresourcemanager.ResourceGroupManager
	var eventHubManagers resourcemanagereventhub.EventHubManagers
	var storageManagers resourcemanagerstorages.StorageManagers
	var keyVaultManager resourcemanagerkeyvaults.KeyVaultManager
	var resourceClient resourcemanagersql.ResourceClient

	if os.Getenv("TEST_CONTROLLER_WITH_MOCKS") == "false" {
		resourceGroupManager = resourcegroupsresourcemanager.AzureResourceGroupManager
		eventHubManagers = resourcemanagereventhub.AzureEventHubManagers
		storageManagers = resourcemanagerstorages.AzureStorageManagers
		keyVaultManager = resourcemanagerkeyvaults.AzureKeyVaultManager
		resourceClient = &resourcemanagersql.GoSDKClient{}
		timeout = time.Second * 320
	} else {
		resourceGroupManager = &resourcegroupsresourcemanagermock.MockResourceGroupManager{}
		eventHubManagers = resourcemanagereventhubmock.MockEventHubManagers
		storageManagers = resourcemanagerstoragesmock.MockStorageManagers
		keyVaultManager = &resourcemanagerkeyvaultsmock.MockKeyVaultManager{}
		resourceClient = &resourcemanagersqlmock.MockGoSDKClient{}
		timeout = time.Second * 60
	}

	err = (&KeyVaultReconciler{
		Client:          k8sManager.GetClient(),
		Log:             ctrl.Log.WithName("controllers").WithName("KeyVault"),
		Recorder:        k8sManager.GetEventRecorderFor("KeyVault-controller"),
		KeyVaultManager: keyVaultManager,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&EventhubReconciler{
		Client:          k8sManager.GetClient(),
		Log:             ctrl.Log.WithName("controllers").WithName("EventHub"),
		Recorder:        k8sManager.GetEventRecorderFor("Eventhub-controller"),
		Scheme:          scheme.Scheme,
		EventHubManager: eventHubManagers.EventHub,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&ResourceGroupReconciler{
		Client:               k8sManager.GetClient(),
		Log:                  ctrl.Log.WithName("controllers").WithName("ResourceGroup"),
		Recorder:             k8sManager.GetEventRecorderFor("ResourceGroup-controller"),
		ResourceGroupManager: resourceGroupManager,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&EventhubNamespaceReconciler{
		Client:                   k8sManager.GetClient(),
		Log:                      ctrl.Log.WithName("controllers").WithName("EventhubNamespace"),
		Recorder:                 k8sManager.GetEventRecorderFor("EventhubNamespace-controller"),
		EventHubNamespaceManager: eventHubManagers.EventHubNamespace,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&ConsumerGroupReconciler{
		Client:               k8sManager.GetClient(),
		Log:                  ctrl.Log.WithName("controllers").WithName("ConsumerGroup"),
		Recorder:             k8sManager.GetEventRecorderFor("ConsumerGroup-controller"),
		ConsumerGroupManager: eventHubManagers.ConsumerGroup,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&AzureSqlServerReconciler{
		Client:         k8sManager.GetClient(),
		Log:            ctrl.Log.WithName("controllers").WithName("AzureSqlServer"),
		Recorder:       k8sManager.GetEventRecorderFor("AzureSqlServer-controller"),
		Scheme:         scheme.Scheme,
		ResourceClient: resourceClient,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&AzureSqlDatabaseReconciler{
		Client:         k8sManager.GetClient(),
		Log:            ctrl.Log.WithName("controllers").WithName("AzureSqlDatabase"),
		Recorder:       k8sManager.GetEventRecorderFor("AzureSqlDatabase-controller"),
		Scheme:         scheme.Scheme,
		ResourceClient: resourceClient,
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

	// Create the ResourceGroup resource
	result, _ := resourceGroupManager.CheckExistence(context.Background(), resourceGroupName)
	if result.Response.StatusCode != 204 {
		_, _ = resourceGroupManager.CreateGroup(context.Background(), resourceGroupName, resourcegroupLocation)
	}

	eventHubNSManager := eventHubManagers.EventHubNamespace
	// Create the Eventhub namespace resource
	_, err = eventHubNSManager.CreateNamespaceAndWait(context.Background(), resourceGroupName, eventhubNamespaceName, namespaceLocation)
	Expect(err).ToNot(HaveOccurred())

	Eventually(func() bool {
		namespace, _ := eventHubManagers.EventHubNamespace.GetNamespace(context.Background(), resourceGroupName, eventhubNamespaceName)
		return namespace.ProvisioningState != nil && *namespace.ProvisioningState == "Succeeded"
	}, 60,
	).Should(BeTrue())

	// Create the Eventhub resource
	_, err = eventHubManagers.EventHub.CreateHub(context.Background(), resourceGroupName, eventhubNamespaceName, eventhubName, int32(7), int32(2), nil)
	Expect(err).ToNot(HaveOccurred())

	// Create the Storage Account and Container
	_, err = storageManagers.Storage.CreateStorage(context.Background(), resourceGroupName, storageAccountName, resourcegroupLocation, azurev1alpha1.StorageSku{
		Name: "Standard_LRS",
	}, "Storage", map[string]*string{}, "", nil)
	Expect(err).ToNot(HaveOccurred())

	_, err = storageManagers.BlobContainer.CreateBlobContainer(context.Background(), resourceGroupName, storageAccountName, blobContainerName)
	Expect(err).ToNot(HaveOccurred())

	tc = testContext{
		k8sClient:             k8sClient,
		resourceGroupName:     resourceGroupName,
		resourceGroupLocation: resourcegroupLocation,
		eventhubNamespaceName: eventhubNamespaceName,
		eventhubName:          eventhubName,
		namespaceLocation:     namespaceLocation,
		storageAccountName:    storageAccountName,
		blobContainerName:     blobContainerName,
		eventHubManagers:      eventHubManagers,
		resourceGroupManager:  resourceGroupManager,
		storageManagers:       storageManagers,
		keyVaultManager:       keyVaultManager,
		timeout:               timeout,
	}
})

var _ = AfterSuite(func() {
	log.Println(fmt.Sprintf("Started common controller test teardown"))
	//clean up the resources created for test
	By("tearing down the test environment")

	// delete the resource group and contained resources
	_, _ = tc.resourceGroupManager.DeleteGroup(context.Background(), tc.resourceGroupName)

	err := testEnv.Stop()
	Expect(err).ToNot(HaveOccurred())
	log.Println(fmt.Sprintf("Finished common controller test teardown"))
})
