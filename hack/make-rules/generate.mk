# ==============================================================================
# Makefile helper functions for generate necessary files
#

# Generate go source files from protobuf files.
.PHONY: gen.protoc
gen.protoc: ## Generate go source files from protobuf files.
	@protoc \
		--proto_path=$(APIROOT) \
		--proto_path=$(APISROOT) \
		--proto_path=$(ROOT_DIR)/third_party/protobuf \
		--go_out=paths=source_relative:$(APIROOT) \
		--go-http_out=paths=source_relative:$(APIROOT) \
		--go-grpc_out=paths=source_relative:$(APIROOT) \
		--go-errors_out=paths=source_relative:$(APIROOT) \
		--go-errors-code_out=paths=source_relative:$(ROOT_DIR)/docs/guide/zh-CN/api/errors-code \
		--validate_out=paths=source_relative,lang=go:$(APIROOT) \
		--openapi_out=fq_schema_naming=true,default_response=false:$(ROOT_DIR)/api/openapi \
		--openapiv2_out=$(ROOT_DIR)/api/openapi \
		--openapiv2_opt=logtostderr=true \
		--openapiv2_opt=json_names_for_fields=false \
		$(shell find $(APIROOT) -name *.proto)


##@ gen.system
.PHONY: gen.systemd
gen.systemd: $(addprefix gen.systemd., $(SERVICES)) ## Generate all systemd unit files.

.PHONY: gen.systemd.%
gen.systemd.%: ## Generate specified systemd unit file.
	$(eval ENV_FILE ?= $(MANIFESTS_DIR)/env.local)
	$(eval GENERATED_SERVICE_DIR := $(OUTPUT_DIR)/systemd)
	$(eval SERVICE := $(lastword $(subst ., ,$*)))
	@echo "===========> Generating $(SERVICE) systemd unit file"
	@$(SCRIPTS_DIR)/gen-service-config.sh $(SERVICE) $(ENV_FILE) \
		$(ROOT_DIR)/configs/systemd.tmpl.service $(GENERATED_SERVICE_DIR)
ifeq ($(V),1)
	echo "DBG: Generating systemd unit file at $(GENERATED_SERVICE_DIR)/$(SERVICE).service"
endif


##@ gen appconfig
.PHONY: gen.appconfig
gen.appconfig: $(addprefix gen.appconfig.,$(SERVICES)) ## Generate all application configuration files.

.PHONY: gen.appconfig.%
gen.appconfig.%: ## Generate specified application configuration file.
	$(eval ENV_FILE ?= $(MANIFESTS_DIR)/env.local)
	$(eval GENERATED_SERVICE_DIR := $(OUTPUT_DIR)/appconfig)
	$(eval SERVICE := $(lastword $(subst ., ,$*)))
	@echo "===========> Generating $(SERVICE) configuration file"
	@$(SCRIPTS_DIR)/gen-service-config.sh $(SERVICE) $(ENV_FILE) \
		$(ROOT_DIR)/configs/appconfig/$(SERVICE).config.tmpl.yaml $(GENERATED_SERVICE_DIR)
ifeq ($(V),1)
	@echo "DBG: Generating $(SERVICE) application configuration file at $(GENERATED_SERVICE_DIR)/$(SERVICE)"
endif