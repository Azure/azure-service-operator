# Virtual Machine Extension Operator

This operator deploys an Azure Virtual Machine Extension into a specified resource group at the specified location. Users can specify the publisher, extension type name, settings, etc. for the VM extension.

Learn more about Azure Virtual Machine Extension [here](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachineextensions).

Here is a [sample YAML](/config/samples/azure_v1alpha1_azurevirtualmachineextension.yaml) to provision a Virtual Machine Extension.

The spec is comprised of the following fields:

* Location
* ResourceGroup
* VMName
* AutoUpgradeMinorVersion
* ForceUpdateTag
* Publisher
* TypeName
* TypeHandlerVersion
* Settings
* ProtectedSettings

### Required Fields

A Virtual Machine Extension needs the following fields to deploy, along with a location and resource group.

* `VMName`
* `AutoUpgradeMinorVersion`
* `ForceUpdateTag`
* `Publisher`
* `TypeName`
* `TypeHandlerVersion`
* `Settings`

### Optional Fields

* `ProtectedSettings`

## Deploy, view and delete resources

You can follow the steps [here](/docs/v1/howto/resourceprovision.md) to deploy, view and delete resources.
