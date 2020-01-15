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
	"log"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

var apimClient apimanagement.APIClient

// Manager represents an API Management type
type Manager struct {
	Log logr.Logger
}

// func init() {
// 	apimClient = getAPIMClient()
// }

// NewManager returns an API Manager type
func NewManager(log logr.Logger) *Manager {
	return &Manager{
		Log: log,
	}
}

// CreateAPI creates an API within an API management service
func (m *Manager) CreateAPI(ctx context.Context, resourcegroup string, apiname string, properties string) (string, error) {
	return "", nil
}

// DeleteAPI deletes an API within an API management service
func (m *Manager) DeleteAPI(ctx context.Context, resourcegroup string, apiname string) (string, error) {
	return "", nil
}

// GetAPI fetches an API within an API management service
func (m *Manager) GetAPI(ctx context.Context, resourcegroup string, servername string, database string) (string, error) {
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
func (m *Manager) GetParents(runtime.Object) ([]resourcemanager.KubeParent, error) {
	return []resourcemanager.KubeParent{}, nil
}

func getAPIMClient() apimanagement.APIClient {
	apimClient := apimanagement.NewAPIClient(config.SubscriptionID())

	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Fatalf("failed to initialize authorizer %v\n", err)
	}
	apimClient.Authorizer = a
	apimClient.AddToUserAgent(config.UserAgent())

	return apimClient
}
