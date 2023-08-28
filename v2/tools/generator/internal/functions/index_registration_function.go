/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// IndexRegistrationFunction is a function for registering an index for a given property chain. The leaf of the property chain
// is expected to implement the genruntime.Indexer interface
type IndexRegistrationFunction struct {
	name             string
	resourceTypeName astmodel.TypeName
	propertyChain    []*astmodel.PropertyDefinition
	indexKey         string
	idFactory        astmodel.IdentifierFactory
}

// NewIndexRegistrationFunction returns a new index registration function
func NewIndexRegistrationFunction(
	idFactory astmodel.IdentifierFactory,
	name string,
	resourceTypeName astmodel.TypeName,
	indexKey string,
	propertyChain []*astmodel.PropertyDefinition,
) *IndexRegistrationFunction {
	return &IndexRegistrationFunction{
		idFactory:        idFactory,
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
	codeGenerationContext *astmodel.CodeGenerationContext,
	_ astmodel.TypeName,
) (*dst.FuncDecl, error) {
	rawObjName := "rawObj"
	objName := "obj"

	// obj, ok := rawObj.(*<type>)
	cast := astbuilder.TypeAssert(
		dst.NewIdent(objName),
		dst.NewIdent(rawObjName),
		astbuilder.PointerTo(f.resourceTypeName.AsType(codeGenerationContext)))

	// if !ok { return nil }
	checkAssert := astbuilder.ReturnIfNotOk(astbuilder.Nil())
	specSelector := astbuilder.Selector(dst.NewIdent(objName), "Spec")

	var stmts []dst.Stmt
	if f.isIndexSingleValue() {
		stmts = f.singleValue(specSelector)
	} else {
		var err error
		stmts, err = f.multipleValues(specSelector)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to generate multiple value index registration function for %s", f.resourceTypeName)
		}
	}

	fn := &astbuilder.FuncDetails{
		Name: f.Name(),
		Body: astbuilder.Statements(
			cast,
			checkAssert,
			stmts),
	}

	fn.AddParameter(rawObjName, astmodel.ControllerRuntimeObjectType.AsType(codeGenerationContext))

	fn.AddReturn(&dst.ArrayType{Elt: dst.NewIdent("string")})

	pkg := codeGenerationContext.MustGetImportedPackageName(f.resourceTypeName.PackageReference())

	fn.AddComments(fmt.Sprintf("an index function for %s.%s %s", pkg, f.resourceTypeName.Name(), f.indexKey))
	return fn.DefineFunc(), nil
}

// singleValue is used when the property path contains no collections
func (f *IndexRegistrationFunction) singleValue(selector *dst.SelectorExpr) []dst.Stmt {
	propNames := make([]string, 0, len(f.propertyChain))
	nilGuards := make([]dst.Stmt, 0, len(f.propertyChain))
	for _, prop := range f.propertyChain {
		propNames = append(propNames, prop.PropertyName().String())
		intermediateSelector := astbuilder.Selector(selector, propNames...)
		nilGuards = append(nilGuards, astbuilder.ReturnIfNil(intermediateSelector, astbuilder.Nil()))
	}

	// return obj.spec.<property>.Index()
	finalSelector := astbuilder.Selector(selector, propNames...)
	ret := astbuilder.Returns(astbuilder.CallExpr(finalSelector, "Index"))

	return astbuilder.Statements(
		nilGuards,
		ret,
	)
}

// multipleValues is used when there are collections in the property path.
func (f *IndexRegistrationFunction) multipleValues(selector *dst.SelectorExpr) ([]dst.Stmt, error) {

	// var result []string
	resultVar := astbuilder.LocalVariableDeclaration(
		"result",
		&dst.ArrayType{
			Elt: dst.NewIdent("string"),
		},
		"")

	locals := astmodel.NewKnownLocalsSet(f.idFactory)

	// This makes a collection of statements that add items to result, defined above.
	stmts, err := f.makeStatements(
		dst.NewIdent("result"),
		locals,
		selector,
		f.propertyChain,
		astbuilder.Returns(astbuilder.Nil()),
	)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to make statements for property chain %v", f.propertyChain)
	}

	ret := astbuilder.Returns(dst.NewIdent("result"))

	return astbuilder.Statements(
		resultVar,
		stmts,
		ret), nil
}

func (f *IndexRegistrationFunction) makeStatements(
	result *dst.Ident,
	locals *astmodel.KnownLocalsSet,
	ident dst.Expr,
	remainingChain []*astmodel.PropertyDefinition,
	nilHandler dst.Stmt,
) ([]dst.Stmt, error) {
	if len(remainingChain) == 0 {
		return astbuilder.Statements(astbuilder.AppendSliceToSlice(result, astbuilder.CallExpr(ident, "Index"))), nil
	}

	p := remainingChain[0]

	t := p.PropertyType()
	if optionalType, ok := astmodel.AsOptionalType(t); ok {
		t = optionalType.Element()
	}

	if _, ok := astmodel.AsTypeName(t); ok {
		ident = astbuilder.Selector(ident, p.PropertyName().String())
		next := astbuilder.IfNil(ident, nilHandler)
		stmts, err := f.makeStatements(result, locals, ident, remainingChain[1:], nilHandler)
		if err != nil {
			return nil, err
		}

		return astbuilder.Statements(next, stmts), nil
	}

	if _, ok := astmodel.AsArrayType(t); ok {
		return f.handleArray(result, locals, ident, remainingChain)
	}

	if _, ok := astmodel.AsMapType(t); ok {
		return f.handleMap(result, locals, ident, remainingChain)
	}

	return nil, errors.Errorf("can't produce index registration function for type %T", t)
}

func (f *IndexRegistrationFunction) handleArray(
	result *dst.Ident,
	locals *astmodel.KnownLocalsSet,
	ident dst.Expr,
	remainingChain []*astmodel.PropertyDefinition,
) ([]dst.Stmt, error) {
	p := remainingChain[0]

	local := locals.CreateSingularLocal(p.PropertyName().String(), "Item")
	nilHandler := astbuilder.Continue()
	loopStatements, err := f.makeStatements(result, locals.Clone(), dst.NewIdent(local), remainingChain[1:], nilHandler)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to make loop statements for array property %s", p.PropertyName())
	}

	loop := astbuilder.IterateOverSlice(
		local,
		astbuilder.Selector(ident, p.PropertyName().String()),
		loopStatements...)
	return astbuilder.Statements(loop), nil
}

func (f *IndexRegistrationFunction) handleMap(
	result *dst.Ident,
	locals *astmodel.KnownLocalsSet,
	ident dst.Expr,
	remainingChain []*astmodel.PropertyDefinition,
) ([]dst.Stmt, error) {
	p := remainingChain[0]

	// key in the loop would always be un-used.
	key := "_"
	value := locals.CreateLocal("value")
	nilHandler := astbuilder.Continue()
	loopStatements, err := f.makeStatements(result, locals.Clone(), dst.NewIdent(value), remainingChain[1:], nilHandler)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to make loop statements for map property %s", p.PropertyName())
	}

	loop := astbuilder.IterateOverMapWithValue(
		key,
		value,
		astbuilder.Selector(ident, p.PropertyName().String()),
		loopStatements...)

	return astbuilder.Statements(loop), nil
}

func (f *IndexRegistrationFunction) isIndexSingleValue() bool {
	for _, prop := range f.propertyChain {
		_, isMap := astmodel.AsMapType(prop.PropertyType())
		_, isArray := astmodel.AsArrayType(prop.PropertyType())
		if isMap || isArray {
			return false
		}
	}

	return true
}
