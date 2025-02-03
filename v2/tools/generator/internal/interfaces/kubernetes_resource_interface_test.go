/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package interfaces

import (
	"regexp"
	"testing"

	"github.com/go-logr/logr"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func Test_createAzureNameFunctionHandlersForValidatedType_givenValidatedType_returnsExpectedResult(t *testing.T) {
	t.Parallel()

	validatedInt := astmodel.NewValidatedType(
		astmodel.IntType,
		astmodel.NumberValidations{})

	defaultString := astmodel.NewValidatedType(
		astmodel.StringType,
		astmodel.StringValidations{
			Patterns: []*regexp.Regexp{
				regexp.MustCompile("^.*/default$"),
			},
		})

	identifierString := astmodel.NewValidatedType(
		astmodel.StringType,
		astmodel.StringValidations{
			Patterns: []*regexp.Regexp{
				regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$"),
			},
		})

	cases := map[string]struct {
		validatedType           *astmodel.ValidatedType
		expectedErrorSubstrings []string
	}{
		"validated int is not accepted": {
			validatedType: validatedInt,
			expectedErrorSubstrings: []string{
				"unable to handle non-string validated definitions",
			},
		},
		"default string works": {
			validatedType: defaultString,
		},
		"identifier string works": {
			validatedType: identifierString,
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			_, err := createAzureNameFunctionHandlersForValidatedType(*c.validatedType, logr.Discard())
			if len(c.expectedErrorSubstrings) == 0 {
				g.Expect(err).To(BeNil())
			} else {
				g.Expect(err).To(HaveOccurred())
				for _, substr := range c.expectedErrorSubstrings {
					g.Expect(err).To(MatchError(ContainSubstring(substr)))
				}
			}
		})
	}
}

func Test_createAzureNameFunctionHandlersForInternalTypeName_givenInternalTypeName_returnsExpectedResult(t *testing.T) {
	t.Parallel()

	intEnum := astmodel.NewEnumType(
		astmodel.IntType,
		astmodel.MakeEnumValue("kilo", "1000"),
		astmodel.MakeEnumValue("mega", "1000000"),
		astmodel.MakeEnumValue("giga", "1000000000"))

	intEnumDef := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "IntEnum"),
		intEnum)

	singleValuedStringEnum := astmodel.NewEnumType(
		astmodel.StringType,
		astmodel.MakeEnumValue("Default", "default"))

	singleValuedStringEnumDef := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "SingleValuedStringEnum"),
		singleValuedStringEnum)

	multipleValueStrunctEnum := astmodel.NewEnumType(
		astmodel.StringType,
		astmodel.MakeEnumValue("Alpha", "alpha"),
		astmodel.MakeEnumValue("Beta", "beta"),
		astmodel.MakeEnumValue("Gamma", "gamma"),
		astmodel.MakeEnumValue("Delta", "delta"))

	multipleValueStrunctEnumDef := astmodel.MakeTypeDefinition(
		astmodel.MakeInternalTypeName(test.Pkg2020, "MultipleValueStrunctEnum"),
		multipleValueStrunctEnum)

	defs := astmodel.MakeTypeDefinitionSetFromDefinitions(
		intEnumDef,
		singleValuedStringEnumDef,
		multipleValueStrunctEnumDef)

	cases := map[string]struct {
		typename                astmodel.InternalTypeName
		expectedErrorSubstrings []string
	}{
		"undefined typename": {
			typename: astmodel.MakeInternalTypeName(test.Pkg2020, "Unknown"),
			expectedErrorSubstrings: []string{
				"unable to resolve type",
				"Unknown",
			},
		},
		"int enum does not have string base type": {
			typename: intEnumDef.Name(),
			expectedErrorSubstrings: []string{
				"unable to handle non-string enum base type",
			},
		},
		"single-valued string enum works": {
			typename: singleValuedStringEnumDef.Name(),
		},
		"multiple-valued string enum works": {
			typename: multipleValueStrunctEnumDef.Name(),
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			_, err := createAzureNameFunctionHandlersForInternalTypeName(c.typename, defs)
			if len(c.expectedErrorSubstrings) == 0 {
				g.Expect(err).To(BeNil())
			} else {
				g.Expect(err).To(HaveOccurred())
				for _, substr := range c.expectedErrorSubstrings {
					g.Expect(err).To(MatchError(ContainSubstring(substr)))
				}
			}
		})
	}
}

func Test_createAzureNameFunctionHandlersForPrimitiveType_givenPrimitiveType_returnsExpectedResult(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		testType                *astmodel.PrimitiveType
		expectedErrorSubstrings []string
	}{
		"string is accepted": {
			testType: astmodel.StringType,
		},
		"int is not accepted": {
			testType: astmodel.IntType,
			expectedErrorSubstrings: []string{
				"cannot use type",
				"int",
				astmodel.AzureNameProperty,
			},
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			_, err := createAzureNameFunctionHandlersForPrimitiveType(c.testType)
			if len(c.expectedErrorSubstrings) == 0 {
				g.Expect(err).To(BeNil())
			} else {
				g.Expect(err).To(HaveOccurred())
				for _, substr := range c.expectedErrorSubstrings {
					g.Expect(err).To(MatchError(ContainSubstring(substr)))
				}
			}
		})
	}
}
