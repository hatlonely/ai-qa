package main

import (
	"fmt"
	"plugin"
)

func main() {
	// 加载插件
	p, err := plugin.Open("../calculator/calculator.so")
	if err != nil {
		fmt.Printf("无法加载插件: %v\n", err)
		return
	}

	// 查找 Add 函数
	addSymbol, err := p.Lookup("Add")
	if err != nil {
		fmt.Printf("无法找到 Add 函数: %v\n", err)
		return
	}

	// 类型断言
	add, ok := addSymbol.(func(int, int) int)
	if !ok {
		fmt.Println("类型转换失败")
		return
	}

	// 调用插件函数
	result := add(10, 5)
	fmt.Printf("10 + 5 = %d\n", result)

	// 查找并调用 Subtract 函数
	subtractSymbol, err := p.Lookup("Subtract")
	if err != nil {
		fmt.Printf("无法找到 Subtract 函数: %v\n", err)
		return
	}
	subtract := subtractSymbol.(func(int, int) int)
	fmt.Printf("10 - 5 = %d\n", subtract(10, 5))

	// 查找并获取 Version 变量
	versionSymbol, err := p.Lookup("Version")
	if err != nil {
		fmt.Printf("无法找到 Version 变量: %v\n", err)
		return
	}
	version := *versionSymbol.(*string)
	fmt.Printf("插件版本: %s\n", version)

	// 查找并调用 Multiply 函数
	multiplySymbol, err := p.Lookup("Multiply")
	if err != nil {
		fmt.Printf("无法找到 Multiply 函数: %v\n", err)
		return
	}
	multiply := multiplySymbol.(func(int, int) int)
	fmt.Printf("10 * 5 = %d\n", multiply(10, 5))
}
