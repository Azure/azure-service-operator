# Public IP Address Operator

This operator deploys an Azure Public IP Address (PIP) into a specified resource group at the specified location. Users can specify IP allocation method, idle timeout, IP address version, and SKU.

Learn more about Azure Public IP Address [here](https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-ip-addresses-overview-arm).

Here is a [sample YAML](/config/samples/azure_v1alpha1_azurepublicipaddress.yaml) to provision a Public IP Address.

The spec is comprised of the following fields:

* Location
* ResourceGroup
* PublicIPAllocationMethod
* IdleTimeoutInMinutes
* PublicIPAddressVersion
* SkuName

### Required Fields

A Public IP Address needs the following fields to deploy, along with a location and resource group.

* `PublicIPAllocationMethod` specify the allocation method for the public IP address, either 'Static' or 'Dynamic'
* `IdleTimeoutInMinutes` specify the idle timeout value (in minutes) for the public IP address
* `PublicIPAddressVersion` specify the version for the public IP address, either 'IPv4' or 'IPv6'
* `SkuName` specify the SKU name for the public IP address, either 'Basic' or 'Standard'

### Optional Fields

Not available.

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
