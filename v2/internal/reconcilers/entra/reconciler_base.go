/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package entra

import (
	"context"
	"fmt"
	"strings"

	msgraphmodels "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
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

type adoptionAllower interface {
	comparable
	AdoptionAllowed() bool
}

// canAdopt checks whether spec allows adoption. A nil spec allows adoption by default.
func (r *reconcilerBase[R, E]) canAdopt[P adoptionAllower](spec P) bool {
	var nilSpec P
	if spec == nilSpec {
		return true
	}

	return spec.AdoptionAllowed()
}

type creationAllower interface {
	comparable
	CreationAllowed() bool
}

// canCreate checks whether spec allows creation. A nil spec allows creation by default.
func (r *reconcilerBase[R, E]) canCreate[P creationAllower](spec P) bool {
	var nilSpec P
	if spec == nilSpec {
		return true
	}

	return spec.CreationAllowed()
}

// loadByID loads an Entra resource by ID using fetch.
// A 404 returns the zero value of E without an error; other results and errors are returned unchanged.
func (r *reconcilerBase[R, E]) loadByID(
	ctx context.Context,
	id string,
	fetch func(context.Context, string) (E, error),
) (E, error) {
	resource, err := fetch(ctx, id)
	if isNotFound(err) {
		var zero E
		return zero, nil
	}

	return resource, err
}

// query finds Entra resources whose field equals value, collecting results from every page.
// fetchPage receives the escaped filter and next-link URL, which is empty for the first page.
// readPage extracts the resources and next link from a response. Any failure discards partial results.
func (r *reconcilerBase[R, E]) query[C comparable](
	ctx context.Context,
	field string,
	value string,
	fetchPage func(context.Context, string, string) (C, error),
	readPage func(C) ([]E, *string),
) ([]E, error) {
	filter := fmt.Sprintf("%s eq '%s'", field, r.escapeODataString(value))
	var resource R
	var nilPage C
	var results []E
	nextLink := ""
	for {
		page, err := fetchPage(ctx, filter, nextLink)
		if err != nil {
			return nil, eris.Wrapf(err, "searching %T by %s %q", resource, field, value)
		}

		if page == nilPage {
			return nil, eris.Errorf("searching %T by %s %q returned no response", resource, field, value)
		}

		values, next := readPage(page)
		results = append(results, values...)

		nextLink = to.Value(next)
		if nextLink == "" {
			return results, nil
		}
	}
}

// escapeODataString escapes a string value for use in an OData filter expression.
// Single quotes in OData are escaped by doubling them.
func (r *reconcilerBase[R, E]) escapeODataString(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
