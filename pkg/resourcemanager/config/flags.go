// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config

import (
	"flag"
)

// AddFlags adds flags applicable to all services.
// Remember to call `flag.Parse()` in your main or TestMain.
func AddFlags() error {
	flag.StringVar(&creds.subscriptionID, "subscription", creds.subscriptionID, "Subscription for tests.")
	flag.StringVar(&locationDefault, "location", locationDefault, "Default location for tests.")
	flag.StringVar(&cloudName, "cloud", cloudName, "Name of Azure cloud.")
	flag.StringVar(&creds.operatorKeyvault, "operatorKeyvault", creds.operatorKeyvault, "Keyvault operator uses to store secrets.")
	flag.BoolVar(&useDeviceFlow, "useDeviceFlow", useDeviceFlow, "Use device-flow grant type rather than client creds.")
	flag.BoolVar(&creds.useManagedIdentity, "useMI", creds.useManagedIdentity, "Use managed identity authentication (aad-pod-identity).")
	flag.BoolVar(&keepResources, "keepResources", keepResources, "Keep resources created by samples.")

	return nil
}
