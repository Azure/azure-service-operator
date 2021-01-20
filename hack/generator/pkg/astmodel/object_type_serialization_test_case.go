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
	"github.com/pkg/errors"
	"k8s.io/klog/v2"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
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

// NewObjectSerializationTestCase creates a new test case for the JSON serialization round-trip-ability of the specified object type
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
func (o ObjectSerializationTestCase) AsFuncs(name TypeName, genContext *CodeGenerationContext) []dst.Decl {

	var errs []error
	properties := o.makePropertyMap()

	// Find all the simple generators (those with no external dependencies)
	simpleGenerators := o.createGenerators(properties, genContext, o.createIndependentGenerator)

	// Find all the complex generators (dependent on other generators we'll be generating elsewhere)
	relatedGenerators := o.createGenerators(properties, genContext, o.createRelatedGenerator)

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

	var result []dst.Decl

	if len(simpleGenerators) == 0 && len(relatedGenerators) == 0 {
		// No properties that we can generate to test - skip the testing completely
		errs = append(errs, errors.Errorf("No property generators for %v", name))
	} else {
		result = append(result,
			o.createTestRunner(genContext),
			o.createTestMethod(genContext),
			o.createGeneratorDeclaration(genContext),
			o.createGeneratorMethod(genContext, len(simpleGenerators) > 0, len(relatedGenerators) > 0))

		if len(simpleGenerators) > 0 {
			result = append(result, o.createGeneratorsFactoryMethod(o.idOfIndependentGeneratorsFactoryMethod(), simpleGenerators, genContext))
		}

		if len(relatedGenerators) > 0 {
			result = append(result, o.createGeneratorsFactoryMethod(o.idOfRelatedGeneratorsFactoryMethod(), relatedGenerators, genContext))
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

	// Merge references required for properties
	for _, prop := range o.objectType.Properties() {
		for _, ref := range prop.PropertyType().RequiredPackageReferences().AsSlice() {
			result.AddImportOfReference(ref)
		}
	}

	// We're not currently creating generators for types in this package, so leave it out
	result.Remove(NewPackageImport(GenRuntimeReference))

	return result
}

// Equals determines if this TestCase is equal to another one
func (o ObjectSerializationTestCase) Equals(_ TestCase) bool {
	panic("implement me")
}

// createTestRunner generates the AST for the test runner itself
func (o ObjectSerializationTestCase) createTestRunner(codegenContext *CodeGenerationContext) dst.Decl {

	const (
		parametersLocal  = "parameters"
		propertiesLocal  = "properties"
		propertyMethod   = "Property"
		testingRunMethod = "TestingRun"
	)

	gopterPackage := codegenContext.MustGetImportedPackageName(GopterReference)
	propPackage := codegenContext.MustGetImportedPackageName(GopterPropReference)
	testingPackage := codegenContext.MustGetImportedPackageName(TestingReference)

	t := dst.NewIdent("t")

	// parameters := gopter.DefaultTestParameters()
	defineParameters := astbuilder.SimpleAssignment(
		dst.NewIdent(parametersLocal),
		token.DEFINE,
		astbuilder.CallQualifiedFunc(gopterPackage, "DefaultTestParameters"))

	// parameters.MaxSize = 10
	configureMaxSize := astbuilder.SimpleAssignment(
		&dst.SelectorExpr{
			X:   dst.NewIdent(parametersLocal),
			Sel: dst.NewIdent("MaxSize"),
		},
		token.ASSIGN,
		astbuilder.IntLiteral(10))

	// properties := gopter.NewProperties(parameters)
	defineProperties := astbuilder.SimpleAssignment(
		dst.NewIdent(propertiesLocal),
		token.DEFINE,
		astbuilder.CallQualifiedFunc(gopterPackage, "NewProperties", dst.NewIdent(parametersLocal)))

	// partial expression: description of the test
	testName := astbuilder.StringLiteralf("Round trip of %v via JSON returns original", o.Subject())

	// partial expression: prop.ForAll(RunTestForX, XGenerator())
	propForAll := astbuilder.CallQualifiedFunc(
		propPackage,
		"ForAll",
		dst.NewIdent(o.idOfTestMethod()),
		astbuilder.CallFunc(o.idOfGeneratorMethod(o.subject)))

	// properties.Property("...", prop.ForAll(RunTestForX, XGenerator())
	defineTestCase := astbuilder.InvokeQualifiedFunc(
		propertiesLocal,
		propertyMethod,
		testName,
		propForAll)

	// properties.TestingRun(t)
	runTests := astbuilder.InvokeQualifiedFunc(propertiesLocal, testingRunMethod, t)

	// Define our function
	fn := astbuilder.NewTestFuncDetails(
		testingPackage,
		o.testName,
		defineParameters,
		configureMaxSize,
		defineProperties,
		defineTestCase,
		runTests)

	return fn.DefineFunc()
}

// createTestMethod generates the AST for a method to run a single test of JSON serialization
func (o ObjectSerializationTestCase) createTestMethod(codegenContext *CodeGenerationContext) dst.Decl {
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

	jsonPackage := codegenContext.MustGetImportedPackageName(JsonReference)
	cmpPackage := codegenContext.MustGetImportedPackageName(CmpReference)
	cmpoptsPackage := codegenContext.MustGetImportedPackageName(CmpOptsReference)
	prettyPackage := codegenContext.MustGetImportedPackageName(PrettyReference)
	diffPackage := codegenContext.MustGetImportedPackageName(DiffReference)

	// bin, err := json.Marshal(subject)
	serialize := astbuilder.SimpleAssignmentWithErr(
		dst.NewIdent(binId),
		token.DEFINE,
		astbuilder.CallQualifiedFunc(jsonPackage, "Marshal", dst.NewIdent(subjectId)))

	// if err != nil { return err.Error() }
	serializeFailed := astbuilder.ReturnIfNotNil(
		dst.NewIdent(errId),
		astbuilder.CallQualifiedFunc("err", "Error"))

	// var actual X
	declare := astbuilder.NewVariable(actualId, o.subject.name)

	// err = json.Unmarshal(bin, &actual)
	deserialize := astbuilder.SimpleAssignment(
		dst.NewIdent("err"),
		token.ASSIGN,
		astbuilder.CallQualifiedFunc(jsonPackage, "Unmarshal",
			dst.NewIdent(binId),
			&dst.UnaryExpr{
				Op: token.AND,
				X:  dst.NewIdent(actualId),
			}))

	// if err != nil { return err.Error() }
	deserializeFailed := astbuilder.ReturnIfNotNil(
		dst.NewIdent(errId),
		astbuilder.CallQualifiedFunc("err", "Error"))

	// match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	equateEmpty := astbuilder.CallQualifiedFunc(cmpoptsPackage, "EquateEmpty")
	compare := astbuilder.SimpleAssignment(
		dst.NewIdent(matchId),
		token.DEFINE,
		astbuilder.CallQualifiedFunc(cmpPackage, "Equal",
			dst.NewIdent(subjectId),
			dst.NewIdent(actualId),
			equateEmpty))

	// if !match { result := diff.Diff(subject, actual); return result }
	prettyPrint := &dst.IfStmt{
		Cond: &dst.UnaryExpr{
			Op: token.NOT,
			X:  dst.NewIdent(matchId),
		},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				astbuilder.SimpleAssignment(
					dst.NewIdent(actualFmtId),
					token.DEFINE,
					astbuilder.CallQualifiedFunc(prettyPackage, "Sprint", dst.NewIdent(actualId))),
				astbuilder.SimpleAssignment(
					dst.NewIdent(subjectFmtId),
					token.DEFINE,
					astbuilder.CallQualifiedFunc(prettyPackage, "Sprint", dst.NewIdent(subjectId))),
				astbuilder.SimpleAssignment(
					dst.NewIdent(resultId),
					token.DEFINE,
					astbuilder.CallQualifiedFunc(diffPackage, "Diff", dst.NewIdent(subjectFmtId), dst.NewIdent(actualFmtId))),
				astbuilder.Returns(dst.NewIdent(resultId)),
			},
		},
	}

	// return ""
	ret := astbuilder.Returns(astbuilder.StringLiteral(""))

	// Create the function
	fn := &astbuilder.FuncDetails{
		Name: o.idOfTestMethod(),
		Returns: []*dst.Field{
			{
				Type: dst.NewIdent("string"),
			},
		},
		Body: []dst.Stmt{
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

func (o ObjectSerializationTestCase) createGeneratorDeclaration(genContext *CodeGenerationContext) dst.Decl {
	comment := fmt.Sprintf(
		"Generator of %v instances for property testing - lazily instantiated by %v()",
		o.Subject(),
		o.idOfGeneratorMethod(o.subject))

	gopterPackage := genContext.MustGetImportedPackageName(GopterReference)

	decl := astbuilder.VariableDeclaration(
		o.idOfSubjectGeneratorGlobal(),
		astbuilder.QualifiedTypeName(gopterPackage, "Gen"),
		comment)

	return decl
}

// createGeneratorMethod generates the AST for a method used to populate our generator cache variable on demand
func (o ObjectSerializationTestCase) createGeneratorMethod(ctx *CodeGenerationContext, haveSimpleGenerators bool, haveRelatedGenerators bool) dst.Decl {

	gopterPackage := ctx.MustGetImportedPackageName(GopterReference)
	genPackage := ctx.MustGetImportedPackageName(GopterGenReference)

	fn := &astbuilder.FuncDetails{
		Name: o.idOfGeneratorMethod(o.subject),
		Returns: []*dst.Field{
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
		dst.NewIdent(o.idOfSubjectGeneratorGlobal()),
		dst.NewIdent(o.idOfSubjectGeneratorGlobal()))
	fn.AddStatements(earlyReturn)

	if haveSimpleGenerators {
		// Create a simple version of the generator that does not reference generators for related types
		// This serves to terminate any dependency cycles that might occur during creation of a more fully fledged generator

		makeIndependentMap := astbuilder.SimpleAssignment(
			dst.NewIdent("independentGenerators"),
			token.DEFINE,
			astbuilder.MakeMap(
				dst.NewIdent("string"),
				astbuilder.QualifiedTypeName(gopterPackage, "Gen")))

		addIndependentGenerators := astbuilder.InvokeFunc(
			o.idOfIndependentGeneratorsFactoryMethod(),
			dst.NewIdent("independentGenerators"))

		createIndependentGenerator := astbuilder.SimpleAssignment(
			dst.NewIdent(o.idOfSubjectGeneratorGlobal()),
			token.ASSIGN,
			astbuilder.CallQualifiedFunc(
				genPackage,
				"Struct",
				astbuilder.CallQualifiedFunc("reflect", "TypeOf", &dst.CompositeLit{Type: o.Subject()}),
				dst.NewIdent("independentGenerators")))

		fn.AddStatements(makeIndependentMap, addIndependentGenerators, createIndependentGenerator)
	}

	if haveRelatedGenerators {
		// Define a local that contains all the simple generators
		// Have to call the factory method twice as the simple generator above has captured the map;
		// if we reuse or modify the map, chaos ensues.

		makeAllMap := astbuilder.SimpleAssignment(
			dst.NewIdent("allGenerators"),
			token.DEFINE,
			astbuilder.MakeMap(
				dst.NewIdent("string"),
				astbuilder.QualifiedTypeName(gopterPackage, "Gen")))

		fn.AddStatements(makeAllMap)

		if haveSimpleGenerators {
			addIndependentGenerators := astbuilder.InvokeFunc(
				o.idOfIndependentGeneratorsFactoryMethod(),
				dst.NewIdent("allGenerators"))
			fn.AddStatements(addIndependentGenerators)
		}

		addRelatedGenerators := astbuilder.InvokeFunc(
			o.idOfRelatedGeneratorsFactoryMethod(),
			dst.NewIdent("allGenerators"))

		createFullGenerator := astbuilder.SimpleAssignment(
			dst.NewIdent(o.idOfSubjectGeneratorGlobal()),
			token.ASSIGN,
			astbuilder.CallQualifiedFunc(
				genPackage,
				"Struct",
				astbuilder.CallQualifiedFunc("reflect", "TypeOf", &dst.CompositeLit{Type: o.Subject()}),
				dst.NewIdent("allGenerators")))

		fn.AddStatements(addRelatedGenerators, createFullGenerator)
	}

	// Return the freshly created (and now cached) generator
	normalReturn := astbuilder.Returns(dst.NewIdent(o.idOfSubjectGeneratorGlobal()))
	fn.AddStatements(normalReturn)

	return fn.DefineFunc()
}

// createGeneratorsFactoryMethod generates the AST for a method creating gopter generators
func (o ObjectSerializationTestCase) createGeneratorsFactoryMethod(
	methodName string, generators []dst.Stmt, ctx *CodeGenerationContext) dst.Decl {

	gopterPackage := ctx.MustGetImportedPackageName(GopterReference)

	mapType := &dst.MapType{
		Key:   dst.NewIdent("string"),
		Value: astbuilder.QualifiedTypeName(gopterPackage, "Gen"),
	}

	fn := &astbuilder.FuncDetails{
		Name: methodName,
		Body: generators,
	}

	fn.AddComments("is a factory method for creating gopter generators")
	fn.AddParameter("gens", mapType)

	return fn.DefineFunc()
}

// createGenerators creates AST fragments for gopter generators to create values for properties
// properties is a map of properties needing generators; properties handled here are removed from the map.
// genPackageName is the name for the gopter/gen package (not hard coded in case it's renamed for conflict resolution)
// factory is a method for creating generators
func (o ObjectSerializationTestCase) createGenerators(
	properties map[PropertyName]*PropertyDefinition,
	genContext *CodeGenerationContext,
	factory func(name string, propertyType Type, genContext *CodeGenerationContext) dst.Expr) []dst.Stmt {

	gensIdent := dst.NewIdent("gens")

	var handled []PropertyName
	var result []dst.Stmt

	// Sort Properties into alphabetical order to ensure we always generate the same code
	var toGenerate []PropertyName
	for name := range properties {
		toGenerate = append(toGenerate, name)
	}
	sort.Slice(toGenerate, func(i, j int) bool {
		return toGenerate[i] < toGenerate[j]
	})

	// Iterate over all properties, creating generators where possible
	for _, name := range toGenerate {
		prop := properties[name]
		g := factory(string(name), prop.PropertyType(), genContext)
		if g != nil {
			insert := astbuilder.InsertMap(
				gensIdent,
				&dst.BasicLit{
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

	return result
}

// is directly supported by a Gopter generator, returning nil if the property type isn't supported.
func (o ObjectSerializationTestCase) createIndependentGenerator(
	name string,
	propertyType Type,
	genContext *CodeGenerationContext) dst.Expr {

	genPackage := genContext.MustGetImportedPackageName(GopterGenReference)

	// Handle simple primitive properties
	switch propertyType {
	case StringType:
		return astbuilder.CallQualifiedFunc(genPackage, "AlphaString")
	case UInt32Type:
		return astbuilder.CallQualifiedFunc(genPackage, "UInt32")
	case IntType:
		return astbuilder.CallQualifiedFunc(genPackage, "Int")
	case FloatType:
		return astbuilder.CallQualifiedFunc(genPackage, "Float32")
	case BoolType:
		return astbuilder.CallQualifiedFunc(genPackage, "Bool")
	}

	switch t := propertyType.(type) {
	case TypeName:
		types := genContext.GetTypesInCurrentPackage()
		def, ok := types[t]
		if ok {
			return o.createIndependentGenerator(def.Name().name, def.theType, genContext)
		}
		return nil

	case *EnumType:
		return o.createEnumGenerator(name, genPackage, t)

	case *OptionalType:
		g := o.createIndependentGenerator(name, t.Element(), genContext)
		if g != nil {
			return astbuilder.CallQualifiedFunc(genPackage, "PtrOf", g)
		}

	case *ArrayType:
		g := o.createIndependentGenerator(name, t.Element(), genContext)
		if g != nil {
			return astbuilder.CallQualifiedFunc(genPackage, "SliceOf", g)
		}

	case *MapType:
		keyGen := o.createIndependentGenerator(name, t.KeyType(), genContext)
		valueGen := o.createIndependentGenerator(name, t.ValueType(), genContext)
		if keyGen != nil && valueGen != nil {
			return astbuilder.CallQualifiedFunc(genPackage, "MapOf", keyGen, valueGen)
		}

	case ValidatedType:
		// TODO: we should restrict the values of generated types
		//       but at the moment this is only used for serialization tests, so doesn't affect
		//       anything
		return o.createIndependentGenerator(name, t.ElementType(), genContext)
	}

	// Not a simple property we can handle here
	return nil
}

// defined within the current package, returning nil if the property type isn't supported.
func (o ObjectSerializationTestCase) createRelatedGenerator(
	name string,
	propertyType Type,
	genContext *CodeGenerationContext) dst.Expr {

	genPackageName := genContext.MustGetImportedPackageName(GopterGenReference)

	switch t := propertyType.(type) {
	case TypeName:
		_, ok := genContext.GetTypesInPackage(t.PackageReference)
		if ok {
			// This is a type we're defining, so we can create a generator for it
			if t.PackageReference.Equals(genContext.CurrentPackage()) {
				// create a generator for a property referencing a type in this package
				return astbuilder.CallFunc(o.idOfGeneratorMethod(t))
			}

			importName := genContext.MustGetImportedPackageName(t.PackageReference)
			return astbuilder.CallQualifiedFunc(importName, o.idOfGeneratorMethod(t))
		}

		//TODO: Should we invoke a generator for stuff from our runtime package?

		return nil

	case *OptionalType:
		g := o.createRelatedGenerator(name, t.Element(), genContext)
		if g != nil {
			return astbuilder.CallQualifiedFunc(genPackageName, "PtrOf", g)
		}

	case *ArrayType:
		g := o.createRelatedGenerator(name, t.Element(), genContext)
		if g != nil {
			return astbuilder.CallQualifiedFunc(genPackageName, "SliceOf", g)
		}

	case *MapType:
		// We only support primitive types as keys
		keyGen := o.createIndependentGenerator(name, t.KeyType(), genContext)
		valueGen := o.createRelatedGenerator(name, t.ValueType(), genContext)
		if keyGen != nil && valueGen != nil {
			return astbuilder.CallQualifiedFunc(genPackageName, "MapOf", keyGen, valueGen)
		}

	case ValidatedType:
		// TODO: we should restrict the values of generated types
		//       but at the moment this is only used for serialization tests, so doesn't affect
		//       anything
		return o.createRelatedGenerator(name, t.ElementType(), genContext)
	}

	// Not a property we can handle here
	return nil
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

func (o ObjectSerializationTestCase) Subject() *dst.Ident {
	return dst.NewIdent(o.subject.name)
}

func (o ObjectSerializationTestCase) createEnumGenerator(enumName string, genPackageName string, enum *EnumType) dst.Expr {
	var values []dst.Expr
	for _, o := range enum.Options() {
		id := GetEnumValueId(enumName, o)
		values = append(values, dst.NewIdent(id))
	}

	return astbuilder.CallQualifiedFunc(genPackageName, "OneConstOf", values...)
}
