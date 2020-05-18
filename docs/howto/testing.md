# Testing

Testing can be done in the following ways:

- `make test-integration-controllers` - Test the controllers end-to-end against Azure. You will need to have an active Kubernetes Cluster and run `make install` before running the tests.

This is the most exhaustive and recommended way to test the controllers.

The controller test suite has included Go Build Tags to allow selectively running tests. Users can specify individual test suites by setting the `BUILD_TAGS` parameter as an environment variable or before the make target: `make BUILD_TAGS=azuresqlservercombined test-integration-controllers`. Test files declare their tags in the `// +build` comment at the top of the file.

- `make test-integration-managers` - Test the Azure SDK resource managers against an existing cluster. This tests the Azure APIs but does not test the Kubernetes controllers.

- `make test-unit` - Runs the API and resource manager unit tests.

## Pipeline test runs

The controller test suite is run in against all PRs as part of the continuous integration pipeline.