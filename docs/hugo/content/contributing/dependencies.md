---
title: Developer Dependencies
linktitle: Dependencies
---
Development of Azure Service Operator depends on a number of development tools and libraries that need to be installed. 

If you prefer to install those dependencies manually (instead of using the `.devcontainer/install-dependencies.sh` script), here is a list of what's required. 

| Dependency | Version | Reference |
|:---------- |:-------:|:--------- |
| AZ | latest | https://docs.microsoft.com/en-us/cli/azure/install-azure-cli |
| AZWI | v1.2.0 | https://github.com/Azure/azure-workload-identity |
| BuildX | v0.11.2 | https://github.com/docker/buildx |
| cmctl | latest | https://cert-manager.io/docs/reference/cmctl |
| controller-gen | v0.14.0 | https://book.kubebuilder.io/reference/controller-gen |
| conversion-gen | v0.28.8 | https://pkg.go.dev/k8s.io/code-generator/cmd/conversion-gen |
| gen-crd-api-reference-docs | 11fe95cb | https://github.com/ahmetb/gen-crd-api-reference-docs |
| Go | 1.22 | https://golang.org/doc/install #
| gofumpt | latest | https://pkg.go.dev/mvdan.cc/gofumpt |
| golangci-lint | 1.51.2 | https://github.com/golangci/golangci-lint |
| Helm | v3.8.0 | https://helm.sh/ |
| htmltest | latest | https://github.com/wjdp/htmltest (but see https://github.com/theunrepentantgeek/htmltest for our custom build )
| hugo | v0.88.1 | https://gohugo.io/ |
| kind | v0.20.0 | https://kind.sigs.k8s.io/ |
| kustomize | v4.5.7 | https://kustomize.io/ |
| PostCSS | latest | https://postcss.org/ |
| setup-envtest | latest | https://book.kubebuilder.io/reference/envtest.html |
| Task | v3.31 | https://taskfile.dev/ |
| Trivy | v0.37.3 | https://trivy.dev/ |
| YQ | v4.13.0 | https://github.com/mikefarah/yq/ |

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
