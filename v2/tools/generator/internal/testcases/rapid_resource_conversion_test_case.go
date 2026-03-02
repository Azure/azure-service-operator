/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcases

import (
	"fmt"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// RapidResourceConversionTestCase represents a test that verifies we can convert from one resource to the hub resource
// and back again, with no loss of data, using the rapid property testing library.
type RapidResourceConversionTestCase struct {
	testName  string
	subject   astmodel.TypeName
	toFn      *functions.ResourceConversionFunction
	fromFn    *functions.ResourceConversionFunction
	idFactory astmodel.IdentifierFactory
}

var _ astmodel.TestCase = &RapidResourceConversionTestCase{}

// NewRapidResourceConversionTestCase creates a new test case for the specified resource
func NewRapidResourceConversionTestCase(
	name astmodel.TypeName,
	resourceType *astmodel.ResourceType,
	idFactory astmodel.IdentifierFactory,
) (*RapidResourceConversionTestCase, error) {
	result := &RapidResourceConversionTestCase{
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
func (r *RapidResourceConversionTestCase) Name() string {
	return r.testName
}

// References returns the set of types to which this test case refers.
func (r *RapidResourceConversionTestCase) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet(
		r.subject,
		r.toFn.Hub())
}

// RequiredImports returns a set of the package imports required by this test case
func (r *RapidResourceConversionTestCase) RequiredImports() *astmodel.PackageImportSet {
	result := astmodel.NewPackageImportSet()

	// Standard Go Packages
	result.AddImportsOfReferences(astmodel.TestingReference)

	// Cmp
	result.AddImportsOfReferences(astmodel.CmpReference, astmodel.CmpOptsReference)

	// Rapid
	result.AddImportOfReference(astmodel.RapidReference)

	// Other References
	result.AddImportOfReference(astmodel.DiffReference)
	result.AddImportOfReference(astmodel.PrettyReference)

	result.AddImportOfReference(r.toFn.Hub().PackageReference())

	return result
}

// AsFuncs renders the current test case and any supporting methods as Go abstract syntax trees
func (r *RapidResourceConversionTestCase) AsFuncs(
	receiver astmodel.TypeName,
	codeGenerationContext *astmodel.CodeGenerationContext,
) ([]dst.Decl, error) {
	testFunc, err := r.createTestFunc(receiver, codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating test function for %s", r.subject.Name())
	}

	return []dst.Decl{testFunc}, nil
}

// Equals determines if this TestCase is equal to another one
func (r *RapidResourceConversionTestCase) Equals(other astmodel.TestCase, override astmodel.EqualityOverrides) bool {
	fn, ok := other.(*RapidResourceConversionTestCase)
	if !ok {
		return false
	}

	return r.testName == fn.testName &&
		r.subject.Equals(fn.subject, override) &&
		r.toFn.Equals(fn.toFn, override) &&
		r.fromFn.Equals(fn.fromFn, override)
}

// createTestFunc generates a single test function that includes both the runner and the test body
//
// func Test_X_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
//
//	t.Parallel()
//	rapid.Check(t, func(t *rapid.T) {
//	    subject := XGenerator().Draw(t, "subject")
//	    copied := subject.DeepCopy()
//	    var hub HubVersion
//	    err := copied.ConvertTo(&hub)
//	    if err != nil { t.Fatalf("ConvertTo: %v", err) }
//	    var actual CurrentVersion
//	    err = actual.ConvertFrom(&hub)
//	    if err != nil { t.Fatalf("ConvertFrom: %v", err) }
//	    match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
//	    if !match { ... t.Errorf(result) }
//	})
//
// }
func (r *RapidResourceConversionTestCase) createTestFunc(
	subject astmodel.TypeName,
	codegenContext *astmodel.CodeGenerationContext,
) (dst.Decl, error) {
	const (
		copiedID     = "copied"
		hubID        = "hub"
		actualID     = "actual"
		actualFmtID  = "actualFmt"
		matchID      = "match"
		subjectID    = "subject"
		subjectFmtID = "subjectFmt"
		resultID     = "result"
		errID        = "err"
	)

	testingPackage := codegenContext.MustGetImportedPackageName(astmodel.TestingReference)
	rapidPackage := codegenContext.MustGetImportedPackageName(astmodel.RapidReference)
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

	// copied := subject.DeepCopy()
	assignCopied := astbuilder.ShortDeclaration(
		copiedID,
		astbuilder.CallQualifiedFunc(subjectID, "DeepCopy"))
	assignCopied.Decorations().Before = dst.NewLine
	astbuilder.AddComment(&assignCopied.Decorations().Start, "// Copy subject to make sure conversion doesn't modify it")

	// var hub OtherType
	hubExpr, err := r.toFn.Hub().AsTypeExpr(codegenContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", r.toFn.Hub())
	}

	declareOther := astbuilder.LocalVariableDeclaration(
		hubID,
		hubExpr,
		"// Convert to our hub version")
	declareOther.Decorations().Before = dst.EmptyLine

	// err := copied.ConvertTo(&hub)
	assignTo := astbuilder.ShortDeclaration(
		errID,
		astbuilder.CallQualifiedFunc(
			copiedID,
			r.toFn.Name(),
			astbuilder.AddrOf(dst.NewIdent(hubID))))

	// if err != nil { t.Fatalf("ConvertTo: %v", err) }
	assignToFailed := createRapidFatalf("ConvertTo", errID)

	// var actual OurType
	subjectExpr, err := subject.AsTypeExpr(codegenContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", subject)
	}

	declareResult := astbuilder.LocalVariableDeclaration(
		actualID,
		subjectExpr,
		"// Convert from our hub version")
	declareResult.Decorations().Before = dst.EmptyLine

	// err = actual.ConvertFrom(&hub)
	assignFrom := astbuilder.SimpleAssignment(
		dst.NewIdent(errID),
		astbuilder.CallQualifiedFunc(
			actualID,
			r.fromFn.Name(),
			astbuilder.AddrOf(dst.NewIdent(hubID))))

	// if err != nil { t.Fatalf("ConvertFrom: %v", err) }
	assignFromFailed := createRapidFatalf("ConvertFrom", errID)

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

	// result := diff.Diff(subjectFmt, actualFmt)
	declareDiff := astbuilder.ShortDeclaration(
		resultID,
		astbuilder.CallQualifiedFunc(diffPackage, "Diff", dst.NewIdent(subjectFmtID), dst.NewIdent(actualFmtID)))

	// t.Errorf(result)
	reportError := astbuilder.CallExprAsStmt(dst.NewIdent("t"), "Errorf", dst.NewIdent(resultID))

	// if !match { ... }
	prettyPrint := astbuilder.SimpleIf(
		astbuilder.NotExpr(dst.NewIdent(matchID)),
		declareActual,
		declareSubject,
		declareDiff,
		reportError)

	// Build the closure body
	closureBody := astbuilder.Statements(
		drawSubject,
		assignCopied,
		declareOther,
		assignTo,
		assignToFailed,
		declareResult,
		assignFrom,
		assignFromFailed,
		compare,
		prettyPrint)

	// t.Parallel()
	t := dst.NewIdent("t")
	declareParallel := astbuilder.CallExprAsStmt(t, "Parallel")

	// rapid.Check(t, func(t *rapid.T) { ... })
	rapidCheck := astbuilder.CallQualifiedFuncAsStmt(
		rapidPackage,
		"Check",
		t,
		&dst.FuncLit{
			Type: &dst.FuncType{
				Params: &dst.FieldList{
					List: []*dst.Field{
						{
							Names: []*dst.Ident{dst.NewIdent("t")},
							Type:  astbuilder.Dereference(astbuilder.QualifiedTypeName(rapidPackage, "T")),
						},
					},
				},
			},
			Body: &dst.BlockStmt{
				List: closureBody,
			},
		})

	fn := astbuilder.NewTestFuncDetails(
		testingPackage,
		r.testName,
		declareParallel,
		rapidCheck)

	fn.AddComments(fmt.Sprintf(
		"tests if a specific instance of %s round trips to the hub storage version and back losslessly",
		r.subject.Name()))

	return fn.DefineFunc(), nil
}
