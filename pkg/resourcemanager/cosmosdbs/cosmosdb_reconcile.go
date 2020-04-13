// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package cosmosdbs

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"

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
	if instance.Status.SpecHash != hash {
		// need to push a create or update
		instance.Status.SpecHash = hash
	} else if instance.Status.Provisioned {
		// provisioned and no changes needed
		instance.Status.RequestedAt = nil
		return true, nil
	} else if instance.Status.Provisioning {
		// get the instance and update status
		db, azerr := m.GetCosmosDB(ctx, instance.Spec.ResourceGroup, instance.Name)
		if azerr == nil {
			instance.Status.ResourceId = *db.ID
			instance.Status.State = *db.ProvisioningState

			if instance.Status.State == "Creating" {
				instance.Status.Message = "Waiting for resource to finish creation"
				return false, nil
			}

			if instance.Status.State == "Succeeded" {
				// provisioning is complete, update the secrets
				if err = m.createOrUpdateAccountKeysSecret(ctx, instance); err != nil {
					instance.Status.Message = err.Error()
					return false, err
				}

				instance.Status.Message = resourcemanager.SuccessMsg
				instance.Status.Provisioning = false
				instance.Status.Provisioned = true
				return true, nil
			}

			if instance.Status.State == "Failed" {
				instance.Status.Message = "Failed to provision CosmosDB"
				instance.Status.Provisioning = false
				return true, nil
			}
		} else if azerr.Type == errhelp.ResourceGroupNotFoundErrorCode {
			instance.Status.Provisioning = false
			instance.Status.Message = fmt.Sprintf("Waiting for resource group '%s' to be available", instance.Spec.ResourceGroup)
			instance.Status.State = "Waiting"
			return false, nil
		} else if azerr.Type == errhelp.ResourceNotFound {
			exists, azerr := m.CheckNameExistsCosmosDB(ctx, instance.Name)
			if azerr != nil {
				instance.Status.Provisioning = false
				instance.Status.Message = "Unexpected error occurred during resource request"
				instance.Status.State = "Failed"
				return false, err
			} else if exists {
				// get request returned resource not found and the name already exists
				// so it must exist in a different resource group, user must fix it
				instance.Status.Provisioning = false
				instance.Status.Message = "CosmosDB name already exists"
				instance.Status.State = "Failed"
				return true, nil
			}
		} else {
			instance.Status.Provisioning = false
			instance.Status.Message = azerr.Error()
			return false, azerr.Original
		}
	}

	instance.Status.Provisioning = true

	tags := helpers.LabelsToTags(instance.GetLabels())
	accountName := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup
	location := instance.Spec.Location
	kind := instance.Spec.Kind
	dbType := instance.Spec.Properties.DatabaseAccountOfferType

	db, azerr := m.CreateOrUpdateCosmosDB(ctx, groupName, accountName, location, kind, dbType, tags)

	// everything is in a created/updated state
	if azerr == nil {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.State = "Succeeded"
		instance.Status.ResourceId = *db.ID
		return true, nil
	}

	switch azerr.Type {

	case errhelp.AsyncOpIncompleteError:
		instance.Status.Message = "Resource request successfully submitted to Azure"
		instance.Status.State = "Creating"

	case errhelp.InvalidResourceLocation:
		instance.Status.Provisioning = false
		instance.Status.Message = azerr.Reason
		return true, nil

	}
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

	accountName := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroup

	// if the resource is in a failed state it was never created or could never be verified
	// so we skip attempting to delete the resrouce from Azure
	if instance.Status.FailedProvisioning {
		return false, nil
	}

	notFoundErrors := []string{
		errhelp.NotFoundErrorCode,              // happens on first request after deletion succeeds
		errhelp.ResourceNotFound,               // happens on subsequent requests after deletion succeeds
		errhelp.ResourceGroupNotFoundErrorCode, // database doesn't exist in this resource group but the name exists globally
	}

	// fetch the latest to inspect provisioning state
	cosmosDB, azerr := m.GetCosmosDB(ctx, groupName, accountName)
	if azerr != nil {
		// deletion finished
		if helpers.ContainsString(notFoundErrors, azerr.Type) {
			return false, nil
		}

		//TODO: are there other errors that need handling here?
		instance.Status.Message = azerr.Error()
		return true, azerr.Original
	}

	instance.Status.State = *cosmosDB.ProvisioningState

	// already deleting the resource, try again later
	if instance.Status.State == "Deleting" {
		return true, nil
	}

	// create must succeed before delete succeeds
	if instance.Status.State == "Creating" {
		return true, nil
	}

	// try to delete the cosmosdb instance & secrets
	_, azerr = m.DeleteCosmosDB(ctx, groupName, accountName)
	m.deleteAccountKeysSecret(ctx, instance)
	if azerr != nil {
		// this is likely to happen on first try due to not waiting for the future to complete
		if azerr.Type == errhelp.AsyncOpIncompleteError {
			instance.Status.Message = "Deletion request submitted successfully"
			return true, nil
		}

		// already deleted
		if helpers.ContainsString(notFoundErrors, azerr.Type) {
			return false, nil
		}

		// unhandled error
		instance.Status.Message = azerr.Error()
		return false, azerr.Original
	}

	// second delete calls succeed immediately
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

func (m *AzureCosmosDBManager) createOrUpdateAccountKeysSecret(ctx context.Context, instance *v1alpha1.CosmosDB) error {
	result, err := m.ListKeys(ctx, instance.Spec.ResourceGroup, instance.ObjectMeta.Name)
	if err != nil {
		return err
	}

	secretKey := types.NamespacedName{
		Name:      instance.Name,
		Namespace: instance.Namespace,
	}
	secretData := map[string][]byte{
		"primaryConnectionString":    []byte(*result.PrimaryMasterKey),
		"secondaryConnectionString":  []byte(*result.SecondaryMasterKey),
		"primaryReadonlyMasterKey":   []byte(*result.PrimaryReadonlyMasterKey),
		"secondaryReadonlyMasterKey": []byte(*result.SecondaryReadonlyMasterKey),
	}

	return m.SecretClient.Upsert(ctx, secretKey, secretData)
}

func (m *AzureCosmosDBManager) deleteAccountKeysSecret(ctx context.Context, instance *v1alpha1.CosmosDB) error {
	secretKey := types.NamespacedName{
		Name:      instance.Name,
		Namespace: instance.Namespace,
	}
	return m.SecretClient.Delete(ctx, secretKey)
}
