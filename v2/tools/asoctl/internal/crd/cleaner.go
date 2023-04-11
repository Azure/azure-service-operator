/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/pkg/errors"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Cleaner struct {
	apiExtensionsClient apiextensionsclient.CustomResourceDefinitionInterface
	client              client.Client
	migrationBackoff    wait.Backoff
	dryRun              bool
}

func NewCleaner(apiExtensionsClient apiextensionsclient.CustomResourceDefinitionInterface, client client.Client, dryRun bool) *Cleaner {
	migrationBackoff := wait.Backoff{
		Duration: 2 * time.Second, // wait 2s between attempts, this will help us in a state of conflict.
		Steps:    3,               // 3 retry on error attempts per object
		Jitter:   0.1,             // Jitter 0.1*duration
	}

	return &Cleaner{
		apiExtensionsClient: apiExtensionsClient,
		client:              client,
		migrationBackoff:    migrationBackoff,
		dryRun:              dryRun,
	}
}

func (c *Cleaner) Run(ctx context.Context) error {
	list, err := c.apiExtensionsClient.List(ctx, v1.ListOptions{})
	if err != nil {
		return errors.Wrap(err, "failed to list CRDs")
	}

	if list == nil || len(list.Items) == 0 {
		return errors.New("found 0 results, make sure you have ASO CRDs installed")

	}

	var updated int
	crdRegexp := regexp.MustCompile(`.*\.azure\.com`)
	deprecatedVersionRegexp := regexp.MustCompile(`v1alpha1api\d{8}(preview)?(storage)?`)
	for _, crd := range list.Items {
		crd := crd

		if !crdRegexp.MatchString(crd.Name) {
			continue
		}

		newStoredVersions, deprecatedVersion := removeMatchingStoredVersions(crd.Status.StoredVersions, deprecatedVersionRegexp)

		// If there is no new version found other than the matched version, we short circuit here, as there is no updated version found in the CRDs
		if len(newStoredVersions) <= 0 {
			return errors.New(fmt.Sprintf("it doesn't look like your version of ASO is one that supports deprecating version %q. Have you upgraded ASO yet?", deprecatedVersion))
		}

		// If the slice was not updated, there is no version to deprecate.
		if len(newStoredVersions) == len(crd.Status.StoredVersions) {
			klog.Infof("Nothing to update for %q\n", crd.Name)
			continue
		}

		// Make sure to use a version that hasn't been deprecated for migration. Deprecated versions will not be in our
		// scheme, and so we cannot List/PUT with them. Instead, use the next available version.
		// TODO: We need to do a better job of selecting a version to use here. If we're not careful, we could
		// TODO: issue a GET + PUT with an older Azure API version and end up losing/removing some properties.
		// TODO: The ideal algorithm would be:
		// TODO: 1. Use storage version to list all CRs. Extract the OriginalGVK field
		// TODO: 2. Swap v1alpha1 -> v1beta1 (for alpha deprecation) and save that as versionToUse for that CR
		// TODO: 3. Issue GET + PUT with versionToUse
		// TODO: Doing the above is tricky though so for now we'll just use the latest stored version
		activeVersion := getVersionFromStoredVersion(newStoredVersions[len(newStoredVersions)-1])
		klog.Infof("Starting cleanup for %q", crd.Name)
		objectsToMigrate, err := c.getObjectsForMigration(ctx, crd, activeVersion)
		if err != nil {
			return err
		}

		err = c.migrateObjects(ctx, objectsToMigrate)
		if err != nil {
			return err
		}

		err = c.updateStorageVersions(ctx, crd, newStoredVersions)
		if err != nil {
			return err
		}

		updated++
	}

	if !c.dryRun {
		klog.Infof("Updated %d CRD(s)\n", updated)

	}

	return nil
}

func (c *Cleaner) updateStorageVersions(
	ctx context.Context,
	crd apiextensions.CustomResourceDefinition,
	newStoredVersions []string) error {

	if c.dryRun {
		klog.Infof("Would update storedVersions for %q CRD to: %s\n", crd.Name, newStoredVersions)
		return nil
	}

	crd.Status.StoredVersions = newStoredVersions
	updatedCrd, err := c.apiExtensionsClient.UpdateStatus(ctx, &crd, v1.UpdateOptions{})
	if err != nil {
		return err
	}
	klog.Infof("Updated %q CRD status storedVersions to : %s\n", crd.Name, updatedCrd.Status.StoredVersions)

	return nil
}

func (c *Cleaner) migrateObjects(ctx context.Context, objectsToMigrate *unstructured.UnstructuredList) error {
	for _, obj := range objectsToMigrate.Items {
		obj := obj
		if c.dryRun {
			klog.V(2).Infof("Would migrate resource %q of kind %q", obj.GetName(), obj.GroupVersionKind().Kind)
			continue
		}

		err := retry.OnError(c.migrationBackoff, isErrorFatal, func() error { return c.client.Update(ctx, &obj) })
		if isErrorFatal(err) {
			return err
		}

		klog.V(2).Infof("Migrated %q of kind %s", obj.GetName(), obj.GroupVersionKind().Kind)
	}

	klog.Infof("Migrated %d resources\n", len(objectsToMigrate.Items))
	return nil
}

func isErrorFatal(err error) bool {
	if err == nil {
		return false
	}

	if apierrors.IsGone(err) { // If resource no longer exists, we don't want to retry
		return false
	} else if apierrors.IsConflict(err) {
		// If resource is already in the state of update, we don't want to retry either.
		// Since, we're also updating resources to achieve version migration, and if we see a conflict in update,
		// that means the resource is already updated and we don't have to do anything more.
		return false
	} else {
		return true
	}
}

func (c *Cleaner) getObjectsForMigration(ctx context.Context, crd apiextensions.CustomResourceDefinition, version string) (*unstructured.UnstructuredList, error) {
	list := &unstructured.UnstructuredList{}

	list.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   crd.Spec.Group,
		Version: version,
		Kind:    crd.Spec.Names.ListKind,
	})

	if err := c.client.List(ctx, list); err != nil {
		return nil, err
	}

	return list, nil
}

// removeMatchingStoredVersions returns a new list of storedVersions by removing the non-storage matched version
func removeMatchingStoredVersions(oldVersions []string, versionRegexp *regexp.Regexp) ([]string, string) {
	newStoredVersions := make([]string, 0, len(oldVersions))
	var matchedStoredVersion string
	for _, version := range oldVersions {
		if versionRegexp.MatchString(version) {
			matchedStoredVersion = version
			continue
		}

		newStoredVersions = append(newStoredVersions, version)
	}

	return newStoredVersions, matchedStoredVersion
}

// getVersionFromStoredVersion returns the public (non-storage) API version for a given version
func getVersionFromStoredVersion(version string) string {
	result := strings.TrimSuffix(version, "storage")
	return result
}
