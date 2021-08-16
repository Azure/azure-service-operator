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

	"github.com/pkg/errors"

	"github.com/gobuffalo/envy"
	"k8s.io/client-go/kubernetes/scheme"
	kscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	resourcemanagerapimgmt "github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimgmt"
	resourcemanagerappinsights "github.com/Azure/azure-service-operator/pkg/resourcemanager/appinsights"
	resourcemanagersqlaction "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlaction"
	resourcemanagersqldb "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqldb"
	resourcemanagersqlfailovergroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfailovergroup"
	resourcemanagersqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfirewallrule"
	resourcemanagersqlmanageduser "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlmanageduser"
	resourcemanagersqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlserver"
	resourcemanagersqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
	resourcemanagersqlvnetrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlvnetrule"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagercosmosdbaccount "github.com/Azure/azure-service-operator/pkg/resourcemanager/cosmosdb/account"
	resourcemanagercosmosdbsqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/cosmosdb/sqldatabase"
	resourcemanagereventhub "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resourcemanagerkeyvaults "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/loadbalancer"
	mysqladmin "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/aadadmin"
	mysqlDatabaseManager "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/database"
	mysqlFirewallManager "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/firewallrule"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/mysqlaaduser"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/mysqluser"
	mysqlServerManager "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/server"
	mysqlvnetrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/vnetrule"
	resourcemanagernic "github.com/Azure/azure-service-operator/pkg/resourcemanager/nic"
	resourcemanagerpip "github.com/Azure/azure-service-operator/pkg/resourcemanager/pip"
	resourcemanagerpsqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/database"
	resourcemanagerpsqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/firewallrule"
	resourcemanagerpsqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/psqluser"
	resourcemanagerpsqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/server"
	rediscacheactions "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches/actions"
	rcfwr "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches/firewallrule"
	rediscaches "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches/redis"
	resourcegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	resourcemanagerblobcontainer "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/blobcontainer"
	resourcemanagerstorageaccount "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/vm"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/vmext"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/vmss"
	resourcemanagervnet "github.com/Azure/azure-service-operator/pkg/resourcemanager/vnet"
	k8sSecrets "github.com/Azure/azure-service-operator/pkg/secrets/kube"
	"github.com/Azure/azure-service-operator/pkg/telemetry"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var testEnv *envtest.Environment

var tc TestContext

func setup() error {
	log.Println(fmt.Sprintf("Starting common controller test setup"))

	// Go ahead and assume that we're deployed in the default namespace for
	// the purpose of these tests
	envy.Set("POD_NAMESPACE", "azureoperator-system")

	// Uncomment the below to run the tests in the old v1 naming mode
	// envy.Set("AZURE_SECRET_NAMING_VERSION", "1")

	err := config.ParseEnvironment()
	if err != nil {
		return err
	}

	resourceGroupName := GenerateTestResourceNameWithRandom(TestResourceGroupPrefix, 6)
	resourceGroupLocation := config.DefaultLocation()

	keyvaultName := GenerateAlphaNumTestResourceNameWithRandom("kv-prime", 5)

	var timeout time.Duration

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
			WebhookInstallOptions: envtest.WebhookInstallOptions{
				Paths: []string{
					"../config/webhook",
				},
			},
		}
		cfg, err = testEnv.Start()
		if err != nil {
			return err
		}
	}

	if cfg == nil {
		return fmt.Errorf("rest config nil")
	}

	err = kscheme.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}
	err = azurev1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}
	err = v1alpha2.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}
	err = v1beta1.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	var k8sManager ctrl.Manager

	targetNamespaces := config.TargetNamespaces()
	var cacheFunc cache.NewCacheFunc
	if targetNamespaces != nil {
		log.Println("Restricting operator cache to namespaces", targetNamespaces)
		cacheFunc = cache.MultiNamespacedCacheBuilder(targetNamespaces)
	}

	// +kubebuilder:scaffold:scheme
	k8sManager, err = ctrl.NewManager(cfg, ctrl.Options{
		Scheme:   scheme.Scheme,
		CertDir:  testEnv.WebhookInstallOptions.LocalServingCertDir,
		Port:     testEnv.WebhookInstallOptions.LocalServingPort,
		NewCache: cacheFunc,
	})
	if err != nil {
		return err
	}

	secretClient := k8sSecrets.New(k8sManager.GetClient(), config.SecretNamingVersion())
	resourceGroupManager := resourcegroupsresourcemanager.NewAzureResourceGroupManager(config.GlobalCredentials())
	keyVaultManager := resourcemanagerkeyvaults.NewAzureKeyVaultManager(config.GlobalCredentials(), k8sManager.GetScheme())
	eventhubClient := resourcemanagereventhub.NewEventhubClient(config.GlobalCredentials(), secretClient, scheme.Scheme)
	consumerGroupClient := resourcemanagereventhub.NewConsumerGroupClient(config.GlobalCredentials())
	azureSqlDatabaseManager := resourcemanagersqldb.NewAzureSqlDbManager(config.GlobalCredentials())

	timeout = time.Second * 780

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
			Client: k8sManager.GetClient(),
			AzureClient: resourcemanagerkeyvaults.NewKeyvaultKeyClient(
				config.GlobalCredentials(),
				keyVaultManager,
			),
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
			Client: k8sManager.GetClient(),
			AzureClient: resourcemanagerappinsights.NewManager(
				config.GlobalCredentials(),
				secretClient,
				scheme.Scheme,
			),
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

	err = (&AppInsightsApiKeyReconciler{
		Reconciler: &AsyncReconciler{
			Client: k8sManager.GetClient(),
			AzureClient: resourcemanagerappinsights.NewAPIKeyClient(
				config.GlobalCredentials(),
				secretClient,
				scheme.Scheme,
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AppInsightsApiKey",
				ctrl.Log.WithName("controllers").WithName("AppInsightsApiKey"),
			),
			Recorder: k8sManager.GetEventRecorderFor("AppInsightsApiKey-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&APIMAPIReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: resourcemanagerapimgmt.NewManager(config.GlobalCredentials()),
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

	err = (&CosmosDBReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: resourcemanagercosmosdbaccount.NewAzureCosmosDBManager(config.GlobalCredentials(), secretClient),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"CosmosDB",
				ctrl.Log.WithName("controllers").WithName("CosmosDB"),
			),
			Recorder: k8sManager.GetEventRecorderFor("CosmosDB-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&CosmosDBSQLDatabaseReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: resourcemanagercosmosdbsqldatabase.NewAzureCosmosDBSQLDatabaseManager(config.GlobalCredentials()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"CosmosDBSQLDatabase",
				ctrl.Log.WithName("controllers").WithName("CosmosDBSQLDatabase"),
			),
			Recorder: k8sManager.GetEventRecorderFor("CosmosDBSQLDatabase-controller"),
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
			Client: k8sManager.GetClient(),
			AzureClient: rediscaches.NewAzureRedisCacheManager(
				config.GlobalCredentials(),
				secretClient,
				scheme.Scheme,
			),
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

	err = (&RedisCacheActionReconciler{
		Reconciler: &AsyncReconciler{
			Client: k8sManager.GetClient(),
			AzureClient: rediscacheactions.NewAzureRedisCacheActionManager(
				config.GlobalCredentials(),
				secretClient,
				scheme.Scheme,
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"RedisCacheAction",
				ctrl.Log.WithName("controllers").WithName("RedisCacheAction"),
			),
			Recorder: k8sManager.GetEventRecorderFor("RedisCacheAction-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&RedisCacheFirewallRuleReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: rcfwr.NewAzureRedisCacheFirewallRuleManager(config.GlobalCredentials()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"RedisCacheFirewallRule",
				ctrl.Log.WithName("controllers").WithName("RedisCacheFirewallRule"),
			),
			Recorder: k8sManager.GetEventRecorderFor("RedisCacheFirewallRule-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&EventhubNamespaceReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: resourcemanagereventhub.NewEventHubNamespaceClient(config.GlobalCredentials()),
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

	err = (&AzureSqlServerReconciler{
		Reconciler: &AsyncReconciler{
			Client: k8sManager.GetClient(),
			AzureClient: resourcemanagersqlserver.NewAzureSqlServerManager(
				config.GlobalCredentials(),
				secretClient,
				scheme.Scheme,
			),
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
			AzureClient: azureSqlDatabaseManager,
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
			AzureClient: resourcemanagersqlfirewallrule.NewAzureSqlFirewallRuleManager(config.GlobalCredentials()),
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
			AzureClient: resourcemanagersqlvnetrule.NewAzureSqlVNetRuleManager(config.GlobalCredentials()),
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
			Client: k8sManager.GetClient(),
			AzureClient: resourcemanagersqlfailovergroup.NewAzureSqlFailoverGroupManager(
				config.GlobalCredentials(),
				secretClient,
				scheme.Scheme,
			),
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
			Client: k8sManager.GetClient(),
			AzureClient: resourcemanagersqluser.NewAzureSqlUserManager(
				config.GlobalCredentials(),
				secretClient,
				scheme.Scheme,
			),
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

	err = (&AzureSQLManagedUserReconciler{
		Reconciler: &AsyncReconciler{
			Client: k8sManager.GetClient(),
			AzureClient: resourcemanagersqlmanageduser.NewAzureSqlManagedUserManager(
				config.GlobalCredentials(),
				secretClient,
				scheme.Scheme,
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSqlManagedUser",
				ctrl.Log.WithName("controllers").WithName("AzureSqlManagedUser"),
			),
			Recorder: k8sManager.GetEventRecorderFor("AzureSqlManagedUser-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&VirtualNetworkReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: resourcemanagervnet.NewAzureVNetManager(config.GlobalCredentials()),
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

	err = (&AzurePublicIPAddressReconciler{
		Reconciler: &AsyncReconciler{
			Client: k8sManager.GetClient(),
			AzureClient: resourcemanagerpip.NewAzurePublicIPAddressClient(
				config.GlobalCredentials(),
				secretClient,
				k8sManager.GetScheme(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"PublicIPAddress",
				ctrl.Log.WithName("controllers").WithName("PublicIPAddress"),
			),
			Recorder: k8sManager.GetEventRecorderFor("PublicIPAddress-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureNetworkInterfaceReconciler{
		Reconciler: &AsyncReconciler{
			Client: k8sManager.GetClient(),
			AzureClient: resourcemanagernic.NewAzureNetworkInterfaceClient(
				config.GlobalCredentials(),
				secretClient,
				k8sManager.GetScheme(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"NetworkInterface",
				ctrl.Log.WithName("controllers").WithName("NetworkInterface"),
			),
			Recorder: k8sManager.GetEventRecorderFor("NetworkInterface-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureVirtualMachineReconciler{
		Reconciler: &AsyncReconciler{
			Client: k8sManager.GetClient(),
			AzureClient: vm.NewAzureVirtualMachineClient(
				config.GlobalCredentials(),
				secretClient,
				k8sManager.GetScheme(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"VirtualMachine",
				ctrl.Log.WithName("controllers").WithName("VirtualMachine"),
			),
			Recorder: k8sManager.GetEventRecorderFor("VirtualMachine-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureVirtualMachineExtensionReconciler{
		Reconciler: &AsyncReconciler{
			Client: k8sManager.GetClient(),
			AzureClient: vmext.NewAzureVirtualMachineExtensionClient(
				config.GlobalCredentials(),
				secretClient,
				k8sManager.GetScheme(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"VirtualMachineExtension",
				ctrl.Log.WithName("controllers").WithName("VirtualMachineExtension"),
			),
			Recorder: k8sManager.GetEventRecorderFor("VirtualMachineExtension-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureLoadBalancerReconciler{
		Reconciler: &AsyncReconciler{
			Client: k8sManager.GetClient(),
			AzureClient: loadbalancer.NewAzureLoadBalancerClient(
				config.GlobalCredentials(),
				secretClient,
				k8sManager.GetScheme(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"LoadBalancer",
				ctrl.Log.WithName("controllers").WithName("LoadBalancer"),
			),
			Recorder: k8sManager.GetEventRecorderFor("LoadBalancer-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureVMScaleSetReconciler{
		Reconciler: &AsyncReconciler{
			Client: k8sManager.GetClient(),
			AzureClient: vmss.NewAzureVMScaleSetClient(
				config.GlobalCredentials(),
				secretClient,
				k8sManager.GetScheme(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"VMScaleSet",
				ctrl.Log.WithName("controllers").WithName("VMScaleSet"),
			),
			Recorder: k8sManager.GetEventRecorderFor("VMScaleSet-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&AzureSqlActionReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: resourcemanagersqlaction.NewAzureSqlActionManager(config.GlobalCredentials(), secretClient, scheme.Scheme),
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
			AzureClient: resourcemanagerblobcontainer.New(config.GlobalCredentials()),
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

	err = (&MySQLServerReconciler{
		Reconciler: &AsyncReconciler{
			Client: k8sManager.GetClient(),
			AzureClient: mysqlServerManager.NewMySQLServerClient(
				config.GlobalCredentials(),
				secretClient,
				k8sManager.GetScheme(),
				k8sManager.GetClient(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLServer",
				ctrl.Log.WithName("controllers").WithName("MySQLServer"),
			),
			Recorder: k8sManager.GetEventRecorderFor("MySQLServer-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&MySQLDatabaseReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: mysqlDatabaseManager.NewMySQLDatabaseClient(config.GlobalCredentials()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLDatabase",
				ctrl.Log.WithName("controllers").WithName("MySQLDatabase"),
			),
			Recorder: k8sManager.GetEventRecorderFor("MySQLDatabase-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&MySQLFirewallRuleReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: mysqlFirewallManager.NewMySQLFirewallRuleClient(config.GlobalCredentials()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLFirewallRule",
				ctrl.Log.WithName("controllers").WithName("MySQLFirewallRule"),
			),
			Recorder: k8sManager.GetEventRecorderFor("MySQLFirewallRule-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&MySQLVNetRuleReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: mysqlvnetrule.NewMySQLVNetRuleClient(config.GlobalCredentials()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLVNetRule",
				ctrl.Log.WithName("controllers").WithName("MySQLVNetRule"),
			),
			Recorder: k8sManager.GetEventRecorderFor("MySQLVNetRule-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&MySQLUserReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: mysqluser.NewMySqlUserManager(config.GlobalCredentials(), secretClient, scheme.Scheme),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLUser",
				ctrl.Log.WithName("controllers").WithName("MySQLUser"),
			),
			Recorder: k8sManager.GetEventRecorderFor("MySQLUser-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	identityFinder := helpers.NewAADIdentityFinder(k8sManager.GetClient(), config.PodNamespace())
	err = (&MySQLAADUserReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: mysqlaaduser.NewMySQLAADUserManager(config.GlobalCredentials(), identityFinder),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLAADUser",
				ctrl.Log.WithName("controllers").WithName("MySQLAADUser"),
			),
			Recorder: k8sManager.GetEventRecorderFor("MySQLAADUser-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&MySQLServerAdministratorReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: mysqladmin.NewMySQLServerAdministratorManager(config.GlobalCredentials()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLServerAdministrator",
				ctrl.Log.WithName("controllers").WithName("MySQLServerAdministrator"),
			),
			Recorder: k8sManager.GetEventRecorderFor("MySQLServerAdministrator-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&PostgreSQLServerReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: resourcemanagerpsqlserver.NewPSQLServerClient(config.GlobalCredentials(), secretClient, k8sManager.GetScheme()),
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
			AzureClient: resourcemanagerpsqldatabase.NewPSQLDatabaseClient(config.GlobalCredentials()),
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
			AzureClient: resourcemanagerpsqlfirewallrule.NewPSQLFirewallRuleClient(config.GlobalCredentials()),
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

	err = (&PostgreSQLUserReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: resourcemanagerpsqluser.NewPostgreSqlUserManager(config.GlobalCredentials(), secretClient, k8sManager.GetScheme()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"PostgreSQLUser",
				ctrl.Log.WithName("controllers").WithName("PostgreSQLUser"),
			),
			Recorder: k8sManager.GetEventRecorderFor("PostgreSQLUser-controller"),
			Scheme:   k8sManager.GetScheme(),
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	err = (&StorageAccountReconciler{
		Reconciler: &AsyncReconciler{
			Client:      k8sManager.GetClient(),
			AzureClient: resourcemanagerstorageaccount.New(config.GlobalCredentials(), secretClient, k8sManager.GetScheme()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"StorageAccount",
				ctrl.Log.WithName("controllers").WithName("StorageAccount"),
			),
			Recorder: k8sManager.GetEventRecorderFor("StorageAccount-controller"),
			Scheme:   scheme.Scheme,
		},
	}).SetupWithManager(k8sManager)
	if err != nil {
		return err
	}

	// Webhooks
	if err = (&azurev1alpha1.AzureSQLUser{}).SetupWebhookWithManager(k8sManager); err != nil {
		return err
	}
	if err = (&azurev1alpha1.AzureSQLManagedUser{}).SetupWebhookWithManager(k8sManager); err != nil {
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
		_, err = resourceGroupManager.CreateGroup(context.Background(), resourceGroupName, resourceGroupLocation)
		if err != nil {
			return errors.Wrap(err, "resource creation failed")
		}
	}

	tc = TestContext{
		k8sClient:             k8sClient,
		secretClient:          secretClient,
		resourceGroupName:     resourceGroupName,
		resourceGroupLocation: resourceGroupLocation,
		keyvaultName:          keyvaultName,
		eventhubClient:        eventhubClient,
		resourceGroupManager:  resourceGroupManager,
		keyVaultManager:       keyVaultManager,
		sqlDbManager:          azureSqlDatabaseManager,
		timeout:               timeout,
		timeoutFast:           time.Minute * 3,
		retry:                 time.Second * 3,
		consumerGroupClient:   consumerGroupClient,
	}

	log.Println("Creating KV:", keyvaultName)
	err = CreateVaultWithAccessPolicies(
		context.Background(),
		config.GlobalCredentials(),
		resourceGroupName,
		keyvaultName,
		resourceGroupLocation,
		config.GlobalCredentials().ClientID(),
	)
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
		log.Println(fmt.Sprintf("could not set up environment: %s\n", err))
		os.Exit(1)
	}

	code = m.Run()

	err = teardown()
	if err != nil {
		log.Println(fmt.Sprintf("could not tear down environment: %s\n; original exit code: %d\n", err, code))
	}

	os.Exit(code)
}

func PanicRecover(t *testing.T) {
	if err := recover(); err != nil {
		t.Logf("caught panic in test: %s", err)
		t.Logf("stacktrace from panic: \n%s", string(debug.Stack()))
		t.Fail()
	}
}
