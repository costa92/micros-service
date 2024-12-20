# ==============================================================================
# Makefile used to aggregate all makefiles for easy management.
#


include hack/make-rules/tools.mk
include hack/make-rules/golang.mk
include hack/make-rules/generate.mk # 生成代码
include hack/make-rules/image.mk # 构建镜像
include hack/make-rules/copyright.mk # 代码检查
include hack/make-rules/swagger.mk # 生成swagger文档