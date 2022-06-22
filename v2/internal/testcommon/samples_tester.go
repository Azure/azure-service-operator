/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/yaml"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const RefsPackage = "refs"
const samplesPath = "../../config/samples"

// exclusions slice contains resources to exclude from test as these resources run into
// "runtime.notRegisteredErr" error from the Scheme.
var exclusions = [...]string{
        "batchaccountpool", 
        "snapshot",
	"sqldatabasecontainerthroughputsetting", 
	"sqldatabasethroughputsetting",
	"storageaccountsqueueservice",
	"webtest", 
	// excluding dbformysql/user as is not an ARM resource
	"user",
}

type SamplesTester struct {
	tc            *KubePerTestContext
	decoder       runtime.Decoder
	UseRandomName bool
}

type SampleObject struct {
	SamplesMap   map[string]*linkedhashmap.Map
	childObjects map[string]genruntime.ARMMetaObject
}

func NewSampleObject() SampleObject {
	return SampleObject{
		SamplesMap:   make(map[string]*linkedhashmap.Map),
		childObjects: make(map[string]genruntime.ARMMetaObject),
	}
}

func NewSamplesTester(tc *KubePerTestContext, useRandomName bool) *SamplesTester {
	return &SamplesTester{
		tc:            tc,
		decoder:       serializer.NewCodecFactory(tc.GetScheme()).UniversalDecoder(),
		UseRandomName: useRandomName,
	}
}

func (t SamplesTester) LoadSamples(rg *resources.ResourceGroup, group string) (*SampleObject, *SampleObject) {

	samples := NewSampleObject()
	depSamples := NewSampleObject()

	groupPath := getGroupPath(group)

	err := filepath.Walk(groupPath,
		func(filePath string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}
			if !info.IsDir() && !isExclusion(filePath) {
				//if !info.IsDir() && !isExclusion(filePath) {

				sample := t.getObjectFromFile(filePath)

				sample.SetNamespace(t.tc.Namespace)

				if !strings.Contains(filePath, RefsPackage) {
					if t.UseRandomName {
						sample.SetName(t.tc.NoSpaceNamer.GenerateName(""))
					}
					handleObject(sample, rg, &samples)
				} else {
					// We don't change name for dependencies as they have fixed refs which we can't access inside the object
					// need to make sure we have right names for the refs in samples
					handleObject(sample, rg, &depSamples)
				}
			}

			return nil
		})

	t.tc.Expect(err).To(BeNil())

	// We process the child objects once we have all the parent objects in the tree already
	handleChildObjects(t.tc, &samples)
	handleChildObjects(t.tc, &depSamples)

	return &samples, &depSamples
}

func (t SamplesTester) getObjectFromFile(path string) genruntime.ARMMetaObject {
	jsonMap := make(map[string]interface{})

	byteData, err := ioutil.ReadFile(path)
	t.tc.Expect(err).To(BeNil())

	jsonBytes, err := yaml.ToJSON(byteData)
	t.tc.Expect(err).To(BeNil())

	err = json.Unmarshal(jsonBytes, &jsonMap)
	t.tc.Expect(err).To(BeNil())

	// We need unstructured object here to fetch the correct GVK for the object
	unstructuredObj := unstructured.Unstructured{Object: jsonMap}

	obj, err := genruntime.NewObjectFromExemplar(&unstructuredObj, t.tc.GetScheme())
	t.tc.Expect(err).To(BeNil())

	byteData, err = json.Marshal(unstructuredObj.Object)
	t.tc.Expect(err).To(BeNil())

	err = runtime.DecodeInto(t.decoder, byteData, obj)

	return obj.(genruntime.ARMMetaObject)
}

func handleObject(sample genruntime.ARMMetaObject, rg *resources.ResourceGroup, samples *SampleObject) {
	if sample.Owner().Kind == resolver.ResourceGroupKind {
		handleParentObject(sample, rg, samples.SamplesMap)
	} else {
		samples.childObjects[sample.GetObjectKind().GroupVersionKind().String()] = sample
	}
}

func handleChildObjects(tc *KubePerTestContext, samples *SampleObject) {
	for gvk, sample := range samples.childObjects {

		specField := reflect.ValueOf(sample.GetSpec())
		originalVersion := specField.MethodByName("OriginalVersion").Call([]reflect.Value{})

		tc.Expect(originalVersion[0]).ToNot(BeNil())
		tc.Expect(originalVersion).ToNot(HaveLen(0))
		tc.Expect(originalVersion[0].String()).ToNot(BeEmpty())

		version := originalVersion[0].String()
		tree, ok := samples.SamplesMap[version]
		// Parent version should always exist
		tc.Expect(ok).To(BeTrue())

		// get the parent object
		val, ok := tree.Get(sample.Owner().Kind)
		// If owner already does not exist, we check if its unvisited in 'childObjects' map.
		// If it does, we wait until we visit the owner and add is to the tree or the test condition fails
		if !ok {
			ownerGVK := constructGVK(sample.Owner().Group, version, sample.Owner().Kind)
			_, ok = samples.childObjects[ownerGVK]
			tc.Expect(ok).To(BeTrue())
			continue
		}

		owner := val.(genruntime.ARMMetaObject)
		sample = setOwnersName(sample, owner.GetName())
		tree.Put(sample.GetObjectKind().GroupVersionKind().Kind, sample)

		// Delete after processing
		delete(samples.childObjects, gvk)
	}

	// We need to do recursion here, as per the logic in this block, we check if we've visited the owner
	// for current resource or not. If not, then we keep trying until we visit all the members of childObjects map.
	if len(samples.childObjects) > 0 {
		handleChildObjects(tc, samples)
	}
}

func handleParentObject(sample genruntime.ARMMetaObject, rg *resources.ResourceGroup, samples map[string]*linkedhashmap.Map) {

	version := sample.GetObjectKind().GroupVersionKind().GroupVersion().Version
	sample = setOwnersName(sample, rg.Name)

	kind := sample.GetObjectKind().GroupVersionKind().Kind
	// TODO: If we have similar Kind in the same group and version, we need to put them in a new directory
	if _, ok := samples[version]; ok {
		samples[version].Put(kind, sample)
	} else {
		linkedHashMap := linkedhashmap.New()
		linkedHashMap.Put(kind, sample)
		samples[version] = linkedHashMap
	}
}

func setOwnersName(sample genruntime.ARMMetaObject, ownerName string) genruntime.ARMMetaObject {
	specField := reflect.ValueOf(sample.GetSpec()).Elem()
	ownerField := specField.FieldByName("Owner").Elem()
	ownerField.FieldByName("Name").SetString(ownerName)

	return sample
}

func isExclusion(path string) bool {
	for _, exclusion := range exclusions {
		if strings.Contains(path, exclusion) {
			return true
		}
	}
	return false
}

func getGroupPath(group string) string {
	return path.Join(samplesPath, group)
}

func constructGVK(group, version, kind string) string {
	return group + "/" + version + ", Kind=" + kind

}
