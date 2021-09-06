/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func TestReadableConversionEndpointSet_CreatePropertyEndpoints_GivenObject_CreatesExpectedEndpoints(t *testing.T) {
	g := NewWithT(t)

	person := astmodel.NewObjectType().
		WithProperties(test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	set := NewWritableConversionEndpointSet()
	set.CreatePropertyEndpoints(person)

	g.Expect(set).To(HaveKey("FullName"))
	g.Expect(set).To(HaveKey("KnownAs"))
	g.Expect(set).To(HaveKey("FamilyName"))
	g.Expect(set).To(HaveLen(3))
}
