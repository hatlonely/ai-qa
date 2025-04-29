package main

// Version 返回插件的版本号
func Version() string {
	return "v2.0.0"
}

// GetInfo 返回插件的详细信息
func GetInfo() map[string]string {
	return map[string]string{
		"version": "v2.0.0",
		"author":  "Plugin Team",
		"date":    "2023-06-01",
		"new":     "添加了新功能",
	}
}

func main() {}
