// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/pkg/common/config"
)

func Test_ParseSyncPeriod_ReturnsNever(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Setenv(config.SyncPeriod, "never") // Can't run in parallel

	dur, err := parseSyncPeriod()

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(dur).To(BeNil()) // Nil means no sync
}

func Test_ParseSyncPeriod_ReturnsDefaultWhenEmpty(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Setenv(config.SyncPeriod, "") // Can't run in parallel

	dur, err := parseSyncPeriod()

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(dur).ToNot(BeNil())
	g.Expect(*dur).To(Equal(1 * time.Hour))
}

func Test_ParseSyncPeriod_ReturnsValue(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Setenv(config.SyncPeriod, "21m") // Can't run in parallel

	dur, err := parseSyncPeriod()

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(dur).ToNot(BeNil())
	g.Expect(*dur).To(Equal(21 * time.Minute))
}
