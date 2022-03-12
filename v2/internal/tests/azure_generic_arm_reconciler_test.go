/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package tests

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_ConvertResourceToARMResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ctx := context.Background()

	scheme, err := testcommon.CreateScheme()
	g.Expect(err).ToNot(HaveOccurred())

	testClient := testcommon.CreateClient(scheme)

	resolver, err := testcommon.CreateResolver(scheme, testClient)
	g.Expect(err).ToNot(HaveOccurred())

	rg := testcommon.CreateResourceGroup()
	g.Expect(testClient.Create(ctx, rg)).To(Succeed())

	account := testcommon.CreateDummyResource()
	g.Expect(testClient.Create(ctx, account)).To(Succeed())

	subscriptionID := "1234"

	resource, err := arm.ConvertToARMResourceImpl(ctx, account, scheme, resolver, subscriptionID)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect("azureName").To(Equal(resource.Spec().GetName()))
	g.Expect("2021-01-01").To(Equal(resource.Spec().GetAPIVersion()))
	g.Expect("Microsoft.Batch/batchAccounts").To(Equal(resource.Spec().GetType()))
}
