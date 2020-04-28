package v1beta1

import (
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *AzureSqlServer) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha1.AzureSqlServer)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup

	return nil
}

func (dst *AzureSqlServer) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha1.AzureSqlServer)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup

	return nil

}
