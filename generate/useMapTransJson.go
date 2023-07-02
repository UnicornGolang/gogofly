package generate

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

var (
	jsonData  map[string]any
)

func GenerateUseMap() {
	loadJsonUseMap()
	parseMap(jsonData, workDir)
}

// 加载项目结构的 json 文件
func loadJsonUseMap() {
	jsonBytes, _ := os.ReadFile(workDir + "/data/structure.json")
	err := json.Unmarshal(jsonBytes, &jsonData)
	if err != nil {
		panic("Load Json Data Error: " + err.Error())
	}
}

// 解析 json 文件
func parseMap(jsonData map[string]any, parentDir string) {
	for k, v := range jsonData {
		switch v := v.(type) {
		case string:
			path := v
			if len(path) < 1 {
				continue
			}
			if parentDir != "" {
				path = parentDir + separator + path
				if k == "text" {
					parentDir = path
				}
			} else {
				parentDir = path
			}
			creatDir(path)
		case []any:
			parseArray(v, parentDir)
		}
	}
}

func parseArray(jsonData []any, parentDir string) {
	for _, v := range jsonData {
		mapV, _ := v.(map[string]any)
		parseMap(mapV, parentDir)
	}
}

func creatDir(path string) {
	if len(path) < 1 {
		return
	}
	fmt.Println(":>" + path)
	// 创建文件夹
	err := os.MkdirAll(path, fs.ModePerm)
	if err != nil {
		panic("Create dir Error: " + err.Error())
	}
}
