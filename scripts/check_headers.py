import itertools
import os
import re
import sys


number_of_lines = 4


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
    regex = re.compile('(copyright|generated)', re.MULTILINE | re.IGNORECASE)
    failed_files = []

    print_colorful('==> Checking copyright headers <==')

    for root, subdirs, files in os.walk('.'):
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
