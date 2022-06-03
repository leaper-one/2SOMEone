package util

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func LoadConfig(path string, config interface{}) interface{} {
	// load config
	File, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("读取配置文件失败 #%v", err)
	}
	err = yaml.Unmarshal(File, config)
	if err != nil {
		log.Fatalf("解析失败: %v", err)
	}

	return config
}
