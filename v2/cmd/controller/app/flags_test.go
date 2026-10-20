// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package app

import (
	"flag"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func TestLeaderElectionFlags(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	flags := InitFlags(flagSet)

	g.Expect(flags.Validate()).To(Succeed())

	g.Expect(flagSet.Parse([]string{"--leader-lease-duration=2m", "--leader-renew-deadline=100s", "--leader-retry-period=20s"})).To(Succeed())
	g.Expect(flags.LeaseDuration).To(Equal(2 * time.Minute))
	g.Expect(flags.RenewDeadline).To(Equal(100 * time.Second))
	g.Expect(flags.RetryPeriod).To(Equal(20 * time.Second))
	g.Expect(flags.Validate()).To(Succeed())
}

func TestLeaderElectionFlagsValidation(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		args          []string
		expectedError string
	}{
		"renew deadline at lease duration": {
			args:          []string{"--leader-lease-duration=10s", "--leader-renew-deadline=10s"},
			expectedError: "leader-lease-duration (10s) must be greater than leader-renew-deadline (10s)",
		},
		"renew deadline beyond lease duration": {
			args:          []string{"--leader-lease-duration=10s", "--leader-renew-deadline=20s"},
			expectedError: "leader-lease-duration (10s) must be greater than leader-renew-deadline (20s)",
		},
		"zero lease duration": {
			args:          []string{"--leader-lease-duration=0s"},
			expectedError: "leader-lease-duration (0s) must be greater than zero",
		},
		"negative lease duration": {
			args:          []string{"--leader-lease-duration=-1s"},
			expectedError: "leader-lease-duration (-1s) must be greater than zero",
		},
		"zero renew deadline": {
			args:          []string{"--leader-renew-deadline=0s"},
			expectedError: "leader-renew-deadline (0s) must be greater than zero",
		},
		"negative renew deadline": {
			args:          []string{"--leader-renew-deadline=-1s"},
			expectedError: "leader-renew-deadline (-1s) must be greater than zero",
		},
		"zero retry period": {
			args:          []string{"--leader-retry-period=0s"},
			expectedError: "leader-retry-period (0s) must be greater than zero",
		},
		"negative retry period": {
			args:          []string{"--leader-retry-period=-1s"},
			expectedError: "leader-retry-period (-1s) must be greater than zero",
		},
		"renew deadline at jitter boundary": {
			args: []string{
				"--leader-lease-duration=15s",
				"--leader-renew-deadline=10800ms",
				"--leader-retry-period=9s",
			},
			expectedError: "leader-renew-deadline (10.8s) must be greater than leader-retry-period (9s) multiplied by the jitter factor (1.2)",
		},
		"renew deadline just above jitter boundary": {
			args: []string{
				"--leader-lease-duration=15s",
				"--leader-renew-deadline=10801ms",
				"--leader-retry-period=9s",
			},
		},
		"large retry period does not overflow jitter validation": {
			args:          []string{"--leader-retry-period=2200000h"},
			expectedError: "must be greater than leader-retry-period",
		},
		"jitter overflow is rejected when retry is below renew deadline": {
			args: []string{
				"--leader-lease-duration=2560000h",
				"--leader-renew-deadline=2500000h",
				"--leader-retry-period=2200000h",
			},
			expectedError: "multiplied by the jitter factor (1.2)",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)
			flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
			flags := InitFlags(flagSet)

			g.Expect(flagSet.Parse(c.args)).To(Succeed())
			err := flags.Validate()
			if c.expectedError == "" {
				g.Expect(err).To(Succeed())
			} else {
				g.Expect(err).To(MatchError(ContainSubstring(c.expectedError)))
			}
		})
	}
}

func TestEnableAzureCLIAuthenticationFlag(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	flags := InitFlags(flagSet)

	g.Expect(flags.EnableAzureCLIAuthentication).To(BeFalse())
	g.Expect(flagSet.Parse([]string{"--enable-azure-cli-authentication"})).To(Succeed())
	g.Expect(flags.EnableAzureCLIAuthentication).To(BeTrue())
}

func TestCRDLabelsFlag(t *testing.T) {
	t.Parallel()
	g := NewWithT(t)
	flagSet := flag.NewFlagSet("test", flag.ContinueOnError)
	flags := InitFlags(flagSet)

	g.Expect(flagSet.Parse([]string{"--crd-labels=cluster.x-k8s.io/provider=infrastructure-azure,example.com/owner=aso"})).To(Succeed())
	g.Expect(flags.CRDLabels).To(Equal("cluster.x-k8s.io/provider=infrastructure-azure,example.com/owner=aso"))

	parsed, err := parseCRDLabels(flags.CRDLabels)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(parsed).To(HaveLen(2))
	g.Expect(parsed).To(HaveKeyWithValue("cluster.x-k8s.io/provider", "infrastructure-azure"))
	g.Expect(parsed).To(HaveKeyWithValue("example.com/owner", "aso"))
}

// TestParseCRDLabels covers the CRD specific behaviour layered on top of labels.ParseMap.
// General label parsing and validation is covered by TestParseMap in the labels package.
func TestParseCRDLabels(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		value         string
		expected      map[string]string
		expectedError string
	}{
		"empty": {
			value:    "",
			expected: map[string]string{},
		},
		"reserved app label": {
			value:         "app.kubernetes.io/name=mine",
			expectedError: `label "app.kubernetes.io/name" is reserved`,
		},
		"reserved version label": {
			value:         "app.kubernetes.io/version=v1.0.0",
			expectedError: `label "app.kubernetes.io/version" is reserved`,
		},
		"reserved old version label": {
			value:         "serviceoperator.azure.com/version=v1.0.0",
			expectedError: `label "serviceoperator.azure.com/version" is reserved`,
		},
		"reserved prefix": {
			value:         "serviceoperator.azure.com/anything=value",
			expectedError: `label "serviceoperator.azure.com/anything" is reserved`,
		},
		"reserved label alongside a valid one": {
			value:         "example.com/owner=aso,app.kubernetes.io/name=mine",
			expectedError: `label "app.kubernetes.io/name" is reserved`,
		},
		"reserved label with surrounding whitespace": {
			value:         " app.kubernetes.io/version = v1.0.0 ",
			expectedError: `label "app.kubernetes.io/version" is reserved`,
		},
		"unreserved app.kubernetes.io label": {
			value:    "app.kubernetes.io/part-of=platform",
			expected: map[string]string{"app.kubernetes.io/part-of": "platform"},
		},
		"valid labels are passed through": {
			value:    "environment=production;example.com/owner=aso",
			expected: map[string]string{"environment": "production", "example.com/owner": "aso"},
		},
		"parse errors are propagated": {
			value:         "not a label=value",
			expectedError: `invalid label key "not a label"`,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)

			actual, err := parseCRDLabels(c.value)

			if c.expectedError != "" {
				g.Expect(err).To(MatchError(ContainSubstring(c.expectedError)))
				return
			}

			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(actual).To(HaveLen(len(c.expected)))
			for key, value := range c.expected {
				g.Expect(actual).To(HaveKeyWithValue(key, value))
			}
		})
	}
}
