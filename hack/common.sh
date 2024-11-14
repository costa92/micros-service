#!/usr/bin/env bash


# shellcheck disable=SC2034 # Variables sourced in other scripts.

# Common utilities, variables and checks for all build scripts.
set -eEuo pipefail

# Unset CDPATH, having it set messes up with script import paths
unset CDPATH

USER_ID=$(id -u)
GROUP_ID=$(id -g)
KUBE_VERBOSE=${PROJECT_VERBOSE:-1}

# This will canonicalize the path
ROOT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")"/.. && pwd -P)

source "${ROOT_DIR}/hack/lib/init.sh"


# Constants
readonly BUILD_IMAGE_REPO=miner-build


DOCKER_OPTS=${DOCKER_OPTS:-""}
IFS=" " read -r -a DOCKER <<< "docker ${DOCKER_OPTS}"




# The variable SERVER_SIDE_COMPONENTS is used to define onex server-side components.
# These components need to installed as a service.
declare -Ax SERVER_SIDE_COMPONENTS=(
  ["pay"]="pay-server"
  ["order"]="order-server"
)


# The variable ONEX_CLIENT_SIDE_COMPONENTSis used to define onex client-side components.
# These components no need to installed as a service, but used as a command line.
declare -Ax CLIENT_SIDE_COMPONENTS=(
  ["ctl"]="ctl"
)



# The variable ALL_COMPONENTS is used to define all onex components.
# 12 useable components (@2024.01.01)
declare -Ax ALL_COMPONENTS
for key in "${!CLIENT_SIDE_COMPONENTS[@]}"; do
  ALL_COMPONENTS["$key"]="${CLIENT_SIDE_COMPONENTS[$key]}"
done
for key in "${!SERVER_SIDE_COMPONENTS[@]}"; do
  ALL_COMPONENTS["$key"]="${SERVER_SIDE_COMPONENTS[$key]}"
done





