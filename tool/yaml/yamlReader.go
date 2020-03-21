package yamlReader

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"path/filepath"
)

const (
	YAMLDEFAULTPATH = "./tool/yaml/config.yaml"
)

type YamlConfigure struct {
	Port string `yaml:"port"`
	Version string `yaml:"version"`
}

type YamlReader struct {
	Configure YamlConfigure
}

func Instance() (*YamlReader, error) {
	var filePath, err = filepath.Abs(YAMLDEFAULTPATH)
	if err != nil {
		return nil, err
	}

	log.Printf("正在获取yaml文件路径[%v] \n", filePath)
	yr := new(YamlReader)
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yaml文件[%v]获取错误：%v\n", filePath, err.Error())
		return nil, err
	}

	tagConfigure := new(YamlConfigure)

	err = yaml.Unmarshal(yamlFile, tagConfigure)

	if err != nil {
		log.Printf("读取数据错误[%v]", err)
		return nil, err
	}

	yr.Configure = *tagConfigure

	return yr, nil
}