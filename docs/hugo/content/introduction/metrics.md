# ASOv2 Metrics

Prometheus metrics are exposed for ASOv2 which can be helpful in diagnosability and tracing.
The metrics exposed fall into two groups: Azure based metrics and reconciler metrics (from `controller-runtime`).

## Disabling the metrics

By default, metrics scraping is turned on for ASOv2. It can be turned off using the Helm chart 
`--set metrics.enable` or by omitting `--metrics-addr` flag from the Deployment yaml.

### ASOv2 Helm Chart

While installing the Helm chart, we can turn the metrics ON and OFF and set the metrics expose address using the 
below settings. Also, we can change the settings inside `values.yaml` file for ASOv2 Helm chart.

   ```
   --set metrics.enable=true/false (default: true)
   --set metrics.address=127.0.0.1:8080 (default)
   ```

### Deployment Yaml

In the deployment yaml, we can turn OFF the metrics by omitting the `metrics-addr` flag. We can also change to use 
a different metrics-addr by changing the default value of that same flag.
    
   ```
    spec:
      containers:
      - args:
        - --metrics-addr=127.0.0.1:8080 (default)
   ```
## Understanding the ASOv2 Metrics

| Metric                                         | Description                                                                                                  | Label 1      | Label 2     | Label 3      |
|------------------------------------------------|--------------------------------------------------------------------------------------------------------------|--------------|-------------|--------------|
| `controller_runtime_reconcile_total`           | A prometheus counter metric with total number of reconcilations per controller.                              | Controller   | Result      |              |
| `controller_runtime_errors_total`              | A prometheus counter metric with total number of errors from reconciler                                      | Controller   |             |              |
| `controller_runtime_reconcile_time_seconds`    | A prometheus histogram metric which keeps track of the duration of reconcilations                            | Controller   |             |              |
| `controller_runtime_max_concurrent_reconciles` | A prometheus gauge metric with number of concurrent reconciles per controller                                | Controller   |             |              |
| `controller_runtime_active_workers`            | A prometheus gauge metric with number of active workers per controller                                       | Controller   |             |              |
| `azure_requests_total`                         | A prometheus counter metric with total number of requests to Azure                                           | ResourceName | RequestType | ResponseCode |
| `azure_failed_requests_total`                  | A prometheus counter metric with total number of failed requests to Azure                                    | ResourceName | RequestType |              |
| `azure_requests_time_seconds`                  | A prometheus histogram metric which keeps track of the duration of round-trip time taken by request to Azure | ResourceName | RequestType |              |


