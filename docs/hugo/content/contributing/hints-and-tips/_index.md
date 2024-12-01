---
title: "Hints and Tips"
type: docs
description: "Extra information for making a successful contribution."
---

## Submodules

This project includes a submodule ([`azure-rest-api-specs`](https://github.com/Azure/azure-rest-api-specs)) referencing the Azure Resource Manager API specifications. This submodule is used to generate the code for the resources in the `v2/api` package.

### Changes to azure-rest-api-specs

Generally speaking, don't make changes to the `azure-rest-api-specs` submodule unless necessary for importing your new resource. We (the project team) regularly keep it up to date so it's unlikely you'll need to make changes to it yourself.

**Motivation**: changes to this submodule may have wide ranging effect - while Azure Resource Manager policies on backward compatible APIs mean that the _shape_ of resources is unlikely to change, documentation on resources and properties often does change. Including those changes in a PR along with other changes might make it unwieldy to review.

If an update is necessary, create a separate PR that just updates the submodule, keeping those changes isolated. Once that's merged, you can create your new resource without any additional debris.

### Error: loading schema from root

If the submodule is missing, you'll see this error:

``` text
error loading schema from root ... 
open /azure-service-operator/v2/specs/azure-resource-manager-schemas/schemas/2019-04-01/deploymentTemplate.json 
no such file or directory
```

This can happen if the repo was not cloned with `--recurse-submodules`.

To resolve this, run `git submodule init` (to set up submodules) and `git submodule update` (to check out the code) and then try building again.

## Testing

### Combine as many resources as possible into a single call to CreateResources()

One of the key features of ASO is that it takes care of sequencing - it works out the correct order of resource creation itself. This is essential, as when someone uses `kubectl` to apply a YAML file containing all the resources for their application, all the resurces get created in the cluster at the same times and ASO has to be able to do the right thing.

It's important that we exercise this in our tests. We've found in the past some resources where additional work was required to make this run smoothly - this is why we have extension points defined in the [`genruntime/extensions package`](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2@v2.10.0/pkg/genruntime/extensions).

Instead of calling `tc.CreateResourceAndWait()` for each resource in turn, declare all the resources required for the test and then make a single call to `tc.CreateResourcesAndWait()` (note the plural in the name) to create them all at once.

## Code formatting and linting

We strongly believe that linting (aka static code analysis) can be very useful for identifying potential issues in the code. We use [golangci-lint](https://golangci-lint.run/) to run a suite of linters on the codebase as a part of our CI pipeline for every PR. We also use [gofumpt](https://github.com/mvdan/gofumpt), a stricter version of `gofmt`, to format the code.

To ensure your code passes CI, we suggest running the following two commands before submitting your pull request. 

First, ensure your code is formatted according to our standards:

```bash
task format-code
```

If any files are modified, commit them.


Then, run the linters:

```bash
task controller:lint
```

The above target is appropriate if you're contributing a new resource or controller feature. If you're contributing to our commandline tool, use `task asoctl:lint`; if you're working on our code generator, use `task generator:lint`.

If any issues are found, fix them before submitting your PR. 
