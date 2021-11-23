#!/bin/bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail

PREFIX=''
KEEP_AGE=0

print_usage() {
  echo "Usage: delete-old-resourcegroups.sh -p <RESOURCE_PREFIX> [-a <KEEP_AGE>]"
  echo "    -p (required): Delete all resource groups with this prefix"
  echo "    -a: Only delete the resource group if it is older than this many seconds. Default is 0 which means delete all matching resource groups."
}

while getopts 'p:a:' flag; do
  case "${flag}" in
    p) PREFIX="${OPTARG}" ;;
    a) KEEP_AGE=${OPTARG} ;;
    *) print_usage
       exit 1 ;;
  esac
done


if [ ! $PREFIX ]; then
  echo "-p flag is required"
  print_usage
  exit 1
fi

# Verify that the KEEP_AGE is an int
# Bash is so clear...
case "$KEEP_AGE" in
    ("" | *[!0-9]*)
        echo 'error (-a not a positive decimal integer number)' >&2
        exit 1
esac

if [ "$KEEP_AGE" -eq 0 ]; then
  RESOURCE_GROUPS=`az group list --query "[*].[name]" -o table | { grep "^${PREFIX}" || true; }`
else
  # [*]: this must match what is specified in the CreateTestResourceGroupDefaultTags function
  RESOURCE_GROUPS=`az group list --query "[*].{Name: name, CreatedAt: tags.CreatedAt}" \
    | jq -r ".[] | select(.Name | test(\"^${PREFIX}\")) | select(.CreatedAt == null or now-(.CreatedAt | fromdate) > ${KEEP_AGE}) | .Name"`
fi

for rgname in ${RESOURCE_GROUPS[@]}; do
  echo "$rgname will be deleted"; \
  az group delete --name $rgname --no-wait --yes; \
done
