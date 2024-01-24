/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

var arraySecretReference = NewArrayType(SecretReferenceType)

func IsTypeSecretReference(t Type) bool {
	isSecretReference := TypeEquals(t, SecretReferenceType)
	isOptionalSecretReference := TypeEquals(t, OptionalSecretReferenceType)
	isSliceSecretReference := IsTypeSecretReferenceSlice(t)
	isMapSecretReference := IsTypeSecretReferenceMap(t)

	return isSecretReference || isOptionalSecretReference || isSliceSecretReference || isMapSecretReference
}

func IsTypeSecretReferenceSlice(t Type) bool {
	return TypeEquals(t, arraySecretReference)
}

func IsTypeSecretReferenceMap(t Type) bool {
	return TypeEquals(t, SecretMapReferenceType) || TypeEquals(t, OptionalSecretMapReferenceType)
}
