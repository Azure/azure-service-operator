/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"context"
	"strings"

	api "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20220401"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

var _ extensions.Importer = &RoleDefinitionExtension{}

func (extension *RoleDefinitionExtension) Import(
	ctx context.Context,
	rsrc genruntime.ImportableResource,
	owner *genruntime.ResourceReference,
	next extensions.ImporterFunc,
) (extensions.ImportResult, error) {
	result, err := next(ctx, rsrc, owner)
	if err != nil {
		return extensions.ImportResult{}, err
	}

	if definition, ok := rsrc.(*api.RoleDefinition); ok {
		// If this role definition is built in, we don't need to export it
		if definition.Spec.Type != nil {
			if strings.EqualFold(*definition.Spec.Type, "BuiltInRole") {
				return extensions.ImportSkipped("role definition is built-in"), nil
			}
		}
	}

	return result, nil
}
