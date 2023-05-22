#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

# Kustomize throws an error if the resources are not in the same root directory.
# Hence, we need this script to put samples in one place for operator-bundle.

set -o errexit
set -o nounset
set -o pipefail

SAMPLES_DIR=$1
OUT_DIR=$2

rm -rf "$OUT_DIR"
mkdir "$OUT_DIR"
echo "resources: " > "$OUT_DIR/kustomization.yaml"

for filepath in $(find "$SAMPLES_DIR" -type f); do

  FILENAME=$(basename $filepath)
  if
  [[ $filepath == *"refs"* ]] ||
  [[ $filepath != *"v1api"* ]] ||
  grep -q "$FILENAME" "$OUT_DIR/kustomization.yaml";
  then
    continue
  fi

  cp $filepath "$OUT_DIR"
  echo " - $FILENAME" >> "$OUT_DIR/kustomization.yaml"

done
