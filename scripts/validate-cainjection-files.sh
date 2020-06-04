#!/usr/bin/env bash

invalid_files=$(grep -L "cert-manager.io" ./config/crd/patches/cainjection*.yaml)

# using a subshell to not modify the IFS variable that determines how the expression "${arr[*]}" is output
# this workaround allows us to reject an array containing empty elements
(IFS='' ; [[ -n "${invalid_files[*]}" ]] ); is_empty=$?

# fail validation if the array is not empty
if [[ $is_empty -eq 0 ]]; then
  echo "Validation check failed for the following files:"
  for file in $invalid_files; do
    echo "  $file"
  done
  echo "Run the following command to fix the files:"
  echo "  sed -i '' 's/certmanager.k8s.io/cert-manager.io/' ./config/crd/patches/cainjection*.yaml"
  exit 1
fi
