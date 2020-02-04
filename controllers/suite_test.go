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

	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	k8sSecrets "github.com/Azure/azure-service-operator/pkg/secrets/kube"
	"k8s.io/client-go/rest"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	resourcemanagerappinsights "github.com/Azure/azure-service-operator/pkg/resourcemanager/appinsights"
	resourcemanagersqldb "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqldb"
	resourcemanagersqlfailovergroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfailovergroup"
	resourcemanagersqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfirewallrule"
	resourcemanagersqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlserver"
	resourcemanagersqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagereventhub "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resourcemanagerkeyvaults "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	resourcemanagerappinsightsmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/appinsights"
	resourcemanagersqlmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/azuresql"
	resourcemanagereventhubmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/eventhubs"
	resourcemanagerkeyvaultsmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/keyvaults"
	resourcemanagerpsqlmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/psql"
	resourcegroupsresourcemanagermock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/resourcegroups"
	resourcemanagerstoragesmock "github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/storages"
	resourcemanagerpsqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/database"
	resourcemanagerpsqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/firewallrule"
	resourcemanagerpsqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/server"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	resourcemanagerstorages "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var testEnv *envtest.Environment

var tc TestContext

func setup() error {
	log.Println(fmt.Sprintf("Starting common controller test setup"))
	defer PanicRecover()

	err := resourcemanagerconfig.ParseEnvironment()
	if err != nil {
		return err
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
		if err != nil {
			return err
		}
	} else {
		testEnv = &envtest.Environment{
			CRDDirectoryPaths: []string{filepath.Join("..", "config", "crd", "bases")},
		}
		cfg, err = testEnv.Start()
		if err != nil {
			return err
		}
	}
	if cfg == nil {
		return fmt.Errorf("rest config nil")
	}

	err = azurev1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	var k8sManager ctrl.Manager

	// +kubebuilder:scaffold:scheme
	k8sManager, err = ctrl.NewManager(cfg, ctrl.Options{
		Scheme: scheme.Scheme,
	})
	if err != nil {
		return err
	}

	secretClient := k8sSecrets.New(k8sManager.GetClient())

	var appInsightsManager resourcemanagerappinsights.ApplicationInsightsManager
	var resourceGroupManager resourcegroupsresourcemanager.ResourceGroupManager
	var eventHubManagers resourcemanagereventhub.EventHubManagers
	var storageManagers resourcemanagerstorages.StorageManagers
	var keyVaultManager resourcemanagerkeyvaults.KeyVaultManager
	var eventhubNamespaceClient resourcemanagereventhub.EventHubNamespaceManager
	var sqlServerManager resourcemanagersqlserver.SqlServerManager
	var sqlDbManager resourcemanagersqldb.SqlDbManager
	var sqlFirewallRuleManager resourcemanagersqlfirewallrule.SqlFirewallRuleManager
	var sqlFailoverGroupManager resourcemanagersqlfailovergroup.SqlFailoverGroupManager
	var sqlUserManager resourcemanagersqluser.SqlUserManager
	var eventhubClient resourcemanagereventhub.EventHubManager
	var psqlServerManager resourcemanagerpsqlserver.PostgreSQLServerManager
	var psqlDatabaseManager resourcemanagerpsqldatabase.PostgreSQLDatabaseManager
	var psqlFirewallRuleManager resourcemanagerpsqlfirewallrule.PostgreSQLFirewallRuleManager
	var consumerGroupClient resourcemanagereventhub.ConsumerGroupManager

	if os.Getenv("TEST_CONTROLLER_WITH_MOCKS") == "false" {
		appInsightsManager = resourcemanagerappinsights.NewManager(ctrl.Log.WithName("appinsightsmanager").WithName("AppInsights"))
		resourceGroupManager = resourcegroupsresourcemanager.NewAzureResourceGroupManager()
		eventHubManagers = resourcemanagereventhub.AzureEventHubManagers
		storageManagers = resourcemanagerstorages.AzureStorageManagers
		keyVaultManager = resourcemanagerkeyvaults.NewAzureKeyVaultManager(ctrl.Log.WithName("controllers").WithName("KeyVault"), k8sManager.GetScheme())
		eventhubClient = resourcemanagereventhub.NewEventhubClient(secretClient, scheme.Scheme)
		psqlServerManager = resourcemanagerpsqlserver.NewPSQLServerClient(ctrl.Log.WithName("psqlservermanager").WithName("PostgreSQLServer"), secretClient, k8sManager.GetScheme())
		psqlDatabaseManager = resourcemanagerpsqldatabase.NewPSQLDatabaseClient(ctrl.Log.WithName("psqldatabasemanager").WithName("PostgreSQLDatabase"))
		psqlFirewallRuleManager = resourcemanagerpsqlfirewallrule.NewPSQLFirewallRuleClient(ctrl.Log.WithName("psqlfirewallrulemanager").WithName("PostgreSQLFirewallRule"))
		eventhubNamespaceClient = resourcemanagereventhub.NewEventHubNamespaceClient(ctrl.Log.WithName("controllers").WithName("EventhubNamespace"))
		sqlServerManager = resourcemanagersqlserver.NewAzureSqlServerManager(ctrl.Log.WithName("sqlservermanager").WithName("AzureSqlServer"))
		sqlDbManager = resourcemanagersqldb.NewAzureSqlDbManager(ctrl.Log.WithName("sqldbmanager").WithName("AzureSqlDb"))
		sqlFirewallRuleManager = resourcemanagersqlfirewallrule.NewAzureSqlFirewallRuleManager(ctrl.Log.WithName("sqlfirewallrulemanager").WithName("AzureSqlFirewallRule"))
		sqlFailoverGroupManager = resourcemanagersqlfailovergroup.NewAzureSqlFailoverGroupManager(ctrl.Log.WithName("sqlfailovergroupmanager").WithName("AzureSqlFailoverGroup"))
		consumerGroupClient = resourcemanagereventhub.NewConsumerGroupClient(ctrl.Log.WithName("controllers").WithName("ConsumerGroup"))
		sqlUserManager = resourcemanagersqluser.NewAzureSqlUserManager(
			ctrl.Log.WithName("sqlusermanager").WithName("AzureSqlUser"),
			secretClient,
			scheme.Scheme,
		)

		timeout = time.Second * 720
	} else {
		appInsightsManager = resourcemanagerappinsightsmock.NewMockAppInsightsManager(scheme.Scheme)
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
		psqlServerManager = resourcemanagerpsqlmock.NewMockPSQLServerClient(ctrl.Log.WithName("psqlservermanager").WithName("PostgreSQLServer"), secretClient, k8sManager.GetScheme())
		psqlDatabaseManager = resourcemanagerpsqlmock.NewMockPostgreSqlDbManager(ctrl.Log.WithName("psqldatabasemanager").WithName("PostgreSQLDatabase"))
		psqlFirewallRuleManager = resourcemanagerpsqlmock.NewMockPostgreSqlFirewallRuleManager(ctrl.Log.WithName("psqlfirewallrulemanager").WithName("PostgreSQLFirewallRule"))
		consumerGroupClient = resourcemanagereventhubmock.NewMockConsumerGroupClient()
		timeout = time.Second * 60
	}

	err = (&KeyVaultReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: keyVaultManager,
			Telemetry: telemetry.InitializePrometheusDefault(
				ctrl.Log.WithName("controllers").WithName("KeyVault"),
				"KeyVault",
			),
			Recorder: k8sManager.GetEventRecorderFor("KeyVault-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AppInsightsReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: appInsightsManager,
			Telemetry: telemetry.InitializePrometheusDefault(
				ctrl.Log.WithName("controllers").WithName("AppInsights"),
				"AppInsights",
			),
			Recorder: k8sManager.GetEventRecorderFor("AppInsights-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

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
	if err != nil {
		return err
	}

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
	if err != nil {
		return err
	}

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
	if err != nil {
		return err
	}

	err = (&ConsumerGroupReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: consumerGroupClient,
			Telemetry: telemetry.InitializePrometheusDefault(
				ctrl.Log.WithName("controllers").WithName("ConsumerGroup"),
				"ConsumerGroup",
			),
			Recorder: k8sManager.GetEventRecorderFor("ConsumerGroup-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureDataLakeGen2FileSystemReconciler{
		Client:            k8sManager.GetClient(),
		Log:               ctrl.Log.WithName("controllers").WithName("AzureDataLakeGen2FileSystem"),
		Recorder:          k8sManager.GetEventRecorderFor("AzureDataLakeGen2FileSystem-controller"),
		FileSystemManager: storageManagers.FileSystem,
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureSqlServerReconciler{
		Client:                k8sManager.GetClient(),
		Log:                   ctrl.Log.WithName("controllers").WithName("AzureSqlServer"),
		Recorder:              k8sManager.GetEventRecorderFor("AzureSqlServer-controller"),
		Scheme:                scheme.Scheme,
		AzureSqlServerManager: sqlServerManager,
		SecretClient:          secretClient,
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureSqlDatabaseReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: sqlDbManager,
			Telemetry: telemetry.InitializePrometheusDefault(
				ctrl.Log.WithName("controllers").WithName("AzureSqlDb"),
				"AzureSqlDb",
			),
			Recorder: k8sManager.GetEventRecorderFor("AzureSqlDb-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

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
	if err != nil {
		return err
	}

	err = (&AzureSqlFailoverGroupReconciler{
		Client:                       k8sManager.GetClient(),
		Log:                          ctrl.Log.WithName("controllers").WithName("AzureSqlFailoverGroup"),
		Recorder:                     k8sManager.GetEventRecorderFor("AzureSqlFailoverGroup-controller"),
		Scheme:                       scheme.Scheme,
		AzureSqlFailoverGroupManager: sqlFailoverGroupManager,
		SecretClient:                 secretClient,
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureSQLUserReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: sqlUserManager,
			Telemetry: telemetry.InitializePrometheusDefault(
				ctrl.Log.WithName("controllers").WithName("AzureSqlUser"),
				"AzureSqlUser",
			),
			Recorder: k8sManager.GetEventRecorderFor("AzureSqlUser-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureSqlActionReconciler{
		Client:                k8sManager.GetClient(),
		Log:                   ctrl.Log.WithName("controllers").WithName("AzureSqlAction"),
		Recorder:              k8sManager.GetEventRecorderFor("AzureSqlAction-controller"),
		Scheme:                scheme.Scheme,
		AzureSqlServerManager: sqlServerManager,
		SecretClient:          secretClient,
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&BlobContainerReconciler{
		Client:         k8sManager.GetClient(),
		Log:            ctrl.Log.WithName("controllers").WithName("BlobContainer"),
		Recorder:       k8sManager.GetEventRecorderFor("BlobContainer-controller"),
		Scheme:         scheme.Scheme,
		StorageManager: storageManagers.BlobContainer,
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&PostgreSQLServerReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: psqlServerManager,
			Telemetry: telemetry.InitializePrometheusDefault(
				ctrl.Log.WithName("controllers").WithName("PostgreSQLServer"),
				"PostgreSQLServer",
			),
			Recorder: k8sManager.GetEventRecorderFor("PostgreSQLServer-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&PostgreSQLDatabaseReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: psqlDatabaseManager,
			Telemetry: telemetry.InitializePrometheusDefault(
				ctrl.Log.WithName("controllers").WithName("PostgreSQLDatabaser"),
				"PostgreSQLDatabase",
			),
			Recorder: k8sManager.GetEventRecorderFor("PostgreSQLDatabase-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&PostgreSQLFirewallRuleReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: psqlFirewallRuleManager,
			Telemetry: telemetry.InitializePrometheusDefault(
				ctrl.Log.WithName("controllers").WithName("PostgreSQLFirewallRule"),
				"PostgreSQLFirewallRule",
			),
			Recorder: k8sManager.GetEventRecorderFor("PostgreSQLFirewallRule-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	go func() {
		err = k8sManager.Start(ctrl.SetupSignalHandler())
		if err != nil {
			log.Fatal(err)
		}
	}()

	k8sClient, err := client.New(cfg, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		return err
	}

	log.Println("Creating RG:", resourceGroupName)
	// Create the ResourceGroup resource
	result, _ := resourceGroupManager.CheckExistence(context.Background(), resourceGroupName)
	if result.Response.StatusCode != 204 {
		_, _ = resourceGroupManager.CreateGroup(context.Background(), resourceGroupName, resourcegroupLocation)
	}

	log.Println("Creating EHNS:", eventhubNamespaceName)
	eventHubNSManager := eventHubManagers.EventHubNamespace

	// Create the Eventhub namespace resource
	_, err = eventHubNSManager.CreateNamespaceAndWait(context.Background(), resourceGroupName, eventhubNamespaceName, namespaceLocation)
	if err != nil {
		return err
	}

	tc = TestContext{
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
		timeoutFast:             time.Minute * 3,
		retry:                   time.Second * 3,
		consumerGroupClient:     consumerGroupClient,
	}

	var pstate *string
	finish := time.Now().Add(tc.timeout)
	for {
		if finish.Before(time.Now()) {
			return fmt.Errorf("time out waiting for eventhub namespace")
		}

		namespace, _ := eventHubManagers.EventHubNamespace.GetNamespace(context.Background(), resourceGroupName, eventhubNamespaceName)
		pstate = namespace.ProvisioningState
		if pstate != nil && *pstate == "Succeeded" {
			break
		}
		time.Sleep(tc.retry)
	}

	log.Println("Creating EH:", eventhubName)
	// Create the Eventhub resource
	_, err = eventHubManagers.EventHub.CreateHub(context.Background(), resourceGroupName, eventhubNamespaceName, eventhubName, int32(7), int32(2), nil)
	if err != nil {
		return err
	}

	log.Println("Creating SA:", storageAccountName)
	// Create the Storage Account and Container
	_, _ = storageManagers.Storage.CreateStorage(context.Background(), resourceGroupName, storageAccountName, resourcegroupLocation, azurev1alpha1.StorageSku{
		Name: "Standard_LRS",
	}, "Storage", map[string]*string{}, "", nil, nil)

	// Storage account needs to be in "Suceeded" state
	// for container create to succeed
	finish = time.Now().Add(tc.timeout)
	for {

		if finish.Before(time.Now()) {
			return fmt.Errorf("time out waiting for storage account")
		}

		result, _ := storageManagers.Storage.GetStorage(context.Background(), resourceGroupName, storageAccountName)
		if result.ProvisioningState == s.Succeeded {
			break
		}
		time.Sleep(tc.retry)
	}

	_, err = storageManagers.BlobContainer.CreateBlobContainer(context.Background(), resourceGroupName, storageAccountName, blobContainerName, containerAccessLevel)
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("finished common controller test setup"))

	return nil
}

func teardown() error {
	log.Println(fmt.Sprintf("Started common controller test teardown"))

	finish := time.Now().Add(tc.timeout)
	for {

		if finish.Before(time.Now()) {
			return fmt.Errorf("time out waiting for rg to be gone")
		}

		_, err := tc.resourceGroupManager.DeleteGroup(context.Background(), tc.resourceGroupName)
		if err != nil {
			if strings.Contains(err.Error(), "asynchronous operation has not completed") {
				break
			}
		} else {
			break
		}
		time.Sleep(tc.retry)
	}

	err := testEnv.Stop()
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("Finished common controller test teardown"))
	return nil
}

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

func PanicRecover() {
	if err := recover(); err != nil {
		fmt.Println("caught panic in test:")
		fmt.Println(err)
		// fmt.Println("attempt to tear down...")
		// err = teardown()
		// if err != nil {
		// 	log.Println(fmt.Sprintf("could not tear down environment: %v\n", err))
		// }
	}
}
