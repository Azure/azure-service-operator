/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcases

import (
	"fmt"
	"go/token"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astbuilder"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/conversions"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/functions"
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

func NewResourceConversionTestCase(
	resource astmodel.TypeDefinition,
	idFactory astmodel.IdentifierFactory) *ResourceConversionTestCase {

	name := resource.Name()
	resourceType, ok := astmodel.AsResourceType(resource.Type())
	if !ok {
		panic(fmt.Sprintf("expected %s to be a resource type", name))
	}

	result := &ResourceConversionTestCase{
		subject:   name,
		idFactory: idFactory,
	}

	conversionImplementation, ok := resourceType.FindInterface(astmodel.ConvertibleInterface)
	if !ok {
		panic(fmt.Sprintf("expected %s to implement conversions.Convertible including ConvertTo() and ConvertFrom()", name))
	}

	// Find ConvertTo and ConvertFrom functions from the implementation
	for _, fn := range conversionImplementation.Functions() {
		if rcfn, ok := fn.(*functions.ResourceConversionFunction); ok {
			if rcfn.Direction() == conversions.ConvertFrom {
				result.fromFn = rcfn
			} else if rcfn.Direction() == conversions.ConvertTo {
				result.toFn = rcfn
			}
		}
	}

	// Fail fast if something goes wrong
	if result.fromFn == nil {
		panic(fmt.Sprintf("expected to find ConvertFrom() on %s", name))
	}
	if result.toFn == nil {
		panic(fmt.Sprintf("expected to find ConvertTo() on %s", name))
	}
	if !result.fromFn.Hub().Equals(result.toFn.Hub()) {

	}

	result.testName = fmt.Sprintf(
		"%s_WhenConvertedToHub_RoundTripsWithoutLoss",
		name.Name())

	return result
}

// Name returns the unique name of this test case
func (p *ResourceConversionTestCase) Name() string {
	return p.testName
}

// References returns the set of types to which this test case refers.
func (p *ResourceConversionTestCase) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet(
		p.subject,
		p.toFn.Hub())
}

// RequiredImports returns a set of the package imports required by this test case
func (p *ResourceConversionTestCase) RequiredImports() *astmodel.PackageImportSet {
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

	result.AddImportOfReference(p.toFn.Hub().PackageReference)

	return result
}

// AsFuncs renders the current test case and any supporting methods as Go abstract syntax trees
// subject is the name of the type under test
// codeGenerationContext contains reference material to use when generating
func (p *ResourceConversionTestCase) AsFuncs(receiver astmodel.TypeName, codeGenerationContext *astmodel.CodeGenerationContext) []dst.Decl {
	return []dst.Decl{
		p.createTestRunner(codeGenerationContext),
		p.createTestMethod(receiver, codeGenerationContext),
	}
}

// Equals determines if this TestCase is equal to another one
func (p *ResourceConversionTestCase) Equals(other astmodel.TestCase) bool {
	fn, ok := other.(*ResourceConversionTestCase)
	if !ok {
		return false
	}

	return p.testName == fn.testName &&
		p.subject.Equals(fn.subject) &&
		p.toFn.Equals(fn.toFn) &&
		p.fromFn.Equals(fn.fromFn)
}

// createTestRunner generates the AST for the test runner itself
func (p *ResourceConversionTestCase) createTestRunner(codegenContext *astmodel.CodeGenerationContext) dst.Decl {
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

	// properties := gopter.NewProperties(parameters)
	defineProperties := astbuilder.ShortDeclaration(
		propertiesLocal,
		astbuilder.CallQualifiedFunc(gopterPackage, "NewProperties", dst.NewIdent(parametersLocal)))

	// partial expression: description of the test
	testName := astbuilder.StringLiteralf("Round trip from %s to hub returns original", p.subject.Name())
	testName.Decs.Before = dst.NewLine

	// partial expression: prop.ForAll(RunTestForX, XGenerator())
	propForAll := astbuilder.CallQualifiedFunc(
		propPackage,
		"ForAll",
		dst.NewIdent(p.idOfTestMethod()),
		astbuilder.CallFunc(idOfGeneratorMethod(p.subject, p.idFactory)))
	propForAll.Decs.Before = dst.NewLine

	// properties.Property("...", prop.ForAll(RunTestForX, XGenerator())
	defineTestCase := astbuilder.InvokeQualifiedFunc(
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
	runTests := astbuilder.InvokeQualifiedFunc(propertiesLocal, testingRunMethod, t, createReporter)

	// Define our function
	fn := astbuilder.NewTestFuncDetails(
		testingPackage,
		p.testName,
		defineParameters,
		configureMaxSize,
		defineProperties,
		defineTestCase,
		runTests)

	return fn.DefineFunc()
}

// createTestMethod generates the AST for a method to run a single test of round trip conversion
func (p *ResourceConversionTestCase) createTestMethod(
	subject astmodel.TypeName,
	codegenContext *astmodel.CodeGenerationContext) dst.Decl {
	const (
		errId        = "err"
		hubId        = "hub"
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

	// var hub OtherType
	declareOther := astbuilder.LocalVariableDeclaration(
		hubId,
		p.toFn.Hub().AsType(codegenContext),
		"// Convert to our hub version")
	declareOther.Decorations().Before = dst.NewLine

	// err := subject.ConvertTo(&hub)
	assignTo := astbuilder.ShortDeclaration(
		errId,
		astbuilder.CallQualifiedFunc(
			subjectId,
			p.toFn.Name(),
			astbuilder.AddrOf(dst.NewIdent(hubId))))

	// if err != nil { return err.Error() }
	assignToFailed := astbuilder.ReturnIfNotNil(
		dst.NewIdent(errId),
		astbuilder.CallQualifiedFunc("err", "Error"))

	// var result OurType
	declareResult := astbuilder.LocalVariableDeclaration(
		actualId,
		subject.AsType(codegenContext),
		"// Convert from our hub version")
	declareResult.Decorations().Before = dst.EmptyLine

	// err = result.ConvertFrom(&hub)
	assignFrom := astbuilder.SimpleAssignment(
		dst.NewIdent(errId),
		astbuilder.CallQualifiedFunc(
			actualId,
			p.fromFn.Name(),
			astbuilder.AddrOf(dst.NewIdent(hubId))))

	// if err != nil { return err.Error() }
	assignFromFailed := astbuilder.ReturnIfNotNil(
		dst.NewIdent(errId),
		astbuilder.CallQualifiedFunc("err", "Error"))

	// match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	equateEmpty := astbuilder.CallQualifiedFunc(cmpoptsPackage, "EquateEmpty")
	compare := astbuilder.ShortDeclaration(
		matchId,
		astbuilder.CallQualifiedFunc(cmpPackage, "Equal",
			dst.NewIdent(subjectId),
			dst.NewIdent(actualId),
			equateEmpty))
	compare.Decorations().Before = dst.EmptyLine
	astbuilder.AddComment(&compare.Decorations().Start, "// Compare actual with what we started with")

	// if !match { result := diff.Diff(subject, actual); return result }
	prettyPrint := &dst.IfStmt{
		Cond: &dst.UnaryExpr{
			Op: token.NOT,
			X:  dst.NewIdent(matchId),
		},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				astbuilder.ShortDeclaration(
					actualFmtId,
					astbuilder.CallQualifiedFunc(prettyPackage, "Sprint", dst.NewIdent(actualId))),
				astbuilder.ShortDeclaration(
					subjectFmtId,
					astbuilder.CallQualifiedFunc(prettyPackage, "Sprint", dst.NewIdent(subjectId))),
				astbuilder.ShortDeclaration(
					resultId,
					astbuilder.CallQualifiedFunc(diffPackage, "Diff", dst.NewIdent(subjectFmtId), dst.NewIdent(actualFmtId))),
				astbuilder.Returns(dst.NewIdent(resultId)),
			},
		},
	}

	// return ""
	ret := astbuilder.Returns(astbuilder.StringLiteral(""))
	ret.Decorations().Before = dst.EmptyLine

	// Create the function
	fn := &astbuilder.FuncDetails{
		Name: p.idOfTestMethod(),
		Body: astbuilder.Statements(
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
		"runs a test to see if a specific instance of %s round trips to the hub storage version back losslessly",
		p.subject))
	fn.AddReturns("string")

	return fn.DefineFunc()
}

func (p *ResourceConversionTestCase) idOfTestMethod() string {
	return p.idFactory.CreateIdentifier(
		fmt.Sprintf("RunResourceConversionTestFor%s", p.subject.Name()),
		astmodel.Exported)
}
