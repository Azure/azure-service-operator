/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

// IsWebhookPackageReference returns true if the reference is to a Webhook package OR to a subpackage of
// a Webhook package, false otherwise.
func IsWebhookPackageReference(reference PackageReference) bool {
	sub, ok := reference.(SubPackageReference)
	if !ok {
		return false
	}

	if sub.name == WebhookPackageName {
		return true
	}

	return IsWebhookPackageReference(sub.parent)
}
