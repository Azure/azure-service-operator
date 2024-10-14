// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

type ServersDatabasesTransparentDataEncryption_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *TransparentDataEncryptionProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Properties of a transparent data encryption.
type TransparentDataEncryptionProperties_STATUS_ARM struct {
	// State: Specifies the state of the transparent data encryption.
	State *TransparentDataEncryptionProperties_State_STATUS_ARM `json:"state,omitempty"`
}

type TransparentDataEncryptionProperties_State_STATUS_ARM string

const (
	TransparentDataEncryptionProperties_State_STATUS_ARM_Disabled = TransparentDataEncryptionProperties_State_STATUS_ARM("Disabled")
	TransparentDataEncryptionProperties_State_STATUS_ARM_Enabled  = TransparentDataEncryptionProperties_State_STATUS_ARM("Enabled")
)

// Mapping from string to TransparentDataEncryptionProperties_State_STATUS_ARM
var transparentDataEncryptionProperties_State_STATUS_ARM_Values = map[string]TransparentDataEncryptionProperties_State_STATUS_ARM{
	"disabled": TransparentDataEncryptionProperties_State_STATUS_ARM_Disabled,
	"enabled":  TransparentDataEncryptionProperties_State_STATUS_ARM_Enabled,
}
