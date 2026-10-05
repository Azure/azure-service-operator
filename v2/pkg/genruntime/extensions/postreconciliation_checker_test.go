// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package extensions

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type recordingPostReconciliationChecker struct {
	received annotations.ResolvedReconcilePolicies
}

func (*recordingPostReconciliationChecker) GetExtendedResources() []genruntime.KubernetesResource {
	return nil
}

func (r *recordingPostReconciliationChecker) PostReconcileCheck(
	ctx context.Context,
	obj genruntime.MetaObject,
	owner genruntime.MetaObject,
	resourceResolver *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	policies annotations.ResolvedReconcilePolicies,
	next PostReconcileCheckFunc,
) (PostReconcileCheckResult, error) {
	r.received = policies
	return next(ctx, obj, owner, resourceResolver, armClient, log, policies)
}

func TestCreatePostReconciliationChecker_PassesPoliciesToExtension(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expected := annotations.ResolvedReconcilePolicies{
		Effective:       annotations.ReconcilePolicySkip,
		NamespacePolicy: annotations.ReconcilePolicyManage,
		Global:          annotations.ReconcilePolicyDetachOnDelete,
	}
	extension := &recordingPostReconciliationChecker{}

	checker, found := CreatePostReconciliationChecker(extension)
	g.Expect(found).To(BeTrue())

	result, err := checker(
		context.Background(),
		nil,
		nil,
		nil,
		nil,
		logr.Discard(),
		expected,
	)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(result.ReconciliationSucceeded()).To(BeTrue())
	g.Expect(extension.received).To(Equal(expected))
}
