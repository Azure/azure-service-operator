# ASOv2 Metrics

We have a set of prometheus metrics exposed for ASOv2 which can be helpful in diagnosability and tracing. 
The metrics exposed, fall into two groups of `Azure based metrics` and `reconciler metrics(controller-runtime)` 
combined.

## Enabling the metrics

By default, metrics scraping is turned on for ASOv2, which can be turned off using helm chart or Deployment yaml.

### ASOv2 Helm Chart:

While installing the helm chart, we can turn the metrics ON and OFF and set the metrics expose address using the 
below settings. Also, we can change the settings inside `values.yaml` file for ASOv2 helm chart.

   ```
   --set metrics.enable=true/false (default: true)
   --set metrics.address=127.0.0.1:8080 (default)
   ```

### Deployment Yaml:

In the deployment yaml, we can turn OFF the metrics by omitting the `metrics-addr` flag below and also set the metrics 
expose address there.
    
   ```
    spec:
      containers:
      - args:
        - --metrics-addr=127.0.0.1:8080 (default)
   ```
## Understanding the ASOv2 Metrics

| Metric                                         | Description                                                                                               | Label 1      | Label 2    | Label 3      |
|------------------------------------------------|-----------------------------------------------------------------------------------------------------------|--------------|------------|--------------|
| `controller_runtime_reconcile_total`           | A prometheus counter metric with total number of reconcilations per controller.                           | Controller   | Result     |              |
| `controller_runtime_errors_total`              | A prometheus counter metric with total number of errors from reconciler                                   | Controller   |            |              |
| `controller_runtime_reconcile_time_seconds`    | A prometheus histogram metric which keeps track of the duration of reconcilations                         | Controller   |            |              |
| `controller_runtime_max_concurrent_reconciles` | A prometheus gauge metric with the number of concurrent reconciles per controller                         | Controller   |            |              |
| `controller_runtime_active_workers`            | A prometheus gauge metric with the number of active workers per controller                                | Controller   |            |              |
| `azure_requests_total`                         | A prometheus counter metric with the record of total number of requests to ARM by increasing the counter. | ResourceName | HTTPMethod | ResponseCode |
| `azure_failed_requests_total`                  | A prometheus counter metric with the record of the number of failed requests to ARM.                      | ResourceName | HTTPMethod |              |
| `azure_requests_time_seconds`                  | A prometheus histogram metric with the record the round-trip time taken by the request to ARM             | ResourceName | HTTPMethod |              |


