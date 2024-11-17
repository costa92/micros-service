#!/usr/bin/env bash

# Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/costa92/micros-service





# This file is not intended to be run automatically. It is meant to be run
# immediately before exporting docs. We do not want to check these documents in
# by default.

set -o errexit
set -o nounset
set -o pipefail

ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/..
source "${ROOT_DIR}/hack/lib/init.sh"

golang::setup_env
onex::util::ensure-temp-dir

BINS=(
  gen-docs
  gen-onex-docs
  gen-man
  gen-yaml
)
make build -C "${ROOT_DIR}" BINS="${BINS[*]}"

# Run all known doc generators (today gendocs and genman for nodectl)
# $1 is the directory to put those generated documents
function generate_docs() {
  local dest="$1"


  # Find binary
  gendocs=$(onex::util::find-binary "gen-docs")
  genonexdocs=$(onex::util::find-binary "gen-onex-docs")
  genman=$(onex::util::find-binary "gen-man")
  genyaml=$(onex::util::find-binary "gen-yaml")
  echo genonexdocs $genonexdocs


  mkdir -p "${dest}/docs/guide/en-US/cmd/onexctl"
  "${gendocs}" "${dest}/docs/guide/en-US/cmd/onexctl/"

  mkdir -p "${dest}/docs/guide/en-US/cmd"
  "${genonexdocs}" "${dest}/docs/guide/en-US/cmd/" "onex-fakeserver"


  mkdir -p "${dest}/docs/man/man1/"
  "${genman}" "${dest}/docs/man/man1/" "onex-fakeserver"
  "${genman}" "${dest}/docs/man/man1/" "onex-usercenter"
  "${genman}" "${dest}/docs/man/man1/" "onex-apiserver"
  "${genman}" "${dest}/docs/man/man1/" "onex-gateway"
  "${genman}" "${dest}/docs/man/man1/" "onex-nightwatch"
  "${genman}" "${dest}/docs/man/man1/" "onex-pump"
  "${genman}" "${dest}/docs/man/man1/" "onex-toyblc"
  "${genman}" "${dest}/docs/man/man1/" "onex-controller-manager"
  "${genman}" "${dest}/docs/man/man1/" "onex-minerset-controller"
  "${genman}" "${dest}/docs/man/man1/" "onex-miner-controller"
  "${genman}" "${dest}/docs/man/man1/" "onexctl"

  mkdir -p "${dest}/docs/guide/en-US/yaml/onexctl/"
  "${genyaml}" "${dest}/docs/guide/en-US/yaml/onexctl/"

  # create the list of generated files
  pushd "${dest}" > /dev/null || return 1
  touch docs/.generated_docs
  find . -type f | cut -sd / -f 2- | LC_ALL=C sort > docs/.generated_docs
  popd > /dev/null || return 1
}

# Removes previously generated docs-- we don't want to check them in. $ONEX_ROOT
# must be set.
function remove_generated_docs() {
  if [ -e "${ROOT_DIR}/docs/.generated_docs" ]; then
    # remove all of the old docs; we don't want to check them in.
    while read -r file; do
      rm "${ROOT_DIR}/${file}" 2>/dev/null || true
    done <"${ROOT_DIR}/docs/.generated_docs"
    # The docs/.generated_docs file lists itself, so we don't need to explicitly
    # delete it.
  fi
}

echo "${ONEX_TEMP}"
# generate into ONEX_TMP
generate_docs "${ONEX_TEMP}"

# remove all of the existing docs in ONEX_ROOT
remove_generated_docs

# Copy fresh docs into the repo.
# the shopt is so that we get docs/.generated_docs from the glob.
shopt -s dotglob
cp -af "${ONEX_TEMP}"/* "${ROOT_DIR}"
shopt -u dotglob