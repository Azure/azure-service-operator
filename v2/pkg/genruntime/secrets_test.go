/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_ValidateSecretDestination_EmptyListValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(genruntime.ValidateSecretDestinations(nil)).To(Succeed())
}

func Test_ValidateSecretDestination_ListWithNilElementsValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.SecretDestination{
		nil,
		nil,
	}
	g.Expect(genruntime.ValidateSecretDestinations(destinations)).To(Succeed())
}

func Test_ValidateSecretDestination_LengthOneListValidates(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destinations := []*genruntime.SecretDestination{
		{Name: "n1", Key: "key1"},
	}
	g.Expect(genruntime.ValidateSecretDestinations(destinations)).To(Succeed())
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
	g.Expect(genruntime.ValidateSecretDestinations(destinations)).To(Succeed())
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
	g.Expect(genruntime.ValidateSecretDestinations(destinations)).To(Succeed())
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
	_, err := genruntime.ValidateSecretDestinations(destinations)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("cannot write more than one secret to destination Name: \"n1\", Key: \"key1\""))
}
