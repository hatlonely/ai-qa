# Golang Plugin 基础示例

这是一个展示 Golang plugin 机制基础用法的简单示例，包含一个计算器插件和一个加载该插件的主程序。

## 目录结构

```
plugin_basics/
├── plugin/          # 插件目录
│   └── plugin.go    # 插件源代码
├── app/             # 主程序目录
│   └── main.go      # 主程序源代码
├── build.sh         # 编译运行脚本
└── README.md        # 本文件
```

## 功能说明

1. 插件提供以下功能：
   - `Add(a, b int) int` - 加法函数
   - `Subtract(a, b int) int` - 减法函数
   - `Multiply(a, b int) int` - 乘法函数
   - `Version` 变量 - 版本信息

2. 主程序加载插件并调用这些函数

## 运行方法

确保您使用的操作系统支持 Go plugin（Linux、macOS、FreeBSD），Windows 不支持。

```bash
# 运行示例
./build.sh
```

## 注意事项

1. 插件和主程序必须使用相同版本的 Go 编译
2. 插件和主程序使用的任何共享类型必须来自相同的包
3. 编译的插件(.so文件)与编译它的Go版本绑定
4. 当主程序和插件依赖相同的第三方库时，必须确保版本完全一致
