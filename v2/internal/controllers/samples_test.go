/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const samplesPath = "../../samples"

// skipTests slice contains the groups to skip from being tested.
var skipTests = []string{
	// TODO: Cache has issues with linked caches being able to delete
	"cache",
	"subscription",    // Can't easily be run/recorded in our standard subscription
	"redhatopenshift", // This requires SP creation
}

// randomNameExclusions slice contains groups for which we don't want to use random names
var randomNameExclusions = []string{
	"authorization",
	"cache",
	"containerservice",
	"compute",
	"cdn",
	"documentdb",
	"insights",
	"network",
	"web",
}

func Test_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()

	if *isLive {
		t.Skip("skipping test in live mode")
	}

	g := NewGomegaWithT(t)

	regex, err := regexp.Compile("^v1(api|beta)[a-z0-9]*$")
	g.Expect(err).To(BeNil())

	_ = filepath.WalkDir(samplesPath,
		func(filePath string, info os.DirEntry, err error) error {
			if info.IsDir() && !testcommon.IsFolderExcluded(filePath, skipTests) {
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
	useRandomName := !testcommon.IsFolderExcluded(groupVersionPath, randomNameExclusions)
	samples, err := testcommon.NewSamplesTester(
		tc.NoSpaceNamer,
		tc.GetScheme(),
		groupVersionPath,
		tc.Namespace,
		useRandomName,
		rg.Name,
		tc.AzureSubscription,
		tc.AzureTenant).
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

	// Create all the resources
	tc.CreateResourcesAndWait(resources...)

	tc.DeleteResourceAndWait(rg)
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
	allDestinations := set.Make[genruntime.SecretDestination]()
	allReferences := set.Make[genruntime.SecretReference]()
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

		allDestinations.AddAll(destinations)
		allReferences.AddAll(references)
	}

	// Find orphaned references
	orphanRefs := set.Make[genruntime.SecretReference]()
	for _, ref := range allReferences.Values() {
		matchingDestination := genruntime.SecretDestination{
			Name: ref.Name,
			Key:  ref.Key,
		}
		matchingSecret := fmt.Sprintf("%s/%s", tc.Namespace, ref.Name)
		if allSecrets.Contains(matchingSecret) {
			continue
		}
		if allDestinations.Contains(matchingDestination) {
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
	return strings.Join(
		[]string{
			"Test",
			strings.Title(group),
			version,
			"CreationAndDeletion",
		}, "_")
}
