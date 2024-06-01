// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package compat

import (
	storage "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/pkg/errors"
)

// Storage version of v1api20230202preview.ServiceMeshProfile_STATUS
// Service mesh profile for a managed cluster.
type ServiceMeshProfile_STATUS struct {
	Istio       *IstioServiceMesh_STATUS `json:"istio,omitempty"`
	Mode        *string                  `json:"mode,omitempty"`
	PropertyBag genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_ServiceMeshProfile_STATUS populates our ServiceMeshProfile_STATUS from the provided source ServiceMeshProfile_STATUS
func (profile *ServiceMeshProfile_STATUS) AssignProperties_From_ServiceMeshProfile_STATUS(source *storage.ServiceMeshProfile_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Istio
	if source.Istio != nil {
		var istio IstioServiceMesh_STATUS
		err := istio.AssignProperties_From_IstioServiceMesh_STATUS(source.Istio)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_IstioServiceMesh_STATUS() to populate field Istio")
		}
		profile.Istio = &istio
	} else {
		profile.Istio = nil
	}

	// Mode
	profile.Mode = genruntime.ClonePointerToString(source.Mode)

	// Update the property bag
	if len(propertyBag) > 0 {
		profile.PropertyBag = propertyBag
	} else {
		profile.PropertyBag = nil
	}

	// Invoke the augmentConversionForServiceMeshProfile_STATUS interface (if implemented) to customize the conversion
	var profileAsAny any = profile
	if augmentedProfile, ok := profileAsAny.(augmentConversionForServiceMeshProfile_STATUS); ok {
		err := augmentedProfile.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_ServiceMeshProfile_STATUS populates the provided destination ServiceMeshProfile_STATUS from our ServiceMeshProfile_STATUS
func (profile *ServiceMeshProfile_STATUS) AssignProperties_To_ServiceMeshProfile_STATUS(destination *storage.ServiceMeshProfile_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(profile.PropertyBag)

	// Istio
	if profile.Istio != nil {
		var istio storage.IstioServiceMesh_STATUS
		err := profile.Istio.AssignProperties_To_IstioServiceMesh_STATUS(&istio)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_IstioServiceMesh_STATUS() to populate field Istio")
		}
		destination.Istio = &istio
	} else {
		destination.Istio = nil
	}

	// Mode
	destination.Mode = genruntime.ClonePointerToString(profile.Mode)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForServiceMeshProfile_STATUS interface (if implemented) to customize the conversion
	var profileAsAny any = profile
	if augmentedProfile, ok := profileAsAny.(augmentConversionForServiceMeshProfile_STATUS); ok {
		err := augmentedProfile.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForServiceMeshProfile_STATUS interface {
	AssignPropertiesFrom(src *storage.ServiceMeshProfile_STATUS) error
	AssignPropertiesTo(dst *storage.ServiceMeshProfile_STATUS) error
}

// Storage version of v1api20230202preview.IstioServiceMesh_STATUS
// Istio service mesh configuration.
type IstioServiceMesh_STATUS struct {
	Components  *IstioComponents_STATUS `json:"components,omitempty"`
	PropertyBag genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_IstioServiceMesh_STATUS populates our IstioServiceMesh_STATUS from the provided source IstioServiceMesh_STATUS
func (mesh *IstioServiceMesh_STATUS) AssignProperties_From_IstioServiceMesh_STATUS(source *storage.IstioServiceMesh_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// CertificateAuthority
	if source.CertificateAuthority != nil {
		propertyBag.Add("CertificateAuthority", *source.CertificateAuthority)
	} else {
		propertyBag.Remove("CertificateAuthority")
	}

	// Components
	if source.Components != nil {
		var component IstioComponents_STATUS
		err := component.AssignProperties_From_IstioComponents_STATUS(source.Components)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_IstioComponents_STATUS() to populate field Components")
		}
		mesh.Components = &component
	} else {
		mesh.Components = nil
	}

	// Revisions
	if len(source.Revisions) > 0 {
		propertyBag.Add("Revisions", source.Revisions)
	} else {
		propertyBag.Remove("Revisions")
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		mesh.PropertyBag = propertyBag
	} else {
		mesh.PropertyBag = nil
	}

	// Invoke the augmentConversionForIstioServiceMesh_STATUS interface (if implemented) to customize the conversion
	var meshAsAny any = mesh
	if augmentedMesh, ok := meshAsAny.(augmentConversionForIstioServiceMesh_STATUS); ok {
		err := augmentedMesh.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_IstioServiceMesh_STATUS populates the provided destination IstioServiceMesh_STATUS from our IstioServiceMesh_STATUS
func (mesh *IstioServiceMesh_STATUS) AssignProperties_To_IstioServiceMesh_STATUS(destination *storage.IstioServiceMesh_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(mesh.PropertyBag)

	// CertificateAuthority
	if propertyBag.Contains("CertificateAuthority") {
		var certificateAuthority storage.IstioCertificateAuthority_STATUS
		err := propertyBag.Pull("CertificateAuthority", &certificateAuthority)
		if err != nil {
			return errors.Wrap(err, "pulling 'CertificateAuthority' from propertyBag")
		}

		destination.CertificateAuthority = &certificateAuthority
	} else {
		destination.CertificateAuthority = nil
	}

	// Components
	if mesh.Components != nil {
		var component storage.IstioComponents_STATUS
		err := mesh.Components.AssignProperties_To_IstioComponents_STATUS(&component)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_IstioComponents_STATUS() to populate field Components")
		}
		destination.Components = &component
	} else {
		destination.Components = nil
	}

	// Revisions
	if propertyBag.Contains("Revisions") {
		var revision []string
		err := propertyBag.Pull("Revisions", &revision)
		if err != nil {
			return errors.Wrap(err, "pulling 'Revisions' from propertyBag")
		}

		destination.Revisions = revision
	} else {
		destination.Revisions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForIstioServiceMesh_STATUS interface (if implemented) to customize the conversion
	var meshAsAny any = mesh
	if augmentedMesh, ok := meshAsAny.(augmentConversionForIstioServiceMesh_STATUS); ok {
		err := augmentedMesh.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForIstioServiceMesh_STATUS interface {
	AssignPropertiesFrom(src *storage.IstioServiceMesh_STATUS) error
	AssignPropertiesTo(dst *storage.IstioServiceMesh_STATUS) error
}

// Storage version of v1api20230202preview.IstioComponents_STATUS
// Istio components configuration.
type IstioComponents_STATUS struct {
	IngressGateways []IstioIngressGateway_STATUS `json:"ingressGateways,omitempty"`
	PropertyBag     genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_IstioComponents_STATUS populates our IstioComponents_STATUS from the provided source IstioComponents_STATUS
func (components *IstioComponents_STATUS) AssignProperties_From_IstioComponents_STATUS(source *storage.IstioComponents_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// EgressGateways
	if len(source.EgressGateways) > 0 {
		propertyBag.Add("EgressGateways", source.EgressGateways)
	} else {
		propertyBag.Remove("EgressGateways")
	}

	// IngressGateways
	if source.IngressGateways != nil {
		ingressGatewayList := make([]IstioIngressGateway_STATUS, len(source.IngressGateways))
		for ingressGatewayIndex, ingressGatewayItem := range source.IngressGateways {
			// Shadow the loop variable to avoid aliasing
			ingressGatewayItem := ingressGatewayItem
			var ingressGateway IstioIngressGateway_STATUS
			err := ingressGateway.AssignProperties_From_IstioIngressGateway_STATUS(&ingressGatewayItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_From_IstioIngressGateway_STATUS() to populate field IngressGateways")
			}
			ingressGatewayList[ingressGatewayIndex] = ingressGateway
		}
		components.IngressGateways = ingressGatewayList
	} else {
		components.IngressGateways = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		components.PropertyBag = propertyBag
	} else {
		components.PropertyBag = nil
	}

	// Invoke the augmentConversionForIstioComponents_STATUS interface (if implemented) to customize the conversion
	var componentsAsAny any = components
	if augmentedComponents, ok := componentsAsAny.(augmentConversionForIstioComponents_STATUS); ok {
		err := augmentedComponents.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_IstioComponents_STATUS populates the provided destination IstioComponents_STATUS from our IstioComponents_STATUS
func (components *IstioComponents_STATUS) AssignProperties_To_IstioComponents_STATUS(destination *storage.IstioComponents_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(components.PropertyBag)

	// EgressGateways
	if propertyBag.Contains("EgressGateways") {
		var egressGateway []storage.IstioEgressGateway_STATUS
		err := propertyBag.Pull("EgressGateways", &egressGateway)
		if err != nil {
			return errors.Wrap(err, "pulling 'EgressGateways' from propertyBag")
		}

		destination.EgressGateways = egressGateway
	} else {
		destination.EgressGateways = nil
	}

	// IngressGateways
	if components.IngressGateways != nil {
		ingressGatewayList := make([]storage.IstioIngressGateway_STATUS, len(components.IngressGateways))
		for ingressGatewayIndex, ingressGatewayItem := range components.IngressGateways {
			// Shadow the loop variable to avoid aliasing
			ingressGatewayItem := ingressGatewayItem
			var ingressGateway storage.IstioIngressGateway_STATUS
			err := ingressGatewayItem.AssignProperties_To_IstioIngressGateway_STATUS(&ingressGateway)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_To_IstioIngressGateway_STATUS() to populate field IngressGateways")
			}
			ingressGatewayList[ingressGatewayIndex] = ingressGateway
		}
		destination.IngressGateways = ingressGatewayList
	} else {
		destination.IngressGateways = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForIstioComponents_STATUS interface (if implemented) to customize the conversion
	var componentsAsAny any = components
	if augmentedComponents, ok := componentsAsAny.(augmentConversionForIstioComponents_STATUS); ok {
		err := augmentedComponents.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForIstioComponents_STATUS interface {
	AssignPropertiesFrom(src *storage.IstioComponents_STATUS) error
	AssignPropertiesTo(dst *storage.IstioComponents_STATUS) error
}

// Storage version of v1api20230202preview.IstioIngressGateway_STATUS
// Istio ingress gateway configuration. For now, we support up to one external ingress gateway named
// `aks-istio-ingressgateway-external` and one internal ingress gateway named `aks-istio-ingressgateway-internal`.
type IstioIngressGateway_STATUS struct {
	Enabled     *bool                  `json:"enabled,omitempty"`
	Mode        *string                `json:"mode,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_IstioIngressGateway_STATUS populates our IstioIngressGateway_STATUS from the provided source IstioIngressGateway_STATUS
func (gateway *IstioIngressGateway_STATUS) AssignProperties_From_IstioIngressGateway_STATUS(source *storage.IstioIngressGateway_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Enabled
	if source.Enabled != nil {
		enabled := *source.Enabled
		gateway.Enabled = &enabled
	} else {
		gateway.Enabled = nil
	}

	// Mode
	gateway.Mode = genruntime.ClonePointerToString(source.Mode)

	// Update the property bag
	if len(propertyBag) > 0 {
		gateway.PropertyBag = propertyBag
	} else {
		gateway.PropertyBag = nil
	}

	// Invoke the augmentConversionForIstioIngressGateway_STATUS interface (if implemented) to customize the conversion
	var gatewayAsAny any = gateway
	if augmentedGateway, ok := gatewayAsAny.(augmentConversionForIstioIngressGateway_STATUS); ok {
		err := augmentedGateway.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_IstioIngressGateway_STATUS populates the provided destination IstioIngressGateway_STATUS from our IstioIngressGateway_STATUS
func (gateway *IstioIngressGateway_STATUS) AssignProperties_To_IstioIngressGateway_STATUS(destination *storage.IstioIngressGateway_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(gateway.PropertyBag)

	// Enabled
	if gateway.Enabled != nil {
		enabled := *gateway.Enabled
		destination.Enabled = &enabled
	} else {
		destination.Enabled = nil
	}

	// Mode
	destination.Mode = genruntime.ClonePointerToString(gateway.Mode)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForIstioIngressGateway_STATUS interface (if implemented) to customize the conversion
	var gatewayAsAny any = gateway
	if augmentedGateway, ok := gatewayAsAny.(augmentConversionForIstioIngressGateway_STATUS); ok {
		err := augmentedGateway.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForIstioIngressGateway_STATUS interface {
	AssignPropertiesFrom(src *storage.IstioIngressGateway_STATUS) error
	AssignPropertiesTo(dst *storage.IstioIngressGateway_STATUS) error
}
