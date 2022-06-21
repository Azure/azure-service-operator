/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Authorization_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "authorization", false)
}

func Test_Batch_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "batch", false)
}

// Had to split cache resources into version groups as the time taken for recording
// as a whole was exceeding 150 minutes
func Test_Cache_Alpha_Redis_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "cache/v1alphaapi", false)
}

func Test_Cache_Beta_Redis_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "cache/v1beta", false)
}

func Test_CDN_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "cdn", true)
}

func Test_ContainerInstance_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "containerinstance", true)
}

func Test_ContainerRegistry_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "containerregistry", true)
}

func Test_ContainerService_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "containerservice", false)
}

func Test_DbforMariaDB_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "dbformariadb", true)
}

// TODO: FlexiblServer is not being able to be created. Hitting RESPONSE 404 as ARM is sending a success
// and still controller can't find the created resource TODO: update me

//func Test_DbForMySQL_Samples_CreationAndDeletion(t *testing.T) {
//	t.Parallel()
//	tc := globalTestContext.ForTest(t)
//	runGroupTest(tc, "dbformysql", false)
//}

//func Test_DbForPostgreSQL_Samples_CreationAndDeletion(t *testing.T) {
//	t.Parallel()
//	tc := globalTestContext.ForTest(t)
//	runGroupTest(tc, "dbforpostgresql", true)
//}

func Test_DocumentDB_MongoDB_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "documentdb/mongodb", false)
}

func Test_DocumentDB_SqlDatabase_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "documentdb/sqldatabase", false)
}

func Test_Eventgrid_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "eventgrid", true)
}

func Test_EventHub_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "eventhub", true)
}

func Test_Insights_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "insights", true)
}

func Test_Keyvault_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "keyvault", true)
}

func Test_ManagedIdentity_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "managedidentity", true)
}

func Test_OperationalInsights_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "operationalinsights", true)
}

func Test_SignalRService_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "signalrservice", true)
}

func Test_ServiceBus_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "servicebus", true)
}

func Test_Storage_Samples_CreationAndDeletion(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	runGroupTest(tc, "storage", true)
}

// TODO: These are the complex groups which have too many nested referenced objects and armID refs
// TODO: which we can't handle currently due to having no capability to figure out the order
// TODO: of installing the references
//func Test_Compute_Samples_CreationAndDeletion(t *testing.T) {
//	t.Parallel()
//	tc := globalTestContext.ForTest(t)
//	runGroupTest(tc, "compute", true)
//}

//func Test_Network_Samples_CreationAndDeletion(t *testing.T) {
//	t.Parallel()
//	tc := globalTestContext.ForTest(t)
//	runGroupTest(tc, "network", true)
//}

func runGroupTest(tc *testcommon.KubePerTestContext, group string, useRandomName bool) {
	rg := tc.NewTestResourceGroup()
	samples, refs := testcommon.NewSamplesTester(tc, useRandomName).LoadSamples(rg, group)

	tc.Expect(samples).ToNot(gomega.BeNil())
	tc.Expect(refs).ToNot(gomega.BeNil())

	tc.Expect(samples).ToNot(gomega.BeZero())

	tc.CreateResourceAndWait(rg)
	defer tc.DeleteResourceAndWait(rg)

	for _, refTree := range refs.SamplesMap {
		createAndDeleteResourceTree(tc, refTree, true, 0)
	}

	for _, resourceTree := range samples.SamplesMap {
		// Check if we have any references for the samples beforehand and Create them
		createAndDeleteResourceTree(tc, resourceTree, false, 0)
	}
}

func createAndDeleteResourceTree(tc *testcommon.KubePerTestContext, hashMap *linkedhashmap.Map, isRef bool, index int) {

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
	createAndDeleteResourceTree(tc, hashMap, isRef, index+1)

	for _, secret := range secrets {
		tc.DeleteResource(secret)
	}

	// no DELETE for child objects, to delete them we must delete its parent
	if resourceObj.Owner().Kind == resolver.ResourceGroupKind && !isRef {
		tc.DeleteResourceAndWait(resourceObj)
	}
}
