// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package apimgmt

import (
	"context"
	"fmt"
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimshared"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

// Manager represents an API Management type
type Manager struct{}

// NewManager returns an API Manager type
func NewManager() *Manager {
	return &Manager{}
}

// CreateAPI creates an API within an API management service
func (m *Manager) CreateAPI(
	ctx context.Context,
	resourceGroupName string,
	apiServiceName string,
	apiID string,
	properties azurev1alpha1.APIProperties,
	eTag string) (apimanagement.APIContract, error) {

	props := &apimanagement.APICreateOrUpdateProperties{
		APIType:                apimanagement.HTTP,
		APIVersion:             to.StringPtr(properties.APIVersion),
		APIRevision:            to.StringPtr(properties.APIRevision),
		APIRevisionDescription: to.StringPtr(properties.APIRevisionDescription),
		APIVersionDescription:  to.StringPtr(properties.APIVersionDescription),
		DisplayName:            to.StringPtr(properties.DisplayName),
		Description:            to.StringPtr(properties.Description),
		IsCurrent:              to.BoolPtr(properties.IsCurrent),
		IsOnline:               to.BoolPtr(properties.IsOnline),
		Path:                   to.StringPtr(properties.Path),
		Protocols:              &[]apimanagement.Protocol{"http"},
		Format:                 apimanagement.ContentFormat(properties.Format),
	}

	params := apimanagement.APICreateOrUpdateParameter{
		APICreateOrUpdateProperties: props,
	}

	// Fetch the parent API Management service the API will reside under
	svc, err := apimshared.GetAPIMgmtSvc(ctx, resourceGroupName, apiServiceName)
	if err != nil {
		// If there is no parent APIM service, we cannot proceed
		return apimanagement.APIContract{}, err
	}

	apiClient, err := apimshared.GetAPIMClient()
	if err != nil {
		return apimanagement.APIContract{}, err
	}

	// Submit the ARM request
	future, err := apiClient.CreateOrUpdate(ctx, resourceGroupName, *svc.Name, apiID, params, eTag)

	if err != nil {
		return apimanagement.APIContract{}, err
	}

	return future.Result(apiClient)
}

// DeleteAPI deletes an API within an API management service
func (m *Manager) DeleteAPI(ctx context.Context, resourceGroupName string, apiServiceName string, apiID string, eTag string, deleteRevisions bool) (autorest.Response, error) {
	apiClient, err := apimshared.GetAPIMClient()
	if err != nil {
		return autorest.Response{
			Response: &http.Response{
				StatusCode: 500,
			},
		}, err
	}
	result, err := apiClient.Get(ctx, resourceGroupName, apiServiceName, apiID)
	if err == nil {
		return apiClient.Delete(ctx, resourceGroupName, apiServiceName, apiID, eTag, &deleteRevisions)
	}
	return result.Response, err
}

// GetAPI fetches an API within an API management service
func (m *Manager) GetAPI(ctx context.Context, resourceGroupName string, apiServiceName string, apiID string) (apimanagement.APIContract, error) {
	apiClient, err := apimshared.GetAPIMClient()
	if err != nil {
		return apimanagement.APIContract{}, err
	}
	contract, err := apiClient.Get(ctx, resourceGroupName, apiServiceName, apiID)
	return contract, err
}

// Ensure executes a desired state check against the resource
func (m *Manager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	// Attempt to fetch the parent API Management service the API will or does reside within
	svc, err := apimshared.GetAPIMgmtSvc(ctx, instance.Spec.ResourceGroup, instance.Spec.APIService)
	if err != nil {
		instance.Status.Message = err.Error()
		// If there is no parent APIM service, we cannot proceed
		instance.Status.Provisioning = false
		return false, nil
	}

	// Attempt to fetch the API
	api, err := m.GetAPI(ctx, instance.Spec.ResourceGroup, instance.Spec.APIService, instance.Spec.APIId)
	if err == nil {
		if api.StatusCode == 200 {
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			return true, nil
		}
	}

	// Submit the ARM request
	contract, err := m.CreateAPI(
		ctx,
		instance.Spec.ResourceGroup,
		*svc.Name,
		instance.Spec.APIId,
		azurev1alpha1.APIProperties{
			APIVersion:             instance.Spec.Properties.APIVersion,
			APIRevision:            instance.Spec.Properties.APIRevision,
			APIRevisionDescription: instance.Spec.Properties.APIRevisionDescription,
			APIVersionDescription:  instance.Spec.Properties.APIVersionDescription,
			DisplayName:            instance.Spec.Properties.DisplayName,
			Description:            instance.Spec.Properties.Description,
			IsCurrent:              instance.Spec.Properties.IsCurrent,
			IsOnline:               instance.Spec.Properties.IsOnline,
			Path:                   instance.Spec.Properties.Path,
			Format:                 instance.Spec.Properties.Format,
		},
		instance.Spec.Properties.APIRevision,
	)

	if err != nil {
		// Set the Message in the case where an unexpected error is returned
		instance.Status.Message = err.Error()

		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			instance.Status.Message = err.Error()
			return false, nil
		}

		instance.Status.Provisioning = false

		return false, err
	}

	if contract.StatusCode == 200 {
		instance.Status.Provisioning = false
		instance.Status.Provisioned = true
		instance.Status.FailedProvisioning = false
		instance.Status.Message = resourcemanager.SuccessMsg
	} else {
		instance.Status.Provisioning = false
		instance.Status.Provisioned = false
		instance.Status.FailedProvisioning = true
	}

	return true, nil
}

// Delete removes an API resource
func (m *Manager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	i, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	_, err = m.DeleteAPI(ctx, i.Spec.ResourceGroup, i.Spec.APIService, i.Spec.APIId, i.Spec.Properties.APIRevision, true)
	if err != nil {
		i.Status.Message = err.Error()

		azerr := errhelp.NewAzureErrorAzureError(err)
		handle := []string{
			errhelp.ResourceNotFound,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		if helpers.ContainsString(handle, azerr.Type) {
			return false, nil
		}

		return true, err
	}
	return false, nil
}

// GetParents fetches the hierarchical parent resource references
func (m *Manager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	i, err := m.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: i.Namespace,
				Name:      i.Spec.ResourceGroup,
			},
			Target: &v1alpha1.ResourceGroup{},
		},
		{
			Key: types.NamespacedName{
				Namespace: i.Namespace,
				Name:      i.Spec.APIService,
			},
		},
	}, nil
}

// GetStatus returns the current status of the resource
func (m *Manager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (m *Manager) convert(obj runtime.Object) (*v1alpha1.APIMgmtAPI, error) {
	i, ok := obj.(*v1alpha1.APIMgmtAPI)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return i, nil
}
