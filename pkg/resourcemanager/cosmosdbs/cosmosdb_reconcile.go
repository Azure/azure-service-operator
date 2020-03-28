// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package cosmosdbs

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
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

	//hash := helpers.Hash256(instance.Spec)
	accountName := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	location := instance.Spec.Location
	kind := instance.Spec.Kind
	dbType := instance.Spec.Properties.DatabaseAccountOfferType

	// get the database to see if it exists
	db, err := m.GetCosmosDB(ctx, groupName, accountName)
	if err == nil {
		//TODO: do we need to handle updates here?
		needsUpdate := false

		// already exists and requires an update
		if needsUpdate {
			instance.Status.Provisioning = true
			instance.Status.Provisioned = false
			instance.Status.ResourceId = *db.ID
			instance.Status.State = *db.ProvisioningState
			return false, nil
		}

		// already exists and doesn't require an update
		instance.Status.Provisioning = false
		instance.Status.Provisioned = true
		instance.Status.ResourceId = *db.ID
		instance.Status.State = *db.ProvisioningState
		return true, nil
	}

	// create the database
	_, err = m.CreateOrUpdateCosmosDB(ctx, groupName, accountName, location, kind, dbType, tags)
	if err == nil {
		// creation succeeded, update state and return
		instance.Status.Provisioning = false
		instance.Status.Provisioned = true
		return true, nil
	}

	// handle errors
	return false, err
}

// Delete drops cosmosdb
func (m *AzureCosmosDBManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	_, err = m.DeleteCosmosDB(ctx, groupName, name)
	if err != nil {
		instance.Status.Message = err.Error()

		//TODO: figure out edge cases

		// couldn't delete resource in azure
		return true, err
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
