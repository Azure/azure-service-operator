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
	groupName := instance.Spec.ResourceGroup
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
	} else if azerr.Type == errhelp.ResourceNotFound {
		exists, azerr := m.CheckNameExistsCosmosDB(ctx, accountName)
		if azerr != nil {
			return unexpectedError(&instance.Status, azerr.Original)
		}
		if exists {
			// get request returned resource not found and the name already exists
			// so it must exist in a different resource group, user must fix it
			return nameAlreadyExists(&instance.Status)
		}
		instance.Status.State = "Creating"
	} else if azerr.Type == errhelp.ResourceGroupNotFoundErrorCode {
		return waitingForParent(&instance.Status, groupName)
	}

	// create the database
	instance.Status.Provisioning = true
	instance.Status.Provisioned = false
	instance.Status.ResourceId = ""
	instance.Status.SpecHash = ""
	instance.Status.Message = ""
	db, azerr = m.CreateOrUpdateCosmosDB(ctx, groupName, accountName, location, kind, dbType, tags)

	// everything is in a created/updated state
	if azerr == nil {
		instance.Status.ResourceId = *db.ID
		instance.Status.SpecHash = hash
		return success(&instance.Status)
	}

	switch azerr.Type {
	case errhelp.AsyncOpIncompleteError:
		return creatingOrUpdating(&instance.Status)
	case errhelp.InvalidResourceLocation:
		instance.Status.Provisioning = false
		return true, fmt.Errorf(azerr.Reason)
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

	// try to delete the cosmosdb instance
	_, azerr = m.DeleteCosmosDB(ctx, groupName, accountName)
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

func success(status *v1alpha1.ASOStatus) (bool, error) {
	status.Provisioned = true
	status.Provisioning = false
	status.Message = resourcemanager.SuccessMsg
	status.State = "Succeeded"
	return true, nil
}

func waitingForParent(status *v1alpha1.ASOStatus, groupName string) (bool, error) {
	status.Provisioned = false
	status.Provisioning = false
	status.Message = fmt.Sprintf("Waiting for resource group '%s' to be available", groupName)
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
	return false, err
}

func nameAlreadyExists(status *v1alpha1.ASOStatus) (bool, error) {
	status.Provisioned = false
	status.Provisioning = false
	status.Message = "cosmosdb name already exists"
	status.State = "Failed"
	return true, fmt.Errorf("cosmosdb name already exists")
}
