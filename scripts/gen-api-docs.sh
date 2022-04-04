#!/usr/bin/env bash

# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.

set -o errexit
set -o nounset
set -o pipefail

APIROOT=$1
OUTPUTDIR=$2
TEMPLATEDIR=$3
PATTERN='^v[0-9]((alpha|beta)[a-z0-9]+)?$'

rm -rf $OUTPUTDIR
mkdir $OUTPUTDIR

# Iterate through the directories
for package in $(find "$APIROOT" -type d); 
do
    PACKAGE_VERSION=$(basename $package)
    GROUPNAME=$(basename $(dirname $package))

    # Filter the main CRD packages matching the pattern and ignore the storage packages
    if [[ $PACKAGE_VERSION =~ $PATTERN ]] && [[ "$PACKAGE_VERSION" != *"storage" ]] 
    then

        echo "generating docs for: $package"
        "gen-crd-api-reference-docs" -config "$TEMPLATEDIR/config.json" \
                -template-dir "$TEMPLATEDIR" \
                -api-dir $package \
                -out-file "$OUTPUTDIR/$GROUPNAME.$PACKAGE_VERSION.md" \
                "$@"
    fi
done

# Hacky way to get through the "plain text html not allowed" hugo error
find $OUTPUTDIR -type f -exec sed -i '1 s/^/---\n---\n/' {} \;
