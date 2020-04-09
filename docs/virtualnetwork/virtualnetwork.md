# Virtual Network Operator

This operator deploys an Azure Virtual Network into a specified resource group at the specified location, in the address space provided. Users are also able to add subnets to their virtual network through the operator.

Learn more about Azure Virtual Networks [here](https://docs.microsoft.com/en-us/azure/virtual-network/virtual-networks-overview).

Here is a [sample YAML](/config/samples/azure_v1alpha1_virtualnetwork.yaml) to provision a Virtual Network.

The spec is comprised of the following fields:

* Location
* ResourceGroup
* AddressSpace
* Subnets
  * SubnetName
  * SubnetAddressPrefix

### Required Fields

A Virtual Network needs the following fields to deploy, along with a location and resource group.

* `AddressSpace` specify an address space for your virtual network

### Optional Fields

You are able to specify a single or multiple subnets for your virtual network.

* `Subnets.SubnetName` specify a name for your subnet
* `Subnets.SubnetAddressPrefix` specify an address space for your subnet

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
