/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

// appendWithPrefix adds a set of strings to an existing slice, prefixing each with a specified string
func appendWithPrefix(slice []string, prefix string, values ...string) []string {
	result := slice
	for _, v := range values {
		result = append(result, prefix+v)
	}

	return result
}
