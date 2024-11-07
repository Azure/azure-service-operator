/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package secrets_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

func Test_ValidateSecretDestination_EmptyListValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	warnings, err := secrets.ValidateDestinations(nil, nil, nil)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateSecretDestination_ListWithNilElementsValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.SecretDestination{
		nil,
		nil,
	}

	warnings, err := secrets.ValidateDestinations(nil, destinations, nil)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateSecretDestinationExpressions_ListWithNilElementsValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*core.DestinationExpression{
		nil,
		nil,
	}

	warnings, err := secrets.ValidateDestinations(nil, nil, destinations)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateSecretDestination_LengthOneListValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.SecretDestination{
		{Name: "n1", Key: "key1"},
	}

	warnings, err := secrets.ValidateDestinations(nil, destinations, nil)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateSecretDestinationExpressions_LengthOneListValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*core.DestinationExpression{
		{Name: "n1", Key: "key1", Value: "resource.status.id"},
	}

	warnings, err := secrets.ValidateDestinations(nil, nil, destinations)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateSecretDestination_ListWithoutCollisionsValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.SecretDestination{
		{Name: "n1", Key: "key1"},
		{Name: "n1", Key: "key2"},
		{Name: "n1", Key: "key3"},
		{Name: "n1", Key: "key4"},
	}

	warnings, err := secrets.ValidateDestinations(nil, destinations, nil)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateSecretDestinationExpressions_ListWithoutCollisionsValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*core.DestinationExpression{
		{Name: "n1", Key: "key1", Value: "resource.status.id"},
		{Name: "n1", Key: "key2", Value: "resource.status.id"},
		{Name: "n1", Key: "key3", Value: "resource.status.id"},
		{Name: "n1", Key: "key4", Value: "resource.status.id"},
	}

	warnings, err := secrets.ValidateDestinations(nil, nil, destinations)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateSecretDestination_ListWithDifferentCasesValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.SecretDestination{
		{Name: "n1", Key: "key1"},
		{Name: "n1", Key: "Key1"},
		{Name: "n1", Key: "key3"},
		{Name: "n1", Key: "key4"},
	}

	warnings, err := secrets.ValidateDestinations(nil, destinations, nil)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}

func Test_ValidateSecretDestination_ListWithCollisionsFailsValidation(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.SecretDestination{
		{Name: "n1", Key: "key1"},
		{Name: "n2", Key: "key1"},
		{Name: "n3", Key: "key1"},
		{Name: "n1", Key: "key1"},
	}

	_, err := secrets.ValidateDestinations(nil, destinations, nil)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("cannot write more than one secret to destination Name: \"n1\", Key: \"key1\""))
}

func Test_ValidateSecretDestinationAndExpressions_CollisionBetweenEachFailsValidation(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.SecretDestination{
		{Name: "n3", Key: "key1"},
		{Name: "n4", Key: "key1"},
		{Name: "n5", Key: "key1"},
	}

	destinationExpressions := []*core.DestinationExpression{
		{Name: "n1", Key: "key1", Value: "resource.status.id"},
		{Name: "n2", Key: "key1", Value: "resource.status.id"},
		{Name: "n3", Key: "key1", Value: "resource.status.id"},
	}

	_, err := secrets.ValidateDestinations(nil, destinations, destinationExpressions)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("cannot write more than one secret to destination Name: \"n3\", Key: \"key1\", Value: \"resource.status.id\""))
}

func Test_ValidateSecretDestinationExpressions_EmptyKeyIgnored(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*core.DestinationExpression{
		{Name: "n1", Value: "resource.status.id"},
		{Name: "n1", Key: "key1", Value: "resource.status.id"},
	}

	warnings, err := secrets.ValidateDestinations(nil, nil, destinations)
	g.Expect(warnings).To(BeNil())
	g.Expect(err).To(BeNil())
}
