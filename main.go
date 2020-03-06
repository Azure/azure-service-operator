// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"flag"
	"os"

	"github.com/Azure/azure-service-operator/controllers"
	"k8s.io/apimachinery/pkg/runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	resourceapimanagement "github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimgmt"
	apimservice "github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimservice"
	resourcemanagerappinsights "github.com/Azure/azure-service-operator/pkg/resourcemanager/appinsights"
	resourcemanagersqldb "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqldb"
	resourcemanagersqlfailovergroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfailovergroup"
	resourcemanagersqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfirewallrule"
	resourcemanagersqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlserver"
	resourcemanagersqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagereventhub "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resourcemanagerkeyvault "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	psqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/database"
	psqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/firewallrule"
	psqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/server"
	resourcemanagerrediscache "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches"
	resourcemanagerresourcegroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	resourcemanagerstorage "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages"
	vnet "github.com/Azure/azure-service-operator/pkg/resourcemanager/vnet"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	keyvaultSecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
	k8sSecrets "github.com/Azure/azure-service-operator/pkg/secrets/kube"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
	kscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
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
	var enableLeaderElection bool
	var secretClient secrets.SecretClient
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")

	flag.Parse()

	ctrl.SetLogger(zap.Logger(true))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
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

	apimManager := resourceapimanagement.NewManager(ctrl.Log.WithName("controllers").WithName("APIManagement"))
	apimServiceManager := apimservice.NewAzureAPIMgmtServiceManager()
	vnetManager := vnet.NewAzureVNetManager(ctrl.Log.WithName("controllers").WithName("VirtualNetwork"))
	resourceGroupManager := resourcemanagerresourcegroup.NewAzureResourceGroupManager()

	redisCacheManager := resourcemanagerrediscache.NewAzureRedisCacheManager(
		ctrl.Log.WithName("rediscachemanager").WithName("RedisCache"),
		secretClient,
		scheme,
	)
	appInsightsManager := resourcemanagerappinsights.NewManager(
		ctrl.Log.WithName("appinsightsmanager").WithName("AppInsights"),
		secretClient,
		scheme,
	)
	eventhubNamespaceClient := resourcemanagereventhub.NewEventHubNamespaceClient(ctrl.Log.WithName("controllers").WithName("EventhubNamespace"))
	consumerGroupClient := resourcemanagereventhub.NewConsumerGroupClient(ctrl.Log.WithName("controllers").WithName("ConsumerGroup"))
	storageManagers := resourcemanagerstorage.AzureStorageManagers
	keyVaultManager := resourcemanagerkeyvault.NewAzureKeyVaultManager(ctrl.Log.WithName("keyvaultmanager").WithName("KeyVault"), mgr.GetScheme())
	keyVaultKeyManager := &resourcemanagerkeyvault.KeyvaultKeyClient{
		KeyvaultClient: keyVaultManager,
	}
	eventhubClient := resourcemanagereventhub.NewEventhubClient(secretClient, scheme)
	sqlServerManager := resourcemanagersqlserver.NewAzureSqlServerManager(
		ctrl.Log.WithName("sqlservermanager").WithName("AzureSqlServer"),
		secretClient,
		scheme,
	)
	sqlDBManager := resourcemanagersqldb.NewAzureSqlDbManager(ctrl.Log.WithName("sqldbmanager").WithName("AzureSqlDb"))
	sqlFirewallRuleManager := resourcemanagersqlfirewallrule.NewAzureSqlFirewallRuleManager(ctrl.Log.WithName("sqlfirewallrulemanager").WithName("AzureSqlFirewallRule"))
	sqlFailoverGroupManager := resourcemanagersqlfailovergroup.NewAzureSqlFailoverGroupManager(
		ctrl.Log.WithName("sqlfailovergroupmanager").WithName("AzureSqlFailoverGroup"),
		secretClient,
		scheme,
	)
	psqlserverclient := psqlserver.NewPSQLServerClient(ctrl.Log.WithName("psqlservermanager").WithName("PostgreSQLServer"), secretClient, mgr.GetScheme())
	psqldatabaseclient := psqldatabase.NewPSQLDatabaseClient(ctrl.Log.WithName("psqldatabasemanager").WithName("PostgreSQLDatabase"))
	psqlfirewallruleclient := psqlfirewallrule.NewPSQLFirewallRuleClient(ctrl.Log.WithName("psqlfirewallrulemanager").WithName("PostgreSQLFirewallRule"))
	sqlUserManager := resourcemanagersqluser.NewAzureSqlUserManager(
		ctrl.Log.WithName("sqlusermanager").WithName("AzureSqlUser"),
		secretClient,
		scheme,
	)

	err = (&controllers.StorageReconciler{
		Client:         mgr.GetClient(),
		Log:            ctrl.Log.WithName("controllers").WithName("Storage"),
		Recorder:       mgr.GetEventRecorderFor("Storage-controller"),
		Scheme:         mgr.GetScheme(),
		StorageManager: storageManagers.Storage,
	}).SetupWithManager(mgr)
	if err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Storage")
		os.Exit(1)
	}
	err = (&controllers.CosmosDBReconciler{
		Client:   mgr.GetClient(),
		Log:      ctrl.Log.WithName("controllers").WithName("CosmosDB"),
		Recorder: mgr.GetEventRecorderFor("CosmosDB-controller"),
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

	if err = (&controllers.AzureSqlActionReconciler{
		Client:                mgr.GetClient(),
		Log:                   ctrl.Log.WithName("controllers").WithName("AzureSqlAction"),
		Recorder:              mgr.GetEventRecorderFor("AzureSqlAction-controller"),
		Scheme:                mgr.GetScheme(),
		AzureSqlServerManager: sqlServerManager,
		SecretClient:          secretClient,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "AzureSqlAction")
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
		Client:         mgr.GetClient(),
		Log:            ctrl.Log.WithName("controllers").WithName("BlobContainer"),
		Recorder:       mgr.GetEventRecorderFor("BlobContainer-controller"),
		Scheme:         mgr.GetScheme(),
		StorageManager: storageManagers.BlobContainer,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "BlobContainer")
		os.Exit(1)
	}

	if err = (&controllers.AzureDataLakeGen2FileSystemReconciler{
		Client:            mgr.GetClient(),
		Log:               ctrl.Log.WithName("controllers").WithName("AzureDataLakeGen2FileSystem"),
		Recorder:          mgr.GetEventRecorderFor("AzureDataLakeGen2FileSystem-controller"),
		FileSystemManager: storageManagers.FileSystem,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "AzureDataLakeGen2FileSystem")
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

	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
