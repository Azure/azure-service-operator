/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	ast "github.com/dave/dst"
	"go/token"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog/v2"
)

// ObjectSerializationTestCase represents a test that the object can be losslessly serialized to
// JSON and back again
type ObjectSerializationTestCase struct {
	testName   string
	subject    TypeName
	objectType *ObjectType
	idFactory  IdentifierFactory
}

var _ TestCase = &ObjectSerializationTestCase{}

// NewObjectSerializationTestCase creates a new test case for the JSON serialization round tripability of the specified object type
func NewObjectSerializationTestCase(
	name TypeName,
	objectType *ObjectType,
	idFactory IdentifierFactory) *ObjectSerializationTestCase {
	testName := fmt.Sprintf("%v_WhenSerializedToJson_DeserializesAsEqual", name.Name())
	return &ObjectSerializationTestCase{
		testName:   testName,
		subject:    name,
		objectType: objectType,
		idFactory:  idFactory,
	}
}

// Name returns the unique name of this test case
func (o ObjectSerializationTestCase) Name() string {
	return o.testName
}

// References returns the set of types to which this test case refers.
func (o ObjectSerializationTestCase) References() TypeNameSet {
	result := NewTypeNameSet()
	return result
}

// AsFuncs renders the current test case and supporting methods as Go abstract syntax trees
// subject is the name of the type under test
// codeGenerationContext contains reference material to use when generating
func (o ObjectSerializationTestCase) AsFuncs(name TypeName, genContext *CodeGenerationContext) []ast.Decl {

	var errs []error
	properties := o.makePropertyMap()

	// Find all the simple generators (those with no external dependencies)
	simpleGenerators, err := o.createGenerators(properties, genContext, o.createIndependentGenerator)
	if err != nil {
		errs = append(errs, err)
	}

	// Find all the complex generators (dependent on other generators we'll be generating elsewhere)
	relatedGenerators, err := o.createGenerators(properties, genContext, o.createRelatedGenerator)
	if err != nil {
		errs = append(errs, err)
	}

	// Remove properties from our runtime
	o.removeByPackage(properties, GenRuntimeReference)

	// Temporarily remove properties related to support for Arbitrary JSON
	// TODO: Add generators for these properties
	o.removeByPackage(properties, ApiExtensionsReference)
	o.removeByPackage(properties, ApiExtensionsJsonReference)

	// Write errors for any properties we don't handle
	for _, p := range properties {
		errs = append(errs, errors.Errorf("No generator created for %v (%v)", p.PropertyName(), p.PropertyType()))
	}

	var result []ast.Decl

	if len(simpleGenerators) == 0 && len(relatedGenerators) == 0 {
		// No properties that we can generate to test - skip the testing completely
		errs = append(errs, errors.Errorf("No property generators for %v", name))
	} else {
		result = append(result,
			o.createTestRunner(),
			o.createTestMethod())

		addMethod := func(decl ast.Decl, err error) {
			if err != nil {
				errs = append(errs, err)
			} else {
				result = append(result, decl)
			}
		}

		addMethod(o.createGeneratorDeclaration(genContext))
		addMethod(o.createGeneratorMethod(genContext, len(simpleGenerators) > 0, len(relatedGenerators) > 0))

		if len(simpleGenerators) > 0 {
			addMethod(o.createGeneratorsFactoryMethod(o.idOfIndependentGeneratorsFactoryMethod(), simpleGenerators, genContext))
		}

		if len(relatedGenerators) > 0 {
			addMethod(o.createGeneratorsFactoryMethod(o.idOfRelatedGeneratorsFactoryMethod(), relatedGenerators, genContext))
		}
	}

	if len(errs) > 0 {
		i := "issues"
		if len(errs) == 1 {
			i = "issue"
		}

		klog.Warningf("Encountered %d %s creating JSON Serialisation test for %s", len(errs), i, name)
		for _, err := range errs {
			klog.Warning(err)
		}
	}

	return result
}

// RequiredImports returns a set of the package imports required by this test case
func (o ObjectSerializationTestCase) RequiredImports() *PackageImportSet {
	result := NewPackageImportSet()

	// Standard Go Packages
	result.AddImportsOfReferences(JsonReference, ReflectReference, TestingReference)

	// Cmp
	result.AddImportsOfReferences(CmpReference, CmpOptsReference)

	// Gopter
	result.AddImportsOfReferences(GopterReference, GopterGenReference, GopterPropReference)

	// Other References
	result.AddImportOfReference(DiffReference)
	result.AddImportOfReference(PrettyReference)

	for _, prop := range o.objectType.Properties() {
		for _, ref := range prop.PropertyType().RequiredPackageReferences().AsSlice() {
			result.AddImportOfReference(ref)
		}
	}

	return result
}

// Equals determines if this TestCase is equal to another one
func (o ObjectSerializationTestCase) Equals(_ TestCase) bool {
	panic("implement me")
}

// createTestRunner generates the AST for the test runner itself
func (o ObjectSerializationTestCase) createTestRunner() ast.Decl {

	const (
		parameters = "parameters"
		properties = "properties"
		property   = "property"
		testingRun = "testingRun"
	)

	t := ast.NewIdent("t")

	// parameters := gopter.DefaultTestParameters()
	defineParameters := astbuilder.SimpleAssignment(
		ast.NewIdent(parameters),
		token.DEFINE,
		astbuilder.CallQualifiedFunc("gopter", "DefaultTestParameters"))

	configureMaxSize := astbuilder.SimpleAssignment(
		&ast.SelectorExpr{
			X:   ast.NewIdent(parameters),
			Sel: ast.NewIdent("MaxSize"),
		},
		token.ASSIGN,
		astbuilder.IntLiteral(10))

	// properties := gopter.NewProperties(parameters)
	defineProperties := astbuilder.SimpleAssignment(
		ast.NewIdent(properties),
		token.DEFINE,
		astbuilder.CallQualifiedFunc("gopter", "NewProperties", ast.NewIdent(parameters)))

	// partial expression: name of the test
	testName := astbuilder.StringLiteralf("Round trip of %v via JSON returns original", o.Subject())

	// partial expression: prop.ForAll(RunTestForX, XGenerator())
	propForAll := astbuilder.CallQualifiedFunc(
		"prop",
		"ForAll",
		ast.NewIdent(o.idOfTestMethod()),
		astbuilder.CallFunc(o.idOfGeneratorMethod(o.subject)))

	// properties.Property("...", prop.ForAll(RunTestForX, XGenerator())
	defineTestCase := astbuilder.InvokeQualifiedFunc(
		properties,
		property,
		testName,
		propForAll)

	// properties.TestingRun(t)
	runTests := astbuilder.InvokeQualifiedFunc(properties, testingRun, t)

	// Define our function
	fn := astbuilder.NewTestFuncDetails(
		o.testName,
		defineParameters,
		configureMaxSize,
		defineProperties,
		defineTestCase,
		runTests)

	return fn.DefineFunc()
}

// createTestMethod generates the AST for a method to run a single test of JSON serialization
func (o ObjectSerializationTestCase) createTestMethod() ast.Decl {
	const (
		binId        = "bin"
		actualId     = "actual"
		actualFmtId  = "actualFmt"
		matchId      = "match"
		subjectId    = "subject"
		subjectFmtId = "subjectFmt"
		resultId     = "result"
		errId        = "err"
	)

	// bin, err := json.Marshal(subject)
	serialize := astbuilder.SimpleAssignmentWithErr(
		ast.NewIdent(binId),
		token.DEFINE,
		astbuilder.CallQualifiedFunc("json", "Marshal", ast.NewIdent(subjectId)))

	// if err != nil { return err.Error() }
	serializeFailed := astbuilder.ReturnIfNotNil(
		ast.NewIdent(errId),
		astbuilder.CallQualifiedFunc("err", "Error"))

	// var actual X
	declare := astbuilder.NewVariable(actualId, o.subject.name)

	// err = json.Unmarshal(bin, &actual)
	deserialize := astbuilder.SimpleAssignment(
		ast.NewIdent("err"),
		token.ASSIGN,
		astbuilder.CallQualifiedFunc("json", "Unmarshal",
			ast.NewIdent(binId),
			&ast.UnaryExpr{
				Op: token.AND,
				X:  ast.NewIdent(actualId),
			}))

	// if err != nil { return err.Error() }
	deserializeFailed := astbuilder.ReturnIfNotNil(
		ast.NewIdent(errId),
		astbuilder.CallQualifiedFunc("err", "Error"))

	// match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	equateEmpty := astbuilder.CallQualifiedFunc("cmpopts", "EquateEmpty")
	compare := astbuilder.SimpleAssignment(
		ast.NewIdent(matchId),
		token.DEFINE,
		astbuilder.CallQualifiedFunc("cmp", "Equal",
			ast.NewIdent(subjectId),
			ast.NewIdent(actualId),
			equateEmpty))

	// if !match { result := diff.Diff(subject, actual); return result }
	prettyPrint := &ast.IfStmt{
		Cond: &ast.UnaryExpr{
			Op: token.NOT,
			X:  ast.NewIdent(matchId),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				astbuilder.SimpleAssignment(
					ast.NewIdent(actualFmtId),
					token.DEFINE,
					astbuilder.CallQualifiedFunc("pretty", "Sprint", ast.NewIdent(actualId))),
				astbuilder.SimpleAssignment(
					ast.NewIdent(subjectFmtId),
					token.DEFINE,
					astbuilder.CallQualifiedFunc("pretty", "Sprint", ast.NewIdent(subjectId))),
				astbuilder.SimpleAssignment(
					ast.NewIdent(resultId),
					token.DEFINE,
					astbuilder.CallQualifiedFunc("diff", "Diff", ast.NewIdent(subjectFmtId), ast.NewIdent(actualFmtId))),
				astbuilder.Returns(ast.NewIdent(resultId)),
			},
		},
	}

	// return ""
	ret := astbuilder.Returns(astbuilder.StringLiteral(""))

	// Create the function
	fn := &astbuilder.FuncDetails{
		Name: o.idOfTestMethod(),
		Returns: []*ast.Field{
			{
				Type: ast.NewIdent("string"),
			},
		},
		Body: []ast.Stmt{
			serialize,
			serializeFailed,
			declare,
			deserialize,
			deserializeFailed,
			compare,
			prettyPrint,
			ret,
		},
	}
	fn.AddParameter("subject", o.Subject())
	fn.AddComments(fmt.Sprintf(
		"runs a test to see if a specific instance of %v round trips to JSON and back losslessly",
		o.Subject()))

	return fn.DefineFunc()
}

func (o ObjectSerializationTestCase) createGeneratorDeclaration(genContext *CodeGenerationContext) (ast.Decl, error) {
	comment := fmt.Sprintf(
		"Generator of %v instances for property testing - lazily instantiated by %v()",
		o.Subject(),
		o.idOfGeneratorMethod(o.subject))

	gopterPackage, err := genContext.GetImportedPackageName(GopterReference)
	if err != nil {
		return nil, errors.Wrap(err, "unable to generate generator declaration")
	}

	decl := astbuilder.VariableDeclaration(
		o.idOfSubjectGeneratorGlobal(),
		astbuilder.QualifiedTypeName(gopterPackage, "Gen"),
		comment)

	return decl, nil
}

// createGeneratorMethod generates the AST for a method used to populate our generator cache variable on demand
func (o ObjectSerializationTestCase) createGeneratorMethod(ctx *CodeGenerationContext, haveSimpleGenerators bool, haveRelatedGenerators bool) (ast.Decl, error) {

	gopterPackage, err := ctx.GetImportedPackageName(GopterReference)
	if err != nil {
		return nil, errors.Wrapf(err, "error looking up import name for %s", GopterReference)
	}

	genPackage, err := ctx.GetImportedPackageName(GopterGenReference)
	if err != nil {
		return nil, errors.Wrapf(err, "error looking up import name for %s", GopterGenReference)
	}

	fn := &astbuilder.FuncDetails{
		Name: o.idOfGeneratorMethod(o.subject),
		Returns: []*ast.Field{
			{
				Type: astbuilder.QualifiedTypeName(gopterPackage, "Gen"),
			},
		},
	}

	fn.AddComments(
		fmt.Sprintf("returns a generator of %v instances for property testing.", o.Subject()),
		fmt.Sprintf("We first initialize %v with a simplified generator based on the fields with primitive types", o.idOfSubjectGeneratorGlobal()),
		"then replacing it with a more complex one that also handles complex fields.",
		"This ensures any cycles in the object graph properly terminate.",
		"The call to gen.Struct() captures the map, so we have to create a new one for the second generator.")

	// If we have already cached our builder, return it immediately
	earlyReturn := astbuilder.ReturnIfNotNil(
		ast.NewIdent(o.idOfSubjectGeneratorGlobal()),
		ast.NewIdent(o.idOfSubjectGeneratorGlobal()))
	fn.AddStatements(earlyReturn)

	if haveSimpleGenerators {
		// Create a simple version of the generator that does not reference generators for related types
		// This serves to terminate any dependency cycles that might occur during creation of a more fully fledged generator

		makeIndependentMap := astbuilder.SimpleAssignment(
			ast.NewIdent("independentGenerators"),
			token.DEFINE,
			astbuilder.MakeMap(
				ast.NewIdent("string"),
				astbuilder.QualifiedTypeName(gopterPackage, "Gen")))

		addIndependentGenerators := astbuilder.InvokeFunc(
			o.idOfIndependentGeneratorsFactoryMethod(),
			ast.NewIdent("independentGenerators"))

		createIndependentGenerator := astbuilder.SimpleAssignment(
			ast.NewIdent(o.idOfSubjectGeneratorGlobal()),
			token.ASSIGN,
			astbuilder.CallQualifiedFunc(
				genPackage,
				"Struct",
				astbuilder.CallQualifiedFunc("reflect", "TypeOf", &ast.CompositeLit{Type: o.Subject()}),
				ast.NewIdent("independentGenerators")))

		fn.AddStatements(makeIndependentMap, addIndependentGenerators, createIndependentGenerator)
	}

	if haveRelatedGenerators {
		// Define a local that contains all the simple generators
		// Have to call the factory method twice as the simple generator above has captured the map;
		// if we reuse or modify the map, chaos ensues.

		makeAllMap := astbuilder.SimpleAssignment(
			ast.NewIdent("allGenerators"),
			token.DEFINE,
			astbuilder.MakeMap(
				ast.NewIdent("string"),
				astbuilder.QualifiedTypeName(gopterPackage, "Gen")))

		fn.AddStatements(makeAllMap)

		if haveSimpleGenerators {
			addIndependentGenerators := astbuilder.InvokeFunc(
				o.idOfIndependentGeneratorsFactoryMethod(),
				ast.NewIdent("allGenerators"))
			fn.AddStatements(addIndependentGenerators)
		}

		addRelatedGenerators := astbuilder.InvokeFunc(
			o.idOfRelatedGeneratorsFactoryMethod(),
			ast.NewIdent("allGenerators"))

		createFullGenerator := astbuilder.SimpleAssignment(
			ast.NewIdent(o.idOfSubjectGeneratorGlobal()),
			token.ASSIGN,
			astbuilder.CallQualifiedFunc(
				genPackage,
				"Struct",
				astbuilder.CallQualifiedFunc("reflect", "TypeOf", &ast.CompositeLit{Type: o.Subject()}),
				ast.NewIdent("allGenerators")))

		fn.AddStatements(addRelatedGenerators, createFullGenerator)
	}

	// Return the freshly created (and now cached) generator
	normalReturn := astbuilder.Returns(ast.NewIdent(o.idOfSubjectGeneratorGlobal()))
	fn.AddStatements(normalReturn)

	return fn.DefineFunc(), nil
}

// createGeneratorsFactoryMethod generates the AST for a method creating gopter generators
func (o ObjectSerializationTestCase) createGeneratorsFactoryMethod(
	methodName string, generators []ast.Stmt, ctx *CodeGenerationContext) (ast.Decl, error) {

	if len(generators) == 0 {
		// No simple properties, don't generate a method
		return nil, nil
	}

	gopterPackage, err := ctx.GetImportedPackageName(GopterReference)
	if err != nil {
		return nil, errors.Wrapf(err, "error looking up import name %s for factory generation", GopterReference)
	}

	mapType := &ast.MapType{
		Key:   ast.NewIdent("string"),
		Value: astbuilder.QualifiedTypeName(gopterPackage, "Gen"),
	}

	fn := &astbuilder.FuncDetails{
		Name: methodName,
		Body: generators,
	}

	fn.AddComments("is a factory method for creating gopter generators")
	fn.AddParameter("gens", mapType)

	return fn.DefineFunc(), nil
}

// createGenerators creates AST fragments for gopter generators to create values for properties
// properties is a map of properties needing generators; properties handled here are removed from the map.
// genPackageName is the name for the gopter/gen package (not hard coded in case it's renamed for conflict resolution)
// factory is a method for creating generators
func (o ObjectSerializationTestCase) createGenerators(
	properties map[PropertyName]*PropertyDefinition,
	genContext *CodeGenerationContext,
	factory func(name string, propertyType Type, genContext *CodeGenerationContext) (ast.Expr, error)) ([]ast.Stmt, error) {

	gensIdent := ast.NewIdent("gens")

	var handled []PropertyName
	var result []ast.Stmt

	// Iterate over all properties, creating generators where possible
	var errs []error
	for name, prop := range properties {
		g, err := factory(string(name), prop.PropertyType(), genContext)
		if err != nil {
			errs = append(errs, err)
		} else if g != nil {
			insert := astbuilder.InsertMap(
				gensIdent,
				&ast.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf("\"%v\"", prop.PropertyName()),
				},
				g)
			result = append(result, insert)
			handled = append(handled, name)
		}
	}

	// Remove properties we've handled from the map
	for _, name := range handled {
		delete(properties, name)
	}

	return result, kerrors.NewAggregate(errs)
}

// is directly supported by a Gopter generator, returning nil if the property type isn't supported.
func (o ObjectSerializationTestCase) createIndependentGenerator(
	name string,
	propertyType Type,
	genContext *CodeGenerationContext) (ast.Expr, error) {

	genPackage, err := genContext.GetImportedPackageName(GopterGenReference)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to generate independent generator for %s", name)
	}

	// Handle simple primitive properties
	switch propertyType {
	case StringType:
		return astbuilder.CallQualifiedFunc(genPackage, "AlphaString"), nil
	case IntType:
		return astbuilder.CallQualifiedFunc(genPackage, "Int"), nil
	case FloatType:
		return astbuilder.CallQualifiedFunc(genPackage, "Float32"), nil
	case BoolType:
		return astbuilder.CallQualifiedFunc(genPackage, "Bool"), nil
	}

	switch t := propertyType.(type) {
	case TypeName:
		types := genContext.GetTypesInCurrentPackage()
		def, ok := types[t]
		if ok {
			return o.createIndependentGenerator(def.Name().name, def.theType, genContext)
		}
		return nil, nil

	case *EnumType:
		return o.createEnumGenerator(name, genPackage, t)

	case *OptionalType:
		g, err := o.createIndependentGenerator(name, t.Element(), genContext)
		if err != nil {
			return nil, err
		} else if g != nil {
			return astbuilder.CallQualifiedFunc(genPackage, "PtrOf", g), nil
		}

	case *ArrayType:
		g, err := o.createIndependentGenerator(name, t.Element(), genContext)
		if err != nil {
			return nil, err
		} else if g != nil {
			return astbuilder.CallQualifiedFunc(genPackage, "SliceOf", g), nil
		}

	case *MapType:
		keyGen, err := o.createIndependentGenerator(name, t.KeyType(), genContext)
		if err != nil {
			return nil, err
		}

		valueGen, err := o.createIndependentGenerator(name, t.ValueType(), genContext)
		if err != nil {
			return nil, err
		}

		if keyGen != nil && valueGen != nil {
			return astbuilder.CallQualifiedFunc(genPackage, "MapOf", keyGen, valueGen), nil
		}
	}

	// Not a simple property we can handle here
	return nil, nil
}

// defined within the current package, returning nil if the property type isn't supported.
func (o ObjectSerializationTestCase) createRelatedGenerator(
	name string,
	propertyType Type,
	genContext *CodeGenerationContext) (ast.Expr, error) {

	genPackageName, err := genContext.GetImportedPackageName(GopterGenReference)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to generate related generator for %s", name)
	}

	switch t := propertyType.(type) {
	case TypeName:
		_, ok := genContext.GetTypesInPackage(t.PackageReference)
		if ok {
			// This is a type we're defining, so we can create a generator for it
			if t.PackageReference.Equals(genContext.CurrentPackage()) {
				// create a generator for a property referencing a type in this package
				return astbuilder.CallFunc(o.idOfGeneratorMethod(t)), nil
			}

			importName, err := genContext.GetImportedPackageName(t.PackageReference)
			if err != nil {
				return nil, err
			}

			return astbuilder.CallQualifiedFunc(importName, o.idOfGeneratorMethod(t)), nil
		}

		//TODO: Should we invoke a generator for stuff from our runtime package?

		return nil, nil

	case *OptionalType:
		g, err := o.createRelatedGenerator(name, t.Element(), genContext)
		if err != nil {
			return nil, err
		} else if g != nil {
			return astbuilder.CallQualifiedFunc(genPackageName, "PtrOf", g), nil
		}

	case *ArrayType:
		g, err := o.createRelatedGenerator(name, t.Element(), genContext)
		if err != nil {
			return nil, err
		} else if g != nil {
			return astbuilder.CallQualifiedFunc(genPackageName, "SliceOf", g), nil
		}

	case *MapType:
		// We only support primitive types as keys
		keyGen, err := o.createIndependentGenerator(name, t.KeyType(), genContext)
		if err != nil {
			return nil, err
		}

		valueGen, err := o.createRelatedGenerator(name, t.ValueType(), genContext)
		if err != nil {
			return nil, err
		}

		if keyGen != nil && valueGen != nil {
			return astbuilder.CallQualifiedFunc(genPackageName, "MapOf", keyGen, valueGen), nil
		}
	}

	// Not a property we can handle here
	return nil, nil
}

func (o *ObjectSerializationTestCase) removeByPackage(
	properties map[PropertyName]*PropertyDefinition,
	ref PackageReference) {

	// Work out which properties need to be removed because their types come from the specified package
	var toRemove []PropertyName
	for name, prop := range properties {
		propertyType := prop.PropertyType()
		refs := propertyType.RequiredPackageReferences()
		if refs.Contains(ref) {
			toRemove = append(toRemove, name)
		}
	}

	// Remove them
	for _, name := range toRemove {
		delete(properties, name)
	}
}

func (o *ObjectSerializationTestCase) makePropertyMap() map[PropertyName]*PropertyDefinition {
	result := make(map[PropertyName]*PropertyDefinition)
	for _, prop := range o.objectType.Properties() {
		result[prop.PropertyName()] = prop
	}
	return result
}

func (o ObjectSerializationTestCase) idOfSubjectGeneratorGlobal() string {
	return o.idOfGeneratorGlobal(o.subject)
}

func (o ObjectSerializationTestCase) idOfTestMethod() string {
	return o.idFactory.CreateIdentifier(
		fmt.Sprintf("RunTestFor%v", o.Subject()),
		Exported)
}

func (o ObjectSerializationTestCase) idOfGeneratorGlobal(name TypeName) string {
	return o.idFactory.CreateIdentifier(
		fmt.Sprintf("cached%vGenerator", name.Name()),
		NotExported)
}

func (o ObjectSerializationTestCase) idOfGeneratorMethod(typeName TypeName) string {
	name := o.idFactory.CreateIdentifier(
		fmt.Sprintf("%vGenerator", typeName.Name()),
		Exported)
	return name
}

func (o ObjectSerializationTestCase) idOfIndependentGeneratorsFactoryMethod() string {
	return o.idFactory.CreateIdentifier(
		fmt.Sprintf("AddIndependentPropertyGeneratorsFor%v", o.Subject()),
		Exported)
}

// idOfRelatedTypesGeneratorsFactoryMethod creates the identifier for the method that creates generators referencing
// other types
func (o ObjectSerializationTestCase) idOfRelatedGeneratorsFactoryMethod() string {
	return o.idFactory.CreateIdentifier(
		fmt.Sprintf("AddRelatedPropertyGeneratorsFor%v", o.Subject()),
		Exported)
}

func (o ObjectSerializationTestCase) Subject() *ast.Ident {
	return ast.NewIdent(o.subject.name)
}

func (o ObjectSerializationTestCase) createEnumGenerator(enumName string, genPackageName string, enum *EnumType) (ast.Expr, error) {
	var values []ast.Expr
	for _, o := range enum.Options() {
		id := GetEnumValueId(enumName, o)
		values = append(values, ast.NewIdent(id))
	}

	return astbuilder.CallQualifiedFunc(genPackageName, "OneConstOf", values...), nil
}
