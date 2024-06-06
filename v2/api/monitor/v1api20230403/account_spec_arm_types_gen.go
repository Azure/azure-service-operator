// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230403

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Account_Spec_ARM struct {
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Account_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-04-03"
func (account Account_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (account *Account_Spec_ARM) GetName() string {
	return account.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Monitor/accounts"
func (account *Account_Spec_ARM) GetType() string {
	return "Microsoft.Monitor/accounts"
}