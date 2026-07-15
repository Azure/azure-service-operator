/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "strings"

// VersionMigrationMode indicates how version migration should be handled
type VersionMigrationMode string

const (
	// VersionMigrationModeNone indicates that no version migration is needed.
	// All packages generated for this group will use the new style, v{date} format.
	//
	// This is the default for any new groups used for new resources, as well
	// as for any golden tests within the generator.
	VersionMigrationModeNone VersionMigrationMode = "none"

	// VersionMigrationModeLegacy indicates that older style versions will be generated
	// for resource versions introduced up to ASO version 2.16.0, and new style versions
	// for resource versions introduced in ASO version 2.17.0 and later.
	//
	// To achieve this, older style versions will be generated for all packages,
	// and a new pipeline stage will *move* resources introduced in ASO version 2.17.0 and later
	// into new packages with new style versions.
	//
	// This mode will be configured for any groups existing prior to ASO version 2.17.0,
	// creating a stable reference point from which we can migrate to new style versions
	// in a controlled manner.
	VersionMigrationModeLegacy VersionMigrationMode = "legacy"

	// VersionMigrationModeHybrid indicates that new style versions will be generated for all
	// resources, with old style versions also generated for resources introduced up to ASO
	// version v2.16.0 for backward compatibility.
	//
	// To achieve this, new style versions will be generated for all packages,
	// and a new pipeline stage will create older style versions for resources existing
	// in ASO version 2.16.0 and earlier to provide a migration path forward for existing users.
	VersionMigrationModeHybrid VersionMigrationMode = "hybrid"
)

// LastLegacyASOVersion is the last ASO release that used the legacy `v1api` prefixed versioning scheme
// for all resources. Resources introduced in this release or earlier that live in Hybrid mode groups
// have both a `v1api`-prefixed variant (retained for backward compatibility) and a new `v`-prefixed
// variant. The `v`-prefixed variant only becomes available in the ASO release in which the group is
// migrated to Hybrid mode - see versionMigrationHybridReleases.
const LastLegacyASOVersion = "v2.16.0"

// versionMigrationHybridReleases records the ASO release in which each Hybrid mode group was
// migrated from Legacy versioning. This mirrors the "Hybrid" column of the migration tracking
// table in issue #4831.
//
// This release is used as the reported "Supported From" version for the new `v`-prefixed variants
// of resources that were introduced in ASO version 2.16.0 or earlier (the pre-migration resources).
// This ensures our documentation correctly reflects that those `v`-prefixed variants only became
// available when the group was migrated - and shows them in the "Next Release" section of our
// documentation until they've actually shipped.
//
// It also delays the "deprecated" flag being applied to the `v1api`-prefixed variants until
// after the migration has actually released; before then, the `v1api` variants are still the
// only ones users can use.
//
// Keys are group names in lower case; values are ASO version strings (e.g. "v2.21.0").
// When migrating a new group to Hybrid mode, add an entry here recording the upcoming ASO release.
var versionMigrationHybridReleases = map[string]string{
	"alertsmanagement": "v2.19.0",
	"apimanagement":    "v2.19.0",
	"app":              "v2.20.0",
	"appconfiguration": "v2.19.0",
	"authorization":    "v2.21.0",
	"batch":            "v2.18.0",
	"cache":            "v2.21.0",
	"compute":          "v2.20.0",
	"datafactory":      "v2.20.0",
	"dbformysql":       "v2.19.0",
	"eventgrid":        "v2.20.0",
	"sql":              "v2.21.0",
	"storage":          "v2.18.0",
	"subscription":     "v2.21.0",
	"synapse":          "v2.19.0",
	"web":              "v2.19.0",
}

// versionMigrationModes contains a mapping of group names to their version migration modes.
// Keys are group names in lower case.
// We'll start by configuring all existing groups to use the Legacy mode, with no migration required for new groups.
// As we move groups to the new versioning scheme, we'll move them to hybrid mode (giving users a migration path forward).
// Once all groups have been fully migrated, we will delete this file.
var versionMigrationModes = map[string]VersionMigrationMode{
	"alertsmanagement": VersionMigrationModeHybrid,
	"apimanagement":    VersionMigrationModeHybrid,
	"app":              VersionMigrationModeHybrid,
	"appconfiguration": VersionMigrationModeHybrid,

	"authorization": VersionMigrationModeHybrid,

	"batch": VersionMigrationModeHybrid,

	"cache":             VersionMigrationModeHybrid,
	"cdn":               VersionMigrationModeLegacy,
	"cognitiveservices": VersionMigrationModeLegacy,

	"compute": VersionMigrationModeHybrid,

	"containerinstance": VersionMigrationModeLegacy,
	"containerregistry": VersionMigrationModeLegacy,
	"containerservice":  VersionMigrationModeLegacy,

	"datafactory": VersionMigrationModeHybrid,

	"dataprotection": VersionMigrationModeLegacy,
	"dbformariadb":   VersionMigrationModeLegacy,

	"dbformysql": VersionMigrationModeHybrid,

	"dbforpostgresql":         VersionMigrationModeLegacy,
	"devices":                 VersionMigrationModeLegacy,
	"documentdb":              VersionMigrationModeLegacy,
	"eventgrid":               VersionMigrationModeHybrid,
	"eventhub":                VersionMigrationModeLegacy,
	"insights":                VersionMigrationModeLegacy,
	"keyvault":                VersionMigrationModeLegacy,
	"kubernetesconfiguration": VersionMigrationModeLegacy,
	"kusto":                   VersionMigrationModeLegacy,
	"machinelearningservices": VersionMigrationModeLegacy,
	"managedidentity":         VersionMigrationModeLegacy,
	"monitor":                 VersionMigrationModeLegacy,
	"network":                 VersionMigrationModeLegacy,
	"network.frontdoor":       VersionMigrationModeLegacy,
	"notificationhubs":        VersionMigrationModeLegacy,
	"operationalinsights":     VersionMigrationModeLegacy,
	"quota":                   VersionMigrationModeLegacy,
	"redhatopenshift":         VersionMigrationModeLegacy,
	"resources":               VersionMigrationModeLegacy,
	"search":                  VersionMigrationModeLegacy,
	"servicebus":              VersionMigrationModeLegacy,
	"signalrservice":          VersionMigrationModeLegacy,
	"sql":                     VersionMigrationModeHybrid,

	"storage": VersionMigrationModeHybrid,

	"subscription": VersionMigrationModeHybrid,

	"synapse": VersionMigrationModeHybrid,
	"web":     VersionMigrationModeHybrid,
}

// VersionPrefixForGroup returns the version prefix to use for the specified group.
func VersionPrefixForGroup(group string) string {
	if m, ok := versionMigrationModes[strings.ToLower(group)]; ok {
		if m == VersionMigrationModeLegacy {
			return "v1api"
		}
	}

	return "v"
}

// VersionMigrationModeForGroup returns the version migration mode for the specified group.
func VersionMigrationModeForGroup(
	group string,
) VersionMigrationMode {
	if m, ok := versionMigrationModes[strings.ToLower(group)]; ok {
		return m
	}

	return VersionMigrationModeNone
}

func FindGroupsByMode(mode VersionMigrationMode) []string {
	var result []string
	for g, m := range versionMigrationModes {
		if m == mode {
			result = append(result, g)
		}
	}

	return result
}

// HybridMigrationReleaseForGroup returns the ASO release in which the specified group was migrated
// to Hybrid mode versioning, and true if such a release is registered.
// Returns ("", false) for groups without a known migration release (including all non-Hybrid groups).
func HybridMigrationReleaseForGroup(group string) (string, bool) {
	release, ok := versionMigrationHybridReleases[strings.ToLower(group)]
	return release, ok
}
