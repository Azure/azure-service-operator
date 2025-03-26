/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package tests

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	dbforpostgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20230601preview"
	dbforpostgresqlarm "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20230601preview/arm"
	dbforpostgresqlstorage "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20240801/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type testData struct {
	scheme         *runtime.Scheme
	client         client.Client
	resolver       *resolver.Resolver
	subscriptionID string
}

func testSetup(g *WithT) *testData {
	scheme, err := testcommon.CreateScheme()
	g.Expect(err).ToNot(HaveOccurred())

	testClient := testcommon.CreateClient(scheme)

	resolver, err := testcommon.CreateResolver(scheme, testClient)
	g.Expect(err).ToNot(HaveOccurred())

	subscriptionID := "00000000-0000-0000-0000-000000000000"

	return &testData{
		scheme:         scheme,
		client:         testClient,
		resolver:       resolver,
		subscriptionID: subscriptionID,
	}
}

func Test_ConvertResourceToARMResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()

	testData := testSetup(g)

	rg := testcommon.CreateResourceGroup()
	g.Expect(testData.client.Create(ctx, rg)).To(Succeed())

	account := testcommon.CreateDummyResource()
	g.Expect(testData.client.Create(ctx, account)).To(Succeed())

	resource, err := arm.ConvertToARMResourceImpl(ctx, account, testData.resolver, testData.subscriptionID)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(resource.Spec().GetName()).To(Equal("azureName"))
	g.Expect(resource.Spec().GetAPIVersion()).To(Equal("2021-01-01"))
	g.Expect(resource.Spec().GetType()).To(Equal("Microsoft.Batch/batchAccounts"))
}

func Test_Conversion_DiscoversConfigMapsNotOnStorageVersion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()

	// ensure that the storage type is the hub type
	var _ conversion.Hub = &dbforpostgresqlstorage.FlexibleServer{}

	testData := testSetup(g)

	rg := testcommon.CreateResourceGroup()
	g.Expect(testData.client.Create(ctx, rg)).To(Succeed())

	configMap := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "myconfig",
			Namespace: rg.Namespace,
		},
		Data: map[string]string{
			"breakfast": "bagel",
		},
	}
	g.Expect(testData.client.Create(ctx, configMap)).To(Succeed())

	db2 := &dbforpostgresql.FlexibleServer{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: rg.Namespace,
			Name:      "mydb",
		},
		Spec: dbforpostgresql.FlexibleServer_Spec{
			AzureName: "mydb",
			Location:  to.Ptr("westus"),
			Owner: &genruntime.KnownResourceReference{
				Name: "myrg",
			},
			DataEncryption: &dbforpostgresql.DataEncryption{
				GeoBackupKeyURIFromConfig: &genruntime.ConfigMapReference{
					Name: "myconfig",
					Key:  "breakfast",
				},
			},
		},
	}

	storageVersion, err := genruntime.ObjAsVersion(
		db2,
		testData.scheme,
		schema.GroupVersionKind{
			Group:   dbforpostgresqlstorage.GroupVersion.Group,
			Version: dbforpostgresqlstorage.GroupVersion.Version,
			Kind:    "FlexibleServer",
		})
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(genruntime.GetOriginalGVK(storageVersion)).To(Equal(schema.GroupVersionKind{
		Group:   dbforpostgresql.GroupVersion.Group,
		Version: dbforpostgresql.GroupVersion.Version,
		Kind:    "FlexibleServer",
	}))

	// Convert db to the storage type to simulate what would happen during reconciliation
	resource, err := arm.ConvertToARMResourceImpl(ctx, storageVersion, testData.resolver, testData.subscriptionID)
	g.Expect(err).ToNot(HaveOccurred())

	armType, ok := resource.Spec().(*dbforpostgresqlarm.FlexibleServer_Spec)
	g.Expect(ok).To(BeTrue(), "Expected resource.Spec() to be of type FlexibleServer_Spec_ARM")

	g.Expect(resource.Spec().GetName()).To(Equal("mydb"))
	g.Expect(resource.Spec().GetAPIVersion()).To(Equal("2023-06-01-preview"))
	g.Expect(resource.Spec().GetType()).To(Equal("Microsoft.DBforPostgreSQL/flexibleServers"))

	// Ensure that the property that exists only in the preview API but not in the storage version was found
	g.Expect(armType.Properties.DataEncryption.GeoBackupKeyURI).To(Equal(to.Ptr("bagel")))
}
