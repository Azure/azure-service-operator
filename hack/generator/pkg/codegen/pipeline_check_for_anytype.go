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

// checkForAnyType returns a stage that will return an error if there
// are any uses of AnyType remaining in the passed defs.
func checkForAnyType() PipelineStage {
	return PipelineStage{
		"Check for rogue AnyTypes",
		func(ctx context.Context, defs astmodel.Types) (astmodel.Types, error) {
			var badNames []astmodel.TypeName
			for name, def := range defs {
				if containsAnyType(def.Type()) {
					badNames = append(badNames, name)
				}
			}

			if len(badNames) == 0 {
				return defs, nil
			}

			packages, err := groupNamesIntoPackages(badNames)
			if err != nil {
				return nil, errors.Wrap(err, "summarising bad types")
			}
			return nil, errors.Errorf("AnyTypes found - add exclusions for: %s", strings.Join(packages, ", "))
		},
	}
}

func containsAnyType(theType astmodel.Type) bool {
	var found bool
	visitor := astmodel.MakeTypeVisitor()
	visitor.VisitPrimitive = func(_ *astmodel.TypeVisitor, it *astmodel.PrimitiveType, _ interface{}) astmodel.Type {
		if it == astmodel.AnyType {
			found = true
		}
		return it
	}
	visitor.Visit(theType, nil)
	return found
}

func groupNamesIntoPackages(names []astmodel.TypeName) ([]string, error) {
	grouped := make(map[string][]string)
	for _, name := range names {
		group, version, err := name.PackageReference.GroupAndPackage()
		if err != nil {
			return nil, err
		}
		groupVersion := group + "/" + version
		grouped[groupVersion] = append(grouped[groupVersion], name.Name())
	}

	var groupNames []string
	for groupName := range grouped {
		groupNames = append(groupNames, groupName)
	}
	sort.Strings(groupNames)

	if klog.V(2).Enabled() {
		for _, groupName := range groupNames {
			sort.Strings(grouped[groupName])
			klog.Infof("%s: %v", groupName, grouped[groupName])
		}
	}

	return groupNames, nil
}
