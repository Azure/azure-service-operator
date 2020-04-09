# Blob Container Operator

This operator deploys a Blob Container into a specified resource group at the specified location. A Storage Account must first be created prior to creating a Blob Container.

Learn more about Azure Blob Storage [here](https://docs.microsoft.com/en-us/azure/storage/blobs/storage-blobs-introduction).

Here is a [sample YAML](/config/samples/azure_v1alpha1_blobcontainer.yaml) to provision a Blob Container.

### Required Fields 

A Blob Container needs the following fields to deploy, along with a location and resource group.

* `AccountName` specify the account name of the Storage Account to add the Blob Container to
* `AccessLevel` specify the access level of the data in the container. Possible values include: 'Container', 'Blob', 'None'

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.