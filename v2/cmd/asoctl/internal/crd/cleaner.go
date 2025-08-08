/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package crd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/internal/crdmanagement"
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
	log logr.Logger,
) *Cleaner {
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
	if c.dryRun {
		c.log.Info("Starting update (dry run)")
	} else {
		c.log.Info("Starting update")
	}

	appLabelRequirement, err := labels.NewRequirement(crdmanagement.ServiceOperatorAppLabel, selection.Equals, []string{crdmanagement.ServiceOperatorAppValue})
	if err != nil {
		return err
	}
	selector := labels.NewSelector()
	selector = selector.Add(*appLabelRequirement)
	crdsWithNewLabel, err := c.apiExtensionsClient.List(ctx, v1.ListOptions{LabelSelector: selector.String()})
	if err != nil {
		return eris.Wrap(err, "failed to list CRDs")
	}

	versionLabelRequirement, err := labels.NewRequirement(crdmanagement.ServiceOperatorVersionLabelOld, selection.Exists, []string{})
	if err != nil {
		return err
	}
	selector = labels.NewSelector()
	selector = selector.Add(*versionLabelRequirement)
	crdsWithOldLabel, err := c.apiExtensionsClient.List(ctx, v1.ListOptions{LabelSelector: selector.String()})
	if err != nil {
		return eris.Wrap(err, "failed to list CRDs")
	}

	var crds []apiextensions.CustomResourceDefinition
	crds = append(crds, crdsWithNewLabel.Items...)
	crds = append(crds, crdsWithOldLabel.Items...)

	var updated int
	var asoCRDsSeen int

	for _, crd := range crds {
		asoCRDsSeen++
		newStoredVersions, deprecatedVersions := crdmanagement.GetDeprecatedStorageVersions(crd)

		// If there is no new version found other than the matched version, we short circuit here, as there is no updated version found in the CRDs
		if len(newStoredVersions) <= 0 {
			// TODO: test?
			return eris.Errorf("it doesn't look like your version of ASO is one that supports deprecating CRD %q, versions %q. Have you upgraded ASO yet?", crd.Name, deprecatedVersions)
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

	if asoCRDsSeen <= 0 {
		return eris.New("found no Azure Service Operator CRDs, make sure you have ASO installed.")
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
	newStoredVersions []string,
) error {
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
			return eris.Wrap(err,
				fmt.Sprintf("migrating %q of kind %s", obj.GetName(), obj.GroupVersionKind().Kind))
		}

		if found {
			originalVersion = strings.Replace(originalVersion, "v1alpha1api", "v1beta", 1)
			err = unstructured.SetNestedField(obj.Object, originalVersion, originalVersionFieldPath...)
			if err != nil {
				return eris.Wrap(err,
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
		// that means the resource is already updated, and we don't have to do anything more.
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
