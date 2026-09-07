/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"reflect"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/Azure/azure-service-operator/v2/api/redhatopenshift/v20260901preview/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	testreflect "github.com/Azure/azure-service-operator/v2/internal/testcommon/reflect"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
)

const sampleKubeconfig = `
apiVersion: v1
kind: Config
clusters:
- cluster:
    server: https://api.hcp.example.com:6443
    certificate-authority-data: Y2FEYXRh
  name: hcp
contexts:
- context:
    cluster: hcp
    user: system-admin
  name: admin
current-context: admin
users:
- name: system-admin
  user:
    client-certificate-data: Y2VydERhdGE=
`

func generateTestKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate test key: %s", err)
	}
	return key
}

func Test_InjectClientKey_SetsClientKeyDataForEveryUser(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	privateKey := generateTestKey(t)
	expectedKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	result, err := injectClientKey(sampleKubeconfig, privateKey)
	g.Expect(err).ToNot(HaveOccurred())

	config, err := clientcmd.Load([]byte(result))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(config.AuthInfos).To(HaveLen(1))

	authInfo, ok := config.AuthInfos["system-admin"]
	g.Expect(ok).To(BeTrue())
	g.Expect(authInfo.ClientKeyData).To(Equal(expectedKeyPEM))

	// The signed client certificate returned by the RP must be preserved unchanged.
	g.Expect(authInfo.ClientCertificateData).To(Equal([]byte("certData")))

	// Unrelated kubeconfig content (clusters, contexts) must survive the round-trip.
	g.Expect(config.Clusters).To(HaveKey("hcp"))
	g.Expect(config.Clusters["hcp"].Server).To(Equal("https://api.hcp.example.com:6443"))
	g.Expect(config.CurrentContext).To(Equal("admin"))
}

func Test_InjectClientKey_SetsSameKeyForMultipleUsers(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	const multiUserKubeconfig = `
apiVersion: v1
kind: Config
clusters:
- cluster:
    server: https://api.hcp.example.com:6443
  name: hcp
contexts:
- context:
    cluster: hcp
    user: system-admin
  name: admin
current-context: admin
users:
- name: system-admin
  user:
    client-certificate-data: Y2VydERhdGE=
- name: other-user
  user:
    client-certificate-data: b3RoZXJDZXJ0
`

	privateKey := generateTestKey(t)
	expectedKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	result, err := injectClientKey(multiUserKubeconfig, privateKey)
	g.Expect(err).ToNot(HaveOccurred())

	config, err := clientcmd.Load([]byte(result))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(config.AuthInfos).To(HaveLen(2))
	for name, authInfo := range config.AuthInfos {
		g.Expect(authInfo.ClientKeyData).To(Equal(expectedKeyPEM), "user %s should have the injected client key", name)
	}
}

func Test_InjectClientKey_InvalidKubeconfig_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	privateKey := generateTestKey(t)

	_, err := injectClientKey("not: [a, valid, kubeconfig", privateKey)
	g.Expect(err).To(HaveOccurred())
}

func Test_SecretsSpecifiedHcp_NilOperatorSpec_ReturnsEmpty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cluster := &storage.HcpOpenShiftCluster{}
	g.Expect(secretsSpecifiedHcp(cluster)).To(BeEmpty())
}

func Test_SecretsSpecifiedHcp_NilSecrets_ReturnsEmpty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cluster := &storage.HcpOpenShiftCluster{
		Spec: storage.HcpOpenShiftCluster_Spec{
			OperatorSpec: &storage.HcpOpenShiftClusterOperatorSpec{},
		},
	}
	g.Expect(secretsSpecifiedHcp(cluster)).To(BeEmpty())
}

// Test_SecretsSpecifiedHcp_AllSecretsSpecifiedAllSecretsReturned guards against a new field being
// added to HcpOpenShiftClusterOperatorSecrets without secretsSpecifiedHcp being updated to report it.
func Test_SecretsSpecifiedHcp_AllSecretsSpecifiedAllSecretsReturned(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	operatorSecrets := &storage.HcpOpenShiftClusterOperatorSecrets{}
	testreflect.PopulateStruct(operatorSecrets)

	cluster := &storage.HcpOpenShiftCluster{
		Spec: storage.HcpOpenShiftCluster_Spec{
			OperatorSpec: &storage.HcpOpenShiftClusterOperatorSpec{
				Secrets: operatorSecrets,
			},
		},
	}

	secretNames := secretsSpecifiedHcp(cluster)
	expectedTags := reflecthelpers.GetJSONTags(reflect.TypeOf(storage.HcpOpenShiftClusterOperatorSecrets{}))
	// $propertyBag exists because this is the storage version, but it isn't a secret.
	expectedTags.Remove("$propertyBag")

	g.Expect(secretNames).To(HaveLen(len(expectedTags)))
}

func Test_SecretsToWriteHcp_NilSecrets_ReturnsNil(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cluster := &storage.HcpOpenShiftCluster{
		Spec: storage.HcpOpenShiftCluster_Spec{
			OperatorSpec: &storage.HcpOpenShiftClusterOperatorSpec{},
		},
	}

	result, err := secretsToWriteHcp(cluster, "irrelevant")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(BeNil())
}

func Test_SecretsToWriteHcp_AdminCredentials_WritesRequestedSecret(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cluster := &storage.HcpOpenShiftCluster{
		ObjectMeta: metav1.ObjectMeta{Namespace: "hcp-ns"},
		Spec: storage.HcpOpenShiftCluster_Spec{
			OperatorSpec: &storage.HcpOpenShiftClusterOperatorSpec{
				Secrets: &storage.HcpOpenShiftClusterOperatorSecrets{
					AdminCredentials: &genruntime.SecretDestination{
						Name: "hcp-secret",
						Key:  "adminCreds",
					},
				},
			},
		},
	}

	result, err := secretsToWriteHcp(cluster, "kubeconfig-contents")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result).To(HaveLen(1))
	g.Expect(result[0].Name).To(Equal("hcp-secret"))
	g.Expect(result[0].Namespace).To(Equal("hcp-ns"))
	g.Expect(result[0].StringData).To(HaveKeyWithValue("adminCreds", "kubeconfig-contents"))
}

func Test_SecretsToWriteHcp_RequestedButEmptyValue_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cluster := &storage.HcpOpenShiftCluster{
		ObjectMeta: metav1.ObjectMeta{Namespace: "hcp-ns"},
		Spec: storage.HcpOpenShiftCluster_Spec{
			OperatorSpec: &storage.HcpOpenShiftClusterOperatorSpec{
				Secrets: &storage.HcpOpenShiftClusterOperatorSecrets{
					AdminCredentials: &genruntime.SecretDestination{
						Name: "hcp-secret",
						Key:  "adminCreds",
					},
				},
			},
		},
	}

	// An empty admin credential string means we asked for the secret but have no value to put in it.
	_, err := secretsToWriteHcp(cluster, "")
	g.Expect(err).To(HaveOccurred())
}

func Test_PreReconcileCheck_ClusterDeleting_BlocksReconcile(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cluster := &storage.HcpOpenShiftCluster{
		Status: storage.HcpOpenShiftCluster_STATUS{
			Properties: &storage.HcpOpenShiftClusterProperties_STATUS{
				ProvisioningState: to.Ptr("Deleting"),
			},
		},
	}

	ext := &HcpOpenShiftClusterExtension{}
	nextCalled := false
	next := func(
		_ context.Context,
		_ genruntime.MetaObject,
		_ *resolver.Resolver,
		_ *genericarmclient.GenericClient,
		_ logr.Logger,
	) (extensions.PreReconcileCheckResult, error) {
		nextCalled = true
		return extensions.ProceedWithReconcile(), nil
	}

	result, err := ext.PreReconcileCheck(context.Background(), cluster, nil, nil, logr.Discard(), next)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.BlockReconciliation()).To(BeTrue())
	g.Expect(nextCalled).To(BeFalse())
}

func Test_PreReconcileCheck_ClusterDeletingCaseInsensitive_BlocksReconcile(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	cluster := &storage.HcpOpenShiftCluster{
		Status: storage.HcpOpenShiftCluster_STATUS{
			Properties: &storage.HcpOpenShiftClusterProperties_STATUS{
				ProvisioningState: to.Ptr("deleting"),
			},
		},
	}

	ext := &HcpOpenShiftClusterExtension{}
	result, err := ext.PreReconcileCheck(context.Background(), cluster, nil, nil, logr.Discard(),
		func(
			_ context.Context,
			_ genruntime.MetaObject,
			_ *resolver.Resolver,
			_ *genericarmclient.GenericClient,
			_ logr.Logger,
		) (extensions.PreReconcileCheckResult, error) {
			return extensions.ProceedWithReconcile(), nil
		})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(result.BlockReconciliation()).To(BeTrue())
}

func Test_PreReconcileCheck_ClusterNotDeleting_CallsNext(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	for _, state := range []*string{nil, to.Ptr("Succeeded"), to.Ptr("Failed")} {
		cluster := &storage.HcpOpenShiftCluster{}
		if state != nil {
			cluster.Status.Properties = &storage.HcpOpenShiftClusterProperties_STATUS{ProvisioningState: state}
		}

		ext := &HcpOpenShiftClusterExtension{}
		nextCalled := false
		next := func(
			_ context.Context,
			_ genruntime.MetaObject,
			_ *resolver.Resolver,
			_ *genericarmclient.GenericClient,
			_ logr.Logger,
		) (extensions.PreReconcileCheckResult, error) {
			nextCalled = true
			return extensions.ProceedWithReconcile(), nil
		}

		result, err := ext.PreReconcileCheck(context.Background(), cluster, nil, nil, logr.Discard(), next)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(result.BlockReconciliation()).To(BeFalse())
		g.Expect(nextCalled).To(BeTrue())
	}
}

func Test_PreReconcileCheck_WrongResourceType_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	wrongType := &storage.HcpOpenShiftClustersNodePool{}

	ext := &HcpOpenShiftClusterExtension{}
	_, err := ext.PreReconcileCheck(context.Background(), wrongType, nil, nil, logr.Discard(),
		func(
			_ context.Context,
			_ genruntime.MetaObject,
			_ *resolver.Resolver,
			_ *genericarmclient.GenericClient,
			_ logr.Logger,
		) (extensions.PreReconcileCheckResult, error) {
			t.Fatal("next should not be called when the resource type assertion fails")
			return extensions.PreReconcileCheckResult{}, nil
		})
	g.Expect(err).To(HaveOccurred())
}
