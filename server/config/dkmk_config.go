package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Server struct {
	Listen string `yaml:"listen"`
}

type Log struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
	Name  string `yaml:"name"`
}

type Database struct {
	DSN     string `yaml:"dsn"`
	MaxIdle int    `yaml:"max_idle"`
	MaxConn int    `yaml:"max_conn"`
}

type Web struct {
	BaseURL    string `yaml:"base_url"`
	App        string `yaml:"app"`
	PublicPath string `yaml:"public_path"`
}

type DKMKConfig struct {
	Log    Log      `yaml:"log"`
	DB     Database `yaml:"db"`
	WEB    Web      `yaml:"web"`
	Server Server   `yaml:"server"`
}

func LoadConfig(filename string, cfg interface{}) error {
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else if err = yaml.Unmarshal(yamlFile, cfg); err != nil {
		return err
	} else {
		return nil
	}
}

var Cfg DKMKConfig

func LoadDKMKConfig(filename string) *DKMKConfig {
	if err := LoadConfig(filename, &Cfg); err != nil {
		logrus.WithError(err).Errorf("config load failed. : %+v", err)
		panic(err)
	}
	return &Cfg
}
