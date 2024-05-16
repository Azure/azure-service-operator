#!/bin/bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail


print_usage() {
  echo "Usage: delete-role-assignment.sh -d <DIRECTORY>"
}

DIR=

while getopts 'd:g:' flag; do
  case "${flag}" in
    d) DIR="${OPTARG}" ;;
    *) print_usage
       exit 1 ;;
  esac
done


if [[ -z "$DIR" ]]; then
  print_usage
  exit 1
fi

if [ -f "$DIR/azure/roleassignmentid.txt" ]; then
  # Need to delete the role assignment as well so we don't leak them
  ROLE_ASSIGNMENT_ID=$(cat $DIR/azure/roleassignmentid.txt)
  echo "Deleting role assignment: ${ROLE_ASSIGNMENT_ID}"
  az role assignment delete --ids "${ROLE_ASSIGNMENT_ID}"
fi
