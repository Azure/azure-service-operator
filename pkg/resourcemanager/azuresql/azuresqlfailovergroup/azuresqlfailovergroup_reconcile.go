// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlfailovergroup

import (
	"context"
	"fmt"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure creates a sqlfailovergroup
func (fg *AzureSqlFailoverGroupManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		fg.SecretClient = options.SecretClient
	}

	instance, err := fg.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	serverName := instance.Spec.Server
	failoverGroupName := instance.ObjectMeta.Name
	sqlFailoverGroupProperties := azuresqlshared.SQLFailoverGroupProperties{
		FailoverPolicy:               instance.Spec.FailoverPolicy,
		FailoverGracePeriod:          instance.Spec.FailoverGracePeriod,
		SecondaryServer:              instance.Spec.SecondaryServer,
		SecondaryServerResourceGroup: instance.Spec.SecondaryServerResourceGroup,
		DatabaseList:                 instance.Spec.DatabaseList,
	}

	resp, err := fg.GetFailoverGroup(ctx, groupName, serverName, failoverGroupName)
	if err == nil {

		if *resp.ReplicationState == "SEEDING" {
			return false, nil
		}

		instance.Status.Provisioning = false
		instance.Status.Provisioned = true
		instance.Status.Message = resourcemanager.SuccessMsg
		instance.Status.ResourceId = *resp.ID
		return true, nil
	}
	instance.Status.Message = fmt.Sprintf("AzureSqlFailoverGroup Get error %s", err.Error())

	_, err = fg.CreateOrUpdateFailoverGroup(ctx, groupName, serverName, failoverGroupName, sqlFailoverGroupProperties)
	if err != nil {
		instance.Status.Message = err.Error()
		catch := []string{
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
			errhelp.ResourceNotFound,
			errhelp.AlreadyExists,
			errhelp.FailoverGroupBusy,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}
		return false, err
	}

	secret, _ := fg.GetOrPrepareSecret(ctx, instance)

	// create or update the secret
	key := types.NamespacedName{Name: instance.ObjectMeta.Name, Namespace: instance.Namespace}
	err = fg.SecretClient.Upsert(
		ctx,
		key,
		secret,
		secrets.WithOwner(instance),
		secrets.WithScheme(fg.Scheme),
	)
	if err != nil {
		return false, err
	}

	// create was received successfully but replication is not done
	return false, nil
}

// Delete drops a sqlfailovergroup
func (fg *AzureSqlFailoverGroupManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		fg.SecretClient = options.SecretClient
	}

	instance, err := fg.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	serverName := instance.Spec.Server
	failoverGroupName := instance.ObjectMeta.Name

	// key for Secret to delete on successful provision
	key := types.NamespacedName{Name: instance.ObjectMeta.Name, Namespace: instance.Namespace}

	_, err = fg.DeleteFailoverGroup(ctx, groupName, serverName, failoverGroupName)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureErrorAzureError(err)

		// these errors are expected
		ignore := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.FailoverGroupBusy,
		}

		// this means the thing doesn't exist
		finished := []string{
			errhelp.ResourceNotFound,
		}

		if helpers.ContainsString(ignore, azerr.Type) {
			return true, nil
		}

		if helpers.ContainsString(finished, azerr.Type) {
			// Best case deletion of secret
			fg.SecretClient.Delete(ctx, key)
			return false, nil
		}
		instance.Status.Message = fmt.Sprintf("AzureSqlFailoverGroup Delete failed with: %s", err.Error())
		return false, err
	}

	instance.Status.Message = fmt.Sprintf("Delete AzureSqlFailoverGroup succeeded")
	// Best case deletion of secret
	fg.SecretClient.Delete(ctx, key)
	return false, nil
}

// GetParents returns the parents of sqlfailovergroup
func (fg *AzureSqlFailoverGroupManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := fg.convert(obj)
	if err != nil {
		return nil, err
	}

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

// GetStatus gets the ASOStatus
func (g *AzureSqlFailoverGroupManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (fg *AzureSqlFailoverGroupManager) convert(obj runtime.Object) (*azurev1alpha1.AzureSqlFailoverGroup, error) {
	local, ok := obj.(*azurev1alpha1.AzureSqlFailoverGroup)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
