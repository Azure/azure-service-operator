# Network Interface Operator

This operator deploys an Azure Network Interface (NIC) into a specified resource group at the specified location. Users can specify underlying public IP address and virtual network configurations in their NIC setup.

Learn more about Azure Network Interface [here](https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-network-interface).

Here is a [sample YAML](/config/samples/azure_v1alpha1_azurenetworkinterface.yaml) to provision a Network Interface.

The spec is comprised of the following fields:

* Location
* ResourceGroup
* VNetName
* SubnetName
* PublicIPAddressName

### Required Fields

A Network Interface needs the following fields to deploy, along with a location and resource group.

* `VNetName` specify the name for the virtual network that the network interface belongs to
* `SubnetName` specify the name for the subnet that the network interface belongs to
* `PublicIPAddressName` specify the name for the public IP address that the network interface uses

### Optional Fields

Not available.

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
