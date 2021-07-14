/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
)

const JSONMarshalFunctionName string = "MarshalJSON"

// OneOfJSONMarshalFunction is a function for marshalling discriminated unions
// (types with only mutually exclusive properties) to JSON
type OneOfJSONMarshalFunction struct {
	oneOfObject *ObjectType
	idFactory   IdentifierFactory // TODO: It's this or pass it in the AsFunc method
}

// NewOneOfJSONMarshalFunction creates a new OneOfJSONMarshalFunction struct
func NewOneOfJSONMarshalFunction(oneOfObject *ObjectType, idFactory IdentifierFactory) *OneOfJSONMarshalFunction {
	return &OneOfJSONMarshalFunction{oneOfObject, idFactory}
}

// Ensure OneOfJSONMarshalFunction implements Function interface correctly
var _ Function = (*OneOfJSONMarshalFunction)(nil)

func (f *OneOfJSONMarshalFunction) Name() string {
	return JSONMarshalFunctionName
}

// Equals determines if this function is equal to the passed in function
func (f *OneOfJSONMarshalFunction) Equals(other Function) bool {
	if o, ok := other.(*OneOfJSONMarshalFunction); ok {
		return f.oneOfObject.Equals(o.oneOfObject)
	}

	return false
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (f *OneOfJSONMarshalFunction) References() TypeNameSet {
	return nil
}

// AsFunc returns the function as a go dst
func (f *OneOfJSONMarshalFunction) AsFunc(
	codeGenerationContext *CodeGenerationContext,
	receiver TypeName) *dst.FuncDecl {

	jsonPackage := codeGenerationContext.MustGetImportedPackageName(JsonReference)

	receiverName := f.idFactory.CreateIdentifier(receiver.name, NotExported)

	var statements []dst.Stmt

	for _, property := range f.oneOfObject.Properties() {

		ifStatement := dst.IfStmt{
			Cond: astbuilder.NotNil(
				astbuilder.Selector(
					dst.NewIdent(receiverName),
					string(property.propertyName))),
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							&dst.CallExpr{
								Fun: astbuilder.Selector(
									dst.NewIdent(jsonPackage),
									"Marshal"),
								Args: []dst.Expr{
									astbuilder.Selector(
										dst.NewIdent(receiverName),
										string(property.propertyName)),
								},
							},
						},
					},
				},
			},
		}

		statements = append(statements, &ifStatement)
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
		receiver.name))
	fn.AddReturns("[]byte", "error")
	return fn.DefineFunc()
}

// RequiredPackageReferences returns a set of references to packages required by this
func (f *OneOfJSONMarshalFunction) RequiredPackageReferences() *PackageReferenceSet {
	return NewPackageReferenceSet(MakeExternalPackageReference("encoding/json"))
}
