/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package test

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// FakeFunction is a fake function that can be used for testing purposes
// Note that there's another FakeFunction in astmodel, but we can't eliminate that one because it would cause a package
// reference cycle
type FakeFunction struct {
	name         string
	Imported     *astmodel.PackageReferenceSet
	Referenced   astmodel.TypeNameSet
	TypeReturned astmodel.Type
	idFactory    astmodel.IdentifierFactory
}

func NewFakeFunction(name string, idFactory astmodel.IdentifierFactory) *FakeFunction {
	return &FakeFunction{
		name:      name,
		idFactory: idFactory,
		Imported:  astmodel.NewPackageReferenceSet(),
	}
}

var _ astmodel.Function = &FakeFunction{}

var _ astmodel.ValueFunction = &FakeFunction{}

func (fake *FakeFunction) Name() string {
	return fake.name
}

func (fake *FakeFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	result := astmodel.NewPackageReferenceSet()
	result.Merge(fake.Imported)
	return result
}

func (fake *FakeFunction) References() astmodel.TypeNameSet {
	return fake.Referenced
}

func (fake *FakeFunction) AsFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
) (*dst.FuncDecl, error) {
	receiverName := fake.idFactory.CreateReceiver(receiver.Name())
	details := astbuilder.FuncDetails{
		ReceiverIdent: receiverName,
		ReceiverType:  astmodel.NewOptionalType(receiver).AsType(codeGenerationContext),
		Name:          fake.name,
	}

	if fake.TypeReturned != nil {
		details.AddReturn(fake.TypeReturned.AsType(codeGenerationContext))
	}

	return details.DefineFunc(), nil
}

func (fake *FakeFunction) Equals(f astmodel.Function, _ astmodel.EqualityOverrides) bool {
	if fake == nil && f == nil {
		return true
	}

	if fake == nil || f == nil {
		return false
	}

	fn, ok := f.(*FakeFunction)
	if !ok {
		return false
	}

	// Check to see if they have the same references
	if !fake.Referenced.Equals(fn.Referenced) {
		return false
	}

	// Check to see if they have the same imports
	if fake.Imported.Length() != fn.Imported.Length() {
		return false
	}

	for _, imp := range fake.Imported.AsSlice() {
		if !fn.Imported.Contains(imp) {
			return false
		}
	}

	return true
}

func (fake *FakeFunction) ReturnType() astmodel.Type {
	return fake.TypeReturned
}
