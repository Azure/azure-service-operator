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
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-service-operator/pkg/secrets"
	k8sSecrets "github.com/Azure/azure-service-operator/pkg/secrets/kube"
	"k8s.io/client-go/rest"

	s "github.com/Azure/azure-sdk-for-go/profiles/latest/storage/mgmt/storage"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	resourcemanagersqldb "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqldb"
	resourcemanagersqlfailovergroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfailovergroup"
	resourcemanagersqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfirewallrule"
	resourcemanagersqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlserver"
	resourcemanagersqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagereventhub "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resourcemanagerkeyvaults "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	resourcemanagersqlmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/azuresql"
	resourcemanagereventhubmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/eventhubs"
	resourcemanagerkeyvaultsmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/keyvaults"
	resourcegroupsresourcemanagermock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/resourcegroups"
	resourcemanagerstoragesmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/storages"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	resourcemanagerstorages "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"

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
	k8sClient               client.Client
	secretClient            secrets.SecretClient
	resourceGroupName       string
	resourceGroupLocation   string
	eventhubNamespaceName   string
	eventhubName            string
	namespaceLocation       string
	storageAccountName      string
	blobContainerName       string
	resourceGroupManager    resourcegroupsresourcemanager.ResourceGroupManager
	eventHubManagers        resourcemanagereventhub.EventHubManagers
	eventhubClient          resourcemanagereventhub.EventHubManager
	storageManagers         resourcemanagerstorages.StorageManagers
	keyVaultManager         resourcemanagerkeyvaults.KeyVaultManager
	sqlServerManager        resourcemanagersqlserver.SqlServerManager
	sqlDbManager            resourcemanagersqldb.SqlDbManager
	sqlFirewallRuleManager  resourcemanagersqlfirewallrule.SqlFirewallRuleManager
	sqlFailoverGroupManager resourcemanagersqlfailovergroup.SqlFailoverGroupManager
	sqlUserManager          resourcemanagersqluser.SqlUserManager
	timeout                 time.Duration
	retry                   time.Duration
}

var tc testContext

func TestAPIs(t *testing.T) {
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
	containerAccessLevel := s.PublicAccessContainer

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

	err = azurev1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	err = azurev1alpha1.AddToScheme(scheme.Scheme)
	Expect(err).NotTo(HaveOccurred())

	// +kubebuilder:scaffold:scheme
	k8sManager, err = ctrl.NewManager(cfg, ctrl.Options{
		Scheme: scheme.Scheme,
	})
	Expect(err).ToNot(HaveOccurred())

	secretClient := k8sSecrets.New(k8sManager.GetClient())
	var resourceGroupManager resourcegroupsresourcemanager.ResourceGroupManager
	var eventHubManagers resourcemanagereventhub.EventHubManagers
	var storageManagers resourcemanagerstorages.StorageManagers
	var keyVaultManager resourcemanagerkeyvaults.KeyVaultManager
	var eventhubNamespaceClient resourcemanagereventhub.EventHubNamespaceManager
<<<<<<< HEAD
	var sqlServerManager resourcemanagersqlserver.SqlServerManager
	var sqlDbManager resourcemanagersqldb.SqlDbManager
	var sqlFirewallRuleManager resourcemanagersqlfirewallrule.SqlFirewallRuleManager
	var sqlFailoverGroupManager resourcemanagersqlfailovergroup.SqlFailoverGroupManager
	var sqlUserManager resourcemanagersqluser.SqlUserManager
=======
	var eventhubClient resourcemanagereventhub.EventHubManager
	var sqlServerManager resourcemanagersql.SqlServerManager
	var sqlDbManager resourcemanagersql.SqlDbManager
	var sqlFirewallRuleManager resourcemanagersql.SqlFirewallRuleManager
	var sqlFailoverGroupManager resourcemanagersql.SqlFailoverGroupManager
	var sqlUserManager resourcemanagersql.SqlUserManager
>>>>>>> f71bd8fbfa2023d28f950025ee953dca05d00057

	if os.Getenv("TEST_CONTROLLER_WITH_MOCKS") == "false" {
		resourceGroupManager = resourcegroupsresourcemanager.NewAzureResourceGroupManager()
		eventHubManagers = resourcemanagereventhub.AzureEventHubManagers
		storageManagers = resourcemanagerstorages.AzureStorageManagers
		keyVaultManager = resourcemanagerkeyvaults.AzureKeyVaultManager
		eventhubNamespaceClient = resourcemanagereventhub.NewEventHubNamespaceClient(ctrl.Log.WithName("controllers").WithName("EventhubNamespace"))
<<<<<<< HEAD
		sqlServerManager = resourcemanagersqlserver.NewAzureSqlServerManager(ctrl.Log.WithName("sqlservermanager").WithName("AzureSqlServer"))
		sqlDbManager = resourcemanagersqldb.NewAzureSqlDbManager(ctrl.Log.WithName("sqldbmanager").WithName("AzureSqlDb"))
		sqlFirewallRuleManager = resourcemanagersqlfirewallrule.NewAzureSqlFirewallRuleManager(ctrl.Log.WithName("sqlfirewallrulemanager").WithName("AzureSqlFirewallRule"))
		sqlFailoverGroupManager = resourcemanagersqlfailovergroup.NewAzureSqlFailoverGroupManager(ctrl.Log.WithName("sqlfailovergroupmanager").WithName("AzureSqlFailoverGroup"))
		sqlUserManager = resourcemanagersqluser.NewAzureSqlUserManager(ctrl.Log.WithName("sqlusermanager").WithName("AzureSqlUser"))
=======
		eventhubClient = resourcemanagereventhub.NewEventhubClient(secretClient, scheme.Scheme)
		sqlServerManager = resourcemanagersql.NewAzureSqlServerManager(ctrl.Log.WithName("sqlservermanager").WithName("AzureSqlServer"))
		sqlDbManager = resourcemanagersql.NewAzureSqlDbManager(ctrl.Log.WithName("sqldbmanager").WithName("AzureSqlDb"))
		sqlFirewallRuleManager = resourcemanagersql.NewAzureSqlFirewallRuleManager(ctrl.Log.WithName("sqlfirewallrulemanager").WithName("AzureSqlFirewallRule"))
		sqlFailoverGroupManager = resourcemanagersql.NewAzureSqlFailoverGroupManager(ctrl.Log.WithName("sqlfailovergroupmanager").WithName("AzureSqlFailoverGroup"))
		sqlUserManager = resourcemanagersql.NewAzureSqlUserManager(ctrl.Log.WithName("sqlusermanager").WithName("AzureSqlUser"))
>>>>>>> f71bd8fbfa2023d28f950025ee953dca05d00057
		timeout = time.Second * 900
	} else {
		resourceGroupManager = &resourcegroupsresourcemanagermock.MockResourceGroupManager{}
		eventHubManagers = resourcemanagereventhubmock.MockEventHubManagers
		storageManagers = resourcemanagerstoragesmock.MockStorageManagers
		keyVaultManager = &resourcemanagerkeyvaultsmock.MockKeyVaultManager{}
		eventhubNamespaceClient = resourcemanagereventhubmock.NewMockEventHubNamespaceClient()
		eventhubClient = resourcemanagereventhubmock.NewMockEventHubClient(secretClient, scheme.Scheme)
		sqlServerManager = resourcemanagersqlmock.NewMockSqlServerManager()
		sqlDbManager = resourcemanagersqlmock.NewMockSqlDbManager()
		sqlFirewallRuleManager = resourcemanagersqlmock.NewMockSqlFirewallRuleManager()
		sqlFailoverGroupManager = resourcemanagersqlmock.NewMockSqlFailoverGroupManager()
		sqlUserManager = resourcemanagersqlmock.NewMockAzureSqlUserManager()

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
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: eventhubClient,
			Telemetry: telemetry.InitializePrometheusDefault(
				ctrl.Log.WithName("controllers").WithName("EventHub"),
				"EventHub",
			),
			Recorder: k8sManager.GetEventRecorderFor("Eventhub-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&ResourceGroupReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: resourceGroupManager,
			Telemetry: telemetry.InitializePrometheusDefault(
				ctrl.Log.WithName("controllers").WithName("ResourceGroup"),
				"ResourceGroup",
			),
			Recorder: k8sManager.GetEventRecorderFor("ResourceGroup-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&EventhubNamespaceReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: eventhubNamespaceClient,
			Telemetry: telemetry.InitializePrometheusDefault(
				ctrl.Log.WithName("controllers").WithName("EventhubNamespace"),
				"EventhubNamespace",
			),
			Recorder: k8sManager.GetEventRecorderFor("EventhubNamespace-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&ConsumerGroupReconciler{
		Client:               k8sManager.GetClient(),
		Log:                  ctrl.Log.WithName("controllers").WithName("ConsumerGroup"),
		Recorder:             k8sManager.GetEventRecorderFor("ConsumerGroup-controller"),
		ConsumerGroupManager: eventHubManagers.ConsumerGroup,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&AzureDataLakeGen2FileSystemReconciler{
		Client:            k8sManager.GetClient(),
		Log:               ctrl.Log.WithName("controllers").WithName("AzureDataLakeGen2FileSystem"),
		Recorder:          k8sManager.GetEventRecorderFor("AzureDataLakeGen2FileSystem-controller"),
		FileSystemManager: storageManagers.FileSystem,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&AzureSqlServerReconciler{
		Client:                k8sManager.GetClient(),
		Log:                   ctrl.Log.WithName("controllers").WithName("AzureSqlServer"),
		Recorder:              k8sManager.GetEventRecorderFor("AzureSqlServer-controller"),
		Scheme:                scheme.Scheme,
		AzureSqlServerManager: sqlServerManager,
		SecretClient:          secretClient,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&AzureSqlDatabaseReconciler{
		Client:            k8sManager.GetClient(),
		Log:               ctrl.Log.WithName("controllers").WithName("AzureSqlDatabase"),
		Recorder:          k8sManager.GetEventRecorderFor("AzureSqlDatabase-controller"),
		Scheme:            scheme.Scheme,
		AzureSqlDbManager: sqlDbManager,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&AzureSqlFirewallRuleReconciler{
		Client: k8sManager.GetClient(),
		Telemetry: telemetry.InitializePrometheusDefault(
			ctrl.Log.WithName("controllers").WithName("AzureSQLFirewallRuleOperator"),
			"AzureSQLFirewallRuleOperator",
		),
		Recorder:                    k8sManager.GetEventRecorderFor("AzureSqlFirewallRule-controller"),
		Scheme:                      scheme.Scheme,
		AzureSqlFirewallRuleManager: sqlFirewallRuleManager,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&AzureSqlFailoverGroupReconciler{
		Client:                       k8sManager.GetClient(),
		Log:                          ctrl.Log.WithName("controllers").WithName("AzureSqlFailoverGroup"),
		Recorder:                     k8sManager.GetEventRecorderFor("AzureSqlFailoverGroup-controller"),
		Scheme:                       scheme.Scheme,
		AzureSqlFailoverGroupManager: sqlFailoverGroupManager,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&AzureSQLUserReconciler{
		Client:              k8sManager.GetClient(),
		Log:                 ctrl.Log.WithName("controllers").WithName("AzureSqlUser"),
		Recorder:            k8sManager.GetEventRecorderFor("AzureSqlUser-controller"),
		Scheme:              scheme.Scheme,
		AzureSqlUserManager: sqlUserManager,
		SecretClient:        secretClient,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&AzureSqlActionReconciler{
		Client:                k8sManager.GetClient(),
		Log:                   ctrl.Log.WithName("controllers").WithName("AzureSqlAction"),
		Recorder:              k8sManager.GetEventRecorderFor("AzureSqlAction-controller"),
		Scheme:                scheme.Scheme,
		AzureSqlServerManager: sqlServerManager,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&BlobContainerReconciler{
		Client:         k8sManager.GetClient(),
		Log:            ctrl.Log.WithName("controllers").WithName("BlobContainer"),
		Recorder:       k8sManager.GetEventRecorderFor("BlobContainer-controller"),
		Scheme:         scheme.Scheme,
		StorageManager: storageManagers.BlobContainer,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&AzureSqlFailoverGroupReconciler{
		Client:                       k8sManager.GetClient(),
		Log:                          ctrl.Log.WithName("controllers").WithName("AzureSqlFailoverGroup"),
		Recorder:                     k8sManager.GetEventRecorderFor("AzureSqlFailoverGroup-controller"),
		Scheme:                       scheme.Scheme,
		AzureSqlFailoverGroupManager: sqlFailoverGroupManager,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	err = (&AzureSqlFirewallRuleReconciler{
		Client: k8sManager.GetClient(),
		Telemetry: telemetry.InitializePrometheusDefault(
			ctrl.Log.WithName("controllers").WithName("AzureSQLFirewallRuleOperator"),
			"AzureSQLFirewallRuleOperator",
		),
		Recorder:                    k8sManager.GetEventRecorderFor("AzureSqlFirewall-controller"),
		Scheme:                      scheme.Scheme,
		AzureSqlFirewallRuleManager: sqlFirewallRuleManager,
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

	tc = testContext{
		k8sClient:               k8sClient,
		secretClient:            secretClient,
		resourceGroupName:       resourceGroupName,
		resourceGroupLocation:   resourcegroupLocation,
		eventhubNamespaceName:   eventhubNamespaceName,
		eventhubName:            eventhubName,
		namespaceLocation:       namespaceLocation,
		storageAccountName:      storageAccountName,
		blobContainerName:       blobContainerName,
		eventHubManagers:        eventHubManagers,
		eventhubClient:          eventhubClient,
		resourceGroupManager:    resourceGroupManager,
		sqlServerManager:        sqlServerManager,
		sqlDbManager:            sqlDbManager,
		sqlFirewallRuleManager:  sqlFirewallRuleManager,
		sqlFailoverGroupManager: sqlFailoverGroupManager,
		sqlUserManager:          sqlUserManager,
		storageManagers:         storageManagers,
		keyVaultManager:         keyVaultManager,
		timeout:                 timeout,
		retry:                   time.Second * 3,
	}

	Eventually(func() bool {
		namespace, _ := eventHubManagers.EventHubNamespace.GetNamespace(context.Background(), resourceGroupName, eventhubNamespaceName)
		return namespace.ProvisioningState != nil && *namespace.ProvisioningState == "Succeeded"
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// Create the Eventhub resource
	_, err = eventHubManagers.EventHub.CreateHub(context.Background(), resourceGroupName, eventhubNamespaceName, eventhubName, int32(7), int32(2), nil)
	Expect(err).ToNot(HaveOccurred())

	// Create the Storage Account and Container
	_, _ = storageManagers.Storage.CreateStorage(context.Background(), resourceGroupName, storageAccountName, resourcegroupLocation, azurev1alpha1.StorageSku{
		Name: "Standard_LRS",
	}, "Storage", map[string]*string{}, "", nil, nil)

	// Storage account needs to be in "Suceeded" state
	// for container create to succeed
	Eventually(func() s.ProvisioningState {
		result, _ := storageManagers.Storage.GetStorage(context.Background(), resourceGroupName, storageAccountName)
		return result.ProvisioningState
	}, tc.timeout, tc.retry,
	).Should(Equal(s.Succeeded))

	_, err = storageManagers.BlobContainer.CreateBlobContainer(context.Background(), resourceGroupName, storageAccountName, blobContainerName, containerAccessLevel)
	Expect(err).ToNot(HaveOccurred())

})

var _ = AfterSuite(func() {
	log.Println(fmt.Sprintf("Started common controller test teardown"))
	//clean up the resources created for test
	By("tearing down the test environment")

	// delete the resource group and contained resources
	Eventually(func() bool {
		_, err := tc.resourceGroupManager.DeleteGroup(context.Background(), tc.resourceGroupName)
		if err != nil {
			log.Println(err)
			log.Println()
			if strings.Contains(err.Error(), "asynchronous operation has not completed") {
				return true
			}
			return false
		}
		return true
	}, tc.timeout, tc.retry,
	).Should(BeTrue())
	err := testEnv.Stop()
	Expect(err).ToNot(HaveOccurred())
	log.Println(fmt.Sprintf("Finished common controller test teardown"))
})
