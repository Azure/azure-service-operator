---
title: Metrics
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

Prometheus metrics are exposed for ASOv2 which can be helpful in diagnosability and tracing.
The metrics exposed fall into two groups: Azure based metrics, and reconciler metrics (from `controller-runtime`).

## Toggling the metrics

By default, secure metrics for ASOv2 are turned on and can be toggled by the following options:

### ASOv2 Helm Chart

    While installing the Helm chart, we can turn the metrics _**on**_ and _**off**_ and set the metrics expose address using the 
    below settings. Also, we can change the settings inside `values.yaml` file for ASOv2 Helm chart.

    ```
    --set metrics.enable=true/false (default: true)
    --set metrics.secure=true/false (default: true)
    --set metrics.profiling=true/false (default: false)
    --set metrics.address=0.0.0.0:8443 (default)
    ```

### Deployment YAML
    
    In the deployment yaml, we can turn _**off**_ the metrics by omitting the `metrics-addr` flag. We can also change to use 
    a different metrics-addr by changing the default value of that same flag.

    ```
    spec:
      containers:
       - args:
         - --metrics-addr=0.0.0.0:8080 (default)    
         - --secure-metrics=true/false (default: true)
         - --profiling-metrics=true/false (default: false)
    ```

## Scraping Metrics Securely via HTTPs using RBAC

A ServiceAccount token is required to scrape metrics securely. The corresponding ServiceAccount needs permissions on the "/metrics" and "debug/pprof" paths. 
This can be achieved e.g. by following the [Kubernetes documentation](https://kubernetes.io/docs/concepts/cluster-administration/system-metrics/).

Follow the steps below to scrape metrics securely.

{{< tabpane text=true left=true >}}
{{% tab header="Helm Chart" %}}
``` Helm Chart
--set metrics.enable=true
--set metrics.secure=true
--set metrics.profiling=true
--set metrics.address=0.0.0.0:8443
```
{{% /tab %}}
{{% tab header="Deployment YAML" %}}
``` Deployment YAML
spec:
containers:
 - args:
   - --metrics-addr=0.0.0.0:8443  
   - --secure-metrics=true 
   - --profiling-metrics=true
```
{{% /tab %}}
{{< /tabpane >}}

Deploy the following RBAC configuration. This creates a role that can scrape metrics.
```
cat << EOT | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: default-metrics
rules:
- nonResourceURLs:
  - "/metrics"
  - "/debug/pprof/*"
  verbs:
  - get
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: default-metrics
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: default-metrics
subjects:
- kind: ServiceAccount
  name: default
  namespace: default
EOT
```
Test locally:
  - Open a port-forward

      ```
      kubectl port-forward deployments/azureserviceoperator-controller-manager -n azureserviceoperator-system 8443
      ```
  - Create a ServiceAccount token and scrape metrics
      ```
      TOKEN=$(kubectl create token default)
      curl https://localhost:8443/metrics --header "Authorization: Bearer $TOKEN" -k
      ```
  
## Understanding the ASOv2 Metrics

See [controller-runtime metrics](https://book-v1.book.kubebuilder.io/beyond_basics/controller_metrics) for more details
about specific controller-runtime metrics.

See [workqueue](https://github.com/kubernetes-sigs/controller-runtime/blob/main/pkg/metrics/workqueue.go) for more details
about the workqueue metrics from controller-runtime.

| Metric                                               | Description                                                                                                    | Metric Type | Label 1    | Label 2     | Label 3      |
|------------------------------------------------------|----------------------------------------------------------------------------------------------------------------|-------------|------------|-------------|--------------|
| `azure_successful_requests_total`                    | Total number of requests to Azure we received responses for. responseCode may be a failure such as 4xx or 5xx. | counter     | resource   | requestType | responseCode |
| `azure_failed_requests_total`                        | Total number of requests which we didn't receive a response from Azure for.                                    | counter     | resource   | requestType |              |
| `azure_requests_time_seconds`                        | Tracks the duration of round-trip time taken by request to Azure.                                              | histogram   | resource   | requestType |              |
| `controller_runtime_reconcile_total`                 | Total number of reconciliations per controller.                                                                | counter     | controller | result      |              |
| `controller_runtime_errors_total`                    | Total number of errors per controller.                                                                         | counter     | controller |             |              |
| `controller_runtime_reconcile_panics_total`          | Total number of panics per controller.                                                                         | counter     | controller |             |              |
| `controller_runtime_terminal_reconcile_errors_total` | Total number of terminal reconciliation errors per controller.                                                 | counter     | controller |             |              |
| `controller_runtime_reconcile_time_seconds`          | Tracks the duration of reconciliations.                                                                        | histogram   | controller |             |              |
| `controller_runtime_max_concurrent_reconciles`       | Number of concurrent reconciles per controller.                                                                | gauge       | controller |             |              |
| `controller_runtime_active_workers`                  | Number of active workers per controller.                                                                       | gauge       | controller |             |              |
| `controller_runtime_webhook_panics_total`            | Total number of webhook panics.                                                                                | counter     |            |             |              |
| `controller_runtime_webhook_requests_in_flight`      | Total number of webhook requests in flight.                                                                    | gauge       | webhook    |             |              |
| `controller_runtime_webhook_requests_total`          | Total number of webhook requests by HTTP status code.                                                          | gauge       | webhook    | code        |              |
| `workqueue_depth`                                    | Total depth of the work queue per controller.                                                                  | gauge       | controller | name        |              |
| `workqueue_adds_total`                               | Total adds handled by the work queue per controller.                                                           | counter     | controller | name        |              |
| `workqueue_queue_duration_seconds`                   | How long items stay in workqueue before being acted upon.                                                      | histogram   | controller | name        |              |
| `workqueue_work_duration_seconds`                    | How long processing an item from the workqueue takes.                                                          | histogram   | controller | name        |              |
| `workqueue_unfinished_work_seconds`                  | How many seconds of work have been done not observed by work_duration. Related to stuck threads.               | gauge       | controller | name        |              |
| `workqueue_longest_running_processor_seconds`        | How many seconds has the longest running processor been running.                                               | gauge       | controller | name        |              |
| `workqueue_retries_total`                            | Total number of retries handled by workqueue.                                                                  | counter     | controller | name        |              |
| `go_gc_duration_seconds`                             | Wall-time pause duration in garbage collection cycles                                                          | summary     | controller |             |              | 
| `go_gc_gogc_percent`                                 | Heap size target percentage configured by the user                                                             | gauge       | controller |             |              |
| `go_gc_gomemlimit_bytes`                             | Runtime memory limit configured by the user                                                                    | gauge       | controller |             |              |
| `go_goroutines`                                      | Number of goroutines that currently exist                                                                      | gauge       | controller |             |              |
| `go_info`                                            | Information about the Go environment                                                                           | gauge       | controller |             |              |
| `go_memstats_alloc_bytes`                            | Number of bytes allocated in heap and currently in use                                                         | gauge       | controller |             |              |
| `go_memstats_alloc_bytes_total`                      | Total number of bytes allocated in heap until now                                                              | counter     | controller |             |              |
| `go_memstats_buck_hash_sys_bytes`                    | Number of bytes used by the profiling bucket hash table                                                        | gauge       | controller |             |              |
| `go_memstats_frees_total`                            | Total number of heap objects frees                                                                             | counter     | controller |             |              |
| `go_memstats_gc_sys_bytes`                           | Number of bytes used for garbage collection system metadata                                                    | gauge       | controller |             |              |
| `go_memstats_heap_alloc_bytes`                       | Number of heap bytes allocated and currently in use                                                            | gauge       | controller |             |              |
| `go_memstats_heap_idle_bytes`                        | Number of heap bytes waiting to be used                                                                        | gauge       | controller |             |              |
| `go_memstats_heap_inuse_bytes`                       | Number of heap bytes that are in use                                                                           | gauge       | controller |             |              |
| `go_memstats_heap_objects`                           | Number of currently allocated objects                                                                          | gauge       | controller |             |              |
| `go_memstats_heap_released_bytes`                    | Number of heap bytes released to OS                                                                            | gauge       | controller |             |              |
| `go_memstats_heap_sys_bytes`                         | Number of heap bytes obtained from system                                                                      | gauge       | controller |             |              |
| `go_memstats_last_gc_time_seconds`                   | Number of seconds since 1970 of last garbage collection                                                        | gauge       | controller |             |              |
| `go_memstats_mallocs_total`                          | Total number of heap objects allocated                                                                         | counter     | controller |             |              |
| `go_memstats_mcache_inuse_bytes`                     | Number of bytes in use by mcache structures                                                                    | gauge       | controller |             |              |
| `go_memstats_mcache_sys_bytes`                       | Number of bytes used for mcache structures obtained from system                                                | gauge       | controller |             |              |
| `go_memstats_mspan_inuse_bytes`                      | Number of bytes in use by mspan structures                                                                     | gauge       | controller |             |              |
| `go_memstats_mspan_sys_bytes`                        | Number of bytes used for mspan structures obtained from system                                                 | gauge       | controller |             |              |
| `go_memstats_next_gc_bytes`                          | Number of heap bytes when next garbage collection will take place                                              | gauge       | controller |             |              |
| `go_memstats_other_sys_bytes`                        | Number of bytes used for other system allocations                                                              | gauge       | controller |             |              |
| `go_memstats_stack_inuse_bytes`                      | Number of bytes obtained from system for stack allocator in non-CGO environments                               | gauge       | controller |             |              |
| `go_memstats_stack_sys_bytes`                        | Number of bytes obtained from system for stack allocator                                                       | gauge       | controller |             |              |
| `go_memstats_sys_bytes`                              | Number of bytes obtained from system                                                                           | gauge       | controller |             |              |
| `go_sched_gomaxprocs_threads`                        | Current runtime.GOMAXPROCS setting                                                                             | gauge       | controller |             |              |

The above table is not comprehensive. For a full set of metrics reported from the pod, query the `/metrics` API via the
`curl` command documented above.

### Labels

Labels are used to differentiate the characteristics of the metric that is being measured. Each metric with distinct labels
is an independent metric. Below are the labels used in ASOv2 metrics:

- **controller**: Each resource being reconciled against Azure ARM has a separate dedicated controller.
- **result**: Reconcile result returned by controller ( error | requeue | requeue_after | success ).
- **resource**: Resource type for which the request is sent, such as `Microsoft.Resources/resourceGroups`.
- **requestType**: HTTP request method ( GET | PUT | DELETE ).
- **responseCode**: HTTP status code in response from Azure.

