# Load Balancer Operator

This operator deploys an Azure Load Balancer (LB) into a specified resource group at the specified location.

Learn more about Azure Network Interface [here](https://docs.microsoft.com/en-us/rest/api/load-balancer/loadbalancers/createorupdate).

Here is a [sample YAML](/config/samples/azure_v1alpha1_azureloadbalancer.yaml) to provision a Load Balancer.

The spec is comprised of the following fields:

* Location
* ResourceGroup
* PublicIPAddressName
* BackendAddressPoolName
* InboundNatPoolName
* FrontendPortRangeStart
* FrontendPortRangeEnd
* BackendPort

### Required Fields

A Network Interface needs the following fields to deploy, along with a location and resource group.

* `PublicIPAddressName` specify the name for the public IP address that the network interface uses
* `BackendAddressPoolName` specify the backend address pool name
* `InboundNatPoolName` specify the inbound nat pool name
* `FrontendPortRangeStart` specify the start of the front end port range
* `FrontendPortRangeEnd` specify the end of the front end port range
* `BackendPort` specify the backend port

### Optional Fields

Not available.

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
