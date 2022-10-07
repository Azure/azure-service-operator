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
)

func TestCollector_DestinationsWithSameConfigMap_Merges(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.ConfigMapDestination{
		Name: "myconfig",
		Key:  "foo",
	}
	destination2 := &genruntime.ConfigMapDestination{
		Name: "myconfig",
		Key:  "bar",
	}
	destination3 := &genruntime.ConfigMapDestination{
		Name: "myconfig",
		Key:  "baz",
	}

	collector := configmaps.NewCollector("ns")
	collector.AddValue(destination1, "value1")
	collector.AddValue(destination2, "value2")
	collector.AddBinaryValue(destination3, []byte("value3"))

	result, err := collector.Values()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(HaveLen(1))
	g.Expect(result[0].Name).To(Equal("myconfig"))
	g.Expect(result[0].Namespace).To(Equal("ns"))
	g.Expect(result[0].Data).To(HaveLen(2))
	g.Expect(result[0].Data["foo"]).To(Equal("value1"))
	g.Expect(result[0].Data["bar"]).To(Equal("value2"))
	g.Expect(result[0].BinaryData["baz"]).To(Equal([]byte("value3")))
}

func TestCollector_DestinationsDifferentConfigMap_DoesNotMerge(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.ConfigMapDestination{
		Name: "myconfig",
		Key:  "foo",
	}
	destination2 := &genruntime.ConfigMapDestination{
		Name: "theirconfig",
		Key:  "bar",
	}

	collector := configmaps.NewCollector("ns")
	collector.AddValue(destination1, "value1")
	collector.AddValue(destination2, "value2")

	result, err := collector.Values()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(HaveLen(2))
	g.Expect(result[0].Name).To(Equal("myconfig"))
	g.Expect(result[0].Namespace).To(Equal("ns"))
	g.Expect(result[0].Data).To(HaveLen(1))
	g.Expect(result[0].Data["foo"]).To(Equal("value1"))

	g.Expect(result[1].Name).To(Equal("theirconfig"))
	g.Expect(result[1].Namespace).To(Equal("ns"))
	g.Expect(result[1].Data).To(HaveLen(1))
	g.Expect(result[1].Data["bar"]).To(Equal("value2"))
}

func TestCollector_SameDestinationSameKey_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.ConfigMapDestination{
		Name: "myconfig",
		Key:  "foo",
	}

	collector := configmaps.NewCollector("ns")
	collector.AddValue(destination1, "value1")
	collector.AddValue(destination1, "value2")

	_, err := collector.Values()
	g.Expect(err).To(HaveOccurred())
}

func TestCollector_BinaryDestinationSameKey_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	destination1 := &genruntime.ConfigMapDestination{
		Name: "myconfig",
		Key:  "foo",
	}

	collector := configmaps.NewCollector("ns")
	collector.AddValue(destination1, "value1")
	collector.AddBinaryValue(destination1, []byte("value2"))

	_, err := collector.Values()
	g.Expect(err).To(HaveOccurred())
}
