#
# These variables should not need tweaking.
#

# ==============================================================================
# Includes

# include the common make file
ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR :=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
endif

APIROOT := $(ROOT_DIR)/pkg/api
MANIFESTS_DIR=$(ROOT_DIR)/manifests
SCRIPTS_DIR=$(ROOT_DIR)/hack

# It's necessary to set this because some environments don't link sh -> bash.
# zh 有些环境不会将sh链接到bash，所以需要设置这个变量, 用于指定shell的路径
SHELL := /usr/bin/env bash -o errexit -o pipefail +o nounset
.SHELLFLAGS = -ec

# It's necessary to set the errexit flags for the bash shell.
# zh: 为bash shell设置errexit标志是必要的
export SHELLOPTS := errexit

# ==============================================================================
# Build options
#
PRJ_SRC_PATH :=github.com/costa92/micros-service

# zh: 用于将逗号分隔的字符串转换为空格分隔的字符串
COMMA := ,
SPACE :=
SPACE +=

# zh: 用于指定输出目录和相关子目录
ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(ROOT_DIR)/_output
$(shell mkdir -p $(OUTPUT_DIR))
endif

ifeq ($(origin LOCALBIN),undefined)
LOCALBIN := $(OUTPUT_DIR)/bin
$(shell mkdir -p $(LOCALBIN))
endif

ifeq ($(origin TOOLS_DIR),undefined)
TOOLS_DIR := $(OUTPUT_DIR)/tools
$(shell mkdir -p $(TOOLS_DIR))
endif

ifeq ($(origin TMP_DIR),undefined)
TMP_DIR := $(OUTPUT_DIR)/tmp
$(shell mkdir -p $(TMP_DIR))
endif

# set the version number. you should not need to do this
# for the majority of scenarios.
ifeq ($(origin VERSION), undefined)
# Current version of the project.
  VERSION := $(shell git describe --tags --always --match='v*')
  ifneq (,$(shell git status --porcelain 2>/dev/null))
    VERSION := $(VERSION)-dirty
  endif
endif

# ==============================================================================
# golang

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
GOPATH ?= $(shell go env GOPATH)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# The OS must be linux when building docker images
# PLATFORMS ?= linux_amd64 linux_arm64
# The OS can be linux/windows/darwin when building binaries
PLATFORMS ?= darwin_amd64 windows_amd64 linux_amd64 linux_arm64

# ==============================================================================
# Makefile settings
#
# We don't need make's built-in rules.
# zh: 我们不需要make的内置规则
# 执行命令： V=1 make go.build  可以打印出所有的命令
MAKEFLAGS += --no-builtin-rules
ifeq ($(V),1)
  # 使用 `$(MAKECMDGOALS)` 打印警告消息，显示 Makefile 的目标。`$(MAKECMDGOALS)` 是一个特殊变量，包含在命令行上指定的目标。
  $(warning ***** starting Makefile for goal(s) "$(MAKECMDGOALS)")
  # 打印当前日期和时间的警告消息。
  $(warning ***** $(shell date))
else
  # If we're not debugging the Makefile, don't echo recipes.]
  MAKEFLAGS += -s --no-print-directory
endif

# =====================================================
FIND := find . ! -path './third_party/*' ! -path './vendor/*' ! -path './.git/*' ! -path './.idea/*' ! -path './_output/*'
XARGS := xargs --no-run-if-empty
