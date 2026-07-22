/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package webhook

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v20230701 "github.com/Azure/azure-service-operator/v2/api/keyvault/v20230701"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func newTestVaultsKeyObj() *v20230701.VaultsKey {
	kty := v20230701.KeyProperties_Kty_RSA
	return &v20230701.VaultsKey{
		ObjectMeta: metav1.ObjectMeta{
			Name: "mykey",
		},
		Spec: v20230701.VaultsKey_Spec{
			AzureName: "mykey",
			Properties: &v20230701.KeyProperties{
				Kty:     &kty,
				KeySize: to.Ptr(2048),
				Attributes: &v20230701.KeyAttributes{
					Enabled: to.Ptr(true),
				},
			},
		},
	}
}

func markCreated(obj *v20230701.VaultsKey) *v20230701.VaultsKey {
	genruntime.SetResourceID(obj, "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.KeyVault/vaults/myvault/keys/mykey")
	return obj
}

func Test_VaultsKey_ValidateNotExportable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	webhook := &VaultsKey{}

	t.Run("create with exportable=true is rejected", func(t *testing.T) {
		obj := newTestVaultsKeyObj()
		obj.Spec.Properties.Attributes.Exportable = to.Ptr(true)

		_, err := webhook.validateNotExportable(context.Background(), obj)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("exportable"))
	})

	t.Run("create with exportable=false is allowed", func(t *testing.T) {
		obj := newTestVaultsKeyObj()
		obj.Spec.Properties.Attributes.Exportable = to.Ptr(false)

		_, err := webhook.validateNotExportable(context.Background(), obj)
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("create with exportable unset is allowed", func(t *testing.T) {
		obj := newTestVaultsKeyObj()

		_, err := webhook.validateNotExportable(context.Background(), obj)
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("update to exportable=true is rejected", func(t *testing.T) {
		newObj := markCreated(newTestVaultsKeyObj())
		newObj.Spec.Properties.Attributes.Exportable = to.Ptr(true)

		_, err := webhook.validateNotExportable(context.Background(), newObj)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("exportable"))
	})
}

func Test_VaultsKey_ValidateImmutable(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	webhook := &VaultsKey{}

	t.Run("not yet created - any change allowed", func(t *testing.T) {
		oldObj := newTestVaultsKeyObj()
		newObj := newTestVaultsKeyObj()
		newObj.Spec.Properties.KeySize = to.Ptr(4096)

		_, err := webhook.validateImmutable(context.Background(), oldObj, newObj)
		g.Expect(err).ToNot(HaveOccurred())
	})

	t.Run("created - changing keySize is rejected", func(t *testing.T) {
		oldObj := markCreated(newTestVaultsKeyObj())
		newObj := markCreated(newTestVaultsKeyObj())
		newObj.Spec.Properties.KeySize = to.Ptr(4096)

		_, err := webhook.validateImmutable(context.Background(), oldObj, newObj)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("spec.properties"))
	})

	t.Run("created - changing tags is rejected", func(t *testing.T) {
		oldObj := markCreated(newTestVaultsKeyObj())
		newObj := markCreated(newTestVaultsKeyObj())
		newObj.Spec.Tags = map[string]string{"foo": "bar"}

		_, err := webhook.validateImmutable(context.Background(), oldObj, newObj)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("spec.tags"))
	})

	t.Run("created - true no-op update is allowed", func(t *testing.T) {
		oldObj := markCreated(newTestVaultsKeyObj())
		newObj := markCreated(newTestVaultsKeyObj())

		_, err := webhook.validateImmutable(context.Background(), oldObj, newObj)
		g.Expect(err).ToNot(HaveOccurred())
	})
}
