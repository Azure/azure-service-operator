// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20220131p "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20220131preview"
	v20220131ps "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20220131preview/storage"
	v20230131 "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	v20230131s "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type FederatedIdentityCredentialExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *FederatedIdentityCredentialExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20220131p.FederatedIdentityCredential{},
		&v20220131ps.FederatedIdentityCredential{},
		&v20230131.FederatedIdentityCredential{},
		&v20230131s.FederatedIdentityCredential{}}
}
