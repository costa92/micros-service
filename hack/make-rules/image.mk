# ==============================================================================
# Makefile helper functions for docker image
#


DOCKER := docker

DOCKER_SUPPORTED_API_VERSION ?= 1.32

EXTRA_ARGS ?= --no-cache
_DOCKER_BUILD_EXTRA_ARGS :=


# Track code version with Docker Label.
DOCKER_LABELS ?= git-describe="$(shell date -u +v%Y%m%d)-$(shell git describe --tags --always --dirty)"

ifdef HTTP_PROXY
_DOCKER_BUILD_EXTRA_ARGS += --build-arg HTTP_PROXY=${HTTP_PROXY}
endif


ifneq ($(EXTRA_ARGS), )
_DOCKER_BUILD_EXTRA_ARGS += $(EXTRA_ARGS)
endif

# Determine image files by looking into cmd/*
CMD_DIRS ?= $(wildcard ${ROOT_DIR}/cmd/*)

# Determine images names by stripping out the dir names.
# Filter out directories without Go files, as these directories cannot be compiled to build a docker image.
IMAGES ?= $(filter-out tools, $(foreach dir, $(CMD_DIRS), $(notdir $(if $(wildcard $(dir)/*.go), $(dir),))))
ifeq (${IMAGES},)
  $(error Could not determine IMAGES, set ONEX_ROOT or run in source dir)
endif


.PHONY: image.verify
image.verify: ## Verify docker version.
	$(eval API_VERSION := $(shell $(DOCKER) version | grep -E 'API version: {1,6}[0-9]' | head -n1 | awk '{print $$3} END { if (NR==0) print 0}' ))
	$(eval PASS := $(shell echo "$(API_VERSION) > $(DOCKER_SUPPORTED_API_VERSION)" | bc))
	@if [ $(PASS) -ne 1 ]; then \
		$(DOCKER) -v ;\
		echo "Unsupported docker version. Docker API version should be greater than $(DOCKER_SUPPORTED_API_VERSION)"; \
		exit 1; \
	fi

.PHONY: image.daemon.verify
image.daemon.verify: ## Verify docker daemon version.
	$(eval PASS := $(shell $(DOCKER) version | grep -q -E 'Experimental: {1,5}true' && echo 1 || echo 0))
	@if [ $(PASS) -ne 1 ]; then \
		echo "Experimental features of Docker daemon is not enabled. Please add \"experimental\": true in '/etc/docker/daemon.json' and then restart Docker daemon."; \
		exit 1; \
	fi


.PHONY: image.build
image.build: image.verify go.build.verify $(addprefix image.build., $(addprefix $(IMAGE_PLAT)., $(IMAGES))) ## Build all docker images.


#  Build docker image for each image.
.PHONY: image.dockerfile
image.dockerfile: $(addprefix image.dockerfile., $(IMAGES)) ## Generate all dockerfiles.

#  Generate Dockerfile for each image.
.PHONY: image.dockerfile.%
image.dockerfile.%: ## Generate specified dockerfiles.
	$(eval IMAGE := $(lastword $(subst ., ,$*)))
	# Set a unified environment variable file
	@$(SCRIPTS_DIR)/gen-dockerfile.sh $(GENERATED_DOCKERFILE_DIR) $(IMAGE)
ifeq ($(V),1)
	echo "DBG: Generating Dockerfile at $(GENERATED_DOCKERFILE_DIR)/$(IMAGE)"
endif

.PHONY: image.build.%
ifneq (${MULTISTAGE},1)
image.build.%: go.build.% image.dockerfile.% ## Build specified docker image.
	$(eval IMAGE := $(word 2,$(subst ., ,$*)))
	$(eval IMAGE_PLAT := $(subst _,/,$(PLATFORM)))
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	$(eval DOCKERFILE := Dockerfile)
	$(eval DST_DIR := $(TMP_DIR)/$(IMAGE))
	# In the SemVer versioning specification, the use of the "+" sign is possible, but container image tag names 
	# do not support the "+" character. Therefore, it is necessary to replace the "+" with "-" in the version number. 
	# For example, the version number "v0.18.0+20240121235656" should be transformed into "v0.18.0-20240121235656" for 
	# use as a container tag name.
	$(eval IMAGE_TAG := $(subst +,-,$(VERSION)))
	@echo "===========> Building docker image $(IMAGE) $(IMAGE_TAG) for $(IMAGE_PLAT)"

	@mkdir -p $(TMP_DIR)/$(IMAGE)
	@cp -r $(OUTPUT_DIR)/platforms/$(IMAGE_PLAT)/$(IMAGE) $(TMP_DIR)/$(IMAGE)/

else
# TODO: onex-allinone, does not support multi-stage builds. Need to adjust it.
image.build.%: image.dockerfile.% ## Build specified docker image in multistage way.
endif
	@export OUTPUT_DIR=$(OUTPUT_DIR)
	@if [ -f  $(ROOT_DIR)/build/docker/$(IMAGE)/build.sh ] ; then \
		DST_DIR=$(DST_DIR) OUTPUT_DIR=$(OUTPUT_DIR) IMAGE_PLAT=${IMAGE_PLAT} \
		$(ROOT_DIR)/build/docker/$(IMAGE)/build.sh ; \
	fi
	$(eval BUILD_SUFFIX := $(_DOCKER_BUILD_EXTRA_ARGS) --pull \
		-f $(GENERATED_DOCKERFILE_DIR)/$(IMAGE)/$(DOCKERFILE) \
		--build-arg OS=$(OS) \
		--build-arg ARCH=$(ARCH) \
		--build-arg goproxy=$($(GO) env GOPROXY) \
		--label $(DOCKER_LABELS) \
		-t $(REGISTRY_PREFIX)/$(IMAGE)-$(ARCH):$(IMAGE_TAG) \
		$(DST_DIR))
	@if [ $(shell $(GO) env GOARCH) != $(ARCH) ] ; then \
		$(MAKE) image.daemon.verify ; \
		$(DOCKER) build --platform $(IMAGE_PLAT) $(BUILD_SUFFIX) ; \
	else \
		$(DOCKER) build $(BUILD_SUFFIX) ; \
	fi
	@-rm -rf $(TMP_DIR)/$(IMAGE)



.PHONY: image.push
image.push: image.verify go.build.verify $(addprefix image.push., $(addprefix $(IMAGE_PLAT)., $(IMAGES))) ## Build and push all docker images to docker registry.

#  Build and push specified docker image.
.PHONY: image.push.multiarch
image.push.multiarch: image.verify go.build.verify $(foreach p,$(PLATFORMS),$(addprefix image.push., $(addprefix $(p)., $(IMAGES)))) ## Build and push all docker with supported arch to docker registry.

#  image 推送
.PHONY: image.push.%
image.push.%: image.build.% ## Build and push specified docker image.
	# NOTICE: The `IMAGE_TAG` variable is inherited from the `image.build.%` makefile rule.
	@echo "===========> Pushing image $(IMAGE) $(IMAGE_TAG) to $(REGISTRY_PREFIX)"
	$(DOCKER) push $(REGISTRY_PREFIX)/$(IMAGE)-$(ARCH):$(IMAGE_TAG)