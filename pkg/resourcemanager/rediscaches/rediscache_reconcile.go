// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package rediscaches

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure creates a rediscache
func (rc *AzureRedisCacheManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		rc.SecretClient = options.SecretClient
	}

	instance, err := rc.convert(obj)
	if err != nil {
		return false, err
	}

	redisName := instance.Name
	groupName := instance.Spec.ResourceGroupName
	name := instance.ObjectMeta.Name

	if len(instance.Spec.SecretName) == 0 {
		instance.Spec.SecretName = redisName
	}

	instance.Status.Provisioning = true

	newRc, err := rc.GetRedisCache(ctx, groupName, name)
	if err == nil {
		if newRc.ProvisioningState == "Succeeded" {
			err = rc.ListKeysAndCreateSecrets(groupName, redisName, instance.Spec.SecretName, instance)
			if err != nil {
				instance.Status.Message = err.Error()
				return false, err
			}
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.State = string(newRc.ProvisioningState)
			instance.Status.ResourceId = *newRc.ID
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			return true, nil
		}
		instance.Status.Message = "RedisCache exists but may not be ready"
		instance.Status.State = string(newRc.ProvisioningState)
		return false, nil
	}
	instance.Status.Message = fmt.Sprintf("RedisCache Get error %s", err.Error())

	result, err := rc.CreateRedisCache(ctx, *instance)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		instance.Status.Provisioning = false

		if result != nil {
			// stop reconciling if the spec is bad
			if result.StatusCode == http.StatusBadRequest {
				return true, nil
			}
		}
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

	return true, nil
}

// Delete drops a rediscache
func (rc *AzureRedisCacheManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		rc.SecretClient = options.SecretClient
	}

	instance, err := rc.convert(obj)
	if err != nil {
		return false, err
	}

	name := instance.ObjectMeta.Name
	groupName := instance.Spec.ResourceGroupName

	if len(instance.Spec.SecretName) == 0 {
		instance.Spec.SecretName = name
	}

	// key for SecretClient to delete secrets on successful deletion
	key := types.NamespacedName{Name: instance.Spec.SecretName, Namespace: instance.Namespace}

	resp, err := rc.GetRedisCache(ctx, groupName, name)
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			// Best case deletion of secrets
			rc.SecretClient.Delete(ctx, key)
			return false, nil
		}
		return false, err
	}

	if resp.ProvisioningState == "Deleting" || resp.ProvisioningState == "Creating" {
		instance.Status.Message = fmt.Sprintf("Async Operation: %s not complete", resp.ProvisioningState)
		return true, nil
	}

	req, err := rc.DeleteRedisCache(ctx, groupName, name)
	if err != nil {
		instance.Status.Message = err.Error()

		if req.Response().StatusCode == http.StatusNotFound {
			// Best case deletion of secrets
			rc.SecretClient.Delete(ctx, key)
			return false, nil
		}

		return true, fmt.Errorf("AzureRedisCacheManager Delete failed with %s", err)
	}

	return true, nil
}

// GetParents returns the parents of rediscache
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

// GetStatus gets the ASOStatus
func (g *AzureRedisCacheManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (rc *AzureRedisCacheManager) convert(obj runtime.Object) (*azurev1alpha1.RedisCache, error) {
	local, ok := obj.(*azurev1alpha1.RedisCache)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
