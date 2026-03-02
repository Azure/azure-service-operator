/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcases

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astbuilder"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/conversions"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/functions"
)

// RapidPropertyAssignmentTestCase represents a test that verifies we can convert from one object in our conversion
// graph to the next one, and back again, with no loss of data, using the rapid property testing library.
type RapidPropertyAssignmentTestCase struct {
	testName  string
	subject   astmodel.TypeName
	toFn      *functions.PropertyAssignmentFunction
	fromFn    *functions.PropertyAssignmentFunction
	idFactory astmodel.IdentifierFactory
}

var _ astmodel.TestCase = &RapidPropertyAssignmentTestCase{}

func NewRapidPropertyAssignmentTestCase(
	name astmodel.TypeName,
	container astmodel.FunctionContainer,
	idFactory astmodel.IdentifierFactory,
) *RapidPropertyAssignmentTestCase {
	result := &RapidPropertyAssignmentTestCase{
		subject:   name,
		idFactory: idFactory,
	}

	// Find Property Assignment functions
	for _, fn := range container.Functions() {
		if !strings.HasPrefix(fn.Name(), conversions.AssignPropertiesMethodPrefix) {
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
func (r *RapidPropertyAssignmentTestCase) Name() string {
	return r.testName
}

// References returns the set of types to which this test case refers.
func (r *RapidPropertyAssignmentTestCase) References() astmodel.TypeNameSet {
	return astmodel.NewTypeNameSet(
		r.subject,
		r.toFn.ParameterType())
}

// RequiredImports returns a set of the package imports required by this test case
func (r *RapidPropertyAssignmentTestCase) RequiredImports() *astmodel.PackageImportSet {
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

	result.AddImportOfReference(r.toFn.ParameterType().PackageReference())

	return result
}

// AsFuncs renders the current test case and any supporting methods as Go abstract syntax trees
func (r *RapidPropertyAssignmentTestCase) AsFuncs(
	receiver astmodel.TypeName,
	codeGenerationContext *astmodel.CodeGenerationContext,
) ([]dst.Decl, error) {
	testFunc, err := r.createTestFunc(receiver, codeGenerationContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating test function for %s", r.testName)
	}

	return []dst.Decl{testFunc}, nil
}

// Equals determines if this TestCase is equal to another one
func (r *RapidPropertyAssignmentTestCase) Equals(other astmodel.TestCase, override astmodel.EqualityOverrides) bool {
	fn, ok := other.(*RapidPropertyAssignmentTestCase)
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
// func Test_X_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
//
//	t.Parallel()
//	rapid.Check(t, func(t *rapid.T) {
//	    subject := XGenerator().Draw(t, "subject")
//	    copied := subject.DeepCopy()
//	    var other OtherType
//	    err := copied.AssignPropertiesTo(&other)
//	    if err != nil { t.Fatalf("AssignTo: %v", err) }
//	    var actual CurrentType
//	    err = actual.AssignPropertiesFrom(&other)
//	    if err != nil { t.Fatalf("AssignFrom: %v", err) }
//	    match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
//	    if !match { ... t.Errorf(result) }
//	})
//
// }
func (r *RapidPropertyAssignmentTestCase) createTestFunc(
	subject astmodel.TypeName,
	codegenContext *astmodel.CodeGenerationContext,
) (dst.Decl, error) {
	const (
		copiedID     = "copied"
		otherID      = "other"
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
	astbuilder.AddComment(&assignCopied.Decorations().Start, "// Copy subject to make sure assignment doesn't modify it")

	// var other OtherType
	parameterTypeExpr, err := r.toFn.ParameterType().AsTypeExpr(codegenContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", r.toFn.ParameterType())
	}

	declareOther := astbuilder.LocalVariableDeclaration(
		otherID,
		parameterTypeExpr,
		"// Use AssignPropertiesTo() for the first stage of conversion")
	declareOther.Decorations().Before = dst.EmptyLine

	// err := copied.AssignPropertiesTo(&other)
	assignTo := astbuilder.ShortDeclaration(
		errID,
		astbuilder.CallQualifiedFunc(
			copiedID,
			r.toFn.Name(),
			astbuilder.AddrOf(dst.NewIdent(otherID))))

	// if err != nil { t.Fatalf("AssignTo: %v", err) }
	assignToFailed := createRapidFatalf("AssignPropertiesTo", errID)

	// var actual OurType
	subjectExpr, err := subject.AsTypeExpr(codegenContext)
	if err != nil {
		return nil, eris.Wrapf(err, "creating type expression for %s", subject)
	}

	declareResult := astbuilder.LocalVariableDeclaration(
		actualID,
		subjectExpr,
		"// Use AssignPropertiesFrom() to convert back to our original type")
	declareResult.Decorations().Before = dst.EmptyLine

	// err = actual.AssignPropertiesFrom(&other)
	assignFrom := astbuilder.SimpleAssignment(
		dst.NewIdent(errID),
		astbuilder.CallQualifiedFunc(
			actualID,
			r.fromFn.Name(),
			astbuilder.AddrOf(dst.NewIdent(otherID))))

	// if err != nil { t.Fatalf("AssignFrom: %v", err) }
	assignFromFailed := createRapidFatalf("AssignPropertiesFrom", errID)

	// match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	equateEmpty := astbuilder.CallQualifiedFunc(cmpoptsPackage, "EquateEmpty")
	compare := astbuilder.ShortDeclaration(
		matchID,
		astbuilder.CallQualifiedFunc(cmpPackage, "Equal",
			dst.NewIdent(subjectID),
			dst.NewIdent(actualID),
			equateEmpty))
	compare.Decorations().Before = dst.EmptyLine
	astbuilder.AddComment(&compare.Decorations().Start, "Check for a match")

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

	// rapid.Check(t, func(t *rapid.T) { ... })
	t := dst.NewIdent("t")

	// t.Parallel()
	declareParallel := astbuilder.CallExprAsStmt(t, "Parallel")

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
		"tests if a specific instance of %s can be assigned to %s and back losslessly",
		r.subject.Name(),
		r.fromFn.ParameterType().PackageReference().PackageName()))

	return fn.DefineFunc(), nil
}

// createRapidFatalf generates: if <errID> != nil { t.Fatalf("<context>: %%v", <errID>) }
func createRapidFatalf(context string, errID string) *dst.IfStmt {
	return &dst.IfStmt{
		Cond: astbuilder.AreNotEqual(dst.NewIdent(errID), dst.NewIdent("nil")),
		Body: astbuilder.StatementBlock(
			astbuilder.CallExprAsStmt(
				dst.NewIdent("t"),
				"Fatalf",
				astbuilder.StringLiteralf("%s: %%v", context),
				dst.NewIdent(errID))),
	}
}
