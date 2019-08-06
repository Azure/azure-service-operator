package config

import (
	"flag"
)

// AddFlags adds flags applicable to all services.
// Remember to call `flag.Parse()` in your main or TestMain.
func AddFlags() error {
	flag.StringVar(&subscriptionID, "subscription", subscriptionID, "Subscription for tests.")
	flag.StringVar(&cloudName, "cloud", cloudName, "Name of Azure cloud.")
	flag.BoolVar(&useDeviceFlow, "useDeviceFlow", useDeviceFlow, "Use device-flow grant type rather than client credentials.")

	return nil
}
