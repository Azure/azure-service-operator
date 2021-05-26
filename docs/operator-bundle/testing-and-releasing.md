# Testing and releasing the operatorhub bundle for ASO

This is a guide to producing and testing the operator bundle after an ASO release.

## Test upgrading from the old bundle to the new one

### Before you begin

You'll need:

* A previous operator hub bundle (for me this is `quay.io/operatorhubio/azure-service-operator:v0.37.0`)
* A newly published operator image that needs wrapping in an operator bundle (eg. `mcr.microsoft.com/k8s/azureserviceoperator:1.0.22275`)
* A cluster with OLM installed (but without ASO installed). Installing OLM:
```sh
curl -sL https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.18.1/install.sh | bash -s v0.18.1
```
* An attached registry that you can push bundle and index images to (in my case this is `xtianregistry.azurecr.io`).
  * If it's an ACR then log into it with:
```sh
az login
az acr login --name <registry-name>
```
* The `opm` tool to build the index image:
  * Follow [these instructions](https://docs.openshift.com/container-platform/4.5/operators/admin/olm-managing-custom-catalogs.html#olm-installing-opm_olm-managing-custom-catalogs) to get the `opm` binary.
  * You can use docker rather than podman to log into the redhat registry, but you'll need to adjust the path to the auth credentials to match docker.
  * The `oc image extract` command is convenient for getting the binary out of the right image layer. If you don't have/want oc installed you can find it manually using `docker image save` and `tar` to find and extract it.

### Make the bundle image and index

1. Run `make generate-operator-bundle PREVIOUS_BUNDLE_VERSION=<the last released bundle version>`, or you can update the default value for `PREVIOUS_BUNDLE_VERSION` in the Makefile and commit the change - this is probably less error-prone in the long run.
2. Build the new bundle image:
    ```sh
    docker build -f bundle.Dockerfile --tag <registry-fqdn>/operatorhubio/azure-service-operator:v<new operator version> .
    ```
3. Push the new bundle to the registry:
    ```sh
    docker push <registry-fqdn>/operatorhubio/azure-service-operator:v<new operator version>
    ```
4. Pull the previous bundle version from the `quay.io` registry:
    ```sh
    docker pull quay.io/operatorhubio/azure-service-operator:v<previous operator version>
    ```
5. Tag the previous bundle for your test registry:
    ```sh
    docker tag quay.io/operatorhubio/azure-service-operator:v<previous operator version> <registry-fqdn>/operatorhubio/azure-service-operator:v<previous operator version>
    ```
6. Push that bundle to your registry as well:
    ```sh
    docker push <registry-fqdn> /operatorhubio/azure-service-operator:v<previous operator version>
    ```
7. Make the index image that points to the two bundles:
    ```sh
    opm -c docker index add  --bundles xtianregistry.azurecr.io/operatorhubio/azure-service-operator:v0.37.0,xtianregistry.azurecr.io/operatorhubio/azure-service-operator:v1.0.22275 --tag xtianregistry.azurecr.io/operatorhubio/azure-service-operator-index:test1
    ```
8. Push the index to the registry:
    ```sh
    docker push xtianregistry.azurecr.io/operatorhubio/azure-service-operator-index:test1
    ```

### Create your test [CatalogSource](https://docs.openshift.com/container-platform/4.7/rest_api/operatorhub_apis/catalogsource-operators-coreos-com-v1alpha1.html)

The catalog source will point to your new index image:
```yaml
apiVersion: operators.coreos.com/v1alpha1
kind: CatalogSource
metadata:
  name: xtian-test
  namespace: olm
spec:
  displayName: Source for testing ASO
  image: xtianregistry.azurecr.io/operatorhubio/azure-service-operator-index:test1
  publisher: xtian
  sourceType: grpc
```

(Replace the registry fqdn and names with your ones.)

You should see that the catalog operator has started a pod from the index image. You can watch it interacting with the index pod using:
```sh
k logs -n olm deploy/catalog-operator -f
```

### Create your [Subscription](https://docs.openshift.com/container-platform/4.7/rest_api/operatorhub_apis/subscription-operators-coreos-com-v1alpha1.html)

This will trigger the OLM to install the operator from the test catalog.
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
  source: xtian-test
  sourceNamespace: olm
  installPlanApproval: Manual
  startingCSV: azure-service-operator.v0.37.0
```

(Adjust for your catalog source and starting version.)

The `installPlanApproval: Manual` setting tells OLM to create an unapproved [InstallPlan](https://docs.openshift.com/container-platform/4.7/rest_api/operatorhub_apis/installplan-operators-coreos-com-v1alpha1.html) that needs to be approved manually before that version of the operator will be installed.
This plus the `startingCSV` field pointing at the previous version of the operator will let you see the upgrade process at each step.

### Approve installation of the old version

You should see an unapproved InstallPlan created in the `operators` namespace for the old version of the operator.
```sh
k get -n operators installplans
```

Edit or patch it to set `spec.approved` to `true`.
```sh
k edit -n operators installplans install-abcde
```

The previous version of the operator will install and start up. You can check progress by looking at the yaml for the install plan and [ClusterServiceVersion](https://docs.openshift.com/container-platform/4.7/rest_api/operatorhub_apis/installplan-operators-coreos-com-v1alpha1.html), the logs from the catalog operator, and eventually the state of the pods in the `operators` namespace.

### Approve installation of the new version

Once the old version of the operator is installed, the OLM will create another install plan to upgrade it to the new version.
Approve that plan in the same way as the previous one and watch progress.

### Troubleshooting

If there's an issue with the new bundle, it will probably show up as a condition on the install plan (as well as logging by the catalog operator).
Problems I've seen in the past include:

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
    k delete -n operators csv azure-service-operator.v<old> azure-service-operator.v<new>
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

## Creating PRs to add the new version to the [community-operators](https://github.com/operator-framework/community-operators) repo

The embedded OpenShift operator store and [OperatorHub](https://operatorhub.io) are both populated from the community-operators repository.
The OpenShift store bundles are under the `community-operators` directory, while OperatorHub ones are under `upstream-community-operators`.

Once the bundle has been tested, separate PRs should be created for each of the two destination directories.

1. Create a fork and local clone of the `community-operators` repo if you haven't already got one.
2. Create a new branch, for example `upstream-azure-service-operator-1.0.23956`.
3. Copy the manifests directory to the destination folder, renaming it with the new version number.
     ```sh
     cp -r bundle/manifests ~/dev/community-operators/upstream-community-operators/azure-service-operator/1.0.23956
     ```
 4. Edit the package file `community-operators/upstream-community-operators/azure-service-operator/azure-service-operator.package.yaml` to point to the new bundle version.
     ```yaml
     channels:
     - currentCSV: azure-service-operator.v1.0.23956
       name: stable
     defaultChannel: stable
     packageName: azure-service-operator
     ```
5. Commit the changes, push them and create a PR.
