// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package cosmosdbs

import (
	"context"
	"fmt"
	"strings"

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

	if instance.Status.SpecHash == hash && instance.Status.Provisioned {
		instance.Status.RequestedAt = nil
		return true, nil
	}
	instance.Status.Provisioned = false

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

	if instance.Status.State == "Creating" {
		// avoid multiple CreateOrUpdate requests while resource is already creating
		return false, nil
	}

	if instance.Status.State == "Succeeded" {
		// provisioning is complete, update the secrets
		if err = m.createOrUpdateAccountKeysSecret(ctx, instance); err != nil {
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
	groupName := instance.Spec.ResourceGroup
	location := instance.Spec.Location
	kind := instance.Spec.Kind
	networkRule := instance.Spec.VirtualNetworkRules

	cosmosDBProperties := v1alpha1.CosmosDBProperties{
		DatabaseAccountOfferType:      instance.Spec.Properties.DatabaseAccountOfferType,
		EnableMultipleWriteLocations:  instance.Spec.Properties.EnableMultipleWriteLocations,
		MongoDBVersion:                instance.Spec.Properties.MongoDBVersion,
		IsVirtualNetworkFilterEnabled: instance.Spec.Properties.IsVirtualNetworkFilterEnabled,
	}

	db, err = m.CreateOrUpdateCosmosDB(ctx, groupName, accountName, location, kind, networkRule, cosmosDBProperties, tags)
	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)
		instance.Status.Message = err.Error()

		switch azerr.Type {
		case errhelp.AsyncOpIncompleteError:
			instance.Status.State = "Creating"
			instance.Status.Message = "Resource request successfully submitted to Azure"
			instance.Status.SpecHash = hash
			return false, nil
		case errhelp.InvalidResourceLocation, errhelp.LocationNotAvailableForResourceType:
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

	if err = m.createOrUpdateAccountKeysSecret(ctx, instance); err != nil {
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

	// if the resource is in a failed state it was never created or could never be verified
	// so we skip attempting to delete the resrouce from Azure
	if instance.Status.FailedProvisioning {
		return false, nil
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
			return false, m.deleteAccountKeysSecret(ctx, instance)
		}

		// unhandled error
		instance.Status.Message = azerr.Error()
		return false, err
	}

	return false, m.deleteAccountKeysSecret(ctx, instance)
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

	err = m.SecretClient.Upsert(ctx, secretKey, secretData)
	if err != nil {
		return err
	}

	return nil
}

func (m *AzureCosmosDBManager) deleteAccountKeysSecret(ctx context.Context, instance *v1alpha1.CosmosDB) error {
	secretKey := types.NamespacedName{
		Name:      instance.Name,
		Namespace: instance.Namespace,
	}
	return m.SecretClient.Delete(ctx, secretKey)
}
