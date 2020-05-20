// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	"github.com/Azure/azure-service-operator/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *AzureSqlFirewallRule) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.AzureSqlFirewallRule)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Server = src.Spec.Server
	dst.Spec.StartIPAddress = src.Spec.StartIPAddress
	dst.Spec.EndIPAddress = src.Spec.EndIPAddress

	// Status
	dst.Status = v1beta1.ASOStatus(src.Status)

	return nil
}

func (dst *AzureSqlFirewallRule) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.AzureSqlFirewallRule)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Server = src.Spec.Server
	dst.Spec.StartIPAddress = src.Spec.StartIPAddress
	dst.Spec.EndIPAddress = src.Spec.EndIPAddress

	// Status
	dst.Status = ASOStatus(src.Status)

	return nil

}
