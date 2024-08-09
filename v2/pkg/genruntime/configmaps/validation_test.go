/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package configmaps_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

func Test_ValidateConfigMapDestination_EmptyListValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	warnings, err := configmaps.ValidateDestinations(nil, nil, nil)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateConfigMapDestination_ListWithNilElementsValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.ConfigMapDestination{
		nil,
		nil,
	}

	warnings, err := configmaps.ValidateDestinations(nil, destinations, nil)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateConfigMapDestinationExpressions_ListWithNilElementsValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*core.DestinationExpression{
		nil,
		nil,
	}

	warnings, err := configmaps.ValidateDestinations(nil, nil, destinations)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateConfigMapDestination_LengthOneListValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.ConfigMapDestination{
		{Name: "n1", Key: "key1"},
	}

	warnings, err := configmaps.ValidateDestinations(nil, destinations, nil)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateConfigMapDestinationExpressions_LengthOneListValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*core.DestinationExpression{
		{Name: "n1", Key: "key1", Value: "resource.status.id"},
	}

	warnings, err := configmaps.ValidateDestinations(nil, nil, destinations)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateConfigMapDestination_ListWithoutCollisionsValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.ConfigMapDestination{
		{Name: "n1", Key: "key1"},
		{Name: "n1", Key: "key2"},
		{Name: "n1", Key: "key3"},
		{Name: "n1", Key: "key4"},
	}

	warnings, err := configmaps.ValidateDestinations(nil, destinations, nil)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateConfigMapDestinationExpressions_ListWithoutCollisionsValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*core.DestinationExpression{
		{Name: "n1", Key: "key1", Value: "resource.status.id"},
		{Name: "n1", Key: "key2", Value: "resource.status.id"},
		{Name: "n1", Key: "key3", Value: "resource.status.id"},
		{Name: "n1", Key: "key4", Value: "resource.status.id"},
	}

	warnings, err := configmaps.ValidateDestinations(nil, nil, destinations)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateConfigMapDestination_ListWithDifferentCasesValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.ConfigMapDestination{
		{Name: "n1", Key: "key1"},
		{Name: "n1", Key: "Key1"},
		{Name: "n1", Key: "key3"},
		{Name: "n1", Key: "key4"},
	}

	warnings, err := configmaps.ValidateDestinations(nil, destinations, nil)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateConfigMapDestination_ListWithCollisionsFailsValidation(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.ConfigMapDestination{
		{Name: "n1", Key: "key1"},
		{Name: "n2", Key: "key1"},
		{Name: "n3", Key: "key1"},
		{Name: "n1", Key: "key1"},
	}
	_, err := configmaps.ValidateDestinations(nil, destinations, nil)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("cannot write more than one configmap value to destination Name: \"n1\", Key: \"key1\""))
}

func Test_ValidateConfigMapDestinationAndExpressions_CollisionBetweenEachFailsValidation(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.ConfigMapDestination{
		{Name: "n3", Key: "key1"},
		{Name: "n4", Key: "key1"},
		{Name: "n5", Key: "key1"},
	}

	destinationExpressions := []*core.DestinationExpression{
		{Name: "n1", Key: "key1", Value: "resource.status.id"},
		{Name: "n2", Key: "key1", Value: "resource.status.id"},
		{Name: "n3", Key: "key1", Value: "resource.status.id"},
	}

	_, err := configmaps.ValidateDestinations(nil, destinations, destinationExpressions)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("cannot write more than one configmap value to destination Name: \"n3\", Key: \"key1\", Value: \"resource.status.id\""))
}

func Test_ValidateConfigMapDestinationExpressions_EmptyKeyIgnored(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*core.DestinationExpression{
		{Name: "n1", Value: "resource.status.id"},
		{Name: "n1", Key: "key1", Value: "resource.status.id"},
	}

	warnings, err := configmaps.ValidateDestinations(nil, nil, destinations)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}
