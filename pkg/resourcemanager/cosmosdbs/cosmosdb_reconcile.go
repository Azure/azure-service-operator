// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package cosmosdbs

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/pollclient"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure ensures that cosmosdb is provisioned as specified
func (m *AzureCosmosDBManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		m.SecretClient = options.SecretClient
	}

	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	hash := helpers.Hash256(instance.Spec)

	if instance.Status.SpecHash == hash && instance.Status.Provisioned {
		instance.Status.RequestedAt = nil
		return true, nil
	}
	instance.Status.Provisioned = false

	if instance.Status.PollingURL != "" {
		pollClient := pollclient.NewPollClient()
		pollResponse, err := pollClient.Get(ctx, instance.Status.PollingURL)
		if err != nil {
			instance.Status.Provisioning = false
			return false, err
		}

		// response is not ready yet
		if pollResponse.Status == "Dequeued" {
			return false, nil
		}

		instance.Status.PollingURL = ""

		if pollResponse.Status == "Failed" {
			instance.Status.Provisioning = false
			instance.Status.Message = pollResponse.Error.Error()
			return true, nil
		}
	}

	// get the instance and update status
	db, err := m.GetCosmosDB(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)

		instance.Status.Message = err.Error()

		switch azerr.Type {
		case errhelp.ResourceGroupNotFoundErrorCode, errhelp.ParentNotFoundErrorCode:
			instance.Status.Provisioning = false
			instance.Status.State = "Waiting"
			return false, nil
		}

	} else {
		instance.Status.ResourceId = *db.ID
		instance.Status.State = *db.ProvisioningState
	}

	if instance.Status.State == "Creating" || instance.Status.State == "Updating" {
		// avoid multiple CreateOrUpdate requests while resource is already creating
		return false, nil
	}

	if instance.Status.State == "Succeeded" {
		// provisioning is complete, update the secrets
		if err = m.createOrUpdateSecret(ctx, instance, db); err != nil {
			instance.Status.Message = err.Error()
			return false, err
		}

		if instance.Status.SpecHash == hash {
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.Provisioning = false
			instance.Status.Provisioned = true
			return true, nil
		}
	}

	if instance.Status.State == "Failed" {
		instance.Status.Message = "Failed to provision CosmosDB"
		instance.Status.Provisioning = false
		instance.Status.Provisioned = false
		return true, nil
	}

	instance.Status.Provisioning = true

	tags := helpers.LabelsToTags(instance.GetLabels())
	accountName := instance.ObjectMeta.Name
	db, pollingUrl, err := m.CreateOrUpdateCosmosDB(ctx, accountName, instance.Spec, tags)
	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)
		instance.Status.Message = err.Error()

		switch azerr.Type {
		case errhelp.AsyncOpIncompleteError:
			instance.Status.State = "Creating"
			instance.Status.Message = "Resource request successfully submitted to Azure"
			instance.Status.SpecHash = hash
			instance.Status.PollingURL = pollingUrl
			return false, nil
		case errhelp.InvalidResourceLocation, errhelp.LocationNotAvailableForResourceType, errhelp.BadRequest:
			instance.Status.Provisioning = false
			instance.Status.Message = azerr.Error()
			return true, nil
		case errhelp.ResourceGroupNotFoundErrorCode, errhelp.ParentNotFoundErrorCode:
			instance.Status.Provisioning = false
		case errhelp.NotFoundErrorCode:
			nameExists, err := m.CheckNameExistsCosmosDB(ctx, accountName)
			if err != nil {
				instance.Status.Message = err.Error()
			}
			if nameExists {
				instance.Status.Provisioning = false
				instance.Status.Message = "CosmosDB Account name already exists"
				return true, nil
			}
		}

		return false, err
	}

	if err = m.createOrUpdateSecret(ctx, instance, db); err != nil {
		instance.Status.Message = err.Error()
		return false, err
	}

	instance.Status.SpecHash = hash
	instance.Status.ResourceId = *db.ID
	instance.Status.State = *db.ProvisioningState
	instance.Status.Provisioned = true
	instance.Status.Provisioning = false
	instance.Status.Message = resourcemanager.SuccessMsg
	return false, nil
}

// Delete drops cosmosdb
func (m *AzureCosmosDBManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		m.SecretClient = options.SecretClient
	}

	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	accountName := instance.ObjectMeta.Name

	// try to delete the cosmosdb instance & secrets
	_, err = m.DeleteCosmosDB(ctx, groupName, accountName)
	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)

		// request submitted or already in progress
		if azerr.Type == errhelp.AsyncOpIncompleteError || (azerr.Type == errhelp.PreconditionFailed && strings.Contains(azerr.Reason, "operation in progress")) {
			instance.Status.State = "Deleting"
			instance.Status.Message = "Deletion request submitted successfully"
			return true, nil
		}

		notFound := []string{
			errhelp.NotFoundErrorCode,
			errhelp.ResourceNotFound,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		if helpers.ContainsString(notFound, azerr.Type) {
			_ = m.deleteSecret(ctx, instance)
			return false, nil
		}

		// unhandled error
		instance.Status.Message = azerr.Error()
		return false, err
	}

	_ = m.deleteSecret(ctx, instance)
	return false, nil
}

// GetParents returns the parents of cosmosdb
func (m *AzureCosmosDBManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &v1alpha1.ResourceGroup{},
		},
	}, nil
}

// GetStatus gets the ASOStatus
func (m *AzureCosmosDBManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (m *AzureCosmosDBManager) convert(obj runtime.Object) (*v1alpha1.CosmosDB, error) {
	db, ok := obj.(*v1alpha1.CosmosDB)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return db, nil
}

func (m *AzureCosmosDBManager) createOrUpdateSecret(ctx context.Context, instance *v1alpha1.CosmosDB, db *documentdb.DatabaseAccount) error {
	connStrResult, err := m.ListConnectionStrings(ctx, instance.Spec.ResourceGroup, instance.ObjectMeta.Name)
	if err != nil {
		return err
	}

	keysResult, err := m.ListKeys(ctx, instance.Spec.ResourceGroup, instance.ObjectMeta.Name)
	if err != nil {
		return err
	}

	secretKey := types.NamespacedName{
		Name:      instance.Name,
		Namespace: instance.Namespace,
	}
	secretData := map[string][]byte{
		"primaryEndpoint":            []byte(*db.DocumentEndpoint),
		"primaryMasterKey":           []byte(*keysResult.PrimaryMasterKey),
		"secondaryMasterKey":         []byte(*keysResult.SecondaryMasterKey),
		"primaryReadonlyMasterKey":   []byte(*keysResult.PrimaryReadonlyMasterKey),
		"secondaryReadonlyMasterKey": []byte(*keysResult.SecondaryReadonlyMasterKey),
	}

	// set all available connection strings in the secret
	if connStrResult.ConnectionStrings != nil {
		for _, cs := range *connStrResult.ConnectionStrings {
			secretData[helpers.RemoveNonAlphaNumeric(*cs.Description)] = []byte(*cs.ConnectionString)
		}
	}

	// set each location's endpoint in the secret
	if db.DatabaseAccountProperties.ReadLocations != nil {
		for _, l := range *db.DatabaseAccountProperties.ReadLocations {
			safeLocationName := helpers.RemoveNonAlphaNumeric(strings.ToLower(*l.LocationName))
			secretData[safeLocationName+"Endpoint"] = []byte(*l.DocumentEndpoint)
		}
	}

	return m.SecretClient.Upsert(ctx, secretKey, secretData)
}

func (m *AzureCosmosDBManager) deleteSecret(ctx context.Context, instance *v1alpha1.CosmosDB) error {
	secretKey := types.NamespacedName{
		Name:      instance.Name,
		Namespace: instance.Namespace,
	}
	return m.SecretClient.Delete(ctx, secretKey)
}
