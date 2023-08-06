/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcases

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// PropertyAssignmentTestCase represents a test that verifies we can convert from one object in our conversion graph
// to the next one, and back again, with no loss of data (lossless conversion via the PropertyBag)
type PropertyAssignmentTestCase struct {
	testName  string                                // The name of this particular test
	subject   astmodel.TypeName                     // The subject this test is going to exercise
	toFn      *functions.PropertyAssignmentFunction // The function to copy properties TO our related instance
	fromFn    *functions.PropertyAssignmentFunction // The function to copy properties FROM our related instance
	idFactory astmodel.IdentifierFactory            // a reference to our common factory for creating identifiers
}

var _ astmodel.TestCase = &PropertyAssignmentTestCase{}

func NewPropertyAssignmentTestCase(
	name astmodel.TypeName,
	container astmodel.FunctionContainer,
	idFactory astmodel.IdentifierFactory) *PropertyAssignmentTestCase {

	result := &PropertyAssignmentTestCase{
		subject:   name,
		idFactory: idFactory,
	}

	// Find Property Assignment functions
	for _, fn := range container.Functions() {
		if !strings.HasPrefix(fn.Name(), conversions.AssignPropertiesMethodPrefix) {
			// We're now using PropertyAssignment functions in other contexts, but only want to generate tests
			// for the originals (and those only because we're allowing hand-written extensions that need testing).
			continue
		}

		if pafn, ok := fn.(*functions.PropertyAssignmentFunction); ok {
			if pafn.Direction() == conversions.ConvertFrom {
				result.fromFn = pafn
			} else if pafn.Direction() == conversions.ConvertTo {
				result.toFn = pafn
			}
		}
	}

	// Fail fast if something goes wrong
	if result.fromFn == nil {
		panic(fmt.Sprintf("expected to find PropertyAssignmentFrom() on %s", name))
	}

	if result.toFn == nil {
		panic(fmt.Sprintf("expected to find PropertyAssignmentTo() on %s", name))
	}

	if !astmodel.TypeEquals(result.fromFn.ParameterType(), result.toFn.ParameterType()) {
		panic(fmt.Sprintf("expected PropertyAssignmentFrom() and PropertyAssignmentTo() on %s to be consistent", name))
	}

	result.testName = fmt.Sprintf(
		"%s_WhenPropertiesConverted_RoundTripsWithoutLoss",
		name.Name())

	return result
}

// Name returns the unique name of this test case
func (p *PropertyAssignmentTestCase) Name() string {
	return p.testName
}

// References returns the set of types to which this test case refers.
func (p *PropertyAssignmentTestCase) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet(
		p.subject,
		p.toFn.ParameterType())
}

// RequiredImports returns a set of the package imports required by this test case
func (p *PropertyAssignmentTestCase) RequiredImports() *astmodel.PackageImportSet {
	result := astmodel.NewPackageImportSet()

	// Standard Go Packages
	result.AddImportsOfReferences(astmodel.OSReference, astmodel.TestingReference)

	// Cmp
	result.AddImportsOfReferences(astmodel.CmpReference, astmodel.CmpOptsReference)

	// Gopter
	result.AddImportsOfReferences(astmodel.GopterReference, astmodel.GopterGenReference, astmodel.GopterPropReference)

	// Other References
	result.AddImportOfReference(astmodel.DiffReference)
	result.AddImportOfReference(astmodel.PrettyReference)

	result.AddImportOfReference(p.toFn.ParameterType().PackageReference())

	return result
}

// AsFuncs renders the current test case and any supporting methods as Go abstract syntax trees
// subject is the name of the type under test
// codeGenerationContext contains reference material to use when generating
func (p *PropertyAssignmentTestCase) AsFuncs(
	receiver astmodel.TypeName, codeGenerationContext *astmodel.CodeGenerationContext) []dst.Decl {
	return []dst.Decl{
		p.createTestRunner(codeGenerationContext),
		p.createTestMethod(receiver, codeGenerationContext),
	}
}

// Equals determines if this TestCase is equal to another one
func (p *PropertyAssignmentTestCase) Equals(other astmodel.TestCase, override astmodel.EqualityOverrides) bool {
	fn, ok := other.(*PropertyAssignmentTestCase)
	if !ok {
		return false
	}

	return p.testName == fn.testName &&
		p.subject.Equals(fn.subject, override) &&
		p.toFn.Equals(fn.toFn, override) &&
		p.fromFn.Equals(fn.fromFn, override)
}

// createTestRunner generates the AST for the test runner itself
func (p *PropertyAssignmentTestCase) createTestRunner(codegenContext *astmodel.CodeGenerationContext) dst.Decl {
	const (
		parametersLocal  = "parameters"
		propertiesLocal  = "properties"
		propertyMethod   = "Property"
		testingRunMethod = "TestingRun"
	)

	parametersLocalId := dst.NewIdent(parametersLocal)

	gopterPackage := codegenContext.MustGetImportedPackageName(astmodel.GopterReference)
	osPackage := codegenContext.MustGetImportedPackageName(astmodel.OSReference)
	propPackage := codegenContext.MustGetImportedPackageName(astmodel.GopterPropReference)
	testingPackage := codegenContext.MustGetImportedPackageName(astmodel.TestingReference)

	t := dst.NewIdent("t")

	// t.Parallel()
	declareParallel := astbuilder.CallExprAsStmt(t, "Parallel")

	// parameters := gopter.DefaultTestParameters()
	defineParameters := astbuilder.ShortDeclaration(
		parametersLocal,
		astbuilder.CallQualifiedFunc(gopterPackage, "DefaultTestParameters"))

	// parameters.MaxSize = 10
	configureMaxSize := astbuilder.QualifiedAssignment(
		parametersLocalId,
		"MaxSize",
		token.ASSIGN,
		astbuilder.IntLiteral(10))

	// properties := gopter.NewProperties(parameters)
	defineProperties := astbuilder.ShortDeclaration(
		propertiesLocal,
		astbuilder.CallQualifiedFunc(gopterPackage, "NewProperties", parametersLocalId))

	// partial expression: description of the test
	testName := astbuilder.StringLiteralf("Round trip from %s to %s via %s & %s returns original",
		p.subject.Name(),
		p.toFn.ParameterType().Name(),
		p.toFn.Name(),
		p.fromFn.Name())
	testName.Decs.Before = dst.NewLine

	// partial expression: prop.ForAll(RunTestForX, XGenerator())
	propForAll := astbuilder.CallQualifiedFunc(
		propPackage,
		"ForAll",
		dst.NewIdent(p.idOfTestMethod()),
		astbuilder.CallFunc(idOfGeneratorMethod(p.subject, p.idFactory)))
	propForAll.Decs.Before = dst.NewLine

	// properties.Property("...", prop.ForAll(RunTestForX, XGenerator())
	defineTestCase := astbuilder.CallQualifiedFuncAsStmt(
		propertiesLocal,
		propertyMethod,
		testName,
		propForAll)

	// properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
	createReporter := astbuilder.CallQualifiedFunc(
		gopterPackage,
		"NewFormatedReporter",
		dst.NewIdent("false"),
		astbuilder.IntLiteral(240),
		astbuilder.Selector(dst.NewIdent(osPackage), "Stdout"))
	runTests := astbuilder.CallQualifiedFuncAsStmt(propertiesLocal, testingRunMethod, t, createReporter)

	// Define our function
	fn := astbuilder.NewTestFuncDetails(
		testingPackage,
		p.testName,
		declareParallel,
		defineParameters,
		configureMaxSize,
		defineProperties,
		defineTestCase,
		runTests)

	return fn.DefineFunc()
}

// createTestMethod generates the AST for a method to run a single test conversion and back again
func (p *PropertyAssignmentTestCase) createTestMethod(
	subject astmodel.TypeName,
	codegenContext *astmodel.CodeGenerationContext) dst.Decl {
	const (
		errId        = "err"
		copiedId     = "copied"
		otherId      = "other"
		actualId     = "actual"
		actualFmtId  = "actualFmt"
		matchId      = "match"
		subjectId    = "subject"
		subjectFmtId = "subjectFmt"
		resultId     = "result"
	)

	cmpPackage := codegenContext.MustGetImportedPackageName(astmodel.CmpReference)
	cmpoptsPackage := codegenContext.MustGetImportedPackageName(astmodel.CmpOptsReference)
	prettyPackage := codegenContext.MustGetImportedPackageName(astmodel.PrettyReference)
	diffPackage := codegenContext.MustGetImportedPackageName(astmodel.DiffReference)

	// copied := subject.DeepCopy()
	assignCopied := astbuilder.ShortDeclaration(
		copiedId,
		astbuilder.CallQualifiedFunc(subjectId, "DeepCopy"))
	assignCopied.Decorations().Before = dst.NewLine
	astbuilder.AddComment(&assignCopied.Decorations().Start, "// Copy subject to make sure assignment doesn't modify it")

	// var other OtherType
	declareOther := astbuilder.LocalVariableDeclaration(
		otherId,
		p.toFn.ParameterType().AsType(codegenContext),
		"// Use AssignPropertiesTo() for the first stage of conversion")
	declareOther.Decorations().Before = dst.EmptyLine

	// err := subject.AssignPropertiesTo(other)
	assignTo := astbuilder.ShortDeclaration(
		errId,
		astbuilder.CallQualifiedFunc(
			copiedId,
			p.toFn.Name(),
			astbuilder.AddrOf(dst.NewIdent(otherId))))

	// if err != nil { return err.Error() }
	assignToFailed := astbuilder.ReturnIfNotNil(
		dst.NewIdent(errId),
		astbuilder.CallQualifiedFunc("err", "Error"))

	// var result OurType
	declareResult := astbuilder.LocalVariableDeclaration(
		actualId,
		subject.AsType(codegenContext),
		"// Use AssignPropertiesFrom() to convert back to our original type")
	declareResult.Decorations().Before = dst.EmptyLine

	// err = result.AssignPropertiesFrom(other)
	assignFrom := astbuilder.SimpleAssignment(
		dst.NewIdent(errId),
		astbuilder.CallQualifiedFunc(
			actualId,
			p.fromFn.Name(),
			astbuilder.AddrOf(dst.NewIdent(otherId))))

	// if err != nil { return err.Error() }
	assignFromFailed := astbuilder.ReturnIfNotNil(
		dst.NewIdent(errId),
		astbuilder.CallQualifiedFunc("err", "Error"))

	// match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	// We include cmpopts.EquateEmpty() to allow empty slices and maps to match nil values
	equateEmpty := astbuilder.CallQualifiedFunc(cmpoptsPackage, "EquateEmpty")
	compare := astbuilder.ShortDeclaration(
		matchId,
		astbuilder.CallQualifiedFunc(cmpPackage, "Equal",
			dst.NewIdent(subjectId),
			dst.NewIdent(actualId),
			equateEmpty))
	compare.Decorations().Before = dst.EmptyLine
	astbuilder.AddComment(&compare.Decorations().Start, "Check for a match")

	// actualFmt := pretty.Sprint(actual)
	declareActual := astbuilder.ShortDeclaration(
		actualFmtId,
		astbuilder.CallQualifiedFunc(prettyPackage, "Sprint", dst.NewIdent(actualId)))

	// subjectFmt := pretty.Sprint(subject)
	declareSubject := astbuilder.ShortDeclaration(
		subjectFmtId,
		astbuilder.CallQualifiedFunc(prettyPackage, "Sprint", dst.NewIdent(subjectId)))

	// result := diff.Diff(subject, actual)
	declareDiff := astbuilder.ShortDeclaration(
		resultId,
		astbuilder.CallQualifiedFunc(diffPackage, "Diff", dst.NewIdent(subjectFmtId), dst.NewIdent(actualFmtId)))

	// return result
	returnDiff := astbuilder.Returns(dst.NewIdent(resultId))

	// if !match {
	//     result := diff.Diff(subject, actual);
	//     return result
	// }
	prettyPrint := astbuilder.SimpleIf(
		astbuilder.NotExpr(dst.NewIdent(matchId)),
		declareActual,
		declareSubject,
		declareDiff,
		returnDiff)

	// return ""
	ret := astbuilder.Returns(astbuilder.StringLiteral(""))
	ret.Decorations().Before = dst.EmptyLine

	// Create the function
	fn := &astbuilder.FuncDetails{
		Name: p.idOfTestMethod(),
		Body: astbuilder.Statements(
			assignCopied,
			declareOther,
			assignTo,
			assignToFailed,
			declareResult,
			assignFrom,
			assignFromFailed,
			compare,
			prettyPrint,
			ret),
	}

	fn.AddParameter("subject", p.subject.AsType(codegenContext))
	fn.AddComments(fmt.Sprintf(
		"tests if a specific instance of %s can be assigned to %s and back losslessly",
		p.subject.Name(),
		p.fromFn.ParameterType().PackageReference().PackageName()))
	fn.AddReturns("string")

	return fn.DefineFunc()
}

func (p *PropertyAssignmentTestCase) idOfTestMethod() string {
	return p.idFactory.CreateIdentifier(
		fmt.Sprintf("RunPropertyAssignmentTestFor%s", p.subject.Name()),
		astmodel.Exported)
}
