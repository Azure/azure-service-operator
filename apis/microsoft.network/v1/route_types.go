/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type (
	// RouteSpecProperties are the resource specific properties
	RouteSpecProperties struct {
		AddressPrefix    string `json:"addressPrefix,omitempty"`
		NextHopIPAddress string `json:"nextHopIpAddress,omitempty"`
		// +kubebuilder:validation:Enum=Internet;None;VirtualAppliance;VirtualNetworkGateway;VnetLocal
		NextHopType string `json:"nextHopType,omitempty"`
	}

	// RouteSpec defines the desired state of Route
	RouteSpec struct {
		// +k8s:conversion-gen=false
		APIVersion string `json:"apiVersion,omitempty"`

		// Properties of the subnet
		Properties *RouteSpecProperties `json:"properties,omitempty"`
	}

	// RouteStatus defines the observed state of Route
	RouteStatus struct {
		ID string `json:"id,omitempty"`
		// +k8s:conversion-gen=false
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:storageversion

	// Route is the Schema for the routes API
	Route struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   RouteSpec   `json:"spec,omitempty"`
		Status RouteStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// RouteList contains a list of Route
	RouteList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []Route `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
