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
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimshared"
	"github.com/Azure/go-autorest/autorest/to"
)

// Manager represents an API Management type
type Manager struct {
	Log       logr.Logger
	APIClient apimanagement.APIClient
}

// NewManager returns an API Manager type
func NewManager(log logr.Logger) *Manager {
	return &Manager{
		Log:       log,
		APIClient: apimshared.GetAPIMClient(),
	}
}

// CreateAPI creates an API within an API management service
func (m *Manager) CreateAPI(
	ctx context.Context,
	resourceGroupName string,
	apiServiceName string,
	apiId string,
	properties v1alpha1.APIProperties,
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

	future, err := m.APIClient.CreateOrUpdate(ctx, resourceGroupName, apiServiceName, apiId, params, eTag)

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
func (m *Manager) DeleteAPI(ctx context.Context, resourcegroup string, apiName string) (string, error) {
	return "", nil
}

// GetAPI fetches an API within an API management service
func (m *Manager) GetAPI(ctx context.Context, resourcegroup string, apiName string) (string, error) {
	return "", nil
}

// Ensure executes a desired state check against the resource
func (m *Manager) Ensure(context.Context, runtime.Object) (bool, error) {
	return true, nil
}

// Delete removes a resource
func (m *Manager) Delete(context.Context, runtime.Object) (bool, error) {
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
