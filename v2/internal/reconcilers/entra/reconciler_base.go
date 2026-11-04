/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"strings"

	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
)

// reconcilerBase is a reusable foundation for Entra resource reconcilers.
type reconcilerBase struct {
	reconcilers.ReconcilerCommon
}

// escapeODataString escapes a string value for use in an OData filter expression.
// Single quotes in OData are escaped by doubling them.
func (r *reconcilerBase) escapeODataString(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
