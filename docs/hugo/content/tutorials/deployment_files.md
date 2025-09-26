---
title: Creating deployment manifests
weight: 10
type: docs
---

## Creating your own deployment files 

Typially, a deployment of a specific type of Azure resource is its own CRD; you can tell what type of resource is deployed by looking at its `kind:`.

The `spec:` field of the deployment file will typically be associated with the parameters you expect when creating the resource in Azure.

### Example AKS Cluster

Lets look at an AKS cluster as an example. The [latest sample](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/containerservice/v1api20240901/v1api20240901_managedcluster.yaml) looks like so:

```yaml
apiVersion: containerservice.azure.com/v1api20240901
kind: ManagedCluster
metadata:
  name: sample-managedcluster-20240901
  namespace: default
spec:
  location: westus3
  owner:
    name: aso-sample-rg
  dnsPrefix: aso
  agentPoolProfiles:
    - name: pool1
      count: 1
      vmSize: Standard_DS2_v2
      osType: Linux
      mode: System
  identity:
    type: SystemAssigned
```

If we wanted to specify our own deployment, we'd replace parameters as we typically fill in for creation in the CLI. Let's say I wanted to create `myAKSCluster` in `eastus`, for instance. It'd look like:

```yaml
apiVersion: containerservice.azure.com/v1api20240901
kind: ManagedCluster
metadata:
  name: myAKSCluster
  namespace: default
spec:
  location: eastus
  owner:
    name: <myResourceGroup>
  dnsPrefix: aso
  agentPoolProfiles:
    - name: agentpool1
      count: 1
      vmSize: <my_vm_size>
      osType: Linux
      mode: System
  identity:
    type: SystemAssigned
```

## More Examples

If you want to create your own deployment files, you can visit our [samples directory](https://github.com/Azure/azure-service-operator/tree/main/v2/samples) to see how different deployments are structured and alter the manifest according to your own specifications.
