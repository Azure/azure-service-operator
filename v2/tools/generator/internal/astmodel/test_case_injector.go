/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// TestCaseInjector is a utility for injecting function definitions into resources and objects
type TestCaseInjector struct {
	// visitor is used to do the actual injection
	visitor TypeVisitor[any]
}

// NewTestCaseInjector creates a new function injector for modifying resources and objects
func NewTestCaseInjector() *TestCaseInjector {
	result := &TestCaseInjector{}

	result.visitor = TypeVisitorBuilder[any]{
		VisitObjectType:   result.injectTestCaseIntoObject,
		VisitResourceType: result.injectTestCaseIntoResource,
	}.Build()

	return result
}

// Inject modifies the passed type definition by injecting the passed function
func (fi *TestCaseInjector) Inject(def TypeDefinition, cases ...TestCase) (TypeDefinition, error) {
	result := def

	for _, fn := range cases {
		var err error
		result, err = fi.visitor.VisitDefinition(result, fn)
		if err != nil {
			return TypeDefinition{}, err
		}
	}

	return result, nil
}

// injectTestCaseIntoObject takes the function provided as a context and includes it on the
// provided object type
func (_ *TestCaseInjector) injectTestCaseIntoObject(
	_ *TypeVisitor[any], ot *ObjectType, ctx any) (Type, error) {
	fn := ctx.(TestCase)
	return ot.WithTestCase(fn), nil
}

// injectTestCaseIntoResource takes the function provided as a context and includes it on the
// provided resource type
func (_ *TestCaseInjector) injectTestCaseIntoResource(
	_ *TypeVisitor[any], rt *ResourceType, ctx any) (Type, error) {
	fn := ctx.(TestCase)
	return rt.WithTestCase(fn), nil
}
