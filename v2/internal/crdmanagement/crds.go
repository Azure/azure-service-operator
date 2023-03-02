// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package crdmanagement

import (
	"context"
	"os"
	"path/filepath"
	"reflect"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/version"
)

const ServiceOperatorVersionLabel = "azure.microsoft.com/aso-version"
const CRDLocation = "crds"

// TODO: Fix logging levels in this whole file
// TODO: The label selector below seems to not be working...

type Manager struct {
	logger     logr.Logger
	kubeClient kubeclient.Client
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
	//err := kubeClient.List(ctx, &list)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list CRDs")
	}

	for _, crd := range list.Items {
		m.logger.V(0).Info("Found a CRD", "CRD", crd.Name)
	}

	return list.Items, nil
}

func (m *Manager) LoadOperatorCRDs(path string) ([]apiextensions.CustomResourceDefinition, error) {
	// Expectation is that every file in this folder is a CRD
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read directory %s", path)
	}

	var results []apiextensions.CustomResourceDefinition

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

		m.logger.V(0).Info("Loaded CRD", "path", filePath, "name", crd.Name)
		results = append(results, crd)
	}

	return results, nil
}

func (m *Manager) FindGoalCRDsNeedingUpdate(
	existing []apiextensions.CustomResourceDefinition,
	goal []apiextensions.CustomResourceDefinition,
	comparators ...func(a apiextensions.CustomResourceDefinition, b apiextensions.CustomResourceDefinition) bool,
) map[string]apiextensions.CustomResourceDefinition {

	needUpdate := make(map[string]apiextensions.CustomResourceDefinition)

	// Build a map so lookup is faster
	existingCRDs := make(map[string]apiextensions.CustomResourceDefinition, len(existing))
	for _, crd := range existing {
		existingCRDs[crd.Name] = crd
	}

	// Every goal CRD should exist and match an existing one
	for _, goalCRD := range goal {
		existingCRD, ok := existingCRDs[goalCRD.Name]
		if !ok {
			// Not found
			needUpdate[goalCRD.Name] = goalCRD
			m.logger.V(0).Info("Found missing CRD", "CRD", goalCRD.Name)
			continue
		}

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

		if !equal {
			m.logger.V(0).Info("Goal CRD does not match existing CRD", "CRD", goalCRD.Name)
			needUpdate[goalCRD.Name] = goalCRD
		}
	}

	return needUpdate
}

func ApplyServiceOperatorVersionLabel(crds []apiextensions.CustomResourceDefinition) []apiextensions.CustomResourceDefinition {
	results := make([]apiextensions.CustomResourceDefinition, 0, len(crds))

	for _, crd := range crds {
		// Apply our version label to it
		if crd.Labels == nil {
			crd.Labels = make(map[string]string)
		}
		crd.Labels[ServiceOperatorVersionLabel] = version.BuildVersion

		results = append(results, crd)
	}

	return results
}

func ignoreCABundle(a apiextensions.CustomResourceDefinition) apiextensions.CustomResourceDefinition {
	if a.Spec.Conversion != nil && a.Spec.Conversion.Webhook != nil &&
		a.Spec.Conversion.Webhook.ClientConfig != nil {
		a.Spec.Conversion.Webhook.ClientConfig.CABundle = nil
	}

	return a
}

func SpecEqual(a apiextensions.CustomResourceDefinition, b apiextensions.CustomResourceDefinition) bool {
	a = ignoreCABundle(a)
	b = ignoreCABundle(b)

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
