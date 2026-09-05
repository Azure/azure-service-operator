// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"strings"
	"time"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	armhcp20260901preview "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshifthcp/armredhatopenshifthcp"
	"github.com/go-logr/logr"
	"github.com/rotisserie/eris"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-service-operator/v2/api/redhatopenshift/v20260901preview/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/extensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
)

var _ extensions.PreReconciliationChecker = &HcpOpenShiftClusterExtension{}

// PreReconcileCheck does a pre-reconcile check to see if the resource is in a state that can be reconciled.
// ARM resources should implement this to avoid reconciliation attempts that cannot possibly succeed.
// Returns ProceedWithReconcile if the reconciliation should go ahead.
// Returns BlockReconcile and a human-readable reason if the reconciliation should be skipped.
// ctx is the current operation context.
// obj is the resource about to be reconciled. The resource's State will be freshly updated.
// owner is the owner of the resource being reconciled.
// resourceResolver allows resolving references to other resources.
// armClient allows access to ARM for any required queries.
// log is the logger for the current operation.
// next is the next (nested) implementation to call.
func (ext *HcpOpenShiftClusterExtension) PreReconcileCheck(ctx context.Context,
	obj genruntime.MetaObject,
	resourceResolver *resolver.Resolver,
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
	next extensions.PreReconcileCheckFunc,
) (extensions.PreReconcileCheckResult, error) {
	// This has to be the current hub storage version of the hcpOpenShiftCluster.
	// It will need to be updated if the hub storage version changes.
	hcpOpenShiftCluster, ok := obj.(*storage.HcpOpenShiftCluster)
	if !ok {
		return extensions.PreReconcileCheckResult{}, eris.Errorf("cannot run on unknown resource type %T, expected *storage.HcpOpenShiftCluster", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = hcpOpenShiftCluster

	// If the hcpOpenShiftCluster is already deleting, we have to wait for that to finish
	// before trying anything else
	if hcpOpenShiftCluster.Status.Properties != nil &&
		hcpOpenShiftCluster.Status.Properties.ProvisioningState != nil &&
		strings.EqualFold(*hcpOpenShiftCluster.Status.Properties.ProvisioningState, "Deleting") {
		return extensions.BlockReconcile("reconcile blocked while hcpOpenShiftCluster is at status deleting"), nil
	}

	return next(ctx, obj, resourceResolver, armClient, log)
}

var _ genruntime.KubernetesSecretExporter = &HcpOpenShiftClusterExtension{}

func (ext *HcpOpenShiftClusterExtension) ExportKubernetesSecrets(
	ctx context.Context,
	obj genruntime.MetaObject,
	additionalSecrets set.Set[string],
	armClient *genericarmclient.GenericClient,
	log logr.Logger,
) (*genruntime.KubernetesSecretExportResult, error) {
	// This has to be the current hub storage version. It will need to be updated
	// if the hub storage version changes.
	typedObj, ok := obj.(*storage.HcpOpenShiftCluster)
	if !ok {
		return nil, eris.Errorf("cannot run on unknown resource type %T, expected *storage.HcpOpenShiftCluster", obj)
	}

	// Type assert that we are the hub type. This will fail to compile if
	// the hub type has been changed but this extension has not
	var _ conversion.Hub = typedObj

	primarySecrets := secretsSpecifiedHcp(typedObj)
	requestedSecrets := set.Union(primarySecrets, additionalSecrets)

	if len(requestedSecrets) == 0 {
		log.V(Debug).Info("No secrets retrieval to perform as operatorSpec is empty")
		return nil, nil
	}

	id, err := genruntime.GetAndParseResourceID(typedObj)
	if err != nil {
		return nil, err
	}

	var adminCredentials string
	if requestedSecrets.Contains(adminCredentialsKey) {
		adminCredentials, err = requestAdminCredential(
			ctx, armClient, id.SubscriptionID, id.ResourceGroupName, typedObj.AzureName(), log,
		)
		if err != nil {
			return nil, err
		}
	}

	secretSlice, err := secretsToWriteHcp(typedObj, adminCredentials)
	if err != nil {
		return nil, err
	}

	resolvedSecrets := map[string]string{}
	if adminCredentials != "" {
		resolvedSecrets[adminCredentialsKey] = adminCredentials
	}
	return &genruntime.KubernetesSecretExportResult{
		Objs:       secrets.SliceToClientObjectSlice(secretSlice),
		RawSecrets: secrets.SelectSecrets(additionalSecrets, resolvedSecrets),
	}, nil
}

// requestAdminCredential retrieves an admin kubeconfig for the ARO-HCP cluster.
//
// It uses the CSR-based admin-credential API: a private key and certificate
// signing request are generated client-side, and the Azure RP returns a
// kubeconfig referencing the signed client certificate (but not the private
// key), which is then injected client-side.
func requestAdminCredential(
	ctx context.Context,
	armClient *genericarmclient.GenericClient,
	subscriptionID string,
	resourceGroupName string,
	clusterName string,
	log logr.Logger,
) (string, error) {
	// Using armClient.ClientOptions() here ensures we share the same HTTP connection, so this is not opening a new
	// connection each time through
	hcpClient, err := armhcp20260901preview.NewHcpOpenShiftClustersClient(subscriptionID, armClient.Creds(), armClient.ClientOptions())
	if err != nil {
		return "", eris.Wrapf(err, "failed to create new HcpOpenShiftClustersClient")
	}

	// Generate a private key and certificate signing request. The RP signs the CSR
	// and returns a kubeconfig referencing the signed certificate; the matching
	// private key never leaves this process and is injected into the kubeconfig below.
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", eris.Wrapf(err, "failed to generate private key for admin credential request")
	}

	csrDER, err := x509.CreateCertificateRequest(rand.Reader, &x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:   "system:customer-break-glass:system-admin",
			Organization: []string{"system:masters"},
		},
	}, privateKey)
	if err != nil {
		return "", eris.Wrapf(err, "failed to create certificate signing request for admin credentials")
	}
	csrPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrDER})

	log.V(Debug).Info("Starting CSR-based BeginRequestAdminCredential")
	poller, err := hcpClient.BeginRequestAdminCredential(
		ctx,
		resourceGroupName,
		clusterName,
		armhcp20260901preview.HcpOpenShiftClusterAdminCredentialRequest{
			CertificateSigningRequest: to.Ptr(string(csrPEM)),
		},
		nil,
	)
	if err != nil {
		return "", eris.Wrapf(err, "failed creating admin credentials")
	}

	log.V(Debug).Info("Waiting for CSR-based admin credential request to complete")
	resp, err := pollAdminCredential(ctx, poller)
	if err != nil {
		return "", err
	}

	log.V(Debug).Info("CSR-based admin credential request completed")
	kubeconfig := to.Value(resp.HcpOpenShiftClusterAdminCredential.Kubeconfig)
	if kubeconfig == "" {
		return "", eris.Errorf(
			"admin credential response for cluster %s in resource group %s contained an empty kubeconfig",
			clusterName, resourceGroupName,
		)
	}

	// The CSR-based API returns a kubeconfig whose user has the signed client
	// certificate but no private key. Inject the private key generated above so the
	// resulting kubeconfig can authenticate.
	return injectClientKey(kubeconfig, privateKey)
}

// pollAdminCredential waits for an admin credential request poller to complete,
// applying a bounded timeout and translating cancellation/timeout into
// descriptive errors.
func pollAdminCredential[T any](ctx context.Context, poller *runtime.Poller[T]) (T, error) {
	pollCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	resp, pollErr := poller.PollUntilDone(pollCtx, &runtime.PollUntilDoneOptions{
		Frequency: 15 * time.Second,
	})
	if pollErr != nil {
		if ctx.Err() != nil {
			return resp, eris.Wrapf(pollErr, "parent context cancelled while waiting for admin credentials")
		}
		if pollCtx.Err() == context.DeadlineExceeded {
			return resp, eris.Wrapf(pollErr, "timed out after 5 minutes waiting for admin credentials to be ready")
		}
		return resp, eris.Wrapf(pollErr, "failed waiting for admin credentials to be ready")
	}
	return resp, nil
}

// injectClientKey loads the given kubeconfig, sets the supplied private key as the
// client key for every user entry, and returns the re-serialized kubeconfig.
func injectClientKey(kubeconfig string, privateKey *rsa.PrivateKey) (string, error) {
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	config, err := clientcmd.Load([]byte(kubeconfig))
	if err != nil {
		return "", eris.Wrapf(err, "failed to load admin kubeconfig")
	}

	for _, authInfo := range config.AuthInfos {
		authInfo.ClientKeyData = privateKeyPEM
	}

	merged, err := clientcmd.Write(*config)
	if err != nil {
		return "", eris.Wrapf(err, "failed to serialize admin kubeconfig")
	}

	return string(merged), nil
}

func secretsSpecifiedHcp(obj *storage.HcpOpenShiftCluster) set.Set[string] {
	if obj.Spec.OperatorSpec == nil || obj.Spec.OperatorSpec.Secrets == nil {
		return nil
	}

	operatorSecrets := obj.Spec.OperatorSpec.Secrets
	result := set.Set[string]{}
	if operatorSecrets.AdminCredentials != nil {
		result.Add(adminCredentialsKey)
	}

	return result
}

func secretsToWriteHcp(obj *storage.HcpOpenShiftCluster, adminCredentials string) ([]*v1.Secret, error) {
	operatorSpecSecrets := obj.Spec.OperatorSpec.Secrets
	if operatorSpecSecrets == nil {
		return nil, nil
	}

	collector := secrets.NewCollector(obj.Namespace)
	collector.AddValue(operatorSpecSecrets.AdminCredentials, adminCredentials)

	return collector.Values()
}
