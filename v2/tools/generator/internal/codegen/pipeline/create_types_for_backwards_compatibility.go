/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
)

const (
	CreateTypesForBackwardCompatibilityStageID = "createTypesForBackwardCompatibility"
	v1betaVersionPrefix                        = "v1beta"
	v1VersionPrefix                            = "v1api"
)

// CreateTypesForBackwardCompatibility returns a pipeline stage that creates copies of types into other packages to
// provide backward compatibility with previous releases of Azure Service Operator.
// Backward compatibility versions are created for all versions to allow users of older versions of the operator to
// easily upgrade.
func CreateTypesForBackwardCompatibility(
	prefix string,
	configuration *config.ObjectModelConfiguration,
) *Stage {
	stage := NewStage(
		CreateTypesForBackwardCompatibilityStageID,
		"Create clones of types for backward compatibility with prior ASO versions",
		func(ctx context.Context, state *State) (*State, error) {
			resources, err := findResourcesRequiringCompatibilityVersion(prefix, state.Definitions(), configuration)
			if err != nil {
				return nil, err
			}

			compatibilityDefs, err := createBackwardCompatibleDefinitions(resources, state.Definitions())
			if err != nil {
				return nil, err
			}

			// Add the new types into our state
			finalDefs := state.Definitions()
			finalDefs.AddTypes(compatibilityDefs)

			return state.WithDefinitions(finalDefs), nil
		})

	//
	// Our pre-requisite support can't currently handle our needs because the earlier stage MUST be present, else the
	// pipeline fails to run. These requirements are more of the "if the stage is present it must run before me" variety
	//
	// ApplyExportFiltersStageID - we want to filter for export first to avoid creating a lot of backward compatibility
	//     types that would immediately be pruned. We also need to avoid issues where the filters might remove the main
	//     API type without removing the compatibility type
	//
	// AddSecretsStageID - any transformations made to accommodate secrets must happen before we clone the type
	//
	//stage.RequiresPrerequisiteStages(
	//	ApplyExportFiltersStageID, // export filters won't correctly prune compatibility types, so we create them afterwards
	//	AddSecretsStageID,         // secrets need to be configured before we copy the type
	//)

	stage.RequiresPostrequisiteStages(
		CreateARMTypesStageID,     // ARM types need to be created for each compatibility type
		CreateStorageTypesStageID, // storage types need to be created too
	)

	return stage
}

func createBackwardCompatibleDefinitions(
	resources astmodel.TypeDefinitionSet,
	definitions astmodel.TypeDefinitionSet,
) (astmodel.TypeDefinitionSet, error) {
	// Find all type definitions needed for the resources specified
	defs, err := astmodel.FindConnectedDefinitions(definitions, resources)
	if err != nil {
		return nil, errors.Wrap(err, "finding types connected to resources requiring backward compatibility")
	}

	// Update the description of each to reflect purpose
	withDescriptions, err := addCompatibilityComments(defs)
	if err != nil {
		return nil, errors.Wrap(err, "changing comments of types created for backward compatibility")
	}

	// Rename all the types into our compatibility namespace
	renames := createBackwardCompatibilityRenameMap(defs, v1betaVersionPrefix)
	visitor := astmodel.NewRenamingVisitor(renames)
	renamed, err := visitor.RenameAll(withDescriptions)
	if err != nil {
		return nil, errors.Wrap(err, "creating types for backward compatibility")
	}

	return renamed, nil
}

func findResourcesRequiringCompatibilityVersion(
	prefix string,
	definitions astmodel.TypeDefinitionSet,
	configuration *config.ObjectModelConfiguration,
) (astmodel.TypeDefinitionSet, error) {
	compat := make(astmodel.TypeDefinitionSet)
	var errs []error
	resources := astmodel.FindResourceDefinitions(definitions)
	for name, def := range resources {

		// Find out when we started supporting this resource
		from, err := configuration.SupportedFrom.Lookup(name)
		if err != nil {
			if config.IsNotConfiguredError(err) {
				// $supportedFrom is not configured, skip this resource
				continue
			}

			// If something else went wrong, keep details
			errs = append(errs, err)
		}

		if strings.HasPrefix(from, prefix) {
			compat.Add(def)
		}
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"finding resources introduced in versions with prefix %q",
			prefix)
	}

	return compat, nil
}

func addCompatibilityComments(defs astmodel.TypeDefinitionSet) (astmodel.TypeDefinitionSet, error) {
	visitor := astmodel.TypeVisitorBuilder{
		VisitObjectType: removePropertyDescriptions,
	}.Build()

	result := make(astmodel.TypeDefinitionSet)

	var errs []error
	for _, def := range defs {
		name := def.Name()
		desc := []string{
			fmt.Sprintf(
				"Deprecated version of %s. Use %s.%s instead",
				name.Name(),
				name.PackageReference().PackageName(),
				name.Name()),
		}

		t, err := visitor.Visit(def.Type(), nil)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		result.Add(def.WithType(t).WithDescription(desc...))
	}

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	return result, nil
}

func removePropertyDescriptions(ot *astmodel.ObjectType) astmodel.Type {
	result := ot
	ot.Properties().ForEach(func(property *astmodel.PropertyDefinition) {
		result = result.WithProperty(property.WithDescription(""))
	})

	return result
}

func createBackwardCompatibilityRenameMap(
	set astmodel.TypeDefinitionSet,
	versionPrefix string,
) map[astmodel.TypeName]astmodel.TypeName {
	result := make(map[astmodel.TypeName]astmodel.TypeName)

	for name := range set {
		if _, ok := result[name]; !ok {
			newName := createBackwardCompatibilityRename(name, versionPrefix)
			result[name] = newName
		}
	}

	return result
}

func createBackwardCompatibilityRename(name astmodel.TypeName, versionPrefix string) astmodel.TypeName {
	var ref astmodel.PackageReference

	switch r := name.PackageReference().(type) {
	case astmodel.LocalPackageReference:
		ref = r.WithVersionPrefix(versionPrefix)
	case astmodel.StoragePackageReference:
		local := r.Local().WithVersionPrefix(versionPrefix)
		ref = astmodel.MakeStoragePackageReference(local)
	default:
		panic(fmt.Sprintf("unexpected package reference type %T", r))
	}
	return name.WithPackageReference(ref)
}
