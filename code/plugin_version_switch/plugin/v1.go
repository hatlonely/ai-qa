package main

// Version 返回插件的版本号
func Version() string {
	return "v1.0.0"
}

// GetInfo 返回插件的详细信息
func GetInfo() map[string]string {
	return map[string]string{
		"version": "v1.0.0",
		"author":  "Plugin Team",
		"date":    "2023-01-01",
	}
}

func main() {}
