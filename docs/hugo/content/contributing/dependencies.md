---
title: Developer Dependencies
linktitle: Dependencies
---
Development of Azure Service Operator depends on a number of development tools and libraries that need to be installed. 

If you prefer to install those dependencies manually (instead of using the `.devcontainer/install-dependencies.sh` script), here is a list of what's required. 

| Dependency | Version | Reference |
|:---------- |:-------:|:--------- |
| AZ | latest | https://docs.microsoft.com/en-us/cli/azure/install-azure-cli |
| AZWI | v1.5.1 | https://github.com/Azure/azure-workload-identity |
| BuildX | v0.29.1 | https://github.com/docker/buildx |
| cmctl | latest | https://cert-manager.io/docs/reference/cmctl |
| controller-gen | v0.19.0 | https://book.kubebuilder.io/reference/controller-gen |
| conversion-gen | v0.34.1 | https://pkg.go.dev/k8s.io/code-generator/cmd/conversion-gen |
| crddoc | latest | https://github.com/theunrepentantgeek/crddoc |
| Go | 1.23 | https://golang.org/doc/install #
| gofumpt | latest | https://pkg.go.dev/mvdan.cc/gofumpt |
| golangci-lint | 2.6.1 | https://github.com/golangci/golangci-lint |
| Helm | v3.19.0 | https://helm.sh/ |
| htmltest | latest | https://github.com/wjdp/htmltest (but see https://github.com/theunrepentantgeek/htmltest for our custom build )
| hugo | v0.152.2 | https://gohugo.io/ |
| kind | v0.30.0 | https://kind.sigs.k8s.io/ |
| kustomize | v4.5.7 | https://kustomize.io/ |
| Pip3 | latest | https://pip.pypa.io/en/stable/installation/ |
| PostCSS | latest | https://postcss.org/ |
| setup-envtest | v0.22.4 | https://book.kubebuilder.io/reference/envtest.html |
| Task | v3.44.0 | https://taskfile.dev/ |
| Trivy | v0.67.2 | https://trivy.dev/ |
| YQ | v4.48.1 | https://github.com/mikefarah/yq/ |

Dependencies are listed alphabetically. Refer to `install-dependencies.sh` for a recommended order of installation.

To update this file:

* Modify the file header content in `docs/v2/dependencies-header.md`;
* Modify the file footer in `docs/v2/dependencies-footer.md`; or
* Modify the dependencies installed by `.devcontainer/install-dependencies.sh`.

Regenerate the file using task:

``` bash
$ task doc:dependencies
```

Finally, submit the update as a PR.
