# Resource Group Operator

The Resource group operator can be used to provision a Resource group given the location and a name. It is currently in `v1beta1`.

Learn more about Azure resource groups [here](https://docs.microsoft.com/azure/azure-resource-manager/management/manage-resource-groups-cli).

Here is a [sample YAML](/config/samples/azure_v1alpha1_resourcegroup.yaml) to provision a resource group.

## Required Fields

A resource group needs the following fields to deploy

* `name` specify the name of the resource group you'd like to create
* `location` specify which region the resource group should be created in

## Deploy, view and delete resources

You can follow the steps [here](/docs/v1/howto/resourceprovision.md) to deploy, view and delete resources.
