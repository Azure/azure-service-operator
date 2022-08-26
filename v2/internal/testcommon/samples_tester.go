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
	"image",
	// Network
	"loadbalancer",
	"virtualnetworkgateway",
	"virtualnetworksvirtualnetworkpeering",
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
	rgName           string
}

type SampleObject struct {
	SamplesMap map[string]genruntime.ARMMetaObject
	RefsMap    map[string]genruntime.ARMMetaObject
}

func NewSampleObject() *SampleObject {
	return &SampleObject{
		SamplesMap: make(map[string]genruntime.ARMMetaObject),
		RefsMap:    make(map[string]genruntime.ARMMetaObject),
	}
}

func NewSamplesTester(noSpaceNamer ResourceNamer, scheme *runtime.Scheme, groupVersionPath string, namespace string, useRandomName bool, rgName string) *SamplesTester {
	return &SamplesTester{
		noSpaceNamer:     noSpaceNamer,
		decoder:          serializer.NewCodecFactory(scheme).UniversalDecoder(),
		scheme:           scheme,
		groupVersionPath: groupVersionPath,
		namespace:        namespace,
		useRandomName:    useRandomName,
		rgName:           rgName,
	}
}

func (t *SamplesTester) LoadSamples() (*SampleObject, error) {
	samples := NewSampleObject()

	err := filepath.Walk(t.groupVersionPath,
		func(filePath string, info os.FileInfo, err error) error {

			if !info.IsDir() && !IsSampleExcluded(filePath, exclusions) {
				sample, err := t.getObjectFromFile(filePath)
				if err != nil {
					return err
				}

				sample.SetNamespace(t.namespace)

				if filepath.Base(filepath.Dir(filePath)) == refsPackage {
					// We don't change name for dependencies as they have fixed refs which we can't access inside the object
					// need to make sure we have right names for the refs in samples
					t.handleObject(sample, samples.RefsMap)
				} else {
					if t.useRandomName {
						sample.SetName(t.noSpaceNamer.GenerateName(""))
					}
					t.handleObject(sample, samples.SamplesMap)
				}
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	// We add ownership once we have all the resources in the map
	err = t.setOwnership(samples.SamplesMap)
	if err != nil {
		return nil, err
	}
	err = t.setOwnership(samples.RefsMap)
	if err != nil {
		return nil, err
	}
	return samples, nil
}

// handleObject handles the sample object by adding it into the samples map. If key already exists, then we append a
// random string to the key and add it to the map so that we don't overwrite the sample. As keys are only used to find
// if owner Kind actually exist in the map so it should be fine to append random string here.
func (t *SamplesTester) handleObject(sample genruntime.ARMMetaObject, samples map[string]genruntime.ARMMetaObject) {
	kind := sample.GetObjectKind().GroupVersionKind().Kind
	_, found := samples[kind]
	if found {
		kind = kind + t.noSpaceNamer.makeRandomStringOfLength(5, t.noSpaceNamer.runes)
	}
	samples[kind] = sample
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

func (t *SamplesTester) setOwnership(samples map[string]genruntime.ARMMetaObject) error {
	for gk, sample := range samples {
		// We don't apply ownership to the resources which have no owner
		if sample.Owner() == nil {
			continue
		}

		var ownersName string
		if sample.Owner().Kind == resolver.ResourceGroupKind {
			ownersName = t.rgName
		} else {
			owner, ok := samples[sample.Owner().Kind]
			if !ok {
				return fmt.Errorf("owner: %s, does not exist for resource '%s'", sample.Owner().Kind, gk)
			}
			ownersName = owner.GetName()
		}

		sample = setOwnersName(sample, ownersName)
	}
	return nil
}

func setOwnersName(sample genruntime.ARMMetaObject, ownerName string) genruntime.ARMMetaObject {
	specField := reflect.ValueOf(sample.GetSpec()).Elem()
	ownerField := specField.FieldByName("Owner").Elem()
	ownerField.FieldByName("Name").SetString(ownerName)

	return sample
}

func IsFolderExcluded(path string, exclusions []string) bool {
	for _, exclusion := range exclusions {
		if strings.Contains(path, exclusion) {
			return true
		}
	}
	return false
}

func IsSampleExcluded(path string, exclusions []string) bool {
	for _, exclusion := range exclusions {
		base := filepath.Base(path)
		baseWithoutAPIVersion := strings.Split(base, "_")[1]
		baseWithoutAPIVersion = strings.TrimSuffix(baseWithoutAPIVersion, filepath.Ext(baseWithoutAPIVersion))
		if baseWithoutAPIVersion == exclusion {
			return true
		}
	}
	return false
}
