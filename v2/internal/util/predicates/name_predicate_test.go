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

func Test_NamePredicate(t *testing.T) {
	t.Parallel()

	predicate := predicates.MakeNamePredicate("name1", "name3")

	matchingName1 := &SampleObj{
		ObjectMeta: metav1.ObjectMeta{
			Name: "name1",
		},
	}
	nonmatchingName1 := &SampleObj{
		ObjectMeta: metav1.ObjectMeta{
			Name: "name2",
		},
	}
	matchingName2 := &SampleObj{
		ObjectMeta: metav1.ObjectMeta{
			Name: "name3",
		},
	}

	tests := []struct {
		name     string
		obj      client.Object
		expected bool
	}{
		{
			name:     "Matched name tracked",
			obj:      matchingName1,
			expected: true,
		},
		{
			name:     "Another matched name tracked",
			obj:      matchingName2,
			expected: true,
		},
		{
			name:     "Unmatched name not tracked",
			obj:      nonmatchingName1,
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
