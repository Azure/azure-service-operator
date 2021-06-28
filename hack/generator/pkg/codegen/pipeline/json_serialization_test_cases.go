/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
	"github.com/Azure/azure-service-operator/hack/generator/pkg/testcases"

	kerrors "k8s.io/apimachinery/pkg/util/errors"
)

func InjectJsonSerializationTests(idFactory astmodel.IdentifierFactory) Stage {

	return MakeStage(
		"jsonTestCases",
		"Add test cases to verify JSON serialization",
		func(ctx context.Context, types astmodel.Types) (astmodel.Types, error) {
			factory := makeObjectSerializationTestCaseFactory(idFactory)
			result := make(astmodel.Types)
			var errs []error
			for _, d := range types {
				updated, err := factory.AddTestTo(d)
				if err != nil {
					errs = append(errs, err)
				} else {
					result[updated.Name()] = updated
				}
			}

			if len(errs) > 0 {
				return nil, kerrors.NewAggregate(errs)
			}

			return result, nil
		})
}

type objectSerializationTestCaseFactory struct {
	visitor   astmodel.TypeVisitor
	idFactory astmodel.IdentifierFactory
}

func makeObjectSerializationTestCaseFactory(idFactory astmodel.IdentifierFactory) objectSerializationTestCaseFactory {
	result := objectSerializationTestCaseFactory{
		idFactory: idFactory,
	}

	result.visitor = astmodel.TypeVisitorBuilder{
		VisitObjectType: result.injectTestCase,
	}.Build()

	return result
}

func (s *objectSerializationTestCaseFactory) AddTestTo(def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	return s.visitor.VisitDefinition(def, def.Name())
}

func (s *objectSerializationTestCaseFactory) injectTestCase(
	_ *astmodel.TypeVisitor, objectType *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
	name := ctx.(astmodel.TypeName)
	testcase := testcases.NewObjectSerializationTestCase(name, objectType, s.idFactory)
	result := objectType.WithTestCase(testcase)
	return result, nil
}
