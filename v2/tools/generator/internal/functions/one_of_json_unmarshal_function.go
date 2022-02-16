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

func (f *OneOfJSONUnmarshalFunction) resolveOneOfMemberToObjectType(t astmodel.Type, definitions astmodel.TypeDefinitionSet) (astmodel.TypeName, *astmodel.ObjectType) {
	// OneOfs are expected to contain properties that are:
	// pointer to typename to objectType

	propType, err := definitions.FullyResolve(t)
	if err != nil {
		panic(err) // type should not contain unresolvable references at this point
	}

	optionalType, ok := propType.(*astmodel.OptionalType)
	if !ok {
		panic(fmt.Sprintf("OneOf contained non-optional type %s", propType.String()))
	}

	typeName, ok := optionalType.Element().(astmodel.TypeName)
	if !ok {
		panic("Expected OneOf to have pointer to TypeName")
	}

	resolvedInnerOptional, err := definitions.FullyResolve(typeName)
	if err != nil {
		panic(err) // definitions should not contain unresolvable references at this point
	}

	propObjType, ok := resolvedInnerOptional.(*astmodel.ObjectType)
	if !ok {
		panic(fmt.Sprintf("OneOf contained non-object type %s", propType.String()))
	}

	return typeName, propObjType
}

type propNameAndType struct {
	propName astmodel.PropertyName
	typeName astmodel.TypeName // the name of the type inside the pointer type
}

func (f *OneOfJSONUnmarshalFunction) getDiscriminatorMapping(
	propName astmodel.PropertyName,
	definitions astmodel.TypeDefinitionSet) map[string]propNameAndType {

	result := make(map[string]propNameAndType)
	for _, prop := range f.oneOfObject.Properties() {
		propObjTypeName, propObjType := f.resolveOneOfMemberToObjectType(prop.PropertyType(), definitions)

		potentialDiscriminatorProp, ok := propObjType.Property(propName)
		if !ok {
			return nil
		}

		potentialDiscriminatorType, err := definitions.FullyResolve(potentialDiscriminatorProp.PropertyType())
		if err != nil {
			panic(err) // should not be unresolvable
		}

		enumType, ok := potentialDiscriminatorType.(*astmodel.EnumType)
		if !ok {
			return nil // if not an enum type cannot be used as discriminator
		}

		enumOptions := enumType.Options()
		if len(enumOptions) != 1 {
			return nil // if enum type has more than one value, cannot be used as discriminator
			// not entirely true since the options could all have distinct sets, but this is good-
			// enough for now
		}

		enumValue := enumOptions[0].Value
		if _, ok := result[enumValue]; ok {
			return nil // if values are not distinct for each member, cannot be used as discriminator
		}

		result[enumValue] = propNameAndType{prop.PropertyName(), propObjTypeName}
	}

	return result
}

func (f *OneOfJSONUnmarshalFunction) determineDiscriminant(definitions astmodel.TypeDefinitionSet) (string, map[string]propNameAndType) {
	// grab out the first member of the OneOf
	var firstMember *astmodel.ObjectType
	for _, prop := range f.oneOfObject.Properties() {
		_, firstMember = f.resolveOneOfMemberToObjectType(prop.PropertyType(), definitions)
		break
	}

	// try to find a discriminator property out of the properties on the first member
	for _, prop := range firstMember.Properties() {
		mapping := f.getDiscriminatorMapping(prop.PropertyName(), definitions)
		if mapping != nil {
			jsonTag, ok := prop.Tag("json")
			if !ok {
				// in reality every property will have a JSON tag
				panic("discriminator property had no JSON tag")
			}

			// first part of JSON tag is the JSON name
			return jsonTag[0], mapping
		}
	}

	panic("unable to determine a discriminator property for oneOf type")
}

// AsFunc returns the function as a go dst
func (f *OneOfJSONUnmarshalFunction) AsFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName) *dst.FuncDecl {

	jsonPackage := codeGenerationContext.MustGetImportedPackageName(astmodel.JsonReference)
	receiverName := f.idFactory.CreateReceiver(receiver.Name())

	allDefinitions := codeGenerationContext.GetAllReachableDefinitions()
	discrimJSONName, valuesMapping := f.determineDiscriminant(allDefinitions)

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
			return &dst.SelectorExpr{X: dst.NewIdent(receiverName), Sel: dst.NewIdent(string(prop.propName))}
		}

		statements = append(statements,
			astbuilder.IfEqual(
				dst.NewIdent(discrimName),
				astbuilder.TextLiteral(value),
				astbuilder.SimpleAssignment(selector(), astbuilder.AddrOf(&dst.CompositeLit{Type: dst.NewIdent(prop.typeName.Name())})),
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
