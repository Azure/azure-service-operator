// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package kube

import (
	"context"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("Kube Secrets Client", func() {

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		It("should create and delete secret in k8s", func() {
			//s := strconv.FormatInt(GinkgoRandomSeed(), 10)
			secretName := "secret" + strconv.FormatInt(GinkgoRandomSeed(), 10)

			var err error
			ctx := context.Background()

			data := map[string][]byte{
				"test":  []byte("data"),
				"sweet": []byte("potato"),
			}

			client := New(K8sClient)

			key := types.NamespacedName{Name: secretName, Namespace: "default"}

			Context("creating secret with secret client", func() {
				err = client.Create(ctx, key, data)
				Expect(err).To(BeNil())
			})

			secret := &v1.Secret{}
			Context("ensuring secret exists using k8s client", func() {
				err = K8sClient.Get(ctx, key, secret)
				Expect(err).To(BeNil())
				d, err := client.Get(ctx, key)
				Expect(err).To(BeNil())

				for k, v := range d {
					Expect(data[k]).To(Equal(v))
				}
			})

			Context("delete secret and ensure it is gone", func() {
				err = client.Delete(ctx, key)
				Expect(err).To(BeNil())

				err = K8sClient.Get(ctx, key, secret)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
