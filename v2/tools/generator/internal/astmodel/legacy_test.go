/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestHybridMigrationReleaseForGroup_WhenRegistered_ReturnsRelease(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// authorization is registered as a Hybrid group with an explicit migration release
	release, ok := HybridMigrationReleaseForGroup("authorization")
	g.Expect(ok).To(BeTrue())
	g.Expect(release).To(Equal("v2.21.0"))
}

func TestHybridMigrationReleaseForGroup_WhenNotRegistered_ReturnsFalse(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// cdn is a Legacy group with no hybrid migration release registered
	_, ok := HybridMigrationReleaseForGroup("cdn")
	g.Expect(ok).To(BeFalse())

	// A group that doesn't exist at all
	_, ok = HybridMigrationReleaseForGroup("nonexistent")
	g.Expect(ok).To(BeFalse())
}

func TestHybridMigrationReleaseForGroup_IsCaseInsensitive(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	release, ok := HybridMigrationReleaseForGroup("Authorization")
	g.Expect(ok).To(BeTrue())
	g.Expect(release).To(Equal("v2.21.0"))
}
