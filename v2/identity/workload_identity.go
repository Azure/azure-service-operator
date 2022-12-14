/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package identity

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

var (
	// #nosec
	TokenFile = "/var/run/secrets/tokens/azure-identity"
)

// The below bit is from https://github.com/Azure/azure-sdk-for-go/issues/15615#issuecomment-1211012677
type workloadIdentityCredential struct {
	assertion     string
	tokenFilePath string
	cred          *azidentity.ClientAssertionCredential
	lastRead      time.Time
	lock          sync.Mutex
}

func NewWorkloadIdentityCredential(tenantID, clientID string) (*workloadIdentityCredential, error) {
	w := &workloadIdentityCredential{tokenFilePath: TokenFile}
	cred, err := azidentity.NewClientAssertionCredential(tenantID, clientID, w.getAssertion, nil)
	if err != nil {
		return nil, err
	}

	w.cred = cred
	return w, nil
}

func (w *workloadIdentityCredential) GetToken(ctx context.Context, opts policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return w.cred.GetToken(ctx, opts)
}

func (w *workloadIdentityCredential) getAssertion(context.Context) (string, error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	if now := time.Now(); w.lastRead.Add(5 * time.Minute).Before(now) {
		content, err := os.ReadFile(w.tokenFilePath)
		if err != nil {
			return "", err
		}

		w.assertion = string(content)
		w.lastRead = now
	}

	return w.assertion, nil
}
