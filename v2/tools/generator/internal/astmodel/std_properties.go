/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// These are some magical field names which we're going to use or generate
const (
	APIVersionProperty               = "APIVersion" // Used by armconversion package
	AzureNameProperty                = "AzureName"
	NameProperty                     = "Name" // Used by armconversion package
	OwnerProperty                    = "Owner"
	SetAzureNameFunc                 = "SetAzureName"
	TypeProperty                     = "Type" // Used by armconversion package
	ScopeProperty                    = "Scope"
	OperatorSpecProperty             = "OperatorSpec"
	OperatorSpecSecretsProperty      = "Secrets"
	OperatorSpecConfigMapsProperty   = "ConfigMaps"
	ConditionsProperty               = "Conditions"
	OptionalConfigMapReferenceSuffix = "ConfigRef"
)

// IsKubernetesResourceProperty returns true if the supplied property name is one of the properties required by the
// KubernetesResource interface.
func IsKubernetesResourceProperty(name PropertyName) bool {
	return name == AzureNameProperty || name == OwnerProperty
}
