/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package genruntime

/*
This file contains manual implementations to reduce code bloat in generated code.
*/

func CloneSliceOfString(slice []string) []string {
	if slice == nil {
		return nil // preserve nils
	}

	result := make([]string, len(slice))
	copy(result, slice)
	return result
}

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
