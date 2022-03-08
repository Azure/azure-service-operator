/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"sort"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

const JSONUnmarshalFunctionName string = "UnmarshalJSON"

// OneOfJSONUnmarshalFunction is a function for unmarshalling discriminated unions
// (types with only mutually exclusive properties) from JSON
type OneOfJSONUnmarshalFunction struct {
	oneOfObject *astmodel.ObjectType
	idFactory   astmodel.IdentifierFactory
}

// NewOneOfJSONUnmarshalFunction creates a new OneOfJSONUnmarshalFunction struct
func NewOneOfJSONUnmarshalFunction(oneOfObject *astmodel.ObjectType, idFactory astmodel.IdentifierFactory) *OneOfJSONUnmarshalFunction {
	return &OneOfJSONUnmarshalFunction{oneOfObject, idFactory}
}

// Ensure OneOfJSONMarshalFunction implements Function interface correctly
var _ astmodel.Function = (*OneOfJSONUnmarshalFunction)(nil)

func (f *OneOfJSONUnmarshalFunction) Name() string {
	return JSONUnmarshalFunctionName
}

// Equals determines if this function is equal to the passed in function
func (f *OneOfJSONUnmarshalFunction) Equals(other astmodel.Function, overrides astmodel.EqualityOverrides) bool {
	if o, ok := other.(*OneOfJSONUnmarshalFunction); ok {
		return f.oneOfObject.Equals(o.oneOfObject, overrides)
	}

	return false
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (f *OneOfJSONUnmarshalFunction) References() astmodel.TypeNameSet {
	return nil
}

// AsFunc returns the function as a go dst
func (f *OneOfJSONUnmarshalFunction) AsFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName) *dst.FuncDecl {

	jsonPackage := codeGenerationContext.MustGetImportedPackageName(astmodel.JsonReference)
	receiverName := f.idFactory.CreateReceiver(receiver.Name())

	allDefinitions := codeGenerationContext.GetAllReachableDefinitions()
	discrimJSONName, valuesMapping := astmodel.DetermineDiscriminantAndValues(f.oneOfObject, allDefinitions)

	paramName := "data"
	mapName := "rawJson"
	discrimName := "discriminator"
	errName := "err"

	statements := []dst.Stmt{
		astbuilder.LocalVariableDeclaration(mapName, &dst.MapType{Key: dst.NewIdent("string"), Value: dst.NewIdent("interface{}")}, ""),
		astbuilder.ShortDeclaration(errName,
			astbuilder.CallQualifiedFunc(jsonPackage, "Unmarshal", dst.NewIdent(paramName), astbuilder.AddrOf(dst.NewIdent(mapName)))),
		astbuilder.CheckErrorAndReturn(),
		astbuilder.ShortDeclaration(discrimName, &dst.IndexExpr{
			X:     dst.NewIdent(mapName),
			Index: astbuilder.StringLiteral(discrimJSONName),
		}),
	}

	// must order this for consistent output
	values := make([]string, 0, len(valuesMapping))
	for value := range valuesMapping {
		values = append(values, value)
	}

	sort.Strings(values)

	for _, value := range values {
		prop := valuesMapping[value]
		selector := func() dst.Expr {
			return &dst.SelectorExpr{X: dst.NewIdent(receiverName), Sel: dst.NewIdent(string(prop.PropertyName))}
		}

		statements = append(statements,
			astbuilder.IfEqual(
				dst.NewIdent(discrimName),
				astbuilder.TextLiteral(value),
				astbuilder.SimpleAssignment(selector(), astbuilder.AddrOf(&dst.CompositeLit{Type: dst.NewIdent(prop.TypeName.Name())})),
				astbuilder.Returns(astbuilder.CallQualifiedFunc(jsonPackage, "Unmarshal", dst.NewIdent(paramName), selector())),
			))
	}

	statements = append(statements, astbuilder.ReturnNoError()) // TODO

	fn := &astbuilder.FuncDetails{
		Name:          f.Name(),
		ReceiverIdent: receiverName,
		ReceiverType:  astbuilder.Dereference(receiver.AsType(codeGenerationContext)),
		Body:          statements,
	}
	fn.AddParameter(paramName, &dst.ArrayType{Elt: dst.NewIdent("byte")})
	fn.AddComments(fmt.Sprintf("unmarshals the %s", receiver.Name()))
	fn.AddReturns("error")
	return fn.DefineFunc()
}

// RequiredPackageReferences returns a set of references to packages required by this
func (f *OneOfJSONUnmarshalFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet(astmodel.JsonReference)
}
