#!/bin/bash

# Check that generated CRDs are not too large for Kubernetes
# This script validates that CRD files don't exceed the 1.5MiB size limit

set -e

echo "Checking CRD sizes..."
FAILED=false
LIMIT_MB=1.5
# Convert 1.5 MiB to bytes: 1.5 * 1024 * 1024 = 1572864
LIMIT_BYTES=1572864

# Default CRD directory - can be overridden by passing as argument
CRD_DIR="${1:-v2/config/crd/generated/bases}"

if [ ! -d "$CRD_DIR" ]; then
    echo "[ERR] CRD directory not found: $CRD_DIR"
    echo "Usage: $0 [CRD_DIRECTORY]"
    exit 1
fi

echo "Checking CRDs against size limit of ${LIMIT_MB}MiB (${LIMIT_BYTES} bytes)..."

WARNING_BYTES=1048576  # 1 MiB warning threshold
TOTAL_CRDS=0
LARGE_CRDS=0
WARNING_CRDS=0

for crd_file in "$CRD_DIR"/*.yaml; do
  if [ -f "$crd_file" ]; then
    TOTAL_CRDS=$((TOTAL_CRDS + 1))
    # Convert YAML to compact JSON and get size using yq
    json_size=$(yq -o=json --indent 0 "$crd_file" | wc -c)
    
    if [ "$json_size" -gt "$LIMIT_BYTES" ]; then
      yaml_size=$(wc -c < "$crd_file")
      # Use python for floating point calculation
      size_mb=$(python3 -c "print(f'{$json_size / 1024 / 1024:.2f}')")
      echo "[ERR] CRD too large: $(basename "$crd_file")"
      echo "      JSON size: ${json_size} bytes (${size_mb}MiB)"
      echo "      YAML size: ${yaml_size} bytes"
      echo "      Limit: ${LIMIT_BYTES} bytes (${LIMIT_MB}MiB)"
      FAILED=true
      LARGE_CRDS=$((LARGE_CRDS + 1))
    elif [ "$json_size" -gt "$WARNING_BYTES" ]; then
      # Use python for floating point calculation
      size_mb=$(python3 -c "print(f'{$json_size / 1024 / 1024:.2f}')")
      echo "[WRN] Large CRD (>1MiB): $(basename "$crd_file") - ${size_mb}MiB"
      WARNING_CRDS=$((WARNING_CRDS + 1))
    fi
    # Suppress output for CRDs under 1MiB to reduce log noise
  fi
done

echo ""
echo "[INF] CRD Size Summary:"
echo "      Total CRDs: ${TOTAL_CRDS}"
echo "      Large CRDs (>1MiB): ${WARNING_CRDS}"
echo "      Oversized CRDs (>${LIMIT_MB}MiB): ${LARGE_CRDS}"

if [ "$FAILED" = true ]; then
  echo ""
  echo "[ERR] Build failed: One or more CRDs exceed the ${LIMIT_MB}MiB size limit"
  exit 1
else
  echo ""
  echo "[INF] All CRDs are within the size limit"
fi
