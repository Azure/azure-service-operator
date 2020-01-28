/*
MIT License

Copyright (c) Microsoft Corporation. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
*/

package apimgmt

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	// "github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimshared"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

// Manager represents an API Management type
type Manager struct {
	Telemetry telemetry.PrometheusTelemetry
	APIClient apimanagement.APIClient
}

// NewManager returns an API Manager type
func NewManager(log logr.Logger) *Manager {
	return &Manager{
		Telemetry: &telemetry.PrometheusClient{},
		APIClient: apimshared.GetAPIMClient(),
	}
}

// CreateAPI creates an API within an API management service
func (m *Manager) CreateAPI(
	ctx context.Context,
	resourceGroupName string,
	apiName string,
	apiServiceName string,
	properties azurev1alpha1.APIProperties,
	eTag string) (*apimanagement.APIContract, error) {

	props := &apimanagement.APICreateOrUpdateProperties{
		APIType:                apimanagement.HTTP,
		APIVersion:             to.StringPtr(properties.APIVersion),
		APIRevision:            to.StringPtr(properties.APIRevision),
		APIRevisionDescription: to.StringPtr(properties.APIRevisionDescription),
		APIVersionDescription:  to.StringPtr(properties.APIVersionDescription),
		APIVersionSetID:        to.StringPtr(properties.APIVersionSetID),
		DisplayName:            to.StringPtr(properties.DisplayName),
		Description:            to.StringPtr(properties.Description),
		IsCurrent:              to.BoolPtr(properties.IsCurrent),
		IsOnline:               to.BoolPtr(properties.IsOnline),
		Path:                   to.StringPtr(properties.Path),
		Format:                 apimanagement.ContentFormat(properties.Format),
	}

	params := apimanagement.APICreateOrUpdateParameter{
		APICreateOrUpdateProperties: props,
	}

	// Fetch the parent API Management service the API will reside under
	svc, err := apimshared.GetAPIMgmtSvc(ctx, resourceGroupName, apiServiceName)
	fmt.Printf("got service reference %v", svc)
	if err != nil {
		// If there is no parent APIM service, we cannot proceed
		m.Telemetry.LogError("failure fetching API management service", err)
		return nil, err
	}

	// Submit the ARM request
	future, err := m.APIClient.CreateOrUpdate(ctx, resourceGroupName, apiName, *svc.ID, params, eTag)

	if err != nil {
		return nil, err
	}

	err = future.WaitForCompletionRef(ctx, m.APIClient.Client)
	if err != nil {
		return nil, err
	}

	result, err := future.Result(m.APIClient)
	return &result, err
}

// DeleteAPI deletes an API within an API management service
func (m *Manager) DeleteAPI(ctx context.Context, resourceGroupName string, apiServiceName string, apiId string, eTag string, deleteRevisions bool) (autorest.Response, error) {
	result, err := m.APIClient.Get(ctx, resourceGroupName, apiServiceName, apiId)
	if err == nil {
		return m.APIClient.Delete(ctx, resourceGroupName, apiServiceName, apiId, eTag, &deleteRevisions)
	}
	return result.Response, nil
}

// GetAPI fetches an API within an API management service
func (m *Manager) GetAPI(ctx context.Context, resourceGroupName string, apiServiceName string, apiId string) (apimanagement.APIContract, error) {
	contract, err := m.APIClient.Get(ctx, resourceGroupName, apiServiceName, apiId)
	return contract, err
}

// Ensure executes a desired state check against the resource
func (m *Manager) Ensure(ctx context.Context, obj runtime.Object) (bool, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	// Set k8s status to provisioning at the beginning of this reconciliation
	instance.Status.Provisioning = true

	// Fetch the API
	api, err := m.GetAPI(ctx, instance.Spec.ResourceGroup, instance.Spec.APIService, instance.Spec.Properties.APIRevision)
	if err == nil {
		instance.Status.State = api.Status

		if api.Status == "Succeeded" {
			instance.Status.Message = api.Status
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			return true, nil
		}

		return false, nil
	}
	return false, nil
}

// Delete removes a resource
func (m *Manager) Delete(ctx context.Context, obj runtime.Object) (bool, error) {
	// i, err := m.convert(obj)
	// if err != nil {
	// 	return false, err
	// }
	// response, err := m.DeleteAPI(ctx, i.Spec.ResourceGroup, i.Spec.APIService, i.Spec.Properties.APIRevision, i.Spec.Properties.APIRevision, true)

	// if err != nil {
	// 	m.Telemetry.LogInfo("Error deleting API", err.Error())
	// 	if !errhelp.IsAsynchronousOperationNotComplete(err) {
	// 		return true, err
	// 	}
	// }
	// i.Status.State = response.Status

	// if err == nil {
	// 	if response.Status != "InProgress" {
	// 		return false, nil
	// 	}
	// }

	return true, nil
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
			Target: &v1alpha1.ApimService{},
		},
	}, nil
}

func (m *Manager) convert(obj runtime.Object) (*v1alpha1.APIMgmt, error) {
	i, ok := obj.(*v1alpha1.APIMgmt)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return i, nil
}
