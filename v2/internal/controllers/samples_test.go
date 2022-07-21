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
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	. "github.com/onsi/gomega"
	"k8s.io/api/core/v1"
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

	useRandomName := !testcommon.IsExclusion(groupVersionPath, randomNameExclusions)
	samples, refs, err := testcommon.NewSamplesTester(tc.NoSpaceNamer, tc.GetScheme(), groupVersionPath, tc.Namespace, useRandomName).LoadSamples()

	tc.Expect(err).To(BeNil())
	tc.Expect(samples).ToNot(BeNil())
	tc.Expect(refs).ToNot(BeNil())

	tc.Expect(samples).ToNot(BeZero())

	rg := tc.NewTestResourceGroup()
	tc.CreateResourceAndWait(rg)

	// Check if we have any references for the samples beforehand and Create them
	createResourceTree(tc, refs.SamplesMap.Values(), true, rg.Name, 0)
	createResourceTree(tc, samples.SamplesMap.Values(), false, rg.Name, 0)

	tc.DeleteResourceAndWait(rg)
}

func createResourceTree(tc *testcommon.KubePerTestContext, resourceChain []interface{}, isRef bool, rgName string, index int) {

	if index >= len(resourceChain) {
		return
	}

	resourceObj := resourceChain[index].(genruntime.ARMMetaObject)
	if resourceObj.Owner() != nil && resourceObj.Owner().Kind == resolver.ResourceGroupKind {
		resourceObj = testcommon.SetOwnersName(resourceObj, rgName)
	}

	findRefsAndCreateSecrets(tc, resourceObj)
	tc.CreateResourceAndWait(resourceObj)

	//Using recursion here to maintain the order of creation and deletion of resources
	createResourceTree(tc, resourceChain, isRef, rgName, index+1)
}

func findRefsAndCreateSecrets(tc *testcommon.KubePerTestContext, resourceObj genruntime.ARMMetaObject) {
	refs, err := reflecthelpers.FindSecretReferences(resourceObj)
	tc.Expect(err).To(BeNil())
	secrets := make(map[string]*v1.Secret)
	for ref := range refs {
		password := tc.Namer.GeneratePasswordOfLength(40)

		secret := &v1.Secret{
			ObjectMeta: tc.MakeObjectMetaWithName(ref.Name),
			StringData: map[string]string{
				ref.Key: password,
			},
		}

		tc.CreateResource(secret)
		secrets[ref.Name] = secret
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
