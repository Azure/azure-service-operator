/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

func IsTypeResourceReference(t Type) bool {
	// Handles optional too
	if ex, ok := AsExternalTypeName(t); ok {
		return TypeEquals(ex, ResourceReferenceType) ||
			TypeEquals(ex, WellknownResourceReferenceType)
	}

	if IsTypeResourceReferenceSlice(t) || IsTypeResourceReferenceMap(t) {
		return true
	}

	return false
}

func IsTypeResourceReferenceSlice(t Type) bool {
	if sl, ok := AsArrayType(t); ok {
		return TypeEquals(sl.Element(), ResourceReferenceType) ||
			TypeEquals(sl.Element(), WellknownResourceReferenceType)
	}

	return false
}

func IsTypeResourceReferenceMap(t Type) bool {
	if mp, ok := AsMapType(t); ok {
		return TypeEquals(mp.ValueType(), ResourceReferenceType) ||
			TypeEquals(mp.ValueType(), WellknownResourceReferenceType)
	}

	return false
}
