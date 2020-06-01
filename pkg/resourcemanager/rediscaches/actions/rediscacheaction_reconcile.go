// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package actions

import (
	"context"
	"fmt"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

func (m *AzureRedisCacheActionManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		m.SecretClient = options.SecretClient
	}

	instance, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	// never re-provision an action
	if instance.Status.Provisioned {
		return true, nil
	}

	rollAllKeys := instance.Spec.ActionName == v1alpha1.RedisCacheActionNameRollAllKeys

	if rollAllKeys || instance.Spec.ActionName == v1alpha1.RedisCacheActionNameRollPrimaryKey {
		if err = m.RegeneratePrimaryAccessKey(ctx, instance.Spec.ResourceGroup, instance.Spec.CacheName); err != nil {
			instance.Status.Message = err.Error()
			return false, err
		}
	}

	if rollAllKeys || instance.Spec.ActionName == v1alpha1.RedisCacheActionNameRollSecondaryKey {
		if err = m.RegenerateSecondaryAccessKey(ctx, instance.Spec.ResourceGroup, instance.Spec.CacheName); err != nil {
			instance.Status.Message = err.Error()
			return false, err
		}
	}

	// regenerate the secret
	cacheInstance := &v1alpha1.RedisCache{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Spec.CacheName,
			Namespace: instance.Namespace,
		},
		Spec: v1alpha1.RedisCacheSpec{
			SecretName:        instance.Spec.SecretName,
			ResourceGroupName: instance.Spec.ResourceGroup,
		},
	}
	if err = m.ListKeysAndCreateSecrets(ctx, cacheInstance); err != nil {
		instance.Status.Provisioned = false
		instance.Status.FailedProvisioning = true
		instance.Status.Message = err.Error()
		return false, err
	}

	// successful return
	instance.Status.Provisioned = true
	instance.Status.FailedProvisioning = false
	instance.Status.Message = resourcemanager.SuccessMsg
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
				Name:      instance.Spec.CacheName,
			},
			Target: &v1alpha1.RedisCache{},
		},
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
