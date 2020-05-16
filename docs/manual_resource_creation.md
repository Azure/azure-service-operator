# Manual Resource Creation
First step in building a code generator is creating handwritten code that represents what
you'd like your code generator to produce. This will walk you through how to create a new
resource the hard way. At some point in the future, the code generator will do this for us.

## [Kubebuilder](https://github.com/kubernetes-sigs/kubebuilder) enters the chat
In this example, we'll be adding the Azure Subnet subresource. A subresource in Azure is one that is
has an ID, Name, Properties, and is owned by a top-level Azure resource. A top-level Azure
resource has all the fields of a subresource, but also Location, Tags, etc...

To create this subresouce, we'll use kubebuilder to generate the Azure API version as well as the
storage version of the resource. For this subresource we are going to use the [2019-11-01 API version](https://github.com/Azure/azure-rest-api-specs/blob/872f6e6a6ae0e21045ae85c7448ae5ffe614f685/specification/network/resource-manager/Microsoft.Network/stable/2019-11-01/virtualNetwork.json#L1569-L1592).

### Create the Azure API version
```bash
kubebuilder create api --controller=false --make=false --resource=true --kind Subnet --group microsoft.network --version v20191101
```

### Create the Storage schema
```bash
kubebuilder create api --controller=false --make=false --resource=true --kind Subnet --group microsoft.network --version v1
```

## Fill in the Structure

### The Azure API version

We'll add the details to match the Azure json schema described in OpenAPI.
```go
package v20191101

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type (
	// SubnetProperties are the properties of the subnet
	SubnetProperties struct {
		// AddressPrefix for the subnet, eg. 10.0.0.0/24
		AddressPrefix string `json:"addressPrefix,omitempty"`
		// AddressPrefixes are a list of address prefixes for a subnet
		AddressPrefixes []string `json:"addressPrefixes,omitempty"`
	}

	// SubnetSpec is a subnet in a Virtual Network
	SubnetSpec struct {
		// ID of the subnet resource
		ID string `json:"id,omitempty"`
		// Properties of the subnet
		Properties SubnetProperties `json:"properties,omitempty"`
	}

	// SubnetStatus defines the observed state of Subnet
	SubnetStatus struct {
		ID                string `json:"id,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true

	// Subnet is the Schema for the subnets API
	Subnet struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`
		Spec   SubnetSpec   `json:"spec,omitempty"`
		Status SubnetStatus `json:"status,omitempty"`
	}
)
```

### The Storage version

Note the difference in the status for the storage schema. 
- The `DeploymentID` enables the controller to track the state of the resource deployment in Azure.
- The `APIVersion` informs the storage schema what Azure API version used to project this resource.
```go
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type (
	// SubnetProperties are the properties of the subnet
	SubnetProperties struct {
		// AddressPrefix for the subnet, eg. 10.0.0.0/24
		AddressPrefix string `json:"addressPrefix,omitempty"`
		// AddressPrefixes are a list of address prefixes for a subnet
		AddressPrefixes []string `json:"addressPrefixes,omitempty"`
	}

	// SubnetSpec is a subnet in a Virtual Network
	SubnetSpec struct {
		// +k8s:conversion-gen=false
		APIVersion string `json:"apiVersion"`
		// Properties of the subnet
		Properties SubnetProperties `json:"properties,omitempty"`
	}

	// SubnetStatus defines the observed state of Subnet
	SubnetStatus struct {
		ID string `json:"id,omitempty"`
		// +k8s:conversion-gen=false
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:storageversion

	// Subnet is the Schema for the subnets API
	Subnet struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`
		Spec   SubnetSpec   `json:"spec,omitempty"`
		Status SubnetStatus `json:"status,omitempty"`
	}
)
```

## Conversions

### Conversion Hub
In the `v1` package, add an entry to the `conversion.go` file to declare the hub interface.
```go
func (*Subnet) Hub(){}
```

### Conversion To / From Hub
In the `v20191101` package, add an entry to the `conversion.go` file to the conversion funcs. Note
the addition of the API version projected into the storage schema.
```go
func (src *Subnet) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.Subnet)

	if err := Convert_v20191101_Subnet_To_v1_Subnet(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *Subnet) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.Subnet)

	if err := Convert_v1_Subnet_To_v20191101_Subnet(src, dst, nil); err != nil {
		return err
	}

	return nil
}
```

## Implement corev1.MetaObject
To link the CRD to it's Azure resource type, you must implement `ResourceType() string`.
```go
func (*Subnet) ResourceType() string {
	return "Microsoft.Network/virtualNetworks/subnets"
}
```

## Webooks
Add / uncomment kustomize patches for webhooks
```yaml
patchesStrategicMerge:
- patches/webhook_in_subnets.yaml
- patches/cainjection_in_subnets.yaml
```

Add validation and defaulting webhook behaviors in `v1.webhook.go`
```go
// log is for logging in this package.
var subnetlog = logf.Log.WithName("subnet-resource")

// +kubebuilder:webhook:path=/mutate-microsoft-network-infra-azure-com-v1-subnet,mutating=true,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=subnets,verbs=create;update,versions=v1,name=default.subnet.infra.azure.com

var _ webhook.Defaulter = &Subnet{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Subnet) Default() {
	subnetlog.Info("default", "name", r.Name)
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-microsoft-network-infra-azure-com-v1-subnet,mutating=false,matchPolicy=Equivalent,failurePolicy=fail,groups=microsoft.network.infra.azure.com,resources=subnets,versions=v1,name=validation.subnet.infra.azure.com

var _ webhook.Validator = &Subnet{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Subnet) ValidateCreate() error {
	subnetlog.Info("validate create", "name", r.Name)
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Subnet) ValidateUpdate(old runtime.Object) error {
	subnetlog.Info("validate update", "name", r.Name)
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Subnet) ValidateDelete() error {
	subnetlog.Info("validate delete", "name", r.Name)
	return nil
}
```

## Controller RBAC
To enable the reconciling controller to be able to manipulate subnet resources, we must add these
kubebuilder markers. The markers are added to the `./controllers/generic_controller.go` file.
```go
// +kubebuilder:rbac:groups=microsoft.network.infra.azure.com,resources=subnets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.network.infra.azure.com,resources=subnets/status,verbs=get;update;patch
```

## Referencing the Subnet in the Virtual Network

The Subnet subresource is owned by the Virtual Network top-level resource. To reference the subnet
in the Virtual Network in the storage schema, the following must be added.

```go
SubnetRefs []azcorev1.KnownTypeReference `json:"subnetRefs,omitempty" group:"microsoft.network.infra.azure.com" kind:"Subnet" owned:"true"`
```

The field above provides metadata to the controller that it needs to reconcile the subnet as a
subresource of the Virtual Network and should be deleted along with the Virtual Network, through 
OwnerReferences.

## The End
As you can see it takes quite a bit of boilerplate to add resources. This is why it would be so
advantageous to provide a way to transform Azure resource schema into K8s CRDs. By adding a little
extra metadata to our CRDs, we can use a generic controller to reconcile with Azure Resource Manager.

There is much to be done to make this easier. This is the way.


