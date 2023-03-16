package identity

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/pkg/errors"
)

func NewClientCertificateCredential(tenantID, clientID string, clientCertificate, password []byte) (*azidentity.ClientCertificateCredential, error) {
	certs, key, err := azidentity.ParseCertificates([]byte(clientCertificate), password)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse certificate for '%s': %v", clientID)
	}

	cred, err := azidentity.NewClientCertificateCredential(tenantID, clientID, certs, key, nil)
	if err != nil {
		return nil, err
	}
	return cred, nil

}
