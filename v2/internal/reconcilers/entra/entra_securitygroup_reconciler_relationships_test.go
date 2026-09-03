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
