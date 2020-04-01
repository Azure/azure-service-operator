# Eventhub Operator

## Resources supported

The Eventhub operator can be used to provision the following resources.

1. Eventhub - Deploys an EventHub instance given the Eventhub namespace, Resource Group and Location.

2. Eventhub namespace - Deploys an EventHub namespace given the Resource Group and Location. Also has the ability to configure Sku, Properties and Network rule.

3. Consumer groups - Deploys a consumer group given the Eventhub, Eventhub namespace and Resource Group.

## Deploying resources

You can follow the steps [here](/docs/development.md) to either run the operator locally or in a real Kubernetes cluster.

You can use the YAML files in the `config/samples` folder to create the resources.

## View and Troubleshoot resource provisioning

To view your created Eventhub resources, refer to the steps [here](viewresources.md)


## How would you use the Eventhub Operator from a real application

TODO: Demo app
