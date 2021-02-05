// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package kube

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/pkg/secrets"
)

func getExpectedSecretName(secretKey secrets.SecretKey, namingScheme secrets.SecretNamingVersion) types.NamespacedName {
	switch namingScheme {
	case secrets.SecretNamingV1:
		return types.NamespacedName{Namespace: secretKey.Namespace, Name: secretKey.Name}
	case secrets.SecretNamingV2:
		return types.NamespacedName{Namespace: secretKey.Namespace, Name: strings.ToLower(secretKey.Kind) + "-" + secretKey.Name}
	default:
		panic("unknown secret naming scheme")
	}
}

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

		supportedSecretNamingSchemes := []secrets.SecretNamingVersion{
			secrets.SecretNamingV1,
			secrets.SecretNamingV2,
		}

		for _, secretNamingScheme := range supportedSecretNamingSchemes {
			secretNamingScheme := secretNamingScheme
			It(fmt.Sprintf("should create and delete secret in k8s with secret naming scheme %q", secretNamingScheme), func() {
				secretName := "secret" + strconv.FormatInt(GinkgoRandomSeed(), 10)

				var err error
				ctx := context.Background()

				data := map[string][]byte{
					"test":  []byte("data"),
					"sweet": []byte("potato"),
				}

				client := New(k8sClient, secretNamingScheme)

				key := secrets.SecretKey{Name: secretName, Namespace: "default", Kind: "Test"}

				Context("creating secret with secret client", func() {
					err = client.Upsert(ctx, key, data)
					Expect(err).To(BeNil())
				})

				secret := &v1.Secret{}
				Context("ensuring secret exists using k8s client", func() {
					err = k8sClient.Get(ctx, getExpectedSecretName(key, secretNamingScheme), secret)
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

					err = k8sClient.Get(ctx, getExpectedSecretName(key, secretNamingScheme), secret)
					Expect(err).ToNot(BeNil())
				})
			})
		}
	})
})
