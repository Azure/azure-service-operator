/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const CreateTypesForBackwardCompatibilityID = "createTypesForBackwardCompatibility"

// CreateTypesForBackwardCompatibility returns a pipeline stage that creates copies of types in other packages to
// provide backward compatibility with previous releases of Azure Service Operator.
// Backward compatibility versions are created for all versions to allow users of older versions of the operator to
// easily upgrade.
func CreateTypesForBackwardCompatibility(prefix string) *Stage {
	stage := NewStage(
		CreateTypesForBackwardCompatibilityID,
		"Create clones of types for backward compatibility with prior ASO versions",
		func(ctx context.Context, state *State) (*State, error) {
			// Work out the new names for all our new types
			renames := createBackwardCompatibilityRenameMap(state.Definitions(), prefix)

			// Rename all the types
			visitor := astmodel.NewRenamingVisitor(renames)

			renamed, err := visitor.RenameAll(state.Definitions())
			if err != nil {
				return nil, errors.Wrap(err, "creating types for backward compatibility")
			}

			// Update the description of each to reflect purpose
			renamedWithDescriptions, err := addCompatibilityComments(renamed)
			if err != nil {
				return nil, errors.Wrap(err, "changing comments of types created for backward compatibility")
			}

			// Add the new types into our state
			defs := state.Definitions()
			defs.AddTypes(renamedWithDescriptions)

			return state.WithDefinitions(defs), nil
		})

	stage.RequiresPrerequisiteStages(CreateStorageTypesStageID)
	stage.RequiresPostrequisiteStages(CreateConversionGraphStageId)
	return stage
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
				"Backward compatibility type for %s.%s",
				name.PackageReference.PackageName(),
				name.Name()),
		}

		t, err := visitor.Visit(def.Type(), nil)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		result.Add(def.WithType(t).WithDescription(desc))
	}

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	return result, nil
}

func removePropertyDescriptions(ot *astmodel.ObjectType) astmodel.Type {
	result := ot
	for _, property := range ot.Properties() {
		result = result.WithProperty(property.WithDescription(""))
	}

	return result
}

func createBackwardCompatibilityRenameMap(
	set astmodel.TypeDefinitionSet,
	versionPrefix string) map[astmodel.TypeName]astmodel.TypeName {
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

	switch r:=name.PackageReference.(type) {
	case astmodel.LocalPackageReference:
		ref = r.WithVersionPrefix(versionPrefix)
		break
	case astmodel.StoragePackageReference:
		local := r.Local().WithVersionPrefix(versionPrefix)
		ref = astmodel.MakeStoragePackageReference(local)
		break
	default:
		panic(fmt.Sprintf("unexpected package reference type %T", r))
	}
	return name.WithPackageReference(ref)
}


