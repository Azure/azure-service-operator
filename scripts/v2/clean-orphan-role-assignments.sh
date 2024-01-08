#!/bin/bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

# This script exists to manually clean up the CI subscription role assignments if there get to be too many.
# We could run it automatically but they're leaked relatively rarely and only clutter the UI (not a security issue as the
# backing identity is deleted).
# We should probably figure out what's creating the assignments and update that to do a best effort delete...
# It's probably make-mi-fic.sh

set -o errexit
set -o nounset
set -o pipefail

DELETE=0

print_usage() {
  echo "Usage: cleanup-orphan-role-assignments.sh [-y]"
  echo "    -y: Actually delete everything. If not specified, the default is dry-run mode"
}

while getopts 'y' flag; do
  case "${flag}" in
    y) DELETE=1 ;;
    *) print_usage
       exit 1 ;;
  esac
done

# This lists all of the role assignments with no principal
IDS=$(az role assignment list --query "([?(principalName == '')].id).join(' ', [*])" | tr -d '"')

COUNT=$(echo ${IDS} | wc -w)

echo "About to delete ${COUNT} role assignments..."

if [[ "$DELETE" -eq 1 ]]; then
  echo "Starting to delete role assignments..."
  for I in $IDS
    do
      echo "Deleting ID ${I}"
      az role assignment delete --ids "${I}" -y
    done
  echo "Done deleting role assignments"
else
  echo "Not deleting role assignments in dry-run mode. Pass -y to actually delete the role assignments"
fi
