/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

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

	// VersionMigrationModeHybrid indicates that new style versions will be generated for all resources, with old style
	//
	//
	// that both old and new style versions will be
	// generated for resource versions introduced up to ASO version 2.16.0, and only new style
	// versions for resource versions introduced in ASO version 2.17.0 and later.
	//
	// To achieve this, new style versions will be generated for all packages,
	// and a new pipeline stage will create older style versions for resources existing
	// in ASO version 2.16.0 and earlier to provide a migration path forward for existing users.
	VersionMigrationModeHybrid VersionMigrationMode = "hybrid"
)

// versionMigrationModes contains a mapping of group names to their version migration modes.
// We'll start by configuring all existing groups to use the Legacy mode, with no migration required for new groups.
// As we move groups to the new versioning scheme, we'll move them to hybrid mode (giving users a migration path forward).
// Once all groups have been fully migrated, we will delete this file.
var versionMigrationModes = map[string]VersionMigrationMode{
	"alertsmanagement":        VersionMigrationModeLegacy,
	"apimanagement":           VersionMigrationModeLegacy,
	"app":                     VersionMigrationModeLegacy,
	"appconfiguration":        VersionMigrationModeLegacy,
	"authorization":           VersionMigrationModeLegacy,
	"batch":                   VersionMigrationModeLegacy,
	"cache":                   VersionMigrationModeLegacy,
	"cdn":                     VersionMigrationModeLegacy,
	"cognitiveservices":       VersionMigrationModeLegacy,
	"compute":                 VersionMigrationModeLegacy,
	"containerinstance":       VersionMigrationModeLegacy,
	"containerregistry":       VersionMigrationModeLegacy,
	"containerservice":        VersionMigrationModeLegacy,
	"datafactory":             VersionMigrationModeLegacy,
	"dataprotection":          VersionMigrationModeLegacy,
	"dbformariadb":            VersionMigrationModeLegacy,
	"dbformysql":              VersionMigrationModeLegacy,
	"dbforpostgresql":         VersionMigrationModeLegacy,
	"devices":                 VersionMigrationModeLegacy,
	"documentdb":              VersionMigrationModeLegacy,
	"entra":                   VersionMigrationModeLegacy,
	"eventgrid":               VersionMigrationModeLegacy,
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
	"resources.m":             VersionMigrationModeLegacy,
	"scheme.g":                VersionMigrationModeLegacy,
	"search":                  VersionMigrationModeLegacy,
	"servicebus":              VersionMigrationModeLegacy,
	"signalrservice":          VersionMigrationModeLegacy,
	"sql":                     VersionMigrationModeLegacy,
	"storage":                 VersionMigrationModeLegacy,
	"subscription":            VersionMigrationModeLegacy,
	"synapse":                 VersionMigrationModeLegacy,
	"web":                     VersionMigrationModeLegacy,
}
