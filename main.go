// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"flag"
	"os"

	aadpodv1 "github.com/Azure/aad-pod-identity/pkg/apis/aadpodidentity/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha2 "github.com/Azure/azure-service-operator/api/v1alpha2"
	azurev1beta1 "github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/controllers"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
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
	mysqladmin "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/aadadmin"
	mysqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/database"
	mysqlfirewall "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/firewallrule"
	mysqlaaduser "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/mysqlaaduser"
	mysqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/mysqluser"
	mysqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/server"
	mysqlvnetrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/vnetrule"
	nic "github.com/Azure/azure-service-operator/pkg/resourcemanager/nic"
	pip "github.com/Azure/azure-service-operator/pkg/resourcemanager/pip"
	psqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/database"
	psqlfirewallrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/firewallrule"
	psqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/psqluser"
	psqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/server"
	psqlvnetrule "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/vnetrule"
	rediscacheactions "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches/actions"
	rcfwr "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches/firewallrule"
	rediscache "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches/redis"
	resourcemanagerresourcegroup "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	blobContainerManager "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/blobcontainer"
	storageaccountManager "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"
	vm "github.com/Azure/azure-service-operator/pkg/resourcemanager/vm"
	vmext "github.com/Azure/azure-service-operator/pkg/resourcemanager/vmext"
	vmss "github.com/Azure/azure-service-operator/pkg/resourcemanager/vmss"
	vnet "github.com/Azure/azure-service-operator/pkg/resourcemanager/vnet"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
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
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")

	flag.Parse()

	ctrl.SetLogger(zap.Logger(true))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:               scheme,
		MetricsBindAddress:   metricsAddr,
		LeaderElection:       enableLeaderElection,
		LivenessEndpointName: "/healthz",
		Port:                 9443,
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

	setupLog.V(0).Info("Configuration details", "Configuration", resourcemanagerconfig.ConfigString())

	err = (&controllers.StorageAccountReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:     mgr.GetClient(),
			ARMFactory: storageaccountManager.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagercosmosdb.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: rediscache.NewARMClient,
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

	if err = (&controllers.RedisCacheActionReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:     mgr.GetClient(),
			ARMFactory: rediscacheactions.NewARMClient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"RedisCacheAction",
				ctrl.Log.WithName("controllers").WithName("RedisCacheAction"),
			),
			Recorder: mgr.GetEventRecorderFor("RedisCacheAction-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "RedisCacheAction")
		os.Exit(1)
	}

	if err = (&controllers.RedisCacheFirewallRuleReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:     mgr.GetClient(),
			ARMFactory: rcfwr.NewARMClient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"RedisCacheFirewallRule",
				ctrl.Log.WithName("controllers").WithName("RedisCacheFirewallRule"),
			),
			Recorder: mgr.GetEventRecorderFor("RedisCacheFirewallRule-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "RedisCacheFirewallRule")
		os.Exit(1)
	}

	err = (&controllers.EventhubReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagereventhub.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagerresourcegroup.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagereventhub.NewNamespaceARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagerkeyvault.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagereventhub.NewConsumerGroupARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagersqlserver.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagersqldb.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagersqlfirewallrule.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagersqlvnetrule.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagersqlaction.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagersqluser.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagersqlmanageduser.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagersqlfailovergroup.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: blobContainerManager.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagerappinsights.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: psqlserver.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: psqldatabase.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: psqlfirewallrule.NewARMClient,
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

	if err = (&controllers.PostgreSQLUserReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:     mgr.GetClient(),
			ARMFactory: psqluser.NewARMClient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"PSQLUser",
				ctrl.Log.WithName("controllers").WithName("PostgreSQLUser"),
			),
			Recorder: mgr.GetEventRecorderFor("PostgreSQLUser-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "PostgreSQLUser")
		os.Exit(1)
	}

	if err = (&controllers.ApimServiceReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:     mgr.GetClient(),
			ARMFactory: apimservice.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: vnet.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourceapimanagement.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagerkeyvault.NewKeyARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: mysqlserver.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: mysqldatabase.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: mysqlfirewall.NewARMClient,
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

	if err = (&controllers.MySQLUserReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:     mgr.GetClient(),
			ARMFactory: mysqluser.NewARMClient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLUser",
				ctrl.Log.WithName("controllers").WithName("MySQLUser"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLUser-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "MySQLUser")
		os.Exit(1)
	}

	identityFinder := helpers.NewAADIdentityFinder(mgr.GetClient(), resourcemanagerconfig.PodNamespace())
	if err = (&controllers.MySQLAADUserReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client: mgr.GetClient(),
			ARMFactory: func(creds resourcemanagerconfig.Credentials, _ secrets.SecretClient, _ *runtime.Scheme) resourcemanager.ARMClient {
				return mysqlaaduser.NewMySQLAADUserManager(creds, identityFinder)
			},
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLAADUser",
				ctrl.Log.WithName("controllers").WithName("MySQLAADUser"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLAADUser-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "MySQLAADUser")
		os.Exit(1)
	}

	if err = (&controllers.MySQLServerAdministratorReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:     mgr.GetClient(),
			ARMFactory: mysqladmin.NewARMClient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"MySQLServerAdministrator",
				ctrl.Log.WithName("controllers").WithName("MySQLServerAdministrator"),
			),
			Recorder: mgr.GetEventRecorderFor("MySQLServerAdministrator-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "MySQLServerAdministrator")
		os.Exit(1)
	}

	if err = (&controllers.AzurePublicIPAddressReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:     mgr.GetClient(),
			ARMFactory: pip.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: nic.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: mysqlvnetrule.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: vm.NewARMClient,
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

	if err = (&controllers.AzureVirtualMachineExtensionReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:     mgr.GetClient(),
			ARMFactory: vmext.NewARMClient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"VirtualMachineExtension",
				ctrl.Log.WithName("controllers").WithName("VirtualMachineExtension"),
			),
			Recorder: mgr.GetEventRecorderFor("VirtualMachineExtension-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "VirtualMachineExtension")
		os.Exit(1)
	}

	if err = (&controllers.PostgreSQLVNetRuleReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:     mgr.GetClient(),
			ARMFactory: psqlvnetrule.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: loadbalancer.NewARMClient,
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
			Client:     mgr.GetClient(),
			ARMFactory: vmss.NewARMClient,
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

	if err = (&controllers.AppInsightsApiKeyReconciler{
		Reconciler: &controllers.AsyncReconciler{
			Client:     mgr.GetClient(),
			ARMFactory: resourcemanagerappinsights.NewAPIKeyARMClient,
			Telemetry: telemetry.InitializeTelemetryDefault(
				"AppInsightsApiKey",
				ctrl.Log.WithName("controllers").WithName("AppInsightsApiKey"),
			),
			Recorder: mgr.GetEventRecorderFor("AppInsightsApiKey-controller"),
			Scheme:   scheme,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "AppInsightsApiKey")
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
	if err = (&azurev1alpha1.MySQLUser{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "MySQLUser")
		os.Exit(1)
	}
	if err = (&azurev1alpha1.MySQLAADUser{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "MySQLAADUser")
		os.Exit(1)
	}
	if err = (&azurev1alpha1.PostgreSQLServer{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "PostgreSQLServer")
		os.Exit(1)
	}

	if err = (&azurev1alpha1.AzureSQLUser{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "AzureSQLUser")
		os.Exit(1)
	}
	if err = (&azurev1alpha1.AzureSQLManagedUser{}).SetupWebhookWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create webhook", "webhook", "AzureSQLManagedUser")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}

}
