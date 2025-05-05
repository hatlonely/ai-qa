#!/bin/bash

set -e  # 遇到错误立即退出

echo "=== 编译插件 ==="
cd plugin
go build -buildmode=plugin -o plugin.so plugin.go

echo "=== 编译主程序 ==="
cd ../app
go build -o app main.go

echo "=== 运行主程序 ==="
./app

echo "=== 演示完成 ==="
