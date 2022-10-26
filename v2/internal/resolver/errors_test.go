/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package resolver

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

func TestOwnerNotFound_Is(t *testing.T) {
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

	tests := []struct {
		name     string
		expected error
		actual   error
		is       bool
	}{
		{
			name:     "is equal to copy",
			expected: core.NewReferenceNotFoundError(fooName, nil),
			actual:   core.NewReferenceNotFoundError(fooName, nil),
			is:       true,
		},
		{
			name:     "is equal to correct type with different NamespacedName instance",
			expected: core.NewReferenceNotFoundError(fooName, nil),
			actual:   core.NewReferenceNotFoundError(fooName2, nil),
			is:       true,
		},
		{
			name:     "is not equal to correct type with different NamespacedName",
			expected: core.NewReferenceNotFoundError(fooName, nil),
			actual:   core.NewReferenceNotFoundError(barName, nil),
			is:       false,
		},
		{
			name:     "is not equal to incorrect type",
			expected: core.NewReferenceNotFoundError(fooName, nil),
			actual:   errors.New("this is a test"),
			is:       false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			g.Expect(tt.is).To(Equal(errors.Is(tt.expected, tt.actual)))
		})
	}
}

func TestOwnerNotFound_AsCorrectType_Works(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fooName := types.NamespacedName{
		Namespace: "default",
		Name:      "foo",
	}

	err := core.NewReferenceNotFoundError(fooName, nil)

	var e *core.ReferenceNotFound
	g.Expect(errors.As(err, &e)).To(BeTrue())
	g.Expect(e).To(Equal(err))
}

func TestOwnerNotFound_AsIncorrectType_Fails(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fooName := types.NamespacedName{
		Namespace: "default",
		Name:      "foo",
	}

	err := core.NewReferenceNotFoundError(fooName, nil)

	var e *apierrors.StatusError
	g.Expect(errors.As(err, &e)).To(BeFalse())
}

func TestOwnerNotFound_RemembersCause(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	fooName := types.NamespacedName{
		Namespace: "default",
		Name:      "foo",
	}

	cause := errors.New("I caused the problem")
	err := errors.WithStack(core.NewReferenceNotFoundError(fooName, cause))

	g.Expect(errors.Cause(err)).To(Equal(cause))

	errorText := err.Error()
	g.Expect(errorText).To(ContainSubstring("I caused the problem"))
	g.Expect(errorText).To(ContainSubstring("default/foo does not exist"))

	// Note that both of the below lines are fragile with respect to line number and will
	// need to be changed if the lines causing the error above are changed.
	g.Expect(StackTraceOf(err)).To(ContainSubstring("errors_test.go:119"))
	g.Expect(StackTraceOf(errors.Cause(err))).To(ContainSubstring("errors_test.go:118"))
}

func StackTraceOf(e error) string {
	var tracer stackTracer
	if errors.As(e, &tracer) {
		var stack strings.Builder
		for _, f := range tracer.StackTrace() {
			stack.WriteString(fmt.Sprintf("%+s:%d\n", f, f))
		}

		return stack.String()
	}

	return ""
}

// stackTracer allows access to the stack trace of an error
// This should be exposed by the errors package, but it is not
type stackTracer interface {
	StackTrace() errors.StackTrace
}
