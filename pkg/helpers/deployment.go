package helpers

// IsDeploymentComplete will dtermine if the deployment is complete
func IsDeploymentComplete(status string) bool {
	switch status {
	case "Succeeded":
		return true
	case "Failed":
		return true
	case "Canceled":
		return true
	}
	return false
}
