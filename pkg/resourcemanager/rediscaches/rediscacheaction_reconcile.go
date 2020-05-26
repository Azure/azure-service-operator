// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package rediscaches

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

func (m *AzureRedisCacheActionManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	actionInstance, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	// never re-provision an action
	if actionInstance.Status.Provisioned {
		return true, nil
	}

	// get the RedisCache instance to see if it's done provisioning and set it as an owner
	cacheInstance := &v1alpha1.RedisCache{}
	cacheName := types.NamespacedName{Name: actionInstance.Spec.CacheName, Namespace: actionInstance.Namespace}
	err = options.Client.Get(ctx, cacheName, cacheInstance)
	if err != nil {
		return false, err
	}

	if cacheInstance.Status.FailedProvisioning || cacheInstance.Status.Provisioning {
		actionInstance.Status.Message = "Waiting for parent RedisCache to finish provisioning"
		return false, nil
	}

	rollAllKeys := actionInstance.Spec.ActionName == v1alpha1.RedisCacheActionNameRollAllKeys

	if rollAllKeys || actionInstance.Spec.ActionName == v1alpha1.RedisCacheActionNameRollPrimaryKey {
		if err = m.RegeneratePrimaryAccessKey(ctx, actionInstance.Spec.ResourceGroup, actionInstance.Spec.CacheName); err != nil {
			actionInstance.Status.Provisioned = false
			actionInstance.Status.FailedProvisioning = true
			return false, err
		}
	}

	if rollAllKeys || actionInstance.Spec.ActionName == v1alpha1.RedisCacheActionNameRollSecondaryKey {
		if err = m.RegenerateSecondaryAccessKey(ctx, actionInstance.Spec.ResourceGroup, actionInstance.Spec.CacheName); err != nil {
			actionInstance.Status.Provisioned = false
			actionInstance.Status.FailedProvisioning = true
			return false, err
		}
	}

	// regenerate the secret
	if err = m.ListKeysAndCreateSecrets(ctx, actionInstance.Spec.ResourceGroup, actionInstance.Spec.CacheName, cacheInstance.Spec.SecretName, cacheInstance); err != nil {
		actionInstance.Status.Provisioned = false
		actionInstance.Status.FailedProvisioning = true
		return false, err
	}

	// successful return
	actionInstance.Status.Provisioned = true
	actionInstance.Status.FailedProvisioning = false
	return true, nil
}

func (m *AzureRedisCacheActionManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	// no deletion necessary for deletion of action
	return false, nil
}

func (m *AzureRedisCacheActionManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
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
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.CacheName,
			},
			Target: &v1alpha1.RedisCache{},
		},
	}, nil
}

// GetStatus gets the ASOStatus
func (m *AzureRedisCacheActionManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (m *AzureRedisCacheActionManager) convert(obj runtime.Object) (*v1alpha1.RedisCacheAction, error) {
	local, ok := obj.(*v1alpha1.RedisCacheAction)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
