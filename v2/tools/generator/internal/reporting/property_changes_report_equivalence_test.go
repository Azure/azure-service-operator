/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func Test_PropertyChangesReport_typesEquivalent_GivenTypes_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	refPkg := astmodel.MakeNamedLocalPackageReference("v", "demo", "ref")
	fooPkg := astmodel.MakeNamedLocalPackageReference("v", "demo", "foo")
	name := astmodel.MakeInternalTypeName(refPkg, "Name")

	cases := map[string]struct {
		left     astmodel.Type
		right    astmodel.Type
		expected bool
	}{
		"Primitive types equal": {
			left:     astmodel.StringType,
			right:    astmodel.StringType,
			expected: true,
		},
		"Primitive types not equal": {
			left:     astmodel.StringType,
			right:    astmodel.IntType,
			expected: false,
		},
		"Primitive type not equal to optional type": {
			left:     astmodel.StringType,
			right:    astmodel.OptionalStringType,
			expected: false,
		},
		"Optional type not equal to primitive type": {
			left:     astmodel.OptionalStringType,
			right:    astmodel.StringType,
			expected: false,
		},
		"InternalTypeName equal to self": {
			left:     name,
			right:    name,
			expected: true,
		},
		"InternalTypeName not equal to different InternalTypeName": {
			left:     name,
			right:    astmodel.MakeInternalTypeName(refPkg, "DifferentName"),
			expected: false,
		},
		"InternalTypeName equal even if in different package": {
			left:     name,
			right:    astmodel.MakeInternalTypeName(fooPkg, "Name"),
			expected: true,
		},
		"InternalTypeName not equal to optional InternalTypeName": {
			left:     name,
			right:    astmodel.NewOptionalType(name),
			expected: false,
		},
		"Optional InternalTypeName not equal to wrapped InternalTypeName": {
			left:     astmodel.NewOptionalType(name),
			right:    name,
			expected: false,
		},
		"Optional InternalTypeName equal to self": {
			left:     astmodel.NewOptionalType(name),
			right:    astmodel.NewOptionalType(name),
			expected: true,
		},
	}

	rpt := NewPropertyChangesReport(
		// No renames
		astmodel.TypeAssociation{},
		// No type definitions
		astmodel.TypeDefinitionSet{},
		// No type renames,
		func(astmodel.InternalTypeName) (string, bool) {
			return "", false
		},
		// No property renames
		func(astmodel.InternalTypeName, astmodel.PropertyName) (string, bool) {
			return "", false
		},
	)

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := rpt.typesEquivalent(c.left, c.right)
			if result != c.expected {
				t.Errorf("expected %v, got %v", c.expected, result)
			}
		})
	}
}
