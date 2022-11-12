package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

var (
	ConfigInformation *ConfigInfo
)

type ConfigInfo struct {
	DBConfigInfo *DBModel `yaml:"db"`
	MQConfigInfo *MQModel `yaml:"mq"`
	Spec         string   `yaml:"spec"`
	Port         string   `yaml:"port"`
}

type DBModel struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

type MQModel struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	QueueName string `yaml:"queue_name"`
}

func LoadConfigInformation() {
	wr, _ := os.Getwd()
	filePath := path.Join(wr, "audit_log_config.yml")
	configData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf(" config file read failed: %s", err)
		os.Exit(-1)
	}
	err = yaml.Unmarshal(configData, &ConfigInformation)
	if err != nil {
		fmt.Printf(" config parse failed: %s", err)
		os.Exit(-1)
	}
}
