# 安装指定工具的目标规则
.PHONY: tools.install.%
tools.install.%: ## 安装指定工具
	@echo "===========> 正在安装 $*"
	@$(MAKE) _install.$*

# 验证指定工具的目标规则
.PHONY: tools.verify.%
tools.verify.%: ## 验证指定工具
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

# 安装 addlicense 工具的目标规则
# make tools.install.addlicense
.PHONY: _install.addlicense
_install.addlicense: ## 安装 addlicense 工具
# 判断 ADDLICENSE_VERSION 是否为空，如果为空则设置为 latest
	$(eval ADDLICENSE_VERSION := $(if $(strip $(ADDLICENSE_VERSION)),$(strip $(ADDLICENSE_VERSION)),latest))
	@$(GO) install github.com/superproj/addlicense@$(ADDLICENSE_VERSION)