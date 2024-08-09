/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// These are some magical field names which we're going to use or generate
const (
	APIVersionProperty                       = "APIVersion" // Used by armconversion package
	AzureNameProperty                        = "AzureName"
	NameProperty                             = "Name" // Used by armconversion package
	OwnerProperty                            = "Owner"
	SetAzureNameFunc                         = "SetAzureName"
	TypeProperty                             = "Type" // Used by armconversion package
	OperatorSpecProperty                     = "OperatorSpec"
	OperatorSpecSecretsProperty              = "Secrets"
	OperatorSpecSecretExpressionsProperty    = "SecretExpressions"
	OperatorSpecConfigMapsProperty           = "ConfigMaps"
	OperatorSpecConfigMapExpressionsProperty = "ConfigMapExpressions"
	ConditionsProperty                       = "Conditions"
	OptionalConfigMapReferenceSuffix         = "FromConfig"
	UserAssignedIdentitiesProperty           = "UserAssignedIdentities"
	UserAssignedIdentitiesTypeName           = "UserAssignedIdentityDetails"
)

// IsKubernetesResourceProperty returns true if the supplied property name is one of the properties required by the
// KubernetesResource interface.
func IsKubernetesResourceProperty(name PropertyName) bool {
	return name == AzureNameProperty || name == OwnerProperty
}

func IsUserAssignedIdentityProperty(prop *PropertyDefinition) (InternalTypeName, bool) {
	if !prop.HasName(UserAssignedIdentitiesProperty) {
		return InternalTypeName{}, false
	}

	arrayType, isArray := prop.PropertyType().(*ArrayType)
	if !isArray {
		return InternalTypeName{}, false
	}

	typeName, ok := AsInternalTypeName(arrayType.Element())
	if !ok {
		return InternalTypeName{}, false
	}

	if typeName.Name() != UserAssignedIdentitiesTypeName {
		return InternalTypeName{}, false
	}

	return typeName, true
}
