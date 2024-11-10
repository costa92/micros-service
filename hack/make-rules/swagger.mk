# ==============================================================================
# Makefile helper functions for swagger
#

.PHONY: swagger.run
swagger.run: tools.verify.swagger
	@echo "===========> Generating swagger API docs"
	#@swagger generate spec --scan-models -w $(ROOT_DIR)/cmd/gen-swagger-type-docs -o $(ROOT_DIR)/api/swagger/kubernetes.yaml
	@swagger mixin `find $(ROOT_DIR)/api/openapi -name "*.swagger.json"` \
		-q                                                    \
		--keep-spec-order                                     \
		--format=yaml                                         \
		--ignore-conflicts                                    \
		-o $(ROOT_DIR)/api/swagger/swagger.yaml
	@echo "Generated at: $(ROOT_DIR)/api/swagger/swagger.yaml"

.PHONY: swagger.serve
swagger.serve: tools.verify.swagger
	@swagger serve -F=redoc --no-open --port 65534 $(ROOT_DIR)/api/swagger/swagger.yaml
