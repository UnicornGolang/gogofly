package generate

import (
	"fmt"
	"os"
	"path/filepath"
)

// 使用 data 中的 structure.json 文件生成项目使用的目录结构
// 1. 首先使用 map 来解析 json 文件
// 2. 使用 struct 来解析 json 文件

var (
	workDir   string // 当前项目的工作目录
	separator string // 当前环境中使用的路径分隔符
)


func init() {
	separator = string(filepath.Separator)
	workDir, _ = os.Getwd()
	// rootDir = workDir[:strings.LastIndex(workDir, separator)]
	// fmt.Println(rootDir)
  fmt.Println("current project root filepath is : " + workDir)
}
