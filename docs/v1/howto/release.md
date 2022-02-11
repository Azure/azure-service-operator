# Releasing Azure Service Operator

The release process for Azure Service Operator has two steps:
1. Release an updated container image for the operator manager.
2. Reference the updated container image in a new Helm chart.

## Releasing an updated container image for the manager

1. Ensure you have a successful pipeline run off `main` that has all the changes you want to be in the release. If not, [schedule a run on Azure DevOps](http://dev.azure.com/azure/azure-service-operator). Note that you need to be signed in on DevOps to schedule a build.
   - The output of this step is a _candidate_ container image.
2. Create a new [DevOps release](https://dev.azure.com/azure/azure-service-operator/_release?_a=releases&view=mine&definitionId=2). **Note**: The release must be created from a build of `main`. The release creation will fail if it is created targetting artifacts from any other branch.
   - Click "create release" and choose the build from step 1 for "Artifacts".
   - Click "Create".
3. Wait for the DevOps release created in step #2 to finish. The result of a successful run of the DevOps release is:
   - The creation of a GitHub release you can access at https://github.com/Azure/azure-service-operator/releases. This release contains:
     - A `setup.yaml` file complete with all of the Azure Service Operator CRDs and deployment configuration which can easily be applied directly to a cluster.
     - `notes.txt` with a link to the Azure Service Operator manager docker image.


## Releasing an updated Helm chart

Because the Helm chart is hosted out of the GitHub repo itself, we cannot update it until an official container image for the ASO manager has been produced. Once the DevOps release has been run and a new image has been published, perform the following steps to generate a new version of the helm chart:

1. Create a new branch of ASO.
2. Update the `version` field of `charts/azure-service-operator/Chart.yaml`. Note that this field is the version of the _chart_, so it should follow semver for the chart. If there's a breaking change in the chart then the major version should be incremented, otherwise incrementing the minor version is fine.
3. Update the `appVersion` field of `charts/azure-service-operator/Chart.yaml` to the version of the container published in the [section above](#Releasing-an-updated-container-image-for-the-manager). Currently, this is the same version as the latest [tag](https://github.com/Azure/azure-service-operator/tags).
4. Run `make helm-chart-manifests`. You should see a new chart `.tgz` get created and the new chart version you defined referenced in `charts/index.yaml`.
5. Submit a PR to ASO with the updated `Chart.yaml`, `index.yaml` and chart `.tgz` file.
