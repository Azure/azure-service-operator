#!/usr/bin/env python3

import itertools
import os
import pathlib
import re
import subprocess
import sys
import time


def strip(text, suffix):
    if suffix and text.endswith(suffix):
        return text[:-len(suffix)]
    return text

def grouper(iterable, n, fillvalue=None):
    "Collect data into fixed-length chunks or blocks"
    # grouper('ABCDEFG', 3, 'x') --> ABC DEF Gxx"
    args = [iter(iterable)] * n
    return itertools.zip_longest(*args, fillvalue=fillvalue)

def get_expected_samples(crd_path):
    # For some reason it's 2-5x faster to do this with yq than it is to do it with pyyaml.
    # The shape of this output is as follows:
    #    name: loadbalancer
    #    group: microsoft.network.azure.com
    #    version: v1alpha1api20201101
    # This can be repeated many times if there are multiple APIVersions. Each triplet represents a single GVK.
    result = subprocess.run(['yq', 'e', '{"name": .spec.names.singular, "group": .spec.group, "version": .spec.versions[] | select(.name != "*storage").name}', str(crd_path)], capture_output=True, text=True)

    if result.returncode != 0:
        print(f"yq failed ({result.returncode}): {result.stderr}")
        sys.exit(1)

    expected_samples = set()
    lines = result.stdout.strip().splitlines()

    for kind, group, version in grouper(lines, 3):
        kind = kind.split(' ')[1]
        group = group.split(' ')[1]
        group = strip(group, ".azure.com")
        version = version.split(' ')[1]
        versionPrefix = re.match('v1(alpha|beta)([0-9][a-z]+)?', version).group()
        filename = f'{version}_{kind}.yaml'
        expected_samples.add(os.path.join(group, versionPrefix, filename))

    return expected_samples

if __name__ == "__main__":

    if len(sys.argv) != 2:
        me = sys.argv[0]
        sys.stderr.write(f"usage: {me} <path to config directory> \n")
        sys.stderr.write(f"example: {me} ./v2/config\n")
        sys.exit(1)

    root_dir = pathlib.Path(sys.argv[1])
    config_dir = root_dir / 'config'
    crd_bases_dir = config_dir / 'crd' / 'bases'
    samples_dir = root_dir / 'samples'

    crd_bases = [entry for entry in crd_bases_dir.iterdir() if entry.is_file()]

    expected_samples = set()
    for crd in crd_bases:
        expected_samples.update(get_expected_samples(crd))

    actual_samples = set()

    for entry in samples_dir.glob("**/*"):
        dir_string = str(entry.relative_to(samples_dir))
        if not entry.is_dir() :
            dir_string_slice = dir_string.split(os.sep)
            if len(dir_string_slice) > 3:
                group = dir_string_slice[0]
                versionPrefix = dir_string_slice[len(dir_string_slice) - 2]
                sample = dir_string_slice[len(dir_string_slice) - 1]
                dir_string = os.path.join(group, versionPrefix, sample)
            actual_samples.add(dir_string)

    difference = expected_samples.difference(actual_samples)
    if difference:
        print(f'Found {len(difference)} missing samples: {difference}')
        sys.exit(1)
