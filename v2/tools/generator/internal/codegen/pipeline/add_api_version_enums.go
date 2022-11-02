/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const AddAPIVersionEnumsStageId = "add-api-version-enums"

func AddAPIVersionEnums() *Stage {
	stage := NewStage(
		AddAPIVersionEnumsStageId,
		"Add enums for API Versions in each package",
		func(ctx context.Context, state *State) (*State, error) {
			newDefs := make(astmodel.TypeDefinitionSet)

			apiVersions := apiVersions{
				generated: make(map[astmodel.PackageReference]apiVersion),
				output:    newDefs,
			}

			for name, def := range state.Definitions() {
				if rt, ok := astmodel.AsResourceType(def.Type()); ok {
					version := apiVersions.Get(name.PackageReference)
					def = def.WithType(rt.WithAPIVersion(version.name, version.value))
				}

				newDefs.Add(def)
			}

			return state.WithDefinitions(newDefs), nil
		},
	)

	return stage
}

type apiVersions struct {
	generated map[astmodel.PackageReference]apiVersion
	output    astmodel.TypeDefinitionSet
}

type apiVersion struct {
	name  astmodel.TypeName
	value astmodel.EnumValue
}

func (vs apiVersions) Get(pr astmodel.PackageReference) apiVersion {
	if v, ok := vs.generated[pr]; ok {
		return v
	}

	name := astmodel.MakeTypeName(pr, "APIVersion") // TODO: constant?
	value := astmodel.MakeEnumValue(
		"Value",
		fmt.Sprintf("%q", apiVersionFromPackageReference(pr)))

	result := apiVersion{name: name, value: value}
	vs.generated[pr] = result
	vs.output.Add(
		astmodel.MakeTypeDefinition(
			name,
			astmodel.NewEnumType(astmodel.StringType, value)))
	return result
}

func apiVersionFromPackageReference(pr astmodel.PackageReference) string {
	localPR, ok := pr.(astmodel.LocalPackageReference)
	if !ok {
		panic("all resources should have local package references")
	}

	return localPR.ApiVersion()
}
