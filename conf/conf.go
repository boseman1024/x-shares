package conf

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Conf struct {
	Init     bool
	FilePath string `yaml:"filePath"`
	FileUrl  string `yaml:"fileUrl"`
}

func (conf *Conf) initConf() bool {
	dirPath, err := filepath.Abs("./")
	yamlFile, err := ioutil.ReadFile(dirPath + "\\conf\\conf.yaml")
	if err != nil {
		log.Println("READ CONF FAILED")
		return false
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Println("CONF INIT FAILED")
		return false
	}
	conf.Init = true
	return true
}

func (conf *Conf) GetConf(key string) (string, error) {
	if !conf.Init {
		status := conf.initConf()
		if !status {
			return "", errors.New("GET CONF FAILED")
		}
	}
	if key == "filePath" {
		return conf.FilePath, nil
	} else if key == "fileUrl" {
		return conf.FileUrl, nil
	}
	return "", errors.New("NO KEY ERROR")
}
