// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config_test

import (
	"reflect"
	"testing"

	// Importing this for side effects is required as it initializes cloud.AzurePublic
	_ "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	. "github.com/onsi/gomega"

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
