package configs

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// Constants of configs
const (
	BuildVersion = "BUILD_VERSION"
)

// Option for configurations
type Option struct {
	Name     string `yaml:"name"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Github struct {
		ClientID     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
	} `yaml:"github"`
	System struct {
		Attachments struct {
			Storage string `yaml:"storage"`
			Path    string `yaml:"path"`
		} `yaml:"attachments"`
	} `yaml:"system"`
	Operators   []string `yaml:"operators"`
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
	opt.OperatorSet = make(map[string]bool)
	for _, operator := range opt.Operators {
		opt.OperatorSet[operator] = true
	}
	AppConfig = &opt
	return nil
}
