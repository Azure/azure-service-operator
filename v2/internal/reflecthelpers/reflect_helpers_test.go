/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reflecthelpers_test

import (
	"context"
	"testing"

	//nolint:staticcheck // ignoring deprecation (SA1019) to unblock CI builds
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"

	// TODO: Do we want to use a sample object rather than a code generated one?
	batch "github.com/Azure/azure-service-operator/v2/api/microsoft.batch/v1alpha1api20210101"

	. "github.com/onsi/gomega"
)

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
