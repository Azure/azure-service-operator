/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package predicates_test

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"

	"github.com/Azure/azure-service-operator/v2/internal/util/predicates"
)

func Test_NamespacePredicate(t *testing.T) {
	t.Parallel()

	predicate := predicates.MakeNamespacePredicate("ns1", "ns3")

	matchingNamespace1 := &SampleObj{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ns1",
		},
	}
	nonmatchingNamespace1 := &SampleObj{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ns2",
		},
	}
	matchingNamespace2 := &SampleObj{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ns3",
		},
	}

	tests := []struct {
		name     string
		obj      client.Object
		expected bool
	}{
		{
			name:     "Matched namespace tracked",
			obj:      matchingNamespace1,
			expected: true,
		},
		{
			name:     "Another natched namespace tracked",
			obj:      matchingNamespace2,
			expected: true,
		},
		{
			name:     "Unmatched namespace not tracked",
			obj:      nonmatchingNamespace1,
			expected: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := predicate.Update(
				event.UpdateEvent{
					ObjectOld: tt.obj,
					ObjectNew: tt.obj,
				})
			g.Expect(result).To(Equal(tt.expected))

			result = predicate.Create(
				event.CreateEvent{
					Object: tt.obj,
				})
			g.Expect(result).To(Equal(tt.expected))

			result = predicate.Delete(
				event.DeleteEvent{
					Object: tt.obj,
				})
			g.Expect(result).To(Equal(tt.expected))

			result = predicate.Generic(
				event.GenericEvent{
					Object: tt.obj,
				})
			g.Expect(result).To(Equal(tt.expected))
		})
	}
}
