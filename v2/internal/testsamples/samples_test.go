/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testsamples

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	asoentra "github.com/Azure/azure-service-operator/v2/api/entra/v1"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const samplesPath = "../../samples"
const redactedEntraID = "11111111-1111-1111-1111-111111111111"

// randomNameExclusions slice contains groups for which we don't want to use random names
var randomNameExclusions = []string{
	"/authorization/",
	"/cache/",
	"/containerservice/",
	"/compute/",
	"/cdn/",
	"/documentdb/",
	"/insights/",
	"/network/",
	"/web/",
	"/app/",
	"/dbforpostgresql/v1api20240801", // Only required starting when we added virtualendpoints support
	"/dbforpostgresql/v20250801",     // Only required starting when we added virtualendpoints support
}

func Test_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)

	regex, err := regexp.Compile("^v(1api)?[a-z0-9]*$")
	g.Expect(err).To(BeNil())

	_ = filepath.WalkDir(samplesPath,
		func(filePath string, info os.DirEntry, err error) error {
			if info.IsDir() && !testcommon.IsSampleFolderExcluded(filePath) {
				basePath := filepath.Base(filePath)
				// proceed only if the base path is the matching versions.
				if regex.MatchString(basePath) {

					testName := getTestName(filepath.Base(filepath.Dir(filePath)), basePath)
					t.Run(testName, func(t *testing.T) {
						t.Parallel()
						tc := globalTestContext.ForTest(t)
						runGroupTest(tc, filePath)
					})

				}
			}
			return err
		})
}

func runGroupTest(tc *testcommon.KubePerTestContext, groupVersionPath string) {
	rg := tc.NewTestResourceGroup()
	useRandomName := !testcommon.PathContains(groupVersionPath, randomNameExclusions)
	samples, err := testcommon.NewSamplesTester(
		tc.NoSpaceNamer,
		tc.GetScheme(),
		groupVersionPath,
		tc.Namespace,
		useRandomName,
		rg.Name,
		tc.AzureSubscription,
		tc.AzureTenant,
	).
		LoadSamples()

	tc.Expect(err).To(BeNil())
	tc.Expect(samples).ToNot(BeNil())
	tc.Expect(samples).ToNot(BeZero())

	if !samples.HasSamples() {
		// No testable samples in this folder, skip
		return
	}

	tc.CreateResourceAndWait(rg)

	refsSlice := processSamples(samples.RefsMap)
	samplesSlice := processSamples(samples.SamplesMap)

	resources := append(refsSlice, samplesSlice...)

	// For secrets we need to look across refs and samples:
	findRefsAndCreateSecrets(tc, resources)

	preRedactionResources, remainingResources := splitResourcesForEntraIDRedaction(resources)
	if len(preRedactionResources) > 0 {
		tc.CreateResourcesAndWait(preRedactionResources...)
		addStatusEntraIDLiteralRedactions(preRedactionResources, tc.WithLiteralRedaction)
	}

	// Create the remaining resources once any runtime Entra ID redactions are registered.
	tc.CreateResourcesAndWait(remainingResources...)
	addStatusEntraIDLiteralRedactions(remainingResources, tc.WithLiteralRedaction)

	tc.DeleteResourceAndWait(rg)
}

func splitResourcesForEntraIDRedaction(resources []client.Object) ([]client.Object, []client.Object) {
	preRedaction := make([]client.Object, 0)
	remaining := make([]client.Object, 0, len(resources))

	for _, resource := range resources {
		if _, ok := resource.(*asoentra.ServicePrincipal); ok {
			preRedaction = append(preRedaction, resource)
			continue
		}

		remaining = append(remaining, resource)
	}

	return preRedaction, remaining
}

func addStatusEntraIDLiteralRedactions(
	resources []client.Object,
	addLiteralRedaction func(value string, replacement string),
) {
	for _, resource := range resources {
		servicePrincipal, ok := resource.(*asoentra.ServicePrincipal)
		if !ok || servicePrincipal.Status.EntraID == nil || *servicePrincipal.Status.EntraID == "" {
			continue
		}

		addLiteralRedaction(*servicePrincipal.Status.EntraID, redactedEntraID)
	}
}

func processSamples(samples map[string]client.Object) []client.Object {
	samplesSlice := make([]client.Object, 0, len(samples))

	for _, resourceObj := range samples {
		obj := resourceObj
		samplesSlice = append(samplesSlice, obj)
	}

	return samplesSlice
}

// findRefsAndCreateSecrets finds all references not matched by a corresponding genruntime.SecretDestination or hardcoded secret
// and generates secrets which correspond to those references
func findRefsAndCreateSecrets(tc *testcommon.KubePerTestContext, resources []client.Object) {
	allDestinationKeys := set.Make[string]() // key is name + "/" + key
	allReferences := make([]genruntime.SecretReference, 0)
	allSecrets := set.Make[string]() // key is namespace + "/" + name

	for _, obj := range resources {
		if secret, ok := obj.(*v1.Secret); ok {
			allSecrets.Add(fmt.Sprintf("%s/%s", secret.Namespace, secret.Name))
			continue
		}

		destinations, err := reflecthelpers.Find[genruntime.SecretDestination](obj)
		tc.Expect(err).To(BeNil())

		references, err := reflecthelpers.FindSecretReferences(obj)
		tc.Expect(err).To(BeNil())

		for _, dest := range destinations {
			allDestinationKeys.Add(fmt.Sprintf("%s/%s", dest.Name, dest.Key))
		}
		allReferences = append(allReferences, references...)
	}

	// Find orphaned references
	orphanRefs := set.Make[genruntime.SecretReference]()
	for _, ref := range allReferences {
		matchingDestinationKey := fmt.Sprintf("%s/%s", ref.Name, ref.Key)
		matchingSecret := fmt.Sprintf("%s/%s", tc.Namespace, ref.Name)
		if allSecrets.Contains(matchingSecret) {
			continue
		}
		if allDestinationKeys.Contains(matchingDestinationKey) {
			continue
		}

		orphanRefs.Add(ref)
	}

	for ref := range orphanRefs {
		password := tc.Namer.GeneratePasswordOfLength(40)

		secret := &v1.Secret{
			ObjectMeta: tc.MakeObjectMetaWithName(ref.Name),
			StringData: map[string]string{
				ref.Key: password,
			},
		}

		err := tc.CheckIfResourceExists(secret)
		if err != nil {
			tc.CreateResource(secret)
		}
	}
}

func getTestName(group string, version string) string {
	var result strings.Builder

	// Common Prefix
	result.WriteString("Test_")

	// Titlecase for each part of the group
	title := cases.Title(language.English)
	for _, part := range strings.Split(group, ".") {
		result.WriteString(title.String(part))
	}
	result.WriteString("_")

	// Append the version
	result.WriteString(version)

	// Common Suffix
	result.WriteString("_CreationAndDeletion")

	return result.String()
}

func TestGetTestName(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		group    string
		version  string
		expected string
	}{
		"simple": {
			group:    "group",
			version:  "version",
			expected: "Test_Group_version_CreationAndDeletion",
		},
		"Network": {
			group:    "network",
			version:  "v1api",
			expected: "Test_Network_v1api_CreationAndDeletion",
		},
		"Frontdoor": {
			group:    "network.frontdoor",
			version:  "v1api",
			expected: "Test_NetworkFrontdoor_v1api_CreationAndDeletion",
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			g.Expect(getTestName(c.group, c.version)).To(Equal(c.expected))
		})
	}
}

func TestAddStatusEntraIDLiteralRedactions(t *testing.T) {
	t.Parallel()

	t.Run("registers service principal entra IDs", func(t *testing.T) {
		t.Parallel()

		g := NewGomegaWithT(t)
		actual := map[string]string{}
		entraID := "22222222-2222-2222-2222-222222222222"
		resources := []client.Object{
			&asoentra.ServicePrincipal{
				Status: asoentra.ServicePrincipalStatus{
					EntraID: &entraID,
				},
			},
		}

		addStatusEntraIDLiteralRedactions(resources, func(value string, replacement string) {
			actual[value] = replacement
		})

		g.Expect(actual).To(Equal(map[string]string{
			entraID: redactedEntraID,
		}))
	})

	t.Run("ignores resources without an entra ID", func(t *testing.T) {
		t.Parallel()

		g := NewGomegaWithT(t)
		actual := map[string]string{}
		resources := []client.Object{
			&asoentra.ServicePrincipal{},
			&v1.Secret{},
		}

		addStatusEntraIDLiteralRedactions(resources, func(value string, replacement string) {
			actual[value] = replacement
		})

		g.Expect(actual).To(BeEmpty())
	})
}

func TestSplitResourcesForEntraIDRedaction(t *testing.T) {
	t.Parallel()

	t.Run("separates service principals from remaining resources", func(t *testing.T) {
		t.Parallel()

		g := NewGomegaWithT(t)
		servicePrincipal := &asoentra.ServicePrincipal{}
		secret := &v1.Secret{}
		preRedaction, remaining := splitResourcesForEntraIDRedaction([]client.Object{
			secret,
			servicePrincipal,
		})

		g.Expect(preRedaction).To(Equal([]client.Object{servicePrincipal}))
		g.Expect(remaining).To(Equal([]client.Object{secret}))
	})

	t.Run("leaves resources unchanged when there is no service principal", func(t *testing.T) {
		t.Parallel()

		g := NewGomegaWithT(t)
		secret := &v1.Secret{}
		preRedaction, remaining := splitResourcesForEntraIDRedaction([]client.Object{secret})

		g.Expect(preRedaction).To(BeEmpty())
		g.Expect(remaining).To(Equal([]client.Object{secret}))
	})
}
