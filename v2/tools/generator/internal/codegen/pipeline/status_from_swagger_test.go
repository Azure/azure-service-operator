/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"runtime"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func Test_ShouldSkipDir_GivenPath_HasExpectedResult(t *testing.T) {
	t.Parallel()

	linuxCases := []struct {
		name       string
		path       string
		shouldSkip bool
	}{
		// Simple paths
		{"Root", "/", false},
		{"Top level", "/foo/", false},
		{"Nested", "/foo/bar/", false},
		// Paths to skip
		{"Skip top level", "/examples/", true},
		{"Skip nested", "/foo/examples/", true},
		{"Skip nested, trailing directory", "/foo/examples/bar/", true},
	}

	windowsCases := []struct {
		name       string
		path       string
		shouldSkip bool
	}{
		// Simple paths
		{"Drive", "D:\\", false},
		{"Top level, Windows", "D:\\foo\\", false},
		{"Nested, Windows", "D:\\foo\\bar\\", false},
		// Paths to skip
		{"Skip top level, Windows", "D:\\examples\\", true},
		{"Skip nested, Windows", "D:\\foo\\examples\\", true},
		{"Skip nested, trailing directory, Windows", "D:\\foo\\examples\\bar\\", true},
	}

	cases := linuxCases

	// If testing on Windows, also test Windows paths
	// Can't test Windows paths on Linux because *reasons*
	if runtime.GOOS == "windows" {
		cases = append(cases, windowsCases...)
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			skipped := shouldSkipDir(c.path)

			g.Expect(skipped).To(Equal(c.shouldSkip))
		})
	}
}

func Test_StructurallyIdentical_RecursesIntoTypeNames(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	pkg := test.MakeLocalPackageReference("abc", "123")
	name := func(n string) astmodel.TypeName { return astmodel.MakeTypeName(pkg, n) }

	types1 := astmodel.MakeTypeDefinitionSet(map[astmodel.TypeName]astmodel.Type{
		name("LeafInt"): astmodel.IntType,
		name("X"): astmodel.NewObjectType().WithProperties(
			astmodel.NewPropertyDefinition("Prop", "prop", name("LeafInt")),
		),
	})

	types2 := astmodel.MakeTypeDefinitionSet(map[astmodel.TypeName]astmodel.Type{
		name("AnotherLeafInt"): astmodel.IntType,
		name("X"): astmodel.NewObjectType().WithProperties(
			astmodel.NewPropertyDefinition("Prop", "prop", name("AnotherLeafInt")),
		),
	})

	x1 := types1[name("X")].Type()
	x2 := types2[name("X")].Type()

	// safety check - they should be not TypeEquals
	g.Expect(astmodel.TypeEquals(x1, x2)).ToNot(BeTrue())

	// actual check - they should be structurallyIdentical
	identical := structurallyIdentical(x1, types1, x2, types2)
	g.Expect(identical).To(BeTrue())
}

func Test_StructurallyIdentical_DistinguishesIdenticalTypeNames(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	pkg := test.MakeLocalPackageReference("abc", "123")
	name := func(n string) astmodel.TypeName { return astmodel.MakeTypeName(pkg, n) }

	types1 := astmodel.MakeTypeDefinitionSet(map[astmodel.TypeName]astmodel.Type{
		name("Leaf"): astmodel.IntType,
		name("X"): astmodel.NewObjectType().WithProperties(
			astmodel.NewPropertyDefinition("Prop", "prop", name("Leaf")),
		),
	})

	types2 := astmodel.MakeTypeDefinitionSet(map[astmodel.TypeName]astmodel.Type{
		name("Leaf"): astmodel.StringType, // string, not int
		name("X"): astmodel.NewObjectType().WithProperties(
			astmodel.NewPropertyDefinition("Prop", "prop", name("Leaf")),
		),
	})

	x1 := types1[name("X")].Type()
	x2 := types2[name("X")].Type()

	// safety check - they should be TypeEquals
	g.Expect(astmodel.TypeEquals(x1, x2)).To(BeTrue())

	// actual check - they should not be structurallyIdentical
	identical := structurallyIdentical(x1, types1, x2, types2)
	g.Expect(identical).To(BeFalse())
}
