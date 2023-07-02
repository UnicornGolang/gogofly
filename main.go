package main

import (
	"gogofly/cmd"
	"gogofly/generate"
)

func premain() {
	// 生成项目的目录结构
	generate.GenerateUseStruct()
}

// @title GoWeb 开发通用框架
// @version v0.0.1
// @description 使用 golang 开发的完整的后端微服务框架
func main() {
	defer cmd.Clean()
	cmd.Start()
}
