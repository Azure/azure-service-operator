/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/yaml"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const refsPackage = "refs"
const samplesPath = "../../config/samples"

// exclusions slice contains resources to exclude from test as these resources run into
// "runtime.notRegisteredErr" error from the Scheme.
var exclusions = [...]string{
	"webtest",
	// excluding dbformysql/user as is not an ARM resource
	"user",
}

type SamplesTester struct {
	noSpaceNamer  ResourceNamer
	decoder       runtime.Decoder
	scheme        *runtime.Scheme
	UseRandomName bool
}

type SampleObject struct {
	SamplesMap   map[string]*linkedhashmap.Map
	childObjects map[string]genruntime.ARMMetaObject
}

func NewSampleObject() *SampleObject {
	return &SampleObject{
		SamplesMap:   make(map[string]*linkedhashmap.Map),
		childObjects: make(map[string]genruntime.ARMMetaObject),
	}
}

func NewSamplesTester(noSpaceNamer ResourceNamer, scheme *runtime.Scheme, useRandomName bool) *SamplesTester {
	return &SamplesTester{
		noSpaceNamer:  noSpaceNamer,
		decoder:       serializer.NewCodecFactory(scheme).UniversalDecoder(),
		scheme:        scheme,
		UseRandomName: useRandomName,
	}
}

func (t *SamplesTester) LoadSamples(rg *resources.ResourceGroup, namespace string, group string) (*SampleObject, *SampleObject, error) {

	samples := NewSampleObject()
	depSamples := NewSampleObject()

	groupPath := getGroupPath(group)

	err := filepath.Walk(groupPath,
		func(filePath string, info os.FileInfo, err error) error {

			if !info.IsDir() && !isExclusion(filePath) {
				sample, err := t.getObjectFromFile(filePath)
				if err != nil {
					return err
				}

				sample.SetNamespace(namespace)

				if !strings.Contains(filePath, refsPackage) {
					if t.UseRandomName {
						sample.SetName(t.noSpaceNamer.GenerateName(""))
					}
					handleObject(sample, rg, samples)
				} else {
					// We don't change name for dependencies as they have fixed refs which we can't access inside the object
					// need to make sure we have right names for the refs in samples
					handleObject(sample, rg, depSamples)
				}
			}

			return nil
		})

	if err != nil {
		return nil, nil, err
	}

	// We process the child objects once we have all the parent objects in the tree already
	err = handleChildObjects(t.scheme, samples)
	if err != nil {
		return nil, nil, err
	}
	err = handleChildObjects(t.scheme, depSamples)
	if err != nil {
		return nil, nil, err
	}
	return samples, depSamples, nil
}

func (t *SamplesTester) getObjectFromFile(path string) (genruntime.ARMMetaObject, error) {
	jsonMap := make(map[string]interface{})

	byteData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	jsonBytes, err := yaml.ToJSON(byteData)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		return nil, err
	}

	// We need unstructured object here to fetch the correct GVK for the object
	unstructuredObj := unstructured.Unstructured{Object: jsonMap}

	obj, err := genruntime.NewObjectFromExemplar(&unstructuredObj, t.scheme)
	if err != nil {
		return nil, err
	}

	err = runtime.DecodeInto(t.decoder, byteData, obj)
	if err != nil {
		return nil, err
	}

	return obj.(genruntime.ARMMetaObject), nil
}

func handleObject(sample genruntime.ARMMetaObject, rg *resources.ResourceGroup, samples *SampleObject) {
	if sample.Owner().Kind == resolver.ResourceGroupKind {
		handleParentObject(sample, rg, samples.SamplesMap)
	} else {
		samples.childObjects[sample.GetObjectKind().GroupVersionKind().String()] = sample
	}
}

func handleChildObjects(scheme *runtime.Scheme, samples *SampleObject) error {
	for gvk, sample := range samples.childObjects {

		sampleGVKAware, ok := sample.(genruntime.GroupVersionKindAware)
		if !ok {
			//return error
			return fmt.Errorf("error in casting '%s' to type '%s'", sample.GetObjectKind().GroupVersionKind().Kind, "genruntime.GroupVersionKindAware")
		}

		version := sampleGVKAware.OriginalGVK().Version
		tree, ok := samples.SamplesMap[version]
		// Parent version should always exist
		if !ok {
			return fmt.Errorf("found resource: %s, for which owner does not exist", gvk)
		}

		// get the parent object
		val, ok := tree.Get(sample.Owner().Kind)
		// If owner already does not exist, we check if its unvisited in 'childObjects' map.
		// If it does, we wait until we visit the owner and add is to the tree or the test condition fails
		if !ok {
			ownerGVK := constructGVK(sample.Owner().Group, version, sample.Owner().Kind)
			_, ok = samples.childObjects[ownerGVK]
			if !ok {
				return fmt.Errorf("owner: %s, does not exist for resource '%s'", ownerGVK, gvk)
			}
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
		err := handleChildObjects(scheme, samples)
		if err != nil {
			return err
		}
	}
	return nil
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

func constructGVK(group string, version string, kind string) string {
	return group + "/" + version + ", Kind=" + kind

}
