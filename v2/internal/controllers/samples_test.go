/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/yaml"
)

const samplesPath = "../../config/samples"

// exclusions slice contains resources to exclude from test as we have disabled these
// resources from being generated in ASOv2 or their configuration is based on other resources.
var exclusions = [...]string{"batchaccountpool", "snapshot",
	"sqldatabasecontainerthroughputsetting", "sqldatabasethroughputsetting",
	"webtest", "storageaccountsqueueservice", "resourcegroup",
	// Problematic ones, their configuration is based on other resources
	"loadbalancer", "virtualmachine",
}

func Test_Samples_CreationAndDeletion(t *testing.T) {

	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	samples, err := loadSamples(tc, rg)
	tc.Expect(err).To(gomega.BeNil())
	tc.Expect(len(samples)).ToNot(gomega.BeZero())

	for gv, resourceTree := range samples {
		fmt.Sprintf("Test for '%s' samples", gv)
		//tc.RunParallelSubtests(
		//	testcommon.Subtest{
		//		Name: fmt.Sprintf("Test for '%s' samples", gv),
		//		Test: func(testContext *testcommon.KubePerTestContext) {
		createAndDeleteResourceTree(tc, resourceTree)
		//		},
		//	},
		//)

	}
}

func createAndDeleteResourceTree(tc *testcommon.KubePerTestContext, resourceTree *treemap.Map) {

	k, v := resourceTree.Min()

	if k != nil && v != nil {
		resourceObj := v.(genruntime.ARMMetaObject)
		tc.CreateResourceAndWait(resourceObj)

		resourceTree.Remove(k)
		createAndDeleteResourceTree(tc, resourceTree)

		// no DELETE for child objects, to delete them we must delete its parent
		if resourceObj.Owner().Kind == resolver.ResourceGroupKind {
			tc.DeleteResourceAndWait(resourceObj)
		}
	}

}

func loadSamples(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) (map[string]*treemap.Map, error) {

	samples := make(map[string]*treemap.Map)
	var childObjects []genruntime.ARMMetaObject

	decoder := serializer.NewCodecFactory(tc.GetScheme()).UniversalDecoder()

	err := filepath.Walk(samplesPath,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}
			if !info.IsDir() && !isExclusion(path) {
				//if !info.IsDir() && strings.Contains(path, "storage") && !isExclusion(path) {

				sample := getObjectFromFile(tc, decoder, path)

				sample.SetNamespace(tc.Namespace)
				sample.SetName(tc.NoSpaceNamer.GenerateName(""))

				if sample.Owner().Kind == resolver.ResourceGroupKind {
					handleParentObject(sample, rg, samples)
				} else {
					childObjects = append(childObjects, sample)
				}
			}

			return nil
		})

	// TODO: Find a better way to handle Child resources. Doing it this way as we need to make sure all the parent
	// objects are already present in the map.
	handleChildObjects(tc, childObjects, samples)

	return samples, err
}

func getObjectFromFile(tc *testcommon.KubePerTestContext, decoder runtime.Decoder, path string) genruntime.ARMMetaObject {

	jsonMap := make(map[string]interface{})

	byteData, err := ioutil.ReadFile(path)
	tc.Expect(err).To(gomega.BeNil())

	jsonBytes, err := yaml.ToJSON(byteData)
	tc.Expect(err).To(gomega.BeNil())

	err = json.Unmarshal(jsonBytes, &jsonMap)
	tc.Expect(err).To(gomega.BeNil())

	// We need unstructured object here to fetch the correct GVK for the object
	unstructuredObj := unstructured.Unstructured{Object: jsonMap}

	obj, err := genruntime.NewObjectFromExemplar(&unstructuredObj, tc.GetScheme())
	tc.Expect(err).To(gomega.BeNil())

	byteData, err = json.Marshal(unstructuredObj.Object)
	tc.Expect(err).To(gomega.BeNil())

	err = runtime.DecodeInto(decoder, byteData, obj)

	return obj.(genruntime.ARMMetaObject)
}

func handleChildObjects(tc *testcommon.KubePerTestContext, childObjects []genruntime.ARMMetaObject, samples map[string]*treemap.Map) {
	for _, sample := range childObjects {
		specField := reflect.ValueOf(sample.GetSpec())
		version := specField.MethodByName("OriginalVersion").Call([]reflect.Value{})

		tc.Expect(version[0]).ToNot(gomega.BeNil())
		tc.Expect(version[0].String()).ToNot(gomega.BeEmpty())

		gv := generateGVString(sample.Owner().Group, version[0].String())

		node, ok := samples[gv]
		tc.Expect(ok).To(gomega.BeTrue())

		// TODO: Fow now, we don't have complex cases and this should pass always. If in future we have nested child objects,
		// We might need to have a look into extending this.
		val, ok := node.Get(sample.Owner().Kind)
		tc.Expect(ok).To(gomega.BeTrue())

		owner := val.(genruntime.ARMMetaObject)

		sample = setOwnersName(sample, owner.GetName())
		samples[gv].Put(sample.GetObjectKind().GroupVersionKind().Kind, sample)
		samples[gv] = node
	}
}

func handleParentObject(sample genruntime.ARMMetaObject, rg *resources.ResourceGroup, samples map[string]*treemap.Map) {
	gv := sample.GetObjectKind().GroupVersionKind().GroupVersion().String()
	sample = setOwnersName(sample, rg.Name)

	kind := sample.GetObjectKind().GroupVersionKind().Kind

	if _, ok := samples[gv]; ok {
		samples[gv].Put(kind, sample)
	} else {
		treeMap := treemap.NewWithStringComparator()
		treeMap.Put(kind, sample)
		samples[gv] = treeMap
	}
}

func isExclusion(path string) bool {

	for _, exc := range exclusions {
		if strings.Contains(path, exc) {
			return true
		}
	}
	return false
}

func setOwnersName(sample genruntime.ARMMetaObject, ownerName string) genruntime.ARMMetaObject {
	specField := reflect.ValueOf(sample.GetSpec()).Elem()
	ownerField := specField.FieldByName("Owner").Elem()
	ownerField.FieldByName("Name").SetString(ownerName)

	return sample
}

func generateGVString(group, version string) string {
	return group + "/" + version
}
