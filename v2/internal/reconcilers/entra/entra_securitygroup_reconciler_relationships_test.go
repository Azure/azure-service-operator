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

	reconcileBothSides := func(
		ctx context.Context,
		reconciler *EntraSecurityGroupReconciler,
		ownersCurrent []string,
		ownersDesired []string,
		ownersAdd func(context.Context, string) error,
		ownersRemove func(context.Context, string) error,
		membersCurrent []string,
		membersDesired []string,
		membersAdd func(context.Context, string) error,
		membersRemove func(context.Context, string) error,
	) error {
		var joinedErr error

		ownersErr := reconciler.reconcileRelationshipSide(
			ctx,
			"owners",
			ownersCurrent,
			ownersDesired,
			ownersAdd,
			ownersRemove,
			logr.Discard(),
		)
		joinedErr = errors.Join(joinedErr, ownersErr)

		membersErr := reconciler.reconcileRelationshipSide(
			ctx,
			"members",
			membersCurrent,
			membersDesired,
			membersAdd,
			membersRemove,
			logr.Discard(),
		)
		joinedErr = errors.Join(joinedErr, membersErr)

		return joinedErr
	}

	t.Run("owners fail and members still run", func(t *testing.T) {
		t.Parallel()
		g := NewGomegaWithT(t)

		reconciler := &EntraSecurityGroupReconciler{}
		var membersCalls []string

		err := reconcileBothSides(
			context.Background(),
			reconciler,
			nil,
			[]string{"owner-a"},
			func(_ context.Context, _ string) error {
				return errors.New("owners list failed")
			},
			func(_ context.Context, _ string) error { return nil },
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
		)

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("owners"))
		g.Expect(membersCalls).To(Equal([]string{"add:member-b", "remove:member-a"}))
	})

	t.Run("members fail and owners still run", func(t *testing.T) {
		t.Parallel()
		g := NewGomegaWithT(t)

		reconciler := &EntraSecurityGroupReconciler{}
		var ownersCalls []string

		err := reconcileBothSides(
			context.Background(),
			reconciler,
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
			nil,
			[]string{"member-a"},
			func(_ context.Context, _ string) error {
				return errors.New("members list failed")
			},
			func(_ context.Context, _ string) error { return nil },
		)

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("members"))
		g.Expect(ownersCalls).To(Equal([]string{"add:owner-b", "remove:owner-a"}))
	})

	t.Run("both sides fail and aggregate both contexts", func(t *testing.T) {
		t.Parallel()
		g := NewGomegaWithT(t)

		reconciler := &EntraSecurityGroupReconciler{}

		err := reconcileBothSides(
			context.Background(),
			reconciler,
			nil,
			[]string{"owner-a"},
			func(_ context.Context, _ string) error {
				return errors.New("owners side failed")
			},
			func(_ context.Context, _ string) error { return nil },
			nil,
			[]string{"member-a"},
			func(_ context.Context, _ string) error {
				return errors.New("members side failed")
			},
			func(_ context.Context, _ string) error { return nil },
		)

		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("owners"))
		g.Expect(err.Error()).To(ContainSubstring("members"))
	})
}
