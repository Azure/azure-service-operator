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

type IndexRegistrationFunction struct {
	name             string
	resourceTypeName astmodel.TypeName
	propertyChain    []*astmodel.PropertyDefinition
	indexKey         string
}

// NewIndexRegistrationFunction returns a new index registration function
func NewIndexRegistrationFunction(
	name string,
	resourceTypeName astmodel.TypeName,
	indexKey string,
	propertyChain []*astmodel.PropertyDefinition,
) *IndexRegistrationFunction {
	return &IndexRegistrationFunction{
		name:             name,
		resourceTypeName: resourceTypeName,
		// TODO: Technically you can derive the key from the chain so maybe should just pass one
		indexKey:      indexKey,
		propertyChain: propertyChain,
	}
}

// Ensure IndexRegistrationFunction implements Function interface correctly
var _ astmodel.Function = &IndexRegistrationFunction{}

func (f *IndexRegistrationFunction) Name() string {
	return f.name
}

// IndexKey returns the index key used in the index function
func (f *IndexRegistrationFunction) IndexKey() string {
	return f.indexKey
}

// Equals determines if this function is equal to the passed in function
func (f *IndexRegistrationFunction) Equals(other astmodel.Function, overrides astmodel.EqualityOverrides) bool {
	if o, ok := other.(*IndexRegistrationFunction); ok {
		return f.resourceTypeName.Equals(o.resourceTypeName, overrides)
	}

	return false
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (f *IndexRegistrationFunction) References() astmodel.TypeNameSet {
	return nil
}

// RequiredPackageReferences returns a set of references to packages required by this
func (f *IndexRegistrationFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return astmodel.NewPackageReferenceSet(astmodel.ControllerRuntimeClient)
}

// AsFunc returns the function as a go dst
func (f *IndexRegistrationFunction) AsFunc(
	genContext *astmodel.CodeGenerationContext,
	_ astmodel.TypeName,
) *dst.FuncDecl {
	rawObjName := "rawObj"
	objName := "obj"

	// obj, ok := rawObj.(*<type>)
	cast := astbuilder.TypeAssert(
		dst.NewIdent(objName),
		dst.NewIdent(rawObjName),
		astbuilder.Dereference(f.resourceTypeName.AsType(genContext)))

	// if !ok { return nil }
	checkAssert := astbuilder.ReturnIfNotOk(astbuilder.Nil())
	specSelector := astbuilder.Selector(dst.NewIdent(objName), "Spec")

	// Create a series of guards to protect against the possibility of a nil dereference
	// if obj.Spec.AdministratorLoginPassword == nil {
	//	return nil
	// }
	propNames := make([]string, 0, len(f.propertyChain))
	nilGuards := make([]dst.Stmt, 0, len(f.propertyChain))
	for _, prop := range f.propertyChain {
		propNames = append(propNames, prop.PropertyName().String())
		intermediateSelector := astbuilder.Selector(specSelector, propNames...)
		nilGuards = append(nilGuards, astbuilder.ReturnIfNil(intermediateSelector, astbuilder.Nil()))
	}

	// return []string{obj.Spec.<property>}
	propNames = append(propNames, "Name") // Add the ".Name" selector to select the secrets name field
	finalSelector := astbuilder.Selector(specSelector, propNames...)
	ret := astbuilder.Returns(astbuilder.SliceLiteral(dst.NewIdent("string"), finalSelector))

	fn := &astbuilder.FuncDetails{
		Name: f.Name(),
		Body: astbuilder.Statements(
			cast,
			checkAssert,
			nilGuards,
			ret),
	}

	fn.AddParameter(rawObjName, astmodel.ControllerRuntimeObjectType.AsType(genContext))

	fn.AddReturn(&dst.ArrayType{Elt: dst.NewIdent("string")})

	pkg := genContext.MustGetImportedPackageName(f.resourceTypeName.PackageReference)

	fn.AddComments(fmt.Sprintf("an index function for %s.%s %s", pkg, f.resourceTypeName.Name(), f.indexKey))
	return fn.DefineFunc()
}
