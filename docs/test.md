# Testing

Testing the full project can be done in two ways:

- `make test` - Test against the Kubernetes integration testing framework.
- `make test-existing` - Test against an existing cluster. This is currently easiest to do run against a kind cluster setup.

Both of these invoke the Ginkgo test runner, running tests in parallel across 4 nodes by default.

## Focus tests

The ginkgo runner makes it simple to test single test cases, or packages in isolation.

- Rename the spec from `Describe` to `FDescribe`, or the individual test case from `It` to `FIt` excludes all other tests at the same level.
- Adding an additional parameter to ginkgo runner `--focus=REGEXP`.

Refer to [this](https://onsi.github.io/ginkgo/#focused-specs) page for more details.