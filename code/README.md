# Golang Plugin 示例集合

这个目录包含多个展示 Golang plugin 机制的示例项目。

## 目录结构

```
code/
├── plugin_basics/           # 基础插件示例
│   ├── calculator/          # 计算器插件
│   ├── main/                # 基础主程序
│   ├── run.sh               # 运行脚本
│   └── README.md            # 说明文档
│
├── plugin_version_switch/   # 插件版本切换示例
│   ├── version_plugin/      # 多版本插件
│   ├── main_app/            # 版本切换主程序
│   ├── build.sh             # 编译脚本
│   └── README.md            # 说明文档
│
└── README.md                # 本文档
```

## 示例说明

### 1. 基础插件示例 (plugin_basics)

展示 Go plugin 机制的基本用法，包含一个简单的计算器插件和一个调用该插件的主程序。

功能包括：
- 导出和调用插件函数
- 访问插件中的变量
- 基本的插件加载过程

详细信息请查看 [plugin_basics/README.md](plugin_basics/README.md)

### 2. 插件版本切换示例 (plugin_version_switch)

展示如何在程序运行期间动态切换不同版本的插件，包含三个版本的插件和一个能够在运行时加载不同版本的主程序。

功能包括：
- 多版本插件编译与管理
- 运行时动态切换插件版本
- 插件加载性能测量

详细信息请查看 [plugin_version_switch/README.md](plugin_version_switch/README.md)

## 运行环境要求

所有示例都要求：
1. Go 1.8 或更高版本
2. Linux、macOS 或 FreeBSD 操作系统（Windows 不支持 Go plugin）
3. 插件和主程序必须使用相同版本的 Go 编译

## 注意事项

1. Go plugin 机制有一些限制，详情请参考各示例的 README 文件
2. 这些示例主要用于演示和学习目的，实际应用中可能需要更复杂的错误处理和优化
