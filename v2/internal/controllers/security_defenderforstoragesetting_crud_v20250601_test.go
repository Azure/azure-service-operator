/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	security "github.com/Azure/azure-service-operator/v2/api/security/v20250601"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Security_DefenderForStorageSetting_v20250601_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()
	account := newStorageAccount20250601(tc, rg)
	setting := &security.DefenderForStorageSetting{
		ObjectMeta: tc.MakeObjectMeta("defender"),
		Spec: security.DefenderForStorageSetting_Spec{
			Owner: &genruntime.ArbitraryOwnerReference{
				Group: "storage.azure.com",
				Kind:  "StorageAccount",
				Name:  account.Name,
			},
			Properties: &security.DefenderForStorageSettingProperties{
				IsEnabled:                         to.Ptr(true),
				OverrideSubscriptionLevelSettings: to.Ptr(true),
				MalwareScanning: &security.MalwareScanningProperties{
					OnUpload: &security.OnUploadProperties{
						IsEnabled:     to.Ptr(true),
						CapGBPerMonth: to.Ptr(1),
					},
				},
				SensitiveDataDiscovery: &security.SensitiveDataDiscoveryProperties{
					IsEnabled: to.Ptr(true),
				},
			},
		},
	}

	// Azure has no DELETE operation for these settings; detach without disabling protection.
	tc.AddAnnotation(&setting.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")
	tc.CreateResourcesAndWait(account, setting)

	tc.Expect(account.Status.Id).ToNot(BeNil())
	armID := *account.Status.Id + "/providers/Microsoft.Security/defenderForStorageSettings/current"
	tc.Expect(setting.Status.Id).To(Equal(to.Ptr(armID)))
	tc.Expect(setting.Status.Properties).ToNot(BeNil())
	tc.Expect(setting.Status.Properties.IsEnabled).To(Equal(to.Ptr(true)))
	tc.Expect(setting.Status.Properties.OverrideSubscriptionLevelSettings).To(Equal(to.Ptr(true)))
	tc.Expect(setting.Status.Properties.MalwareScanning).ToNot(BeNil())
	tc.Expect(setting.Status.Properties.MalwareScanning.OnUpload).ToNot(BeNil())
	tc.Expect(setting.Status.Properties.MalwareScanning.OnUpload.IsEnabled).To(Equal(to.Ptr(true)))
	tc.Expect(setting.Status.Properties.MalwareScanning.OnUpload.CapGBPerMonth).To(Equal(to.Ptr(1)))
	tc.Expect(setting.Status.Properties.SensitiveDataDiscovery).ToNot(BeNil())
	tc.Expect(setting.Status.Properties.SensitiveDataDiscovery.IsEnabled).To(Equal(to.Ptr(true)))

	old := setting.DeepCopy()
	setting.Spec.Properties.MalwareScanning.OnUpload.CapGBPerMonth = to.Ptr(2)
	tc.PatchResourceAndWait(old, setting)
	tc.Expect(setting.Status.Properties.MalwareScanning.OnUpload.CapGBPerMonth).To(Equal(to.Ptr(2)))

	tc.DeleteResourceAndWait(setting)
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armID, string(security.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())

	// A different Kubernetes object must adopt and update the same existing "current" settings.
	adopted := &security.DefenderForStorageSetting{
		ObjectMeta: tc.MakeObjectMeta("adopted"),
		Spec:       setting.Spec,
	}
	adopted.Spec.Properties.MalwareScanning.OnUpload.CapGBPerMonth = to.Ptr(1)
	tc.AddAnnotation(&adopted.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")
	tc.CreateResourceAndWait(adopted)
	tc.Expect(adopted.Status.Id).To(Equal(to.Ptr(armID)))
	tc.Expect(adopted.Status.Properties).ToNot(BeNil())
	tc.Expect(adopted.Status.Properties.MalwareScanning).ToNot(BeNil())
	tc.Expect(adopted.Status.Properties.MalwareScanning.OnUpload).ToNot(BeNil())
	tc.Expect(adopted.Status.Properties.MalwareScanning.OnUpload.CapGBPerMonth).To(Equal(to.Ptr(1)))
}
