/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	azruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/tools/generator/pkg/naming"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	"strings"
)

type ARMResourceImporterFactory interface {
	CreateForARMID(armID string) (resourceImporter, error)
	Client() *azruntime.Pipeline
	Config() cloud.ServiceConfiguration
}

type armResourceImporterFactory struct {
	resourceImporterFactory
	armClient *azruntime.Pipeline
	armConfig cloud.ServiceConfiguration
}

var _ ARMResourceImporterFactory = &armResourceImporterFactory{}

// CreateForArmId creates a resourceImporter for the specified ARM ID
func (f *armResourceImporterFactory) CreateForARMID(armID string) (resourceImporter, error) {
	obj, err := f.createBlankObjectFromARMID(armID)
	if err != nil {
		return nil, err
	}

	armMeta, ok := obj.(genruntime.ARMMetaObject)
	if !ok {
		return nil, errors.Errorf(
			"unable to create blank resource, expected %s to identify an ARM object", armID)
	}

	return newARMResourceImporter(armID, armMeta, f, f.armClient, f.armConfig), nil
}

func (f *armResourceImporterFactory) Client() *azruntime.Pipeline {
	return f.armClient
}

func (f *armResourceImporterFactory) Config() cloud.ServiceConfiguration {
	return f.armConfig
}

func (f *armResourceImporterFactory) createBlankObjectFromARMID(armID string) (runtime.Object, error) {
	gvk, err := f.groupVersionKindFromARMID(armID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get GVK for blank resource")
	}

	obj, err := f.createBlankObjectFromGVK(gvk)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create blank resource")
	}

	if mo, ok := obj.(genruntime.ARMMetaObject); ok {
		name, err := f.nameFromARMID(armID)
		if err != nil {
			return nil, errors.Wrap(err, "unable to get name for blank resource")
		}

		mo.SetName(name)
	}

	return obj, nil
}

// groupVersionKindFromARMID returns the GroupVersionKind for the resource we're importing
func (f *armResourceImporterFactory) groupVersionKindFromARMID(armID string) (schema.GroupVersionKind, error) {
	gk, err := f.groupKindFromARMID(armID)
	if err != nil {
		return schema.GroupVersionKind{},
			errors.Wrap(err, "unable to determine GroupVersionKind for the resource")
	}

	return f.selectVersionFromGK(gk)
}

// groupKindFromARMID parses a GroupKind from the resource URL, allowing us to look up the actual resource
func (f *armResourceImporterFactory) groupKindFromARMID(armID string) (schema.GroupKind, error) {
	id, err := f.resourceIdFromARMID(armID)
	if err != nil {
		return schema.GroupKind{},
			errors.Wrap(err, "unable to parse GroupKind")
	}

	return schema.GroupKind{
		Group: f.groupFromID(id),
		Kind:  f.kindFromID(id),
	}, nil
}

// resourceIdFromARMID parses an ARM ID from the supplied resource path
func (f *armResourceImporterFactory) resourceIdFromARMID(armID string) (*arm.ResourceID, error) {
	id, err := arm.ParseResourceID(armID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse ARM ID from path %s", armID)
	}

	return id, nil
}

// groupFromID extracts an ASO group name from the ARM ID
func (*armResourceImporterFactory) groupFromID(id *arm.ResourceID) string {
	parts := strings.Split(id.ResourceType.Namespace, ".")
	last := len(parts) - 1
	group := strings.ToLower(parts[last]) + ".azure.com"
	klog.V(3).Infof("Group: %s", group)
	return group
}

// kindFromID extracts an ASO kind from the ARM ID
func (*armResourceImporterFactory) kindFromID(id *arm.ResourceID) string {
	if len(id.ResourceType.Types) != 1 {
		panic("Don't currently know how to handle nested resources")
	}

	kind := naming.Singularize(id.ResourceType.Types[0])
	klog.V(3).Infof("Kind: %s", kind)
	return kind
}

func (f *armResourceImporterFactory) nameFromARMID(armID string) (string, error) {
	id, err := f.resourceIdFromARMID(armID)
	if err != nil {
		return "", errors.Wrap(err, "unable to parse name")
	}

	klog.V(3).Infof("Name: %s", id.Name)
	return id.Name, nil
}
