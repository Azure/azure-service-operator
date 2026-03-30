#!/bin/bash
set -e

# Patterns of files to ignore when checking for code changes.
# These are git pathspec patterns passed to "git diff" with :(exclude,glob).
# With glob magic, "*" matches within a single directory (does not cross "/"),
# while "**" matches across directory boundaries.
# A trailing "/" matches the directory recursively.
IGNORE_FILTERS=(
  "docs/"
  "README.md"
  "ROADMAP.md"
  "SECURITY.md"
  "SUPPORT.md"
  "v2/README.md"
  "scripts/*"  # Only ignore direct children of scripts/, not subdirectories like scripts/v2/
  "scripts/v2/check-changes.sh"
  ".agents/"
  ".github/instructions/"
  ".github/copilot-instructions.md"
)

# Build pathspec exclusions from IGNORE_FILTERS
EXCLUSIONS=()
for FILTER in "${IGNORE_FILTERS[@]}"; do
  EXCLUSIONS+=(":(exclude,glob)${FILTER}")
done

ALL_CHANGED_FILES=$(git diff HEAD origin/main --name-only)
CHANGED_FILES=$(git diff HEAD origin/main --name-only -- . "${EXCLUSIONS[@]}")

ALL_COUNT=0
for FILE in $ALL_CHANGED_FILES; do
  ALL_COUNT=$((ALL_COUNT+1))
done

NON_IGNORED_COUNT=0
echo "Checking for file changes..."
for FILE in $CHANGED_FILES; do
  echo "Source code file ${FILE} changed"
  NON_IGNORED_COUNT=$((NON_IGNORED_COUNT+1))
done

IGNORED_COUNT=$((ALL_COUNT - NON_IGNORED_COUNT))

echo "" # Blank line for readability
echo "$IGNORED_COUNT match(es) for ignore filter '${IGNORE_FILTERS[*]}' found."
echo "$NON_IGNORED_COUNT match(es) for changed source code files found."
if [[ $NON_IGNORED_COUNT -gt 0 ]]; then
  echo "code-changed=true" >> $GITHUB_OUTPUT
else
  echo "code-changed=false" >> $GITHUB_OUTPUT
fi
