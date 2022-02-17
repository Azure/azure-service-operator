/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package secrets_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

func TestCollector_DestinationsWithSameSecret_Merges(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "foo",
	}
	destination2 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "bar",
	}

	collector := secrets.NewSecretCollector("ns")
	collector.AddSecretValue(destination1, "secret1")
	collector.AddSecretValue(destination2, "secret2")

	result := collector.Secrets()
	g.Expect(result).To(HaveLen(1))
	g.Expect(result[0].Name).To(Equal("mysecret"))
	g.Expect(result[0].Namespace).To(Equal("ns"))
	g.Expect(result[0].StringData).To(HaveLen(2))
	g.Expect(result[0].StringData["foo"]).To(Equal("secret1"))
	g.Expect(result[0].StringData["bar"]).To(Equal("secret2"))
}

func TestCollector_DestinationsDifferentSecret_DoesNotMerge(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "foo",
	}
	destination2 := &genruntime.SecretDestination{
		Name: "theirsecret",
		Key:  "bar",
	}

	collector := secrets.NewSecretCollector("ns")
	collector.AddSecretValue(destination1, "secret1")
	collector.AddSecretValue(destination2, "secret2")

	result := collector.Secrets()
	g.Expect(result).To(HaveLen(2))
	g.Expect(result[0].Name).To(Equal("mysecret"))
	g.Expect(result[0].Namespace).To(Equal("ns"))
	g.Expect(result[0].StringData).To(HaveLen(1))
	g.Expect(result[0].StringData["foo"]).To(Equal("secret1"))

	g.Expect(result[1].Name).To(Equal("theirsecret"))
	g.Expect(result[1].Namespace).To(Equal("ns"))
	g.Expect(result[1].StringData).To(HaveLen(1))
	g.Expect(result[1].StringData["bar"]).To(Equal("secret2"))
}

func TestCollector_SameDestinationSameKey_Overwrites(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "foo",
	}

	collector := secrets.NewSecretCollector("ns")
	// Note: We expect that the collector overwrites if there is a collision, but
	// this isn't a failure mode we want to encourage in actual user specified secret
	// destinations. We may decide to reject this at the webhook level.
	collector.AddSecretValue(destination1, "secret1")
	collector.AddSecretValue(destination1, "secret2")

	result := collector.Secrets()
	g.Expect(result).To(HaveLen(1))
	g.Expect(result[0].Name).To(Equal("mysecret"))
	g.Expect(result[0].Namespace).To(Equal("ns"))
	g.Expect(result[0].StringData).To(HaveLen(1))
	g.Expect(result[0].StringData["foo"]).To(Equal("secret2"))
}
