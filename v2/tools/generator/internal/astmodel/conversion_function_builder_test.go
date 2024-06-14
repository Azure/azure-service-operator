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
	}

	pkgs := make(map[InternalPackageReference]*PackageDefinition)

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
			}

			stmts, err := builder.BuildConversion(params)
			g.Expect(err).To(Succeed())

			assertStmtsExpected(t, stmts)
		})
	}
}
