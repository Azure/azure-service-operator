# Virtual Machine Scale Set Operator

This operator deploys an Azure Virtual Machine Scale Set (VMSS) into a specified resource group at the specified location. Users can specify platform image, size, user name and public SSH key, etc. for the VMSS.

Learn more about Azure Virtual Machine [here](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachinescalesets).

Here is a [sample YAML](/config/samples/azure_v1alpha1_azurevmscaleset.yaml) to provision a VMSS.

The spec is comprised of the following fields:

* Location
* ResourceGroup
* VMSize
* Capacity
* AdminUserName
* SshPublicKeyData
* PlatformImageURN
* VirtualNetworkName
* SubnetName
* LoadBalancerName
* BackendAddressPoolName
* InboundNatPoolName

### Required Fields

A Virtual Machine needs the following fields to deploy, along with a location and resource group.

* `VMSize` specify the VM size for the virtual machine scale set
* `Capacity` specify the number of instances in the VMSS
* `AdminUserName` specify the user name for the virtual machine scale set
* `SshPublicKeyData` specify the SSH public key data for loging into the virtual machine scale set
* `PlatformImageURN` specify the platform image's uniform resource name (URN) in the 'publisher:offer:sku:version' format.
* `VirtualNetworkName` specify the virtual network
* `SubnetName` specify the subnet
* `LoadBalancerName` specify the load balancer
* `BackendAddressPoolName` specify the backend address pool
* `InboundNatPoolName` specify the inbound nat pool

### Optional Fields

Not available.

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
