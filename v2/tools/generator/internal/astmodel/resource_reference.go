/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

var arrayResourceReferenceType = NewArrayType(ResourceReferenceType)
var mapResourceReferenceType = NewMapType(StringType, ResourceReferenceType)

func IsTypeResourceReference(t Type) bool {
	isResourceReference := TypeEquals(t, ResourceReferenceType)
	isOptionalResourceReference := TypeEquals(t, OptionalResourceReferenceType)
	isSliceResourceReference := IsTypeResourceReferenceSlice(t)
	isMapResourceReference := IsTypeResourceReferenceMap(t)

	return isResourceReference || isOptionalResourceReference || isSliceResourceReference || isMapResourceReference
}

func IsTypeResourceReferenceSlice(t Type) bool {
	return TypeEquals(t, arrayResourceReferenceType)
}

func IsTypeResourceReferenceMap(t Type) bool {
	return TypeEquals(t, mapResourceReferenceType)
}
