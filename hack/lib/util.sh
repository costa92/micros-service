#!/usr/bin/env bash

# Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/costa92/micros-service



function onex::util::sourced_variable {
  # Call this function to tell shellcheck that a variable is supposed to
  # be used from other calling context. This helps quiet an "unused
  # variable" warning from shellcheck and also document your code.
  true
}

function onex::util::sortable_date() {
  date "+%Y%m%d-%H%M%S"
}



# Example:  onex::util::trap_add 'echo "in trap DEBUG"' DEBUG
# See: http://stackoverflow.com/questions/3338030/multiple-bash-traps-for-the-same-signal
function onex::util::trap_add() {
  local trap_add_cmd
  trap_add_cmd=$1
  shift

  for trap_add_name in "$@"; do
    local existing_cmd
    local new_cmd

    # Grab the currently defined trap commands for this trap
    existing_cmd=$(trap -p "${trap_add_name}" |  awk -F"'" '{print $2}')

    if [[ -z "${existing_cmd}" ]]; then
      new_cmd="${trap_add_cmd}"
    else
      new_cmd="${trap_add_cmd};${existing_cmd}"
    fi

    # Assign the test. Disable the shellcheck warning telling that trap
    # commands should be single quoted to avoid evaluating them at this
    # point instead evaluating them at run time. The logic of adding new
    # commands to a single trap requires them to be evaluated right away.
    # shellcheck disable=SC2064
    trap "${new_cmd}" "${trap_add_name}"
  done
}


# Opposite of onex::util::ensure-temp-dir()
function onex::util::cleanup-temp-dir() {
  rm -rf "${ONEX_TEMP}"
}

# Create a temp dir that'll be deleted at the end of this bash session.
#
# Vars set:
#   ONEX_TEMP
function onex::util::ensure-temp-dir() {
  if [[ -z ${ONEX_TEMP-} ]]; then
    ONEX_TEMP=$(mktemp -d 2>/dev/null || mktemp -d -t onex.XXXXXX)
    onex::util::trap_add onex::util::cleanup-temp-dir EXIT
  fi
}

function onex::util::host_os() {
  local host_os
  case "$(uname -s)" in
    Darwin)
      host_os=darwin
      ;;
    Linux)
      host_os=linux
      ;;
    *)
      onex::log::error "Unsupported host OS.  Must be Linux or Mac OS X."
      exit 1
      ;;
  esac
  echo "${host_os}"
}

function onex::util::host_arch() {
  local host_arch
  case "$(uname -m)" in
    x86_64*)
      host_arch=amd64
      ;;
    i?86_64*)
      host_arch=amd64
      ;;
    amd64*)
      host_arch=amd64
      ;;
    aarch64*)
      host_arch=arm64
      ;;
    arm64*)
      host_arch=arm64
      ;;
    arm*)
      host_arch=arm
      ;;
    i?86*)
      host_arch=x86
      ;;
    s390x*)
      host_arch=s390x
      ;;
    ppc64le*)
      host_arch=ppc64le
      ;;
    *)
      onex::log::error "Unsupported host arch. Must be x86_64, 386, arm, arm64, s390x or ppc64le."
      exit 1
      ;;
  esac
  echo "${host_arch}"
}

# This figures out the host platform without relying on golang.  We need this as
# we don't want a golang install to be a prerequisite to building yet we need
# this info to figure out where the final binaries are placed.
function onex::util::host_platform() {
  echo "$(onex::util::host_os)/$(onex::util::host_arch)"
}

# looks for $1 in well-known output locations for the platform ($2)
# $ONEX_ROOT must be set
function onex::util::find-binary-for-platform() {
  local -r lookfor="$1"
  local -r platform="$2"
  local locations=(
    "${ROOT_DIR}/_output/bin/${lookfor}"
    "${ROOT_DIR}/_output/dockerized/bin/${platform}/${lookfor}"
    "${ROOT_DIR}/_output/local/bin/${platform}/${lookfor}"
    "${ROOT_DIR}/platforms/${platform}/${lookfor}"
  )

  # if we're looking for the host platform, add local non-platform-qualified search paths
  if [[ "${platform}" = "$(onex::util::host_platform)" ]]; then
    locations+=(
      "${ROOT_DIR}/_output/local/go/bin/${lookfor}"
      "${ROOT_DIR}/_output/dockerized/go/bin/${lookfor}"
    );
  fi

  # looks for $1 in the $PATH
  if which "${lookfor}" >/dev/null; then
    local -r local_bin="$(which "${lookfor}")"
    locations+=( "${local_bin}"  );
  fi

  # List most recently-updated location.
  local -r bin=$( (ls -t "${locations[@]}" 2>/dev/null || true) | head -1 )

  if [[ -z "${bin}" ]]; then
    onex::log::error "Failed to find binary ${lookfor} for platform ${platform}"
    return 1
  fi

  echo -n "${bin}"
}

# looks for $1 in well-known output locations for the host platform
# $ONEX_ROOT must be set
function onex::util::find-binary() {
  onex::util::find-binary-for-platform "$1" "$(onex::util::host_platform)"
}
