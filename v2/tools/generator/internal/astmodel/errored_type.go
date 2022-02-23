/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"strings"

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

var _ MetaType = &ErroredType{}

func NewErroredType(t Type, errors []string, warnings []string) *ErroredType {
	result := &ErroredType{
		inner:    nil,
		errors:   errors,
		warnings: warnings,
	}

	return result.WithType(t) // using WithType ensures warnings and errors get merged if needed
}

// Errors returns the errors stored in this ErroredType
func (e *ErroredType) Errors() []string {
	return append([]string(nil), e.errors...)
}

// Warnings returns the warnings stored in this ErroredType
func (e *ErroredType) Warnings() []string {
	return append([]string(nil), e.warnings...)
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

func (e *ErroredType) Equals(t Type, overrides EqualityOverrides) bool {
	if e == t {
		return true // short-circuit
	}

	other, ok := t.(*ErroredType)
	if !ok {
		return false
	}

	return ((e.inner == nil && other.inner == nil) || e.inner.Equals(other.inner, overrides)) &&
		stringSlicesEqual(e.warnings, other.warnings) &&
		stringSlicesEqual(e.errors, other.errors)
}

func propertyNameSlicesEqual(l, r []PropertyName) bool {
	if len(l) != len(r) {
		return false
	}

	if len(l) > 0 && (&l[0] == &r[0]) {
		return true // reference equality
	}

	for ix := range l {
		if string(l[ix]) != string(r[ix]) {
			return false
		}
	}

	return true
}

func stringSlicesEqual(l, r []string) bool {
	if len(l) != len(r) {
		return false
	}

	if len(l) > 0 && (&l[0] == &r[0]) {
		return true // reference equality
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

// AsZero renders an expression for the "zero" value of the type
// by delegating to the inner type
func (e *ErroredType) AsZero(definitions TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr {
	return e.inner.AsZero(definitions, ctx)
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

// Unwrap returns the type contained within the error type
func (e *ErroredType) Unwrap() Type {
	return e.inner
}

// WriteDebugDescription adds a description of the current errored type to the passed builder,
// builder receives the full description, including the nested type, errors and warnings
// definitions is for resolving named types
func (e *ErroredType) WriteDebugDescription(builder *strings.Builder, definitions TypeDefinitionSet) {
	builder.WriteString("Error[")
	if e.inner != nil {
		e.inner.WriteDebugDescription(builder, definitions)
	} else {
		builder.WriteString("<missing>")
	}

	for _, e := range e.errors {
		builder.WriteString("|")
		builder.WriteString(e)
	}

	for _, w := range e.warnings {
		builder.WriteString("|")
		builder.WriteString(w)
	}

	builder.WriteString("]")
}
