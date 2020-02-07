/*
Copyright 2019 microsoft.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azuresqldb

import (
	"context"
	"fmt"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure creates an AzureSqlDb
func (db *AzureSqlDbManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.EnsureOption) (bool, error) {

	instance, err := db.convert(obj)
	if err != nil {
		return false, err
	}

	location := instance.Spec.Location
	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	dbName := instance.ObjectMeta.Name
	dbEdition := instance.Spec.Edition

	azureSqlDatabaseProperties := azuresqlshared.SQLDatabaseProperties{
		DatabaseName: dbName,
		Edition:      dbEdition,
	}

	instance.Status.Provisioning = true

	_, err = db.CreateOrUpdateDB(ctx, groupName, location, server, azureSqlDatabaseProperties)
	if err != nil {
		instance.Status.Message = err.Error()
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}

		return true, fmt.Errorf("AzureSqlDb CreateOrUpdate error %v", err)
	}

	resp, err := db.GetDB(ctx, groupName, server, dbName)
	if err != nil {
		instance.Status.Message = err.Error()
		catch := []string{
			errhelp.NotFoundErrorCode,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}
		return false, fmt.Errorf("AzureSqlDb GetDB error %v", err)
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true
	instance.Status.State = string(*resp.Status)
	instance.Status.Message = resourcemanager.SuccessMsg

	return true, nil
}

// Delete drops a AzureSqlDb
func (db *AzureSqlDbManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := db.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	server := instance.Spec.Server
	dbName := instance.ObjectMeta.Name

	_, err = db.DeleteDB(ctx, groupName, server, dbName)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			// Database does not exist
			return false, nil
		}

		return true, fmt.Errorf("AzureSqlDb delete error %v", err)
	}

	return false, nil
}

// GetParents returns the parents of AzureSqlDatabase
func (db *AzureSqlDbManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := db.convert(obj)
	if err != nil {
		return nil, err
	}

	rgKey := types.NamespacedName{Name: instance.Spec.ResourceGroup, Namespace: instance.Namespace}
	key := types.NamespacedName{Name: instance.Spec.Server, Namespace: instance.Namespace}

	return []resourcemanager.KubeParent{
		{Key: key, Target: &azurev1alpha1.AzureSqlServer{}},
		{Key: rgKey, Target: &azurev1alpha1.ResourceGroup{}},
	}, nil
}

func (*AzureSqlDbManager) convert(obj runtime.Object) (*azurev1alpha1.AzureSqlDatabase, error) {
	local, ok := obj.(*azurev1alpha1.AzureSqlDatabase)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
