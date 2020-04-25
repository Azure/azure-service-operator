# Virtual Machine Operator

This operator deploys an Azure Virtual Machine (VM) into a specified resource group at the specified location. Users can specify platform image, size, user name and public SSH key, etc. for the VM.

Learn more about Azure Virtual Machine [here](https://docs.microsoft.com/en-us/rest/api/compute/virtualmachines).

Here is a [sample YAML](/config/samples/azure_v1alpha1_azurevirtualmachine.yaml) to provision a Virtual Machine.

The spec is comprised of the following fields:

* Location
* ResourceGroup
* VMSize
* AdminUserName
* SshPublicKeyData
* NetworkInterfaceName
* PlatformImageURN

### Required Fields

A Virtual Machine needs the following fields to deploy, along with a location and resource group.

* `VMSize` specify the VM size for the virtual machine
* `AdminUserName` specify the user name for the virtual machine
* `SshPublicKeyData` specify the SSH public key data for loging into the virtual machine
* `NetworkInterfaceName` specify the network interface that the VM will use
* `PlatformImageURN` specify the platform image's uniform resource name (URN) in the 'publisher:offer:sku:version' format.

### Optional Fields

Not available.

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
