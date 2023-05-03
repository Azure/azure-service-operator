// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlfailovergroup

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"github.com/Azure/go-autorest/autorest"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
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
func (m *AzureSqlFailoverGroupManager) GetServer(ctx context.Context, subscriptionID string, resourceGroupName string, serverName string) (result sql.Server, err error) {
	serversClient, err := azuresqlshared.GetGoServersClient(azuresqlshared.GetSubscriptionCredentials(m.Creds, subscriptionID))
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
func (m *AzureSqlFailoverGroupManager) GetDB(
	ctx context.Context,
	subscriptionID string,
	resourceGroupName string,
	serverName string,
	databaseName string,
) (sql.Database, error) {
	dbClient, err := azuresqlshared.GetGoDbClient(azuresqlshared.GetSubscriptionCredentials(m.Creds, subscriptionID))
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

// TODO: Delete this?
// GetFailoverGroup retrieves a failover group
func (m *AzureSqlFailoverGroupManager) GetFailoverGroup(
	ctx context.Context,
	subscriptionID string,
	resourceGroupName string,
	serverName string,
	failovergroupname string,
) (sql.FailoverGroup, error) {
	failoverGroupsClient, err := azuresqlshared.GetGoFailoverGroupsClient(azuresqlshared.GetSubscriptionCredentials(m.Creds, subscriptionID))
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
func (m *AzureSqlFailoverGroupManager) DeleteFailoverGroup(
	ctx context.Context,
	subscriptionID string,
	resourceGroupName string,
	serverName string,
	failoverGroupName string,
) (result autorest.Response, err error) {

	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the server exists, if it doesn't then short-circuit
	_, err = m.GetServer(ctx, subscriptionID, resourceGroupName, serverName)
	if err != nil {
		return result, nil
	}

	failoverGroupsClient, err := azuresqlshared.GetGoFailoverGroupsClient(azuresqlshared.GetSubscriptionCredentials(m.Creds, subscriptionID))
	if err != nil {
		return result, err
	}

	// check to see if the failover group exists, if it doesn't then short-circuit
	_, err = failoverGroupsClient.Get(ctx, resourceGroupName, serverName, failoverGroupName)
	if err != nil {
		return result, nil
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

// TransformToSQLFailoverGroup translates the Kubernetes shaped v1beta1.AzureSqlFailoverGroup into the Azure SDK sql.FailoverGroup.
// This function makes a number of remote calls and so should be called sparingly.
func (m *AzureSqlFailoverGroupManager) TransformToSQLFailoverGroup(ctx context.Context, instance *v1beta1.AzureSqlFailoverGroup) (sql.FailoverGroup, error) {
	secondaryServer, err := m.GetServer(ctx, instance.Spec.SecondaryServerSubscriptionID, instance.Spec.SecondaryServerResourceGroup, instance.Spec.SecondaryServer)
	if err != nil {
		return sql.FailoverGroup{}, err
	}

	secServerResourceID := secondaryServer.ID
	partnerServerInfo := sql.PartnerInfo{
		ID:              secServerResourceID,
		ReplicationRole: sql.Secondary,
	}

	partnerServerInfoArray := []sql.PartnerInfo{partnerServerInfo}

	var databaseIDArray []string

	// Parse the Databases in the Databaselist and form array of Resource IDs
	for _, each := range instance.Spec.DatabaseList {
		database, err := m.GetDB(ctx, instance.Spec.SubscriptionID, instance.Spec.ResourceGroup, instance.Spec.Server, each)
		if err != nil {
			return sql.FailoverGroup{}, err
		}
		databaseIDArray = append(databaseIDArray, *database.ID)
	}

	failoverPolicy, err := azuresqlshared.TranslateFailoverPolicy(instance.Spec.FailoverPolicy)
	if err != nil {
		return sql.FailoverGroup{}, err
	}

	// Construct FailoverGroupProperties struct
	failoverGracePeriod := &instance.Spec.FailoverGracePeriod
	// TODO: This is a bit of a hack right now because the spec doesn't allow us to omit this field, but it
	// TODO: cannot be specified when using Manual mode
	if failoverPolicy == sql.Manual {
		failoverGracePeriod = nil
	}
	failoverGroupProperties := sql.FailoverGroupProperties{
		ReadWriteEndpoint: &sql.FailoverGroupReadWriteEndpoint{
			FailoverPolicy:                         failoverPolicy,
			FailoverWithDataLossGracePeriodMinutes: failoverGracePeriod,
		},
		PartnerServers: &partnerServerInfoArray,
		Databases:      &databaseIDArray,
	}

	failoverGroup := sql.FailoverGroup{
		FailoverGroupProperties: &failoverGroupProperties,
	}

	return failoverGroup, nil
}

// CreateOrUpdateFailoverGroup creates a failover group
func (m *AzureSqlFailoverGroupManager) CreateOrUpdateFailoverGroup(
	ctx context.Context,
	subscriptionID string,
	resourceGroup string,
	server string,
	failoverGroupName string,
	failoverGroupProperties sql.FailoverGroup) (sql.FailoverGroupsCreateOrUpdateFuture, error) {

	failoverGroupsClient, err := azuresqlshared.GetGoFailoverGroupsClient(azuresqlshared.GetSubscriptionCredentials(m.Creds, subscriptionID))
	if err != nil {
		return sql.FailoverGroupsCreateOrUpdateFuture{}, err
	}

	return failoverGroupsClient.CreateOrUpdate(
		ctx,
		resourceGroup,
		server,
		failoverGroupName,
		failoverGroupProperties)
}

func (m *AzureSqlFailoverGroupManager) NewSecret(instance *v1beta1.AzureSqlFailoverGroup) map[string][]byte {
	failoverGroupName := instance.ObjectMeta.Name
	azureSQLPrimaryServer := instance.Spec.Server
	azureSQLSecondaryServer := instance.Spec.SecondaryServer

	secret := map[string][]byte{}

	// TODO: In a future version we should consider moving these values to properties on status or something rather than
	// TODO: indirecting them through KeyVault, since really none of these values are actually secrets
	secret["azureSqlPrimaryServer"] = []byte(azureSQLPrimaryServer)
	secret["readWriteListenerEndpoint"] = []byte(failoverGroupName + "." + config.Environment().SQLDatabaseDNSSuffix)
	secret["azureSqlSecondaryServer"] = []byte(azureSQLSecondaryServer)
	secret["readOnlyListenerEndpoint"] = []byte(failoverGroupName + ".secondary." + config.Environment().SQLDatabaseDNSSuffix)

	return secret
}

func doReadWriteEndpointsMatch(expected *sql.FailoverGroupReadWriteEndpoint, actual *sql.FailoverGroupReadWriteEndpoint) bool {
	if expected == nil && actual == nil {
		return true
	}

	if (expected == nil) != (actual == nil) {
		return false
	}

	if expected.FailoverPolicy != actual.FailoverPolicy {
		return false
	}

	if expected.FailoverWithDataLossGracePeriodMinutes == nil && actual.FailoverWithDataLossGracePeriodMinutes != nil ||
		expected.FailoverWithDataLossGracePeriodMinutes != nil && actual.FailoverWithDataLossGracePeriodMinutes == nil {
		return false
	}

	if expected.FailoverWithDataLossGracePeriodMinutes == nil && actual.FailoverWithDataLossGracePeriodMinutes == nil {
		return true
	}

	return *expected.FailoverWithDataLossGracePeriodMinutes == *actual.FailoverWithDataLossGracePeriodMinutes
}

func doDatabasesMatch(expectedDatabases *[]string, actualDatabases *[]string) bool {
	if (expectedDatabases == nil) != (actualDatabases == nil) {
		return false
	}
	if expectedDatabases == nil && actualDatabases == nil {
		return true
	}
	expected := *expectedDatabases
	actual := *actualDatabases

	if len(expected) != len(actual) {
		return false
	}
	for i, v1 := range expected {
		v2 := actual[i]
		if v1 != v2 {
			return false
		}
	}

	return true
}

func doPartnerServersMatch(expectedPartnerServers *[]sql.PartnerInfo, actualPartnerServers *[]sql.PartnerInfo) bool {
	if (expectedPartnerServers == nil) != (actualPartnerServers == nil) {
		return false
	}
	if expectedPartnerServers != nil && actualPartnerServers != nil {
		if len(*expectedPartnerServers) != len(*actualPartnerServers) {
			return false
		}
		for i, v1 := range *expectedPartnerServers {
			v2 := (*actualPartnerServers)[i]
			if v1.ID == nil && v2.ID != nil ||
				v1.ID != nil && v2.ID == nil {
				return false
			}
			if v1.ID != nil && v2.ID != nil {
				if *v1.ID != *v2.ID {
					return false
				}
			}

		}
	}

	return true
}

func DoesResourceMatchAzure(expected sql.FailoverGroup, actual sql.FailoverGroup) bool {
	if len(expected.Tags) != len(actual.Tags) {
		return false
	}
	for k, v := range expected.Tags {
		if v != actual.Tags[k] {
			return false
		}
	}

	if (expected.FailoverGroupProperties == nil) != (actual.FailoverGroupProperties == nil) {
		return false
	}
	if expected.FailoverGroupProperties != nil && actual.FailoverGroupProperties != nil {
		// We care about ReadWriteEndpoint, PartnerServers, and Databases
		if !doReadWriteEndpointsMatch(expected.FailoverGroupProperties.ReadWriteEndpoint, actual.FailoverGroupProperties.ReadWriteEndpoint) {
			return false
		}
		if !doDatabasesMatch(expected.FailoverGroupProperties.Databases, actual.FailoverGroupProperties.Databases) {
			return false
		}
		if !doPartnerServersMatch(expected.FailoverGroupProperties.PartnerServers, actual.FailoverGroupProperties.PartnerServers) {
			return false
		}
	}

	return true
}
