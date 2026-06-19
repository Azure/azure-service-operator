/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

import (
	"context"
	"errors"
	"testing"

	"github.com/go-logr/logr"
	. "github.com/onsi/gomega"
)

func TestReconcileRelationshipSide_AddBeforeRemove_AndSkipRemoveWhenAddFails(t *testing.T) {
	t.Parallel()

	t.Run("adds happen before removes", func(t *testing.T) {
		t.Parallel()
		g := NewGomegaWithT(t)

		var calls []string
		reconciler := &EntraSecurityGroupReconciler{}
		err := reconciler.reconcileRelationshipSide(
			context.Background(),
			"owners",
			[]string{"owner-a", "owner-b"},
			[]string{"owner-b", "owner-c", "owner-d"},
			func(_ context.Context, id string) error {
				calls = append(calls, "add:"+id)
				return nil
			},
			func(_ context.Context, id string) error {
				calls = append(calls, "remove:"+id)
				return nil
			},
			logr.Discard(),
		)

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(calls).To(Equal([]string{
			"add:owner-c",
			"add:owner-d",
			"remove:owner-a",
		}))
	})

	t.Run("remove is skipped when add fails", func(t *testing.T) {
		t.Parallel()
		g := NewGomegaWithT(t)

		var calls []string
		reconciler := &EntraSecurityGroupReconciler{}
		err := reconciler.reconcileRelationshipSide(
			context.Background(),
			"members",
			[]string{"member-a"},
			[]string{"member-b"},
			func(_ context.Context, id string) error {
				calls = append(calls, "add:"+id)
				return errors.New("boom")
			},
			func(_ context.Context, id string) error {
				calls = append(calls, "remove:"+id)
				return nil
			},
			logr.Discard(),
		)

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("members add member-b"))
		g.Expect(calls).To(Equal([]string{"add:member-b"}))
	})
}

func TestReconcileRelationshipSide_CrossSideBestEffort(t *testing.T) {
	t.Parallel()

	t.Run("owners fail and members still run", testCrossSideOwnersFailMembersStillRun)
	t.Run("members fail and owners still run", testCrossSideMembersFailOwnersStillRun)
	t.Run("both sides fail and aggregate both contexts", testCrossSideBothSidesFail)
}

func testCrossSideOwnersFailMembersStillRun(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	reconciler := &EntraSecurityGroupReconciler{}
	ctx := context.Background()
	var membersCalls []string

	ownersErr := reconciler.reconcileRelationshipSide(
		ctx, "owners",
		nil, []string{"owner-a"},
		func(_ context.Context, _ string) error { return errors.New("owners list failed") },
		func(_ context.Context, _ string) error { return nil },
		logr.Discard(),
	)
	membersErr := reconciler.reconcileRelationshipSide(
		ctx, "members",
		[]string{"member-a"}, []string{"member-b"},
		func(_ context.Context, id string) error {
			membersCalls = append(membersCalls, "add:"+id)
			return nil
		},
		func(_ context.Context, id string) error {
			membersCalls = append(membersCalls, "remove:"+id)
			return nil
		},
		logr.Discard(),
	)
	err := errors.Join(ownersErr, membersErr)

	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("owners"))
	g.Expect(membersCalls).To(Equal([]string{"add:member-b", "remove:member-a"}))
}

func testCrossSideMembersFailOwnersStillRun(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	reconciler := &EntraSecurityGroupReconciler{}
	ctx := context.Background()
	var ownersCalls []string

	ownersErr := reconciler.reconcileRelationshipSide(
		ctx, "owners",
		[]string{"owner-a"}, []string{"owner-b"},
		func(_ context.Context, id string) error {
			ownersCalls = append(ownersCalls, "add:"+id)
			return nil
		},
		func(_ context.Context, id string) error {
			ownersCalls = append(ownersCalls, "remove:"+id)
			return nil
		},
		logr.Discard(),
	)
	membersErr := reconciler.reconcileRelationshipSide(
		ctx, "members",
		nil, []string{"member-a"},
		func(_ context.Context, _ string) error { return errors.New("members list failed") },
		func(_ context.Context, _ string) error { return nil },
		logr.Discard(),
	)
	err := errors.Join(ownersErr, membersErr)

	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("members"))
	g.Expect(ownersCalls).To(Equal([]string{"add:owner-b", "remove:owner-a"}))
}

func testCrossSideBothSidesFail(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	reconciler := &EntraSecurityGroupReconciler{}
	ctx := context.Background()

	ownersErr := reconciler.reconcileRelationshipSide(
		ctx, "owners",
		nil, []string{"owner-a"},
		func(_ context.Context, _ string) error { return errors.New("owners side failed") },
		func(_ context.Context, _ string) error { return nil },
		logr.Discard(),
	)
	membersErr := reconciler.reconcileRelationshipSide(
		ctx, "members",
		nil, []string{"member-a"},
		func(_ context.Context, _ string) error { return errors.New("members side failed") },
		func(_ context.Context, _ string) error { return nil },
		logr.Discard(),
	)
	err := errors.Join(ownersErr, membersErr)

	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("owners"))
	g.Expect(err.Error()).To(ContainSubstring("members"))
}
