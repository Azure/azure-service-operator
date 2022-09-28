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
	"regexp"
	"strings"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/Azure/azure-service-operator/v2/internal/resolver"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const (
	refsPackage          = "refs"
	defaultResourceGroup = "aso-sample-rg"
)

// exclusions slice contains RESOURCES to exclude from test
var exclusions = []string{
	// Excluding webtest as it contains hidden link reference
	"webtest",

	// Excluding dbformysql/user as is not an ARM resource
	"user",

	// Excluding sqlroleassignment because it needs to reference a principalId
	"sqlroleassignment",
}

type SamplesTester struct {
	noSpaceNamer      ResourceNamer
	decoder           runtime.Decoder
	scheme            *runtime.Scheme
	groupVersionPath  string
	namespace         string
	useRandomName     bool
	rgName            string
	azureSubscription string
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

func NewSamplesTester(
	noSpaceNamer ResourceNamer,
	scheme *runtime.Scheme,
	groupVersionPath string,
	namespace string,
	useRandomName bool,
	rgName string,
	azureSubscription string) *SamplesTester {
	return &SamplesTester{
		noSpaceNamer:      noSpaceNamer,
		decoder:           serializer.NewCodecFactory(scheme).UniversalDecoder(),
		scheme:            scheme,
		groupVersionPath:  groupVersionPath,
		namespace:         namespace,
		useRandomName:     useRandomName,
		rgName:            rgName,
		azureSubscription: azureSubscription,
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
	err = t.setOwnershipAndReferences(samples.SamplesMap)
	if err != nil {
		return nil, err
	}
	err = t.setOwnershipAndReferences(samples.RefsMap)
	if err != nil {
		return nil, err
	}
	return samples, nil
}

// handleObject handles the sample object by adding it into the samples map. If key already exists, then we append a
// random string to the key and add it to the map so that we don't overwrite the sample. As keys are only used to find
// if owner Kind actually exist in the map, so it should be fine to append random string here.
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

func (t *SamplesTester) setOwnershipAndReferences(samples map[string]genruntime.ARMMetaObject) error {
	for gk, sample := range samples {
		// We don't apply ownership to the resources which have no owner
		if sample.Owner() == nil {
			continue
		}

		var ownersName string
		if sample.Owner().Kind == resolver.ResourceGroupKind {
			ownersName = t.rgName
		} else if t.useRandomName {
			owner, ok := samples[sample.Owner().Kind]
			if !ok {
				return fmt.Errorf("owner: %s, does not exist for resource '%s'", sample.Owner().Kind, gk)
			}
			ownersName = owner.GetName()
		}

		if ownersName != "" {
			var err error
			sample = setOwnersName(sample, ownersName)
			sample, err = t.findAndSetARMIDReferences(sample)
			if err != nil {
				return err
			}
		}
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

func (t *SamplesTester) findAndSetARMIDReferences(obj genruntime.ARMMetaObject) (genruntime.ARMMetaObject, error) {
	resourceRefType := reflect.TypeOf(genruntime.ResourceReference{})
	visitor := reflecthelpers.NewReflectVisitor()
	visitor.VisitStruct = func(this *reflecthelpers.ReflectVisitor, it reflect.Value, ctx interface{}) error {

		if it.Type() == resourceRefType {
			if it.CanInterface() {
				reference := it.Interface().(genruntime.ResourceReference)
				if reference.ARMID != "" {
					armIDField := it.FieldByName("ARMID")
					if !armIDField.CanSet() {
						return errors.New("cannot set 'ARMID' field of 'genruntime.ResourceReference'")
					}

					//Regex to match '/00000000-0000-0000-0000-000000000000/' strings, to replace with the subscriptionID
					subMatcher := regexp.MustCompile("\\/([0]+-?)+\\/")
					armIDString := armIDField.String()
					armIDString = strings.ReplaceAll(armIDString, defaultResourceGroup, t.rgName)
					armIDString = subMatcher.ReplaceAllString(armIDString, fmt.Sprint("/", t.azureSubscription, "/"))

					armIDField.SetString(armIDString)
				}
			} else {
				// This should be impossible given how the visitor works
				return errors.New("genruntime.ResourceReference field was unexpectedly nil")
			}
			return nil
		}

		return reflecthelpers.IdentityVisitStruct(this, it, ctx)
	}

	err := visitor.Visit(obj, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "scanning for references of type %s", resourceRefType.String())
	}

	return obj, nil
}
