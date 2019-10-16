package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Constants of configs
const (
	BuildVersion = "BUILD_VERSION"
)
//Database is used to defining database related information .
type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

// Option for configurations
type Option struct {
	Name     string `yaml:"name"`
	Database  `yaml:"database"`
	CSVFileName string `yaml:"csvFileName"`
	CSVSubFileName string `yaml:"csvSubFileName"`
	Environment string
	OperatorSet map[string]bool
}

// AppConfig is the configs for the whole application
var AppConfig *Option

// Init is using to initialize the configs
func Init(file, env string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	var options map[string]Option
	err = yaml.Unmarshal(data, &options)
	if err != nil {
		return err
	}
	opt := options[env]
	opt.Environment = env
	/*opt.OperatorSet = make(map[string]bool)
	for _, operator := range opt.Operators {
		log.Println("operator..............",operator)
		opt.OperatorSet[operator] = true
	}*/
	AppConfig = &opt
	return nil
}
