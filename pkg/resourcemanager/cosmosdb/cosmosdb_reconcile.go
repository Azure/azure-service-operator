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
	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	// convert kube labels to expected tag format
	tags := map[string]*string{}
	for k, v := range instance.GetLabels() {
		value := v
		tags[k] = &value
	}

	hash := helpers.Hash256(instance.Spec)
	accountName := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	location := instance.Spec.Location
	kind := instance.Spec.Kind
	dbType := instance.Spec.Properties.DatabaseAccountOfferType

	// get the database to see if it exists
	db, azerr := m.GetCosmosDB(ctx, groupName, accountName)
	if azerr == nil {
		instance.Status.ResourceId = *db.ID
		instance.Status.State = *db.ProvisioningState

		switch *db.ProvisioningState {
		case "Creating":
			return creatingOrUpdating(&instance.Status)
		case "Succeeded":
			if instance.Status.SpecHash == hash {
				return success(&instance.Status)
			}
			// needs update, allow to fall through
			instance.Status.State = "Updating"
		case "Deleting":
			// previous delete hasn't finished, try again later
			return waitingForDelete(&instance.Status)
		default:
			//TODO: how should we handle unknown states?
			return true, nil
		}
	} else if azerr.Code == 404 {
		// resource not found, need to create it
		instance.Status.State = "Creating"
	}

	// create the database
	db, azerr = m.CreateOrUpdateCosmosDB(ctx, groupName, accountName, location, kind, dbType, tags)

	// everything is in a created/updated state
	if azerr == nil {
		instance.Status.ResourceId = *db.ID
		instance.Status.SpecHash = hash
		return success(&instance.Status)
	}

	switch azerr.Type {
	case errhelp.ResourceGroupNotFoundErrorCode:
		return waitingForParent(&instance.Status)
	case errhelp.AsyncOpIncompleteError:
		return creatingOrUpdating(&instance.Status)
	default:
		return unexpectedError(&instance.Status, err)
	}
}

// Delete drops cosmosdb
func (m *AzureCosmosDBManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName

	resp, azerr := m.DeleteCosmosDB(ctx, groupName, name)
	if azerr != nil {
		instance.Status.Message = azerr.Error()

		//TODO: figure out edge cases
		fmt.Printf("Status = %s\n", resp.Status)

		// couldn't delete resource in azure
		return true, err
	}

	// already deleted,
	if resp.StatusCode == 204 {
		return false, nil
	}

	// deleted successfully
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
				Name:      instance.Spec.ResourceGroupName,
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

func success(status *v1alpha1.ASOStatus) (bool, error) {
	status.Provisioned = true
	status.Provisioning = false
	status.Message = "Resource successfully provisioned"
	status.State = "Succeeded"
	return true, nil
}

func waitingForParent(status *v1alpha1.ASOStatus) (bool, error) {
	status.Provisioned = false
	status.Provisioning = false
	status.Message = "Waiting for resource group to be available"
	status.State = "Waiting"
	return false, nil
}

func waitingForDelete(status *v1alpha1.ASOStatus) (bool, error) {
	status.Provisioned = false
	status.Provisioning = false
	status.Message = "Waiting for previous delete to finish"
	status.State = "Waiting"
	return false, nil
}

func creatingOrUpdating(status *v1alpha1.ASOStatus) (bool, error) {
	status.Provisioned = false
	status.Provisioning = true
	status.Message = "Resource request successfully submitted to Azure"
	return false, nil
}

func unexpectedError(status *v1alpha1.ASOStatus, err error) (bool, error) {
	status.Provisioned = false
	status.Provisioning = false
	status.Message = "Unexpected error occurred during resource request"
	status.State = "Failed"
	status.FailedProvisioning = true
	return false, err
}
