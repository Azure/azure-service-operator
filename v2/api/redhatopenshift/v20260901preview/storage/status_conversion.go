/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	hubstorage "github.com/Azure/azure-service-operator/v2/api/redhatopenshift/v20261001/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const activeVersionsProperty = "activeVersions"

func (properties *HcpOpenShiftClusterProperties_STATUS) AssignPropertiesFrom(src *hubstorage.HcpOpenShiftClusterProperties_STATUS) error {
	if src.Status == nil {
		properties.Status = nil
		return nil
	}

	status := &ResourceStatus_STATUS{
		Conditions:  conditionsFromHub(src.Status.Conditions),
		PropertyBag: genruntime.NewPropertyBag(src.Status.PropertyBag),
	}
	if src.Status.ActiveVersions != nil {
		if err := status.PropertyBag.Add(activeVersionsProperty, src.Status.ActiveVersions); err != nil {
			return err
		}
	}
	properties.Status = status
	return nil
}

func (properties *HcpOpenShiftClusterProperties_STATUS) AssignPropertiesTo(dst *hubstorage.HcpOpenShiftClusterProperties_STATUS) error {
	if properties.Status == nil {
		dst.Status = nil
		return nil
	}

	status := &hubstorage.ClusterResourceStatus_STATUS{
		Conditions:  conditionsToHub(properties.Status.Conditions),
		PropertyBag: genruntime.NewPropertyBag(properties.Status.PropertyBag),
	}
	if status.PropertyBag.Contains(activeVersionsProperty) {
		if err := status.PropertyBag.Pull(activeVersionsProperty, &status.ActiveVersions); err != nil {
			return err
		}
	}
	dst.Status = status
	return nil
}

func (properties *NodePoolProperties_STATUS) AssignPropertiesFrom(src *hubstorage.NodePoolProperties_STATUS) error {
	if src.Status == nil {
		properties.Status = nil
		return nil
	}

	status := &ResourceStatus_STATUS{
		Conditions:  conditionsFromHub(src.Status.Conditions),
		PropertyBag: genruntime.NewPropertyBag(src.Status.PropertyBag),
	}
	if src.Status.ActiveVersions != nil {
		if err := status.PropertyBag.Add(activeVersionsProperty, src.Status.ActiveVersions); err != nil {
			return err
		}
	}
	properties.Status = status
	return nil
}

func (properties *NodePoolProperties_STATUS) AssignPropertiesTo(dst *hubstorage.NodePoolProperties_STATUS) error {
	if properties.Status == nil {
		dst.Status = nil
		return nil
	}

	status := &hubstorage.NodePoolResourceStatus_STATUS{
		Conditions:  conditionsToHub(properties.Status.Conditions),
		PropertyBag: genruntime.NewPropertyBag(properties.Status.PropertyBag),
	}
	if status.PropertyBag.Contains(activeVersionsProperty) {
		if err := status.PropertyBag.Pull(activeVersionsProperty, &status.ActiveVersions); err != nil {
			return err
		}
	}
	dst.Status = status
	return nil
}

func (properties *ExternalAuthProperties_STATUS) AssignPropertiesFrom(src *hubstorage.ExternalAuthProperties_STATUS) error {
	if src.Status == nil {
		properties.Status = nil
		return nil
	}

	properties.Status = &ResourceStatus_STATUS{
		Conditions:  conditionsFromHub(src.Status.Conditions),
		PropertyBag: genruntime.NewPropertyBag(src.Status.PropertyBag),
	}
	return nil
}

func (properties *ExternalAuthProperties_STATUS) AssignPropertiesTo(dst *hubstorage.ExternalAuthProperties_STATUS) error {
	if properties.Status == nil {
		dst.Status = nil
		return nil
	}

	dst.Status = &hubstorage.ExternalAuthResourceStatus_STATUS{
		Conditions:  conditionsToHub(properties.Status.Conditions),
		PropertyBag: genruntime.NewPropertyBag(properties.Status.PropertyBag),
	}
	return nil
}

func conditionsFromHub(source []hubstorage.Condition_STATUS) []Condition_STATUS {
	if source == nil {
		return nil
	}

	result := make([]Condition_STATUS, len(source))
	for i := range source {
		result[i] = Condition_STATUS{
			LastTransitionTime: genruntime.ClonePointerToString(source[i].LastTransitionTime),
			Message:            genruntime.ClonePointerToString(source[i].Message),
			PropertyBag:        genruntime.NewPropertyBag(source[i].PropertyBag),
			Reason:             genruntime.ClonePointerToString(source[i].Reason),
			Status:             genruntime.ClonePointerToString(source[i].Status),
			Type:               genruntime.ClonePointerToString(source[i].Type),
		}
	}
	return result
}

func conditionsToHub(source []Condition_STATUS) []hubstorage.Condition_STATUS {
	if source == nil {
		return nil
	}

	result := make([]hubstorage.Condition_STATUS, len(source))
	for i := range source {
		result[i] = hubstorage.Condition_STATUS{
			LastTransitionTime: genruntime.ClonePointerToString(source[i].LastTransitionTime),
			Message:            genruntime.ClonePointerToString(source[i].Message),
			PropertyBag:        genruntime.NewPropertyBag(source[i].PropertyBag),
			Reason:             genruntime.ClonePointerToString(source[i].Reason),
			Status:             genruntime.ClonePointerToString(source[i].Status),
			Type:               genruntime.ClonePointerToString(source[i].Type),
		}
	}
	return result
}
