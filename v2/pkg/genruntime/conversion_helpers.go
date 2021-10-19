/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

/*
This file contains manual implementations to reduce code bloat in generated code.
*/

// CloneSliceOfString clones the given []string. It is used (indirectly)
// by the generator when generating property conversions.
func CloneSliceOfString(slice []string) []string {
	if slice == nil {
		return nil // preserve nils
	}

	result := make([]string, len(slice))
	copy(result, slice)
	return result
}

// CloneMapOfStringToString clones the given map[string]string. It is used
// (indirectly) by the generator when generating property conversions.
func CloneMapOfStringToString(input map[string]string) map[string]string {
	if input == nil {
		return nil // preserve nils
	}

	result := make(map[string]string, len(input))
	for k, v := range input {
		result[k] = v
	}

	return result
}
