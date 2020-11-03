// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config

import (
	"flag"
)

// AddFlags adds flags applicable to all services.
// Remember to call `flag.Parse()` in your main or TestMain.
func AddFlags() error {
	flag.StringVar(&credentials.subscriptionID, "subscription", credentials.subscriptionID, "Subscription for tests.")
	flag.StringVar(&locationDefault, "location", locationDefault, "Default location for tests.")
	flag.StringVar(&cloudName, "cloud", cloudName, "Name of Azure cloud.")
	flag.StringVar(&credentials.operatorKeyvault, "operatorKeyvault", credentials.operatorKeyvault, "Keyvault operator uses to store secrets.")
	flag.BoolVar(&useDeviceFlow, "useDeviceFlow", useDeviceFlow, "Use device-flow grant type rather than client credentials.")
	flag.BoolVar(&credentials.useMI, "useMI", credentials.useMI, "Use MI authentication (aad-pod-identity).")
	flag.BoolVar(&keepResources, "keepResources", keepResources, "Keep resources created by samples.")

	return nil
}
