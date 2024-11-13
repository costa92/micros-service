#!/usr/bin/env bash


function onex::util::sourced_variable {
  # Call this function to tell shellcheck that a variable is supposed to
  # be used from other calling context. This helps quiet an "unused
  # variable" warning from shellcheck and also document your code.
  true
}

function onex::util::sortable_date() {
  date "+%Y%m%d-%H%M%S"
}