#!/usr/bin/env bash


set -e

ROOT=$(dirname "${BASH_SOURCE}")/..
function test_make() {
  cd $ROOT && echo PWD: $PWD
  make
  cd ..
}

test_make