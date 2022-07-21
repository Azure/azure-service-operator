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
	"path/filepath"
	"reflect"
	"strings"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const refsPackage = "refs"

//TODO(super-harsh) Need a take on helper methods to perform reflection/parsing and write up a few
// helper methods which can find and replace the distinct property values for the excluded resources below.

// exclusions slice contains RESOURCES to exclude from test
var exclusions = []string{
	// Excluding webtest as it contains hidden link reference
	"webtest",

	// Below resources are excluded as they contain internal references or armIDs.
	// Compute
	"virtualmachine",
	"virtualmachinescaleset",
	// Network
	"loadbalancer",
	"networkinterface",
	"virtualnetworkgateway",
	"virtualnetworksubnet",
	"virtualnetworksvirtualnetworkpeering",
	"image",
	// MachineLearningServices
	"workspacescompute",

	// Excluding dbformysql/user as is not an ARM resource
	"user",
}

type SamplesTester struct {
	noSpaceNamer     ResourceNamer
	decoder          runtime.Decoder
	scheme           *runtime.Scheme
	groupVersionPath string
	namespace        string
	useRandomName    bool
}

type SampleObject struct {
	// Using linkedhashmap.Map here to maintain the order(owner->children) of resources we put into the map
	// to make sure we always have an owner resource before a child resource.
	SamplesMap   *linkedhashmap.Map
	childObjects map[string]genruntime.ARMMetaObject
}

func NewSampleObject() *SampleObject {
	return &SampleObject{
		SamplesMap:   linkedhashmap.New(),
		childObjects: make(map[string]genruntime.ARMMetaObject),
	}
}

func NewSamplesTester(noSpaceNamer ResourceNamer, scheme *runtime.Scheme, groupVersionPath string, namespace string, useRandomName bool) *SamplesTester {
	return &SamplesTester{
		noSpaceNamer:     noSpaceNamer,
		decoder:          serializer.NewCodecFactory(scheme).UniversalDecoder(),
		scheme:           scheme,
		groupVersionPath: groupVersionPath,
		namespace:        namespace,
		useRandomName:    useRandomName,
	}
}

func (t *SamplesTester) LoadSamples() (*SampleObject, *SampleObject, error) {
	samples := NewSampleObject()
	refSamples := NewSampleObject()

	err := filepath.Walk(t.groupVersionPath,
		func(filePath string, info os.FileInfo, err error) error {

			if !info.IsDir() && !IsExclusion(filePath, exclusions) {
				sample, err := t.getObjectFromFile(filePath)
				if err != nil {
					return err
				}

				sample.SetNamespace(t.namespace)

				if filepath.Base(filepath.Dir(filePath)) == refsPackage {
					// We don't change name for dependencies as they have fixed refs which we can't access inside the object
					// need to make sure we have right names for the refs in samples
					t.handleObject(sample, refSamples, t.noSpaceNamer)
				} else {
					if t.useRandomName {
						sample.SetName(t.noSpaceNamer.GenerateName(""))
					}
					t.handleObject(sample, samples, t.noSpaceNamer)
				}
			}

			return nil
		})

	if err != nil {
		return nil, nil, err
	}

	// We process the child objects once we have all the parent objects in the linkedMap already
	err = t.handleChildObjects(t.scheme, samples)
	if err != nil {
		return nil, nil, err
	}
	err = t.handleChildObjects(t.scheme, refSamples)
	if err != nil {
		return nil, nil, err
	}
	return samples, refSamples, nil
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

func (t *SamplesTester) handleObject(sample genruntime.ARMMetaObject, samples *SampleObject, namer ResourceNamer) {
	if sample.Owner() == nil || sample.Owner().Kind == resolver.ResourceGroupKind {
		putIntoLinkedMap(sample, samples.SamplesMap, t.noSpaceNamer)
	} else {
		gk := sample.GetObjectKind().GroupVersionKind().GroupKind().String()
		_, found := samples.childObjects[gk]
		if found {
			gk = gk + namer.makeRandomStringOfLength(3, namer.runes)
		}
		samples.childObjects[gk] = sample
	}
}

func (t *SamplesTester) handleChildObjects(scheme *runtime.Scheme, samples *SampleObject) error {
	for gk, sample := range samples.childObjects {

		linkedMap := samples.SamplesMap
		// get the parent object
		val, ok := linkedMap.Get(sample.Owner().Kind)

		// If owner already does not exist, we check if its unvisited in 'childObjects' map.
		// If it does, we wait until we visit the owner and add is to the linkedMap or the test condition fails
		if !ok {
			ownerGK := constructGroupKind(sample.Owner().Group, sample.Owner().Kind)
			_, ok = samples.childObjects[ownerGK]
			if !ok {
				return fmt.Errorf("owner: %s, does not exist for resource '%s'", ownerGK, gk)
			}
			continue
		}

		owner := val.(genruntime.ARMMetaObject)
		sample = SetOwnersName(sample, owner.GetName())
		putIntoLinkedMap(sample, linkedMap, t.noSpaceNamer)

		// Delete after processing
		delete(samples.childObjects, gk)
	}

	// We need to do recursion here, as per the logic in this block, we check if we've visited the owner
	// for current resource or not. If not, then we keep trying until we visit all the members of childObjects map.
	if len(samples.childObjects) > 0 {
		err := t.handleChildObjects(scheme, samples)
		if err != nil {
			return err
		}
	}
	return nil
}

func putIntoLinkedMap(sample genruntime.ARMMetaObject, samplesMap *linkedhashmap.Map, namer ResourceNamer) {
	kind := sample.GetObjectKind().GroupVersionKind().Kind
	_, found := samplesMap.Get(kind)
	// if found the same kind in a samples directory, we just simply add some random chars in the end of the kind
	// as to avoid overwriting and  to maintain the order.
	if found {
		kind = kind + namer.makeRandomStringOfLength(3, namer.runes)
	}
	samplesMap.Put(kind, sample)
}

func SetOwnersName(sample genruntime.ARMMetaObject, ownerName string) genruntime.ARMMetaObject {
	specField := reflect.ValueOf(sample.GetSpec()).Elem()
	ownerField := specField.FieldByName("Owner").Elem()
	ownerField.FieldByName("Name").SetString(ownerName)

	return sample
}

func IsExclusion(path string, exclusions []string) bool {
	for _, exclusion := range exclusions {
		filepath.Dir(path)
		if strings.Contains(path, exclusion) {
			return true
		}
	}
	return false
}

func constructGroupKind(group string, kind string) string {
	return kind + "." + group

}
