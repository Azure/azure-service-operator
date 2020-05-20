// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"flag"
	"os"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/controllers"
	"github.com/Azure/go-autorest/autorest/azure/auth"

	kscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	healthz "sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha2 "github.com/Azure/azure-service-operator/api/v1alpha2"
	azurev1beta1 "github.com/Azure/azure-service-operator/api/v1beta1"
	resourceapimanagement "github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimgmt"
	apimservice "github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimservice"
	resourcemanagerappinsights "github.com/Azure/azure-service-operator/pkg/resourcemanager/appinsights"
	resourcemanagersqlaction "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlaction"
	resourcemanagersqldb "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqldb"
	resourcemanagersqlfailovergroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfailovergroup"
	resourcemanagersqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfirewallrule"
	resourcemanagersqlmanageduser "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlmanageduser"
	resourcemanagersqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlserver"
	resourcemanagersqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
	resourcemanagersqlvnetrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlvnetrule"
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagercosmosdb "github.com/Azure/azure-service-operator/pkg/resourcemanager/cosmosdbs"
	resourcemanagereventhub "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resourcemanagerkeyvault "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	loadbalancer "github.com/Azure/azure-service-operator/pkg/resourcemanager/loadbalancer"
	mysqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/database"
	mysqlfirewall "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/firewallrule"
	mysqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/server"
	mysqlvnetrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/vnetrule"
	nic "github.com/Azure/azure-service-operator/pkg/resourcemanager/nic"
	pip "github.com/Azure/azure-service-operator/pkg/resourcemanager/pip"
	psqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/database"
	psqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/firewallrule"
	psqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/server"
	psqlvnetrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/vnetrule"
	resourcemanagerrediscache "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches"
	resourcemanagerresourcegroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	blobContainerManager "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/blobcontainer"
	storageaccountManager "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"
	vm "github.com/Azure/azure-service-operator/pkg/resourcemanager/vm"
	vmss "github.com/Azure/azure-service-operator/pkg/resourcemanager/vmss"
	vnet "github.com/Azure/azure-service-operator/pkg/resourcemanager/vnet"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	keyvaultSecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
	k8sSecrets "github.com/Azure/azure-service-operator/pkg/secrets/kube"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"

	// +kubebuilder:scaffold:imports
	"net/http"
)

var (
	masterURL, kubeconfig, resources, clusterName               string
	cloudName, tenantID, subscriptionID, clientID, clientSecret string
	useAADPodIdentity                                           bool

	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {

	_ = kscheme.AddToScheme(scheme)
	_ = azurev1alpha1.AddToScheme(scheme)
	_ = azurev1beta1.AddToScheme(scheme)
	_ = azurev1alpha2.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=events,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=azuresqlusers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=get;update;patch

func main() {
	var metricsAddr string
	var healthAddr string
	var enableLeaderElection bool
	var secretClient secrets.SecretClient
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&healthAddr, "health-addr", ":8081", "The address the health endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")

	flag.Parse()

	ctrl.SetLogger(zap.Logger(true))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		HealthProbeBindAddress: healthAddr,
		LeaderElection:         enableLeaderElection,
		LivenessEndpointName:   "/healthz",
	})

	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	err = resourcemanagerconfig.ParseEnvironment()
	if err != nil {
		setupLog.Error(err, "unable to parse settings required to provision resources in Azure")
		os.Exit(1)
	}

	keyvaultName := resourcemanagerconfig.OperatorKeyvault()

	if keyvaultName == "" {
		setupLog.Info("Keyvault name is empty")
		secretClient = k8sSecrets.New(mgr.GetClient())
	} else {
		setupLog.Info("Instantiating secrets client for keyvault " + keyvaultName)
		secretClient = keyvaultSecrets.New(keyvaultName)
	}

	apimManager := resourceapimanagement.NewManager()
	apimServiceManager := apimservice.NewAzureAPIMgmtServiceManager()
	vnetManager := vnet.NewAzureVNetManager()
	resourceGroupManager := resourcemanagerresourcegroup.NewAzureResourceGroupManager()

	redisCacheManager := resourcemanagerrediscache.NewAzureRedisCacheManager(
		secretClient,
		scheme,
	)
	appInsightsManager := resourcemanagerappinsights.NewManager(
		secretClient,
		scheme,
	)
	eventhubNamespaceClient := resourcemanagereventhub.NewEventHubNamespaceClient()
	consumerGroupClient := resourcemanagereventhub.NewConsumerGroupClient()
	cosmosDBClient := resourcemanagercosmosdb.NewAzureCosmosDBManager(
		secretClient,
	)
	keyVaultManager := resourcemanagerkeyvault.NewAzureKeyVaultManager(mgr.GetScheme())
	keyVaultKeyManager := &resourcemanagerkeyvault.KeyvaultKeyClient{
		KeyvaultClient: keyVaultManager,
	}
	eventhubClient := resourcemanagereventhub.NewEventhubClient(secretClient, scheme)
	sqlServerManager := resourcemanagersqlserver.NewAzureSqlServerManager(
		secretClient,
		scheme,
	)
	sqlDBManager := resourcemanagersqldb.NewAzureSqlDbManager()
	sqlFirewallRuleManager := resourcemanagersqlfirewallrule.NewAzureSqlFirewallRuleManager()
	sqlVNetRuleManager := resourcemanagersqlvnetrule.NewAzureSqlVNetRuleManager()
	sqlFailoverGroupManager := resourcemanagersqlfailovergroup.NewAzureSqlFailoverGroupManager(
		secretClient,
		scheme,
	)
	psqlserverclient := psqlserver.NewPSQLServerClient(secretClient, mgr.GetScheme())
	psqldatabaseclient := psqldatabase.NewPSQLDatabaseClient()
	psqlfirewallruleclient := psqlfirewallrule.NewPSQLFirewallRuleClient()
	sqlUserManager := resourcemanagersqluser.NewAzureSqlUserManager(
		secretClient,
		scheme,
	)
	sqlManagedUserManager := resourcemanagersqlmanageduser.NewAzureSqlManagedUserManager(
		secretClient,
		scheme,
	)
	sqlActionManager := resourcemanagersqlaction.NewAzureSqlActionManager(secretClient, scheme)

	var AzureHealthCheck healthz.Checker = func(_ *http.Request) error {
		_, err := auth.NewAuthorizerFromEnvironment()
		if err != nil {
			return err
		}

		return nil
	}
	if err := mgr.AddHealthzCheck("azurehealthz", AzureHealthCheck); err != nil {
		setupLog.Error(err, "problem running health check to azure autorizer")
	}

	err = (&controllers.StorageAccountReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: storageaccountManager.New(secretClient, scheme),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"StorageAccount",
				ctrl.Log.WithName("controllers").WithName("StorageAccount"),
			),
			Recorder: mgr.GetEventRecorderFor("StorageAccount-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "StorageAccount")
		os.Exit(1)
	}
	err = (&controllers.CosmosDBReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: cosmosDBClient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"CosmosDB",
				ctrl.Log.WithName("controllers").WithName("CosmosDB"),
			),
			Recorder: mgr.GetEventRecorderFor("CosmosDB-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "CosmosDB")
		os.Exit(1)
	}
	err = (&controllers.RedisCacheReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: redisCacheManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"RedisCache",
				ctrl.Log.WithName("controllers").WithName("RedisCache"),
			),
			Recorder: mgr.GetEventRecorderFor("RedisCache-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "RedisCache")
		os.Exit(1)
	}

	err = (&controllers.EventhubReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: eventhubClient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"Eventhub",
				ctrl.Log.WithName("controllers").WithName("Eventhub"),
			),
			Recorder: mgr.GetEventRecorderFor("Eventhub-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Eventhub")
		os.Exit(1)
	}

	err = (&controllers.ResourceGroupReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: resourceGroupManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"ResourceGroup",
				ctrl.Log.WithName("controllers").WithName("ResourceGroup"),
			),
			Recorder: mgr.GetEventRecorderFor("ResourceGroup-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ResourceGroup")
		os.Exit(1)
	}

	err = (&controllers.EventhubNamespaceReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: eventhubNamespaceClient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"EventhubNamespace",
				ctrl.Log.WithName("controllers").WithName("EventhubNamespace"),
			),
			Recorder: mgr.GetEventRecorderFor("EventhubNamespace-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "EventhubNamespace")
		os.Exit(1)
	}

	err = (&controllers.KeyVaultReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: keyVaultManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"KeyVault",
				ctrl.Log.WithName("controllers").WithName("KeyVault"),
			),
			Recorder: mgr.GetEventRecorderFor("KeyVault-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "KeyVault")
		os.Exit(1)
	}

	err = (&controllers.ConsumerGroupReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: consumerGroupClient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"ConsumerGroup",
				ctrl.Log.WithName("controllers").WithName("ConsumerGroup"),
			),
			Recorder: mgr.GetEventRecorderFor("ConsumerGroup-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ConsumerGroup")
		os.Exit(1)
	}

	if err = (&controllers.AzureSqlServerReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: sqlServerManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSqlServer",
				ctrl.Log.WithName("controllers").WithName("AzureSqlServer"),
			),
			Recorder: mgr.GetEventRecorderFor("AzureSqlServer-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "AzureSqlServer")
		os.Exit(1)
	}

	/* Azure Sql Database */
	err = (&controllers.AzureSqlDatabaseReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: sqlDBManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSqlDb",
				ctrl.Log.WithName("controllers").WithName("AzureSqlDb"),
			),
			Recorder: mgr.GetEventRecorderFor("AzureSqlDb-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "AzureSqlDb")
		os.Exit(1)
	}

	if err = (&controllers.AzureSqlFirewallRuleReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: sqlFirewallRuleManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSQLFirewallRuleOperator",
				ctrl.Log.WithName("controllers").WithName("AzureSQLFirewallRuleOperator"),
			),
			Recorder: mgr.GetEventRecorderFor("SqlFirewallRule-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SqlFirewallRule")
		os.Exit(1)
	}

	if err = (&controllers.AzureSQLVNetRuleReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: sqlVNetRuleManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSQLVNetRuleOperator",
				ctrl.Log.WithName("controllers").WithName("AzureSQLVNetRuleOperator"),
			),
			Recorder: mgr.GetEventRecorderFor("SqlVnetRule-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SqlVNetRule")
		os.Exit(1)
	}

	if err = (&controllers.AzureSqlActionReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: sqlActionManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSQLActionOperator",
				ctrl.Log.WithName("controllers").WithName("AzureSQLActionOperator"),
			),
			Recorder: mgr.GetEventRecorderFor("SqlAction-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "SqlAction")
		os.Exit(1)
	}

	if err = (&controllers.AzureSQLUserReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: sqlUserManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSQLUser",
				ctrl.Log.WithName("controllers").WithName("AzureSQLUser"),
			),
			Recorder: mgr.GetEventRecorderFor("AzureSQLUser-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "AzureSQLUser")
		os.Exit(1)
	}

	if err = (&controllers.AzureSQLManagedUserReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: sqlManagedUserManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSQLManagedUser",
				ctrl.Log.WithName("controllers").WithName("AzureSQLManagedUser"),
			),
			Recorder: mgr.GetEventRecorderFor("AzureSQLManagedUser-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "AzureSQLManagedUser")
		os.Exit(1)
	}

	if err = (&controllers.AzureSqlFailoverGroupReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: sqlFailoverGroupManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AzureSqlFailoverGroup",
				ctrl.Log.WithName("controllers").WithName("AzureSqlFailoverGroup"),
			),
			Recorder: mgr.GetEventRecorderFor("AzureSqlFailoverGroup-controller"),
			Scheme:   mgr.GetScheme(),
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "AzureSqlFailoverGroup")
		os.Exit(1)
	}

	if err = (&controllers.BlobContainerReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: blobContainerManager.New(),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"BlobContainer",
				ctrl.Log.WithName("controllers").WithName("BlobContainer"),
			),
			Recorder: mgr.GetEventRecorderFor("BlobContainer-controller"),
			Scheme:   mgr.GetScheme(),
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "BlobContainer")
		os.Exit(1)
	}

	if err = (&controllers.AppInsightsReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: appInsightsManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AppInsights",
				ctrl.Log.WithName("controllers").WithName("AppInsights"),
			),
			Recorder: mgr.GetEventRecorderFor("AppInsights-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "AppInsights")
		os.Exit(1)
	}

	if err = (&controllers.PostgreSQLServerReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: psqlserverclient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"PostgreSQLServer",
				ctrl.Log.WithName("controllers").WithName("PostgreSQLServer"),
			),
			Recorder: mgr.GetEventRecorderFor("PostgreSQLServer-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "PostgreSQLServer")
		os.Exit(1)
	}

	if err = (&controllers.PostgreSQLDatabaseReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: psqldatabaseclient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"PostgreSQLDatabase",
				ctrl.Log.WithName("controllers").WithName("PostgreSQLDatabase"),
			),
			Recorder: mgr.GetEventRecorderFor("PostgreSQLDatabase-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "PostgreSQLDatabase")
		os.Exit(1)
	}

	if err = (&controllers.PostgreSQLFirewallRuleReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: psqlfirewallruleclient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"PostgreSQLFirewallRule",
				ctrl.Log.WithName("controllers").WithName("PostgreSQLFirewallRule"),
			),
			Recorder: mgr.GetEventRecorderFor("PostgreSQLFirewallRule-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "PostgreSQLFirewallRule")
		os.Exit(1)
	}

	if err = (&controllers.ApimServiceReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: apimServiceManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"ApimService",
				ctrl.Log.WithName("controllers").WithName("ApimService"),
			),
			Recorder: mgr.GetEventRecorderFor("ApimService-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ApimService")
		os.Exit(1)
	}

	if err = (&controllers.VirtualNetworkReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: vnetManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"VNet",
				ctrl.Log.WithName("controllers").WithName("VNet"),
			),
			Recorder: mgr.GetEventRecorderFor("VNet-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "VNet")
		os.Exit(1)
	}

	if err = (&controllers.APIMAPIReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: apimManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"APIManagement",
				ctrl.Log.WithName("controllers").WithName("APIManagement"),
			),
			Recorder: mgr.GetEventRecorderFor("APIManagement-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "APIMgmtAPI")
		os.Exit(1)
	}

	if err = (&controllers.KeyVaultKeyReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: keyVaultKeyManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"KeyVaultKey",
				ctrl.Log.WithName("controllers").WithName("KeyVaultKey"),
			),
			Recorder: mgr.GetEventRecorderFor("KeyVaultKey-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "KeyVaultKey")
		os.Exit(1)
	}

	if err = (&controllers.MySQLServerReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			AzureClient: mysqlserver.NewMySQLServerClient(
				secretClient,
				mgr.GetScheme(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLServer",
				ctrl.Log.WithName("controllers").WithName("MySQLServer"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLServer-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "MySQLServer")
		os.Exit(1)
	}
	if err = (&controllers.MySQLDatabaseReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: mysqldatabase.NewMySQLDatabaseClient(),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLDatabase",
				ctrl.Log.WithName("controllers").WithName("MySQLDatabase"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLDatabase-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "MySQLDatabase")
		os.Exit(1)
	}
	if err = (&controllers.MySQLFirewallRuleReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: mysqlfirewall.NewMySQLFirewallRuleClient(),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLFirewallRule",
				ctrl.Log.WithName("controllers").WithName("MySQLFirewallRule"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLFirewallRule-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "MySQLFirewallRule")
		os.Exit(1)
	}

	if err = (&controllers.AzurePublicIPAddressReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			AzureClient: pip.NewAzurePublicIPAddressClient(
				secretClient,
				mgr.GetScheme(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"PublicIPAddress",
				ctrl.Log.WithName("controllers").WithName("PublicIPAddress"),
			),
			Recorder: mgr.GetEventRecorderFor("PublicIPAddress-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "PublicIPAddress")
		os.Exit(1)
	}

	if err = (&controllers.AzureNetworkInterfaceReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			AzureClient: nic.NewAzureNetworkInterfaceClient(
				secretClient,
				mgr.GetScheme(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"NetworkInterface",
				ctrl.Log.WithName("controllers").WithName("NetworkInterface"),
			),
			Recorder: mgr.GetEventRecorderFor("NetworkInterface-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "NetworkInterface")
		os.Exit(1)
	}

	if err = (&controllers.MySQLVNetRuleReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: mysqlvnetrule.NewMySQLVNetRuleClient(),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLVNetRule",
				ctrl.Log.WithName("controllers").WithName("MySQLVNetRule"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLVNetRule-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "MySQLVNetRule")
		os.Exit(1)
	}

	if err = (&controllers.AzureVirtualMachineReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			AzureClient: vm.NewAzureVirtualMachineClient(
				secretClient,
				mgr.GetScheme(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"VirtualMachine",
				ctrl.Log.WithName("controllers").WithName("VirtualMachine"),
			),
			Recorder: mgr.GetEventRecorderFor("VirtualMachine-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "VirtualMachine")
		os.Exit(1)
	}

	if err = (&controllers.PostgreSQLVNetRuleReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: psqlvnetrule.NewPostgreSQLVNetRuleClient(),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"PostgreSQLVNetRule",
				ctrl.Log.WithName("controllers").WithName("PostgreSQLVNetRule"),
			),
			Recorder: mgr.GetEventRecorderFor("PostgreSQLVNetRule-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "PostgreSQLVNetRule")
		os.Exit(1)
	}

	if err = (&controllers.AzureLoadBalancerReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			AzureClient: loadbalancer.NewAzureLoadBalancerClient(
				secretClient,
				mgr.GetScheme(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"LoadBalancer",
				ctrl.Log.WithName("controllers").WithName("LoadBalancer"),
			),
			Recorder: mgr.GetEventRecorderFor("LoadBalancer-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "LoadBalancer")
		os.Exit(1)
	}

	if err = (&controllers.AzureVMScaleSetReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			AzureClient: vmss.NewAzureVMScaleSetClient(
				secretClient,
				mgr.GetScheme(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"VMScaleSet",
				ctrl.Log.WithName("controllers").WithName("VMScaleSet"),
			),
			Recorder: mgr.GetEventRecorderFor("VMScaleSet-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "VMScaleSet")
		os.Exit(1)
	}

	if err = (&v1alpha1.AzureSqlServer{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "AzureSqlServer")
		os.Exit(1)
	}
	if err = (&azurev1alpha1.AzureSqlDatabase{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "AzureSqlDatabase")
		os.Exit(1)
	}
	if err = (&azurev1alpha1.AzureSqlFirewallRule{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "AzureSqlFirewallRule")
		os.Exit(1)
	}
	if err = (&azurev1alpha1.AzureSqlFailoverGroup{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "AzureSqlFailoverGroup")
		os.Exit(1)
	}
	if err = (&azurev1alpha1.BlobContainer{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "BlobContainer")
		os.Exit(1)
	}

	if err = (&azurev1alpha1.MySQLServer{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "MySQLServer")
		os.Exit(1)
	}
	if err = (&azurev1alpha1.PostgreSQLServer{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "PostgreSQLServer")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}

}
