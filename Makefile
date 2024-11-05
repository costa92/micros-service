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


.PHONY: targets
targets: Makefile ## Show all Sub-makefile targets.
	@for mk in `echo $(MAKEFILE_LIST) | sed 's/Makefile //g'`; do echo -e \\n\\033[35m$$mk\\033[0m; awk -F':.*##' '/^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 }' $$mk;done;
