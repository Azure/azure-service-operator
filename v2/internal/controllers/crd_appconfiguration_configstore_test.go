/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/api/appconfiguration/v1beta20220501"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// If recording this test, might need to manually purge the old App config: az appconfig purge --name asotest-confstore-fsrajl

func Test_AppConfiguration_ConfigurationStore_CRUD(t *testing.T) {
	t.Parallel()

	createmodeDefault := v1beta20220501.ConfigurationStoreProperties_CreateMode_Default
	publicNetworkAccessDisabled := v1beta20220501.ConfigurationStoreProperties_PublicNetworkAccess_Disabled
	publicNetworkAccessEnabled := v1beta20220501.ConfigurationStoreProperties_PublicNetworkAccess_Enabled

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	cs := &v1beta20220501.ConfigurationStore{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: tc.MakeObjectMeta("confstore"),
		Spec: v1beta20220501.ConfigurationStores_Spec{
			CreateMode: &createmodeDefault,
			Location:   tc.AzureRegion,
			Owner:      testcommon.AsOwner(rg),
			Sku: &v1beta20220501.Sku{
				Name: to.StringPtr("standard"),
			},
			PublicNetworkAccess: &publicNetworkAccessDisabled,
		},
		Status: v1beta20220501.ConfigurationStore_STATUS{},
	}

	tc.CreateResourceAndWait(cs)

	tc.ExportAsSample(cs)

	armId := *cs.Status.Id
	old := cs.DeepCopy()
	cs.Spec.PublicNetworkAccess = &publicNetworkAccessEnabled
	tc.PatchResourceAndWait(old, cs)
	tc.Expect(string(*cs.Status.PublicNetworkAccess)).To(gomega.Equal(string(publicNetworkAccessEnabled)))

	tc.DeleteResourceAndWait(cs)

	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(v1beta20220501.APIVersion_Value))
	tc.Expect(err).ToNot(gomega.HaveOccurred())
	tc.Expect(exists).To(gomega.BeFalse())
}
