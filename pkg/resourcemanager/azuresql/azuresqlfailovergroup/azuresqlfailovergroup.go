// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlfailovergroup

import (
	"context"
	"net/http"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	sql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/go-autorest/autorest"
	"github.com/go-logr/logr"
)

type AzureSqlFailoverGroupManager struct {
	Log          logr.Logger
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureSqlFailoverGroupManager(log logr.Logger, secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureSqlFailoverGroupManager {
	return &AzureSqlFailoverGroupManager{
		Log:          log,
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

// GetServer returns a SQL server
func (f *AzureSqlFailoverGroupManager) GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error) {
	serversClient := azuresqlshared.GetGoServersClient()

	return serversClient.Get(
		ctx,
		resourceGroupName,
		serverName,
	)
}

// GetDB retrieves a database
func (f *AzureSqlFailoverGroupManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (sql.Database, error) {
	dbClient := azuresqlshared.GetGoDbClient()

	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
		"serviceTierAdvisors, transparentDataEncryption",
	)
}

// GetFailoverGroup retrieves a failover group
func (f *AzureSqlFailoverGroupManager) GetFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string) (sql.FailoverGroup, error) {
	failoverGroupsClient := azuresqlshared.GetGoFailoverGroupsClient()

	return failoverGroupsClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		failovergroupname,
	)
}

// DeleteFailoverGroup deletes a failover group
func (sdk *AzureSqlFailoverGroupManager) DeleteFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result autorest.Response, err error) {

	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}
	// check to see if the server exists, if it doesn't then short-circuit
	_, err = sdk.GetServer(ctx, resourceGroupName, serverName)
	if err != nil {
		return result, nil
	}
	// check to see if the failover group exists, if it doesn't then short-circuit
	_, err = sdk.GetFailoverGroup(ctx, resourceGroupName, serverName, failoverGroupName)
	if err != nil {
		return result, nil
	}
	failoverGroupsClient := azuresqlshared.GetGoFailoverGroupsClient()
	future, err := failoverGroupsClient.Delete(
		ctx,
		resourceGroupName,
		serverName,
		failoverGroupName,
	)
	if err != nil {
		return result, err
	}

	return future.Result(failoverGroupsClient)
}

// CreateOrUpdateFailoverGroup creates a failover group
func (sdk *AzureSqlFailoverGroupManager) CreateOrUpdateFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string, properties azuresqlshared.SQLFailoverGroupProperties) (result sql.FailoverGroupsCreateOrUpdateFuture, err error) {
	failoverGroupsClient := azuresqlshared.GetGoFailoverGroupsClient()

	// Construct a PartnerInfo object from the server name
	// Get resource ID from the servername to use

	server, err := sdk.GetServer(ctx, properties.SecondaryServerResourceGroup, properties.SecondaryServer)
	if err != nil {
		return result, nil
	}

	secServerResourceID := server.ID
	partnerServerInfo := sql.PartnerInfo{
		ID:              secServerResourceID,
		ReplicationRole: sql.Secondary,
	}

	partnerServerInfoArray := []sql.PartnerInfo{partnerServerInfo}

	var databaseIDArray []string

	// Parse the Databases in the Databaselist and form array of Resource IDs
	for _, each := range properties.DatabaseList {
		database, err := sdk.GetDB(ctx, resourceGroupName, serverName, each)
		if err != nil {
			return result, err
		}
		databaseIDArray = append(databaseIDArray, *database.ID)
	}

	// Construct FailoverGroupProperties struct
	failoverGroupProperties := sql.FailoverGroupProperties{
		ReadWriteEndpoint: &sql.FailoverGroupReadWriteEndpoint{
			FailoverPolicy:                         azuresqlshared.TranslateFailoverPolicy(properties.FailoverPolicy),
			FailoverWithDataLossGracePeriodMinutes: &properties.FailoverGracePeriod,
		},
		PartnerServers: &partnerServerInfoArray,
		Databases:      &databaseIDArray,
	}

	failoverGroup := sql.FailoverGroup{
		FailoverGroupProperties: &failoverGroupProperties,
	}

	return failoverGroupsClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		serverName,
		failovergroupname,
		failoverGroup)
}

func (f *AzureSqlFailoverGroupManager) GetOrPrepareSecret(ctx context.Context, instance *azurev1alpha1.AzureSqlFailoverGroup) (map[string][]byte, error) {
	failovergroupname := instance.ObjectMeta.Name
	azuresqlprimaryserver := instance.Spec.Server
	azuresqlsecondaryserver := instance.Spec.SecondaryServer

	secret := map[string][]byte{}

	key := types.NamespacedName{Name: failovergroupname, Namespace: instance.Namespace}

	if stored, err := f.SecretClient.Get(ctx, key); err == nil {
		f.Log.Info("secret already exists, pulling stored values")
		return stored, nil
	}

	f.Log.Info("secret not found, generating values for new secret")

	secret["azureSqlPrimaryServer"] = []byte(azuresqlprimaryserver)
	secret["readWriteListenerEndpoint"] = []byte(failovergroupname + ".database.windows.net")
	secret["azureSqlSecondaryServer"] = []byte(azuresqlsecondaryserver)
	secret["readOnlyListenerEndpoint"] = []byte(failovergroupname + ".secondary.database.windows.net")

	return secret, nil
}
