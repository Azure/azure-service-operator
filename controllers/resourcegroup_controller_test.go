/*
Copyright 2019 microsoft.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	resourcemanagerconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

	resoucegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestResourceGroup(t *testing.T) {
	//g := NewGomegaWithT(t)
	RegisterTestingT(t)
	resourcemanagerconfig.LoadSettings()

	const timeout = time.Second * 240
	const poll = time.Second * 10
	resourceGroupName := "t-rg-dev-" + helpers.RandomString(10)

	// Create the Resourcegroup object and expect the Reconcile to be created
	resourceGroupInstance := &azurev1.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourceGroupName,
			Namespace: "default",
		},
		Spec: azurev1.ResourceGroupSpec{
			Location: "westus",
		},
	}

	// send the resourceGroup to kubernetes
	err := k8sClient.Create(context.Background(), resourceGroupInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	// prep query gor Get
	// @todo consider random namespaces?
	resourceGroupNamespacedName := types.NamespacedName{Name: resourceGroupName, Namespace: "default"}

	// wait until entity has been submitted
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
		return resourceGroupInstance.IsSubmitted()
	}, timeout,
	).Should(BeTrue())

	// wait until resource is provisioned
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
		return resourceGroupInstance.Status.Provisioned == true
	}, timeout,
	).Should(BeTrue())

	// verify cloud resource exists in Azure
	Eventually(func() bool {
		result, _ := resoucegroupsresourcemanager.CheckExistence(context.Background(), resourceGroupName)
		return result.Response.StatusCode == 204
	}, timeout, poll,
	).Should(BeTrue())

	// delete resoruce group and then verify
	k8sClient.Delete(context.Background(), resourceGroupInstance)

	// has the operator set the proper status for deletion?
	Eventually(func() bool {
		_ = k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
		return resourceGroupInstance.IsBeingDeleted()
	}, timeout,
	).Should(BeTrue())

	// is the resource now gone from kubernetes?
	Eventually(func() bool {
		err := k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
		if err == nil {
			err = fmt.Errorf("")
		}
		return strings.Contains(err.Error(), "not found")
	}, timeout,
	).Should(BeTrue())

	// make sure the resource is gone from Azure
	Eventually(func() bool {
		result, _ := resoucegroupsresourcemanager.CheckExistence(context.Background(), resourceGroupName)
		return result.Response.StatusCode == 404
	}, timeout, poll,
	).Should(BeTrue())
}
