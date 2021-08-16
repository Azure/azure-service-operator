// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"flag"
	"os"

	aadpodv1 "github.com/Azure/aad-pod-identity/pkg/apis/aadpodidentity/v1"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Azure/azure-service-operator/controllers"
	"github.com/Azure/azure-service-operator/pkg/helpers"

	kscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha2 "github.com/Azure/azure-service-operator/api/v1alpha2"
	azurev1beta1 "github.com/Azure/azure-service-operator/api/v1beta1"
	resourceapimanagement "github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimgmt"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimservice"
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
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagercosmosdbaccount "github.com/Azure/azure-service-operator/pkg/resourcemanager/cosmosdb/account"
	resourcemanagercosmosdbsqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/cosmosdb/sqldatabase"
	resourcemanagereventhub "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	resourcemanagerkeyvault "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/loadbalancer"
	mysqladmin "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/aadadmin"
	mysqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/database"
	mysqlfirewall "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/firewallrule"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/mysqlaaduser"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/mysqluser"
	mysqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/server"
	mysqlvnetrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/vnetrule"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/nic"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/pip"
	psqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/database"
	psqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/firewallrule"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/psqluser"
	psqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/server"
	psqlvnetrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/vnetrule"
	rediscacheactions "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches/actions"
	rcfwr "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches/firewallrule"
	rediscache "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches/redis"
	resourcemanagerresourcegroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	blobContainerManager "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/blobcontainer"
	storageaccountManager "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/vm"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/vmext"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/vmss"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/vnet"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	keyvaultSecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
	k8sSecrets "github.com/Azure/azure-service-operator/pkg/secrets/kube"
	"github.com/Azure/azure-service-operator/pkg/telemetry"
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
	_ = azurev1beta1.AddToScheme(scheme)
	_ = azurev1alpha2.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme

	// We need to query AAD identity types, so add them to the scheme
	aadPodIdentityGroupVersion := schema.GroupVersion{Group: aadpodv1.CRDGroup, Version: aadpodv1.CRDVersion}
	scheme.AddKnownTypes(aadPodIdentityGroupVersion,
		&aadpodv1.AzureIdentity{},
		&aadpodv1.AzureIdentityList{},
		&aadpodv1.AzureIdentityBinding{},
		&aadpodv1.AzureIdentityBindingList{},
		&aadpodv1.AzureAssignedIdentity{},
		&aadpodv1.AzureAssignedIdentityList{},
		&aadpodv1.AzurePodIdentityException{},
		&aadpodv1.AzurePodIdentityExceptionList{})
	metav1.AddToGroupVersion(scheme, aadPodIdentityGroupVersion)
}

// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=azure.microsoft.com,resources=events,verbs=get;list;watch;create;update;patch;delete
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

	ctrl.SetLogger(zap.New(func(o *zap.Options) {
		o.Development = true
	}))

	err := resourcemanagerconfig.ParseEnvironment()
	if err != nil {
		setupLog.Error(err, "unable to parse settings required to provision resources in Azure")
		os.Exit(1)
	}

	setupLog.V(0).Info("Configuration details", "Configuration", resourcemanagerconfig.ConfigString())

	targetNamespaces := resourcemanagerconfig.TargetNamespaces()
	var cacheFunc cache.NewCacheFunc
	if targetNamespaces != nil {
		cacheFunc = cache.MultiNamespacedCacheBuilder(targetNamespaces)
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:               scheme,
		MetricsBindAddress:   metricsAddr,
		NewCache:             cacheFunc,
		LeaderElection:       enableLeaderElection,
		LivenessEndpointName: "/healthz",
		Port:                 9443,
	})

	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	keyvaultName := resourcemanagerconfig.GlobalCredentials().OperatorKeyvault()

	if keyvaultName == "" {
		setupLog.Info("Keyvault name is empty")
		secretClient = k8sSecrets.New(mgr.GetClient(), config.SecretNamingVersion())
	} else {
		setupLog.Info("Instantiating secrets client for keyvault " + keyvaultName)
		secretClient = keyvaultSecrets.New(keyvaultName, config.GlobalCredentials(), config.SecretNamingVersion())
	}

	if config.SelectedMode().IncludesWatchers() {
		if err := registerReconcilers(mgr, secretClient); err != nil {
			setupLog.Error(err, "unable to create controller")
			os.Exit(1)
		}
	}

	if config.SelectedMode().IncludesWebhooks() {
		if err := registerWebhooks(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook")
			os.Exit(1)
		}
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}

}

func registerReconcilers(mgr manager.Manager, secretClient secrets.SecretClient) error {
	// TODO(creds-refactor): construction of these managers will need
	// to move into the AsyncReconciler.Reconcile so that it can use the correct
	// creds based on the namespace of the specific resource being reconciled.
	apimManager := resourceapimanagement.NewManager(config.GlobalCredentials())
	apimServiceManager := apimservice.NewAzureAPIMgmtServiceManager(config.GlobalCredentials())
	vnetManager := vnet.NewAzureVNetManager(config.GlobalCredentials())
	resourceGroupManager := resourcemanagerresourcegroup.NewAzureResourceGroupManager(config.GlobalCredentials())

	redisCacheManager := rediscache.NewAzureRedisCacheManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)
	redisCacheActionManager := rediscacheactions.NewAzureRedisCacheActionManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)

	redisCacheFirewallRuleManager := rcfwr.NewAzureRedisCacheFirewallRuleManager(config.GlobalCredentials())
	appInsightsManager := resourcemanagerappinsights.NewManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)
	eventhubNamespaceClient := resourcemanagereventhub.NewEventHubNamespaceClient(config.GlobalCredentials())
	consumerGroupClient := resourcemanagereventhub.NewConsumerGroupClient(config.GlobalCredentials())
	cosmosDBClient := resourcemanagercosmosdbaccount.NewAzureCosmosDBManager(
		config.GlobalCredentials(),
		secretClient,
	)
	cosmosDBSQLDatabaseClient := resourcemanagercosmosdbsqldatabase.NewAzureCosmosDBSQLDatabaseManager(config.GlobalCredentials())
	keyVaultManager := resourcemanagerkeyvault.NewAzureKeyVaultManager(config.GlobalCredentials(), mgr.GetScheme())
	keyVaultKeyManager := resourcemanagerkeyvault.NewKeyvaultKeyClient(config.GlobalCredentials(), keyVaultManager)
	eventhubClient := resourcemanagereventhub.NewEventhubClient(config.GlobalCredentials(), secretClient, scheme)
	sqlServerManager := resourcemanagersqlserver.NewAzureSqlServerManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)
	sqlDBManager := resourcemanagersqldb.NewAzureSqlDbManager(config.GlobalCredentials())
	sqlFirewallRuleManager := resourcemanagersqlfirewallrule.NewAzureSqlFirewallRuleManager(config.GlobalCredentials())
	sqlVNetRuleManager := resourcemanagersqlvnetrule.NewAzureSqlVNetRuleManager(config.GlobalCredentials())
	sqlFailoverGroupManager := resourcemanagersqlfailovergroup.NewAzureSqlFailoverGroupManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)
	psqlserverclient := psqlserver.NewPSQLServerClient(config.GlobalCredentials(), secretClient, mgr.GetScheme())
	psqldatabaseclient := psqldatabase.NewPSQLDatabaseClient(config.GlobalCredentials())
	psqlfirewallruleclient := psqlfirewallrule.NewPSQLFirewallRuleClient(config.GlobalCredentials())
	psqlusermanager := psqluser.NewPostgreSqlUserManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)
	sqlUserManager := resourcemanagersqluser.NewAzureSqlUserManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)
	sqlManagedUserManager := resourcemanagersqlmanageduser.NewAzureSqlManagedUserManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)
	sqlActionManager := resourcemanagersqlaction.NewAzureSqlActionManager(config.GlobalCredentials(), secretClient, scheme)

	err := (&controllers.StorageAccountReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: storageaccountManager.New(config.GlobalCredentials(), secretClient, scheme),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"StorageAccount",
				ctrl.Log.WithName("controllers").WithName("StorageAccount"),
			),
			Recorder: mgr.GetEventRecorderFor("StorageAccount-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr)
	if err != nil {
		return errors.Wrap(err, "controller StorageAccount")
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
		return errors.Wrap(err, "controller CosmosDB")
	}

	err = (&controllers.CosmosDBSQLDatabaseReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: cosmosDBSQLDatabaseClient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"CosmosDBSQLDatabase",
				ctrl.Log.WithName("controllers").WithName("CosmosDBSQLDatabase"),
			),
			Recorder: mgr.GetEventRecorderFor("CosmosDBSQLDatabase-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr)
	if err != nil {
		return errors.Wrap(err, "controller CosmosDBSQLDatabase")
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
		return errors.Wrap(err, "controller RedisCache")
	}

	if err = (&controllers.RedisCacheActionReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: redisCacheActionManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"RedisCacheAction",
				ctrl.Log.WithName("controllers").WithName("RedisCacheAction"),
			),
			Recorder: mgr.GetEventRecorderFor("RedisCacheAction-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller RedisCacheAction")
	}

	if err = (&controllers.RedisCacheFirewallRuleReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: redisCacheFirewallRuleManager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"RedisCacheFirewallRule",
				ctrl.Log.WithName("controllers").WithName("RedisCacheFirewallRule"),
			),
			Recorder: mgr.GetEventRecorderFor("RedisCacheFirewallRule-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller RedisCacheFirewallRule")
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
		return errors.Wrap(err, "controller Eventhub")
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
		return errors.Wrap(err, "controller ResourceGroup")
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
		return errors.Wrap(err, "controller EventhubNamespace")
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
		return errors.Wrap(err, "controller KeyVault")
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
		return errors.Wrap(err, "controller ConsumerGroup")
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
		return errors.Wrap(err, "controller AzureSqlServer")
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
		return errors.Wrap(err, "controller AzureSqlDb")
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
		return errors.Wrap(err, "controller SqlFirewallRule")
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
		return errors.Wrap(err, "controller SqlVNetRule")
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
		return errors.Wrap(err, "controller SqlAction")
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
		return errors.Wrap(err, "controller AzureSQLUser")
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
		return errors.Wrap(err, "controller AzureSQLManagedUser")
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
		return errors.Wrap(err, "controller AzureSqlFailoverGroup")
	}

	if err = (&controllers.BlobContainerReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: blobContainerManager.New(config.GlobalCredentials()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"BlobContainer",
				ctrl.Log.WithName("controllers").WithName("BlobContainer"),
			),
			Recorder: mgr.GetEventRecorderFor("BlobContainer-controller"),
			Scheme:   mgr.GetScheme(),
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller BlobContainer")
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
		return errors.Wrap(err, "controller AppInsights")
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
		return errors.Wrap(err, "controller PostgreSQLServer")
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
		return errors.Wrap(err, "controller PostgreSQLDatabase")
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
		return errors.Wrap(err, "controller PostgreSQLFirewallRule")
	}

	if err = (&controllers.PostgreSQLUserReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: psqlusermanager,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"PSQLUser",
				ctrl.Log.WithName("controllers").WithName("PostgreSQLUser"),
			),
			Recorder: mgr.GetEventRecorderFor("PostgreSQLUser-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller PostgreSQLUser")
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
		return errors.Wrap(err, "controller ApimService")
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
		return errors.Wrap(err, "controller VNet")
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
		return errors.Wrap(err, "controller APIMgmtAPI")
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
		return errors.Wrap(err, "controller KeyVaultKey")
	}

	if err = (&controllers.MySQLServerReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			AzureClient: mysqlserver.NewMySQLServerClient(
				config.GlobalCredentials(),
				secretClient,
				mgr.GetScheme(),
				mgr.GetClient(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLServer",
				ctrl.Log.WithName("controllers").WithName("MySQLServer"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLServer-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller MySQLServer")
	}
	if err = (&controllers.MySQLDatabaseReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: mysqldatabase.NewMySQLDatabaseClient(config.GlobalCredentials()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLDatabase",
				ctrl.Log.WithName("controllers").WithName("MySQLDatabase"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLDatabase-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller MySQLDatabase")
	}
	if err = (&controllers.MySQLFirewallRuleReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: mysqlfirewall.NewMySQLFirewallRuleClient(config.GlobalCredentials()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLFirewallRule",
				ctrl.Log.WithName("controllers").WithName("MySQLFirewallRule"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLFirewallRule-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller MySQLFirewallRule")
	}

	if err = (&controllers.MySQLUserReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: mysqluser.NewMySqlUserManager(config.GlobalCredentials(), secretClient, scheme),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLUser",
				ctrl.Log.WithName("controllers").WithName("MySQLUser"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLUser-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller MySQLUser")
	}

	// Use the API reader rather than using mgr.GetClient(), because
	// the client might be restricted by target namespaces, while we
	// need to read from the operator namespace.
	identityFinder := helpers.NewAADIdentityFinder(mgr.GetAPIReader(), config.PodNamespace())
	if err = (&controllers.MySQLAADUserReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: mysqlaaduser.NewMySQLAADUserManager(config.GlobalCredentials(), identityFinder),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLAADUser",
				ctrl.Log.WithName("controllers").WithName("MySQLAADUser"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLAADUser-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller MySQLAADUser")
	}

	if err = (&controllers.MySQLServerAdministratorReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: mysqladmin.NewMySQLServerAdministratorManager(config.GlobalCredentials()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLServerAdministrator",
				ctrl.Log.WithName("controllers").WithName("MySQLServerAdministrator"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLServerAdministrator-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller MySQLServerAdministrator")
	}

	if err = (&controllers.AzurePublicIPAddressReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			AzureClient: pip.NewAzurePublicIPAddressClient(
				config.GlobalCredentials(),
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
		return errors.Wrap(err, "controller PublicIPAddress")
	}

	if err = (&controllers.AzureNetworkInterfaceReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			AzureClient: nic.NewAzureNetworkInterfaceClient(
				config.GlobalCredentials(),
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
		return errors.Wrap(err, "controller NetworkInterface")
	}

	if err = (&controllers.MySQLVNetRuleReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: mysqlvnetrule.NewMySQLVNetRuleClient(config.GlobalCredentials()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLVNetRule",
				ctrl.Log.WithName("controllers").WithName("MySQLVNetRule"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLVNetRule-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller MySQLVNetRule")
	}

	if err = (&controllers.AzureVirtualMachineReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			AzureClient: vm.NewAzureVirtualMachineClient(
				config.GlobalCredentials(),
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
		return errors.Wrap(err, "controller VirtualMachine")
	}

	if err = (&controllers.AzureVirtualMachineExtensionReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			AzureClient: vmext.NewAzureVirtualMachineExtensionClient(
				config.GlobalCredentials(),
				secretClient,
				mgr.GetScheme(),
			),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"VirtualMachineExtension",
				ctrl.Log.WithName("controllers").WithName("VirtualMachineExtension"),
			),
			Recorder: mgr.GetEventRecorderFor("VirtualMachineExtension-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller VirtualMachineExtension")
	}

	if err = (&controllers.PostgreSQLVNetRuleReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: psqlvnetrule.NewPostgreSQLVNetRuleClient(config.GlobalCredentials()),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"PostgreSQLVNetRule",
				ctrl.Log.WithName("controllers").WithName("PostgreSQLVNetRule"),
			),
			Recorder: mgr.GetEventRecorderFor("PostgreSQLVNetRule-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller PostgreSQLVNetRule")
	}

	if err = (&controllers.AzureLoadBalancerReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			AzureClient: loadbalancer.NewAzureLoadBalancerClient(
				config.GlobalCredentials(),
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
		return errors.Wrap(err, "controller LoadBalancer")
	}

	if err = (&controllers.AzureVMScaleSetReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			AzureClient: vmss.NewAzureVMScaleSetClient(
				config.GlobalCredentials(),
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
		return errors.Wrap(err, "controller VMScaleSet")
	}

	if err = (&controllers.AppInsightsApiKeyReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: resourcemanagerappinsights.NewAPIKeyClient(config.GlobalCredentials(), secretClient, scheme),
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AppInsightsApiKey",
				ctrl.Log.WithName("controllers").WithName("AppInsightsApiKey"),
			),
			Recorder: mgr.GetEventRecorderFor("AppInsightsApiKey-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		return errors.Wrap(err, "controller AppInsightsApiKey")
	}
	return nil
}

func registerWebhooks(mgr manager.Manager) error {
	if err := (&v1alpha1.AzureSqlServer{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering AzureSqlServer webhook")
	}
	if err := (&azurev1alpha1.AzureSqlDatabase{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering AzureSqlDatabase webhook")
	}
	if err := (&azurev1alpha1.AzureSqlFirewallRule{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering AzureSqlFirewallRule webhook")
	}
	if err := (&azurev1alpha1.AzureSqlFailoverGroup{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering AzureSqlFailoverGroup webhook")
	}
	if err := (&azurev1alpha1.BlobContainer{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering BlobContainer webhook")
	}

	if err := (&azurev1alpha1.MySQLServer{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering MySQLServer webhook")
	}
	if err := (&azurev1alpha1.MySQLUser{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering MySQLUser webhook")
	}
	if err := (&azurev1alpha1.MySQLAADUser{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering MySQLAADUser webhook")
	}
	if err := (&azurev1alpha1.PostgreSQLServer{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering PostgreSQLServer webhook")
	}

	if err := (&azurev1alpha1.AzureSQLUser{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering AzureSQLUser webhook")
	}
	if err := (&azurev1alpha1.AzureSQLManagedUser{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering AzureSQLManagedUser webhook")
	}
	return nil
}
