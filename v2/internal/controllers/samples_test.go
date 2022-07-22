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

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	. "github.com/onsi/gomega"
	"k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const samplesPath = "../../config/samples"

// skipTests slice contains the groups to skip from being tested.
var skipTests = []string{
	// TODO: Will re-record test .. as resource was having some issues in creation.
	"cache",
}

// randomNameExclusions slice contains groups for which we don't want to use random names
var randomNameExclusions = []string{
	"authorization",
	"cache",
	"containerservice",
	"documentdb",
	"network",
}

func Test_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	regex, err := regexp.Compile("^v1(alpha|beta)[a-z0-9]*$")
	g.Expect(err).To(BeNil())

	_ = filepath.WalkDir(samplesPath,
		func(filePath string, info os.DirEntry, err error) error {
			if info.IsDir() && !testcommon.IsExclusion(filePath, skipTests) {
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
	useRandomName := !testcommon.IsExclusion(groupVersionPath, randomNameExclusions)
	samples, err := testcommon.NewSamplesTester(tc.NoSpaceNamer, tc.GetScheme(), groupVersionPath, tc.Namespace, useRandomName, rg.Name).LoadSamples()

	tc.Expect(err).To(BeNil())
	tc.Expect(samples).ToNot(BeNil())

	tc.Expect(samples).ToNot(BeZero())

	tc.CreateResourceAndWait(rg)

	refsSlice := processSamples(tc, samples.RefsMap)
	samplesSlice := processSamples(tc, samples.SamplesMap)

	// Check if we have any references for the samples beforehand and Create them
	tc.CreateResourcesAndWait(refsSlice...)
	tc.CreateResourcesAndWait(samplesSlice...)

	tc.DeleteResourceAndWait(rg)
}

func processSamples(tc *testcommon.KubePerTestContext, samples map[string]genruntime.ARMMetaObject) []client.Object {
	var samplesSlice []client.Object

	for _, resourceObj := range samples {
		r := resourceObj.(client.Object)

		findRefsAndCreateSecrets(tc, resourceObj)
		samplesSlice = append(samplesSlice, r)
	}

	return samplesSlice
}

func findRefsAndCreateSecrets(tc *testcommon.KubePerTestContext, resourceObj genruntime.ARMMetaObject) {
	refs, err := reflecthelpers.FindSecretReferences(resourceObj)
	tc.Expect(err).To(BeNil())
	for ref := range refs {
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
