#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail

APIROOT=$1
OUTPUTDIR=$2
PATTERN='^v[0-9]((alpha|beta)[a-z0-9]+)?$'

GOPATH=/Users/$USER/go
mkdir -p $GOPATH/bin
GOBIN="$GOPATH/bin"

#Check if 'gen-crd-api-reference-docs' tool is installed, if not, then install. 
if ! command -v $GOBIN/gen-crd-api-reference-docs  &> /dev/null
then
    echo "Installing gen-crd-api-reference-docs in $GOBIN" 
    #TODO: Replace this with the new release tag.
    go install github.com/ahmetb/gen-crd-api-reference-docs@v0.3.1-0.20220223025230-af7c5e0048a3
fi

#Refresh the directory
rm -rf $OUTPUTDIR
mkdir $OUTPUTDIR

#Iterate through the directories
for package in $(find "$APIROOT" -type d); 
do
    PACKAGE_VERSION=$(basename $package)
    GROUPNAME=$(basename $(dirname $package))

    #Filter the main CRD packages
    if [[ $PACKAGE_VERSION =~ $PATTERN ]] && [[ "$PACKAGE_VERSION" != *"storage" ]] 
    then

        echo "generating docs for: $package"
        "${GOBIN}/gen-crd-api-reference-docs" -config "../docs/api/template/config.json" \
                -template-dir "../docs/api/template" \
                -api-dir $package \
                -out-file "$OUTPUTDIR/$GROUPNAME.$PACKAGE_VERSION.md" \
                "$@"
                
    fi
done

find $OUTPUTDIR -type f -exec sed -i '' '1 s/^/---\n---\n/' {} \;