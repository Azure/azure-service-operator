/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"
)

type ErroredType struct {
	inner    Type
	errors   []string
	warnings []string
}

var _ Type = &ErroredType{}

func NewErroredType(t Type, errors []string, warnings []string) *ErroredType {
	result := &ErroredType{
		inner:    nil,
		errors:   errors,
		warnings: warnings,
	}

	return result.WithType(t) // using WithType ensures warnings and errors get merged if needed
}

func (e *ErroredType) InnerType() Type {
	return e.inner
}

func (e *ErroredType) WithType(t Type) *ErroredType {
	if otherError, ok := t.(*ErroredType); ok {
		// nested errors merge errors & warnings
		e.inner = otherError.inner
		e.errors = append(e.errors, otherError.errors...)
		e.warnings = append(e.warnings, otherError.warnings...)
	} else {
		e.inner = t
	}

	return e
}

func (e *ErroredType) Equals(t Type) bool {
	other, ok := t.(*ErroredType)
	if !ok {
		return false
	}

	if e == other {
		return true // short-circuit
	}

	return ((e.inner == nil && other.inner == nil) || e.inner.Equals(other.inner)) &&
		stringSlicesEqual(e.warnings, other.warnings) &&
		stringSlicesEqual(e.errors, other.errors)
}

func stringSlicesEqual(l []string, r []string) bool {
	if len(l) != len(r) {
		return false
	}

	for ix := range l {
		if l[ix] != r[ix] {
			return false
		}
	}

	return true
}

func (e *ErroredType) References() TypeNameSet {
	if e.inner == nil {
		return nil
	}

	return e.inner.References()
}

func (e *ErroredType) RequiredPackageReferences() *PackageReferenceSet {
	if e.inner == nil {
		return NewPackageReferenceSet()
	}

	return e.inner.RequiredPackageReferences()
}

func (e *ErroredType) handleWarningsAndErrors() {
	for _, warning := range e.warnings {
		klog.Warning(warning)
	}

	if len(e.errors) > 0 {
		var errs []error
		for _, err := range e.errors {
			errs = append(errs, errors.New(err))
		}

		if len(errs) == 1 {
			panic(errs[0])
		} else {
			panic(kerrors.NewAggregate(errs))
		}
	}
}

func (e *ErroredType) AsDeclarations(cgc *CodeGenerationContext, dc DeclarationContext) []dst.Decl {
	e.handleWarningsAndErrors()
	if e.inner == nil {
		return nil
	}

	return e.inner.AsDeclarations(cgc, dc)
}

func (e *ErroredType) AsType(cgc *CodeGenerationContext) dst.Expr {
	e.handleWarningsAndErrors()
	if e.inner == nil {
		return nil
	}

	return e.inner.AsType(cgc)
}

func (e *ErroredType) String() string {
	if e.inner == nil {
		return "(error hole)"
	}

	has := "warnings"
	if len(e.errors) > 0 {
		has = "errors"
	}

	return fmt.Sprintf("%s (has %s)", e.inner.String(), has)
}
