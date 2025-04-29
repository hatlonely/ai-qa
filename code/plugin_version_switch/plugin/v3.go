package main

// Version 返回插件的版本号
func Version() string {
	return "v3.0.0"
}

// GetInfo 返回插件的详细信息
func GetInfo() map[string]string {
	return map[string]string{
		"version": "v3.0.0",
		"author":  "Plugin Team",
		"date":    "2023-12-01",
		"new":     "重大更新和性能优化",
		"status":  "stable",
	}
}

func main() {}
