# Build all by default, even if it's not first
.DEFAULT_GOAL := help

include hack/make-rules/common.mk
include hack/make-rules/all.mk

# ==============================================================================
# Usage
define USAGE_OPTIONS

\033[35mOptions:\033[0m

  BINS             The binaries to build. Default is all of cmd.

endef
export USAGE_OPTIONS
# ==============================================================================

## --------------------------------------
## Generate / Manifests
## --------------------------------------

##@ Generate
.PHONY: protoc
protoc: ## Generate api proto files.
	$(MAKE) gen.protoc


##@ Build
.PHONY: build
build: tidy ## Build the operator.
	$(MAKE) go.build


.PHONY: build.multiarch 
build.multiarch: ## Build source code for multiple platforms. See option PLATFORMS. make build.multiarch PLATFORMS=linux/amd64,linux/arm64  BINS=order-server
	$(MAKE) go.build.multiarch 


##@ Image
.PHONY: build.image 
build.image: ## Build docker image. 
	$(MAKE) image.build


.PHONY: image.multiarch
image.multiarch: ## Build docker images for multiple platforms. See option PLATFORMS.
	$(MAKE) image.build.multiarch


.PHONY: push
push: ## Build docker images for host arch and push images to registry.
	$(MAKE) image.push



.PHONY: tidy
tidy:
	@$(GO) mod tidy


## add-copyright
.PHONY: add-copyright
add-copyright: ## Ensures source code files have copyright license headers. 
	$(MAKE) copyright.add


.PHONY: swagger
#swagger: gen.protoc
swagger: ## Generate and aggregate swagger document.
	@$(MAKE) swagger.run

.PHONY: swagger.serve
serve-swagger: ## Serve swagger spec and docs at 65534.
	@$(MAKE) swagger.serve


.PHONY: targets
targets: Makefile ## Show all Sub-makefile targets.
	@for mk in `echo $(MAKEFILE_LIST) | sed 's/Makefile //g'`; do echo -e \\n\\033[35m$$mk\\033[0m; awk -F':.*##' '/^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 }' $$mk;done;


# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk command is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php
.PHONY: help
help: Makefile ## Display this help info.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<TARGETS> <OPTIONS>\033[0m\n\n\033[35mTargets:\033[0m\n"} /^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' Makefile #$(MAKEFILE_LIST)
	@echo -e "$$USAGE_OPTIONS"
