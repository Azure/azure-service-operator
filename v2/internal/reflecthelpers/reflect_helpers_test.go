/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reflecthelpers_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"

	//nolint:staticcheck // ignoring deprecation (SA1019) to unblock CI builds
	resources "github.com/Azure/azure-service-operator/v2/api/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"

	// TODO: Do we want to use a sample object rather than a code generated one?
	batch "github.com/Azure/azure-service-operator/v2/api/microsoft.batch/v1alpha1api20210101"

	. "github.com/onsi/gomega"
)

func createDummyResource() *batch.BatchAccount {
	return &batch.BatchAccount{
		Spec: batch.BatchAccounts_Spec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}
}

type DummyStruct struct{}

func Test_ConvertResourceToARMResource(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := context.Background()

	scheme, err := testcommon.CreateScheme()
	g.Expect(err).ToNot(HaveOccurred())

	testClient := testcommon.CreateClient(scheme)

	subscriptionID := "1234"
	resolver, err := testcommon.CreateResolver(scheme, testClient)
	g.Expect(err).ToNot(HaveOccurred())

	rg := testcommon.CreateResourceGroup()
	g.Expect(testClient.Create(ctx, rg)).To(Succeed())
	account := createDummyResource()
	g.Expect(testClient.Create(ctx, account)).To(Succeed())

	resource, err := reflecthelpers.ConvertResourceToARMResource(ctx, resolver, account, subscriptionID)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect("azureName").To(Equal(resource.Spec().GetName()))
	g.Expect("2021-01-01").To(Equal(resource.Spec().GetAPIVersion()))
	g.Expect("Microsoft.Batch/batchAccounts").To(Equal(resource.Spec().GetType()))
}

func Test_FindReferences(t *testing.T) {
	g := NewGomegaWithT(t)
	ctx := context.Background()

	scheme, err := testcommon.CreateScheme()
	g.Expect(err).ToNot(HaveOccurred())

	testClient := testcommon.CreateClient(scheme)

	rg := testcommon.CreateResourceGroup()
	g.Expect(testClient.Create(ctx, rg)).To(Succeed())

	account := testcommon.CreateDummyResource()
	ref := genruntime.ResourceReference{ARMID: "test"}
	account.Spec.KeyVaultReference = &batch.KeyVaultReference{
		Reference: ref,
	}
	g.Expect(testClient.Create(ctx, account)).To(Succeed())

	refs, err := reflecthelpers.FindResourceReferences(account)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(refs).To(HaveLen(1))
	g.Expect(refs).To(HaveKey(ref))
}

func Test_NewStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	account := testcommon.CreateDummyResource()

	status, err := reflecthelpers.NewEmptyStatus(account)
	g.Expect(err).To(BeNil())
	g.Expect(status).To(BeAssignableToTypeOf(&batch.BatchAccount_Status{}))
}

func Test_EmptyArmResourceStatus(t *testing.T) {
	g := NewGomegaWithT(t)

	account := testcommon.CreateDummyResource()

	status, err := reflecthelpers.NewEmptyArmResourceStatus(account)
	g.Expect(err).To(BeNil())
	g.Expect(status).To(BeAssignableToTypeOf(&batch.BatchAccount_StatusARM{}))
}
