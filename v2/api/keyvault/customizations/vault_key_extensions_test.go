/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

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
