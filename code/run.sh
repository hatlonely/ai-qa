#!/bin/bash

set -e  # 遇到错误立即退出

echo "=== 编译插件 ==="
cd calculator
go build -buildmode=plugin -o calculator.so calculator.go

echo "=== 编译主程序 ==="
cd ../main
go build -o main main.go

echo "=== 运行主程序 ==="
./main

echo "=== 演示完成 ==="
