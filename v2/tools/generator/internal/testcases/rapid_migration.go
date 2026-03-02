/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcases

import "github.com/Azure/azure-service-operator/v2/internal/set"

// gopterGroups is the set of groups that still use gopter-based property tests.
// To migrate a group to rapid, remove it from this set.
var gopterGroups = set.Make(
	"alertsmanagement",
	"apimanagement",
	"app",
	"appconfiguration",
	"authorization",
	"batch",
	"cache",
	"cdn",
	"cognitiveservices",
	"compute",
	"containerinstance",
	"containerregistry",
	"containerservice",
	"datafactory",
	"dataprotection",
	"dbformariadb",
	"dbformysql",
	"dbforpostgresql",
	"devices",
	"documentdb",
	"entra",
	"eventgrid",
	"eventhub",
	"insights",
	"keyvault",
	"kubernetesconfiguration",
	"kusto",
	"machinelearningservices",
	"managedidentity",
	"monitor",
	"network",
	"network.frontdoor",
	"notificationhubs",
	"operationalinsights",
	"quota",
	"redhatopenshift",
	"resources",
	"search",
	"servicebus",
	"signalrservice",
	"sql",
	"storage",
	"subscription",
	"synapse",
	"web",
)

// UseRapidForGroup returns true if the given group should use rapid-based property tests.
// Groups NOT in the gopterGroups set get rapid tests.
func UseRapidForGroup(group string) bool {
	return !gopterGroups.Contains(group)
}

// UseGopterForGroup returns true if the given group should use gopter-based property tests.
// Groups IN the gopterGroups set keep gopter tests.
func UseGopterForGroup(group string) bool {
	return gopterGroups.Contains(group)
}
