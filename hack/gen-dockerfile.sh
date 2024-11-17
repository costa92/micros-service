#!/usr/bin/env bash

# Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/costa92/micros-service



ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/..
source "${ROOT_DIR}/hack/lib/init.sh"


#  判断参数个数是否为2
if [ $# -ne 2 ];then
    onex::log::error "Usage: gen-dockerfile.sh ${DOCKERFILE_DIR} ${IMAGE_NAME}"
    exit 1
fi


DOCKERFILE_DIR=$1/$2
IMAGE_NAME=$2

#  环境变量
ALL_IN_ONE_IMAGE_NAME=onex-allinone
ENV_FILE=${ENV_FILE:-${ROOT_DIR}/manifests/env.local}
source ${ENV_FILE}


declare -A envs


function cat_dockerfile()
{
	cat << 'EOF'
# syntax=docker/dockerfile:1.4


# Dockerfile generated by hack/gen-dockerfile.sh. DO NOT EDIT.

# Build the IMAGE_NAME binary
# Run this with docker build --build-arg prod_image=<golang:x.y.z>
# Default <prod_image> is BASE_IMAGE
ARG prod_image=BASE_IMAGE

FROM ${prod_image}
LABEL maintainer="<costa9293@gmail.com>"

WORKDIR /opt/project

# Note: the <prod_image> is required to support
# setting timezone otherwise the build will fail
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone

COPY IMAGE_NAME /opt/project/bin/

ENTRYPOINT ["/opt/project/bin/IMAGE_NAME"]
EOF
}


function cat_multistage_dockerfile()
{
	cat << 'EOF'
# syntax=docker/dockerfile:1.4


# Dockerfile generated by scripts/gen-dockerfile.sh. DO NOT EDIT.

# Build the IMAGE_NAME binary

# Run this with docker build --build-arg prod_image=<golang:x.y.z>
# Default <prod_image> is BASE_IMAGE
ARG prod_image=BASE_IMAGE

# Ignore Hadolint rule "Always tag the version of an image explicitly."
# It's an invalid finding since the image is explicitly set in the Makefile.
# https://github.com/hadolint/hadolint/wiki/DL3006
# hadolint ignore=DL3006
FROM golang:1.20 as builder
WORKDIR /workspace

# Run this with docker build --build-arg goproxy=$(go env GOPROXY) to override the goproxy
ARG goproxy=https://proxy.golang.org
ARG OS
ARG ARCH

# Run this with docker build.
ENV GOPROXY=$goproxy

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN --mount=type=cache,target=/go/pkg/mod \
  go mod download

# Copy the sources
COPY api/ api/
COPY cmd/IMAGE_NAME cmd/IMAGE_NAME
COPY pkg/ pkg/
COPY internal/ internal/
COPY third_party/ third_party/

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
RUN CGO_ENABLED=0 GOOS=${OS:-linux} GOARCH=${ARCH} go build -a ./cmd/IMAGE_NAME

# Production image
FROM ${prod_image}
LABEL maintainer="<costa9293@gmail.com>"

WORKDIR /opt/project

# Note: the <prod_image> is required to support
# setting timezone otherwise the build will fail
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo "Asia/Shanghai" > /etc/timezone

COPY --from=builder /workspace/IMAGE_NAME /opt/project/bin/
# Use uid of nonroot user (65532) because kubernetes expects numeric user when applying pod security policies
USER 65532
ENTRYPOINT ["/opt/project/bin/IMAGE_NAME"]
EOF
}



function cat_allinone_dockerfile()
{
	cat << EOF
# syntax=docker/dockerfile:1.4

# Dockerfile generated by hack/gen-dockerfile.sh. DO NOT EDIT.

# Build the onex-allinone binary
# Run this with docker build --build-arg prod_image=<golang:x.y.z>
# Default <prod_image> is systemd-debian:12
# jrei/systemd-debian:12 因为镜像内置的 Linux 工具太少，不方便排障，所以不适用
# 使用 centos 镜像，也可以使你有途径了解另一个知名的 Linux 发行版：CentOS
ARG prod_image=centos:centos8

FROM \${prod_image}
LABEL maintainer="<costa9293@gmail.com>"

WORKDIR ${INSTALL_DIR}

# Note: the <prod_image> is required to support
# setting timezone otherwise the build will fail
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone

# 用环境变量替换
COPY bin/* ${BIN_DIR}/
COPY appconfig/* ${CONFIG_DIR}/
COPY cert ${CONFIG_DIR}/cert
COPY config ${CONFIG_DIR}/
COPY systemd/* /etc/systemd/system/
COPY onex-admin.sh ${INSTALL_DIR}/

# 设置开机启动
# RUN systemctl enable onex-usercenter
# RUN systemctl enable onex-apiserver
# RUN systemctl enable onex-gateway
# RUN systemctl enable onex-nightwatch
# RUN systemctl enable onex-pump
# RUN systemctl enable onex-toyblc
# RUN systemctl enable onex-controller-manager
# RUN systemctl enable onex-minerset-controller
# RUN systemctl enable onex-miner-controller
# RUN systemctl enable onex-cacheserver

# 将 OneX 二进制文件加入到 Linux 的搜索路径中
ENV PATH=/opt/project/bin:\${PATH}

ENTRYPOINT ["/usr/sbin/init"]
EOF
}

function get_base_image() {
  declare -A map=(
    ["onex-fake-miner"]="debian:trixie"
    ["onexctl"]="debian:trixie"
    [${ALL_IN_ONE_IMAGE_NAME}]="systemd-debian:12"
  )

  base_image=${map[$1]}
  echo ${base_image:-debian:trixie}
}

cat_func=cat_dockerfile
[[ ! -d ${DOCKERFILE_DIR} ]] && mkdir -p ${DOCKERFILE_DIR}

[[ ${IMAGE_NAME} == "${ALL_IN_ONE_IMAGE_NAME}" ]] && cat_func=cat_allinone_dockerfile

BASE_IMAGE=$(get_base_image ${IMAGE_NAME})

# generate dockerfile
eval ${cat_func}| sed -e "s/BASE_IMAGE/${BASE_IMAGE}/g" -e "s/IMAGE_NAME/${IMAGE_NAME}/g" > ${DOCKERFILE_DIR}/Dockerfile

# generate multi-stage dockerfile
# onex-allinone does not need multiple stages.
if [[ ${IMAGE_NAME} != "${ALL_IN_ONE_IMAGE_NAME}" ]];then
    cat_multistage_dockerfile | \
        sed -e "s/BASE_IMAGE/${BASE_IMAGE}/g" -e "s/IMAGE_NAME/${IMAGE_NAME}/g" > ${DOCKERFILE_DIR}/Dockerfile.multistage
fi