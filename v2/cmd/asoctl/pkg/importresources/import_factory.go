/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// importFactory is a helper struct that's used to create things.
type importFactory struct {
	scheme *runtime.Scheme
}

func newImportFactory(
	scheme *runtime.Scheme,
) *importFactory {
	return &importFactory{
		scheme: scheme,
	}
}

// createBlankObjectFromGVK creates a blank object of from a given GVK.
func (f *importFactory) createBlankObjectFromGVK(gvk schema.GroupVersionKind) (runtime.Object, error) {
	obj, err := f.scheme.New(gvk)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create blank resource")
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
			errors.Errorf(
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
			return schema.GroupVersionKind{}, errors.Wrapf(err, "unable to create blank resource for GVK %s", gvk)
		}

		if _, ok := obj.(genruntime.ImportableResource); ok {
			if result != nil {
				return schema.GroupVersionKind{},
					errors.Errorf(
						"multiple known versions for Group %s, Kind %s implement genruntime.ImportableResource",
						gk.Group,
						gk.Kind)
			}

			result = &gvk
		}
	}

	if result == nil {
		return schema.GroupVersionKind{},
			errors.Errorf(
				"no known versions for Group %s, Kind %s implement genruntime.ImportableResource",
				gk.Group,
				gk.Kind)
	}

	return *result, nil
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
