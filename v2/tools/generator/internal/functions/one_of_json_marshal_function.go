/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"

	"github.com/dave/dst"

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
func NewOneOfJSONMarshalFunction(oneOfObject *astmodel.ObjectType, idFactory astmodel.IdentifierFactory) *OneOfJSONMarshalFunction {
	return &OneOfJSONMarshalFunction{oneOfObject, idFactory}
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
func (f *OneOfJSONMarshalFunction) References() astmodel.TypeNameSet[astmodel.TypeName] {
	return nil
}

// AsFunc returns the function as a go dst
func (f *OneOfJSONMarshalFunction) AsFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
) *dst.FuncDecl {
	jsonPackage := codeGenerationContext.MustGetImportedPackageName(astmodel.JsonReference)

	receiverName := f.idFactory.CreateReceiver(receiver.Name())

	props := f.oneOfObject.Properties().AsSlice()
	statements := make([]dst.Stmt, 0, len(props))
	for _, property := range props {
		ret := astbuilder.Returns(
			astbuilder.CallQualifiedFunc(
				jsonPackage,
				"Marshal",
				astbuilder.Selector(dst.NewIdent(receiverName), string(property.PropertyName()))))
		ifStatement := astbuilder.IfNotNil(
			astbuilder.Selector(
				dst.NewIdent(receiverName),
				string(property.PropertyName())),
			ret)
		statements = append(statements, ifStatement)
	}

	finalReturnStatement := &dst.ReturnStmt{
		Results: []dst.Expr{astbuilder.Nil(), astbuilder.Nil()},
	}
	statements = append(statements, finalReturnStatement)

	fn := &astbuilder.FuncDetails{
		Name:          f.Name(),
		ReceiverIdent: receiverName,
		ReceiverType:  receiver.AsType(codeGenerationContext),
		Body:          statements,
	}

	fn.AddComments(fmt.Sprintf(
		"defers JSON marshaling to the first non-nil property, because %s represents a discriminated union (JSON OneOf)",
		receiver.Name()))
	fn.AddReturns("[]byte", "error")
	return fn.DefineFunc()
}

// RequiredPackageReferences returns a set of references to packages required by this
func (f *OneOfJSONMarshalFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet(astmodel.JsonReference)
}
