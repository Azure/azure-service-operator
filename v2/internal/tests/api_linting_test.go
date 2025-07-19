/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package tests

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"

	"github.com/Azure/azure-service-operator/v2/api"
	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"
)

// TypeGroupInfo represents information about a type group (same group + kind, different versions)
type TypeGroupInfo struct {
	GroupKind schema.GroupKind
	Types     []*registration.KnownType
}

// groupTypesByGroupKind groups the known types by their GroupKind (group + kind)
// This allows us to identify related types across different API versions
func groupTypesByGroupKind(knownTypes []*registration.KnownType) (map[schema.GroupKind]*TypeGroupInfo, error) {
	groups := make(map[schema.GroupKind]*TypeGroupInfo)
	scheme := api.CreateScheme()

	for _, knownType := range knownTypes {
		// Get the GroupVersionKind for this type
		gvk, err := apiutil.GVKForObject(knownType.Obj, scheme)
		if err != nil {
			return nil, fmt.Errorf("creating GVK for obj %s: %w", knownType, err)
		}
		groupKind := gvk.GroupKind()

		// Exclude storage types
		if strings.HasSuffix(gvk.Version, "storage") {
			continue
		}

		// Create or get the group info
		if groups[groupKind] == nil {
			groups[groupKind] = &TypeGroupInfo{
				GroupKind: groupKind,
				Types:     []*registration.KnownType{},
			}
		}

		knownType.Obj.GetObjectKind().SetGroupVersionKind(gvk)

		// Add this type to the group
		groups[groupKind].Types = append(groups[groupKind].Types, knownType)
	}

	return groups, nil
}

// implementsDefaulter checks if a resource type implements the genruntime.Defaulter interface
func implementsDefaulter(obj any) bool {
	// Use type assertion to check if the object ok genruntime.Defaulter
	_, ok := obj.(genruntime.Defaulter)
	return ok
}

// implementsValidator checks if a resource type implements the genruntime.Validator[T] interface.
// Since genruntime.Validator is a generic interface we can't use a type assertion or simple reflection.
// Instead, we check if the type has the required methods with the correct signatures.
// Note that this is horrible, but Go doesn't have a way to check generic interfaces directly so this is
// the best that we can do.
func implementsValidator(obj any) bool {
	objType := reflect.TypeOf(obj)
	if objType == nil {
		return false
	}

	// Check if the type has the three required methods:
	// CreateValidations() []func(ctx context.Context, obj T) (admission.Warnings, error)
	// UpdateValidations() []func(ctx context.Context, oldObj T, newObj T) (admission.Warnings, error)
	// DeleteValidations() []func(ctx context.Context, obj T) (admission.Warnings, error)

	hasCreate := implementsValidatorMethod(objType, "CreateValidations")
	hasUpdate := implementsValidatorMethod(objType, "UpdateValidations")
	hasDelete := implementsValidatorMethod(objType, "DeleteValidations")

	if !hasCreate || !hasUpdate || !hasDelete {
		return false
	}

	return true
}

func implementsValidatorMethod(objType reflect.Type, methodName string) bool {
	method, hasMethod := objType.MethodByName(methodName)
	if !hasMethod {
		return false
	}

	methodType := method.Type

	// Check that the method has 1 input (receiver) and 1 output (slice of validation functions)
	if methodType.NumIn() != 1 || methodType.NumOut() != 1 {
		return false
	}

	// Check that the return type is a slice
	returnType := methodType.Out(0)
	if returnType.Kind() != reflect.Slice {
		return false
	}

	// Check that the slice element types are function types
	elemType := returnType.Elem()
	if elemType.Kind() != reflect.Func {
		return false
	}

	return true
}

// validateConsistency checks if all types in a group consistently pass the checkFunc.
// Returns an error if types within a group + kind are inconsistent.
func validateConsistency(
	group *TypeGroupInfo,
	selector func(*registration.KnownType) any,
	checkName string,
	checkFunc func(any) bool,
) error {
	var passingTypes []string
	var failingTypes []string

	for _, knownType := range group.Types {
		typeName := knownType.Obj.GetObjectKind().GroupVersionKind().Version
		obj := selector(knownType)

		if obj == nil {
			// If the selector returns nil, we skip this type
			continue
		}

		if checkFunc(obj) {
			passingTypes = append(passingTypes, typeName)
		} else {
			failingTypes = append(failingTypes, typeName)
		}
	}

	if len(passingTypes) == 0 || len(failingTypes) == 0 {
		return nil
	}

	// There's an inconsistency - return error
	return fmt.Errorf(
		"%[1]s inconsistency for %[2]s: [%[3]s](len=%[4]d) had %[1]s=true , but [%[5]s](len=%[6]d) had %[1]s=false",
		checkName,
		group.GroupKind.String(),
		strings.Join(passingTypes, ", "),
		len(passingTypes),
		strings.Join(failingTypes, ", "),
		len(failingTypes))
}

func Test_APILinting_Defaulter_Consistency(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Get all known types and group them by GroupKind
	knownTypes := controllers.GetKnownTypes()
	groups, err := groupTypesByGroupKind(knownTypes)
	g.Expect(err).ToNot(HaveOccurred())

	defaulterSelector := func(kt *registration.KnownType) any {
		return kt.Defaulter
	}
	// Check each group for Defaulter interface consistency
	for _, group := range groups {
		// Skip groups with only one type (no consistency check needed)
		if len(group.Types) <= 1 {
			continue
		}

		g.Expect(validateConsistency(group, defaulterSelector, "genruntime.Defaulter", implementsDefaulter)).To(Succeed())
	}
}

func Test_APILinting_Validator_Consistency(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Get all known types and group them by GroupKind
	knownTypes := controllers.GetKnownTypes()
	groups, err := groupTypesByGroupKind(knownTypes)
	g.Expect(err).ToNot(HaveOccurred())

	validatorSelector := func(kt *registration.KnownType) any {
		return kt.Validator
	}

	// Check each group for Validator interface consistency
	for _, group := range groups {
		// Skip groups with only one type (no consistency check needed)
		if len(group.Types) <= 1 {
			continue
		}

		g.Expect(validateConsistency(group, validatorSelector, "genruntime.Validator", implementsValidator)).To(Succeed())
	}
}
