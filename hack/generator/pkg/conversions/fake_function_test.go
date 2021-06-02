/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// FakeFunction is a fake function that can be used for testing purposes
type FakeFunction struct {
	name       string
	Imported   *astmodel.PackageReferenceSet
	Referenced astmodel.TypeNameSet
	returnType astmodel.Type
	idFactory  astmodel.IdentifierFactory
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

func (fake *FakeFunction) AsFunc(generationContext *astmodel.CodeGenerationContext, reciever astmodel.TypeName) *dst.FuncDecl {
	recieverName := fake.idFactory.CreateIdentifier(reciever.Name(), astmodel.NotExported)
	details := astbuilder.FuncDetails{
		ReceiverIdent: recieverName,
		ReceiverType:  astmodel.NewOptionalType(reciever).AsType(generationContext),
		Name:          fake.name,
	}

	if fake.returnType != nil {
		details.AddReturn(fake.returnType.AsType(generationContext))
	}

	return details.DefineFunc()
}

func (fake *FakeFunction) Equals(f astmodel.Function) bool {
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
	return fake.returnType
}

