#!/usr/bin/env python3

import re
import subprocess
import sys

def to_semver(git_version: str):
    """Converts git version to semver.

    `git describe` will build tags which we then convert:
    >>> to_semver("v2.0.0") # exact tag match
    'v2.0.0'
    >>> to_semver("v2.0.0-dirty") # exact tag match, uncommitted changes
    'v2.0.0-dirty'
    >>> to_semver("v2.0.0-1-g63121f04") # 1 commit since tag match
    'v2.0.0-rev.1'
    >>> to_semver("v2.0.0-1-g63121f04-dirty") # 1 commit since tag match, uncommitted-changes
    'v2.0.0-rev.1-dirty'

    Prerelease versions:
    >>> to_semver("v2.0.0-alpha.0") # exact tag match with prerelease
    'v2.0.0-alpha.0'
    >>> to_semver("v2.0.0-alpha.0-pre.0") # exact tag match with multiple prerelease bits
    'v2.0.0-alpha.0-pre.0'
    >>> to_semver("v2.0.0-alpha.0-pre.0-1-g63121f04") # with multiple prerelease bits and commit
    'v2.0.0-alpha.0-pre.0-rev.1'
    >>> to_semver("v2.0.0-alpha.1-dirty") # exact tag match, uncommitted changes
    'v2.0.0-alpha.1-dirty'
    >>> to_semver("v2.0.0-beta-1-g63121f04") # 1 commit since tag match
    'v2.0.0-beta-rev.1'
    >>> to_semver("v2.0.0-rc.0-1-g63121f04-dirty") # 1 commit since tag match, uncommitted changes
    'v2.0.0-rc.0-rev.1-dirty'
    """

    match = re.match(r"^(?P<version>[^-]+)(?P<prerelease>-(?!\d).+?)?(-(?P<rev>\d+)-(?P<tag>[^-]+))?(?P<dirty>-dirty)?$", git_version)
    gs = match.groupdict()
    version = gs['version']

    if gs['prerelease'] is not None:
        version += gs['prerelease']

    if gs['rev'] is not None:
        # note that we don't add '+gitcommit' as that is invalid in a Docker tag despite being semver-appropriate
        version += f'-rev.{gs["rev"]}' 

    if gs['dirty'] is not None:
        version += "-dirty"

    return version

if __name__ == "__main__":

    if len(sys.argv) != 2:
        me = sys.argv[0]
        sys.stderr.write(f"usage: {me} <tag-prefix>\n")
        sys.stderr.write(f"example: {me} v2\n")
        import doctest
        print(f"Running tests: {doctest.testmod()}")
        sys.exit(1)

    version_prefix = sys.argv[1]

    git_describe = subprocess.run(["git", "describe", "--tags", "--dirty", "--match", f"{version_prefix}*"], capture_output=True, text=True)

    if git_describe.returncode != 0:
        print(f"git describe failed ({git_describe.returncode}): {git_describe.stderr}")
        sys.exit(1)

    git_version = git_describe.stdout.strip()
    print(to_semver(git_version))
