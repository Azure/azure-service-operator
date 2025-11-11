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
	"github.com/rotisserie/eris"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const (
	refsPackage          = "refs"
	defaultResourceGroup = "aso-sample-rg"
)

// Regex to match '/00000000-0000-0000-0000-000000000000' strings, to replace with the subscriptionID
var subRegex = regexp.MustCompile(`/([0]+-?)+`)

// An empty GUID, used to replace the subscriptionID and tenantID in the sample files
var emptyGUID = uuid.Nil.String()

var wholeSampleExclusions = []*regexp.Regexp{
	regexp.MustCompile(`/cache/`),                               // Cache has issues with linked caches being able to delete
	regexp.MustCompile(`/subscription/`),                        // Can't easily be run/recorded in our standard subscription
	regexp.MustCompile(`/redhatopenshift/`),                     // This requires SP creation
	regexp.MustCompile(`/documentdb/sqldatabase/v1api20210515`), // This is blocked by corp policy (can't set DisableLocalAuth)
}

var exclusions = []*regexp.Regexp{
	// ------------------------------
	// Individual resource exclusions
	// ------------------------------
	regexp.MustCompile(`insights/.*_webtest.yaml`), // Excluding webtest as it contains hidden link reference

	// db users aren't ARM resources
	regexp.MustCompile(`sql/.*_user.yaml`),
	regexp.MustCompile(`dbformysql/.*_user.yaml`),
	regexp.MustCompile(`dbformysql/.*_user_aad.yaml`),
	regexp.MustCompile(`dbforpostgresql/.*_user.yaml`),

	// Excluding sql serversadministrator and serversazureadonlyauthentication as they both require AAD auth
	// which the samples recordings aren't using.
	regexp.MustCompile(`sql/.*_serversadministrator.yaml`),
	regexp.MustCompile(`sql/.*_serversazureadonlyauthentication.yaml`),

	// Requires creating multiple linked SQL servers which is hard to do in the samples
	regexp.MustCompile(`sql/.*_serversfailovergroup.yaml`),

	// TODO: Unable to test diskencryptionsets sample since it requires keyvault/key URI.
	// TODO: we don't support Keyvault/Keys to automate the process
	regexp.MustCompile(`compute/.*_diskencryptionset.yaml`),

	// Excluding APIM Product and Subscription as we need to pass deleteSubscription flag to delete the subscription
	// when we delete the Product. https://github.com/Azure/azure-service-operator/issues/3408
	regexp.MustCompile(`apimanagement/.*_api.yaml`),
	regexp.MustCompile(`apimanagement/.*_apiversionset.yaml`),
	regexp.MustCompile(`apimanagement/.*_product.yaml`),
	regexp.MustCompile(`apimanagement/.*_subscription.yaml`),
	regexp.MustCompile(`apimanagement/.*_productpolicy.yaml`),
	regexp.MustCompile(`apimanagement/.*_productapi.yaml`),

	// Excluding cdn secret as it requires KV secrets
	regexp.MustCompile(`cdn/.*_secret.yaml`),

	// Excluding SignalR CustomDomain and CustomCertificate because they require KV secrets/certs
	regexp.MustCompile(`signalrservice/.*_customdomain.yaml`),
	regexp.MustCompile(`signalrservice/.*_customcertificate.yaml`),

	// [Issue #3091] Exclude backupvaultsbackupinstance as it requires role assignments to be created after backup instance is created to make it land into protection configured state.
	regexp.MustCompile(`dataprotection/.*_backupvaultsbackupinstance.yaml`),

	regexp.MustCompile(`network/.*_virtualnetworkgateway.yaml`), // blocks RG deletion and causes networking tests to fail

	// Excluding quota as Azure Quota API does not support deletion - quotas are read-only system resources
	regexp.MustCompile(`quota/.*_quota.yaml`),
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
	SamplesMap map[string]client.Object
	RefsMap    map[string]client.Object
}

func (s *SampleObject) HasSamples() bool {
	return len(s.SamplesMap) > 0
}

func NewSampleObject() *SampleObject {
	return &SampleObject{
		SamplesMap: make(map[string]client.Object),
		RefsMap:    make(map[string]client.Object),
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
			if !info.IsDir() && !IsSampleFileExcluded(filePath) {
				sample, err := t.getObjectFromFile(filePath)
				if err != nil {
					return eris.Wrapf(err, "loading sample from %s", filePath)
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
		return nil, eris.Wrapf(err, "loading samples in %s", t.groupVersionPath)
	}

	// We add ownership once we have all the resources in the map
	err = t.setSamplesOwnershipAndReferences(samples.SamplesMap, samples.RefsMap)
	if err != nil {
		return nil, eris.Wrap(err, "updating ownership of samples")
	}

	err = t.setRefsOwnershipAndReferences(samples.RefsMap)
	if err != nil {
		return nil, eris.Wrap(err, "updating ownership of refs")
	}

	return samples, nil
}

// handleObject handles the sample object by adding it into the samples map. If key already exists, then we append a
// random string to the key and add it to the map so that we don't overwrite the sample. As keys are only used to find
// if owner Kind actually exist in the map, so it should be fine to append random string here.
func (t *SamplesTester) handleObject(sample client.Object, samples map[string]client.Object) {
	kind := sample.GetObjectKind().GroupVersionKind().Kind
	_, found := samples[kind]
	if found {
		kind = kind + t.noSpaceNamer.makeRandomStringOfLength(5, t.noSpaceNamer.runes)
	}
	samples[kind] = sample
}

func (t *SamplesTester) getObjectFromFile(path string) (client.Object, error) {
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

	decoder := json.NewDecoder(bytes.NewReader(jsonBytes))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(obj)
	if err != nil {
		return nil, eris.Wrapf(err, "while decoding %s", path)
	}

	return obj, nil
}

func (t *SamplesTester) setSamplesOwnershipAndReferences(samples map[string]client.Object, refs map[string]client.Object) error {
	for gk, sample := range samples {
		asoType, ok := sample.(genruntime.ARMMetaObject)
		if !ok {
			continue
		}

		// We don't apply ownership to the resources which have no owner
		if asoType.Owner() == nil {
			continue
		}

		// We only set the owner name if Owner.Kind is ResourceGroup(as we have random rg names) or if we're using random names for resources.
		// Otherwise, we let it be the same as on samples.
		var ownersName string
		if asoType.Owner().Kind == resolver.ResourceGroupKind {
			ownersName = t.rgName
		} else if t.useRandomName {
			// Check if the owner exists in refs, then continue. We don't use random names for refs so its correct anyway.
			_, found := refs[asoType.Owner().Kind]
			if found {
				continue
			}

			var owner client.Object
			owner, found = samples[asoType.Owner().Kind]
			if !found {
				return fmt.Errorf("owner: %s, does not exist for resource '%s'", asoType.Owner().Kind, gk)
			}

			ownersName = owner.GetName()
		}

		if ownersName != "" {
			asoType = setOwnersName(asoType, ownersName)
		}

		err := t.updateFieldsForTest(asoType)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *SamplesTester) setRefsOwnershipAndReferences(samples map[string]client.Object) error {
	for _, sample := range samples {
		asoType, ok := sample.(genruntime.ARMMetaObject)
		if !ok {
			continue
		}

		// We don't apply ownership to the resources which have no owner
		if asoType.Owner() == nil {
			continue
		}

		// We only set the owner name if Owner.Kind is ResourceGroup(as we have random rg names) or if we're using random names for resources.
		// Otherwise, we let it be the same as on samples.
		var ownersName string
		if asoType.Owner().Kind != resolver.ResourceGroupKind {
			continue
		}

		ownersName = t.rgName

		if ownersName != "" {
			asoType = setOwnersName(asoType, ownersName)
		}

		err := t.updateFieldsForTest(asoType)
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

func PathContains(path string, matches []string) bool {
	for _, match := range matches {
		if strings.Contains(path, match) {
			return true
		}
	}

	return false
}

func IsSampleFolderExcluded(path string) bool {
	for _, exclusion := range wholeSampleExclusions {
		if exclusion.MatchString(path) {
			return true
		}
	}

	return false
}

func IsSampleFileExcluded(path string) bool {
	// Exclude everything that's not a yaml file
	ext := filepath.Ext(path)
	if ext != ".yaml" && ext != ".yml" {
		return true
	}

	// Check all exclusions
	for _, exclusion := range exclusions {
		if exclusion.MatchString(path) {
			return true
		}
	}

	return false
}

// updateFieldsForTest uses ReflectVisitor to update ARMReferences, SubscriptionIDs, and TenantIDs.
func (t *SamplesTester) updateFieldsForTest(obj genruntime.ARMMetaObject) error {
	visitor := reflecthelpers.NewReflectVisitor()
	visitor.VisitStruct = t.visitStruct

	err := visitor.Visit(obj, t.rgName)
	if err != nil {
		return eris.Wrapf(err, "updating fields for test")
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
		t.conditionalAssignString(field, emptyGUID, t.azureSubscription)
	}

	// Replace the empty-guid value in any armID field
	if field := it.FieldByNameFunc(isField("armId")); field.IsValid() {
		t.replaceString(field, emptyGUID, t.azureSubscription)
	}

	// Set the value of any TenantID Field that's got an empty GUID as the value
	if field := it.FieldByNameFunc(isField("tenantID")); field.IsValid() {
		t.conditionalAssignString(field, emptyGUID, t.azureTenant)
	}

	return reflecthelpers.IdentityVisitStruct(this, it, ctx)
}

// isField is a helper used to find fields by case-insensitive matches
func isField(field string) func(name string) bool {
	return func(name string) bool {
		return strings.EqualFold(name, field)
	}
}

func (t *SamplesTester) conditionalAssignString(field reflect.Value, match string, value string) {
	if field.Kind() == reflect.Ptr {
		field = field.Elem()
	}

	if field.Kind() != reflect.String {
		return
	}

	if field.String() == match {
		field.SetString(value)
	}
}

func (t *SamplesTester) replaceString(field reflect.Value, old string, new string) {
	if field.Kind() == reflect.Ptr {
		field = field.Elem()
	}

	if field.Kind() != reflect.String {
		return
	}

	val := field.String()
	val = strings.ReplaceAll(val, old, new)
	field.SetString(val)
}

// visitResourceReference checks and sets the SubscriptionID and ResourceGroup name for ARM references to current values
func (t *SamplesTester) visitResourceReference(_ *reflecthelpers.ReflectVisitor, it reflect.Value, ctx any) error {
	if !it.CanInterface() {
		// This should be impossible given how the visitor works
		panic("genruntime.ResourceReference field was unexpectedly nil")
	}

	ownersName := ctx.(string)

	reference := it.Interface().(genruntime.ResourceReference)
	if reference.ARMID != "" {
		armIDField := it.FieldByName("ARMID")
		if !armIDField.CanSet() {
			return eris.New("cannot set 'ARMID' field of 'genruntime.ResourceReference'")
		}

		armIDString := armIDField.String()
		armIDString = strings.ReplaceAll(armIDString, defaultResourceGroup, t.rgName)
		armIDString = subRegex.ReplaceAllString(armIDString, fmt.Sprint("/", t.azureSubscription))

		armIDField.SetString(armIDString)
	} else if reference.Kind == "ResourceGroup" && ownersName != "" { // If we're referring to a resourceGroup, it needs to be updated to refer to the random one
		// TODO: We're making the assumption that every reference of type ResourceGroup is by definition referring
		// TODO: to the randomly generated RG name, but it's possible at some future date we have multiple resourceGroups
		// TODO: floating around. If that happens we may need to update this logic to be a bit more discerning.
		nameField := it.FieldByName("Name")
		if !nameField.CanSet() {
			return eris.New("cannot set 'Name' field of 'genruntime.ResourceReference'")
		}

		nameField.SetString(ownersName)
	}

	return nil
}
