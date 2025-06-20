#!/bin/bash
if [ -z "$FROM_MAKEFILE" ]; then
  echo -e "Error: 此脚本仅允许通过 Makefile 执行\n请使用 make new_mono <服务名称>" >&2
  exit 1
fi

# 检查参数数量
if [ $# -ne 1 ]; then
    echo "<服务名称>缺失"
    exit 1
fi

service_name=$1
echo "创建服务 $service_name ..."

current_dir=$(pwd)

echo "复制模板并修改名称..."
cp -r "$current_dir/.template/app/layout" "$current_dir/app/$service_name"
# 修改文件夹名称 用于wire
mv "$current_dir/app/$service_name/cmd/layout" "$current_dir/app/$service_name/cmd/$service_name"
# 替换api的proto路径
find "$current_dir/app/$service_name/api/v1" -name "*.proto" | xargs sed -i '' "s/layout/$service_name/g"
# 替换conf的proto路径
find "$current_dir/app/$service_name/internal/conf" -name "*.proto" | xargs sed -i '' "s/layout/$service_name/g"
# 遍历所有.go文件（排除wire_gen.go和*pb.go）
find "$current_dir/app/$service_name" -name "*.go" ! -name "wire_gen.go" ! -name "*pb.go" -type f | while read file; do
    # echo "处理文件: $file"
    sed -i "" "s/layout/$service_name/g" "$file"
done

echo "创建服务 $service_name 完成"