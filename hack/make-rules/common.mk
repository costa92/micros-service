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

# zh: 用于将逗号分隔的字符串转换为空格分隔的字符串
COMMA := ,
SPACE :=
SPACE +=


# ==============================================================================
# golang

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
GOPATH ?= $(shell go env GOPATH)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif
# ==============================================================================


# =====================================================
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
