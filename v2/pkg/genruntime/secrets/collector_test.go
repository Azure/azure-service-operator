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
	destination3 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "baz",
	}

	collector := secrets.NewCollector("ns")
	collector.AddValue(destination1, "secret1")
	collector.AddValue(destination2, "secret2")
	collector.AddBinaryValue(destination3, []byte("secret3"))

	result, err := collector.Values()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(HaveLen(1))
	g.Expect(result[0].Name).To(Equal("mysecret"))
	g.Expect(result[0].Namespace).To(Equal("ns"))
	g.Expect(result[0].StringData).To(HaveLen(2))
	g.Expect(result[0].StringData["foo"]).To(Equal("secret1"))
	g.Expect(result[0].StringData["bar"]).To(Equal("secret2"))
	g.Expect(result[0].Data["baz"]).To(Equal([]byte("secret3")))
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

	collector := secrets.NewCollector("ns")
	collector.AddValue(destination1, "secret1")
	collector.AddValue(destination2, "secret2")

	result, err := collector.Values()
	g.Expect(err).ToNot(HaveOccurred())
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

func TestCollector_MissingValue_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "foo",
	}

	collector := secrets.NewCollector("ns")
	collector.AddValue(destination1, "")

	_, err := collector.Values()
	g.Expect(err).To(HaveOccurred())
}

func TestCollector_SameDestinationSameKey_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "foo",
	}

	collector := secrets.NewCollector("ns")
	collector.AddValue(destination1, "secret1")
	collector.AddValue(destination1, "secret2")

	_, err := collector.Values()
	g.Expect(err).To(HaveOccurred())
}

func TestCollector_BinaryDestinationSameKey_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "foo",
	}

	collector := secrets.NewCollector("ns")
	collector.AddValue(destination1, "value1")
	collector.AddBinaryValue(destination1, []byte("value2"))

	_, err := collector.Values()
	g.Expect(err).To(HaveOccurred())
}

func TestCollector_AnnotationsApplied(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "foo",
		Annotations: map[string]string{
			"reflector.v1/reflect": "true",
		},
	}

	collector := secrets.NewCollector("ns")
	collector.AddValue(destination, "secret1")

	result, err := collector.Values()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(HaveLen(1))
	g.Expect(result[0].Annotations).To(Equal(map[string]string{
		"reflector.v1/reflect": "true",
	}))
}

func TestCollector_LabelsApplied(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "foo",
		Labels: map[string]string{
			"app": "myapp",
		},
	}

	collector := secrets.NewCollector("ns")
	collector.AddValue(destination, "secret1")

	result, err := collector.Values()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(HaveLen(1))
	g.Expect(result[0].Labels).To(Equal(map[string]string{
		"app": "myapp",
	}))
}

func TestCollector_MultipleDestinationsSameSecretSameAnnotations_Merges(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "foo",
		Annotations: map[string]string{
			"reflector.v1/reflect": "true",
		},
		Labels: map[string]string{
			"app": "myapp",
		},
	}
	destination2 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "bar",
		Annotations: map[string]string{
			"reflector.v1/reflect": "true",
		},
		Labels: map[string]string{
			"app": "myapp",
		},
	}

	collector := secrets.NewCollector("ns")
	collector.AddValue(destination1, "secret1")
	collector.AddValue(destination2, "secret2")

	result, err := collector.Values()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(HaveLen(1))
	g.Expect(result[0].StringData).To(HaveLen(2))
	g.Expect(result[0].Annotations).To(Equal(map[string]string{
		"reflector.v1/reflect": "true",
	}))
	g.Expect(result[0].Labels).To(Equal(map[string]string{
		"app": "myapp",
	}))
}

func TestCollector_MultipleDestinationsSameSecretDifferentAnnotations_MergesDistinctKeys(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "foo",
		Annotations: map[string]string{
			"annotation1": "value1",
		},
	}
	destination2 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "bar",
		Annotations: map[string]string{
			"annotation2": "value2",
		},
	}

	collector := secrets.NewCollector("ns")
	collector.AddValue(destination1, "secret1")
	collector.AddValue(destination2, "secret2")

	result, err := collector.Values()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(HaveLen(1))
	g.Expect(result[0].Annotations).To(Equal(map[string]string{
		"annotation1": "value1",
		"annotation2": "value2",
	}))
}

func TestCollector_ConflictingAnnotations_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "foo",
		Annotations: map[string]string{
			"mykey": "value1",
		},
	}
	destination2 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "bar",
		Annotations: map[string]string{
			"mykey": "value2",
		},
	}

	collector := secrets.NewCollector("ns")
	collector.AddValue(destination1, "secret1")
	collector.AddValue(destination2, "secret2")

	_, err := collector.Values()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring(`annotation collision for secret "mysecret": annotation "mykey" has conflicting values "value1" and "value2"`))
}

func TestCollector_ConflictingLabels_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "foo",
		Labels: map[string]string{
			"app": "myapp",
		},
	}
	destination2 := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "bar",
		Labels: map[string]string{
			"app": "otherapp",
		},
	}

	collector := secrets.NewCollector("ns")
	collector.AddValue(destination1, "secret1")
	collector.AddValue(destination2, "secret2")

	_, err := collector.Values()
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring(`label collision for secret "mysecret": label "app" has conflicting values "myapp" and "otherapp"`))
}

func TestCollector_NoAnnotationsOrLabels_HasNilMetadata(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination := &genruntime.SecretDestination{
		Name: "mysecret",
		Key:  "foo",
	}

	collector := secrets.NewCollector("ns")
	collector.AddValue(destination, "secret1")

	result, err := collector.Values()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(HaveLen(1))
	g.Expect(result[0].Annotations).To(BeNil())
	g.Expect(result[0].Labels).To(BeNil())
}
