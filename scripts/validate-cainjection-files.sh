#!/usr/bin/env bash

# cainjection*
# cert-manager.io/inject-ca-from: $(CERTIFICATE_NAMESPACE)/$(CERTIFICATE_NAME)

# go to project root directory
cd $(dirname "${BASH_SOURCE[0]}")/..

old_certmgr_str="certmanager.k8s.io/inject-ca-from"
old_certmgr_files=$(grep -Hino "$old_certmgr_str" ./config/crd/patches/cainjection*.yaml)
if [ ${#old_certmgr_files[@]} -gt 0 ]; then
  echo "Validation check failed for the following files:"
  for file in $old_certmgr_files; do
    echo "  $file"
  done
  echo "Run the following command to fix the files:"
  echo "  sed -i '' 's/certmanager.k8s.io/certmanager.io/' ./config/crd/patches/cainjection*.yaml"
  exit 1
fi
