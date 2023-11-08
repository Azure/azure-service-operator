/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

func TestTransformSelector_AppliesToType_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()

	mapSelector := &TransformSelector{
		Map: &MapSelector{
			Key: TransformSelector{
				Name: NewFieldMatcher("string"),
			},
			Value: TransformSelector{
				Name: NewFieldMatcher("any"),
			},
		},
	}

	mapType := astmodel.NewMapType(astmodel.StringType, astmodel.AnyType)

	exactNameSelector := &TransformSelector{
		Group:    NewFieldMatcher("definitions"),
		Version:  NewFieldMatcher("v1"),
		Name:     NewFieldMatcher("ResourceCopy"),
		Optional: true,
	}

	wildcardNameSelector := &TransformSelector{
		Group:    NewFieldMatcher("definitions"),
		Version:  NewFieldMatcher("*"),
		Name:     NewFieldMatcher("ResourceCopy"),
		Optional: true,
	}

	nameType := astmodel.NewOptionalType(
		astmodel.MakeInternalTypeName(
			test.MakeLocalPackageReference("definitions", "v1"),
			"ResourceCopy"))

	objectSelector := &TransformSelector{
		Object: true,
	}

	objectType := astmodel.NewObjectType()

	cases := []struct {
		name        string
		target      *TransformSelector
		subject     astmodel.Type
		expectation bool
	}{
		{"Map selector matches map", mapSelector, mapType, true},
		{"Map selector does not match name", mapSelector, astmodel.StringType, false},
		{"Map selector does not match object", mapSelector, objectType, false},
		{"Name selector matches name", exactNameSelector, nameType, true},
		{"Name selector does not match map", exactNameSelector, mapType, false},
		{"Name selector does not match object", exactNameSelector, objectType, false},
		{"Wildcard selector matches name", wildcardNameSelector, nameType, true},
		{"Wildcard selector does not match map", wildcardNameSelector, mapType, false},
		{"Wildcard selector does not match object", wildcardNameSelector, objectType, false},
		{"Object selector matches object", objectSelector, objectType, true},
		{"Object selector does not match map", objectSelector, mapType, false},
		{"Object selector does not match name", objectSelector, nameType, false},
	}

	for _, c := range cases {
		c := c
		t.Run(
			c.name,
			func(t *testing.T) {
				t.Parallel()
				g := NewGomegaWithT(t)

				g.Expect(c.target.AppliesToType(c.subject)).To(Equal(c.expectation))
			})
	}
}
