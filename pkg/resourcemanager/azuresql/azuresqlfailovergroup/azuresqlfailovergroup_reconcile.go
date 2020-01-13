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

package azuresqlfailovergroup

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

// Ensure creates a sqlfailovergroup
func (fg *AzureSqlFailoverGroupManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := fg.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	serverName := instance.Spec.Server
	failoverGroupName := instance.ObjectMeta.Name
	failoverPolicy := instance.Spec.FailoverPolicy
	failoverGracePeriod := instance.Spec.FailoverGracePeriod
	secondaryServer := instance.Spec.SecondaryServerName
	secondaryResourceGroup := instance.Spec.SecondaryServerResourceGroup
	databaseList := instance.Spec.DatabaseList
	sqlFailoverGroupProperties := azuresqlshared.SQLFailoverGroupProperties{
		FailoverPolicy:               failoverPolicy,
		FailoverGracePeriod:          failoverGracePeriod,
		SecondaryServerName:          secondaryServer,
		SecondaryServerResourceGroup: secondaryResourceGroup,
		DatabaseList:                 databaseList,
	}

	_, err = fg.CreateOrUpdateFailoverGroup(ctx, groupName, serverName, failoverGroupName, sqlFailoverGroupProperties)
	if err != nil {
		catch := []string{
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			instance.Status.Message = fmt.Sprintf("Got ignorable error of type %v", azerr.Type)
			return false, nil
		}
		instance.Status.Message = fmt.Sprintf("AzureSqlFailoverGroup CreateOrUpdate error: %s", err.Error())
		return true, err
	}

	_, err = fg.GetFailoverGroup(ctx, groupName, serverName, failoverGroupName)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("AzureSqlFailoverGroup Get error: %s", err.Error())
		return true, errhelp.NewAzureError(err)
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true
	instance.Status.Message = "AzureSqlFailoverGroup Successfully Provisioned"

	return true, nil
}

// Delete drops a sqlfailovergroup
func (fg *AzureSqlFailoverGroupManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := fg.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	serverName := instance.Spec.Server
	failoverGroupName := instance.ObjectMeta.Name

	resp, err := fg.DeleteFailoverGroup(ctx, groupName, serverName, failoverGroupName)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("AzureSqlFailoverGroup Delete failed with: %s", err.Error())
		return false, errhelp.NewAzureError(err)
	}
	if resp.StatusCode == 200 {
		instance.Status.Message = fmt.Sprintf("Delete AzureSqlFailoverGroup succeeded")
	}
	return true, nil
}

// GetParents returns the parents of sqlfailovergroup
func (fg *AzureSqlFailoverGroupManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := fg.convert(obj)
	if err != nil {
		return nil, err
	}

	// add db
	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Server,
			},
			Target: &azurev1alpha1.AzureSqlServer{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil

}

func (fg *AzureSqlFailoverGroupManager) convert(obj runtime.Object) (*azurev1alpha1.AzureSqlFailoverGroup, error) {
	local, ok := obj.(*azurev1alpha1.AzureSqlFailoverGroup)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
