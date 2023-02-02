/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import (
	"context"
	"regexp"
	"time"

	"github.com/Azure/azure-service-operator/v2/internal/logging"
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
	migrationBackoff := wait.Backoff{ // TODO: Still need to see if we want exponential backoff or linear is fine
		Duration: 2 * time.Second, // wait 2s between attempts, this will help us in a state of conflict.
		Steps:    3,               // 3 retry on error attempts per object
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

		newStoredVersions := removeMatchingStoredVersions(crd.Status.StoredVersions, deprecatedVersionRegexp)

		if len(newStoredVersions) > 0 && len(newStoredVersions) != len(crd.Status.StoredVersions) {
			klog.Infof("starting cleanup for '%s'", crd.Name)
			objectsToMigrate, err := c.getObjectsForMigration(ctx, crd, deprecatedVersionRegexp)
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
		} else {
			klog.Infof("nothing to update for '%s'\n", crd.Name)
		}
	}

	if !c.dryRun {
		klog.Infof("updated %d CRD(s)\n", updated)

	}

	return nil
}

func (c *Cleaner) updateStorageVersions(
	ctx context.Context,
	crd apiextensions.CustomResourceDefinition,
	newStoredVersions []string) error {

	if c.dryRun {
		klog.Infof("new storeVersions for '%s' CRD status storedVersions: %s\n", crd.Name, newStoredVersions)
		return nil
	}

	crd.Status.StoredVersions = newStoredVersions
	updatedCrd, err := c.apiExtensionsClient.UpdateStatus(ctx, &crd, v1.UpdateOptions{})
	if err != nil {
		return err
	}
	klog.Infof("updated '%s' CRD status storedVersions to : %s\n", crd.Name, updatedCrd.Status.StoredVersions)

	return nil
}

func (c *Cleaner) migrateObjects(ctx context.Context, objectsToMigrate *unstructured.UnstructuredList) error {
	for _, obj := range objectsToMigrate.Items {
		obj := obj
		if c.dryRun {
			// TODO: verbose
			klog.V(logging.Verbose).Infof("resource '%s' to migrate for kind '%s'", obj.GetName(), obj.GroupVersionKind().Kind)
			continue
		}

		err := retry.OnError(c.migrationBackoff, isErrorFatal, func() error { return c.client.Update(ctx, &obj) })
		if isErrorFatal(err) {
			return err
		}
		klog.V(logging.Verbose).Infof("migrated '%s' for %s\n", obj.GetName(), obj.GroupVersionKind().Kind)
	}

	return nil
}

func isErrorFatal(err error) bool {
	if err == nil {
		return false
	}

	if apierrors.IsGone(err) { // If resource no longer exists, we don't want to retry
		return false
	} else if apierrors.IsConflict(err) { // If resource is already in the state of update, we don't want to retry either
		return false
	} else {
		return true
	}
}

func (c *Cleaner) getObjectsForMigration(ctx context.Context, crd apiextensions.CustomResourceDefinition, versionRegexp *regexp.Regexp) (*unstructured.UnstructuredList, error) {
	list := &unstructured.UnstructuredList{}

	if ok, version := requireMigrate(crd, versionRegexp); ok {
		list.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   crd.Spec.Group,
			Version: version,
			Kind:    crd.Spec.Names.ListKind,
		})

		if err := c.client.List(ctx, list); err != nil {
			return nil, err
		}
	}

	return list, nil
}

// requireMigrate checks if resources under CRD need migration. Requirement for migration would depend on if the deprecated version is the storage version.
func requireMigrate(crd apiextensions.CustomResourceDefinition, versionRegexp *regexp.Regexp) (bool, string) {
	for _, version := range crd.Spec.Versions {
		if version.Storage && versionRegexp.MatchString(version.Name) {
			return true, version.Name
		}
	}
	return false, ""
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
