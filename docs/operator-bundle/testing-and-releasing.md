# Testing and releasing the operatorhub bundle for ASO

This is a guide to producing and testing the operator bundle after an Azure Service Operator (ASO) release.
It's primarily intended as a reference for the ASO development team.

## Test upgrading from the old bundle to the new one

Creating a bundle for a new version of the operator requires specifying a `replaces` field that indicates to [OLM](https://olm.operatorframework.io/) which version of the bundle it should be installed over.
When creating a PR to the community-operators repo, one of the items on the checklist asks whether upgrading the bundle from the previous version has been tested.
This section explains how to be able to check that box in good conscience.

### Before you begin

You will need:

* A previous operator hub bundle (currently this is `quay.io/operatorhubio/azure-service-operator:v0.37.0`).
This is the latest released operator bundle, which might not correspond to the last version of the operator image (if some versions have been skipped).
* A newly published operator image that needs packaging into an operator bundle (for example `mcr.microsoft.com/k8s/azureserviceoperator:1.0.23956`)

### Set up your cluster

1. Create an [AKS cluster with ACR integration](https://docs.microsoft.com/en-us/azure/aks/cluster-container-registry-integration#create-a-new-aks-cluster-with-acr-integration).
(In theory these steps could work with a `kind` cluster, but the registry management will need some changing.)
2. Install OLM:
    ```sh
    curl -sL https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.18.1/install.sh | bash -s v0.18.1
    ```
3. Install cert-manager:
    ```sh
    k apply -f https://github.com/jetstack/cert-manager/releases/download/v1.12.1/cert-manager.yaml
    ```
4. Install the `opm` tool to build the index image:
    * Follow [these instructions](https://docs.openshift.com/container-platform/4.5/operators/admin/olm-managing-custom-catalogs.html#olm-installing-opm_olm-managing-custom-catalogs) to get the `opm` binary.
    * Use docker rather than podman to log into the redhat registry - the path for auth credentials will be `~/.docker/config.json` if it's installed from the docker package archive (as opposed to from a snap).
    * The `oc image extract` command is convenient for getting the binary out of the right image layer. If you don't have/want oc installed you can find it manually using `docker image save` and `tar` to find and extract it.
5. Log in to your ACR:
    ```sh
    az login
    az acr login --name <registry-name>
    ```

### Make the bundle image and index

An operator bundle is a docker image containing the yaml of the CRDs and a [ClusterServiceVersion](https://docs.openshift.com/container-platform/4.7/rest_api/operatorhub_apis/clusterserviceversion-operators-coreos-com-v1alpha1.html) that defines the operator webhooks and deployment.
The bundle index is a docker image that contains a registry of bundles (potentially in different versions) - for our testing we'll be making a minimal one that just contains the current version and the new version of the ASO bundle.

0. Update the `PREVIOUS_BUNDLE_VERSION` value in the Makefile if it's not right (and commit the change).
1. Run `make generate-operator-bundle` - this will produce the source files for a new bundle with the same version as the new operator version.
2. Set these environment variables to enable copying commands from the steps below:
     ```sh
     export REGISTRY_FQDN=<your registry fqdn>
     export PREVIOUS_BUNDLE_VERSION=<previous bundle version>
     export NEW_BUNDLE_VERSION=<new bundle version>
     ```
3. Build the new bundle image:
    ```sh
    docker build -f bundle.Dockerfile --tag ${REGISTRY_FQDN}/operatorhubio/azure-service-operator:v${NEW_BUNDLE_VERSION} .
    ```
4. Push the new bundle to the registry:
    ```sh
    docker push ${REGISTRY_FQDN}/operatorhubio/azure-service-operator:v${NEW_BUNDLE_VERSION}
    ```
5. Pull the previous bundle version from the `quay.io` registry:
    ```sh
    docker pull quay.io/operatorhubio/azure-service-operator:v${PREVIOUS_BUNDLE_VERSION}
    ```
6. Tag the previous bundle for your test registry:
    ```sh
    docker tag quay.io/operatorhubio/azure-service-operator:v${PREVIOUS_BUNDLE_VERSION} ${REGISTRY_FQDN}/operatorhubio/azure-service-operator:v${PREVIOUS_BUNDLE_VERSION}
    ```
7. Push that bundle to your registry as well:
    ```sh
    docker push ${REGISTRY_FQDN}/operatorhubio/azure-service-operator:v${PREVIOUS_BUNDLE_VERSION}
    ```
8. Make the index image that points to the two bundles:
    ```sh
    opm -c docker index add \
        --bundles ${REGISTRY_FQDN}/operatorhubio/azure-service-operator:v${PREVIOUS_BUNDLE_VERSION},${REGISTRY_FQDN}/operatorhubio/azure-service-operator:v${NEW_BUNDLE_VERSION} \
        --tag ${REGISTRY_FQDN}/operatorhubio/azure-service-operator-index:test1
    ```
9. Push the index to the registry:
    ```sh
    docker push ${REGISTRY_FQDN}/operatorhubio/azure-service-operator-index:test1
    ```

### Create your test [CatalogSource](https://docs.openshift.com/container-platform/4.7/rest_api/operatorhub_apis/catalogsource-operators-coreos-com-v1alpha1.html)

The catalog source will point to your new index image:
```yaml
apiVersion: operators.coreos.com/v1alpha1
kind: CatalogSource
metadata:
  name: aso-test
  namespace: olm
spec:
  displayName: Source for testing ASO
  image: <registry-fqdn>/operatorhubio/azure-service-operator-index:test1
  publisher: aso
  sourceType: grpc
```

Replace the registry name with yours, and apply with `k apply -f catalogsource.yaml`.

You should see that the catalog operator has started a pod from the index image. You can watch it interacting with the index pod using:
```sh
k logs -n olm deploy/catalog-operator -f
```

### Create your [Subscription](https://docs.openshift.com/container-platform/4.7/rest_api/operatorhub_apis/subscription-operators-coreos-com-v1alpha1.html)

This will trigger the OLM to install the operator from the test catalog, starting with the previous version.
```yaml
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  labels:
    operators.coreos.com/azure-service-operator.operators: ""
  name: my-azure-service-operator
  namespace: operators
spec:
  channel: stable
  name: azure-service-operator
  source: aso-test
  sourceNamespace: olm
  installPlanApproval: Manual
  startingCSV: azure-service-operator.v<previous bundle version>
```

Adjust for your starting version and create this with `k apply -f subscription.yaml`.

The `installPlanApproval: Manual` setting tells OLM to create an unapproved [InstallPlan](https://docs.openshift.com/container-platform/4.7/rest_api/operatorhub_apis/installplan-operators-coreos-com-v1alpha1.html) that needs to be approved manually before that version of the operator will be installed.
This plus the `startingCSV` field pointing at the previous version of the operator will let you see the upgrade process at each step.

### Approve installation of the old version

You should see an unapproved InstallPlan created in the `operators` namespace for the old version of the operator bundle.
```sh
k get -n operators installplans
```

Edit or patch it to set `spec.approved` to `true`.
```sh
k edit -n operators installplans install-<hash>
```

The previous version of the operator will install and start up. You can check progress by looking at the yaml for the install plan and [ClusterServiceVersion](https://docs.openshift.com/container-platform/4.7/rest_api/operatorhub_apis/installplan-operators-coreos-com-v1alpha1.html), the logs from the catalog operator, and eventually the state of the pods in the `operators` namespace.

### Approve installation of the new version

Once the old version of the operator is installed, the OLM will create another install plan to upgrade it to the new version.
Approve that plan in the same way as the previous one and watch progress.

### Troubleshooting

If there's an issue with the new bundle, it will probably show up as a condition on the install plan (as well as logging by the catalog operator).
Problems encountered in the past include:

* Validation errors in the CSV (particularly the `webhookdefinitions` section).
* Validation errors when applying the CRDs (eg. incorrect conversion fields after the v1beta1 -> v1 switch)
* Incorrect container image reference in deployment - this manifested as the install plan stuck in Installing (no errors) but the deployed pods stuck failing to pull the image.

To back out the installed operator and retry:

1. Fix the error and push a new bundle image as above - be sure to push the bundle image to the same version tag since that's what the index refers to.
2. Delete the subscription (will also remove the install plans):
    ```sh
    k delete -n operators subscriptions my-azure-service-operator
    ```
3. Delete the CSVs (this will also remove any deployments/services/pods etc.):
    ```sh
    k delete -n operators csv azure-service-operator.v${PREVIOUS_BUNDLE_VERSION} azure-service-operator.v${NEW_BUNDLE_VERSION}
    ```
4. Remove the CRDs:
    ```sh
    k api-resources --api-group azure.microsoft.com -o name | xargs kubectl delete customresourcedefinition
    ```
5. The catalog operator creates a configmap of each bundle image to load the CSV and CRDs from.
This won't be recreated from the updated image if it already exists (even though the content is out of date), so you need to delete it.
They're named with hashes but you can look at the `olm.sourceImage` annotation to find the right one to delete.
    ```sh
    k get -n olm configmap -o jsonpath='{range .items[*]}{@.metadata.name} {@.metadata.annotations.olm\.sourceImage}{"\n"}{end}'
    k delete -n olm configmap <selected-hash>
    ```
6. Then recreating the subscription as above will restart the install/upgrade process.

## Creating PRs to add the new version to the operator repos

The embedded OpenShift operator store is populated from the [community-operators-prod](https://github.com/redhat-openshift-ecosystem/community-operators-prod) repository,
while [OperatorHub](https://operatorhub.io) is populated from the [community-operators](https://github.com/k8s-operatorhub/community-operators) repository.
(Under their respective `operators` directoryies.)

Once the bundle has been tested, PRs should be created for each repo.

1. Create a fork and local clone of the repo if you haven't already got one.
2. Create a new branch, for example `azure-service-operator-${NEW_BUNDLE_VERSION}`.
3. Copy the manifests directory to the destination folder, renaming it with the new version number.
     ```sh
     cp -r bundle/manifests ~/dev/community-operators/operators/azure-service-operator/${NEW_BUNDLE_VERSION}
     ```
 4. Edit the package file `community-operators/operators/azure-service-operator/azure-service-operator.package.yaml` to point to the new bundle version.
     ```yaml
     channels:
     - currentCSV: azure-service-operator.v<new bundle version>
       name: stable
     defaultChannel: stable
     packageName: azure-service-operator
     ```
5. Commit the changes, push them and create a PR.
