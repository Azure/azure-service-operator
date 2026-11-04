/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"strings"

	msgraphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"

	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// reconcilerBase is a reusable foundation for Entra resource reconcilers.
// R is the ASO resource type and E is the corresponding Entra API type.
type reconcilerBase[R genruntime.EntraMetaObject, E msgraphmodels.DirectoryObjectable] struct {
	reconcilers.ReconcilerCommon
}

// escapeODataString escapes a string value for use in an OData filter expression.
// Single quotes in OData are escaped by doubling them.
func (r *reconcilerBase[R, E]) escapeODataString(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
