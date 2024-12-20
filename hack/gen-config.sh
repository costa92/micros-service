#!/usr/bin/env bash

# Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/costa92/micros-service


# 示例：gen-config.sh hack/environment.sh configs/miner-apiserver.yaml

env_file="$1"
template_file="$2"

ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/..

source "${ROOT_DIR}/hack/lib/init.sh"

if [ $# -ne 2 ];then
    onex::log::error "Usage: gen-config.sh manifests/env.local configs/onex.service.tmpl"
    exit 1
fi

source "${env_file}"

declare -A envs

set +u
for env in $(sed -n 's/^[^#].*${\(.*\)}.*/\1/p' ${template_file})
do
    if [ -z "$(eval echo \$${env})" ];then
        onex::log::error "environment variable '${env}' not set"
        missing=true
    fi
done

if [ "${missing}" ];then
    onex::log::error 'You may run `source manifests/env.local` to set these environment'
    exit 1
fi

eval "cat << EOF
$(cat ${template_file})
EOF"