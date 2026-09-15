/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"context"
	"encoding/base64"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys"
	"github.com/go-logr/logr"

	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20230701/storage"
	keys "github.com/Azure/azure-service-operator/v2/api/keyvault/v20230701/storage"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
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

func Test_KeyUpdatesNeeded_NothingManagedMeansNoUpdate(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// A spec that sets nothing mutable requires no update, whatever the key looks like
	key := &keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			AzureName: "my-key",
			Properties: &keys.KeyProperties{
				Kty: to.Ptr("RSA"),
			},
		},
	}
	actual := azkeys.KeyBundle{
		Attributes: &azkeys.KeyAttributes{Enabled: to.Ptr(false)},
		Tags:       map[string]*string{"unmanaged": to.Ptr("tag")},
		Key: &azkeys.JSONWebKey{
			KeyOps: []*azkeys.KeyOperation{to.Ptr(azkeys.KeyOperationEncrypt)},
		},
	}

	_, changed, err := keyUpdatesNeeded(key, actual)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeFalse())
}

func Test_KeyUpdatesNeeded_MatchingSpecMeansNoUpdate(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	exp := 1893456000 // 2030-01-01T00:00:00Z
	key := &keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			AzureName: "my-key",
			Tags:      map[string]string{"env": "test"},
			Properties: &keys.KeyProperties{
				KeyOps: []string{"encrypt", "decrypt"},
				Attributes: &keys.KeyAttributes{
					Enabled: to.Ptr(true),
					Exp:     to.Ptr(exp),
				},
			},
		},
	}
	actual := azkeys.KeyBundle{
		Attributes: &azkeys.KeyAttributes{
			Enabled: to.Ptr(true),
			Expires: to.Ptr(time.Unix(int64(exp), 0)),
		},
		Tags: map[string]*string{"env": to.Ptr("test")},
		Key: &azkeys.JSONWebKey{
			// Order deliberately differs from the spec; keyOps compare as a set
			KeyOps: []*azkeys.KeyOperation{
				to.Ptr(azkeys.KeyOperationDecrypt),
				to.Ptr(azkeys.KeyOperationEncrypt),
			},
		},
	}

	_, changed, err := keyUpdatesNeeded(key, actual)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeFalse())
}

func Test_KeyUpdatesNeeded_DivergedPropertiesAreUpdated(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	key := &keys.VaultKey{
		Spec: keys.VaultKey_Spec{
			AzureName: "my-key",
			Tags:      map[string]string{"env": "prod"},
			Properties: &keys.KeyProperties{
				KeyOps: []string{"sign", "verify"},
				Attributes: &keys.KeyAttributes{
					Enabled: to.Ptr(true),
				},
			},
		},
	}
	actual := azkeys.KeyBundle{
		Attributes: &azkeys.KeyAttributes{Enabled: to.Ptr(false)},
		Tags:       map[string]*string{"env": to.Ptr("test")},
		Key: &azkeys.JSONWebKey{
			KeyOps: []*azkeys.KeyOperation{to.Ptr(azkeys.KeyOperationEncrypt)},
		},
	}

	params, changed, err := keyUpdatesNeeded(key, actual)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeTrue())
	g.Expect(params.KeyOps).To(HaveLen(2))
	g.Expect(params.KeyAttributes).ToNot(BeNil())
	g.Expect(params.KeyAttributes.Enabled).To(HaveValue(BeTrue()))
	// Expiry was not set in the spec, so the update must not touch it
	g.Expect(params.KeyAttributes.Expires).To(BeNil())
	g.Expect(params.Tags).To(HaveKeyWithValue("env", HaveValue(Equal("prod"))))
}

func Test_ReleasePolicyUpdateNeeded(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	policyJSON := []byte(`{"anyOf":[]}`)
	encoded := base64.RawURLEncoding.EncodeToString(policyJSON)

	spec := &keys.KeyReleasePolicy{Data: to.Ptr(encoded)}

	// Matching policy: no update
	_, changed, err := releasePolicyUpdateNeeded(spec, &azkeys.KeyReleasePolicy{EncodedPolicy: policyJSON})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeFalse())

	// Different policy: update with the decoded bytes
	desired, changed, err := releasePolicyUpdateNeeded(spec, &azkeys.KeyReleasePolicy{EncodedPolicy: []byte(`{}`)})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeTrue())
	g.Expect(desired.EncodedPolicy).To(Equal(policyJSON))

	// No policy on the key at all: update
	_, changed, err = releasePolicyUpdateNeeded(spec, nil)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(changed).To(BeTrue())

	// Invalid base64 in the spec is an error, not a silent no-op
	_, _, err = releasePolicyUpdateNeeded(&keys.KeyReleasePolicy{Data: to.Ptr("!!not-base64!!")}, nil)
	g.Expect(err).To(HaveOccurred())
}

func Test_RotationPolicyUpdateNeeded(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	spec := &keys.RotationPolicy{
		Attributes: &keys.KeyRotationPolicyAttributes{ExpiryTime: to.Ptr("P2Y")},
		LifetimeActions: []keys.LifetimeAction{
			{
				// ARM spells the action lowercase; the data plane answers with "Rotate".
				// The comparison must be case-insensitive per the service contract.
				Action:  &keys.Action{Type: to.Ptr("rotate")},
				Trigger: &keys.Trigger{TimeAfterCreate: to.Ptr("P90D")},
			},
		},
	}

	matching := azkeys.KeyRotationPolicy{
		Attributes: &azkeys.KeyRotationPolicyAttributes{ExpiryTime: to.Ptr("P2Y")},
		LifetimeActions: []*azkeys.LifetimeAction{
			{
				Action:  &azkeys.LifetimeActionType{Type: to.Ptr(azkeys.KeyRotationPolicyActionRotate)},
				Trigger: &azkeys.LifetimeActionTrigger{TimeAfterCreate: to.Ptr("P90D")},
			},
			{
				// Service-added default notify action must be tolerated, or every
				// reconcile would apply the policy again forever
				Action:  &azkeys.LifetimeActionType{Type: to.Ptr(azkeys.KeyRotationPolicyActionNotify)},
				Trigger: &azkeys.LifetimeActionTrigger{TimeBeforeExpiry: to.Ptr("P30D")},
			},
		},
	}

	_, changed := rotationPolicyUpdateNeeded(spec, matching)
	g.Expect(changed).To(BeFalse())

	// Different trigger: update, carrying the spec's policy with canonical action casing
	diverged := azkeys.KeyRotationPolicy{
		Attributes: &azkeys.KeyRotationPolicyAttributes{ExpiryTime: to.Ptr("P2Y")},
		LifetimeActions: []*azkeys.LifetimeAction{
			{
				Action:  &azkeys.LifetimeActionType{Type: to.Ptr(azkeys.KeyRotationPolicyActionRotate)},
				Trigger: &azkeys.LifetimeActionTrigger{TimeAfterCreate: to.Ptr("P180D")},
			},
		},
	}

	desired, changed := rotationPolicyUpdateNeeded(spec, diverged)
	g.Expect(changed).To(BeTrue())
	g.Expect(desired.Attributes.ExpiryTime).To(HaveValue(Equal("P2Y")))
	g.Expect(desired.LifetimeActions).To(HaveLen(1))
	g.Expect(desired.LifetimeActions[0].Action.Type).To(HaveValue(Equal(azkeys.KeyRotationPolicyActionRotate)))

	// Missing expiryTime on the actual policy: update
	_, changed = rotationPolicyUpdateNeeded(spec, azkeys.KeyRotationPolicy{LifetimeActions: matching.LifetimeActions})
	g.Expect(changed).To(BeTrue())
}

func Test_VaultKeyExtension_PostReconcileCheck_RejectsUnexpectedResourceType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ext := &VaultKeyExtension{}

	_, err := ext.PostReconcileCheck(
		context.Background(), &keyvault.Vault{}, nil, nil, nil, logr.Discard(), annotations.ResolvedReconcilePolicies{}, nil,
	)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("unexpected resource type"))
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
