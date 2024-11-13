#!/usr/bin/env bash


set -o errexit
set +o nounset
set -o pipefail

# Short-circuit if init.sh has already been sourced
[[ $(type -t onex::init::loaded) == function ]] && return 0

# Unset CDPATH so that path interpolation can work correctly
# https://github.com/minerrnetes/minerrnetes/issues/52255
unset CDPATH

# Default use go modules
export GO111MODULE=on


# The root of the build/dist directory
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"


ROOT_OUTPUT_SUBPATH="${ROOT_OUTPUT_SUBPATH:-_output}"
ROOT_OUTPUT="${ROOT_DIR}/${ROOT_OUTPUT_SUBPATH}"

source "${ROOT_DIR}/hack/lib/util.sh"
source "${ROOT_DIR}/hack/lib/logging.sh"


onex::log::install_errexit