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

// Test_VaultKey_SupportedOperations_Allowlist guards the safety-critical routing of VaultKey
// deletion. The reconciler routes deletion to DeleteNotPossibleInAzure, and so to the
// VaultKeyExtension's deleteMode handling, only because ResourceOperationDelete is absent from the
// supported operations, which the generator derives mechanically from the swagger. A future spec
// bump that adds DELETE to Microsoft.KeyVault/vaults/keys would otherwise silently turn every
// VaultKey deletion into an ARM delete of live key material; this test makes that change explicit.
//
// Both the versioned and the storage (hub) type are asserted: the reconciler operates on the hub.
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

// Test_VaultKey_SecuritySensitiveFieldSets fails when regeneration changes the field set of the
// types that shape a key's material (type, size, curve, exportability, release and rotation
// policy). Such a change alters the security surface of the resource and the checks in the
// validating webhook and VaultKeyExtension, so it should be reviewed deliberately rather than pass
// as part of a routine regeneration diff. Update the expected list once the change is understood.
func Test_VaultKey_SecuritySensitiveFieldSets(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		value    any
		expected []string
	}{
		"VaultKey_Spec":    {VaultKey_Spec{}, []string{"AzureName", "OperatorSpec", "Owner", "Properties", "Tags"}},
		"KeyProperties":    {KeyProperties{}, []string{"Attributes", "CurveName", "KeyOps", "KeySize", "Kty", "Release_Policy", "RotationPolicy"}},
		"KeyAttributes":    {KeyAttributes{}, []string{"Enabled", "Exp", "Exportable", "Nbf"}},
		"KeyReleasePolicy": {KeyReleasePolicy{}, []string{"ContentType", "Data"}},
		"RotationPolicy":   {RotationPolicy{}, []string{"Attributes", "LifetimeActions"}},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			typ := reflect.TypeOf(c.value)
			names := make([]string, 0, typ.NumField())
			for i := 0; i < typ.NumField(); i++ {
				names = append(names, typ.Field(i).Name)
			}
			sort.Strings(names)

			expected := append([]string(nil), c.expected...)
			sort.Strings(expected)

			g.Expect(names).To(Equal(expected))
		})
	}
}
