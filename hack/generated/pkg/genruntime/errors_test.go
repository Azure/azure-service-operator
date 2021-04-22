/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

func TestOwnerNotFound_Is(t *testing.T) {
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

	tests := []struct {
		name     string
		expected error
		actual   error
		is       bool
	}{
		{
			name:     "is equal to copy",
			expected: NewReferenceNotFoundError(fooName, nil),
			actual:   NewReferenceNotFoundError(fooName, nil),
			is:       true,
		},
		{
			name:     "is equal to correct type with different NamespacedName instance",
			expected: NewReferenceNotFoundError(fooName, nil),
			actual:   NewReferenceNotFoundError(fooName2, nil),
			is:       true,
		},
		{
			name:     "is not equal to correct type with different NamespacedName",
			expected: NewReferenceNotFoundError(fooName, nil),
			actual:   NewReferenceNotFoundError(barName, nil),
			is:       false,
		},
		{
			name:     "is not equal to incorrect type",
			expected: NewReferenceNotFoundError(fooName, nil),
			actual:   errors.New("this is a test"),
			is:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)

			g.Expect(tt.is).To(Equal(errors.Is(tt.expected, tt.actual)))
		})
	}
}

func TestOwnerNotFound_AsCorrectType_Works(t *testing.T) {
	g := NewWithT(t)

	fooName := types.NamespacedName{
		Namespace: "default",
		Name:      "foo",
	}

	err := NewReferenceNotFoundError(fooName, nil)

	var e *ReferenceNotFound
	g.Expect(errors.As(err, &e)).To(BeTrue())
	g.Expect(e).To(Equal(err))
}

func TestOwnerNotFound_AsIncorrectType_Fails(t *testing.T) {
	g := NewWithT(t)

	fooName := types.NamespacedName{
		Namespace: "default",
		Name:      "foo",
	}

	err := NewReferenceNotFoundError(fooName, nil)

	var e *apierrors.StatusError
	g.Expect(errors.As(err, &e)).To(BeFalse())
}

func TestOwnerNotFound_RemembersCause(t *testing.T) {
	g := NewWithT(t)

	fooName := types.NamespacedName{
		Namespace: "default",
		Name:      "foo",
	}

	cause := errors.New("I caused the problem")
	err := errors.WithStack(NewReferenceNotFoundError(fooName, cause))

	g.Expect(errors.Cause(err)).To(Equal(cause))

	fmtedErr := fmt.Sprintf("%+v", err)
	g.Expect(fmtedErr).To(ContainSubstring("I caused the problem"))
	g.Expect(fmtedErr).To(ContainSubstring("default/foo does not exist"))
	// Note that both of the below lines are fragile with respect to line number and will
	// need to be changed if the lines causing the error above are changed.
	g.Expect(fmtedErr).To(ContainSubstring("errors_test.go:110"))
	g.Expect(fmtedErr).To(ContainSubstring("errors_test.go:111"))
}
