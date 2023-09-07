/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// TypeAssociation defines an association between two internal types
type TypeAssociation map[InternalTypeName]InternalTypeName

func (ts TypeAssociation) Equals(other TypeAssociation) bool {
	if len(ts) != len(other) {
		// Different sizes, not equal
		return false
	}

	for k, otherVal := range other {
		val, ok := ts[k]
		if !ok {
			// Missing key, not equal
			return false
		}

		if !TypeEquals(val, otherVal) {
			// Values don't match, not equal
			return false
		}
	}

	return true
}
