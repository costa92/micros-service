#!/usr/bin/env bash


# Ensure the go tool exists and is a viable version.
# Inputs:
#   env-var GO_VERSION is the desired go version to use, downloading it if needed (defaults to content of .go-version)
#   env-var FORCE_HOST_GO set to a non-empty value uses the go version in the $PATH and skips ensuring $GO_VERSION is used
golang::verify_go_version() {
  # default GO_VERSION to content of .go-version
  GO_VERSION="${GO_VERSION:-"$(cat "${ROOT_DIR}/.go-version")"}"
  if [ "${GOTOOLCHAIN:-auto}" != 'auto' ]; then
    # no-op, just respect GOTOOLCHAIN
    :
  elif [ -n "${FORCE_HOST_GO:-}" ]; then
    # ensure existing host version is used, like before GOTOOLCHAIN existed
    export GOTOOLCHAIN='local'
  else
    # otherwise, we want to ensure the go version matches GO_VERSION
    GOTOOLCHAIN="go${GO_VERSION}"
    export GOTOOLCHAIN
    # if go is either not installed or too old to respect GOTOOLCHAIN then use gimme
    if ! (command -v go >/dev/null && [ "$(go version | cut -d' ' -f3)" = "${GOTOOLCHAIN}" ]); then
      export GIMME_ENV_PREFIX=${GIMME_ENV_PREFIX:-"${LOCAL_OUTPUT_ROOT}/.gimme/envs"}
      export GIMME_VERSION_PREFIX=${GIMME_VERSION_PREFIX:-"${LOCAL_OUTPUT_ROOT}/.gimme/versions"}
      # eval because the output of this is shell to set PATH etc.
      eval "$("${ONEX_ROOT}/third_party/gimme/gimme" "${GO_VERSION}")"
    fi
  fi

  if [[ -z "$(command -v go)" ]]; then
    onex::log::usage_from_stdin <<EOF
Can't find 'go' in PATH, please fix and retry.
See http://golang.org/doc/install for installation instructions.
EOF
    return 2
  fi

  local go_version
  IFS=" " read -ra go_version <<< "$(go version)"
  local minimum_go_version
  minimum_go_version=go1.20
  if [[ "${minimum_go_version}" != $(echo -e "${minimum_go_version}\n${go_version[2]}" | sort -s -t. -k 1,1 -k 2,2n -k 3,3n | head -n1) && "${go_version[2]}" != "devel" ]]; then
    onex::log::usage_from_stdin <<EOF
Detected go version: ${go_version[*]}.
OneX requires ${minimum_go_version} or greater.
Please install ${minimum_go_version} or later.
EOF
    return 2
  fi
}




# onex::golang::setup_env will check that the `go` commands is available in
# ${PATH}. It will also check that the Go version is good enough for the
# Node build.
#
# Outputs:
#   env-var GOBIN is unset (we want binaries in a predictable place)
#   env-var GO15VENDOREXPERIMENT=1
#   env-var GO111MODULE=on
#   current directory is within GOPATH
golang::setup_env() {
  golang::verify_go_version

  # Set GOROOT so binaries that parse code can work properly.
  GOROOT=$(go env GOROOT)
  export GOROOT

  # Unset GOBIN in case it already exists in the current session.
  unset GOBIN

  # This seems to matter to some tools
  export GO15VENDOREXPERIMENT=1

  # Open go module feature
  export GO111MODULE=on

  # This is for sanity.  Without it, user umasks leak through into release
  # artifacts.
  umask 0022
}
