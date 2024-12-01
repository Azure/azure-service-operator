// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type AuthorizationProvidersAuthorizationsAccessPolicy_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of the Authorization Contract.
	Properties *AuthorizationAccessPolicyContractProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &AuthorizationProvidersAuthorizationsAccessPolicy_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (policy AuthorizationProvidersAuthorizationsAccessPolicy_Spec) GetAPIVersion() string {
	return "2023-05-01-preview"
}

// GetName returns the Name of the resource
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_Spec) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/authorizationProviders/authorizations/accessPolicies"
func (policy *AuthorizationProvidersAuthorizationsAccessPolicy_Spec) GetType() string {
	return "Microsoft.ApiManagement/service/authorizationProviders/authorizations/accessPolicies"
}

// Authorization Access Policy details.
type AuthorizationAccessPolicyContractProperties struct {
	// AppIds: The allowed Azure Active Directory Application IDs
	AppIds []string `json:"appIds,omitempty"`

	// ObjectId: The Object Id
	ObjectId *string `json:"objectId,omitempty" optionalConfigMapPair:"ObjectId"`

	// TenantId: The Tenant Id
	TenantId *string `json:"tenantId,omitempty" optionalConfigMapPair:"TenantId"`
}
