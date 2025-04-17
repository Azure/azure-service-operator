/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const JSONMarshalFunctionName string = "MarshalJSON"

// OneOfJSONMarshalFunction is a function for marshalling discriminated unions
// (types with only mutually exclusive properties) to JSON
type OneOfJSONMarshalFunction struct {
	oneOfObject *astmodel.ObjectType
	idFactory   astmodel.IdentifierFactory
}

// NewOneOfJSONMarshalFunction creates a new OneOfJSONMarshalFunction struct
func NewOneOfJSONMarshalFunction(
	oneOfObject *astmodel.ObjectType,
	idFactory astmodel.IdentifierFactory,
) *OneOfJSONMarshalFunction {
	return &OneOfJSONMarshalFunction{
		oneOfObject,
		idFactory,
	}
}

// Ensure OneOfJSONMarshalFunction implements Function interface correctly
var _ astmodel.Function = (*OneOfJSONMarshalFunction)(nil)

func (f *OneOfJSONMarshalFunction) Name() string {
	return JSONMarshalFunctionName
}

// Equals determines if this function is equal to the passed in function
func (f *OneOfJSONMarshalFunction) Equals(other astmodel.Function, overrides astmodel.EqualityOverrides) bool {
	if o, ok := other.(*OneOfJSONMarshalFunction); ok {
		return f.oneOfObject.Equals(o.oneOfObject, overrides)
	}

	return false
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (f *OneOfJSONMarshalFunction) References() astmodel.TypeNameSet {
	return nil
}

// AsFunc returns the function as a go dst
func (f *OneOfJSONMarshalFunction) AsFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.InternalTypeName,
) (*dst.FuncDecl, error) {
	jsonPackage := codeGenerationContext.MustGetImportedPackageName(astmodel.JSONReference)

	receiverName := f.idFactory.CreateReceiver(receiver.Name())
	receiverExpr, err := receiver.AsTypeExpr(codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", receiver)
	}

	// Find any root properties that need special handling
	props := f.oneOfObject.Properties().AsSlice()
	rootProperties := astmodel.NewPropertySet()
	defs := codeGenerationContext.GetDefinitionsInCurrentPackage()
	for _, prop := range props {
		if !astmodel.IsOneOfLeafProperty(prop, defs) {
			rootProperties.Add(prop)
		}
	}

	// For each leaf property, create a check with a marshal call
	statements := make([]dst.Stmt, 0, len(props))
	for _, property := range props {
		if rootProperties.ContainsProperty(property.PropertyName()) {
			continue
		}

		receiverIdent := dst.NewIdent(receiverName)
		propName := string(property.PropertyName())

		ret := astbuilder.Returns(
			astbuilder.CallQualifiedFunc(
				jsonPackage,
				"Marshal",
				astbuilder.Selector(receiverIdent, propName)))

		ifStatement := astbuilder.IfNotNil(
			astbuilder.Selector(
				receiverIdent,
				string(property.PropertyName())),
			ret)
		ifStatement.Decorations().After = dst.EmptyLine
		statements = append(statements, ifStatement)
	}

	finalReturnStatement := astbuilder.Returns(astbuilder.Nil(), astbuilder.Nil())

	fn := &astbuilder.FuncDetails{
		Name:          f.Name(),
		ReceiverIdent: receiverName,
		ReceiverType:  receiverExpr,
		Body:          astbuilder.Statements(statements, finalReturnStatement),
	}

	fn.AddComments(fmt.Sprintf(
		"defers JSON marshaling to the first non-nil property, because %s represents a discriminated union (JSON OneOf)",
		receiver.Name()))
	fn.AddReturns("[]byte", "error")
	return fn.DefineFunc(), nil
}

// RequiredPackageReferences returns a set of references to packages required by this
func (f *OneOfJSONMarshalFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet(astmodel.JSONReference)
}
