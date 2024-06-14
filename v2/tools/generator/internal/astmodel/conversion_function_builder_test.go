/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/dave/dst"
)

func TestConversionFunctionBuilderBuildConversion_GivenSourceAndDestinationTypes_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()

	enumName := MakeInternalTypeName(pkg, "GreekEnum")
	enumType := NewEnumType(
		StringType,
		MakeEnumValue("Alpha", "alpha"),
		MakeEnumValue("Beta", "beta"),
		MakeEnumValue("Gamma", "gamma"),
		MakeEnumValue("Delta", "delta"))
	enumDef := MakeTypeDefinition(enumName, enumType)

	aliasName := MakeInternalTypeName(pkg, "alias")
	aliasDef := MakeTypeDefinition(aliasName, StringType)

	pkgDef := NewPackageDefinition(pkg)
	pkgDef.AddDefinition(enumDef)
	pkgDef.AddDefinition(aliasDef)

	pkgs := map[InternalPackageReference]*PackageDefinition{
		pkg: pkgDef,
	}

	cases := map[string]struct {
		sourceType      Type
		destinationType Type
	}{
		"strings": {
			sourceType:      StringType,
			destinationType: StringType,
		},
		"from optional string": {
			sourceType:      NewOptionalType(StringType),
			destinationType: StringType,
		},
		"to optional string": {
			sourceType:      StringType,
			destinationType: NewOptionalType(StringType),
		},
		"to enum from string": {
			sourceType:      StringType,
			destinationType: enumName,
		},
		"from enum to string": {
			sourceType:      enumName,
			destinationType: StringType,
		},
		"between identical enums": {
			sourceType:      enumName,
			destinationType: enumName,
		},
		"from optional enum to string": {
			sourceType:      NewOptionalType(enumName),
			destinationType: StringType,
		},
		"to optional enum from string": {
			sourceType:      StringType,
			destinationType: NewOptionalType(enumName),
		},
		"to alias from string": {
			sourceType:      StringType,
			destinationType: aliasName,
		},
		"from alias to string": {
			sourceType:      aliasName,
			destinationType: StringType,
		},
	}

	for name, c := range cases {
		c := c

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			imports := NewPackageImportSet()
			context := NewCodeGenerationContext(pkg, imports, pkgs)

			idFactory := NewIdentifierFactory()
			builder := NewConversionFunctionBuilder(idFactory, context)

			params := ConversionParameters{
				Source:          dst.NewIdent("source"),
				SourceType:      c.sourceType,
				Destination:     dst.NewIdent("destination"),
				DestinationType: c.destinationType,
				NameHint:        "test",
				Locals:          NewKnownLocalsSet(idFactory),
			}

			stmts, err := builder.BuildConversion(params)
			g.Expect(err).To(Succeed())

			assertStmtsExpected(t, stmts)
		})
	}
}
