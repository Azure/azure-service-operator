# Testing

Testing the full project can be done in two ways:

- `make api-test` - Runs the Kubernetes API tests against the CRDs
- `make test` - Test against the Kubernetes integration testing framework.
- `make test-existing-managers` - Test the resource managers against an existing cluster. This is currently the easiest way to run tests against a kind cluster setup.
- `make test-existing-controllers` - Test the controllers against an existing cluster. 

Some test suites (primarily the controller suite) have included Go Build Tags to allow selectively running tests. Users can specify individual test suites by setting the `BUILD_TAGS` parameter as an environment variable or before the make target: `make BUILD_TAGS=azuresqlservercombined test-existing-controllers`. Test files declare their tags in the `// +build` comment at the top of the file.
