#!/usr/bin/env bash

CURRENT_DIR="."

# jpmcb: The openapi client generates an api_contribution_service.go file incorrectly
# which only contains service endpoints for "repos" and should generally be discarded as duplicate
rm client/api_contribution_service.go

# Apply all the patch files
for patch in "${CURRENT_DIR}"/patches/*; do
    echo "Applying patch: $patch"
    git apply "$patch"
    if [[ $? -ne 0 ]]; then
        echo "Error applying patch: $patch"
        exit 1
    fi
done
