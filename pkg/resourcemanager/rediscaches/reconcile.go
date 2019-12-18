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
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

func (rc *AzureRedisCacheManager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := rc.convert(obj)
	if err != nil {
		return false, err
	}

	location := instance.Spec.Location
	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName
	sku := instance.Spec.Properties.Sku
	enableNonSSLPort := instance.Spec.Properties.EnableNonSslPort

	instance.Status.Provisioning = true

	_, err = rc.CreateRedisCache(ctx, groupName, name, location, sku, enableNonSSLPort, nil)
	if err != nil {
		instance.Status.Provisioning = false
		instance.Status.Message = err.Error()
		return false, fmt.Errorf("Redis Cache create error %v", err)
	}

	instance.Status.Provisioning = false
	instance.Status.Provisioned = true

	return true, nil
}

func (rc *AzureRedisCacheManager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {

	instance, err := rc.convert(obj)
	if err != nil {
		return false, err
	}

	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName

	_, err = rc.DeleteRedisCache(ctx, groupName, name)
	if err != nil {
		if errhelp.IsStatusCode204(err) {
			return true, fmt.Errorf("RedisCache delete error %v", err)
		}
	}
	return true, nil
}

func (rc *AzureRedisCacheManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := rc.convert(obj)
	if err != nil {
		return nil, err
	}

	key := types.NamespacedName{Namespace: instance.Namespace, Name: instance.Spec.ResourceGroupName}

	return []resourcemanager.KubeParent{
		{Key: key, Target: &v1alpha1.ResourceGroup{}},
	}, nil
}

func (rc *AzureRedisCacheManager) convert(obj runtime.Object) (*azurev1alpha1.RedisCache, error) {
	local, ok := obj.(*azurev1alpha1.RedisCache)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
