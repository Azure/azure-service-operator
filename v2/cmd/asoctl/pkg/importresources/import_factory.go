/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"crypto/rand"
	"math/big"
	"regexp"
	"strings"
	"sync"
	"unicode"

	"github.com/rotisserie/eris"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// importFactory is a helper struct that's used to create things.
type importFactory struct {
	scheme        *runtime.Scheme
	usedNames     map[schema.GroupKind]set.Set[string]
	usedNamesLock sync.Mutex
}

func newImportFactory(
	scheme *runtime.Scheme,
) *importFactory {
	return &importFactory{
		scheme:    scheme,
		usedNames: make(map[schema.GroupKind]set.Set[string]),
	}
}

// createBlankObjectFromGVK creates a blank object of from a given GVK.
func (f *importFactory) createBlankObjectFromGVK(gvk schema.GroupVersionKind) (runtime.Object, error) {
	obj, err := f.scheme.New(gvk)
	if err != nil {
		return nil, eris.Wrap(err, "unable to create blank resource")
	}

	obj.GetObjectKind().SetGroupVersionKind(gvk)
	return obj, nil
}

// selectVersionFromGK selects the latest version of a given GroupKind.
// The latest stable version will be selected if it exists, otherwise the latest preview version will be selected.
func (f *importFactory) selectVersionFromGK(gk schema.GroupKind) (schema.GroupVersionKind, error) {
	knownVersions := f.scheme.VersionsForGroupKind(gk)
	if len(knownVersions) == 0 {
		return schema.GroupVersionKind{},
			eris.Errorf(
				"no known versions for Group %s, Kind %s",
				gk.Group,
				gk.Kind)
	}

	// Scan for the GVK that implements genruntime.ImportableResource
	// We expect there to be exactly one
	var result *schema.GroupVersionKind
	for _, gv := range knownVersions {
		gvk := gk.WithVersion(gv.Version)
		obj, err := f.createBlankObjectFromGVK(gvk)
		if err != nil {
			return schema.GroupVersionKind{}, eris.Wrapf(err, "unable to create blank resource for GVK %s", gvk)
		}

		if _, ok := obj.(genruntime.ImportableResource); ok {
			if result != nil {
				return schema.GroupVersionKind{},
					eris.Errorf(
						"multiple known versions for Group %s, Kind %s implement genruntime.ImportableResource",
						gk.Group,
						gk.Kind)
			}

			result = &gvk
		}
	}

	if result == nil {
		return schema.GroupVersionKind{},
			eris.Errorf(
				"no known versions for Group %s, Kind %s implement genruntime.ImportableResource",
				gk.Group,
				gk.Kind)
	}

	return *result, nil
}

// createUniqueKubernetesName creates a name that is unique within a given resource type.
func (f *importFactory) createUniqueKubernetesName(
	name string,
	gk schema.GroupKind,
) string {
	// Only one object of a given kind can have a given name at a time.
	// See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/

	// Protect against concurrent access to the usedNames map
	f.usedNamesLock.Lock()
	defer f.usedNamesLock.Unlock()

	used, knownGK := f.usedNames[gk]
	if !knownGK {
		used = set.Make[string]()
		f.usedNames[gk] = used
	}

	baseName := f.createKubernetesName(name)
	suffix := ""
	for {
		uniqueName := baseName
		if suffix != "" {
			uniqueName = baseName + "-" + suffix
		}

		if !used.Contains(uniqueName) {
			used.Add(uniqueName)
			return uniqueName
		}

		suffix = f.createSuffix(6)
	}
}

var suffixRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func (f *importFactory) createSuffix(size int) string {
	buffer := make([]rune, 0, size)

	for range size {
		rnge := big.NewInt(int64(len(suffixRunes)))
		index, err := rand.Int(rand.Reader, rnge)
		if err != nil {
			// This should never happen
			panic(err)
		}

		i := int(index.Int64())
		buffer = append(buffer, suffixRunes[i])
	}

	return string(buffer)
}

var kubernetesNameRegex = regexp.MustCompile("-+")

// createKubernetesName creates a name compliant with kubernetes naming conventions.
func (f *importFactory) createKubernetesName(
	name string,
) string {
	// Most resource types require a name that can be used as a DNS subdomain name as defined in RFC 1123.
	// This means the name must:
	// o contain no more than 253 characters
	// o contain only lowercase alphanumeric characters, '-' or '.'
	// o start with an alphanumeric character
	// o end with an alphanumeric character
	// See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/

	buffer := make([]rune, 0, len(name))

	for _, r := range name {
		switch {
		case unicode.IsLetter(r):
			// Transform letters to lowercase
			buffer = append(buffer, unicode.ToLower(r))

		case unicode.IsNumber(r):
			// Keep numbers as they are
			buffer = append(buffer, r)

		default:
			// Any other characters are converted to hyphens
			buffer = append(buffer, '-')
		}
	}

	result := string(buffer)

	// Collapse any sequences of hyphens into a single hyphen
	result = kubernetesNameRegex.ReplaceAllString(result, "-")

	// Remove any leading or trailing hyphens
	result = strings.Trim(result, "-")

	return result
}
