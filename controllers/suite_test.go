// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"testing"
	"time"

	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	k8sSecrets "github.com/Azure/azure-service-operator/pkg/secrets/kube"
	"k8s.io/client-go/rest"

	resourcemanagerapimgmt "github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimgmt"
	resourcemanagerappinsights "github.com/Azure/azure-service-operator/pkg/resourcemanager/appinsights"
	resourcemanagersqlaction "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlaction"
	resourcemanagersqldb "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqldb"
	resourcemanagersqlfailovergroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfailovergroup"
	resourcemanagersqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfirewallrule"
	resourcemanagersqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlserver"
	resourcemanagersqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
	resourcemanagersqlvnetrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlvnetrule"
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagereventhub "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resourcemanagerkeyvaults "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	resourcemanagerpsqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/database"
	resourcemanagerpsqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/firewallrule"
	resourcemanagerpsqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/server"
	resourcemanagerrediscaches "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	resourcemanagerstorages "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"
	resourcemanagerblobcontainer "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/blobcontainer"
	resourcemanagerstorageaccount "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"
	resourcemanagervnet "github.com/Azure/azure-service-operator/pkg/resourcemanager/vnet"
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

	err := resourcemanagerconfig.ParseEnvironment()
	if err != nil {
		return err
	}

	resourceGroupName := GenerateTestResourceName("rg-prime")
	resourcegroupLocation := resourcemanagerconfig.DefaultLocation()

	eventhubNamespaceName := GenerateTestResourceName("evns-prime")
	eventhubName := GenerateTestResourceName("ev-prime")
	namespaceLocation := resourcemanagerconfig.DefaultLocation()

	storageAccountName := GenerateAlphaNumTestResourceName("saprime")
	blobContainerName := GenerateTestResourceName("blob-prime")
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
	var apiMgmtManager resourcemanagerapimgmt.APIManager
	var resourceGroupManager resourcegroupsresourcemanager.ResourceGroupManager
	var eventHubManagers resourcemanagereventhub.EventHubManagers
	var storageManagers resourcemanagerstorages.StorageManagers
	var eventhubNamespaceClient resourcemanagereventhub.EventHubNamespaceManager
	var sqlServerManager resourcemanagersqlserver.SqlServerManager
	var sqlDbManager resourcemanagersqldb.SqlDbManager
	var sqlFirewallRuleManager resourcemanagersqlfirewallrule.SqlFirewallRuleManager
	var sqlFailoverGroupManager resourcemanagersqlfailovergroup.SqlFailoverGroupManager
	var sqlUserManager resourcemanagersqluser.SqlUserManager
	var sqlActionManager resourcemanagersqlaction.SqlActionManager
	var eventhubClient resourcemanagereventhub.EventHubManager
	var psqlServerManager resourcemanagerpsqlserver.PostgreSQLServerManager
	var psqlDatabaseManager resourcemanagerpsqldatabase.PostgreSQLDatabaseManager
	var psqlFirewallRuleManager resourcemanagerpsqlfirewallrule.PostgreSQLFirewallRuleManager
	var consumerGroupClient resourcemanagereventhub.ConsumerGroupManager
	var sqlVNetRuleManager resourcemanagersqlvnetrule.SqlVNetRuleManager

	appInsightsManager = resourcemanagerappinsights.NewManager(
		secretClient,
		scheme.Scheme,
	)
	apiMgmtManager = resourcemanagerapimgmt.NewManager()
	resourceGroupManager = resourcegroupsresourcemanager.NewAzureResourceGroupManager()
	eventHubManagers = resourcemanagereventhub.AzureEventHubManagers
	storageManagers = resourcemanagerstorages.AzureStorageManagers
	storageAccountManager := resourcemanagerstorageaccount.New()
	blobContainerManager := resourcemanagerblobcontainer.New()
	keyVaultManager := resourcemanagerkeyvaults.NewAzureKeyVaultManager(k8sManager.GetScheme())
	keyVaultKeyManager := &resourcemanagerkeyvaults.KeyvaultKeyClient{
		KeyvaultClient: keyVaultManager,
	}

	virtualNetworkManager := resourcemanagervnet.NewAzureVNetManager()

	eventhubClient = resourcemanagereventhub.NewEventhubClient(secretClient, scheme.Scheme)
	psqlServerManager = resourcemanagerpsqlserver.NewPSQLServerClient(secretClient, k8sManager.GetScheme())
	psqlDatabaseManager = resourcemanagerpsqldatabase.NewPSQLDatabaseClient()
	psqlFirewallRuleManager = resourcemanagerpsqlfirewallrule.NewPSQLFirewallRuleClient()
	eventhubNamespaceClient = resourcemanagereventhub.NewEventHubNamespaceClient()

	sqlServerManager = resourcemanagersqlserver.NewAzureSqlServerManager(
		secretClient,
		scheme.Scheme,
	)
	redisCacheManager := resourcemanagerrediscaches.NewAzureRedisCacheManager(
		ctrl.Log.WithName("rediscachemanager").WithName("RedisCache"),
		secretClient,
		scheme.Scheme,
	)
	sqlDbManager = resourcemanagersqldb.NewAzureSqlDbManager()
	sqlFirewallRuleManager = resourcemanagersqlfirewallrule.NewAzureSqlFirewallRuleManager()
	sqlVNetRuleManager = resourcemanagersqlvnetrule.NewAzureSqlVNetRuleManager()
	sqlFailoverGroupManager = resourcemanagersqlfailovergroup.NewAzureSqlFailoverGroupManager(
		secretClient,
		scheme.Scheme,
	)
	consumerGroupClient = resourcemanagereventhub.NewConsumerGroupClient()
	sqlUserManager = resourcemanagersqluser.NewAzureSqlUserManager(
		secretClient,
		scheme.Scheme,
	)
	sqlActionManager = resourcemanagersqlaction.NewAzureSqlActionManager(secretClient, scheme.Scheme)

	timeout = time.Second * 720

	err = (&KeyVaultReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: keyVaultManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"KeyVault",
				ctrl.Log.WithName("controllers").WithName("KeyVault"),
			),
			Recorder: k8sManager.GetEventRecorderFor("KeyVault-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&KeyVaultKeyReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: keyVaultKeyManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"KeyVaultKey",
				ctrl.Log.WithName("controllers").WithName("KeyVaultKey"),
			),
			Recorder: k8sManager.GetEventRecorderFor("KeyVaultKey-controller"),
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
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AppInsights",
				ctrl.Log.WithName("controllers").WithName("AppInsights"),
			),
			Recorder: k8sManager.GetEventRecorderFor("AppInsights-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&APIMAPIReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: apiMgmtManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"ApiMgmt",
				ctrl.Log.WithName("controllers").WithName("ApiMgmt"),
			),
			Recorder: k8sManager.GetEventRecorderFor("ApiMgmt-controller"),
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
			Telemetry: telemetry.InitializeTelemetryDefault(
				"EventHub",
				ctrl.Log.WithName("controllers").WithName("EventHub"),
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
			Telemetry: telemetry.InitializeTelemetryDefault(
				"ResourceGroup",
				ctrl.Log.WithName("controllers").WithName("ResourceGroup"),
			),
			Recorder: k8sManager.GetEventRecorderFor("ResourceGroup-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&RedisCacheReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: redisCacheManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"RedisCache",
				ctrl.Log.WithName("controllers").WithName("RedisCache"),
			),
			Recorder: k8sManager.GetEventRecorderFor("RedisCache-controller"),
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
			Telemetry: telemetry.InitializeTelemetryDefault(
				"EventhubNamespace",
				ctrl.Log.WithName("controllers").WithName("EventhubNamespace"),
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
			Telemetry: telemetry.InitializeTelemetryDefault(
				"ConsumerGroup",
				ctrl.Log.WithName("controllers").WithName("ConsumerGroup"),
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
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: sqlServerManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSqlServer",
				ctrl.Log.WithName("controllers").WithName("AzureSqlServer"),
			),
			Recorder: k8sManager.GetEventRecorderFor("AzureSqlServer-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureSqlDatabaseReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: sqlDbManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSqlDb",
				ctrl.Log.WithName("controllers").WithName("AzureSqlDb"),
			),
			Recorder: k8sManager.GetEventRecorderFor("AzureSqlDb-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureSqlFirewallRuleReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: sqlFirewallRuleManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSQLFirewallRuleOperator",
				ctrl.Log.WithName("controllers").WithName("AzureSQLFirewallRuleOperator"),
			),
			Recorder: k8sManager.GetEventRecorderFor("AzureSqlFirewallRule-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureSQLVNetRuleReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: sqlVNetRuleManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSQLVNetRuleOperator",
				ctrl.Log.WithName("controllers").WithName("AzureSQLVNetRuleOperator"),
			),
			Recorder: k8sManager.GetEventRecorderFor("AzureSqlVNetRule-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureSqlFailoverGroupReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: sqlFailoverGroupManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSqlFailoverGroup",
				ctrl.Log.WithName("controllers").WithName("AzureSqlFailoverGroup"),
			),
			Recorder: k8sManager.GetEventRecorderFor("AzureSqlFailoverGroup-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureSQLUserReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: sqlUserManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSqlUser",
				ctrl.Log.WithName("controllers").WithName("AzureSqlUser"),
			),
			Recorder: k8sManager.GetEventRecorderFor("AzureSqlUser-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&VirtualNetworkReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: virtualNetworkManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"VirtualNetwork",
				ctrl.Log.WithName("controllers").WithName("VirtualNetwork"),
			),
			Recorder: k8sManager.GetEventRecorderFor("VirtualNetwork-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureSqlActionReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: sqlActionManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSqlAction",
				ctrl.Log.WithName("controllers").WithName("AzureSqlAction"),
			),
			Recorder: k8sManager.GetEventRecorderFor("AzureSqlAction-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&BlobContainerReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: blobContainerManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"BlobContainer",
				ctrl.Log.WithName("controllers").WithName("BlobContainer"),
			),
			Recorder: k8sManager.GetEventRecorderFor("BlobContainer-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&PostgreSQLServerReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: psqlServerManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"PostgreSQLServer",
				ctrl.Log.WithName("controllers").WithName("PostgreSQLServer"),
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
			Telemetry: telemetry.InitializeTelemetryDefault(
				"PostgreSQLDatabase",
				ctrl.Log.WithName("controllers").WithName("PostgreSQLDatabase"),
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
			Telemetry: telemetry.InitializeTelemetryDefault(
				"PostgreSQLFirewallRule",
				ctrl.Log.WithName("controllers").WithName("PostgreSQLFirewallRule"),
			),
			Recorder: k8sManager.GetEventRecorderFor("PostgreSQLFirewallRule-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&StorageReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: storageAccountManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"Storage",
				ctrl.Log.WithName("controllers").WithName("Storage"),
			),
			Recorder: k8sManager.GetEventRecorderFor("Storage-controller"),
			Scheme:   scheme.Scheme,
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
		redisCacheManager:       redisCacheManager,
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
	_, _ = storageAccountManager.CreateStorage(context.Background(), resourceGroupName, storageAccountName, resourcegroupLocation, azurev1alpha1.StorageSku{
		Name: "Standard_LRS",
	}, "Storage", map[string]*string{}, "", nil, nil)

	// Storage account needs to be in "Suceeded" state
	// for container create to succeed
	finish = time.Now().Add(tc.timeout)
	for {

		if finish.Before(time.Now()) {
			return fmt.Errorf("time out waiting for storage account")
		}

		result, _ := storageAccountManager.GetStorage(context.Background(), resourceGroupName, storageAccountName)
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

func PanicRecover(t *testing.T) {
	if err := recover(); err != nil {
		t.Logf("caught panic in test: %v", err)
		t.Logf("stacktrace from panic: \n%s", string(debug.Stack()))
		t.Fail()
	}
}
