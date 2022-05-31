package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
	}
}

func loadConfig(path string) *Config {
	// load config
	var config Config
	File, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("读取配置文件失败 #%v", err)
	}
	err = yaml.Unmarshal(File, &config)
	if err != nil {
		log.Fatalf("解析失败: %v", err)
	}

	return &config
}
