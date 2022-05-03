#!/usr/bin/env python3

import itertools
import os
import re
import sys

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


# Adds zero-width-non-break spaces to long URL-like strings that aren't links
def addBreaks(path):
    # Read the file into memory
    with open(path, 'r') as f:
        lines = f.readlines()

    # Look for Resource-like URLs and tweak them to allow word breaks
    for index, line in enumerate(lines):
        # be cautious about which lines we modify
        if '/subscriptions/' not in line:
            continue
        if '<a>' in line:
            continue

        lines[index] = line.replace('/', '/&#x200b;')

    # Write the file back to disk
    with open(path, 'w') as f:
        f.truncate()
        f.writelines(lines)


if __name__ == '__main__':
    if len(sys.argv) == 2:
        name = sys.argv[1]
        print(f'Adding non-break-zero-width spaces to {name}')
        addBreaks(name)
