package main

// Add 导出的加法函数
func Add(a, b int) int {
	return a + b
}

// Subtract 导出的减法函数
func Subtract(a, b int) int {
	return a - b
}

// 这个变量和下面的函数都会被导出
var Version = "1.0.0"

func Multiply(a, b int) int {
	return a * b
}

// 插件必须有main函数，但可以为空
func main() {}
