/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys"
	"github.com/go-logr/logr"

	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20230701/storage"
	keys "github.com/Azure/azure-service-operator/v2/api/keyvault/v20230701/storage"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_VaultKeyDeleteMode_DefaultsToDetach(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(vaultKeyDeleteMode(&keys.VaultKey{})).To(Equal(DeleteMode_Detach))

	g.Expect(vaultKeyDeleteMode(&keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			OperatorSpec: &keys.VaultKeyOperatorSpec{},
		},
	})).To(Equal(DeleteMode_Detach))

	g.Expect(vaultKeyDeleteMode(&keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			OperatorSpec: &keys.VaultKeyOperatorSpec{
				DeleteMode: to.Ptr(DeleteMode_Delete),
			},
		},
	})).To(Equal(DeleteMode_Delete))
}

// Detach must not require any Azure connectivity at all: it succeeds with no resolver and no ARM
// client, proving no data-plane or ARM call can be involved in the default deletion path.
func Test_VaultKeyExtension_Delete_DetachTouchesNothing(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ext := &VaultKeyExtension{}
	key := &keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			AzureName: "my-key",
		},
	}

	result, err := ext.Delete(context.Background(), logr.Discard(), nil, nil, key, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.Completed()).To(BeTrue())
}

func Test_VaultKeyExtension_Delete_RejectsUnexpectedResourceType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ext := &VaultKeyExtension{}

	_, err := ext.Delete(context.Background(), logr.Discard(), nil, nil, &keyvault.Vault{}, nil)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("unexpected resource type"))
}

func Test_VaultKeyExtension_Delete_RejectsUnknownDeleteMode(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ext := &VaultKeyExtension{}
	key := &keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			AzureName: "my-key",
			OperatorSpec: &keys.VaultKeyOperatorSpec{
				DeleteMode: to.Ptr("obliterate"),
			},
		},
	}

	_, err := ext.Delete(context.Background(), logr.Discard(), nil, nil, key, nil)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("deleteMode"))
}

func Test_VaultKeyCreateMode_DefaultsToDefault(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	g.Expect(vaultKeyCreateMode(&keys.VaultKey{})).To(Equal(CreateMode_Default))

	g.Expect(vaultKeyCreateMode(&keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			OperatorSpec: &keys.VaultKeyOperatorSpec{
				CreateMode: to.Ptr(CreateMode_CreateOrRecover),
			},
		},
	})).To(Equal(CreateMode_CreateOrRecover))
}

func Test_VaultKeyExtension_PreReconcileCheck_RejectsUnexpectedResourceType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ext := &VaultKeyExtension{}

	_, err := ext.PreReconcileCheck(context.Background(), &keyvault.Vault{}, nil, nil, logr.Discard(), nil)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("unexpected resource type"))
}

func Test_IntrinsicMismatch(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	keyWith := func(props *keys.KeyProperties) *keys.VaultKey {
		return &keys.VaultKey{
			Spec: keys.VaultKey_Spec{
				AzureName:  "my-key",
				Properties: props,
			},
		}
	}

	rsa2048 := &azkeys.JSONWebKey{
		Kty: to.Ptr(azkeys.KeyTypeRSA),
		N:   make([]byte, 256), // 2048-bit modulus
	}
	ecP256 := &azkeys.JSONWebKey{
		Kty: to.Ptr(azkeys.KeyTypeEC),
		Crv: to.Ptr(azkeys.CurveNameP256),
	}

	// A spec with no properties can't conflict
	g.Expect(intrinsicMismatch(keyWith(nil), rsa2048)).To(BeEmpty())

	// Matching properties pass
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{
		Kty:     to.Ptr("RSA"),
		KeySize: to.Ptr(2048),
	}), rsa2048)).To(BeEmpty())
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{
		Kty:       to.Ptr("EC"),
		CurveName: to.Ptr("P-256"),
	}), ecP256)).To(BeEmpty())

	// Unset spec properties match anything
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{}), rsa2048)).To(BeEmpty())

	// Mismatches are reported per property
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{
		Kty: to.Ptr("EC"),
	}), rsa2048)).To(ContainSubstring("kty"))
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{
		Kty:     to.Ptr("RSA"),
		KeySize: to.Ptr(4096),
	}), rsa2048)).To(ContainSubstring("keySize"))
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{
		Kty:       to.Ptr("EC"),
		CurveName: to.Ptr("P-384"),
	}), ecP256)).To(ContainSubstring("curveName"))

	// keySize can't be verified for keys without an RSA modulus, so it's left to the service
	g.Expect(intrinsicMismatch(keyWith(&keys.KeyProperties{
		Kty:     to.Ptr("EC"),
		KeySize: to.Ptr(2048),
	}), ecP256)).To(BeEmpty())
}

func Test_VaultURLFromKeyURI(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	vaultURL, err := vaultURLFromKeyURI("https://myvault.vault.azure.net/keys/my-key")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(vaultURL).To(Equal("https://myvault.vault.azure.net"))

	// Sovereign clouds use a different DNS suffix; nothing may be hardcoded
	vaultURL, err = vaultURLFromKeyURI("https://myvault.vault.azure.cn/keys/my-key/abc123")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(vaultURL).To(Equal("https://myvault.vault.azure.cn"))

	_, err = vaultURLFromKeyURI("not-a-uri")
	g.Expect(err).To(HaveOccurred())
}
