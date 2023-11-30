package labels

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

const OwnerNameLabel = "serviceoperator.azure.com/owner-name"

func SetOwnerNameLabel(obj genruntime.ARMMetaObject) {
	if obj.Owner() != nil && obj.Owner().Name != "" {
		genruntime.AddLabel(obj, OwnerNameLabel, obj.Owner().Name)
	}
}
