/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	app "github.com/Azure/azure-service-operator/v2/api/app/v1api20240301"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_App_ContainerApps_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	managedEnv := newManagedEnvironment(tc, rg)
	tc.CreateResourcesAndWait(managedEnv)

	containerApp := newContainerApp(tc, managedEnv, rg)

	tc.CreateResourceAndWait(containerApp)

	tc.Expect(containerApp.Status.Id).ToNot(BeNil())
	armID := *containerApp.Status.Id

	old := containerApp.DeepCopy()
	containerApp.Spec.Tags = map[string]string{"foo": "bar"}
	tc.PatchResourceAndWait(old, containerApp)

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "AuthConfig CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				ContainerApps_AuthConfig_CRUD(tc, containerApp)
			},
		},
	)

	tc.DeleteResourceAndWait(containerApp)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armID,
		string(app.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func ContainerApps_AuthConfig_CRUD(tc *testcommon.KubePerTestContext, containerApp *app.ContainerApp) {
	authConfig := &app.AuthConfig{
		ObjectMeta: tc.MakeObjectMetaWithName("current"),
		Spec: app.AuthConfig_Spec{
			HttpSettings: &app.HttpSettings{
				RequireHttps: to.Ptr(true),
			},
			Owner: testcommon.AsOwner(containerApp),
		},
	}

	tc.CreateResourceAndWait(authConfig)
}

func newContainerApp(tc *testcommon.KubePerTestContext, managedEnv *app.ManagedEnvironment, rg *resources.ResourceGroup) *app.ContainerApp {
	containerApp := &app.ContainerApp{
		ObjectMeta: tc.MakeObjectMeta("app"),
		Spec: app.ContainerApp_Spec{
			Configuration: &app.Configuration{
				Ingress: &app.Ingress{
					AllowInsecure: to.Ptr(false),
					TargetPort:    to.Ptr(80),
				},
			},
			EnvironmentReference: tc.MakeReferenceFromResource(managedEnv),
			Location:             tc.AzureRegion,
			Owner:                testcommon.AsOwner(rg),
			Template: &app.Template{
				Containers: []app.Container{
					{
						Image: to.Ptr("nginx:latest"),
						Name:  to.Ptr("nginx"),
						VolumeMounts: []app.VolumeMount{
							{
								MountPath:  to.Ptr("/usr/share/nginx/html"),
								VolumeName: to.Ptr("shared"),
							},
						},
					},
				},
				InitContainers: []app.BaseContainer{
					{
						Args: []string{
							"-c",
							"echo Hello World",
						},
						Command: []string{"/bin/sh"},
						Image:   to.Ptr("debian:latest"),
						Name:    to.Ptr("debian"),
						VolumeMounts: []app.VolumeMount{
							{
								MountPath:  to.Ptr("/shared"),
								VolumeName: to.Ptr("shared"),
							},
						},
					},
				},
				Volumes: []app.Volume{
					{
						Name:        to.Ptr("shared"),
						StorageType: to.Ptr(app.Volume_StorageType_EmptyDir),
					},
				},
			},
		},
	}
	return containerApp
}
