# Virtual Network Operator

This operator will deploy an Azure Virtual Network to a resource group and location, in the address space provided. Users are able to add subnets to their virtual network.

Learn more about Azure Virtual Networks [here](https://docs.microsoft.com/en-us/azure/virtual-network/virtual-networks-overview).

A virtual network is comprised of the following fields:

* Location
* ResourceGroup
* Address Space
* Subnets
  * SubnetName
  * SubnetAddressPrefix

## Deploy

Follow the steps [here](/docs/development.md) or [here](/docs/deploy.md) to either run the operator locally or in a real Kubernetes cluster.

You can find a sample Virtual Network YAML [here](/config/samples/azure_v1alpha1_virtualnetwork.yaml).