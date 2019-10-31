package reconciler

import "strings"

const (
	AccessPermissionAnnotation = "/access-permissions"
)

type AccessPermissions string

func (p AccessPermissions) create() bool { return containsOrEmpty(p, "c") }
func (p AccessPermissions) update() bool { return containsOrEmpty(p, "u") }
func (p AccessPermissions) delete() bool { return containsOrEmpty(p, "d") }

// Note that all permissions are enabled by default
func containsOrEmpty(p AccessPermissions, c string) bool {
	return strings.Contains(strings.ToLower(string(p)), c) || p == ""
}
