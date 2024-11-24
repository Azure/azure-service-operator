/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package resolver

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/rotisserie/eris"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

func TestReferenceNotFoundError_Is(t *testing.T) {
	t.Parallel()

	fooName := types.NamespacedName{
		Namespace: "default",
		Name:      "foo",
	}
	fooName2 := types.NamespacedName{
		Namespace: "default",
		Name:      "foo",
	}
	barName := types.NamespacedName{
		Namespace: "default",
		Name:      "bar",
	}

	tests := map[string]struct {
		expected error
		actual   error
		is       bool
	}{
		"is equal to copy": {
			expected: core.NewReferenceNotFoundError(fooName, nil),
			actual:   core.NewReferenceNotFoundError(fooName, nil),
			is:       true,
		},
		"is equal to correct type with different NamespacedName instance": {
			expected: core.NewReferenceNotFoundError(fooName, nil),
			actual:   core.NewReferenceNotFoundError(fooName2, nil),
			is:       true,
		},
		"is not equal to correct type with different NamespacedName": {
			expected: core.NewReferenceNotFoundError(fooName, nil),
			actual:   core.NewReferenceNotFoundError(barName, nil),
			is:       false,
		},
		"is not equal to incorrect type": {
			expected: core.NewReferenceNotFoundError(fooName, nil),
			actual:   eris.New("this is a test"),
			is:       false,
		},
	}

	for n, c := range tests {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			g.Expect(eris.Is(c.expected, c.actual)).To(Equal(c.is))
		})
	}
}

func TestReferenceNotFoundError_AsWithCorrectType_Works(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fooName := types.NamespacedName{
		Namespace: "default",
		Name:      "foo",
	}

	err := core.NewReferenceNotFoundError(fooName, nil)

	var e *core.ReferenceNotFound
	g.Expect(eris.As(err, &e)).To(BeTrue())
	g.Expect(e).To(Equal(err))
}

func TestReferenceNotFoundError_AsIncorrectType_Fails(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fooName := types.NamespacedName{
		Namespace: "default",
		Name:      "foo",
	}

	err := core.NewReferenceNotFoundError(fooName, nil)

	var e *apierrors.StatusError
	g.Expect(eris.As(err, &e)).To(BeFalse())
}

func TestReferenceNotFoundError_RemembersCause(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fooName := types.NamespacedName{
		Namespace: "default",
		Name:      "foo",
	}

	cause := eris.New("I caused the problem")
	mid := core.NewReferenceNotFoundError(fooName, cause)
	err := eris.Wrap(mid, "adding context")

	g.Expect(eris.Cause(err)).To(Equal(cause))

	errorText := err.Error()
	g.Expect(errorText).To(ContainSubstring("I caused the problem"))
	g.Expect(errorText).To(ContainSubstring("default/foo does not exist"))
	g.Expect(errorText).To(ContainSubstring("adding context"))

	// Check that the error is formatted correctly
	g.Expect(eris.ToString(err, true)).To(ContainSubstring("errors_test.go:115"))
	g.Expect(eris.ToString(eris.Cause(err), true)).To(ContainSubstring("errors_test.go:113"))
}
