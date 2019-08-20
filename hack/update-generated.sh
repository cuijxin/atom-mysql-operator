#!/bin/bash -e

set -o errexit
set -o nounset
set -o pipefail

CODEGEN_PKG=${GOPATH}/src/k8s.io/code-generator

GROUPS_VERSION="mysql:v1alpha1"

${CODEGEN_PKG}/generate-groups.sh \
    all \
    atom-mysql-operator/pkg/generated \
    atom-mysql-operator/pkg/apis \
    "${GROUPS_VERSION}" \
    "$@"

