.PHONY: tools.install.%
tools.install.%: ## Install a specified tool.
	@echo "===========> Installing $*"
	@$(MAKE) _install.$*

.PHONY: tools.verify.%
tools.verify.%: ## Verify a specified tool.
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi


# make tools.install.addlicense
.PHONY: _install.addlicense
_install.addlicense: ## Install addlicense.
# 判断 ADDLICENSE_VERSION 是否为空 如果为空则设置为latest
	$(eval ADDLICENSE_VERSION := $(if $(strip $(ADDLICENSE_VERSION)),$(strip $(ADDLICENSE_VERSION)),latest))
	@$(GO) install github.com/superproj/addlicense@$(ADDLICENSE_VERSION)