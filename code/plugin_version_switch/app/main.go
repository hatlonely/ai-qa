package main

import (
	"bufio"
	"fmt"
	"os"
	"plugin"
	"strings"
	"time"
)

// 插件版本的路径
var pluginPaths = map[string]string{
	"v1": "../plugin/v1.so",
	"v2": "../plugin/v2.so",
	"v3": "../plugin/v3.so",
}

// 加载插件并获取版本信息
func loadPlugin(version string) {
	// 获取插件路径
	path, ok := pluginPaths[version]
	if !ok {
		fmt.Printf("错误: 未知版本 %s\n", version)
		return
	}

	fmt.Printf("正在加载插件版本: %s (路径: %s)\n", version, path)

	// 打开插件
	start := time.Now()
	p, err := plugin.Open(path)
	if err != nil {
		fmt.Printf("加载插件失败: %v\n", err)
		return
	}
	loadTime := time.Since(start)

	// 获取 Version 函数
	versionSymbol, err := p.Lookup("Version")
	if err != nil {
		fmt.Printf("查找 Version 函数失败: %v\n", err)
		return
	}
	versionFunc := versionSymbol.(func() string)

	// 获取 GetInfo 函数
	infoSymbol, err := p.Lookup("GetInfo")
	if err != nil {
		fmt.Printf("查找 GetInfo 函数失败: %v\n", err)
		return
	}
	getInfoFunc := infoSymbol.(func() map[string]string)

	// 调用函数
	ver := versionFunc()
	info := getInfoFunc()

	// 显示结果
	fmt.Printf("\n== 插件信息 ==\n")
	fmt.Printf("版本号: %s\n", ver)
	fmt.Printf("加载耗时: %v\n", loadTime)
	fmt.Printf("详细信息:\n")
	for k, v := range info {
		fmt.Printf("  %s: %s\n", k, v)
	}
	fmt.Println("============\n")
}

func main() {
	fmt.Println("插件版本切换演示")
	fmt.Println("------------------")
	fmt.Println("可用命令:")
	fmt.Println("load v1  - 加载版本1")
	fmt.Println("load v2  - 加载版本2")
	fmt.Println("load v3  - 加载版本3")
	fmt.Println("exit     - 退出程序")
	fmt.Println("------------------")

	// 读取用户输入
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		if strings.HasPrefix(input, "load ") {
			// 提取版本号
			version := strings.TrimSpace(strings.TrimPrefix(input, "load"))
			loadPlugin(version)
		} else {
			fmt.Println("未知命令。请使用 'load v1', 'load v2', 'load v3' 或 'exit'")
		}
	}

	fmt.Println("程序已退出")
}
