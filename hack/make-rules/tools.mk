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

# 安装 goimports 工具的目标规则
# make tools.install.goimports
.PHONY: _install.goimports
_install.goimports: ## 安装 goimports 工具
# 判断 GOIMPORTS_VERSION 是否为空，如果为空则设置为 latest
	$(eval GOIMPORTS_VERSION := $(if $(strip $(GOIMPORTS_VERSION)),$(strip $(GOIMPORTS_VERSION)),latest))
	@$(GO) install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)

# 安装 go-mod-outdated 
# make tools.install.go-mod-outdated
.PHONY: _install.go-mod-outdated
_install.go-mod-outdated: ## 安装 go-mod-outdated 工具
# 判断 GO_MOD_OUTDATED_VERSION 是否为空，如果
	$(eval GO_MOD_OUTDATED_VERSION := $(if $(strip $(GO_MOD_OUTDATED_VERSION)),$(strip $(GO_MOD_OUTDATED_VERSION)),latest))
	@$(GO) install github.com/psampaz/go-mod-outdated@$(GO_MOD_OUTDATED_VERSION)

# 安装 go-swagger 工具的目标规则
.PHONY: _install.swagger
_install.swagger:
	@$(GO) install github.com/go-swagger/go-swagger/cmd/swagger@$(GO_SWAGGER_VERSION)

# 安装 yq 工具的目标规则
.PHONY: _install.yq
_install.yq:
	@$(GO) install github.com/mikefarah/yq/v4@$(YQ_VERSION)