/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package codegen

import (
	"context"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

// filterOutDefinitionsUsingAnyType returns a stage that will check for any definitions
// containing AnyTypes. It accepts a set of packages that we expect to contain types
// with AnyTypes. Those packages will be quietly filtered out of the output of the
// stage, but if there are more AnyTypes in other packages they'll be reported as an
// error. The stage will also return an error if there are packages that we expect
// to have AnyTypes but turn out not to, ensuring that we clean up our configuration
// as the schemas are fixed and our handling improves.
func filterOutDefinitionsUsingAnyType(packages []string) PipelineStage {
	return checkForAnyType("Filter out rogue definitions using AnyTypes", packages)
}

// ensureDefinitionsDoNotUseAnyTypes returns a stage that will check for any
// definitions containing AnyTypes. The stage will return errors for each type
// found that uses an AnyType.
func ensureDefinitionsDoNotUseAnyTypes() PipelineStage {
	return checkForAnyType("Catch rogue definitions using AnyTypes", []string{})
}

func checkForAnyType(description string, packages []string) PipelineStage {
	expectedPackages := make(map[string]struct{}, len(packages))
	for _, p := range packages {
		expectedPackages[p] = struct{}{}
	}

	return MakePipelineStage(
		"rogueCheck",
		description,
		func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {
			var badNames []astmodel.TypeName
			output := make(astmodel.Types)
			for name, def := range defs {
				if containsAnyType(def.Type()) {
					badNames = append(badNames, name)
				}

				packageName := packageName(name)

				// We only want to include this type in the output if
				// it's not in a package that we know contains
				// AnyTypes.
				if _, found := expectedPackages[packageName]; found {
					continue
				}
				output.Add(def)
			}

			badPackages, err := collectBadPackages(badNames, expectedPackages)
			if err != nil {
				return nil, errors.Wrap(err, "summarising bad types")
			}

			if len(badPackages) > 0 {
				return nil, errors.Errorf("AnyTypes found - add exclusions for: %s", strings.Join(badPackages, ", "))
			}

			return output, nil
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

	visitor := astmodel.TypeVisitorBuilder{
		VisitPrimitive: detectAnyType,
	}.Build()

	_, _ = visitor.Visit(theType, nil)
	return found
}

func packageName(name astmodel.TypeName) string {
	if localRef, ok := name.PackageReference.AsLocalPackage(); ok {
		group := localRef.Group()
		version := localRef.Version()

		return group + "/" + version
	}

	return name.PackageReference.PackageName()
}

func collectBadPackages(
	names []astmodel.TypeName,
	expectedPackages map[string]struct{},
) ([]string, error) {
	grouped := make(map[string][]string)
	for _, name := range names {
		groupVersion := packageName(name)
		grouped[groupVersion] = append(grouped[groupVersion], name.Name())
	}

	var groupNames []string
	for groupName := range grouped {
		// Only complain about this package if it's one we don't know about.
		if _, found := expectedPackages[groupName]; found {
			delete(expectedPackages, groupName)
			continue
		}
		groupNames = append(groupNames, groupName)
	}
	sort.Strings(groupNames)

	if klog.V(2).Enabled() {
		for _, groupName := range groupNames {
			sort.Strings(grouped[groupName])
			klog.Infof("%s: %v", groupName, grouped[groupName])
		}
	}

	// Complain if there were some packages where we expected problems
	// but didn't see any.
	if len(expectedPackages) > 0 {
		var leftovers []string
		for value := range expectedPackages {
			leftovers = append(leftovers, value)
		}
		sort.Strings(leftovers)
		return nil, errors.Errorf(
			"no AnyTypes found in: %s", strings.Join(leftovers, ", "))
	}

	return groupNames, nil
}
