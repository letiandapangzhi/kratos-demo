# Kratos Demo

kratos实践，基于kratos的[beer-shop](https://github.com/go-kratos/beer-shop)多微服务项目，调整了目录结构

## make工具
```shell
# 根据各服务api的proto文件，更新各服务的pb文件
make api

# 根据各服务internal/conf的proto文件，更新各服务的pb文件
make conf

# 更新各服务的wire依赖注入
make wire

# 更新所有服务的api config wire  
make all

# 接入gorm的gen，更新各服务的数据
# 定义在internal/data/gen/generate.go
make gorm_gen

# 新创建服务，基于.template目录 SERVICE_NAME不定义默认demo
# 新增服务后同步执行了make all
# 配置xxx/configs/config.yaml数据库，xxx/internal/data/gen/generate.go确认服务用到的表->make gorm_gen
make new_mono SERVICE_NAME=xxx

```

## 本地开发
### dbeaver连接
驱动设置：
勾选"Allow Public Key Retrieval"

SSL模式选择"DISABLED"（开发环境）

