package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/sathishkumar64/funbook/masterservice/internal/configs"
	"go.uber.org/zap"
)

func main() {

	var options struct {
		Config      string `short:"c" long:"config" description:"Where's the config file place, default /masterservice/internal/configs/config.yaml"`
		Environment string `short:"e" long:"environment" default:"development"`
	}
	p := flags.NewParser(&options, flags.Default)

	if _, err := p.Parse(); err != nil {
		log.Panicln(err)
	}

	if options.Config == "" {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Panicln(err)
		}
		back := ".."
		if strings.Contains(dir, "cmd") {
			back = "../.."
		}
		options.Config = path.Join(dir, back, "/masterservice/internal/configs/config.yaml")
	}

	if err := configs.Init(options.Config, options.Environment); err != nil {
		log.Panicln(err)
	}

	logger, err := zap.NewDevelopment()

	config := configs.AppConfig

	if config.Environment == "production" {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Panicln(err)
	}

	logger.Info("Testing config.", zap.String("Testing this....", config.Name))
}
