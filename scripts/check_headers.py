#!/usr/bin/env python3

import itertools
import os
import re
import sys


number_of_lines = 8


# Turn colors in this script off by setting the NO_COLOR variable in your
# environment to any value:
#
# $ NO_COLOR=1 python3 check_headers.py
if os.getenv('NO_COLOR') is None:
    header = '\033[1;33m'
    reset = '\033[0m'
else:
    header = ''
    reset = ''


def print_colorful(text):
    print(f'{header}{text}{reset}')


# Returns true if the file passes validation, false if it fails
def is_file_compliant(regex, path):
    if not path.endswith('.go'):
        return True

    with open(path, 'r') as f:
        combined = ''.join(itertools.islice(f, number_of_lines))

        if not regex.search(combined):
            return False
    
    return True


if __name__ == '__main__':
    msftRegex = 'Copyright \(c\) Microsoft Corporation\.(\s)*(.){0,7}Licensed under the MIT License\.'

    # We have a few files that were copied from CAPI and modified some. We need to maintain attribution and leave
    # their license as is. See <> for an issue discussing if they could be moved into a place we could depend on.
    capiRegex = 'Copyright \d{4} The Kubernetes Authors\.(\s)*(.){0,7}Licensed under the Apache License, Version 2\.0'
    regex = re.compile('({}|{})'.format(msftRegex, capiRegex), re.MULTILINE | re.IGNORECASE)
    failed_files = []

    print_colorful('==> Checking copyright headers <==')

    for root, subdirs, files in os.walk('.', topdown=True):
        files = [f for f in files if not f[0] == '.'] # exclude hidden files
        subdirs[:] = [d for d in subdirs if not d[0] == '.'] # exclude hidden dirs, e.g. .cache

        for filename in files:
            path = os.path.join(root, filename)
            if not is_file_compliant(regex, path):
                failed_files.append(path)

    if failed_files:
        print('The following files are missing the required copyright header:')
        for name in failed_files:
            print(f'\t{name}')

    print_colorful('Done checking copyright headers')

    if failed_files:
        sys.exit(1)
