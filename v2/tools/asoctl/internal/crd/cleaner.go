/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import (
	"context"
	"fmt"
	"regexp"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/tools/asoctl/internal/client"
)

func CleanDeprecatedCRDVersions(ctx context.Context) error {
	cl, err := client.NewClient()
	if err != nil {
		return err
	}

	list, err := cl.CustomResourceDefinitions().List(ctx, v1.ListOptions{})
	if err != nil {
		return err
	}

	if list.Items == nil || len(list.Items) == 0 {
		fmt.Println("found 0 results, make sure you have ASO CRDs installed")
	}

	var updated bool
	for _, item := range list.Items {
		crdName := item.Name
		crdRegexp := regexp.MustCompile(".*\\.azure\\.com")
		deprecatedVersionRegexp := regexp.MustCompile("v1alpha1api.{8}.*")

		if !crdRegexp.MatchString(crdName) {
			continue
		}

		crd, err := cl.CustomResourceDefinitions().Get(ctx, crdName, v1.GetOptions{})
		if err != nil {
			return err
		}

		found, newStoredVersions := getNewVersionSet(crd.Status.StoredVersions, deprecatedVersionRegexp)

		if found {
			crd.Status.StoredVersions = newStoredVersions
			crd, err = cl.CustomResourceDefinitions().UpdateStatus(ctx, crd, v1.UpdateOptions{})
			updated = true
			if err != nil {
				return err
			}

			fmt.Printf("updated '%v' CRD status storedVersions to : %v\n", crdName, crd.Status.StoredVersions)
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
