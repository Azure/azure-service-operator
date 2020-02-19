# Eventhub Operator

## Resources supported

The Eventhub operator can be used to provision the following resources.

1. Eventhub - Deploys an EventHub instance given the Eventhub namespace, Resource Group and Location.

2. Eventhub namespace - Deploys an EventHub namespace given the Resource Group and Location.

3. Consumer groups - Deploys a consumer group given the Eventhub, Eventhub namespace and Resource Group.

## Deploying resources

You can follow the steps [here](/docs/development.md) to either run the operator locally or in a real Kubernetes cluster.

You can use the YAML files in the `config/samples` folder to create the resources.

**Note**  Don't forget to set the Service Principal ID, Service Principal secret, Tenant ID and Subscription ID as environment variables

## View and Troubleshoot resource provisioning

To view your created Eventhub resources, refer to the steps [here](viewresources.md)

## Help

1. If the secret for the Eventhub in k8s gets deleted accidentally, the reconcile for the parent eventhub is triggered and secret gets created again.
2. If EventhubNamespace and Eventhub are deleted in Azure, then we need to delete the objects in k8s for the resources to be recreated. Reason being, if we apply the same manifest k8s does it recognise it as a change and the reconcile is not triggered.

## How would you use the Eventhub Operator from a real application

TODO: Demo app
