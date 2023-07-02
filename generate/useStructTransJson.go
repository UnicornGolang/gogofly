package generate

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

var rootNode Node

func GenerateUseStruct() {
	useStructTransJson()
	paeseNode(rootNode, workDir)
}

type Node struct {
	Text     string `json:"text"`
	Children []Node `json:"children"`
}

// 使用 struct 来转换 JSON 对象
func useStructTransJson() {
	jsonBytes, _ := os.ReadFile(workDir + "/data/structure.json")
	err := json.Unmarshal(jsonBytes, &rootNode)
	if err != nil {
		panic("Load Json File error: " + err.Error())
	}
}

// 解析读取出来的结构体
func paeseNode(node Node, parentDir string) {
	if node.Text != "" {
		makedir(node, parentDir)
		parentDir = parentDir + separator + node.Text
	}
	for _, v := range node.Children {
		paeseNode(v, parentDir)
	}
}

// 创建
func makedir(node Node, parentDir string) {
	basedir := parentDir + separator + node.Text
  fmt.Println(":>", basedir)
  err := os.MkdirAll(basedir, fs.ModePerm)
  if err != nil {
    panic("create dir error: "+ err.Error())
  }
}
