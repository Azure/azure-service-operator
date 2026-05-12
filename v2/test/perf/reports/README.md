## Perf Reports

This directory contains historical performance reports generated with the `v2/tools/collect-metrics` tool. They can be plotted 
with `task controller:plot-perf-metrics` and examined to look at historical performance behavior.

The "startup" data was gathered by running the following commands:
- Launch the controller in an AKS cluster 
- `go run ./tools/collect-metrics/ --output-folder ./test/perf/reports`
- `k delete pod -n azureserviceoperator-system -l app.kubernetes.io/name=azure-service-operator`
- Wait a bit and ctrl+c the metrics collection
