/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// CreateDeploymentName generates a unique deployment name
func CreateDeploymentName() (string, error) {
	// no status yet, so start provisioning
	deploymentUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	deploymentName := fmt.Sprintf("%s_%d_%s", "k8s", time.Now().Unix(), deploymentUUID.String())
	return deploymentName, nil
}
