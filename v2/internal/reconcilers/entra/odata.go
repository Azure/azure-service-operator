/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import "strings"

// escapeODataString escapes a string value for use in an OData filter expression.
// Single quotes in OData are escaped by doubling them
func escapeODataString(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
