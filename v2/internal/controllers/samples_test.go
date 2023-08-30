/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
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
	"subscription", // Can't easily be run/recorded in our standard subscription
}

// randomNameExclusions slice contains groups for which we don't want to use random names
var randomNameExclusions = []string{
	"authorization",
	"cache",
	"containerservice",
	"compute",
	"documentdb",
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

	// For secrets we need to look across refs and samples:
	findRefsAndCreateSecrets(tc, samples.SamplesMap, samples.RefsMap)

	refsSlice := processSamples(samples.RefsMap)
	samplesSlice := processSamples(samples.SamplesMap)

	// Check if we have any references for the samples beforehand and Create them
	tc.CreateResourcesAndWait(refsSlice...)
	tc.CreateResourcesAndWait(samplesSlice...)

	tc.DeleteResourceAndWait(rg)
}

func processSamples(samples map[string]genruntime.ARMMetaObject) []client.Object {
	samplesSlice := make([]client.Object, 0, len(samples))

	for _, resourceObj := range samples {
		obj := resourceObj.(client.Object)
		samplesSlice = append(samplesSlice, obj)
	}

	return samplesSlice
}

// findRefsAndCreateSecrets finds all references not matched by a corresponding genruntime.SecretDestination and generates
// secrets which correspond to those references
func findRefsAndCreateSecrets(tc *testcommon.KubePerTestContext, samples map[string]genruntime.ARMMetaObject, refs map[string]genruntime.ARMMetaObject) {
	allDestinations := set.Make[genruntime.SecretDestination]()
	allReferences := set.Make[genruntime.SecretReference]()

	// Join samples and refs
	resources := make([]genruntime.ARMMetaObject, 0, len(samples)+len(refs))
	for _, v := range samples {
		resources = append(resources, v)
	}
	for _, v := range refs {
		resources = append(resources, v)
	}

	for _, resourceObj := range resources {
		destinations, err := reflecthelpers.Find[genruntime.SecretDestination](resourceObj)
		tc.Expect(err).To(BeNil())

		references, err := reflecthelpers.FindSecretReferences(resourceObj)
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
