# Virtual Machine Operator

This operator deploys an Azure Disk into a specified resource group at the specified location. Users can specify the disk CreateOption and the disk size. Currently the only CreateOption that we support is 'Empty'.

Learn more about Azure Virtual Machine [here](https://docs.microsoft.com/en-us/rest/api/compute/disks).

Here is a [sample YAML](/config/samples/azure_v1alpha1_azuredisk.yaml) to provision a Virtual Machine.

The spec is comprised of the following fields:

* Location
* ResourceGroup
* CreateOption
* DiskSizeGB

### Required Fields

A Virtual Machine needs the following fields to deploy, along with a location and resource group.

* `CreateOption` specify the CreateOption. Currently the operator only supports 'Empty'.
* `DiskSizeGB` specify the Disk Size in GB.

### Optional Fields

Not available.

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
