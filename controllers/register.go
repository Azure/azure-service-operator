// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimgmt"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimservice"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/appinsights"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlaction"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqldb"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfailovergroup"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlfirewallrule"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlmanageduser"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlserver"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlvnetrule"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	cosmosdbaccount "github.com/Azure/azure-service-operator/pkg/resourcemanager/cosmosdb/account"
	cosmosdbsqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/cosmosdb/sqldatabase"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
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
	redisactions "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches/actions"
	rcfwr "github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches/firewallrule"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/rediscaches/redis"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/blobcontainer"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/vm"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/vmext"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/vmss"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/vnet"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/azure-service-operator/pkg/telemetry"
)

func RegisterReconcilers(mgr manager.Manager, scheme *runtime.Scheme, secretClient secrets.SecretClient) error {
	// TODO(creds-refactor): construction of these managers will need
	// to move into the AsyncReconciler.Reconcile so that it can use the correct
	// creds based on the namespace of the specific resource being reconciled.
	apimManager := apimgmt.NewManager(config.GlobalCredentials())
	apimServiceManager := apimservice.NewAzureAPIMgmtServiceManager(config.GlobalCredentials())
	vnetManager := vnet.NewAzureVNetManager(config.GlobalCredentials())
	resourceGroupManager := resourcegroups.NewAzureResourceGroupManager(config.GlobalCredentials())

	redisCacheManager := redis.NewAzureRedisCacheManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)
	redisCacheActionManager := redisactions.NewAzureRedisCacheActionManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)

	redisCacheFirewallRuleManager := rcfwr.NewAzureRedisCacheFirewallRuleManager(config.GlobalCredentials())
	appInsightsManager := appinsights.NewManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)
	eventhubNamespaceClient := eventhubs.NewEventHubNamespaceClient(config.GlobalCredentials())
	consumerGroupClient := eventhubs.NewConsumerGroupClient(config.GlobalCredentials())
	cosmosDBClient := cosmosdbaccount.NewAzureCosmosDBManager(
		config.GlobalCredentials(),
		secretClient,
	)
	cosmosDBSQLDatabaseClient := cosmosdbsqldatabase.NewAzureCosmosDBSQLDatabaseManager(config.GlobalCredentials())
	keyVaultManager := keyvaults.NewAzureKeyVaultManager(config.GlobalCredentials(), mgr.GetScheme())
	keyVaultKeyManager := keyvaults.NewKeyvaultKeyClient(config.GlobalCredentials(), keyVaultManager)
	eventhubClient := eventhubs.NewEventhubClient(config.GlobalCredentials(), secretClient, scheme)
	sqlServerManager := azuresqlserver.NewAzureSqlServerManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)
	sqlDBManager := azuresqldb.NewAzureSqlDbManager(config.GlobalCredentials())
	sqlFirewallRuleManager := azuresqlfirewallrule.NewAzureSqlFirewallRuleManager(config.GlobalCredentials())
	sqlVNetRuleManager := azuresqlvnetrule.NewAzureSqlVNetRuleManager(config.GlobalCredentials())
	sqlFailoverGroupManager := azuresqlfailovergroup.NewAzureSqlFailoverGroupManager(
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
	sqlUserManager := azuresqluser.NewAzureSqlUserManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)
	sqlManagedUserManager := azuresqlmanageduser.NewAzureSqlManagedUserManager(
		config.GlobalCredentials(),
		secretClient,
		scheme,
	)
	sqlActionManager := azuresqlaction.NewAzureSqlActionManager(config.GlobalCredentials(), secretClient, scheme)

	err := (&StorageAccountReconciler{
		Reconciler: &AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: storageaccount.NewManager(config.GlobalCredentials(), secretClient, scheme),
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
	err = (&CosmosDBReconciler{
		Reconciler: &AsyncReconciler{
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

	err = (&CosmosDBSQLDatabaseReconciler{
		Reconciler: &AsyncReconciler{
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

	err = (&RedisCacheReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&RedisCacheActionReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&RedisCacheFirewallRuleReconciler{
		Reconciler: &AsyncReconciler{
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

	err = (&EventhubReconciler{
		Reconciler: &AsyncReconciler{
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

	err = (&ResourceGroupReconciler{
		Reconciler: &AsyncReconciler{
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

	err = (&EventhubNamespaceReconciler{
		Reconciler: &AsyncReconciler{
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

	err = (&KeyVaultReconciler{
		Reconciler: &AsyncReconciler{
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

	err = (&ConsumerGroupReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AzureSqlServerReconciler{
		Reconciler: &AsyncReconciler{
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
	err = (&AzureSqlDatabaseReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AzureSqlFirewallRuleReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AzureSQLVNetRuleReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AzureSqlActionReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AzureSQLUserReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AzureSQLManagedUserReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AzureSqlFailoverGroupReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&BlobContainerReconciler{
		Reconciler: &AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: blobcontainer.NewManager(config.GlobalCredentials()),
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

	if err = (&AppInsightsReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&PostgreSQLServerReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&PostgreSQLDatabaseReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&PostgreSQLFirewallRuleReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&PostgreSQLUserReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&ApimServiceReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&VirtualNetworkReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&APIMAPIReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&KeyVaultKeyReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&MySQLServerReconciler{
		Reconciler: &AsyncReconciler{
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
	if err = (&MySQLDatabaseReconciler{
		Reconciler: &AsyncReconciler{
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
	if err = (&MySQLFirewallRuleReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&MySQLUserReconciler{
		Reconciler: &AsyncReconciler{
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
	if err = (&MySQLAADUserReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&MySQLServerAdministratorReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AzurePublicIPAddressReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AzureNetworkInterfaceReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&MySQLVNetRuleReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AzureVirtualMachineReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AzureVirtualMachineExtensionReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&PostgreSQLVNetRuleReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AzureLoadBalancerReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AzureVMScaleSetReconciler{
		Reconciler: &AsyncReconciler{
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

	if err = (&AppInsightsApiKeyReconciler{
		Reconciler: &AsyncReconciler{
			Client:      mgr.GetClient(),
			AzureClient: appinsights.NewAPIKeyClient(config.GlobalCredentials(), secretClient, scheme),
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

func RegisterWebhooks(mgr manager.Manager) error {
	if err := (&v1alpha1.AzureSqlServer{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering AzureSqlServer webhook")
	}
	if err := (&v1alpha1.AzureSqlDatabase{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering AzureSqlDatabase webhook")
	}
	if err := (&v1alpha1.AzureSqlFirewallRule{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering AzureSqlFirewallRule webhook")
	}
	if err := (&v1alpha1.AzureSqlFailoverGroup{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering AzureSqlFailoverGroup webhook")
	}
	if err := (&v1alpha1.BlobContainer{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering BlobContainer webhook")
	}

	if err := (&v1alpha1.MySQLServer{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering MySQLServer webhook")
	}
	if err := (&v1alpha1.MySQLUser{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering v1alpha1 MySQLUser webhook")
	}
	if err := (&v1alpha2.MySQLUser{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering v1alpha2 MySQLUser webhook")
	}
	if err := (&v1alpha1.MySQLAADUser{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering v1alpha1 MySQLAADUser webhook")
	}
	if err := (&v1alpha2.MySQLAADUser{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering v1alpha2 MySQLAADUser webhook")
	}
	if err := (&v1alpha1.PostgreSQLServer{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering PostgreSQLServer webhook")
	}

	if err := (&v1alpha1.AzureSQLUser{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering AzureSQLUser webhook")
	}
	if err := (&v1alpha1.AzureSQLManagedUser{}).SetupWebhookWithManager(mgr); err != nil {
		return errors.Wrap(err, "registering AzureSQLManagedUser webhook")
	}
	return nil
}
