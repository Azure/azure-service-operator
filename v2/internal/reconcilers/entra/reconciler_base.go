/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"strings"

	msgraphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// reconcilerBase is a reusable foundation for Entra resource reconcilers.
// R is the ASO resource type and E is the corresponding Entra API type.
type reconcilerBase[R genruntime.EntraMetaObject, E msgraphmodels.DirectoryObjectable] struct {
	reconcilers.ReconcilerCommon
}

// asResource asserts that obj is of resource type R.
// If it isn't, the error identifies both the expected and actual types.
func (r *reconcilerBase[R, E]) asResource(obj genruntime.MetaObject) (R, error) {
	typedObj, ok := obj.(R)
	if !ok {
		var zero R
		return zero, eris.Errorf("cannot modify resource that is not of type %T. Type is %T", zero, obj)
	}

	return typedObj, nil
}

// escapeODataString escapes a string value for use in an OData filter expression.
// Single quotes in OData are escaped by doubling them.
func (r *reconcilerBase[R, E]) escapeODataString(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
