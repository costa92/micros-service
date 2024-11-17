#!/usr/bin/env bash

# Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/costa92/micros-service

set -e
# Install the protoc-gen-go-errors-code binary in $GOPATH/bin
go build -o $GOPATH/bin/protoc-gen-go-errors-code