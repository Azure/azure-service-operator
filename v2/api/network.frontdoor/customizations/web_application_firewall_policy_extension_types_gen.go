// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20220501 "github.com/Azure/azure-service-operator/v2/api/network.frontdoor/v1api20220501"
	v20220501s "github.com/Azure/azure-service-operator/v2/api/network.frontdoor/v1api20220501/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type WebApplicationFirewallPolicyExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *WebApplicationFirewallPolicyExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20220501.WebApplicationFirewallPolicy{},
		&v20220501s.WebApplicationFirewallPolicy{}}
}
