
GO := go
# Minimum supported go version.
GO_MINIMUM_VERSION ?= 1.23

ifeq ($(PRJ_SRC_PATH),)
	$(error the variable PRJ_SRC_PATH must be set prior to including golang.mk)
endif

ifeq ($(ROOT_DIR),)
	$(error the variable ROOT_DIR must be set prior to including golang.mk)
endif

# CMD_DIRS is a list of all directories in the cmd directory.
CMD_DIRS := $(wildcard $(ROOT_DIR)/cmd/*)
# Filter out directories without Go files, as these directories cannot be compiled.
COMMANDS := $(filter-out $(wildcard %.md), $(foreach dir, $(CMD_DIRS), $(if $(wildcard $(dir)/*.go), $(dir),)))
BINS ?= $(foreach cmd,${COMMANDS},$(notdir ${cmd}))


ifeq (${COMMANDS},)
  $(error Could not determine COMMANDS, set ROOT_DIR or run in source dir)
endif
ifeq (${BINS},)
  $(error Could not determine BINS, set ROOT_DIR or run in source dir)
endif

.PHONY: go.build.verify
go.build.verify: ## Verify supported go versions.
ifneq ($(shell $(GO) version|awk -v min=$(GO_MINIMUM_VERSION) '{gsub(/go/,"",$$3);if($$3 >= min){print 0}else{print 1}}'), 0)
	$(error unsupported go version. Please install a go version which is greater than or equal to '$(GO_MINIMUM_VERSION)')
endif

.PHONY: go.build.%
go.build.%: ## Build specified applications with platform, os and arch.
	$(eval COMMAND := $(word 2,$(subst ., ,$*)))
	$(eval PLATFORM := $(word 1,$(subst ., ,$*)))
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	#@ONEX_GIT_VERSION=$(VERSION) $(SCRIPTS_DIR)/build.sh $(COMMAND) $(PLATFORM)
	@if grep -q "func main()" $(ROOT_DIR)/cmd/$(COMMAND)/*.go &>/dev/null; then \
		echo "===========> Building binary $(COMMAND) $(VERSION) for $(OS) $(ARCH)" ; \
		CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) $(GO) build $(GO_BUILD_FLAGS) \
		-o $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)/$(COMMAND)$(GO_OUT_EXT) $(PRJ_SRC_PATH)/cmd/$(COMMAND) ; \
	fi

#  go.build: $(addprefix go.build., $(addprefix $(PLATFORM)., $(BINS))) ## Build all applications.
.PHONY: go.build # make go.build PLATFORM=linux BINS=order-server 
go.build: $(addprefix go.build., $(addprefix $(PLATFORM)., $(BINS))) ## make go.build PLATFORM=linux BINS=order-server 

# 编译所有支持的平台
.PHONY: go.build.multiarch
go.build.multiarch: $(foreach p,$(PLATFORMS),$(addprefix go.build., $(addprefix $(p)., $(BINS)))) ## Build all applications with all supported arch.

# 编译当前环境下的CURRENT_PLATFORM 和 CURRENT_OS
.PHONY: go.build.current
go.build.current: $(addprefix go.build., $(addprefix $(CURRENT_PLATFORM)., $(BINS))) ## Build all applications with current platform. eg: make go.build.current BINS=order-server 

# 编译所有支持的平台
.PHONY: go.updates 
go.updates: tools.verify.go-mod-outdated ## Find outdated dependencies.
	@$(GO) list -u -m -json all | go-mod-outdated -update -direct