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

func Test_KubernetesConfiguration_FluxConfiguration_20230501_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	adminUsername := "adminUser"
	sshPublicKey, err := tc.GenerateSSHKey(2048)
	tc.Expect(err).ToNot(HaveOccurred())

	cluster := NewManagedCluster20240402preview(tc, rg, adminUsername, sshPublicKey)

	tc.CreateResourceAndWait(cluster)

	// We need flux extension before adding flux configurations
	fluxExtension := &kubernetesconfiguration.Extension{
		ObjectMeta: tc.MakeObjectMeta("extension"),
		Spec: kubernetesconfiguration.Extension_Spec{
			AutoUpgradeMinorVersion: to.Ptr(true),
			ExtensionType:           to.Ptr("microsoft.flux"),
			Owner:                   tc.AsExtensionOwner(cluster),
		},
	}

	tc.CreateResourceAndWait(fluxExtension)

	// Example from https://learn.microsoft.com/en-us/rest/api/kubernetesconfiguration/flux-configurations/create-or-update?view=rest-kubernetesconfiguration-2023-05-01
	flux := &kubernetesconfiguration.FluxConfiguration{
		ObjectMeta: tc.MakeObjectMeta("flux"),
		Spec: kubernetesconfiguration.FluxConfiguration_Spec{
			Owner:      tc.AsExtensionOwner(cluster),
			SourceKind: to.Ptr(kubernetesconfiguration.SourceKindDefinition_GitRepository),
			GitRepository: &kubernetesconfiguration.GitRepositoryDefinition{
				RepositoryRef: &kubernetesconfiguration.RepositoryRefDefinition{
					Branch: to.Ptr("master"),
				},
				Url: to.Ptr("https://github.com/Azure/arc-k8s-demo"),
			},
			Kustomizations: map[string]kubernetesconfiguration.KustomizationDefinition{
				"kustomization-1": {},
			},
			Namespace: to.Ptr("flux-ns"),
		},
	}

	tc.CreateResourceAndWait(flux)

	tc.Expect(flux.Status.Id).ToNot(BeNil())
	armId := *flux.Status.Id

	tc.DeleteResourceAndWait(flux)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(kubernetesconfiguration.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
