/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	ast "github.com/dave/dst"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"
)

type ErroredType struct {
	inner    Type
	errors   []string
	warnings []string
}

func (errored ErroredType) WithType(t Type) ErroredType {
	errored.inner = t
	return errored
}

var _ Type = ErroredType{}

func (errored ErroredType) Equals(t Type) bool {
	other, ok := t.(ErroredType)
	if !ok {
		return false
	}

	return errored.inner.Equals(other.inner) &&
		stringSlicesEqual(errored.warnings, other.warnings) &&
		stringSlicesEqual(errored.errors, other.errors)
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

func (errored ErroredType) References() TypeNameSet {
	return errored.inner.References()
}

func (errored ErroredType) RequiredPackageReferences() *PackageReferenceSet {
	return errored.inner.RequiredPackageReferences()
}

func (errored ErroredType) handleWarningsAndErrors() {
	for _, warning := range errored.warnings {
		klog.Warning(warning)
	}

	if len(errored.errors) > 0 {
		var es []error
		for _, e := range errored.errors {
			es = append(es, errors.New(e))
		}

		if len(es) == 1 {
			panic(errored.errors[0])
		} else {
			panic(kerrors.NewAggregate(es))
		}
	}
}

func (errored ErroredType) AsDeclarations(cgc *CodeGenerationContext, dc DeclarationContext) []ast.Decl {
	errored.handleWarningsAndErrors()
	return errored.inner.AsDeclarations(cgc, dc)
}

func (errored ErroredType) AsType(cgc *CodeGenerationContext) ast.Expr {
	errored.handleWarningsAndErrors()
	return errored.inner.AsType(cgc)
}

func (errored ErroredType) String() string {
	has := "warnings"
	if len(errored.errors) > 0 {
		has = "errors"
	}

	return fmt.Sprintf("%s (has %s)", errored.inner.String(), has)
}
