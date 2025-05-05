#!/bin/bash

set -e  # 遇到错误立即退出

echo "=== 编译多版本插件 ==="
cd plugin

echo "编译 v1 插件..."
go build -buildmode=plugin -o v1.so v1.go

echo "编译 v2 插件..."
go build -buildmode=plugin -o v2.so v2.go

echo "编译 v3 插件..."
go build -buildmode=plugin -o v3.so v3.go

echo "=== 编译主程序 ==="
cd ../app
go build -o app main.go

echo "=== 编译完成 ==="
echo "运行以下命令启动程序:"
echo "cd app && ./app"
