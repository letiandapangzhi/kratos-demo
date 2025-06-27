.PHONY: init
# 初始化依赖
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: config
# 生成config的pb文件
config:
	find app -type d -depth 1 -not -name "layout" -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) config'

.PHONY: api
# 生成api的pb文件
api:
	find app -type d -depth 1 -not -name "layout" -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) api'

.PHONY: wire
# wire依赖注入生成
wire:
	find app -type d -depth 1 -not -name "layout" -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) wire'

.PHONY: all
# 执行所有操作
all: config api wire

.PHONY: exec_new_mono_sh
# 执行shell脚本
SERVICE_NAME ?= demo
exec_new_mono_sh:
	@export FROM_MAKEFILE=1; \
	sh ./.template/new_mono_repo.sh $(SERVICE_NAME)

.PHONY: new_mono
# 创建新的服务 make new_mono SERVICE_NAME=demo
new_mono: exec_new_mono_sh all

.PHONY: gorm_gen
gorm_gen:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) gorm_gen'