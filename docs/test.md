# Testing

Testing the full project can be done in two ways:

- `make api-test` - Runs the Kubernetes API tests against the CRDs
- `make test` - Test against the Kubernetes integration testing framework.
- `make test-existing-managers` - Test the resource managers against an existing cluster. This is currently the easiest way to run tests against a kind cluster setup.
- `make test-existing-controllers` - Test the controllers against an existing cluster. By default this will run every controller test, but you can edit the `tags` parameter in the makefile to specify individual test suites. Each controller test file declares its tags in the `// +build` comment at the top of the file.


