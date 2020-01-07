# Resource Group Operator

The Resource group operator can be used to provision a Resource group given the location and a name

## Deploying resources

You can follow the steps [here](/docs/development.md) to either run the operator locally or in a real Kubernetes cluster.

You can use the YAML files in the `config/samples` folder to create the resources.

Here is the sample YAML file for creating a  Resource group.

```yaml
apiVersion: azure.microsoft.com/v1alpha1
kind: ResourceGroup
metadata:
  name: resourcegroup-sample-1907
spec:
  location: "westus"
```

This would create a Resource group by the name `resourcegroup-sample-1907` in the `westus` location.

## View and Troubleshoot resource provisioning

To view your created resource group, refer to the steps [here](viewresources.md)

## Delete a Resource Group

You can delete the resource group using the below command.

```shell
kubectl delete ResourceGroup <resourcegroupname>
```
