/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"strings"
	"testing"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

var cases = []struct {
	Group          string
	UseRandomName  bool
	DeleteChildren bool
}{
	{
		Group:         "authorization",
		UseRandomName: false,
	},
	{
		Group:         "batch",
		UseRandomName: true,
	},
	{
		Group:          "cache",
		UseRandomName:  false,
		DeleteChildren: true,
	},
	{
		Group:         "cdn",
		UseRandomName: true,
	},
	{
		Group:         "containerinstance",
		UseRandomName: true,
	},
	{
		Group:         "containerregistry",
		UseRandomName: true,
	},
	{
		Group:          "containerservice",
		UseRandomName:  false,
		DeleteChildren: true,
	},
	{
		Group:         "dbformariadb",
		UseRandomName: true,
	},
	{
		Group:          "documentdb/mongodb",
		UseRandomName:  false,
		DeleteChildren: false,
	},
	{
		Group:          "documentdb/sqldatabase",
		UseRandomName:  false,
		DeleteChildren: false,
	},
	{
		Group:         "eventgrid",
		UseRandomName: true,
	},
	{
		Group:         "eventhub",
		UseRandomName: true,
	},
	{
		Group:         "insights",
		UseRandomName: true,
	},
	{
		Group:         "keyvault",
		UseRandomName: true,
	},
	{
		Group:         "managedidentity",
		UseRandomName: true,
	},
	{
		Group:         "operationalinsights",
		UseRandomName: true,
	},
	{
		Group:         "servicebus",
		UseRandomName: true,
	},
	{
		Group:         "signalrservice",
		UseRandomName: true,
	},
	{
		Group:         "storage/blobservice",
		UseRandomName: true,
	},
	{
		Group:         "storage/queueservice",
		UseRandomName: true,
	},

	// TODO: These are the complex groups which have too many nested referenced objects and armID refs
	// TODO: which we can't handle currently due to having no capability to figure out the order
	// TODO: of installing the references
	//{
	//	Group: "compute",
	//	UseRandomName: true,
	//},
	//{
	//	Group: "network",
	//	UseRandomName: true,
	//},

	// TODO: FlexiblServer is not being able to be created. Hitting RESPONSE 404 as ARM is sending a success
	// and still controller can't find the created resource.
	//{
	//	Group: "dbformysql",
	//	UseRandomName: false,
	//},
	//{
	//	Group: "dbforpostgresql",
	//	UseRandomName: true,
	//},
}

func Test_Samples_CreationAndDeletion(t *testing.T) {
	for _, test := range cases {
		test := test
		testName := strings.Join(
			[]string{
				"Test",
				strings.Title(strings.ReplaceAll(test.Group, "/", "_")),
				"CreationAndDeletion",
			}, "_")
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			tc := globalTestContext.ForTest(t)
			runGroupTest(
				tc,
				test.Group,
				test.UseRandomName,
				test.DeleteChildren,
			)
		})
	}

}

func runGroupTest(tc *testcommon.KubePerTestContext, group string, useRandomName bool, deleteChildren bool) {
	rg := tc.NewTestResourceGroup()

	samples, refs, err := testcommon.NewSamplesTester(tc.NoSpaceNamer, tc.GetScheme(), useRandomName).LoadSamples(rg, tc.Namespace, group)

	tc.Expect(err).To(BeNil())
	tc.Expect(samples).ToNot(BeNil())
	tc.Expect(refs).ToNot(BeNil())

	tc.Expect(samples).ToNot(BeZero())

	tc.CreateResourceAndWait(rg)
	defer tc.DeleteResourceAndWait(rg)

	for _, refTree := range refs.SamplesMap {
		createAndDeleteResourceTree(tc, refTree, true, false, 0)
	}

	for _, resourceTree := range samples.SamplesMap {
		// Check if we have any references for the samples beforehand and Create them
		createAndDeleteResourceTree(tc, resourceTree, false, deleteChildren, 0)
	}
}

func createAndDeleteResourceTree(tc *testcommon.KubePerTestContext, hashMap *linkedhashmap.Map, isRef bool, deleteChildren bool, index int) {

	var secrets []*v1.Secret
	vals := hashMap.Values()
	if index >= hashMap.Size() {
		return
	}

	resourceObj := vals[index].(genruntime.ARMMetaObject)
	refs, err := reflecthelpers.FindSecretReferences(resourceObj)
	if err != nil {
		return
	}

	for ref, _ := range refs {
		password := tc.Namer.GeneratePasswordOfLength(40)

		secret := &v1.Secret{
			ObjectMeta: tc.MakeObjectMetaWithName(ref.Name),
			StringData: map[string]string{
				ref.Key: password,
			},
		}

		tc.CreateResource(secret)
		secrets = append(secrets, secret)
	}

	tc.CreateResourceAndWait(resourceObj)

	// Using recursion here to maintain the order of creation and deletion of resources
	createAndDeleteResourceTree(tc, hashMap, isRef, deleteChildren, index+1)

	for _, secret := range secrets {
		tc.DeleteResource(secret)
	}

	if !isRef {
		// No DELETE for child objects all, to delete them we must delete its parent
		if resourceObj.Owner().Kind == resolver.ResourceGroupKind || deleteChildren {
			tc.DeleteResourceAndWait(resourceObj)
		}
	}
}
