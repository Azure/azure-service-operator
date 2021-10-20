/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/token"
	"sort"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
)

const JSONUnmarshalFunctionName string = "UnmarshalJSON"

// OneOfJSONUnmarshalFunction is a function for unmarshalling discriminated unions
// (types with only mutually exclusive properties) from JSON
type OneOfJSONUnmarshalFunction struct {
	oneOfObject *ObjectType
	idFactory   IdentifierFactory // TODO: It's this or pass it in the AsFunc method
}

// NewOneOfJSONUnmarshalFunction creates a new OneOfJSONUnmarshalFunction struct
func NewOneOfJSONUnmarshalFunction(oneOfObject *ObjectType, idFactory IdentifierFactory) *OneOfJSONUnmarshalFunction {
	return &OneOfJSONUnmarshalFunction{oneOfObject, idFactory}
}

// Ensure OneOfJSONMarshalFunction implements Function interface correctly
var _ Function = (*OneOfJSONUnmarshalFunction)(nil)

func (f *OneOfJSONUnmarshalFunction) Name() string {
	return JSONUnmarshalFunctionName
}

// Equals determines if this function is equal to the passed in function
func (f *OneOfJSONUnmarshalFunction) Equals(other Function, overrides EqualityOverrides) bool {
	if o, ok := other.(*OneOfJSONMarshalFunction); ok {
		return f.oneOfObject.Equals(o.oneOfObject, overrides)
	}

	return false
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (f *OneOfJSONUnmarshalFunction) References() TypeNameSet {
	return nil
}

func (f *OneOfJSONUnmarshalFunction) resolveOneOfMemberToObjectType(t Type, types Types) (TypeName, *ObjectType) {
	propType, err := types.FullyResolve(t)
	if err != nil {
		panic(err) // type should not contain unresolvable references at this point
	}

	optionalType, ok := propType.(*OptionalType)
	if !ok {
		panic(fmt.Sprintf("OneOf contained non-optional type %s", propType.String()))
	}

	typeName, ok := optionalType.Element().(TypeName)
	if !ok {
		panic("Expected OneOf to have pointer to TypeName")
	}

	resolvedInnerOptional, err := types.FullyResolve(typeName)
	if err != nil {
		panic(err) // types should not contain unresolvable references at this point
	}

	propObjType, ok := resolvedInnerOptional.(*ObjectType)
	if !ok {
		panic(fmt.Sprintf("OneOf contained non-object type %s", propType.String()))
	}

	return typeName, propObjType
}

type propNameAndType struct {
	propName PropertyName
	typeName TypeName
}

func (f *OneOfJSONUnmarshalFunction) getDiscriminatorMapping(propName PropertyName, types Types) map[string]propNameAndType {
	result := make(map[string]propNameAndType)
	for _, prop := range f.oneOfObject.properties {
		propObjTypeName, propObjType := f.resolveOneOfMemberToObjectType(prop.propertyType, types)

		potentialDiscriminatorProp, ok := propObjType.properties[propName]
		if !ok {
			return nil
		}

		potentialDiscriminatorType, err := types.FullyResolve(potentialDiscriminatorProp.PropertyType())
		if err != nil {
			panic(err) // should not be unresolvable
		}

		enumType, ok := potentialDiscriminatorType.(*EnumType)
		if !ok {
			return nil
		}

		if len(enumType.options) != 1 {
			return nil
		}

		enumValue := enumType.options[0].Value
		if _, ok := result[enumValue]; ok {
			return nil
		}

		result[enumValue] = propNameAndType{prop.propertyName, propObjTypeName}
	}

	return result
}

func (f *OneOfJSONUnmarshalFunction) determineDiscriminant(allTypes Types) (string, map[string]propNameAndType) {
	var firstMember *ObjectType
	for _, prop := range f.oneOfObject.properties {
		_, firstMember = f.resolveOneOfMemberToObjectType(prop.PropertyType(), allTypes)
		break
	}

	for _, prop := range firstMember.properties {
		mapping := f.getDiscriminatorMapping(prop.PropertyName(), allTypes)
		if mapping != nil {
			return prop.tags["json"][0], mapping
		}
	}

	panic("unable to determine a discriminator property for oneOf type")
}

// AsFunc returns the function as a go dst
func (f *OneOfJSONUnmarshalFunction) AsFunc(
	codeGenerationContext *CodeGenerationContext,
	receiver TypeName) *dst.FuncDecl {

	jsonPackage := codeGenerationContext.MustGetImportedPackageName(JsonReference)
	receiverName := f.idFactory.CreateIdentifier(receiver.name, NotExported)

	allTypes := codeGenerationContext.GetAllReachableTypes()
	discrimJSONName, valuesMapping := f.determineDiscriminant(allTypes)

	paramName := "data"
	mapName := "rawJson"
	discrimName := "discriminator"
	errName := "err"

	statements := []dst.Stmt{
		astbuilder.LocalVariableDeclaration(mapName, &dst.MapType{Key: dst.NewIdent("string"), Value: dst.NewIdent("interface{}")}, ""),
		astbuilder.ShortDeclaration(errName,
			astbuilder.CallQualifiedFunc(jsonPackage, "Unmarshal", dst.NewIdent(paramName), &dst.UnaryExpr{Op: token.AND, X: dst.NewIdent(mapName)})),
		astbuilder.CheckErrorAndReturn(),
		astbuilder.ShortDeclaration(discrimName, &dst.IndexExpr{
			X:     dst.NewIdent(mapName),
			Index: &dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", discrimJSONName)},
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
			return &dst.SelectorExpr{X: dst.NewIdent(receiverName), Sel: dst.NewIdent(string(prop.propName))}
		}

		statements = append(statements,
			astbuilder.IfEqual(
				dst.NewIdent(discrimName),
				&dst.BasicLit{Kind: token.STRING, Value: value},
				astbuilder.SimpleAssignment(selector(), &dst.UnaryExpr{
					Op: token.AND,
					X:  &dst.CompositeLit{Type: dst.NewIdent(prop.typeName.Name())},
				}),
				astbuilder.Returns(astbuilder.CallQualifiedFunc(jsonPackage, "Unmarshal", dst.NewIdent(paramName), selector())),
			))
	}

	statements = append(statements, astbuilder.ReturnNoError()) // TODO

	fn := &astbuilder.FuncDetails{
		Name:          f.Name(),
		ReceiverIdent: receiverName,
		ReceiverType:  receiver.AsType(codeGenerationContext),
		Body:          statements,
	}
	fn.AddParameter(paramName, &dst.ArrayType{Elt: dst.NewIdent("byte")})
	fn.AddComments(fmt.Sprintf("unmarshals the %s", receiver.name))
	fn.AddReturns("error")
	return fn.DefineFunc()
}

// RequiredPackageReferences returns a set of references to packages required by this
func (f *OneOfJSONUnmarshalFunction) RequiredPackageReferences() *PackageReferenceSet {
	return NewPackageReferenceSet(MakeExternalPackageReference("encoding/json"))
}
