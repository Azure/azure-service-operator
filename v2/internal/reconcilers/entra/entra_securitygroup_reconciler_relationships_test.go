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
	"github.com/google/uuid"
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
			[]uuid.UUID{uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"), uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb")},
			[]uuid.UUID{uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"), uuid.MustParse("cccccccc-cccc-cccc-cccc-cccccccccccc"), uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd")},
			func(_ context.Context, id uuid.UUID) error {
				calls = append(calls, "add:"+id.String())
				return nil
			},
			func(_ context.Context, id uuid.UUID) error {
				calls = append(calls, "remove:"+id.String())
				return nil
			},
			logr.Discard(),
		)

		g.Expect(err).ToNot(HaveOccurred())

		// The precise order of these operations is not guaranteed, but we can check that all the expected calls were made.
		g.Expect(calls).To(HaveLen(3))
		g.Expect(calls).To(ContainElement("add:cccccccc-cccc-cccc-cccc-cccccccccccc"))
		g.Expect(calls).To(ContainElement("add:dddddddd-dddd-dddd-dddd-dddddddddddd"))
		g.Expect(calls).To(ContainElement("remove:aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"))
	})

	t.Run("remove is skipped when add fails", func(t *testing.T) {
		t.Parallel()
		g := NewGomegaWithT(t)

		var calls []string
		reconciler := &EntraSecurityGroupReconciler{}
		err := reconciler.reconcileRelationship(
			context.Background(),
			"members",
			[]uuid.UUID{uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")},
			[]uuid.UUID{uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb")},
			func(_ context.Context, id uuid.UUID) error {
				calls = append(calls, "add:"+id.String())
				return errors.New("boom")
			},
			func(_ context.Context, id uuid.UUID) error {
				calls = append(calls, "remove:"+id.String())
				return nil
			},
			logr.Discard(),
		)

		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(MatchError(ContainSubstring("add bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb to members")))
		g.Expect(calls).To(Equal([]string{"add:bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"}))
	})
}
