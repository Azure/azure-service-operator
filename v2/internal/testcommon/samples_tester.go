/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"

	"github.com/Azure/azure-service-operator/v2/internal/resolver"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const (
	refsPackage          = "refs"
	defaultResourceGroup = "aso-sample-rg"
)

// Regex to match '/00000000-0000-0000-0000-000000000000/' strings, to replace with the subscriptionID
var subRegex = regexp.MustCompile("\\/([0]+-?)+\\/")

// An empty GUID, used to replace the subscriptionID and tenantID in the sample files
var emptyGuid = uuid.Nil.String()

// exclusions slice contains RESOURCES to exclude from test
var exclusions = []string{
	// Excluding webtest as it contains hidden link reference
	"webtest",

	// Excluding dbformysql/user as is not an ARM resource
	"user",

	// Excluding sql serversadministrator and serversazureadonlyauthentication as they both require AAD auth
	// which the samples recordings aren't using.
	"serversadministrator",
	"serversazureadonlyauthentication",
	"serversfailovergroup", // Requires creating multiple linked SQL servers which is hard to do in the samples

	// TODO: Unable to test diskencryptionsets sample since it requires keyvault/key URI.
	// TODO: we don't support Keyvault/Keys to automate the process
	"diskencryptionset",
}

type SamplesTester struct {
	noSpaceNamer      ResourceNamer
	scheme            *runtime.Scheme
	groupVersionPath  string
	namespace         string
	useRandomName     bool
	rgName            string
	azureSubscription string
	azureTenant       string
}

type SampleObject struct {
	SamplesMap map[string]genruntime.ARMMetaObject
	RefsMap    map[string]genruntime.ARMMetaObject
}

func (s *SampleObject) HasSamples() bool {
	return len(s.SamplesMap) > 0
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
	azureSubscription string,
	azureTenant string,
) *SamplesTester {
	return &SamplesTester{
		noSpaceNamer:      noSpaceNamer,
		scheme:            scheme,
		groupVersionPath:  groupVersionPath,
		namespace:         namespace,
		useRandomName:     useRandomName,
		rgName:            rgName,
		azureSubscription: azureSubscription,
		azureTenant:       azureTenant,
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

	byteData, err := os.ReadFile(path)
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

	decorder := json.NewDecoder(bytes.NewReader(jsonBytes))
	decorder.DisallowUnknownFields()
	err = decorder.Decode(obj)
	if err != nil {
		return nil, errors.Wrapf(err, "while decoding %s", path)
	}

	return obj.(genruntime.ARMMetaObject), nil
}

func (t *SamplesTester) setOwnershipAndReferences(samples map[string]genruntime.ARMMetaObject) error {
	for gk, sample := range samples {
		// We don't apply ownership to the resources which have no owner
		if sample.Owner() == nil {
			continue
		}

		// We only set the owner name if Owner.Kind is ResourceGroup(as we have random rg names) or if we're using random names for resources.
		// Otherwise, we let it be the same as on samples.
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
			sample = setOwnersName(sample, ownersName)
		}

		err := t.updateFieldsForTest(sample)
		if err != nil {
			return err
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

// updateFieldsForTest uses ReflectVisitor to update ARMReferences, SubscriptionIDs, and TenantIDs.
func (t *SamplesTester) updateFieldsForTest(obj genruntime.ARMMetaObject) error {
	visitor := reflecthelpers.NewReflectVisitor()
	visitor.VisitStruct = t.visitStruct

	err := visitor.Visit(obj, nil)
	if err != nil {
		return errors.Wrapf(err, "updating fields for test")
	}

	return nil
}

// visitStruct checks and sets the SubscriptionID and ResourceGroup name for ARM references to current values
func (t *SamplesTester) visitStruct(this *reflecthelpers.ReflectVisitor, it reflect.Value, ctx any) error {

	// Configure any ResourceReference we find
	if it.Type() == reflect.TypeOf(genruntime.ResourceReference{}) {
		return t.visitResourceReference(this, it, ctx)
	}

	// Set the value of any SubscriptionID Field that's got an empty GUID as the value
	if field := it.FieldByNameFunc(isField("subscriptionID")); field.IsValid() {
		t.assignString(field, t.azureSubscription)
	}

	// Set the value of any TenantID Field that's got an empty GUID as the value
	if field := it.FieldByNameFunc(isField("tenantID")); field.IsValid() {
		t.assignString(field, t.azureTenant)
	}

	return reflecthelpers.IdentityVisitStruct(this, it, ctx)
}

// isField is a helper used to find fields by case-insensitive matches
func isField(field string) func(name string) bool {
	return func(name string) bool {
		return strings.EqualFold(name, field)
	}
}

func (t *SamplesTester) assignString(field reflect.Value, value string) {
	if field.Kind() == reflect.String &&
		// Set simple string field
		field.String() == emptyGuid {
		field.SetString(value)
		return
	}

	if field.Kind() == reflect.Ptr &&
		field.Elem().Kind() == reflect.String &&
		field.Elem().String() == emptyGuid {
		// Set pointer to string field
		field.Elem().SetString(value)
	}
}

// visitResourceReference checks and sets the SubscriptionID and ResourceGroup name for ARM references to current values
func (t *SamplesTester) visitResourceReference(_ *reflecthelpers.ReflectVisitor, it reflect.Value, _ any) error {
	if !it.CanInterface() {
		// This should be impossible given how the visitor works
		panic("genruntime.ResourceReference field was unexpectedly nil")
	}

	reference := it.Interface().(genruntime.ResourceReference)
	if reference.ARMID == "" {
		return nil
	}

	armIDField := it.FieldByName("ARMID")
	if !armIDField.CanSet() {
		return errors.New("cannot set 'ARMID' field of 'genruntime.ResourceReference'")
	}

	armIDString := armIDField.String()
	armIDString = strings.ReplaceAll(armIDString, defaultResourceGroup, t.rgName)
	armIDString = subRegex.ReplaceAllString(armIDString, fmt.Sprint("/", t.azureSubscription, "/"))

	armIDField.SetString(armIDString)

	return nil
}
