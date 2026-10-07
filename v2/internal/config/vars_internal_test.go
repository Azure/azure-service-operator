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

func Test_AllowMultiEnvManagement_DefaultsToFalse(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Setenv(config.AllowMultiEnvManagement, "") // Can't run in parallel

	cfg, err := ReadFromEnvironment()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(cfg.AllowMultiEnvManagement).To(BeFalse())
}

func Test_AllowMultiEnvManagement_ReadsTrue(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Setenv(config.AllowMultiEnvManagement, "true") // Can't run in parallel

	cfg, err := ReadFromEnvironment()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(cfg.AllowMultiEnvManagement).To(BeTrue())
}

func Test_AllowMultiEnvManagement_ReadsFalse(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Setenv(config.AllowMultiEnvManagement, "false") // Can't run in parallel

	cfg, err := ReadFromEnvironment()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(cfg.AllowMultiEnvManagement).To(BeFalse())
}

func Test_AllowMultiEnvManagement_IncludedInString(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := Values{
		AllowMultiEnvManagement: true,
	}
	s := cfg.String()
	g.Expect(s).To(ContainSubstring("AllowMultiEnvManagement:true"))

	cfg.AllowMultiEnvManagement = false
	s = cfg.String()
	g.Expect(s).To(ContainSubstring("AllowMultiEnvManagement:false"))
}

func Test_FederatedTokenFilePath_DefaultsToEmpty(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Setenv(config.AzureFederatedTokenFile, "") // Can't run in parallel

	cfg, err := ReadFromEnvironment()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(cfg.FederatedTokenFilePath).To(BeEmpty())
}

func Test_FederatedTokenFilePath_ReadsValue(t *testing.T) {
	g := NewGomegaWithT(t)
	const customTokenPath = "/var/run/secrets/azure/tokens/azure-identity-token" // #nosec G101 -- file path, not a credential
	t.Setenv(config.AzureFederatedTokenFile, "  "+customTokenPath+"  ")          // Can't run in parallel

	cfg, err := ReadFromEnvironment()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(cfg.FederatedTokenFilePath).To(Equal(customTokenPath)) // whitespace-trimmed
}

func Test_FederatedTokenFilePath_IncludedInString(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := Values{
		FederatedTokenFilePath: "/some/path",
	}
	s := cfg.String()
	g.Expect(s).To(ContainSubstring("FederatedTokenFilePath:/some/path"))
}

func Test_WorkloadIdentityAuthMode_DefaultsToRelaxed(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Setenv(config.AzureWorkloadIdentityAuthMode, "") // Can't run in parallel

	cfg, err := ReadFromEnvironment()
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(cfg.WorkloadIdentityAuthMode).To(Equal(WorkloadIdentityAuthModeRelaxed))
}

func Test_WorkloadIdentityAuthMode_ReadsSupportedValues(t *testing.T) {
	tests := []struct {
		value    string
		expected WorkloadIdentityAuthMode
	}{
		{value: "relaxed", expected: WorkloadIdentityAuthModeRelaxed},
		{value: "strict", expected: WorkloadIdentityAuthModeStrict},
	}

	for _, test := range tests {
		t.Run(test.value, func(t *testing.T) {
			g := NewGomegaWithT(t)
			t.Setenv(config.AzureWorkloadIdentityAuthMode, test.value) // Can't run in parallel

			cfg, err := ReadFromEnvironment()
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(cfg.WorkloadIdentityAuthMode).To(Equal(test.expected))
		})
	}
}

func Test_WorkloadIdentityAuthMode_RejectsUnsupportedValue(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Setenv(config.AzureWorkloadIdentityAuthMode, "invalid") // Can't run in parallel

	_, err := ReadFromEnvironment()
	g.Expect(err).To(MatchError(ContainSubstring("invalid workload identity auth mode")))
}

func Test_WorkloadIdentityAuthMode_IsIncludedInString(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := Values{WorkloadIdentityAuthMode: WorkloadIdentityAuthModeStrict}
	g.Expect(cfg.String()).To(ContainSubstring("WorkloadIdentityAuthMode:strict"))
}
