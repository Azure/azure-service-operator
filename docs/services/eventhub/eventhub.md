# Event Hubs Operator

## Resources supported

The Event Hubs operator can be used to provision the following resources.

1. Event Hubs - Deploys an Event Hubs instance given the Event Hubs namespace, Resource Group and Location.

2. Event Hubs namespace - Deploys an Event Hubs namespace given the resource group and location. Also has the ability to configure SKU, properties, and network rules.

3. Consumer groups - Deploys a consumer group given the event hub, Event Hubs namespace and resource group.

### Event Hubs - Deployment output

The Event Hubs operator deploys an event hub in the specified namespace according to the spec.

As an output of deployment, the operator stores a JSON formatted secret with the following fields. For more details on where the secrets are stored, look [here](/docs/topics/secrets.md)

- `primaryConnectionString`
- `secondaryConnectionString`
- `primaryKey`
- `secondaryKey`
- `sharedaccessKey`
- `eventhubNamespace`
- `eventhubName`

## Deploy, view and delete resources

You can follow the steps [here](/docs/topics/resourceprovision.md) to deploy, view and delete resources.

<!-- ## How would you use the Event Hubs Operator to support a real application?

TODO: Demo app -->
