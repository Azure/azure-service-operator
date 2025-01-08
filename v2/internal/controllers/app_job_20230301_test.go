/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	app "github.com/Azure/azure-service-operator/v2/api/app/v1api20240301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_App_Job_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	managedEnv := newManagedEnvironment(tc, rg)
	tc.CreateResourcesAndWait(managedEnv)

	job := &app.Job{
		ObjectMeta: tc.MakeObjectMeta("job"),
		Spec: app.Job_Spec{
			Configuration: &app.JobConfiguration{
				EventTriggerConfig: nil,
				ManualTriggerConfig: &app.JobConfiguration_ManualTriggerConfig{
					Parallelism:            to.Ptr(1),
					ReplicaCompletionCount: to.Ptr(1),
				},
				ReplicaRetryLimit: to.Ptr(10),
				ReplicaTimeout:    to.Ptr(10),
				TriggerType:       to.Ptr(app.JobConfiguration_TriggerType_Manual),
			},
			EnvironmentReference: tc.MakeReferenceFromResource(managedEnv),
			Location:             tc.AzureRegion,
			Owner:                testcommon.AsOwner(rg),
			Template: &app.JobTemplate{
				Containers: []app.Container{
					{
						Image: to.Ptr("mcr.microsoft.com/k8se/quickstart-jobs:latest"),
						Name:  to.Ptr("testcontainerappsjob"),
					},
				},
			},
		},
	}

	tc.CreateResourcesAndWait(job)

	tc.Expect(job.Status.Id).ToNot(BeNil())
	armId := *job.Status.Id

	old := job.DeepCopy()
	job.Spec.Tags = map[string]string{"foo": "bar"}
	tc.PatchResourceAndWait(old, job)

	tc.DeleteResourceAndWait(job)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(app.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
