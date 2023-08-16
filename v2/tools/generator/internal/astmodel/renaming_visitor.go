/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

// RenamingVisitor is a visitor for performing simple TypeName renames
type RenamingVisitor struct {
	f func(TypeName) TypeName
	// renames map[TypeName]TypeName
	visitor TypeVisitor[any]
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

	r.visitor = TypeVisitorBuilder[any]{
		VisitTypeName:     r.updateTypeName,
		VisitResourceType: r.updateResourceOwner,
	}.Build()

	return r
}

func (r *RenamingVisitor) updateTypeName(this *TypeVisitor[any], it TypeName, ctx any) (Type, error) {
	newName := r.f(it)
	return IdentityVisitOfTypeName(this, newName, ctx)
}

func (r *RenamingVisitor) updateResourceOwner(this *TypeVisitor[any], it *ResourceType, ctx any) (Type, error) {
	if it.Owner() != nil {
		// TODO: Should this actually happen in TypeVisitor itself?
		it = it.WithOwner(r.f(it.Owner()))
	}
	return IdentityVisitOfResourceType(this, it, ctx)
}

// Rename applies the renames to the Type whose definition is passed in.
func (r *RenamingVisitor) Rename(t Type) (Type, error) {
	return r.visitor.Visit(t, nil)
}

// RenameDefinition applies the renames to the TypeDefinition’s Type and TypeName.
func (r *RenamingVisitor) RenameDefinition(td TypeDefinition) (TypeDefinition, error) {
	return r.visitor.VisitDefinition(td, nil)
}

// RenameAll applies the renames to all the type definitions in the provided set
func (r *RenamingVisitor) RenameAll(definitions TypeDefinitionSet) (TypeDefinitionSet, error) {
	result := make(TypeDefinitionSet)
	var errs []error

	for _, def := range definitions {
		renamed, err := r.visitor.VisitDefinition(def, nil)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		if result.Contains(renamed.Name()) {
			errs = append(errs, errors.Errorf("a definition for %s already exists", renamed.Name()))
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
