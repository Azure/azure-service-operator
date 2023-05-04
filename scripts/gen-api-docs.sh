#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT=$(realpath $0)
SCRIPTDIR=$(dirname "$SCRIPT")
APIROOT=$1
OUTPUTDIR=$2
TEMPLATEDIR=$3

PATTERN='^v[0-9]((alpha|beta)[a-z0-9]+)?$'

# Create the output folder if it's missing
mkdir "$OUTPUTDIR" --parents

# Delete all markdown except _index.md files
find "$OUTPUTDIR" -type f -name '[a-z]*.md'  -delete

# Iterate through the directories
for package in $(find "$APIROOT" -type d); 
do
    PACKAGE_VERSION=$(basename "$package")
    GROUPNAME=$(basename $(dirname $package))

    # Filter the main CRD packages matching the pattern and ignore the storage packages
    if [[ $PACKAGE_VERSION =~ $PATTERN ]] && [[ "$PACKAGE_VERSION" != *"storage" ]] 
    then

        echo "generating docs for: $package"
        mkdir "$OUTPUTDIR/$GROUPNAME" --parents
        "gen-crd-api-reference-docs" -config "$TEMPLATEDIR/config.json" \
                -template-dir "$TEMPLATEDIR" \
                -api-dir $package \
                -out-file "$OUTPUTDIR/$GROUPNAME/$PACKAGE_VERSION.md" \
                "$@"
        $SCRIPTDIR/wrap-resource-urls.py "$OUTPUTDIR/$GROUPNAME/$PACKAGE_VERSION.md"

    fi
done
