#!/bin/bash

# Copyright The Shipwright Contributors
# 
# SPDX-License-Identifier: Apache-2.0

#
# Installs Tekton Pipelines operator.
#

set -eu

TEKTON_VERSION="${TEKTON_VERSION:-v0.18.1}"

TEKTON_HOST="github.com"
TEKTON_HOST_PATH="tektoncd/pipeline/releases/download"

echo "# Deploying Tekton Pipelines Operator '${TEKTON_VERSION}'"

kubectl apply \
    --filename="https://${TEKTON_HOST}/${TEKTON_HOST_PATH}/${TEKTON_VERSION}/release.yaml" \
    --output="yaml"
