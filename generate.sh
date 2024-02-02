#!/usr/bin/env bash

# This script will generate the OpenAPI Go client code for the OpenSauced API
# - Requires Docker to run

API_VERSION="beta"

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v7.0.0 generate \
    --additional-properties packageName=client \
    --generator-name go \
    --git-repo-id go-api \
    --git-user-id open-sauced \
    --input-spec https://raw.githubusercontent.com/open-sauced/api/${API_VERSION}/swagger.yaml \
    --output /local/client \
    --skip-validate-spec

pushd client
    go fmt ./...
popd
