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

	"github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/go-autorest/autorest"
)

// APIManager is an interface for API Management
type APIManager interface {
	CreateAPI(ctx context.Context, resourceGroupName string, apiServiceName string, apiID string, properties azurev1alpha1.APIProperties, eTag string) (apimanagement.APIContract, error)
	DeleteAPI(ctx context.Context, resourceGroupName string, apiServiceName string, apiID string, eTag string, deleteRevisions bool) (autorest.Response, error)
	GetAPI(ctx context.Context, resourceGroupName string, apiServiceName string, apiID string) (apimanagement.APIContract, error)
	resourcemanager.ARMClient
}
