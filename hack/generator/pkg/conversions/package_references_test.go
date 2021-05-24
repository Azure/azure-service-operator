package conversions

import "github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"

func makeTestLocalPackageReference(group string, version string) astmodel.LocalPackageReference {
	return astmodel.MakeLocalPackageReference("github.com/Azure/azure-service-operator/hack/generated/_apis", group, version)
}
