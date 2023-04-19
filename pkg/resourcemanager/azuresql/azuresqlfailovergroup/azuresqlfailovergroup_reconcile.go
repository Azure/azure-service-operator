// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlfailovergroup

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/pollclient"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

// Ensure creates a sqlfailovergroup
func (fg *AzureSqlFailoverGroupManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	secretClient := fg.SecretClient
	if options.SecretClient != nil {
		secretClient = options.SecretClient
	}

	instance, err := fg.convert(obj)
	if err != nil {
		return false, err
	}

	pClient := pollclient.NewPollClient(fg.Creds)
	lroPollResult, err := pClient.PollLongRunningOperationIfNeeded(ctx, &instance.Status, api.PollingURLKindCreateOrUpdate)
	if err != nil {
		instance.Status.Message = err.Error()
		return false, err
	}
	if lroPollResult == pollclient.PollResultTryAgainLater {
		// Need to wait a bit before trying again
		return false, nil
	}
	if lroPollResult == pollclient.PollResultBadRequest {
		// Reached a terminal state
		instance.Status.SetFailedProvisioning(instance.Status.Message)
		return true, nil
	}

	failoverGroupsClient, err := azuresqlshared.GetGoFailoverGroupsClient(azuresqlshared.GetSubscriptionCredentials(fg.Creds, instance.Spec.SubscriptionID))
	if err != nil {
		return false, errors.Wrapf(err, "failed to create failovergroup client")
	}

	notFoundErrorCodes := []string{
		errhelp.ParentNotFoundErrorCode,
		errhelp.ResourceGroupNotFoundErrorCode,
		errhelp.NotFoundErrorCode,
		errhelp.ResourceNotFound,
	}

	failoverGroupName := instance.ObjectMeta.Name
	azureFailoverGroup, err := failoverGroupsClient.Get(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, failoverGroupName)
	if err != nil {
		azerr := errhelp.NewAzureError(err)
		if azerr.Type != errhelp.ResourceNotFound {
			instance.Status.Message = fmt.Sprintf("AzureSqlFailoverGroup Get error %s", err.Error())
			return errhelp.IsErrorFatal(err, notFoundErrorCodes, nil)
		}
	}

	failoverGroupProperties, err := fg.TransformToSQLFailoverGroup(ctx, instance)
	if err != nil {
		instance.Status.Message = err.Error()
		return errhelp.IsErrorFatal(err, notFoundErrorCodes, nil)
	}

	// We found a failover group, check to make sure that it matches what we have locally
	resourceMatchesAzure := DoesResourceMatchAzure(failoverGroupProperties, azureFailoverGroup)

	if resourceMatchesAzure {
		if *azureFailoverGroup.ReplicationState == "SEEDING" {
			instance.Status.Message = "FailoverGroup replicationState is SEEDING"
			return false, nil
		}

		secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
		_, err := secretClient.Get(ctx, secretKey)
		// We make the same assumption many other places in the code make which is that if we cannot
		// get the secret it must not exist.
		if err != nil {
			// Create a new secret
			secret := fg.NewSecret(instance)

			// create or update the secret
			err = secretClient.Upsert(
				ctx,
				secretKey,
				secret,
				secrets.WithOwner(instance),
				secrets.WithScheme(fg.Scheme),
			)
			if err != nil {
				instance.Status.Message = fmt.Sprintf("failed to update secret: %s", err.Error())
				return false, err
			}
		}

		instance.Status.SetProvisioned(resourcemanager.SuccessMsg)
		instance.Status.ResourceId = *azureFailoverGroup.ID
		return true, nil
	}

	// We need to reconcile as our state didn't match Azure
	instance.Status.SetProvisioning("")

	future, err := fg.CreateOrUpdateFailoverGroup(
		ctx,
		instance.Spec.SubscriptionID,
		instance.Spec.ResourceGroup,
		instance.Spec.Server,
		failoverGroupName,
		failoverGroupProperties)
	if err != nil {
		instance.Status.Message = err.Error()
		allowedErrors := []string{
			errhelp.AsyncOpIncompleteError,
			errhelp.AlreadyExists,
			errhelp.FailoverGroupBusy,
		}
		unrecoverableErrors := []string{
			errhelp.InvalidFailoverGroupRegion,
		}
		return errhelp.IsErrorFatal(err, append(allowedErrors, notFoundErrorCodes...), unrecoverableErrors)
	}
	instance.Status.SetProvisioning("Resource request successfully submitted to Azure")
	instance.Status.SetPollingURL(future.PollingURL(), api.PollingURLKindCreateOrUpdate)

	// Need to poll the polling URL, so not done yet!
	return false, nil
}

// Delete drops a sqlfailovergroup
func (fg *AzureSqlFailoverGroupManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	secretClient := fg.SecretClient
	if options.SecretClient != nil {
		secretClient = options.SecretClient
	}

	instance, err := fg.convert(obj)
	if err != nil {
		return false, err
	}

	groupName := instance.Spec.ResourceGroup
	serverName := instance.Spec.Server
	failoverGroupName := instance.ObjectMeta.Name

	// key for Secret to delete on successful provision
	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}

	_, err = fg.DeleteFailoverGroup(ctx, instance.Spec.SubscriptionID, groupName, serverName, failoverGroupName)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureError(err)

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
			secretClient.Delete(ctx, secretKey)
			return false, nil
		}
		instance.Status.Message = fmt.Sprintf("AzureSqlFailoverGroup Delete failed with: %s", err.Error())
		return false, err
	}

	instance.Status.Message = fmt.Sprintf("Delete AzureSqlFailoverGroup succeeded")
	// Best case deletion of secret
	secretClient.Delete(ctx, secretKey)
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
	st := azurev1alpha1.ASOStatus(instance.Status)

	return &st, nil
}

func (fg *AzureSqlFailoverGroupManager) convert(obj runtime.Object) (*v1beta1.AzureSqlFailoverGroup, error) {
	local, ok := obj.(*v1beta1.AzureSqlFailoverGroup)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
