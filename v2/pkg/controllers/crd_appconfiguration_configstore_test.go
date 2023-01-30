/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appconfig "github.com/Azure/azure-service-operator/v2/api/appconfiguration/v1beta20220501"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// If recording this test, might need to manually purge the old App config: az appconfig purge --name asotest-confstore-fsrajl

func Test_AppConfiguration_ConfigurationStore_CRUD(t *testing.T) {
	t.Parallel()

	createModeDefault := appconfig.ConfigurationStoreProperties_CreateMode_Default
	publicNetworkAccessDisabled := appconfig.ConfigurationStoreProperties_PublicNetworkAccess_Disabled
	publicNetworkAccessEnabled := appconfig.ConfigurationStoreProperties_PublicNetworkAccess_Enabled

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	cs := &appconfig.ConfigurationStore{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: tc.MakeObjectMeta("confstore"),
		Spec: appconfig.ConfigurationStore_Spec{
			CreateMode: &createModeDefault,
			Location:   tc.AzureRegion,
			Owner:      testcommon.AsOwner(rg),
			Sku: &appconfig.Sku{
				Name: to.StringPtr("standard"),
			},
			PublicNetworkAccess: &publicNetworkAccessDisabled,
		},
		Status: appconfig.ConfigurationStore_STATUS{},
	}

	tc.CreateResourceAndWait(cs)

	armId := *cs.Status.Id
	old := cs.DeepCopy()
	cs.Spec.PublicNetworkAccess = &publicNetworkAccessEnabled
	tc.PatchResourceAndWait(old, cs)
	tc.Expect(string(*cs.Status.PublicNetworkAccess)).To(gomega.Equal(string(publicNetworkAccessEnabled)))

	tc.RunSubtests(
		testcommon.Subtest{
			Name: "WriteWorkspacesSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				ConfigStore_WriteSecrets(tc, cs)
			},
		})

	tc.DeleteResourceAndWait(cs)

	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(appconfig.APIVersion_Value))
	tc.Expect(err).ToNot(gomega.HaveOccurred())
	tc.Expect(exists).To(gomega.BeFalse())
}

func ConfigStore_WriteSecrets(tc *testcommon.KubePerTestContext, cs *appconfig.ConfigurationStore) {
	old := cs.DeepCopy()
	csKeysSecret := "cskeyssecret"
	cs.Spec.OperatorSpec = &appconfig.ConfigurationStoreOperatorSpec{

		Secrets: &appconfig.ConfigurationStoreOperatorSecrets{
			PrimaryConnectionString:           &genruntime.SecretDestination{Name: csKeysSecret, Key: "primaryConnectionString"},
			SecondaryConnectionString:         &genruntime.SecretDestination{Name: csKeysSecret, Key: "secondaryConnectionString"},
			PrimaryReadOnlyConnectionString:   &genruntime.SecretDestination{Name: csKeysSecret, Key: "primaryReadOnlyConnectionString"},
			SecondaryReadOnlyConnectionString: &genruntime.SecretDestination{Name: csKeysSecret, Key: "secondaryReadOnlyConnectionString"},
		},
	}
	tc.PatchResourceAndWait(old, cs)

	tc.ExpectSecretHasKeys(
		csKeysSecret,
		"primaryConnectionString",
		"secondaryConnectionString",
		"primaryReadOnlyConnectionString",
		"secondaryReadOnlyConnectionString")
}
