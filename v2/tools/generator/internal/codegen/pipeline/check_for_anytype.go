/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"sort"
	"strings"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// CheckForAnyTypeStageID is the unique identifier for this stage
const CheckForAnyTypeStageID = "rogueCheck"

// FilterOutDefinitionsUsingAnyType returns a stage that will check for any definitions
// containing AnyTypes. It accepts a set of packages that we expect to contain types
// with AnyTypes. Those packages will be quietly filtered out of the output of the
// stage, but if there are more AnyTypes in other packages they'll be reported as an
// error. The stage will also return an error if there are packages that we expect
// to have AnyTypes but turn out not to, ensuring that we clean up our configuration
// as the schemas are fixed and our handling improves.
func FilterOutDefinitionsUsingAnyType(packages []string) *Stage {
	return checkForAnyType("Filter out rogue definitions using AnyTypes", packages)
}

// EnsureDefinitionsDoNotUseAnyTypes returns a stage that will check for any
// definitions containing AnyTypes. The stage will return errors for each type
// found that uses an AnyType.
func EnsureDefinitionsDoNotUseAnyTypes() *Stage {
	return checkForAnyType("Check for rogue definitions using AnyTypes", []string{})
}

func checkForAnyType(description string, packages []string) *Stage {
	expectedPackages := set.Make[string]()
	for _, p := range packages {
		expectedPackages.Add(p)
	}

	return NewStage(
		CheckForAnyTypeStageID,
		description,
		func(ctx context.Context, state *State) (*State, error) {
			defs := state.Definitions()
			var badNames []astmodel.InternalTypeName
			output := make(astmodel.TypeDefinitionSet)
			for name, def := range defs {
				if containsAnyType(def.Type()) {
					badNames = append(badNames, name)
				}

				// We only want to include this type in the output if
				// it's not in a package that we know contains
				// AnyTypes.
				if expectedPackages.Contains(name.InternalPackageReference().FolderPath()) {
					continue
				}

				output.Add(def)
			}

			badPackages, err := collectBadPackages(badNames, expectedPackages)
			if err != nil {
				return nil, eris.Wrap(err, "summarising bad types")
			}

			if len(badPackages) > 0 {
				return nil, eris.Errorf("AnyTypes found - add exclusions for: %s", strings.Join(badPackages, ", "))
			}

			return state.WithDefinitions(output), nil
		})
}

func containsAnyType(theType astmodel.Type) bool {
	var found bool
	detectAnyType := func(it *astmodel.PrimitiveType) astmodel.Type {
		if it == astmodel.AnyType {
			found = true
		}

		return it
	}

	visitor := astmodel.TypeVisitorBuilder[any]{
		VisitPrimitive: detectAnyType,
	}.Build()

	_, _ = visitor.Visit(theType, nil)
	return found
}

func collectBadPackages(
	names []astmodel.InternalTypeName,
	expectedPackages set.Set[string],
) ([]string, error) {
	grouped := make(map[string][]string)
	for _, name := range names {
		packagePath := name.InternalPackageReference().FolderPath()
		grouped[packagePath] = append(grouped[packagePath], name.Name())
	}

	var groupNames []string //nolint:prealloc // unlikely case
	for groupName := range grouped {
		// Only complain about this package if it's one we don't know about.
		if expectedPackages.Contains(groupName) {
			expectedPackages.Remove(groupName)
			continue
		}

		groupNames = append(groupNames, groupName)
	}
	sort.Strings(groupNames)

	// Complain if there were some packages where we expected problems
	// but didn't see any.
	if len(expectedPackages) > 0 {
		var leftovers []string
		for value := range expectedPackages {
			leftovers = append(leftovers, value)
		}
		sort.Strings(leftovers)
		return nil, eris.Errorf(
			"no AnyTypes found in: %s", strings.Join(leftovers, ", "))

	}

	return groupNames, nil
}
