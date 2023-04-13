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

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Cleaner struct {
	apiExtensionsClient apiextensionsclient.CustomResourceDefinitionInterface
	client              client.Client
	migrationBackoff    wait.Backoff
	dryRun              bool
	log                 logr.Logger
}

func NewCleaner(
	apiExtensionsClient apiextensionsclient.CustomResourceDefinitionInterface,
	client client.Client,
	dryRun bool,
	log logr.Logger) *Cleaner {
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
		log:                 log,
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
			c.log.Info(
				"Nothing to update",
				"crd-name", crd.Name)
			continue
		}

		// Make sure to use a version that hasn't been deprecated for migration. Deprecated versions will not be in our
		// scheme, and so we cannot List/PUT with them. Instead, use the next available version.
		activeVersion := newStoredVersions[len(newStoredVersions)-1]
		c.log.Info(
			"Starting cleanup",
			"crd-name", crd.Name)

		objectsToMigrate, err := c.getObjectsForMigration(ctx, crd, activeVersion)
		if err != nil {
			return err
		}

		err = c.migrateObjects(ctx, objectsToMigrate, activeVersion)
		if err != nil {
			return err
		}

		err = c.updateStorageVersions(ctx, crd, newStoredVersions)
		if err != nil {
			return err
		}

		updated++
	}

	if c.dryRun {
		c.log.Info("Update finished (dry run)")
	} else {
		c.log.Info(
			"Update finished",
			"crd-count", updated)
	}

	return nil
}

func (c *Cleaner) updateStorageVersions(
	ctx context.Context,
	crd apiextensions.CustomResourceDefinition,
	newStoredVersions []string) error {

	if c.dryRun {
		c.log.Info(
			"Would update storedVersions",
			"crd-name", crd.Name,
			"storedVersions", newStoredVersions)
		return nil
	}

	crd.Status.StoredVersions = newStoredVersions
	updatedCrd, err := c.apiExtensionsClient.UpdateStatus(ctx, &crd, v1.UpdateOptions{})
	if err != nil {
		return err
	}
	c.log.Info(
		"Updated CRD status storedVersions",
		"crd-name", crd.Name,
		"storedVersions", updatedCrd.Status.StoredVersions)

	return nil
}

func (c *Cleaner) migrateObjects(ctx context.Context, objectsToMigrate *unstructured.UnstructuredList) error {
	for _, obj := range objectsToMigrate.Items {
		obj := obj
		if c.dryRun {
			c.log.Info(
				"Would migrate resource",
				"name", obj.GetName(),
				"kind", obj.GroupVersionKind().Kind)
			continue
		}

		originalVersionFieldPath := []string{"spec", "originalVersion"}

		originalVersion, found, err := unstructured.NestedString(obj.Object, originalVersionFieldPath...)
		if err != nil {
			return errors.Wrap(err,
				fmt.Sprintf("migrating %q of kind %s", obj.GetName(), obj.GroupVersionKind().Kind))
		}

		if found {
			originalVersion = strings.Replace(originalVersion, "v1alpha1api", "v1beta", 1)
			err = unstructured.SetNestedField(obj.Object, originalVersion, originalVersionFieldPath...)
			if err != nil {
				return errors.Wrap(err,
					fmt.Sprintf("migrating %q of kind %s", obj.GetName(), obj.GroupVersionKind().Kind))
			}
		} else {
			// If we don't find the originalVersion, it may not have been set.
			// This can happen for some resources such as ResourceGroup which were handcrafted in versions prior to v2.0.0 and thus didn't have a StorageVersion.
			c.log.Info(
				"originalVersion not found. Continuing with the latest.",
				"name", obj.GetName(),
				"kind", obj.GroupVersionKind().Kind)
		}

		err = retry.OnError(c.migrationBackoff, isErrorFatal, func() error { return c.client.Update(ctx, &obj) })
		if isErrorFatal(err) {
			return err
		}

		c.log.Info(
			"Migrated resource",
			"name", obj.GetName(),
			"kind", obj.GroupVersionKind().Kind)
	}

	c.log.Info(
		"Migration finished",
		"resource-count", len(objectsToMigrate.Items))

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
