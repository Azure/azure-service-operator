// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config_test

import (
	"crypto/tls"
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	// Importing this for side effects is required as it initializes cloud.AzurePublic
	_ "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"

	"github.com/Azure/azure-service-operator/v2/internal/config"
)

func Test_String_HasAllKeys(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// This makes sure that if config.Values or testConfig changes
	// then cfgToKey is updated to match.
	s := config.Values{}.String()

	testConfigType := reflect.TypeOf(config.Values{})

	for _, field := range reflect.VisibleFields(testConfigType) {
		g.Expect(s).To(ContainSubstring(field.Name + ":"))
	}
}

func Test_Cloud_DefaultsToAzure(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := config.Values{}
	cld := cfg.Cloud()
	g.Expect(cld).To(Equal(cloud.AzurePublic))
}

func Test_Cloud_AuthorityHostSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := config.Values{
		AzureAuthorityHost: "host",
	}
	cld := cfg.Cloud()
	g.Expect(cld.ActiveDirectoryAuthorityHost).To(Equal(cfg.AzureAuthorityHost))
	g.Expect(cld.Services).To(HaveLen(1))
	g.Expect(cld.Services[cloud.ResourceManager]).To(Equal(cloud.AzurePublic.Services[cloud.ResourceManager]))
}

func Test_Cloud_ResourceManagerEndpointSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := config.Values{
		ResourceManagerEndpoint: "endpoint",
	}
	cld := cfg.Cloud()
	g.Expect(cld.ActiveDirectoryAuthorityHost).To(Equal(cloud.AzurePublic.ActiveDirectoryAuthorityHost))
	g.Expect(cld.Services).To(HaveLen(1))
	g.Expect(cld.Services[cloud.ResourceManager].Endpoint).To(Equal(cfg.ResourceManagerEndpoint))
	g.Expect(cld.Services[cloud.ResourceManager].Audience).To(Equal(cloud.AzurePublic.Services[cloud.ResourceManager].Audience))
}

func Test_Cloud_ResourceManagerAudienceSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cfg := config.Values{
		ResourceManagerAudience: "audience",
	}
	cld := cfg.Cloud()
	g.Expect(cld.ActiveDirectoryAuthorityHost).To(Equal(cloud.AzurePublic.ActiveDirectoryAuthorityHost))
	g.Expect(cld.Services).To(HaveLen(1))
	g.Expect(cld.Services[cloud.ResourceManager].Endpoint).To(Equal(cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint))
	g.Expect(cld.Services[cloud.ResourceManager].Audience).To(Equal(cfg.ResourceManagerAudience))
}

func Test_ParseTLSMinVersion(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		input       string
		expected    uint16
		expectError bool
	}{
		{name: "TLS12 is valid", input: "VersionTLS12", expected: tls.VersionTLS12},
		{name: "TLS13 is valid", input: "VersionTLS13", expected: tls.VersionTLS13},
		{name: "matches the default", input: config.DefaultTLSMinVersion, expected: tls.VersionTLS12},
		{name: "unsupported version is rejected", input: "TLSabcd", expectError: true},
		{name: "TLS11 is rejected", input: "VersionTLS11", expectError: true},
		{name: "empty value is rejected", input: "", expectError: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			v, err := config.ParseTLSMinVersion(c.input)
			if c.expectError {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err).To(MatchError(ContainSubstring(c.input)))
				return
			}

			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(v).To(Equal(c.expected))
		})
	}
}

func TestValidate(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	vOk := config.Values{
		PodNamespace:            "test-namespace",
		MaxConcurrentReconciles: 1,
		OperatorMode:            3,
		TargetNamespaces:        []string{},
		DefaultReconcilePolicy:  "detach-on-delete",
	}

	g.Expect(vOk.Validate()).NotTo(HaveOccurred())

	vKoNamespace := config.Values{
		PodNamespace: "",
	}

	err := vKoNamespace.Validate()
	g.Expect(err).To(MatchError(ContainSubstring("POD_NAMESPACE")))
	g.Expect(err).To(MatchError(ContainSubstring("missing value")))

	g.Expect(vOk.Validate()).NotTo(HaveOccurred())

	vKoConcurrentReconciles := config.Values{
		PodNamespace:            "test-namespace",
		MaxConcurrentReconciles: 0,
		DefaultReconcilePolicy:  "detach-on-delete",
	}

	g.Expect(vKoConcurrentReconciles.Validate().Error()).Error().Should(Equal("MAX_CONCURRENT_RECONCILES must be at least 1"))

	vKoDefaultReconcilePolicy := config.Values{
		PodNamespace:            "test-namespace",
		MaxConcurrentReconciles: 1,
		DefaultReconcilePolicy:  "detach",
	}

	g.Expect(vKoDefaultReconcilePolicy.Validate().Error()).Error().Should(Equal("DEFAULT_RECONCILE_POLICY must be set to any of (detach-on-delete, manage, skip)"))

	vKoOperatorMode := config.Values{
		PodNamespace:            "test-namespace",
		MaxConcurrentReconciles: 1,
		DefaultReconcilePolicy:  "detach-on-delete",
		TargetNamespaces:        []string{"to-be-watched"},
		OperatorMode:            2,
	}

	g.Expect(vKoOperatorMode.Validate().Error()).Error().Should(Equal("AZURE_TARGET_NAMESPACES must include watchers to specify target namespaces"))
}
