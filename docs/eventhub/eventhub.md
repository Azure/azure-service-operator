# Eventhub Operator

## Resources supported

The Eventhub operator can be used to provision the following resources.

1. Eventhub - Deploys an EventHub instance given the Eventhub namespace, Resource Group and Location.

2. Eventhub namespace - Deploys an EventHub namespace given the Resource Group and Location. Also has the ability to configure Sku, Properties and Network rule.

3. Consumer groups - Deploys a consumer group given the Eventhub, Eventhub namespace and Resource Group.

### Eventhub - Deployment output

The Eventhub operator deploys an Eventhub in the specified namespace according to the Spec.

As an output of deployment, the operator stores a JSON formatted secret with the following fields. For more details on where the secrets are stored, look [here](/docs/secrets.md)

- `primaryConnectionString`
- `secondaryConnectionString`
- `primaryKey`
- `secondaryKey`
- `sharedaccessKey`
- `eventhubNamespace`
- `eventhubName`

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.

## How would you use the Eventhub Operator from a real application

TODO: Demo app
