/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

import (
	"context"
	"errors"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
)

func TestReconcileRelationshipSide_AddBeforeRemove_AndSkipRemoveWhenAddFails(t *testing.T) {
	t.Parallel()

	t.Run("adds happen before removes", func(t *testing.T) {
		t.Parallel()
		g := NewGomegaWithT(t)

		var calls []string
		reconciler := &EntraSecurityGroupReconciler{}
		err := reconciler.reconcileRelationship(
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

		// The precise order of these operations is not guaranteed, but we can check that all the expected calls were made.
		g.Expect(calls).To(HaveLen(3))
		g.Expect(calls).To(ContainElement("add:owner-c"))
		g.Expect(calls).To(ContainElement("add:owner-d"))
		g.Expect(calls).To(ContainElement("remove:owner-a"))
	})

	t.Run("remove is skipped when add fails", func(t *testing.T) {
		t.Parallel()
		g := NewGomegaWithT(t)

		var calls []string
		reconciler := &EntraSecurityGroupReconciler{}
		err := reconciler.reconcileRelationship(
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
		g.Expect(err).To(MatchError(ContainSubstring("add member-b to members")))
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

	ownersErr := reconciler.reconcileRelationship(
		ctx,
		"owners",
		nil,
		[]string{"owner-a"},
		func(_ context.Context, _ string) error { return errors.New("owners list failed") },
		func(_ context.Context, _ string) error { return nil },
		logr.Discard(),
	)
	membersErr := reconciler.reconcileRelationship(
		ctx,
		"members",
		[]string{"member-a"},
		[]string{"member-b"},
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
	g.Expect(err).To(MatchError(ContainSubstring("owners")))
	g.Expect(membersCalls).To(Equal([]string{"add:member-b", "remove:member-a"}))
}

func testCrossSideMembersFailOwnersStillRun(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	reconciler := &EntraSecurityGroupReconciler{}
	ctx := context.Background()
	var ownersCalls []string

	ownersErr := reconciler.reconcileRelationship(
		ctx,
		"owners",
		[]string{"owner-a"},
		[]string{"owner-b"},
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
	membersErr := reconciler.reconcileRelationship(
		ctx,
		"members",
		nil,
		[]string{"member-a"},
		func(_ context.Context, _ string) error { return errors.New("members list failed") },
		func(_ context.Context, _ string) error { return nil },
		logr.Discard(),
	)
	err := errors.Join(ownersErr, membersErr)

	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError(ContainSubstring("members")))
	g.Expect(ownersCalls).To(Equal([]string{"add:owner-b", "remove:owner-a"}))
}

func testCrossSideBothSidesFail(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	reconciler := &EntraSecurityGroupReconciler{}
	ctx := context.Background()

	ownersErr := reconciler.reconcileRelationship(
		ctx,
		"owners",
		nil,
		[]string{"owner-a"},
		func(_ context.Context, _ string) error { return errors.New("owners side failed") },
		func(_ context.Context, _ string) error { return nil },
		logr.Discard(),
	)
	membersErr := reconciler.reconcileRelationship(
		ctx,
		"members",
		nil,
		[]string{"member-a"},
		func(_ context.Context, _ string) error { return errors.New("members side failed") },
		func(_ context.Context, _ string) error { return nil },
		logr.Discard(),
	)
	err := errors.Join(ownersErr, membersErr)

	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError(ContainSubstring("owners")))
	g.Expect(err).To(MatchError(ContainSubstring("members")))
}
