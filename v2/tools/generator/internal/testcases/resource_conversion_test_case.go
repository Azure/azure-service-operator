/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcases

import (
	"fmt"
	"go/token"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// ResourceConversionTestCase represents a test that verifies we can convert from one resource to the hub resource
// (based on our conversion graph), and back again, with no loss of data (lossless conversion via the PropertyBag).
// This will be a multistep conversion, building on the PropertyAssignment functions.
type ResourceConversionTestCase struct {
	testName  string                                // The name of this particular test
	subject   astmodel.TypeName                     // The subject this test is going to exercise
	toFn      *functions.ResourceConversionFunction // The function to convert TO our hub instance
	fromFn    *functions.ResourceConversionFunction // The function to convert FROM our hub instance
	idFactory astmodel.IdentifierFactory            // a reference to our common factory for creating identifiers
}

var _ astmodel.TestCase = &ResourceConversionTestCase{}

// NewResourceConversionTestCase creates a new test case for the specified resource
func NewResourceConversionTestCase(
	name astmodel.TypeName,
	resourceType *astmodel.ResourceType,
	idFactory astmodel.IdentifierFactory,
) (*ResourceConversionTestCase, error) {
	result := &ResourceConversionTestCase{
		subject:   name,
		idFactory: idFactory,
	}

	conversionImplementation, ok := resourceType.FindInterface(astmodel.ConvertibleInterface)
	if !ok {
		return nil, eris.Errorf("expected %s to implement conversions.Convertible including ConvertTo() and ConvertFrom()", name)
	}

	// Find ConvertTo and ConvertFrom functions from the implementation
	for _, implementationFunction := range conversionImplementation.Functions() {
		if fn, ok := implementationFunction.(*functions.ResourceConversionFunction); ok {
			if fn.Direction() == conversions.ConvertFrom {
				result.fromFn = fn
			} else if fn.Direction() == conversions.ConvertTo {
				result.toFn = fn
			}
		}
	}

	// Fail fast if something goes wrong
	if result.fromFn == nil {
		return nil, eris.Errorf("expected to find function ConvertFrom() on %s", name)
	}

	if result.toFn == nil {
		return nil, eris.Errorf("expected to find function ConvertTo() on %s", name)
	}

	if !astmodel.TypeEquals(result.fromFn.Hub(), result.toFn.Hub()) {
		return nil, eris.Errorf(
			"expected ConvertFrom(%s) and ConvertTo(%s) on %s to have the same parameter type",
			result.fromFn.Hub(),
			result.toFn.Hub(),
			name)
	}

	result.testName = fmt.Sprintf(
		"%s_WhenConvertedToHub_RoundTripsWithoutLoss",
		name.Name())

	return result, nil
}

// Name returns the unique name of this test case
func (tc *ResourceConversionTestCase) Name() string {
	return tc.testName
}

// References returns the set of types to which this test case refers.
func (tc *ResourceConversionTestCase) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet(
		tc.subject,
		tc.toFn.Hub())
}

// RequiredImports returns a set of the package imports required by this test case
func (tc *ResourceConversionTestCase) RequiredImports() *astmodel.PackageImportSet {
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

	result.AddImportOfReference(tc.toFn.Hub().PackageReference())

	return result
}

// AsFuncs renders the current test case and any supporting methods as Go abstract syntax trees
// subject is the name of the type under test
// codeGenerationContext contains reference material to use when generating
func (tc *ResourceConversionTestCase) AsFuncs(
	receiver astmodel.TypeName,
	codeGenerationContext *astmodel.CodeGenerationContext,
) ([]dst.Decl, error) {
	testRunner := tc.createTestRunner(codeGenerationContext)
	testMethod, err := tc.createTestMethod(receiver, codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating test method for %s", tc.subject.Name())
	}

	return []dst.Decl{
		testRunner,
		testMethod,
	}, nil
}

// Equals determines if this TestCase is equal to another one
func (tc *ResourceConversionTestCase) Equals(other astmodel.TestCase, override astmodel.EqualityOverrides) bool {
	fn, ok := other.(*ResourceConversionTestCase)
	if !ok {
		return false
	}

	return tc.testName == fn.testName &&
		tc.subject.Equals(fn.subject, override) &&
		tc.toFn.Equals(fn.toFn, override) &&
		tc.fromFn.Equals(fn.fromFn, override)
}

// createTestRunner generates the AST for the test runner itself
//
// parameters := gopter.DefaultTestParameters()
// parameters.MaxSize = 10
// parameters.MinSuccessfulTests = 10
// properties := gopter.NewProperties(parameters)
// properties.Property("...", prop.ForAll(RunTestForX, XGenerator())
// properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
func (tc *ResourceConversionTestCase) createTestRunner(codegenContext *astmodel.CodeGenerationContext) dst.Decl {
	const (
		parametersLocal  = "parameters"
		propertiesLocal  = "properties"
		propertyMethod   = "Property"
		testingRunMethod = "TestingRun"
	)

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
		dst.NewIdent(parametersLocal),
		"MaxSize",
		token.ASSIGN,
		astbuilder.IntLiteral(10))

	// parameters.MinSuccessfulTests = 10
	configureMinSuccessfulTests := astbuilder.QualifiedAssignment(
		dst.NewIdent(parametersLocal),
		"MinSuccessfulTests",
		token.ASSIGN,
		astbuilder.IntLiteral(10))

	// properties := gopter.NewProperties(parameters)
	defineProperties := astbuilder.ShortDeclaration(
		propertiesLocal,
		astbuilder.CallQualifiedFunc(gopterPackage, "NewProperties", dst.NewIdent(parametersLocal)))

	// partial expression: description of the test
	testName := astbuilder.StringLiteralf("Round trip from %s to hub returns original", tc.subject.Name())
	testName.Decs.Before = dst.NewLine

	// partial expression: prop.ForAll(RunTestForX, XGenerator())
	propForAll := astbuilder.CallQualifiedFunc(
		propPackage,
		"ForAll",
		dst.NewIdent(tc.idOfTestMethod()),
		astbuilder.CallFunc(idOfGeneratorMethod(tc.subject, tc.idFactory)))
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
		tc.testName,
		declareParallel,
		defineParameters,
		configureMaxSize,
		configureMinSuccessfulTests,
		defineProperties,
		defineTestCase,
		runTests)

	return fn.DefineFunc()
}

// createTestMethod generates the AST for a method to run a single test of round trip conversion
//
// var hub OtherType
// err := subject.ConvertTo(&hub)
//
//	if err != nil {
//	    return err.Error()
//	}
//
// var result OurType
// err = result.ConvertFrom(&hub)
//
//	if err != nil {
//	    return err.Error()
//	}
//
// match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
//
//	if !match {
//	    result := diff.Diff(subject, actual);
//	    return result
//	}
//
// return ""
func (tc *ResourceConversionTestCase) createTestMethod(
	subject astmodel.TypeName,
	codegenContext *astmodel.CodeGenerationContext,
) (dst.Decl, error) {
	const (
		errID        = "err"
		hubID        = "hub"
		actualID     = "actual"
		actualFmtID  = "actualFmt"
		matchID      = "match"
		subjectID    = "subject"
		subjectFmtID = "subjectFmt"
		copiedID     = "copied"
		resultID     = "result"
	)

	cmpPackage := codegenContext.MustGetImportedPackageName(astmodel.CmpReference)
	cmpoptsPackage := codegenContext.MustGetImportedPackageName(astmodel.CmpOptsReference)
	prettyPackage := codegenContext.MustGetImportedPackageName(astmodel.PrettyReference)
	diffPackage := codegenContext.MustGetImportedPackageName(astmodel.DiffReference)

	// copied := subject.DeepCopy()
	assignCopied := astbuilder.ShortDeclaration(
		copiedID,
		astbuilder.CallQualifiedFunc(subjectID, "DeepCopy"))
	assignCopied.Decorations().Before = dst.NewLine
	astbuilder.AddComment(&assignCopied.Decorations().Start, "// Copy subject to make sure conversion doesn't modify it")

	// var hub OtherType
	hubExpr, err := tc.toFn.Hub().AsTypeExpr(codegenContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", tc.toFn.Hub())
	}

	declareOther := astbuilder.LocalVariableDeclaration(
		hubID,
		hubExpr,
		"// Convert to our hub version")
	declareOther.Decorations().Before = dst.EmptyLine

	// err := subject.ConvertTo(&hub)
	assignTo := astbuilder.ShortDeclaration(
		errID,
		astbuilder.CallQualifiedFunc(
			copiedID,
			tc.toFn.Name(),
			astbuilder.AddrOf(dst.NewIdent(hubID))))

	// if err != nil { return err.Error() }
	assignToFailed := astbuilder.ReturnIfNotNil(
		dst.NewIdent(errID),
		astbuilder.CallQualifiedFunc("err", "Error"))

	// var result OurType
	subjectExpr, err := subject.AsTypeExpr(codegenContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", subject)
	}

	declareResult := astbuilder.LocalVariableDeclaration(
		actualID,
		subjectExpr,
		"// Convert from our hub version")
	declareResult.Decorations().Before = dst.EmptyLine

	// err = result.ConvertFrom(&hub)
	assignFrom := astbuilder.SimpleAssignment(
		dst.NewIdent(errID),
		astbuilder.CallQualifiedFunc(
			actualID,
			tc.fromFn.Name(),
			astbuilder.AddrOf(dst.NewIdent(hubID))))

	// if err != nil { return err.Error() }
	assignFromFailed := astbuilder.ReturnIfNotNil(
		dst.NewIdent(errID),
		astbuilder.CallQualifiedFunc("err", "Error"))

	// match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	equateEmpty := astbuilder.CallQualifiedFunc(cmpoptsPackage, "EquateEmpty")
	compare := astbuilder.ShortDeclaration(
		matchID,
		astbuilder.CallQualifiedFunc(cmpPackage, "Equal",
			dst.NewIdent(subjectID),
			dst.NewIdent(actualID),
			equateEmpty))
	compare.Decorations().Before = dst.EmptyLine
	astbuilder.AddComment(&compare.Decorations().Start, "// Compare actual with what we started with")

	// actualFmt := pretty.Sprint(actual)
	declareActual := astbuilder.ShortDeclaration(
		actualFmtID,
		astbuilder.CallQualifiedFunc(prettyPackage, "Sprint", dst.NewIdent(actualID)))

	// subjectFmt := pretty.Sprint(subject)
	declareSubject := astbuilder.ShortDeclaration(
		subjectFmtID,
		astbuilder.CallQualifiedFunc(prettyPackage, "Sprint", dst.NewIdent(subjectID)))

	// result := diff.Diff(subject, actual)
	declareDiff := astbuilder.ShortDeclaration(
		resultID,
		astbuilder.CallQualifiedFunc(diffPackage, "Diff", dst.NewIdent(subjectFmtID), dst.NewIdent(actualFmtID)))

	// return result
	returnDiff := astbuilder.Returns(dst.NewIdent(resultID))

	// if !match {
	//     result := diff.Diff(subject, actual);
	//     return result
	// }
	prettyPrint := astbuilder.SimpleIf(
		astbuilder.NotExpr(dst.NewIdent(matchID)),
		declareActual,
		declareSubject,
		declareDiff,
		returnDiff)

	// return ""
	ret := astbuilder.Returns(astbuilder.StringLiteral(""))
	ret.Decorations().Before = dst.EmptyLine

	// Create the function
	fn := &astbuilder.FuncDetails{
		Name: tc.idOfTestMethod(),
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

	subjectExpr, err = tc.subject.AsTypeExpr(codegenContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", tc.subject)
	}

	fn.AddParameter("subject", subjectExpr)
	fn.AddComments(fmt.Sprintf(
		"tests if a specific instance of %s round trips to the hub storage version and back losslessly",
		tc.subject.Name()))
	fn.AddReturns("string")

	return fn.DefineFunc(), nil
}

func (tc *ResourceConversionTestCase) idOfTestMethod() string {
	return tc.idFactory.CreateIdentifier(
		fmt.Sprintf("RunResourceConversionTestFor%s", tc.subject.Name()),
		astmodel.Exported)
}
