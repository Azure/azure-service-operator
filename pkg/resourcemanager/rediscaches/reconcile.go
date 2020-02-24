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

package rediscaches

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure checks the desired state of the operator
func (rc *AzureRedisCacheManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := rc.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroupName
	name := instance.ObjectMeta.Name
	location := instance.Spec.Location
	sku := instance.Spec.Properties.Sku
	enableNonSSLPort := instance.Spec.Properties.EnableNonSslPort

	instance.Status.Provisioning = true

	redisClient, err := getRedisCacheClient()
	if err != nil {
		return false, err
	}
	resp, err := redisClient.Get(ctx, groupName, name)
	if err == nil {
		if resp.ProvisioningState == "Succeeded" {
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			return true, nil
		}
		instance.Status.Message = "RedisCache exists but may not be ready"
		instance.Status.State = string(resp.ProvisioningState)
		return false, nil
	}
	instance.Status.Message = fmt.Sprintf("RedisCache Get error %s", err.Error())

	_, err = rc.CreateRedisCache(ctx, groupName, name, location, sku, enableNonSSLPort, nil)
	if err != nil {
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false

		catch := []string{
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.AlreadyExists,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
			errhelp.BadRequest,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}
		return false, err
	}

	// Check to see if secret already exists
	secret, err := rc.GetOrPrepareSecret(ctx, instance)
	// create or update the secret
	key := types.NamespacedName{Name: instance.ObjectMeta.Name, Namespace: instance.Namespace}
	err = rc.SecretClient.Upsert(
		ctx,
		key,
		secret,
		secrets.WithOwner(instance),
		secrets.WithScheme(rc.Scheme),
	)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("GetOrPrepareSecrets failed with err %s", err.Error())
		return false, err
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true
	instance.Status.Message = resourcemanager.SuccessMsg

	return true, nil
}

// Delete removes a RedisCache resource
func (rc *AzureRedisCacheManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := rc.convert(obj)
	if err != nil {
		return false, err
	}

	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName

	_, err = rc.DeleteRedisCache(ctx, groupName, name)
	if err != nil {
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.CreationPending,
			errhelp.AsyncOpIncompleteError,
		}
		err = errhelp.NewAzureError(err)
		if azerr, ok := err.(*errhelp.AzureError); ok {
			if helpers.ContainsString(catch, azerr.Type) {
				return false, nil
			}
		}
		return false, fmt.Errorf("AzureRedisCacheManager Delete failed with %s", err)
	}
	return false, nil
}

// GetParents fetches dependent ARM resources
func (rc *AzureRedisCacheManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := rc.convert(obj)
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

func (rc *AzureRedisCacheManager) convert(obj runtime.Object) (*azurev1alpha1.RedisCache, error) {
	local, ok := obj.(*azurev1alpha1.RedisCache)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
