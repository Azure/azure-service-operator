/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import (
	"context"
	"fmt"
	"regexp"

	apiextensionsV1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CleanDeprecatedCRDVersions(ctx context.Context, cl apiextensionsV1.CustomResourceDefinitionInterface) error {
	var crdRegexp = regexp.MustCompile(".*\\.azure\\.com")
	var deprecatedVersionRegexp = regexp.MustCompile("v1alpha1api.{8}.*")

	list, err := cl.List(ctx, v1.ListOptions{})
	if err != nil {
		return err
	}

	if list.Items == nil || len(list.Items) == 0 {
		fmt.Println("found 0 results, make sure you have ASO CRDs installed")
	}

	var updated bool
	for _, item := range list.Items {
		crdName := item.Name

		if !crdRegexp.MatchString(crdName) {
			continue
		}

		crd, err := cl.Get(ctx, crdName, v1.GetOptions{})
		if err != nil {
			return err
		}

		found, newStoredVersions := getNewVersionSet(crd.Status.StoredVersions, deprecatedVersionRegexp)

		if found {
			crd.Status.StoredVersions = newStoredVersions
			crd, err = cl.UpdateStatus(ctx, crd, v1.UpdateOptions{})
			if err != nil {
				return err
			}

			updated = true
			fmt.Printf("updated '%s' CRD status storedVersions to : %s\n", crdName, crd.Status.StoredVersions)
		}
	}

	if !updated {
		fmt.Println("updated 0 CRDs, no deprecated versions found")
	}

	return nil
}

func getNewVersionSet(oldStoredVersions []string, versionRegexp *regexp.Regexp) (bool, []string) {
	var newStoredVersions []string
	var found bool

	for _, version := range oldStoredVersions {
		if versionRegexp.MatchString(version) {
			found = true
			continue
		}

		newStoredVersions = append(newStoredVersions, version)
	}

	return found, newStoredVersions
}
