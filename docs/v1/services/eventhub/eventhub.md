# Event Hubs Operator

## Resources supported

The Event Hubs operator can be used to provision the following resources.

1. Event Hubs - Deploys an Event Hubs instance given the Event Hubs namespace, Resource Group and Location. 
   1. [Sample YAML file](config/samples/azure_v1alpha1_eventhub.yaml)

2. Event Hubs namespace - Deploys an Event Hubs namespace given the resource group and location. Also has the ability to configure SKU, properties, and network rules.
   1. [Sample YAML file](config/samples/azure_v1alpha1_eventhub_namespace.yaml)

3. Consumer groups - Deploys a consumer group given the event hub, Event Hubs namespace and resource group. 
   1. [Sample YAML file](config/samples/azure_v1alpha1_capture.yaml)

### Event Hubs - Deployment output

The Event Hubs operator deploys an event hub in the specified namespace according to the spec.

As an output of deployment, the operator stores a JSON formatted secret with the following fields. For more details on where the secrets are stored, look [here](/docs/v1/howto/secrets.md).

- `primaryConnectionString`
- `secondaryConnectionString`
- `primaryKey`
- `secondaryKey`
- `sharedaccessKey`
- `eventhubNamespace`
- `eventhubName`

## Deploy, view and delete resources

You can follow the steps [here](/docs/v1/howto/resourceprovision.md) to deploy, view and delete resources.

## Sample application

You can use this example to better understand how to use the Event Hubs Operator in a real application. We'll deploy a pair of producer and consumer apps to Kubernetes that will send and receive messages from Event Hubs respectively. Both the apps are written in [Go](https://golang.org/) using the [Sarama library for Kafka](https://github.com/Shopify/sarama/).
 
### Setup Azure Event Hubs

Create an Azure Resource Group

> Replace `location` if required

```shell
cat <<EOF | kubectl apply -f -
apiVersion: azure.microsoft.com/v1alpha1
kind: ResourceGroup
metadata:
  name: eh-aso-rg
spec:
  location: southeastasia
EOF
```

Use `kubectl get resourcegroups/eh-aso-rg` to confirm that the Resource Group was created (`PROVISIONED` status should show up as `true`).

Create Event Hubs namespace

> Replace `location` if required

```shell
cat <<EOF | kubectl apply -f -
apiVersion: azure.microsoft.com/v1alpha1
kind: EventhubNamespace
metadata:
  name: eh-aso-ns
spec:
  location: southeastasia
  resourceGroup: eh-aso-rg
  sku:
    name: Standard
    tier: Standard
    capacity: 1
  properties:
    kafkaEnabled: true
    isAutoInflateEnabled: false
    maximumThroughputUnits: 0
EOF
```

Wait for the namespace to be created and use `kubectl get eventhubnamespaces/eh-aso-ns` to confirm (`PROVISIONED` status should show up as `true`).

Create Event Hub

> Replace `location` if required

```shell
cat <<EOF | kubectl apply -f -
apiVersion: azure.microsoft.com/v1alpha1
kind: Eventhub
metadata:
  name: eh-aso-hub
spec:
  secretName: eh-secret
  location: southeastasia
  resourceGroup: eh-aso-rg
  namespace: eh-aso-ns
  properties:
    messageRetentionInDays: 7
    partitionCount: 3
  authorizationRule:
    name: RootManageSharedAccessKey
    rights:
      - Listen
      - Manage
      - Send
EOF
```

Wait For the Event Hub to be created and use `kubectl get eventhubs/eh-aso-hub` to confirm (`PROVISIONED` status should show up as `true`).

Create Consumer Group

```shell
cat <<EOF | kubectl apply -f -
apiVersion: azure.microsoft.com/v1alpha1
kind: ConsumerGroup
metadata:
  name: eh-aso-cg
spec:
  resourceGroup: eh-aso-rg
  namespace: eh-aso-ns
  eventHub: eh-aso-hub
EOF
```

Wait For the Consumer Group to be created and use `kubectl get consumergroups/eh-aso-cg` to confirm (`PROVISIONED` status should show up as `true`).

### Deploy applications

Deploy the Consumer application

```shell
cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: eh-consumer
spec:
  replicas: 1
  selector:
    matchLabels:
      app: eh-consumer
  template:
    metadata:
      labels:
        app: eh-consumer
    spec:
      nodeSelector:
        "beta.kubernetes.io/os": linux
      containers:
        - name: eh-consumer
          image: abhirockzz/eh-kafka-consumer
          env:
            - name: EVENTHUBS_CONNECTION_STRING
              valueFrom:
                secretKeyRef:
                  name: eh-secret
                  key: primaryConnectionString
            - name: EVENTHUBS_NAMESPACE
              valueFrom:
                secretKeyRef:
                  name: eh-secret
                  key: eventhubNamespace
            - name: EVENTHUBS_BROKER
              value: \$(EVENTHUBS_NAMESPACE).servicebus.windows.net:9093
            - name: EVENTHUBS_TOPIC
              valueFrom:
                secretKeyRef:
                  name: eh-secret
                  key: eventhubName
            - name: EVENTHUBS_CONSUMER_GROUPID
              value: eh-aso-cg
EOF
```

Wait for the consumer application to start

```shell
kubectl get pods -l=app=eh-consumer -w

//Keep a track of the logs for the consumer app
kubectl logs -f $(kubectl get pods -l=app=eh-consumer --output=jsonpath={.items..metadata.name})
```

You should see something similar to:

```bash
Event Hubs broker [eh-aso-ns.servicebus.windows.net:9093]
Sarama client consumer group ID eh-aso-cg
new consumer group created
Event Hubs topic eh-aso-hub
Waiting for program to exit
Partition allocation - map[eh-aso-hub:[0 1 2]]
```

Deploy the Producer application

```shell
cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: eh-producer
spec:
  replicas: 1
  selector:
    matchLabels:
      app: eh-producer
  template:
    metadata:
      labels:
        app: eh-producer
    spec:
      nodeSelector:
        "beta.kubernetes.io/os": linux
      containers:
        - name: eh-producer
          image: abhirockzz/eh-kafka-producer
          env:
            - name: EVENTHUBS_CONNECTION_STRING
              valueFrom:
                secretKeyRef:
                  name: eh-secret
                  key: primaryConnectionString
            - name: EVENTHUBS_NAMESPACE
              valueFrom:
                secretKeyRef:
                  name: eh-secret
                  key: eventhubNamespace
            - name: EVENTHUBS_BROKER
              value: \$(EVENTHUBS_NAMESPACE).servicebus.windows.net:9093
            - name: EVENTHUBS_TOPIC
              valueFrom:
                secretKeyRef:
                  name: eh-secret
                  key: eventhubName
EOF
```

Once the producer application is up and running, the consumer should kick in, start receiving the messages and print them to the console (check the consumer app terminal)

> In case you want to check producer logs: `kubectl logs -f $(kubectl get pods -l=app=eh-producer --output=jsonpath={.items..metadata.name})`

### Clean-up

Delete the applications:

```shell
kubectl delete deployment/eh-producer
kubectl delete deployment/eh-consumer
```

To delete Event Hubs namespace (and associated event hub and consumer group), you can simply delete the Azure Resource Group

```shell
kubectl delete resourcegroups/eh-aso-rg
```
