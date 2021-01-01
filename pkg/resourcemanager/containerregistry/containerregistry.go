// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package containerregistry

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type azureContainerRegistryManager struct {
	Scheme       *runtime.Scheme
	Creds        config.Credentials
	SecretClient secrets.SecretClient
}

func (a azureContainerRegistryManager) CreateRegistry(ctx context.Context, instance *azurev1alpha1.AzureContainerRegistry) (containerregistry.Registry, error) {
	containerRegistryClient, err := getContainerRegistryClient(a.Creds)
	if err != nil {
		return containerregistry.Registry{}, err
	}

	containerRegistryNameCheck := containerregistry.RegistryNameCheckRequest{
		Name: to.StringPtr(instance.Name),
		Type: to.StringPtr("Microsoft.ContainerRegistry/registries"),
	}
	result, err := containerRegistryClient.CheckNameAvailability(ctx, containerRegistryNameCheck)
	if err != nil {
		return containerregistry.Registry{}, err
	}
	if !*result.NameAvailable {
		return containerregistry.Registry{}, fmt.Errorf(*result.Reason)
	}

	registry := toRegistryType(instance)

	future, err := containerRegistryClient.Create(ctx, instance.Spec.ResourceGroup, instance.Name, *registry)
	if err != nil {
		return containerregistry.Registry{}, err
	}

	return future.Result(containerRegistryClient)
}

func (a azureContainerRegistryManager) UpdateRegistry(ctx context.Context, instance *azurev1alpha1.AzureContainerRegistry) (containerregistry.Registry, error) {
	containerRegistryClient, err := getContainerRegistryClient(a.Creds)
	if err != nil {
		return containerregistry.Registry{}, err
	}

	registryUpdateParams := toRegistryUpdateParams(instance)

	future, err := containerRegistryClient.Update(ctx, instance.Spec.ResourceGroup, instance.Name, *registryUpdateParams)
	if err != nil {
		return containerregistry.Registry{}, err
	}

	return future.Result(containerRegistryClient)
}

func (a azureContainerRegistryManager) DeleteRegistry(ctx context.Context, groupName string, registryName string) (result autorest.Response, err error) {
	containerRegistryClient, err := getContainerRegistryClient(a.Creds)
	if err != nil {
		return autorest.Response{}, err
	}

	future, err := containerRegistryClient.Delete(ctx, groupName, registryName)
	if err != nil {
		return autorest.Response{}, err
	}

	return future.Result(containerRegistryClient)
}

func (a azureContainerRegistryManager) GetRegistry(ctx context.Context, groupName string, registryName string) (result containerregistry.Registry, err error) {
	containerRegistryClient, err := getContainerRegistryClient(a.Creds)
	if err != nil {
		return containerregistry.Registry{}, err
	}
	return containerRegistryClient.Get(ctx, groupName, registryName)
}

func (a azureContainerRegistryManager) ListCredentials(ctx context.Context, groupName string, registryName string) (result containerregistry.RegistryListCredentialsResult, err error) {
	containerRegistryClient, err := getContainerRegistryClient(a.Creds)
	if err != nil {
		return containerregistry.RegistryListCredentialsResult{}, err
	}
	return containerRegistryClient.ListCredentials(ctx, groupName, registryName)
}

func getContainerRegistryClient(creds config.Credentials) (containerregistry.RegistriesClient, error) {
	containerRegistriesClient := containerregistry.NewRegistriesClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return containerRegistriesClient, err
	}
	containerRegistriesClient.Authorizer = a
	containerRegistriesClient.AddToUserAgent(config.UserAgent())
	return containerRegistriesClient, nil
}

func (a azureContainerRegistryManager) Ensure(ctx context.Context, object runtime.Object, option ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := a.convert(object)
	if err != nil {
		return true, err
	}

	// hash the spec
	hash := helpers.Hash256(instance.Spec)

	instance.Status.Provisioning = true
	instance.Status.FailedProvisioning = false
	exists := false
	containerRegistry, err := a.GetRegistry(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		exists = true
		if instance.Status.SpecHash == hash {
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			instance.Status.ResourceId = *containerRegistry.ID
			return true, nil
		}
		instance.Status.SpecHash = hash
		instance.Status.ContainsUpdate = true
	}

	if !exists {
		containerRegistry, err = a.CreateRegistry(ctx, instance)
	} else {
		containerRegistry, err = a.UpdateRegistry(ctx, instance)
	}

	if err != nil {
		return HandleCreationError(instance, err)
	}

	instance.Status.State = containerRegistry.Response.Status
	if containerRegistry.ID != nil {
		instance.Status.ResourceId = *containerRegistry.ID
	}
	instance.Status.ContainsUpdate = false
	instance.Status.Provisioned = true
	instance.Status.Provisioning = false
	instance.Status.Message = resourcemanager.SuccessMsg

	// upsert secrets if they exist
	err = a.StoreSecrets(ctx, instance.Spec.ResourceGroup, instance.Name, instance)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a azureContainerRegistryManager) StoreSecrets(ctx context.Context, groupName string, registryName string, instance *azurev1alpha1.AzureContainerRegistry) error {
	credentials, err := a.ListCredentials(ctx, groupName, registryName)
	if err != nil {
		return fmt.Errorf("error retrieving credentials whilst trying to write them to secret %v", err)
	}

	// We don't always have credentials if admin_enabled is set to false so gracefully handle missing usernames/passwords
	if credentials.Username == nil || credentials.Passwords == nil || *credentials.Username == "" || len(*credentials.Passwords) == 0 {
		return nil
	}

	key := types.NamespacedName{
		Name:      fmt.Sprintf("containerregistry-%s", registryName),
		Namespace: instance.Namespace,
	}

	data := map[string][]byte{
		"username": []byte(*credentials.Username),
	}

	for i, password := range *credentials.Passwords {
		if i == 0 {
			data["password"] = []byte(*password.Value)
		} else {
			data[fmt.Sprintf("password%v", i)] = []byte(*password.Value)
		}
	}

	return a.SecretClient.Upsert(ctx,
		key,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(a.Scheme),
	)
}

func (a azureContainerRegistryManager) Delete(ctx context.Context, object runtime.Object, option ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := a.convert(object)
	if err != nil {
		return true, err
	}

	_, err = a.GetRegistry(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		_, err = a.DeleteRegistry(ctx, instance.Spec.ResourceGroup, instance.Name)
		if err != nil {
			catch := []string{
				errhelp.AsyncOpIncompleteError,
			}
			gone := []string{
				errhelp.ResourceGroupNotFoundErrorCode,
				errhelp.ParentNotFoundErrorCode,
				errhelp.NotFoundErrorCode,
				errhelp.ResourceNotFound,
			}
			azerr := errhelp.NewAzureError(err)
			if helpers.ContainsString(catch, azerr.Type) {
				return true, nil
			} else if helpers.ContainsString(gone, azerr.Type) {
				return false, nil
			}
			return true, err
		}
		return true, nil
	}

	return false, nil
}

func (a azureContainerRegistryManager) GetParents(object runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := a.convert(object)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil
}

func (a azureContainerRegistryManager) GetStatus(obj runtime.Object) (*azurev1alpha1.ASOStatus, error) {
	instance, err := a.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (a azureContainerRegistryManager) convert(object runtime.Object) (*azurev1alpha1.AzureContainerRegistry, error) {
	local, ok := object.(*azurev1alpha1.AzureContainerRegistry)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", object.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func HandleCreationError(instance *azurev1alpha1.AzureContainerRegistry, err error) (bool, error) {
	// let the user know what happened
	instance.Status.Message = errhelp.StripErrorTimes(errhelp.StripErrorIDs(err))
	instance.Status.Provisioning = false
	// errors we expect might happen that we are ok with waiting for
	catch := []string{
		errhelp.ResourceGroupNotFoundErrorCode,
		errhelp.ParentNotFoundErrorCode,
		errhelp.NotFoundErrorCode,
		errhelp.AsyncOpIncompleteError,
	}

	catchUnrecoverableErrors := []string{
		errhelp.AccountNameInvalid,
		errhelp.AlreadyExists,
		errhelp.InvalidAccessPolicy,
		errhelp.BadRequest,
		errhelp.LocationNotAvailableForResourceType,
	}

	azerr := errhelp.NewAzureError(err)
	if helpers.ContainsString(catch, azerr.Type) {
		// most of these error technically mean the resource is actually not provisioning
		switch azerr.Type {
		case errhelp.AsyncOpIncompleteError:
			instance.Status.Provisioning = true
		}
		// reconciliation is not done but error is acceptable
		return false, nil
	}

	if helpers.ContainsString(catchUnrecoverableErrors, azerr.Type) {
		// Unrecoverable error, so stop reconcilation
		switch azerr.Type {
		case errhelp.AlreadyExists:
			timeNow := metav1.NewTime(time.Now())
			if timeNow.Sub(instance.Status.RequestedAt.Time) < (30 * time.Second) {
				instance.Status.Provisioning = true
				return false, nil
			}

		}
		instance.Status.Message = "Reconcilation hit unrecoverable error " + err.Error()
		return true, nil
	}

	if azerr.Code == http.StatusForbidden {
		// permission errors when applying access policies are generally worth waiting on
		return false, nil
	}

	// reconciliation not done and we don't know what happened
	return false, err
}

func toRegistryType(instance *azurev1alpha1.AzureContainerRegistry) *containerregistry.Registry {
	return &containerregistry.Registry{
		Sku: &containerregistry.Sku{
			Name: containerregistry.SkuName(instance.Spec.Sku.Name),
		},
		RegistryProperties: &containerregistry.RegistryProperties{
			AdminUserEnabled: to.BoolPtr(instance.Spec.AdminUserEnabled),
		},
		Location: to.StringPtr(instance.Spec.Location),
	}
}

func toRegistryUpdateParams(instance *azurev1alpha1.AzureContainerRegistry) *containerregistry.RegistryUpdateParameters {
	return &containerregistry.RegistryUpdateParameters{
		Sku: &containerregistry.Sku{
			Name: containerregistry.SkuName(instance.Spec.Sku.Name),
		},
		RegistryPropertiesUpdateParameters: &containerregistry.RegistryPropertiesUpdateParameters{
			AdminUserEnabled: to.BoolPtr(instance.Spec.AdminUserEnabled),
		},
	}
}
