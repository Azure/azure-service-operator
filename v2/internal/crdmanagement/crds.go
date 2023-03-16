// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package crdmanagement

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/yaml"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
)

// ServiceOperatorVersionLabel is the label the CRDs have on them containing the ASO version. This value must match the value
// injected by config/crd/labels.yaml
const ServiceOperatorVersionLabel = "serviceoperator.azure.com/version"
const CRDLocation = "crds"

const certMgrInjectCAFromAnnotation = "cert-manager.io/inject-ca-from"

type Manager struct {
	logger     logr.Logger
	kubeClient kubeclient.Client

	crds []apiextensions.CustomResourceDefinition
}

func NewManager(logger logr.Logger, kubeClient kubeclient.Client) *Manager {
	return &Manager{
		logger:     logger,
		kubeClient: kubeClient,
	}
}

func (m *Manager) ListOperatorCRDs(ctx context.Context) ([]apiextensions.CustomResourceDefinition, error) {
	list := apiextensions.CustomResourceDefinitionList{}

	selector := labels.NewSelector()
	requirement, err := labels.NewRequirement(ServiceOperatorVersionLabel, selection.Exists, nil)
	if err != nil {
		return nil, err
	}
	selector = selector.Add(*requirement)

	match := client.MatchingLabelsSelector{
		Selector: selector,
	}
	err = m.kubeClient.List(ctx, &list, match)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list CRDs")
	}

	for _, crd := range list.Items {
		m.logger.V(Verbose).Info("Found an existing CRD", "CRD", crd.Name)
	}

	return list.Items, nil
}

func (m *Manager) LoadOperatorCRDs(path string, podNamespace string) ([]apiextensions.CustomResourceDefinition, error) {
	if len(m.crds) > 0 {
		// Nothing to do as they're already loaded. Pod has to restart for them to change
		return m.crds, nil
	}

	crds, err := m.loadCRDs(path)
	if err != nil {
		return nil, err
	}
	crds = m.fixCRDNamespaceRefs(crds, podNamespace)

	m.crds = crds
	return crds, nil
}

// FindMatchingCRDs finds the CRDs in "goal" that are in "existing" AND compare as equal according to the comparators with
// the corresponding CRD in "goal"
func (m *Manager) FindMatchingCRDs(
	existing []apiextensions.CustomResourceDefinition,
	goal []apiextensions.CustomResourceDefinition,
	comparators ...func(a apiextensions.CustomResourceDefinition, b apiextensions.CustomResourceDefinition) bool,
) map[string]apiextensions.CustomResourceDefinition {

	matching := make(map[string]apiextensions.CustomResourceDefinition)

	// Build a map so lookup is faster
	existingCRDs := make(map[string]apiextensions.CustomResourceDefinition, len(existing))
	for _, crd := range existing {
		existingCRDs[crd.Name] = crd
	}

	// Every goal CRD should exist and match an existing one
	for _, goalCRD := range goal {

		// Note that if the CRD is not found, we will get back a default initialized CRD.
		// We run the comparators on that as they may match, especially if the comparator is something like
		// "specs are not equal"
		existingCRD := existingCRDs[goalCRD.Name]

		// Deepcopy to ensure that modifications below don't persist
		existingCRD = *existingCRD.DeepCopy()
		goalCRD = *goalCRD.DeepCopy()

		equal := true
		for _, c := range comparators {
			if c(existingCRD, goalCRD) == false {
				equal = false
				break
			}
		}

		if equal {
			matching[goalCRD.Name] = goalCRD
		}
	}

	return matching
}

// FindNonMatchingCRDs finds the CRDs in "goal" that are not in "existing" OR are in "existing" but mismatch with the "goal"
// based on the comparator functions.
func (m *Manager) FindNonMatchingCRDs(
	existing []apiextensions.CustomResourceDefinition,
	goal []apiextensions.CustomResourceDefinition,
	comparators ...func(a apiextensions.CustomResourceDefinition, b apiextensions.CustomResourceDefinition) bool,
) map[string]apiextensions.CustomResourceDefinition {

	// Just invert the comparators and call FindMatchingCRDs
	invertedComparators := make([]func(a apiextensions.CustomResourceDefinition, b apiextensions.CustomResourceDefinition) bool, 0, len(comparators))
	for _, c := range comparators {
		c := c
		invertedComparators = append(
			invertedComparators,
			func(a apiextensions.CustomResourceDefinition, b apiextensions.CustomResourceDefinition) bool {
				return !c(a, b)
			})
	}

	return m.FindMatchingCRDs(existing, goal, invertedComparators...)
}

func (m *Manager) ApplyCRDs(ctx context.Context, path string, podNamespace string, patterns []string) error {
	goalCRDs, err := m.LoadOperatorCRDs(path, podNamespace)
	if err != nil {
		return err
	}
	m.logger.V(Info).Info("Goal CRDs", "count", len(goalCRDs))

	existingCRDs, err := m.ListOperatorCRDs(ctx)
	if err != nil {
		return err
	}
	m.logger.V(Info).Info("Existing CRDs", "count", len(existingCRDs))

	goalCRDsWithDifferentVersion := m.FindNonMatchingCRDs(existingCRDs, goalCRDs, VersionEqual)
	goalCRDsWithDifferentSpec := m.FindNonMatchingCRDs(existingCRDs, goalCRDs, SpecEqual)

	// The same CRD may be in both sets, but we don't want to apply twice, so combine the sets
	crdsToApply := make(map[string]apiextensions.CustomResourceDefinition)
	for name, crd := range goalCRDsWithDifferentVersion {
		crdsToApply[name] = crd
	}
	for name, crd := range goalCRDsWithDifferentSpec {
		crdsToApply[name] = crd
	}

	if len(crdsToApply) > 0 {
		for _, crd := range crdsToApply {
			crd := crd

			_, isDifferentVersion := goalCRDsWithDifferentVersion[crd.Name]
			_, isDifferentSpec := goalCRDsWithDifferentSpec[crd.Name]
			m.logger.V(Verbose).Info("Found CRD in need of update", "CRD", crd.Name, "SpecDifferent", isDifferentSpec, "VersionDifferent", isDifferentVersion)

			toApply := &apiextensions.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name: crd.Name,
				},
			}
			result, err := controllerutil.CreateOrUpdate(ctx, m.kubeClient, toApply, func() error {
				resourceVersion := toApply.ResourceVersion
				*toApply = crd
				toApply.ResourceVersion = resourceVersion

				return nil
			})
			if err != nil {
				return errors.Wrapf(err, "failed to apply CRD %s", crd.Name)
			}

			m.logger.V(Debug).Info("Successfully applied CRD", "name", crd.Name, "result", result)
		}

		// If we make it to here, we have successfully updated all the CRDs we needed to. We need to kill the pod and let it restart so
		// that the new shape CRDs can be reconciled.
		m.logger.V(Status).Info("Restarting operator pod after updating CRDs", "count", len(crdsToApply))
		os.Exit(0) // TODO: Should this be nonzero?
	}

	m.logger.V(Status).Info("Successfully reconciled InstalledResourceDefinitions CRDs.")
	return nil
}

func (m *Manager) loadCRDs(path string) ([]apiextensions.CustomResourceDefinition, error) {
	// Expectation is that every file in this folder is a CRD
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read directory %s", path)
	}

	results := make([]apiextensions.CustomResourceDefinition, 0, len(entries))

	for _, entry := range entries {
		if entry.IsDir() {
			continue // Ignore directories
		}

		filePath := filepath.Join(path, entry.Name())
		var content []byte
		content, err = os.ReadFile(filePath)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read %s", filePath)
		}

		crd := apiextensions.CustomResourceDefinition{}
		err = yaml.Unmarshal(content, &crd)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal %s to CRD", filePath)
		}

		m.logger.V(Verbose).Info("Loaded CRD", "path", filePath, "name", crd.Name)
		results = append(results, crd)
	}

	return results, nil
}

func (m *Manager) fixCRDNamespaceRefs(crds []apiextensions.CustomResourceDefinition, namespace string) []apiextensions.CustomResourceDefinition {
	results := make([]apiextensions.CustomResourceDefinition, 0, len(crds))

	for _, crd := range crds {
		crd = fixCRDNamespace(crd, namespace)
		results = append(results, crd)
	}

	return results
}

// fixCRDNamespace fixes up namespace references in the CRD to match the provided namespace.
// This could in theory be done with a string replace across the JSON representation of the CRD, but that's risky given
// we don't know what else might have the "azureserviceoperator-system" string in it. Instead, we hardcode specific places
// we know need to be fixed up. This is more brittle in the face of namespace additions but has the advantage of guaranteeing
// that we can't break our own CRDs with a string replace gone awry.
func fixCRDNamespace(crd apiextensions.CustomResourceDefinition, namespace string) apiextensions.CustomResourceDefinition {
	result := crd.DeepCopy()

	// Set spec.conversion.webhook.clientConfig.service.namespace
	if result.Spec.Conversion != nil &&
		result.Spec.Conversion.Webhook != nil &&
		result.Spec.Conversion.Webhook.ClientConfig != nil &&
		result.Spec.Conversion.Webhook.ClientConfig.Service != nil {
		result.Spec.Conversion.Webhook.ClientConfig.Service.Namespace = namespace
	}

	// Set cert-manager.io/inject-ca-from
	if len(result.Labels) > 0 {
		if injectCAFrom, ok := result.Labels[certMgrInjectCAFromAnnotation]; ok {
			split := strings.Split(injectCAFrom, "/")
			if len(split) == 2 {
				result.Labels[certMgrInjectCAFromAnnotation] = fmt.Sprintf("%s/%s", namespace, split[1])
			}
		}
	}

	return *result
}

func ignoreCABundle(a apiextensions.CustomResourceDefinition) apiextensions.CustomResourceDefinition {
	if a.Spec.Conversion != nil && a.Spec.Conversion.Webhook != nil &&
		a.Spec.Conversion.Webhook.ClientConfig != nil {
		a.Spec.Conversion.Webhook.ClientConfig.CABundle = nil
	}

	return a
}

func ignoreConversionWebhook(a apiextensions.CustomResourceDefinition) apiextensions.CustomResourceDefinition {
	if a.Spec.Conversion != nil && a.Spec.Conversion.Webhook != nil {
		a.Spec.Conversion.Webhook = nil
	}

	return a
}

func SpecEqual(a apiextensions.CustomResourceDefinition, b apiextensions.CustomResourceDefinition) bool {
	a = ignoreCABundle(a)
	b = ignoreCABundle(b)

	return reflect.DeepEqual(a.Spec, b.Spec)
}

func SpecEqualIgnoreConversionWebhook(a apiextensions.CustomResourceDefinition, b apiextensions.CustomResourceDefinition) bool {
	a = ignoreConversionWebhook(a)
	b = ignoreConversionWebhook(b)

	return reflect.DeepEqual(a.Spec, b.Spec)
}

func VersionEqual(a apiextensions.CustomResourceDefinition, b apiextensions.CustomResourceDefinition) bool {
	if a.Labels == nil && b.Labels == nil {
		return true
	}

	if a.Labels == nil || b.Labels == nil {
		return false
	}

	aVersion, aOk := a.Labels[ServiceOperatorVersionLabel]
	bVersion, bOk := b.Labels[ServiceOperatorVersionLabel]

	if !aOk && !bOk {
		return true
	}

	return aVersion == bVersion
}
