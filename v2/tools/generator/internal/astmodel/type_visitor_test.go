/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	"github.com/pkg/errors"

	. "github.com/onsi/gomega"
)

func Test_Visit_GivenCountingTypeVisitor_ReturnsExpectedCounts(t *testing.T) {
	t.Parallel()

	arrType := NewArrayType(StringType)
	mapType := NewMapType(StringType, StringType)
	optType := OptionalIntType

	intProp := NewPropertyDefinition("number", "number", IntType)
	arrProp := NewPropertyDefinition("arrType", "arrType", arrType)
	mapProp := NewPropertyDefinition("mapProp", "mapProp", mapType)
	optProp := NewPropertyDefinition("optType", "optType", optType)

	objType := EmptyObjectType.WithProperties(intProp, arrProp, mapProp, optProp)

	resType := NewResourceType(objType, StringType)

	cases := []struct {
		name     string
		theType  Type
		expected int
	}{
		// Simple types get visited once
		{"Strings are visited once", StringType, 1},
		{"Integers are visited once", IntType, 1},
		{"Booleans are visited once", BoolType, 1},
		// Composite types require more visits
		{"Arrays require two visits", arrType, 2},
		{"Map types require three visits", mapType, 3},
		{"Optional types require two visits", optType, 2},
		// Complex types require even more visits
		{"Object types require one visit plus visiting each property", objType, 9},
		{"Resource types require one visit plus visiting spec and status", resType, 11},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			v := MakeCountingTypeVisitor()
			_, err := v.Visit(c.theType, nil)

			g.Expect(v.VisitCount).To(Equal(c.expected))
			g.Expect(err).To(BeNil())
		})
	}
}

type CountingTypeVisitor struct {
	TypeVisitor
	VisitCount int
}

func MakeCountingTypeVisitor() *CountingTypeVisitor {
	result := &CountingTypeVisitor{}

	countEachName := func(this *TypeVisitor, it TypeName, ctx interface{}) (Type, error) {
		result.VisitCount++
		return IdentityVisitOfTypeName(this, it, ctx)
	}

	countEachArray := func(this *TypeVisitor, it *ArrayType, ctx interface{}) (Type, error) {
		result.VisitCount++
		return IdentityVisitOfArrayType(this, it, ctx)
	}

	countEachPrimitive := func(this *TypeVisitor, it *PrimitiveType, ctx interface{}) (Type, error) {
		result.VisitCount++
		return IdentityVisitOfPrimitiveType(this, it, ctx)
	}

	countEachObject := func(this *TypeVisitor, it *ObjectType, ctx interface{}) (Type, error) {
		result.VisitCount++
		return IdentityVisitOfObjectType(this, it, ctx)
	}

	countEachMap := func(this *TypeVisitor, it *MapType, ctx interface{}) (Type, error) {
		result.VisitCount++
		return IdentityVisitOfMapType(this, it, ctx)
	}

	countEachEnum := func(this *TypeVisitor, it *EnumType, ctx interface{}) (Type, error) {
		result.VisitCount++
		return IdentityVisitOfEnumType(this, it, ctx)
	}

	countEachOptional := func(this *TypeVisitor, it *OptionalType, ctx interface{}) (Type, error) {
		result.VisitCount++
		return IdentityVisitOfOptionalType(this, it, ctx)
	}

	countEachResource := func(this *TypeVisitor, it *ResourceType, ctx interface{}) (Type, error) {
		result.VisitCount++
		return IdentityVisitOfResourceType(this, it, ctx)
	}

	result.TypeVisitor = TypeVisitorBuilder{
		VisitTypeName:     countEachName,
		VisitArrayType:    countEachArray,
		VisitPrimitive:    countEachPrimitive,
		VisitObjectType:   countEachObject,
		VisitMapType:      countEachMap,
		VisitEnumType:     countEachEnum,
		VisitOptionalType: countEachOptional,
		VisitResourceType: countEachResource,
	}.Build()

	return result
}

func TestIdentityVisitorReturnsEqualResult(t *testing.T) {
	t.Parallel()

	mapOfStringToString := NewMapType(StringType, StringType)
	mapOfStringToBool := NewMapType(StringType, BoolType)

	fullName := NewPropertyDefinition("FullName", "full-name", StringType)
	familyName := NewPropertyDefinition("FamilyName", "family-name", StringType)
	knownAs := NewPropertyDefinition("KnownAs", "known-as", StringType)

	transform := NewFakeFunction("Transform")
	transmogrify := NewFakeFunction("Transmogrify")
	skew := NewFakeFunction("Skew")

	person := NewObjectType().WithProperties(fullName, familyName, knownAs)
	individual := NewObjectType().WithProperties(fullName).
		WithFunction(transform).
		WithFunction(transmogrify).
		WithFunction(skew)

	stringEnum := NewEnumType(StringType)

	oneOfType := NewOneOfType("x",BoolType,StringType)

	allofType := NewAllOfType()

	resource := NewResourceType(person, individual)

	iface := NewInterfaceType(transform)

	cases := []struct {
		name    string
		subject Type
	}{
		{"Strings", StringType},
		{"Integers", IntType},
		{"Booleans", BoolType},
		{"Map of String to String", mapOfStringToString},
		{"Map of String to Bool", mapOfStringToBool},
		{"Object type with properties", person},
		{"Object type with functions", individual},
		{"Resource type", resource},
		{"Interface type", iface},
		{"Enum type", stringEnum},
		{"One of type", oneOfType},
		{"All of Type",allofType},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			v := TypeVisitorBuilder{}.Build()
			result, err := v.Visit(c.subject, nil)

			g.Expect(err).To(BeNil())
			g.Expect(TypeEquals(result, c.subject)).To(BeTrue())

			// Test both ways to be paranoid
			// If this test fails while the previous line passes, there's likely a bug in .Equals()
			g.Expect(TypeEquals(c.subject, result)).To(BeTrue())
		})
	}
}

func TestMakeTypeVisitorWithInjectedFunctions(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		subject       Type
		configure     func(builder *TypeVisitorBuilder)
		expectedType  Type
		expectedError string
	}{
		{
			"PrimitiveTypeHandler",
			StringType,
			func(builder *TypeVisitorBuilder) {
				builder.VisitPrimitive = func(tv *TypeVisitor, p *PrimitiveType, _ interface{}) (Type, error) {
					return p, errors.New(p.name)
				}
			},
			StringType,
			"string",
		},
		{
			"PrimitiveTypeSimplified",
			StringType,
			func(builder *TypeVisitorBuilder) {
				builder.VisitPrimitive = func(p *PrimitiveType) (Type, error) {
					return p, errors.New(p.name)
				}
			},
			StringType,
			"string",
		},
		{
			"OptionalTypeHandler",
			OptionalIntType,
			func(builder *TypeVisitorBuilder) {
				builder.VisitOptionalType = func(tv *TypeVisitor, ot *OptionalType, _ interface{}) (Type, error) {
					return ot, errors.New(ot.String())
				}
			},
			OptionalIntType,
			"optional",
		},
		{
			"OptionalTypeSimplified",
			OptionalIntType,
			func(builder *TypeVisitorBuilder) {
				builder.VisitOptionalType = func(ot *OptionalType) (Type, error) {
					return ot, errors.New(ot.String())
				}
			},
			OptionalIntType,
			"optional",
		},
		{
			"EnumTypeHandler",
			NewEnumType(StringType),
			func(builder *TypeVisitorBuilder) {
				builder.VisitEnumType = func(tv *TypeVisitor, et *EnumType, _ interface{}) (Type, error) {
					return et, errors.New(et.String())
				}
			},
			NewEnumType(StringType),
			"enum",
		},
		{
			"EnumTypeSimplified",
			NewEnumType(StringType),
			func(builder *TypeVisitorBuilder) {
				builder.VisitEnumType = func(et *EnumType) (Type, error) {
					return et, errors.New(et.String())
				}
			},
			NewEnumType(StringType),
			"enum",
		},
		{
			"ObjectTypeHandler",
			NewObjectType(),
			func(builder *TypeVisitorBuilder) {
				builder.VisitObjectType = func(tv *TypeVisitor, ot *ObjectType, _ interface{}) (Type, error) {
					return ot, errors.New(ot.String())
				}
			},
			NewObjectType(),
			"object",
		},
		{
			"ObjectTypeSimplified",
			NewObjectType(),
			func(builder *TypeVisitorBuilder) {
				builder.VisitObjectType = func(ot *ObjectType) (Type, error) {
					return ot, errors.New(ot.String())
				}
			},
			NewObjectType(),
			"object",
		},
		{
			"OneOfTypeHandler",
			NewOneOfType("x",BoolType,StringType),
			func(builder *TypeVisitorBuilder) {
				builder.VisitOneOfType = func(tv *TypeVisitor, ot *OneOfType, _ interface{}) (Type, error) {
					return ot, errors.New(ot.String())
				}
			},
			NewOneOfType("x",BoolType,StringType),
			"oneOf",
		},
		{
			"OneOfTypeSimplified",
			NewOneOfType("x",BoolType,StringType),
			func(builder *TypeVisitorBuilder) {
				builder.VisitOneOfType = func(ot *OneOfType) (Type, error) {
					return ot, errors.New(ot.String())
				}
			},
			NewOneOfType("x",BoolType,StringType),
			"oneOf",
		},
		{
			"AllOfTypeHandler",
			NewAllOfType(),
			func(builder *TypeVisitorBuilder) {
				builder.VisitAllOfType = func(tv *TypeVisitor, at *AllOfType, _ interface{}) (Type, error) {
					return at, errors.New(at.String())
				}
			},
			NewAllOfType(),
			"allOf",
		},
		{
			"AllOfTypeSimplified",
			NewAllOfType(),
			func(builder *TypeVisitorBuilder) {
				builder.VisitAllOfType = func(at *AllOfType) (Type, error) {
					return at, errors.New(at.String())
				}
			},
			NewAllOfType(),
			"allOf",
		},
		{
			"MapTypeHandler",
			NewMapType(StringType, IntType),
			func(builder *TypeVisitorBuilder) {
				builder.VisitMapType = func(tv *TypeVisitor, mt *MapType, _ interface{}) (Type, error) {
					return mt, errors.New(mt.String())
				}
			},
			NewMapType(StringType, IntType),
			"map",
		},
		{
			"MapTypeSimplified",
			NewMapType(StringType, IntType),
			func(builder *TypeVisitorBuilder) {
				builder.VisitMapType = func(mt *MapType) (Type, error) {
					return mt, errors.New(mt.String())
				}
			},
			NewMapType(StringType, IntType),
			"map",
		},
		{
			"ArrayTypeHandler",
			NewArrayType(StringType),
			func(builder *TypeVisitorBuilder) {
				builder.VisitArrayType = func(tv *TypeVisitor, at *ArrayType, _ interface{}) (Type, error) {
					return at, errors.New(at.String())
				}
			},
			NewArrayType(StringType),
			"[]string",
		},
		{
			"ArrayTypeSimplified",
			NewArrayType(StringType),
			func(builder *TypeVisitorBuilder) {
				builder.VisitArrayType = func(at *ArrayType) (Type, error) {
					return at, errors.New(at.String())
				}
			},
			NewArrayType(StringType),
			"[]string",
		},
		{
			"FlaggedTypeHandler",
			StorageFlag.ApplyTo(StringType),
			func(builder *TypeVisitorBuilder) {
				builder.VisitFlaggedType = func(tv *TypeVisitor, ft *FlaggedType, _ interface{}) (Type, error) {
					return ft, errors.New(ft.String())
				}
			},
			StorageFlag.ApplyTo(StringType),
			"storage",
		},
		{
			"FlaggedTypeSimplified",
			StorageFlag.ApplyTo(StringType),
			func(builder *TypeVisitorBuilder) {
				builder.VisitFlaggedType = func(ft *FlaggedType) (Type, error) {
					return ft, errors.New(ft.String())
				}
			},
			StorageFlag.ApplyTo(StringType),
			"storage",
		},
	}

	// TODO: Pending tests for
	// visitTypeName:
	// visitResourceType:
	// visitValidatedType:
	// visitErroredType:

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			var builder TypeVisitorBuilder
			c.configure(&builder)
			v := builder.Build()
			typ, err := v.Visit(c.subject, nil)
			g.Expect(TypeEquals(typ, c.expectedType)).To(BeTrue())
			g.Expect(err.Error()).To(ContainSubstring(c.expectedError))
		})
	}
}
