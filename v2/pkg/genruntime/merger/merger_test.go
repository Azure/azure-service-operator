/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package merger_test

import (
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/merger"
)

var secretS1 = &v1.Secret{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "s1",
		Namespace: "testnamespace",
	},
	StringData: map[string]string{
		"key1": "value1",
	},
}

var secretS1SameKey = &v1.Secret{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "s1",
		Namespace: "testnamespace",
	},
	StringData: map[string]string{
		"key1": "value1",
	},
}

var secretS1NewKey = &v1.Secret{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "s1",
		Namespace: "testnamespace",
	},
	StringData: map[string]string{
		"key2": "value2",
	},
}

var secretDifferentNamespace = &v1.Secret{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "s3",
		Namespace: "othernamespace",
	},
	StringData: map[string]string{
		"key2": "value2",
	},
}

var configMapC1 = &v1.ConfigMap{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "c1",
		Namespace: "testnamespace",
	},
	Data: map[string]string{
		"key1": "value1",
	},
}

var configMapC1NewKey = &v1.ConfigMap{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "c1",
		Namespace: "testnamespace",
	},
	Data: map[string]string{
		"key2": "value2",
	},
}

var configMapC2 = &v1.ConfigMap{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "c2",
		Namespace: "testnamespace",
	},
	Data: map[string]string{
		"key1": "value1",
	},
}

var otherResource = &v1.Namespace{
	ObjectMeta: metav1.ObjectMeta{
		Name: "c2",
	},
}

func TestMerge_MergesDuplicateConfigMapsAndSecrets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	merged, err := merger.MergeObjects([]client.Object{
		secretS1,
		secretS1NewKey,
		configMapC1,
		configMapC1NewKey,
		configMapC2,
	})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(merged).To(HaveLen(3))
	g.Expect(merged[0]).To(BeAssignableToTypeOf(&v1.Secret{}))
	s1 := merged[0].(*v1.Secret)
	g.Expect(s1.StringData).To(HaveLen(2))
	g.Expect(merged[1]).To(BeAssignableToTypeOf(&v1.ConfigMap{}))
	c1 := merged[1].(*v1.ConfigMap)
	g.Expect(c1.Data).To(HaveLen(2))
	g.Expect(merged[2]).To(BeAssignableToTypeOf(&v1.ConfigMap{}))
	c2 := merged[2].(*v1.ConfigMap)
	g.Expect(c2.Data).To(Equal(configMapC2.Data))
}

func TestMerge_PreservesOtherResources(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	merged, err := merger.MergeObjects([]client.Object{
		secretS1,
		secretS1NewKey,
		configMapC1,
		configMapC1NewKey,
		configMapC2,
		otherResource,
	})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(merged).To(HaveLen(4))
}

func TestMerge_SameDestinationKey_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	_, err := merger.MergeObjects([]client.Object{
		secretS1,
		secretS1SameKey,
		configMapC1,
		configMapC1NewKey,
		configMapC2,
	})
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("failed merging secrets: key collision, entry exists for key key1 in StringData"))
}

func TestMerge_DifferentNamespaces_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	_, err := merger.MergeObjects([]client.Object{
		secretS1,
		secretDifferentNamespace,
		configMapC1,
	})
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(Equal("cannot merge objects from different namespaces: testnamespace : othernamespace"))
}
