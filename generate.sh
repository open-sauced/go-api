#!/usr/bin/env bash

# This script will generate the OpenAPI Go client code for the OpenSauced API
# - Requires Docker to run

# This release is identical to v1.52.0, it just contains the fast forwarded
# swagger.yaml document
API_VERSION="v1.52.1-beta.1"

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.0.0 generate \
    --additional-properties packageName=client \
    --generator-name go \
    --git-repo-id go-api \
    --git-user-id open-sauced \
    --input-spec https://raw.githubusercontent.com/open-sauced/api/${API_VERSION}/swagger.yaml \
    --output /local/client \
    --skip-validate-spec

pushd client
    go fmt client/...
popd
