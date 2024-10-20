# Build all by default, even if it's not first
.DEFAULT_GOAL := help

include hack/make-rules/common.mk
include hack/make-rules/all.mk

## --------------------------------------
## Generate / Manifests
## --------------------------------------

##@ Generate

.PHONY: gen


.PHONY: protoc
protoc: ## Generate api proto files.
	$(MAKE) gen.protoc