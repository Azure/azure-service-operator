/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import (
	"context"
	"fmt"
	"regexp"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CleanDeprecatedCRDVersions(ctx context.Context, cl apiextensions.CustomResourceDefinitionInterface) error {
	var crdRegexp = regexp.MustCompile(`.*\.azure\.com`)
	var deprecatedVersionRegexp = regexp.MustCompile(`v1alpha1api\d{8}(preview)?(storage)?`)

	list, err := cl.List(ctx, v1.ListOptions{})
	if err != nil {
		return err
	}

	if list == nil || len(list.Items) == 0 {
		fmt.Println("found 0 results, make sure you have ASO CRDs installed")
	}

	var updated int
	for _, item := range list.Items {
		item := item
		crdName := item.Name

		if !crdRegexp.MatchString(crdName) {
			continue
		}

		newStoredVersions := removeMatchingStoredVersions(item.Status.StoredVersions, deprecatedVersionRegexp)

		//
		if len(newStoredVersions) > 0 && len(newStoredVersions) != len(item.Status.StoredVersions) {
			item.Status.StoredVersions = newStoredVersions
			// It's fine to update the storedVersions and remove the deprecated versions this way.
			// Users can still use the existing old v1alpha1api versioned resources and would not require to migrate
			// due to the conversion webhook implemented.
			updatedCrd, err := cl.UpdateStatus(ctx, &item, v1.UpdateOptions{})
			if err != nil {
				return err
			}

			updated++
			fmt.Printf("updated '%s' CRD status storedVersions to : %s\n", crdName, updatedCrd.Status.StoredVersions)
		}
	}

	fmt.Printf("updated %d CRD(s), no deprecated versions found\n", updated)

	return nil
}

func removeMatchingStoredVersions(oldStoredVersions []string, versionRegexp *regexp.Regexp) []string {
	newStoredVersions := make([]string, 0, len(oldStoredVersions))

	for _, version := range oldStoredVersions {
		if versionRegexp.MatchString(version) {
			continue
		}

		newStoredVersions = append(newStoredVersions, version)
	}

	return newStoredVersions
}
