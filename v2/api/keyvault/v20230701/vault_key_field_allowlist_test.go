/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v20230701

import (
	"reflect"
	"sort"
	"testing"

	. "github.com/onsi/gomega"

	storage "github.com/Azure/azure-service-operator/v2/api/keyvault/v20230701/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// This test is a safety net, not a style test. VaultKey's spec surfaces (directly or indirectly)
// key material and key-material-affecting configuration (kty, key size, curve, exportability, release
// policy, rotation policy, etc). If a future codegen run against an updated Key Vault swagger silently
// adds, renames, or removes a field on any of these types, that is a MATERIAL CHANGE to the security
// surface of this resource (e.g. a new field could be a new way to exfiltrate key material, weaken a
// key, or bypass the immutability/exportability checks in the VaultKey validating webhook) and should
// never pass silently through code review. This test intentionally fails loudly - and requires the
// allowlist below to be updated by hand - whenever the field set of these types changes, forcing an
// explicit, deliberate review of the change.
func fieldNamesOf(t *testing.T, value any) []string {
	t.Helper()
	typ := reflect.TypeOf(value)
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		t.Fatalf("expected a struct type, got %s", typ.Kind())
	}

	names := make([]string, 0, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		names = append(names, typ.Field(i).Name)
	}
	sort.Strings(names)
	return names
}

func Test_VaultKey_Spec_FieldAllowlist(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := []string{
		"AzureName",
		"OperatorSpec",
		"Owner",
		"Properties",
		"Tags",
	}
	sort.Strings(expected)

	g.Expect(fieldNamesOf(t, VaultKey_Spec{})).To(Equal(expected))
}

func Test_KeyProperties_FieldAllowlist(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := []string{
		"Attributes",
		"CurveName",
		"KeyOps",
		"KeySize",
		"Kty",
		"Release_Policy",
		"RotationPolicy",
	}
	sort.Strings(expected)

	g.Expect(fieldNamesOf(t, KeyProperties{})).To(Equal(expected))
}

func Test_KeyAttributes_FieldAllowlist(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := []string{
		"Enabled",
		"Exp",
		"Exportable",
		"Nbf",
	}
	sort.Strings(expected)

	g.Expect(fieldNamesOf(t, KeyAttributes{})).To(Equal(expected))
}

func Test_KeyReleasePolicy_FieldAllowlist(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := []string{
		"ContentType",
		"Data",
	}
	sort.Strings(expected)

	g.Expect(fieldNamesOf(t, KeyReleasePolicy{})).To(Equal(expected))
}

func Test_RotationPolicy_FieldAllowlist(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := []string{
		"Attributes",
		"LifetimeActions",
	}
	sort.Strings(expected)

	g.Expect(fieldNamesOf(t, RotationPolicy{})).To(Equal(expected))
}

// Test_VaultKey_SupportedOperations_Allowlist is a safety net, not a style test. VaultKey
// deliberately relies on the ARM API exposing no DELETE verb: DetermineDeleteAction routes
// deletion to the built-in DeleteNotPossibleInAzure handling precisely because
// ResourceOperationDelete is absent from this list, which is what guarantees ASO can never
// destroy live key material. SupportedOperations is derived mechanically from the swagger
// (see getSupportedOperations in the generator's swagger_type_extractor.go), with no
// configuration override, so a future spec bump that adds DELETE to
// Microsoft.KeyVault/vaults/keys would silently turn VaultKey into a resource that issues
// real ARM deletes against live cryptographic keys - inside an otherwise unremarkable
// regeneration diff. This test forces that change to be explicit and reviewed.
//
// Both the versioned type and the storage (hub) type are asserted: the reconciler operates
// on the hub type, so that is the one that strictly must be guarded; the versioned type is
// included for symmetry.
func Test_VaultKey_SupportedOperations_Allowlist(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect((&VaultKey{}).GetSupportedOperations()).To(ConsistOf(
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	))

	g.Expect((&storage.VaultKey{}).GetSupportedOperations()).To(ConsistOf(
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	))
}
