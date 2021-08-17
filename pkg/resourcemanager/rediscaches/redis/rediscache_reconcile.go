// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package redis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure creates a rediscache
func (rc *AzureRedisCacheManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	secretClient := rc.SecretClient
	if options.SecretClient != nil {
		secretClient = options.SecretClient
	}

	instance, err := rc.convert(obj)
	if err != nil {
		return false, err
	}

	name := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
	groupName := instance.Spec.ResourceGroupName

	// if an error occurs thats ok as it means that it doesn't exist yet
	newRc, err := rc.GetRedisCache(ctx, groupName, name.Name)
	if err == nil {

		// succeeded! so end reconcilliation successfully
		if newRc.ProvisioningState == "Succeeded" {
			err = rc.ListKeysAndCreateSecrets(ctx, secretClient, instance)
			if err != nil {
				instance.Status.Message = err.Error()
				return false, err
			}

			if newRc.StaticIP != nil {
				instance.Status.Output = *newRc.StaticIP
			}

			instance.Status.ResourceId = *newRc.ID
			instance.Status.State = string(newRc.ProvisioningState)
			instance.Status.SetProvisioned(resourcemanager.SuccessMsg)
			return true, nil
		}

		// the redis cache exists but has not provisioned yet - so keep waiting
		instance.Status.Message = "RedisCache exists but may not be ready"
		instance.Status.State = string(newRc.ProvisioningState)
		return false, nil
	}

	// actually provision the redis cache
	instance.Status.SetProvisioning("")
	_, err = rc.CreateRedisCache(ctx, *instance)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		azerr := errhelp.NewAzureError(err)

		catchInProgress := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.AlreadyExists,
		}
		catchKnownError := []string{
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.NameNotAvailable,
		}

		// handle the error
		if helpers.ContainsString(catchInProgress, azerr.Type) {
			instance.Status.SetProvisioning("RedisCache exists but may not be ready")
			return false, nil
		} else if helpers.ContainsString(catchKnownError, azerr.Type) {
			return false, nil
		} else {
			// serious error occured, end reconcilliation and mark it as failed
			instance.Status.SetFailedProvisioning(fmt.Sprintf("Error occurred creating the RedisCache: %s", errhelp.StripErrorIDs(err)))
			return true, nil
		}
	}

	return false, nil
}

// Delete drops a rediscache
func (rc *AzureRedisCacheManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	secretClient := rc.SecretClient
	if options.SecretClient != nil {
		secretClient = options.SecretClient
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
	secretKey := secrets.SecretKey{Name: instance.Spec.SecretName, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}

	resp, err := rc.GetRedisCache(ctx, groupName, name)
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			// Best case deletion of secrets
			secretClient.Delete(ctx, secretKey)
			return false, nil
		}
		return false, err
	}

	if resp.ProvisioningState == "Deleting" || resp.ProvisioningState == "Creating" {
		instance.Status.Message = fmt.Sprintf("Async Operation: %s not complete", resp.ProvisioningState)
		return true, nil
	}

	_, err = rc.DeleteRedisCache(ctx, groupName, name)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureError(err)

		ignorableErr := []string{
			errhelp.AsyncOpIncompleteError,
		}

		finished := []string{
			errhelp.ResourceNotFound,
		}
		if helpers.ContainsString(ignorableErr, azerr.Type) {
			return true, nil
		}
		if helpers.ContainsString(finished, azerr.Type) {
			// Best case deletion of secrets
			secretClient.Delete(ctx, secretKey)
			return false, nil
		}
		return true, err
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
