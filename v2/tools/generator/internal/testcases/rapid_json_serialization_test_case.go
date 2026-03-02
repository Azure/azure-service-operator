/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcases

import (
	"fmt"
	"go/token"
	"sort"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// RapidJSONSerializationTestCase represents a test that the object can be losslessly serialized to
// JSON and back again, using the rapid property testing library.
type RapidJSONSerializationTestCase struct {
	testName  string
	subject   astmodel.TypeName
	container astmodel.PropertyContainer
	isOneOf   bool
	idFactory astmodel.IdentifierFactory
}

var _ astmodel.TestCase = &RapidJSONSerializationTestCase{}

// NewRapidJSONSerializationTestCase creates a new test case for the JSON serialization round-trip-ability
// of the specified object type, using the rapid property testing library.
func NewRapidJSONSerializationTestCase(
	name astmodel.TypeName,
	container astmodel.PropertyContainer,
	isOneOf bool,
	idFactory astmodel.IdentifierFactory,
) *RapidJSONSerializationTestCase {
	testName := fmt.Sprintf("%s_WhenSerializedToJson_DeserializesAsEqual", name.Name())
	return &RapidJSONSerializationTestCase{
		testName:  testName,
		subject:   name,
		isOneOf:   isOneOf,
		container: container,
		idFactory: idFactory,
	}
}

// Name returns the unique name of this test case
func (r *RapidJSONSerializationTestCase) Name() string {
	return r.testName
}

// References returns the set of types to which this test case refers.
func (r *RapidJSONSerializationTestCase) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet()
}

// RequiredImports returns a set of the package imports required by this test case
func (r *RapidJSONSerializationTestCase) RequiredImports() *astmodel.PackageImportSet {
	result := astmodel.NewPackageImportSet()

	// Standard Go Packages
	result.AddImportsOfReferences(
		astmodel.JSONReference, astmodel.TestingReference)

	// Cmp
	result.AddImportsOfReferences(astmodel.CmpReference, astmodel.CmpOptsReference)

	// Rapid
	result.AddImportOfReference(astmodel.RapidReference)

	// Other References
	result.AddImportOfReference(astmodel.DiffReference)
	result.AddImportOfReference(astmodel.PrettyReference)

	// Merge references required for properties
	r.container.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
		for ref := range prop.PropertyType().RequiredPackageReferences().All() {
			result.AddImportOfReference(ref)
		}
	})

	// We're not currently creating generators for types in this package, so leave it out
	result.Remove(astmodel.NewPackageImport(astmodel.GenRuntimeReference))

	return result
}

// AsFuncs renders the current test case and supporting methods as Go abstract syntax trees
func (r *RapidJSONSerializationTestCase) AsFuncs(
	_ astmodel.TypeName,
	genContext *astmodel.CodeGenerationContext,
) ([]dst.Decl, error) {
	properties := r.container.Properties().Copy()

	// Collect all generators (independent and related) for properties
	independentGens := r.collectGeneratorAssignments(properties, genContext, r.createIndependentGenerator)
	relatedGens := r.collectGeneratorAssignments(properties, genContext, r.createRelatedGenerator)

	// Remove properties from our runtime
	r.removeByPackage(properties, astmodel.GenRuntimeReference)
	r.removeByPackage(properties, astmodel.GenRuntimeConfigMapsReference)
	r.removeByPackage(properties, astmodel.GenRuntimeSecretsReference)
	r.removeByPackage(properties, astmodel.GenRuntimeCoreReference)

	// Remove API machinery properties
	r.removeByPackage(properties, astmodel.APIMachineryRuntimeReference)
	r.removeByPackage(properties, astmodel.APIMachinerySchemaReference)

	// Temporarily remove properties related to support for Arbitrary JSON
	r.removeByPackage(properties, astmodel.APIExtensionsReference)
	r.removeByPackage(properties, astmodel.APIExtensionsJSONReference)
	r.removeByPackage(properties, astmodel.GenRuntimeConditionsReference)

	// Write errors for any properties we don't handle
	errs := make([]error, 0, len(properties))
	for _, p := range properties {
		errs = append(errs, eris.Errorf("no generator created for %s (%s)", p.PropertyName(), p.PropertyType()))
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, err
	}

	// Merge independent and related generators
	allGens := make([]generatorAssignment, 0, len(independentGens)+len(relatedGens))
	allGens = append(allGens, independentGens...)
	allGens = append(allGens, relatedGens...)

	// Sort by property name for deterministic output
	sort.Slice(allGens, func(i, j int) bool {
		return allGens[i].propertyName < allGens[j].propertyName
	})

	result := []dst.Decl{
		r.createTestRunner(genContext),
		r.createTestMethod(genContext),
		r.createGeneratorDeclaration(genContext),
		r.createGeneratorMethod(genContext, allGens),
	}

	return result, nil
}

// Equals determines if this TestCase is equal to another one
func (r *RapidJSONSerializationTestCase) Equals(other astmodel.TestCase, overrides astmodel.EqualityOverrides) bool {
	otherTC, ok := other.(*RapidJSONSerializationTestCase)
	if !ok {
		return false
	}

	return r.testName == otherTC.testName &&
		r.subject.Equals(otherTC.subject, overrides) &&
		r.container == otherTC.container
}

// generatorAssignment holds a property name and its corresponding rapid generator expression,
// used for building the rapid.Custom() closure.
type generatorAssignment struct {
	propertyName string
	fieldName    string
	genExpr      dst.Expr
}

// createTestRunner generates the AST for the test runner itself
//
// func Test_X_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
//
//	t.Parallel()
//	rapid.Check(t, RunJSONSerializationTestForX)
//
// }
func (r *RapidJSONSerializationTestCase) createTestRunner(codegenContext *astmodel.CodeGenerationContext) dst.Decl {
	testingPackage := codegenContext.MustGetImportedPackageName(astmodel.TestingReference)
	rapidPackage := codegenContext.MustGetImportedPackageName(astmodel.RapidReference)

	t := dst.NewIdent("t")

	// t.Parallel()
	declareParallel := astbuilder.CallExprAsStmt(t, "Parallel")

	// rapid.Check(t, RunJSONSerializationTestForX)
	rapidCheck := astbuilder.CallQualifiedFuncAsStmt(
		rapidPackage,
		"Check",
		t,
		dst.NewIdent(r.idOfTestMethod()))

	fn := astbuilder.NewTestFuncDetails(
		testingPackage,
		r.testName,
		declareParallel,
		rapidCheck)

	return fn.DefineFunc()
}

// createTestMethod generates the AST for a method to run a single test of JSON serialization
//
// func RunJSONSerializationTestForX(t *rapid.T) {
//
//	subject := XGenerator().Draw(t, "subject")
//	bin, err := json.Marshal(subject)
//	if err != nil { t.Fatal(err) }
//	var actual X
//	err = json.Unmarshal(bin, &actual)
//	if err != nil { t.Fatal(err) }
//	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
//	if !match { ... t.Errorf(result) }
//
// }
func (r *RapidJSONSerializationTestCase) createTestMethod(codegenContext *astmodel.CodeGenerationContext) dst.Decl {
	const (
		binID        = "bin"
		actualID     = "actual"
		actualFmtID  = "actualFmt"
		matchID      = "match"
		subjectID    = "subject"
		subjectFmtID = "subjectFmt"
		resultID     = "result"
		errID        = "err"
	)

	rapidPackage := codegenContext.MustGetImportedPackageName(astmodel.RapidReference)
	jsonPackage := codegenContext.MustGetImportedPackageName(astmodel.JSONReference)
	cmpPackage := codegenContext.MustGetImportedPackageName(astmodel.CmpReference)
	cmpoptsPackage := codegenContext.MustGetImportedPackageName(astmodel.CmpOptsReference)
	prettyPackage := codegenContext.MustGetImportedPackageName(astmodel.PrettyReference)
	diffPackage := codegenContext.MustGetImportedPackageName(astmodel.DiffReference)

	// subject := XGenerator().Draw(t, "subject")
	drawSubject := astbuilder.ShortDeclaration(
		subjectID,
		astbuilder.CallExpr(
			astbuilder.CallFunc(idOfGeneratorMethod(r.subject, r.idFactory)),
			"Draw",
			dst.NewIdent("t"),
			astbuilder.StringLiteral("subject")))

	// bin, err := json.Marshal(subject)
	serialize := astbuilder.SimpleAssignmentWithErr(
		dst.NewIdent(binID),
		token.DEFINE,
		astbuilder.CallQualifiedFunc(jsonPackage, "Marshal", dst.NewIdent(subjectID)))
	astbuilder.AddComment(&serialize.Decs.Start, "// Serialize to JSON")
	serialize.Decorations().Before = dst.NewLine

	// if err != nil { t.Fatal(err) }
	serializeFailed := r.createFatalIfNotNil(errID)

	// var actual X
	declare := astbuilder.NewVariable(actualID, r.subject.Name())
	declare.Decorations().Before = dst.EmptyLine
	astbuilder.AddComment(&declare.Decorations().Start, "// Deserialize back into memory")

	// err = json.Unmarshal(bin, &actual)
	deserialize := astbuilder.SimpleAssignment(
		dst.NewIdent(errID),
		astbuilder.CallQualifiedFunc(jsonPackage, "Unmarshal",
			dst.NewIdent(binID),
			astbuilder.AddrOf(dst.NewIdent(actualID))))

	// if err != nil { t.Fatal(err) }
	deserializeFailed := r.createFatalIfNotNil(errID)

	// match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	equateEmpty := astbuilder.CallQualifiedFunc(cmpoptsPackage, "EquateEmpty")
	compare := astbuilder.ShortDeclaration(
		matchID,
		astbuilder.CallQualifiedFunc(cmpPackage, "Equal",
			dst.NewIdent(subjectID),
			dst.NewIdent(actualID),
			equateEmpty))
	compare.Decorations().Before = dst.EmptyLine
	astbuilder.AddComment(&compare.Decorations().Start, "// Check for outcome")

	// actualFmt := pretty.Sprint(actual)
	declareActual := astbuilder.ShortDeclaration(
		actualFmtID,
		astbuilder.CallQualifiedFunc(prettyPackage, "Sprint", dst.NewIdent(actualID)))

	// subjectFmt := pretty.Sprint(subject)
	declareSubject := astbuilder.ShortDeclaration(
		subjectFmtID,
		astbuilder.CallQualifiedFunc(prettyPackage, "Sprint", dst.NewIdent(subjectID)))

	// result := diff.Diff(subjectFmt, actualFmt)
	declareDiff := astbuilder.ShortDeclaration(
		resultID,
		astbuilder.CallQualifiedFunc(diffPackage, "Diff", dst.NewIdent(subjectFmtID), dst.NewIdent(actualFmtID)))

	// t.Error(result)
	reportError := astbuilder.CallExprAsStmt(dst.NewIdent("t"), "Error", dst.NewIdent(resultID))

	// if !match { ... t.Errorf(result) }
	prettyPrint := astbuilder.SimpleIf(
		astbuilder.NotExpr(dst.NewIdent(matchID)),
		declareActual,
		declareSubject,
		declareDiff,
		reportError)

	// Create the function
	fn := &astbuilder.FuncDetails{
		Name: r.idOfTestMethod(),
		Body: astbuilder.Statements(
			drawSubject,
			serialize,
			serializeFailed,
			declare,
			deserialize,
			deserializeFailed,
			compare,
			prettyPrint),
	}

	fn.AddParameter("t", astbuilder.Dereference(astbuilder.QualifiedTypeName(rapidPackage, "T")))
	fn.AddComments(fmt.Sprintf(
		"runs a test to see if a specific instance of %s round trips to JSON and back losslessly",
		r.Subject()))

	return fn.DefineFunc()
}

// createGeneratorDeclaration creates the global variable declaration for the generator cache
//
// var xGenerator *rapid.Generator[X]
func (r *RapidJSONSerializationTestCase) createGeneratorDeclaration(genContext *astmodel.CodeGenerationContext) dst.Decl {
	comment := fmt.Sprintf(
		"// Generator of %s instances for property testing - lazily instantiated by %s()",
		r.Subject(),
		idOfGeneratorMethod(r.subject, r.idFactory))

	rapidPackage := genContext.MustGetImportedPackageName(astmodel.RapidReference)

	// *rapid.Generator[X]
	generatorType := r.rapidGeneratorType(rapidPackage, r.Subject())

	decl := astbuilder.VariableDeclaration(
		r.idOfSubjectGeneratorGlobal(),
		generatorType,
		comment)

	return decl
}

// createGeneratorMethod generates the AST for the generator factory method
//
// func XGenerator() *rapid.Generator[X] {
//
//	if xGenerator != nil {
//	    return xGenerator
//	}
//	xGenerator = rapid.Custom(func(t *rapid.T) X {
//	    var result X
//	    result.Field1 = rapid.String().Draw(t, "Field1")
//	    ...
//	    return result
//	})
//	return xGenerator
//
// }
func (r *RapidJSONSerializationTestCase) createGeneratorMethod(
	ctx *astmodel.CodeGenerationContext,
	allGens []generatorAssignment,
) dst.Decl {
	if r.isOneOf {
		return r.createGeneratorMethodForOneOf(ctx, allGens)
	}

	return r.createGeneratorMethodForObject(ctx, allGens)
}

// createGeneratorMethodForObject generates the generator method for regular object types
func (r *RapidJSONSerializationTestCase) createGeneratorMethodForObject(
	ctx *astmodel.CodeGenerationContext,
	allGens []generatorAssignment,
) dst.Decl {
	rapidPkg := ctx.MustGetImportedPackageName(astmodel.RapidReference)

	generatorGlobalID := r.idOfSubjectGeneratorGlobal()

	// *rapid.Generator[X]
	returnType := r.rapidGeneratorType(rapidPkg, r.Subject())

	fn := &astbuilder.FuncDetails{
		Name: idOfGeneratorMethod(r.subject, r.idFactory),
		Returns: []*dst.Field{
			{
				Type: returnType,
			},
		},
	}

	fn.AddComments(fmt.Sprintf("returns a generator of %s instances for property testing.", r.Subject()))

	// if xGenerator != nil { return xGenerator }
	earlyReturn := astbuilder.ReturnIfNotNil(
		dst.NewIdent(generatorGlobalID),
		dst.NewIdent(generatorGlobalID))

	fn.AddStatements(earlyReturn)

	// Build the generator expression
	var generatorExpr dst.Expr
	if len(allGens) == 0 {
		// No generatable properties — use rapid.Just(X{}) to avoid
		// "Custom generator not calling any of the built-in generators" panic
		generatorExpr = astbuilder.CallQualifiedFunc(
			rapidPkg,
			"Just",
			&dst.CompositeLit{
				Type: r.Subject(),
			})
	} else {
		// Build the rapid.Custom closure body
		closureBody := r.buildCustomClosureBody(allGens)
		// rapid.Custom(func(t *rapid.T) X { ... })
		generatorExpr = r.buildRapidCustomCall(rapidPkg, closureBody)
	}

	// xGenerator = <generatorExpr>
	assignGenerator := astbuilder.SimpleAssignment(dst.NewIdent(generatorGlobalID), generatorExpr)
	assignGenerator.Decorations().Before = dst.EmptyLine

	// return xGenerator
	ret := astbuilder.Returns(dst.NewIdent(generatorGlobalID))
	ret.Decorations().Before = dst.EmptyLine

	fn.AddStatements(assignGenerator, ret)

	return fn.DefineFunc()
}

// createGeneratorMethodForOneOf generates the generator method for OneOf types using rapid.OneOf()
func (r *RapidJSONSerializationTestCase) createGeneratorMethodForOneOf(
	ctx *astmodel.CodeGenerationContext,
	allGens []generatorAssignment,
) dst.Decl {
	rapidPkg := ctx.MustGetImportedPackageName(astmodel.RapidReference)

	generatorGlobalID := r.idOfSubjectGeneratorGlobal()

	// *rapid.Generator[X]
	returnType := r.rapidGeneratorType(rapidPkg, r.Subject())

	fn := &astbuilder.FuncDetails{
		Name: idOfGeneratorMethod(r.subject, r.idFactory),
		Returns: []*dst.Field{
			{
				Type: returnType,
			},
		},
	}

	fn.AddComments(fmt.Sprintf("returns a generator of %s instances for property testing.", r.Subject()))

	// if xGenerator != nil { return xGenerator }
	earlyReturn := astbuilder.ReturnIfNotNil(
		dst.NewIdent(generatorGlobalID),
		dst.NewIdent(generatorGlobalID))

	fn.AddStatements(earlyReturn)

	gensName := "gens"

	// var gens []*rapid.Generator[X]
	gensType := &dst.ArrayType{
		Elt: r.rapidGeneratorType(rapidPkg, r.Subject()),
	}
	declareGens := astbuilder.LocalVariableDeclaration(gensName, gensType, "")
	declareGens.Decs.Before = dst.EmptyLine
	declareGens.Decs.Start = dst.Decorations{"// handle OneOf by choosing only one field to instantiate"}

	fn.AddStatements(declareGens)

	// For each generator, create a rapid.Custom() that sets just that one field
	for _, gen := range allGens {
		// rapid.Custom(func(t *rapid.T) X {
		//     var result X
		//     result.Field = <gen>.Draw(t, "Field")
		//     return result
		// })
		drawExpr := astbuilder.CallExpr(gen.genExpr, "Draw", dst.NewIdent("t"), astbuilder.StringLiteral(gen.fieldName))
		fieldAssign := astbuilder.QualifiedAssignment(
			dst.NewIdent("result"),
			gen.fieldName,
			token.ASSIGN,
			drawExpr)

		singleFieldBody := astbuilder.Statements(
			astbuilder.NewVariable("result", r.subject.Name()),
			fieldAssign,
			astbuilder.Returns(dst.NewIdent("result")))

		singleFieldCustomCall := r.buildRapidCustomCall(rapidPkg, singleFieldBody)

		appendStmt := astbuilder.AppendItemToSlice(dst.NewIdent(gensName), singleFieldCustomCall)
		fn.AddStatements(appendStmt)
	}

	// xGenerator = rapid.OneOf(gens...)
	oneOfCall := &dst.CallExpr{
		Fun:      astbuilder.QualifiedTypeName(rapidPkg, "OneOf"),
		Args:     astbuilder.Expressions(dst.NewIdent(gensName)),
		Ellipsis: true,
	}
	assignGenerator := astbuilder.SimpleAssignment(dst.NewIdent(generatorGlobalID), oneOfCall)

	// return xGenerator
	ret := astbuilder.Returns(dst.NewIdent(generatorGlobalID))
	ret.Decorations().Before = dst.EmptyLine

	fn.AddStatements(assignGenerator, ret)

	return fn.DefineFunc()
}

// buildCustomClosureBody creates the body of the rapid.Custom() closure
// var result X
// result.Field1 = rapid.String().Draw(t, "Field1")
// ...
// return result
func (r *RapidJSONSerializationTestCase) buildCustomClosureBody(allGens []generatorAssignment) []dst.Stmt {
	stmts := make([]dst.Stmt, 0, len(allGens)+2)

	// var result X
	declareResult := astbuilder.NewVariable("result", r.subject.Name())
	stmts = append(stmts, declareResult)

	// result.Field = <gen>.Draw(t, "Field")
	for _, gen := range allGens {
		drawExpr := astbuilder.CallExpr(gen.genExpr, "Draw", dst.NewIdent("t"), astbuilder.StringLiteral(gen.fieldName))
		assign := astbuilder.QualifiedAssignment(
			dst.NewIdent("result"),
			gen.fieldName,
			token.ASSIGN,
			drawExpr)
		stmts = append(stmts, assign)
	}

	// return result
	stmts = append(stmts, astbuilder.Returns(dst.NewIdent("result")))

	return stmts
}

// buildRapidCustomCall creates a rapid.Custom(func(t *rapid.T) X { ... }) expression
func (r *RapidJSONSerializationTestCase) buildRapidCustomCall(rapidPkg string, closureBody []dst.Stmt) *dst.CallExpr {
	return astbuilder.CallQualifiedFunc(
		rapidPkg,
		"Custom",
		&dst.FuncLit{
			Type: &dst.FuncType{
				Params: &dst.FieldList{
					List: []*dst.Field{
						{
							Names: []*dst.Ident{dst.NewIdent("t")},
							Type:  astbuilder.Dereference(astbuilder.QualifiedTypeName(rapidPkg, "T")),
						},
					},
				},
				Results: &dst.FieldList{
					List: []*dst.Field{
						{
							Type: r.Subject(),
						},
					},
				},
			},
			Body: &dst.BlockStmt{
				List: closureBody,
			},
		})
}

// rapidGeneratorType creates the AST for *rapid.Generator[T]
func (r *RapidJSONSerializationTestCase) rapidGeneratorType(rapidPkg string, subjectType dst.Expr) dst.Expr {
	return &dst.StarExpr{
		X: &dst.IndexExpr{
			X: astbuilder.QualifiedTypeName(rapidPkg, "Generator"),
			Index: &dst.Ident{
				Name: subjectType.(*dst.Ident).Name,
			},
		},
	}
}

// collectGeneratorAssignments collects generator assignments for properties.
// Properties handled here are removed from the map.
func (r *RapidJSONSerializationTestCase) collectGeneratorAssignments(
	properties map[astmodel.PropertyName]*astmodel.PropertyDefinition,
	genContext *astmodel.CodeGenerationContext,
	factory func(name string, propertyType astmodel.Type, genContext *astmodel.CodeGenerationContext) dst.Expr,
) []generatorAssignment {
	var handled []astmodel.PropertyName
	var result []generatorAssignment

	// Sort Properties into alphabetical order to ensure we always generate the same code
	toGenerate := make([]astmodel.PropertyName, 0, len(properties))
	for name := range properties {
		toGenerate = append(toGenerate, name)
	}
	sort.Slice(toGenerate, func(i, j int) bool {
		return toGenerate[i] < toGenerate[j]
	})

	for _, name := range toGenerate {
		prop := properties[name]
		g := factory(string(name), prop.PropertyType(), genContext)
		if g != nil {
			result = append(result, generatorAssignment{
				propertyName: string(name),
				fieldName:    string(prop.PropertyName()),
				genExpr:      g,
			})
			handled = append(handled, name)
		}
	}

	// Remove properties we've handled from the map
	for _, name := range handled {
		delete(properties, name)
	}

	return result
}

// createIndependentGenerator creates a rapid generator for a property whose type is directly supported.
func (r *RapidJSONSerializationTestCase) createIndependentGenerator(
	name string,
	propertyType astmodel.Type,
	genContext *astmodel.CodeGenerationContext,
) dst.Expr {
	rapidPackage := genContext.MustGetImportedPackageName(astmodel.RapidReference)

	// Handle simple primitive properties
	switch propertyType {
	case astmodel.StringType:
		return astbuilder.CallQualifiedFunc(rapidPackage, "String")
	case astmodel.UInt32Type:
		return astbuilder.CallQualifiedFunc(rapidPackage, "Uint32")
	case astmodel.IntType:
		return astbuilder.CallQualifiedFunc(rapidPackage, "Int")
	case astmodel.FloatType:
		return astbuilder.CallQualifiedFunc(rapidPackage, "Float64")
	case astmodel.BoolType:
		return astbuilder.CallQualifiedFunc(rapidPackage, "Bool")
	}

	switch t := propertyType.(type) {
	case astmodel.InternalTypeName:
		defs := genContext.GetDefinitionsInCurrentPackage()
		def, ok := defs[t]
		if ok {
			vt, isValidated := def.Type().(*astmodel.ValidatedType)

			g := r.createIndependentGenerator(def.Name().Name(), def.Type(), genContext)
			if !isValidated || g == nil {
				return g
			}

			// typename pointing to validated type needs to cast the result back to the typename.
			// Use rapid.Map(generator, castFn) — a standalone function, NOT a method.
			// generates: rapid.Map(generator, func(it <elementType>) <ResultType> { return <ResultType>(it) })
			genMap := astbuilder.CallQualifiedFunc(
				rapidPackage,
				"Map",
				g,
				&dst.FuncLit{
					Type: &dst.FuncType{
						Params: &dst.FieldList{List: []*dst.Field{
							{
								Names: []*dst.Ident{dst.NewIdent("it")},
								Type:  dst.NewIdent(vt.ElementType().String()),
							},
						}},
						Results: &dst.FieldList{List: []*dst.Field{{Type: dst.NewIdent(t.Name())}}},
					},
					Body: astbuilder.StatementBlock(astbuilder.Returns(astbuilder.CallFunc(t.Name(), dst.NewIdent("it")))),
				})

			return genMap
		}
		return nil

	case *astmodel.EnumType:
		return r.createEnumGenerator(name, rapidPackage, t)

	case *astmodel.OptionalType:
		g := r.createIndependentGenerator(name, t.Element(), genContext)
		if g != nil {
			// rapid.Ptr(g, true)
			return astbuilder.CallQualifiedFunc(rapidPackage, "Ptr", g, dst.NewIdent("true"))
		}

	case *astmodel.ArrayType:
		g := r.createIndependentGenerator(name, t.Element(), genContext)
		if g != nil {
			return astbuilder.CallQualifiedFunc(rapidPackage, "SliceOf", g)
		}

	case *astmodel.MapType:
		keyGen := r.createIndependentGenerator(name, t.KeyType(), genContext)
		valueGen := r.createIndependentGenerator(name, t.ValueType(), genContext)
		if keyGen != nil && valueGen != nil {
			return astbuilder.CallQualifiedFunc(rapidPackage, "MapOf", keyGen, valueGen)
		}

	case *astmodel.ValidatedType:
		return r.createIndependentGenerator(name, t.ElementType(), genContext)
	}

	// Not a simple property we can handle here
	return nil
}

// createRelatedGenerator creates a rapid generator for a property whose type is defined within the current package.
func (r *RapidJSONSerializationTestCase) createRelatedGenerator(
	name string,
	propertyType astmodel.Type,
	genContext *astmodel.CodeGenerationContext,
) dst.Expr {
	rapidPackage := genContext.MustGetImportedPackageName(astmodel.RapidReference)

	switch t := propertyType.(type) {
	case astmodel.InternalTypeName:
		_, ok := genContext.GetDefinitionsInPackage(t.InternalPackageReference())
		if ok {
			// This is a type we're defining, so we can create a generator for it
			if t.PackageReference().Equals(genContext.CurrentPackage()) {
				return astbuilder.CallFunc(idOfGeneratorMethod(t, r.idFactory))
			}

			importName := genContext.MustGetImportedPackageName(t.PackageReference())
			return astbuilder.CallQualifiedFunc(importName, idOfGeneratorMethod(t, r.idFactory))
		}

		return nil

	case *astmodel.OptionalType:
		g := r.createRelatedGenerator(name, t.Element(), genContext)
		if g != nil {
			if r.isOneOf {
				// For OneOf members, force non-nil pointer using rapid.Map()
				typeName, ok := t.Element().(astmodel.TypeName)
				if !ok {
					panic(fmt.Sprintf("expected OneOf to contain pointer to TypeName but had: %s", typeName.String()))
				}

				// generates: rapid.Map(g, func(it T) *T { return &it })
				genMap := astbuilder.CallQualifiedFunc(
					rapidPackage,
					"Map",
					g,
					&dst.FuncLit{
						Type: &dst.FuncType{
							Params:  &dst.FieldList{List: []*dst.Field{{Names: []*dst.Ident{dst.NewIdent("it")}, Type: dst.NewIdent(typeName.Name())}}},
							Results: &dst.FieldList{List: []*dst.Field{{Type: astbuilder.Dereference(dst.NewIdent(typeName.Name()))}}},
						},
						Body: astbuilder.StatementBlock(astbuilder.Returns(astbuilder.AddrOf(dst.NewIdent("it")))),
					})

				genMap.Decs.End = []string{"// generate one case for OneOf type"}

				return genMap
			}

			// otherwise generate a pointer to the type that may be nil
			return astbuilder.CallQualifiedFunc(rapidPackage, "Ptr", g, dst.NewIdent("true"))
		}

	case *astmodel.ArrayType:
		g := r.createRelatedGenerator(name, t.Element(), genContext)
		if g != nil {
			return astbuilder.CallQualifiedFunc(rapidPackage, "SliceOf", g)
		}

	case *astmodel.MapType:
		keyGen := r.createIndependentGenerator(name, t.KeyType(), genContext)
		valueGen := r.createRelatedGenerator(name, t.ValueType(), genContext)
		if keyGen != nil && valueGen != nil {
			return astbuilder.CallQualifiedFunc(rapidPackage, "MapOf", keyGen, valueGen)
		}

	case *astmodel.ValidatedType:
		return r.createRelatedGenerator(name, t.ElementType(), genContext)
	}

	// Not a property we can handle here
	return nil
}

func (r *RapidJSONSerializationTestCase) removeByPackage(
	properties map[astmodel.PropertyName]*astmodel.PropertyDefinition,
	ref astmodel.PackageReference,
) {
	var toRemove []astmodel.PropertyName
	for name, prop := range properties {
		propertyType := prop.PropertyType()
		refs := propertyType.RequiredPackageReferences()
		if refs.Contains(ref) {
			toRemove = append(toRemove, name)
		}
	}

	for _, name := range toRemove {
		delete(properties, name)
	}
}

func (r *RapidJSONSerializationTestCase) idOfSubjectGeneratorGlobal() string {
	return r.idFactory.CreateIdentifier(
		fmt.Sprintf("%sGenerator", r.subject.Name()),
		astmodel.NotExported)
}

func (r *RapidJSONSerializationTestCase) idOfTestMethod() string {
	return r.idFactory.CreateIdentifier(
		fmt.Sprintf("RunJSONSerializationTestFor%s", r.Subject()),
		astmodel.Exported)
}

func (r *RapidJSONSerializationTestCase) Subject() *dst.Ident {
	return dst.NewIdent(r.subject.Name())
}

func (r *RapidJSONSerializationTestCase) createEnumGenerator(enumName string, rapidPkg string, enum *astmodel.EnumType) dst.Expr {
	opts := enum.Options()
	values := make([]dst.Expr, 0, len(opts))
	for _, o := range opts {
		id := astmodel.GetEnumValueID(enumName, o)
		values = append(values, dst.NewIdent(id))
	}

	// rapid.SampledFrom([]T{v1, v2, ...})
	// Build the slice literal for the enum values
	sliceLit := &dst.CompositeLit{
		Type: &dst.ArrayType{
			Elt: dst.NewIdent(enumName),
		},
		Elts: values,
	}

	return astbuilder.CallQualifiedFunc(rapidPkg, "SampledFrom", sliceLit)
}

// createFatalIfNotNil generates:
//
//	if <id> != nil { t.Fatal(<id>) }
func (r *RapidJSONSerializationTestCase) createFatalIfNotNil(id string) *dst.IfStmt {
	return &dst.IfStmt{
		Cond: astbuilder.AreNotEqual(dst.NewIdent(id), dst.NewIdent("nil")),
		Body: astbuilder.StatementBlock(
			astbuilder.CallExprAsStmt(dst.NewIdent("t"), "Fatal", dst.NewIdent(id))),
	}
}
