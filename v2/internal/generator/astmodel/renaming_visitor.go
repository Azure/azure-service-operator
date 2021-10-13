/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import kerrors "k8s.io/apimachinery/pkg/util/errors"

// RenamingVisitor is a visitor for performing simple TypeName renames
type RenamingVisitor struct {
	f func(TypeName) TypeName
	// renames map[TypeName]TypeName
	visitor TypeVisitor
}

// NewRenamingVisitor creates a new visitor which performs the renames specified
func NewRenamingVisitor(renames map[TypeName]TypeName) *RenamingVisitor {
	rename := func(name TypeName) TypeName {
		typeName := name
		rename, ok := renames[name]
		if ok {
			typeName = rename
		}

		return typeName
	}

	return NewRenamingVisitorFromLambda(rename)
}

// NewRenamingVisitorFromLambda creates a new visitor which performs renames using the specified lambda
func NewRenamingVisitorFromLambda(f func(name TypeName) TypeName) *RenamingVisitor {
	r := &RenamingVisitor{
		f: f,
	}

	rename := func(this *TypeVisitor, it TypeName, ctx interface{}) (Type, error) {
		newName := f(it)
		return IdentityVisitOfTypeName(this, newName, ctx)
	}

	r.visitor = TypeVisitorBuilder{
		VisitTypeName: rename,
	}.Build()

	return r
}

// Rename applies the renames to the Type
func (r *RenamingVisitor) Rename(t Type) (Type, error) {
	return r.visitor.Visit(t, nil)
}

// RenameAll applies the renames to the Types
func (r *RenamingVisitor) RenameAll(types Types) (Types, error) {
	result := make(Types)
	var errs []error

	for _, def := range types {
		renamed, err := r.visitor.VisitDefinition(def, nil)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		result.Add(renamed)
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, err
	}

	return result, nil
}
