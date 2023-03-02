/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package predicates_test

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"

	"github.com/Azure/azure-service-operator/v2/internal/util/predicates"
)

type SampleObj struct {
	metav1.ObjectMeta
	metav1.TypeMeta
}

func (s *SampleObj) DeepCopyObject() runtime.Object {
	// This isn't a full implementation of deep copy but we don't need one
	objCopy := *s
	return &objCopy
}

var _ client.Object = &SampleObj{}

func Test_SelectAnnotationsChangedPredicate_DetectsChanges(t *testing.T) {
	t.Parallel()

	annotationKey := "foobar"
	predicate := predicates.MakeSelectAnnotationChangedPredicate(
		map[string]predicates.HasAnnotationChanged{
			annotationKey: predicates.HasBasicAnnotationChanged,
		})

	empty := &SampleObj{}
	withWatchedAnnotation1 := &SampleObj{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotationKey: "1234",
			},
		},
	}
	withWatchedAnnotation2 := &SampleObj{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotationKey: "5678",
			},
		},
	}
	withUnwatchedAnnotation1 := &SampleObj{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"whatever": "1234",
			},
		},
	}
	withUnwatchedAnnotation2 := &SampleObj{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"whatever": "5678",
			},
		},
	}
	withWatchedAndUnwatchedAnnotation := &SampleObj{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotationKey: "1234",
				"whatever":    "1234",
			},
		},
	}

	tests := []struct {
		name     string
		old      client.Object
		new      client.Object
		expected bool
	}{
		{
			name:     "Watched annotation added to empty annotation set",
			old:      empty,
			new:      withWatchedAnnotation1,
			expected: true,
		},
		{
			name:     "Watched annotation added to populated annotation set",
			old:      withUnwatchedAnnotation1,
			new:      withWatchedAndUnwatchedAnnotation,
			expected: true,
		},
		{
			name:     "Watched annotation removed",
			old:      withWatchedAnnotation1,
			new:      empty,
			expected: true,
		},
		{
			name:     "Watched annotation removed",
			old:      withWatchedAndUnwatchedAnnotation,
			new:      withUnwatchedAnnotation1,
			expected: true,
		},
		{
			name:     "Watched annotation changed",
			old:      withWatchedAnnotation1,
			new:      withWatchedAnnotation2,
			expected: true,
		},
		{
			name:     "Unwatched annotation added",
			old:      empty,
			new:      withUnwatchedAnnotation1,
			expected: false,
		},
		{
			name:     "Unwatched annotation changed",
			old:      withUnwatchedAnnotation1,
			new:      withUnwatchedAnnotation2,
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
					ObjectOld: tt.old,
					ObjectNew: tt.new,
				})
			g.Expect(result).To(Equal(tt.expected))
		})
	}
}

func Test_SelectAnnotationsChangedPredicate_MissingAnnotationPassesNilToHandler(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	newValue := "1234"
	annotationKey := "foobar"

	empty := &SampleObj{}
	withWatchedAnnotation := &SampleObj{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotationKey: newValue,
			},
		},
	}

	handler1 := func(old *string, new *string) bool {
		g.Expect(old).To(BeNil())
		g.Expect(new).To(Equal(&newValue))
		return false
	}
	handler2 := func(old *string, new *string) bool {
		g.Expect(old).To(Equal(&newValue))
		g.Expect(new).To(BeNil())
		return false
	}
	testPredicateReceivesExpectedValue(handler1, empty, withWatchedAnnotation)
	testPredicateReceivesExpectedValue(handler2, withWatchedAnnotation, empty)
}

func testPredicateReceivesExpectedValue(handler predicates.HasAnnotationChanged, old client.Object, new client.Object) {
	annotationKey := "foobar"

	predicate := predicates.MakeSelectAnnotationChangedPredicate(
		map[string]predicates.HasAnnotationChanged{
			annotationKey: handler,
		})

	predicate.Update(
		event.UpdateEvent{
			ObjectOld: old,
			ObjectNew: new,
		})
}
