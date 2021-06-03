// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package sqldatabase

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2021-03-15/documentdb"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/cosmosdb"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/pollclient"
)

func NewAzureCosmosDBSQLDatabaseManager(creds config.Credentials) *AzureCosmosDBSQLDatabaseManager {
	return &AzureCosmosDBSQLDatabaseManager{
		Creds: creds,
	}
}

var _ resourcemanager.ARMClient = &AzureCosmosDBSQLDatabaseManager{}

type AzureCosmosDBSQLDatabaseManager struct {
	Creds config.Credentials
}

/*
 * ARMClient implementation
 */

func (m *AzureCosmosDBSQLDatabaseManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	cosmosDBSQLDatabaseClient, err := cosmosdb.GetCosmosDBSQLDatabaseClient(m.Creds)
	if err != nil {
		return false, errors.Wrapf(err, "failed to create cosmos DB SQL database client")
	}

	pClient := pollclient.NewPollClient(m.Creds)
	lroPollResult, err := pClient.PollLongRunningOperationIfNeededV1Alpha1(ctx, &instance.Status)
	if err != nil {
		instance.Status.Message = err.Error()
		return false, err
	}

	if lroPollResult == pollclient.PollResultTryAgainLater {
		// Need to wait a bit before trying again
		return false, nil
	}
	if lroPollResult == pollclient.PollResultBadRequest {
		// Reached a terminal state
		instance.Status.SetFailedProvisioning(instance.Status.Message)
		return true, nil
	}

	notFoundErrorCodes := []string{
		errhelp.ParentNotFoundErrorCode,
		errhelp.ResourceGroupNotFoundErrorCode,
		errhelp.NotFoundErrorCode,
		errhelp.ResourceNotFound,
	}

	// Get the existing resource
	sqlDatabase, err := cosmosDBSQLDatabaseClient.GetSQLDatabase(
		ctx,
		instance.Spec.ResourceGroup,
		instance.Spec.CosmosDBAccount,
		instance.ObjectMeta.Name)
	if err != nil {
		azerr := errhelp.NewAzureError(err)
		if azerr.Type != errhelp.ResourceNotFound && azerr.Type != errhelp.NotFoundErrorCode {
			instance.Status.Message = fmt.Sprintf("failed GET-ing CosmosDBSQLDatabase: %s", err.Error())
			return errhelp.IsErrorFatal(err, notFoundErrorCodes, nil)
		}
	}

	// TODO: There's a bug in Cosmos DB SQL Database where the throughput settings configured on
	// TODO: the DB don't actually get shown on the DB, instead they are shown on the throughput child resource.
	// TODO: We monitor that resource too so we can diff against it
	sqlDatabaseThroughputSettings, err := cosmosDBSQLDatabaseClient.GetSQLDatabaseThroughput(
		ctx,
		instance.Spec.ResourceGroup,
		instance.Spec.CosmosDBAccount,
		instance.ObjectMeta.Name)
	if err != nil {
		azerr := errhelp.NewAzureError(err)
		if azerr.Type != errhelp.ResourceNotFound && azerr.Type != errhelp.NotFoundErrorCode {
			instance.Status.Message = fmt.Sprintf("failed GET-ing CosmosDBSQLDatabaseThroughput: %s", err.Error())
			return errhelp.IsErrorFatal(err, notFoundErrorCodes, nil)
		}
	}

	// Transform our spec to Azure shape
	transformed := m.transformToCosmosDBSQLDatabase(instance)

	// Check if we're in the goal state already
	resourceMatchesAzure := DoesResourceMatchAzure(transformed, sqlDatabase, sqlDatabaseThroughputSettings)
	if resourceMatchesAzure {
		instance.Status.SetProvisioned(resourcemanager.SuccessMsg)
		instance.Status.ResourceId = *sqlDatabase.ID
		return true, nil
	}

	// Drive to goal state
	instance.Status.SetProvisioning("")
	future, err := cosmosDBSQLDatabaseClient.CreateUpdateSQLDatabase(
		ctx,
		instance.Spec.ResourceGroup,
		instance.Spec.CosmosDBAccount,
		instance.ObjectMeta.Name,
		transformed)
	if err != nil {
		instance.Status.Message = err.Error()
		allowedErrors := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.AlreadyExists,
			errhelp.FailoverGroupBusy,
		}
		unrecoverableErrors := []string{
			errhelp.InvalidFailoverGroupRegion,
		}
		return errhelp.IsErrorFatal(err, append(allowedErrors, notFoundErrorCodes...), unrecoverableErrors)
	}
	instance.Status.SetProvisioning("Resource request successfully submitted to Azure")
	instance.Status.PollingURL = future.PollingURL()

	// Need to poll the polling URL, so not done yet!
	return false, nil
}

func (m *AzureCosmosDBSQLDatabaseManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	cosmosDBSQLDatabaseClient, err := cosmosdb.GetCosmosDBSQLDatabaseClient(m.Creds)
	if err != nil {
		return false, errors.Wrapf(err, "failed to create cosmos DB SQL database client")
	}

	_, err = cosmosDBSQLDatabaseClient.DeleteSQLDatabase(ctx, instance.Spec.ResourceGroup, instance.Spec.CosmosDBAccount, instance.ObjectMeta.Name)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureError(err)

		// these errors are expected
		ignore := []string{
			errhelp.AsyncOpIncompleteError,
		}

		// this means the thing doesn't exist
		finished := []string{
			errhelp.ResourceNotFound,
			errhelp.ResourceGroupNotFoundErrorCode,
		}

		if helpers.ContainsString(ignore, azerr.Type) {
			return true, nil
		}

		if helpers.ContainsString(finished, azerr.Type) {
			return false, nil
		}
		instance.Status.Message = fmt.Sprintf("delete failed with: %s", err.Error())
		return false, err
	}

	// Not done yet, keep trying
	return false, nil
}

func (m *AzureCosmosDBSQLDatabaseManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.CosmosDBAccount,
			},
			Target: &v1alpha1.CosmosDB{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &v1alpha1.ResourceGroup{},
		},
	}, nil
}

func (m *AzureCosmosDBSQLDatabaseManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}

	return &instance.Status, nil
}

/*
 * Helper functions
 */

func (m *AzureCosmosDBSQLDatabaseManager) convert(obj runtime.Object) (*v1alpha1.CosmosDBSQLDatabase, error) {
	db, ok := obj.(*v1alpha1.CosmosDBSQLDatabase)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return db, nil
}

func makeTags(tags map[string]string) map[string]*string {
	result := make(map[string]*string)
	for key, value := range tags {
		result[key] = &value
	}

	return result
}

func (m *AzureCosmosDBSQLDatabaseManager) transformToCosmosDBSQLDatabase(instance *v1alpha1.CosmosDBSQLDatabase) documentdb.SQLDatabaseCreateUpdateParameters {
	var autoscaleSettings *documentdb.AutoscaleSettings
	if instance.Spec.AutoscaleSettings != nil && instance.Spec.AutoscaleSettings.MaxThroughput != nil {
		autoscaleSettings = &documentdb.AutoscaleSettings{
			MaxThroughput: instance.Spec.AutoscaleSettings.MaxThroughput,
		}
	}

	var createUpdateOptions *documentdb.CreateUpdateOptions
	if instance.Spec.Throughput != nil || autoscaleSettings != nil {
		createUpdateOptions = &documentdb.CreateUpdateOptions{
			Throughput:        instance.Spec.Throughput,
			AutoscaleSettings: autoscaleSettings,
		}
	}

	return documentdb.SQLDatabaseCreateUpdateParameters{
		//Location: &instance.Spec.Location, // TODO: Is this really optional? TODO: if this passes with this commented out, maybe just remove it?
		Tags: makeTags(instance.Spec.Tags),
		SQLDatabaseCreateUpdateProperties: &documentdb.SQLDatabaseCreateUpdateProperties{
			Resource: &documentdb.SQLDatabaseResource{
				ID: &instance.ObjectMeta.Name,
			},
			Options: createUpdateOptions,
		},
	}
}
