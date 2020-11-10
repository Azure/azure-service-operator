// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlfailovergroup

import (
	"context"
	"net/http"

	"github.com/Azure/azure-service-operator/api/v1beta1"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	sql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"github.com/Azure/go-autorest/autorest"
)

type AzureSqlFailoverGroupManager struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureSqlFailoverGroupManager(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureSqlFailoverGroupManager {
	return &AzureSqlFailoverGroupManager{
		Creds:        creds,
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

// GetServer returns a SQL server
func (m *AzureSqlFailoverGroupManager) GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error) {
	serversClient, err := azuresqlshared.GetGoServersClient(m.Creds)
	if err != nil {
		return sql.Server{}, err
	}

	return serversClient.Get(
		ctx,
		resourceGroupName,
		serverName,
	)
}

// GetDB retrieves a database
func (m *AzureSqlFailoverGroupManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (sql.Database, error) {
	dbClient, err := azuresqlshared.GetGoDbClient(m.Creds)
	if err != nil {
		return sql.Database{}, err
	}

	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
	)
}

// GetFailoverGroup retrieves a failover group
func (m *AzureSqlFailoverGroupManager) GetFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string) (sql.FailoverGroup, error) {
	failoverGroupsClient, err := azuresqlshared.GetGoFailoverGroupsClient(m.Creds)
	if err != nil {
		return sql.FailoverGroup{}, err
	}

	return failoverGroupsClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		failovergroupname,
	)
}

// DeleteFailoverGroup deletes a failover group
func (m *AzureSqlFailoverGroupManager) DeleteFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result autorest.Response, err error) {
	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the server exists, if it doesn't then short-circuit
	_, err = m.GetServer(ctx, resourceGroupName, serverName)
	if err != nil {
		return result, nil
	}
	// check to see if the failover group exists, if it doesn't then short-circuit
	_, err = m.GetFailoverGroup(ctx, resourceGroupName, serverName, failoverGroupName)
	if err != nil {
		return result, nil
	}

	failoverGroupsClient, err := azuresqlshared.GetGoFailoverGroupsClient(m.Creds)
	if err != nil {
		return result, err
	}

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
func (m *AzureSqlFailoverGroupManager) CreateOrUpdateFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string, properties azuresqlshared.SQLFailoverGroupProperties) (result sql.FailoverGroupsCreateOrUpdateFuture, err error) {
	failoverGroupsClient, err := azuresqlshared.GetGoFailoverGroupsClient(m.Creds)
	if err != nil {
		return sql.FailoverGroupsCreateOrUpdateFuture{}, err
	}

	// Construct a PartnerInfo object from the server name
	// Get resource ID from the servername to use

	server, err := m.GetServer(ctx, properties.SecondaryServerResourceGroup, properties.SecondaryServer)
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
		database, err := m.GetDB(ctx, resourceGroupName, serverName, each)
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

func (m *AzureSqlFailoverGroupManager) GetOrPrepareSecret(ctx context.Context, instance *v1beta1.AzureSqlFailoverGroup) (map[string][]byte, error) {
	failovergroupname := instance.ObjectMeta.Name
	azuresqlprimaryserver := instance.Spec.Server
	azuresqlsecondaryserver := instance.Spec.SecondaryServer

	secret := map[string][]byte{}

	key := types.NamespacedName{Name: failovergroupname, Namespace: instance.Namespace}

	if stored, err := m.SecretClient.Get(ctx, key); err == nil {
		return stored, nil
	}

	secret["azureSqlPrimaryServer"] = []byte(azuresqlprimaryserver)
	secret["readWriteListenerEndpoint"] = []byte(failovergroupname + "." + config.Environment().SQLDatabaseDNSSuffix)
	secret["azureSqlSecondaryServer"] = []byte(azuresqlsecondaryserver)
	secret["readOnlyListenerEndpoint"] = []byte(failovergroupname + ".secondary." + config.Environment().SQLDatabaseDNSSuffix)

	return secret, nil
}
