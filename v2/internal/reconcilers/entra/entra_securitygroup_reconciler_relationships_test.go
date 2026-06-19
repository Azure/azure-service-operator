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
