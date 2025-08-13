/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crdmanagement

import (
	"regexp"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

var deprecatedVersionRegexp = regexp.MustCompile(`((v1alpha1api|v1beta)\d{8}(preview)?(storage)?|v1beta1)`) // handcrafted (non-ARM) resources have v1beta1 version

// deprecatedVersionsMap is a map of crd name to a collection of deprecated versions
var deprecatedVersionsMap = map[string][]string{
	"trustedaccessrolebindings.containerservice.azure.com": {"v1api20230202previewstorage"},
}

// IsVersionDeprecated checks if a given version of a CRD is deprecated.
func IsVersionDeprecated(crd apiextensions.CustomResourceDefinition, version string) bool {
	if deprecatedVersionRegexp.MatchString(version) {
		return true
	}

	for _, deprecatedVersion := range deprecatedVersionsMap[crd.Name] {
		if deprecatedVersion == version {
			return true
		}
	}

	return false
}

// GetDeprecatedStorageVersions returns a new list of storedVersions by removing the deprecated versions
// The first return value is the list of versions that are NOT deprecated.
// The second return value is the list of versions that ARE deprecated.
func GetDeprecatedStorageVersions(crd apiextensions.CustomResourceDefinition) ([]string, []string) {
	storedVersions := crd.Status.StoredVersions

	newStoredVersions := make([]string, 0, len(storedVersions))
	var removedVersions []string
	for _, version := range storedVersions {
		if IsVersionDeprecated(crd, version) {
			removedVersions = append(removedVersions, version)
			continue
		}

		newStoredVersions = append(newStoredVersions, version)
	}

	return newStoredVersions, removedVersions
}

// GetDeprecatedStorageVersions returns a map of CRD names to their deprecated storage versions.
func GetAllDeprecatedStorageVersions(crds map[string]apiextensions.CustomResourceDefinition) map[string][]string {
	result := make(map[string][]string)
	for _, crd := range crds {
		_, deprecatedVersions := GetDeprecatedStorageVersions(crd)
		if len(deprecatedVersions) > 0 {
			result[crd.Name] = deprecatedVersions
		}
	}
	return result
}
