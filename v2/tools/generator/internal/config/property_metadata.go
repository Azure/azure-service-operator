/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

// PropertyMetaData contains additional information about a specific property and forms part of a heirarchy
// containing information to supplement the schema and swagger sources consumed by the generator.
//
// ┌──────────────────┐       ┌──────────────────┐       ┌──────────────────┐       ╔══════════════════╗
// │                  │       │                  │       │                  │       ║                  ║
// │  GroupMetaData   │───────│ VersionMetaData  │───────│   KindMetaData   │───────║ PropertyMetaData ║
// │                  │1  1..n│                  │1  1..n│                  │1  1..n║                  ║
// └──────────────────┘       └──────────────────┘       └──────────────────┘       ╚══════════════════╝
//
type PropertyMetaData struct {

}
