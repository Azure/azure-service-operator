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
	for _, crd := range list.Items {
		crd := crd

		if !crdRegexp.MatchString(crd.Name) {
			continue
		}

		newStoredVersions := removeMatchingStoredVersions(crd.Status.StoredVersions, deprecatedVersionRegexp)

		if len(newStoredVersions) > 0 && len(newStoredVersions) != len(crd.Status.StoredVersions) {
			crd.Status.StoredVersions = newStoredVersions
			// It's fine to update the storedVersions and remove the deprecated versions this way.
			// Users can still use the existing old v1alpha1api versioned resources and would not require to migrate
			// due to the conversion webhook implemented.
			updatedCrd, err := cl.UpdateStatus(ctx, &crd, v1.UpdateOptions{})
			if err != nil {
				return err
			}

			updated++
			fmt.Printf("updated '%s' CRD status storedVersions to : %s\n", crd.Name, updatedCrd.Status.StoredVersions)
		}
	}

	fmt.Printf("updated %d CRD(s)\n", updated)

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
