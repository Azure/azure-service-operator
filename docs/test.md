# Testing

Testing the full project can be done in two ways:

- `make api-test` - Runs the Kubernetes API tests against the CRDs
- `make test` - Test against the Kubernetes integration testing framework.
- `make test-existing` - Test against an existing cluster. This is currently easiest to run tests against a kind cluster setup.
