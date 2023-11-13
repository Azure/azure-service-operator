/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	kubernetesconfiguration "github.com/Azure/azure-service-operator/v2/api/kubernetesconfiguration/v1api20230501"

	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_KubernetesConfiguration_Extension_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	adminUsername := "adminUser"
	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

	cluster := NewManagedCluster20230202preview(tc, rg, adminUsername, sshPublicKey)

	tc.CreateResourceAndWait(cluster)

	extension := &kubernetesconfiguration.Extension{
		ObjectMeta: tc.MakeObjectMeta("extension"),
		Spec: kubernetesconfiguration.Extension_Spec{
			AutoUpgradeMinorVersion: to.Ptr(true),
			ExtensionType:           to.Ptr("microsoft.flux"),
			Identity: &kubernetesconfiguration.Identity{
				Type: to.Ptr(kubernetesconfiguration.Identity_Type_SystemAssigned),
			},
			Owner: tc.AsExtensionOwner(cluster),
			Scope: &kubernetesconfiguration.Scope{
				Cluster: &kubernetesconfiguration.ScopeCluster{
					ReleaseNamespace: to.Ptr("kube-system"),
				},
			},
		},
	}

	tc.CreateResourceAndWait(extension)

	tc.Expect(extension.Status.Id).ToNot(BeNil())
	armId := *extension.Status.Id

	tc.DeleteResourceAndWait(extension)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(kubernetesconfiguration.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
