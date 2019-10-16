package main

import (
	"context"
	"github.com/sathishkumar64/funbook/masterservice/internal/bulkupload"
	"github.com/sathishkumar64/funbook/masterservice/internal/durable"
	"github.com/sathishkumar64/funbook/masterservice/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/sathishkumar64/funbook/masterservice/internal/configs"
	"go.uber.org/zap"
)


func startService(ctx context.Context,db *mongo.Client, logger *zap.Logger,csvFileName string,csvSubFileName string) error {
	database := durable.WrapDatabase(db)
	service.RegisterCategory(ctx,database,csvFileName)
	service.RegisterSubCategory(ctx ,database ,csvSubFileName)
	//service.RegisterService(ctx,database,csvFileName,csvSubFileName)
	return nil
}



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
		options.Config = path.Join(dir, back, "/internal/configs/config.yaml")
	}

	if err := configs.Init(options.Config, options.Environment); err != nil {
		log.Panicln(err)
	}

	logger, err := zap.NewDevelopment()

	config := configs.AppConfig
	ctx := context.Background()

	logger.Info("Testing database...",
		zap.String("CSVFileName....", config.CSVFileName),
		zap.String("CSVSubFileName....", config.CSVSubFileName))

	db := durable.OpenDatabaseClient(ctx, &durable.ConnectionInfo{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
	//	Name:     config.Database.Name,
	})
	defer db.Disconnect(ctx)

	if config.Environment == "production" {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Panicln(err)
	}
	wEXTFileName:= bulkupload.ParseFileName(ctx,config.CSVFileName)
	woEXTSubFileName:= bulkupload.ParseFileName(ctx,config.CSVSubFileName)
	startService(ctx,db, logger,wEXTFileName,woEXTSubFileName)

	logger.Info("Testing config.", zap.String("Testing this....", config.Name))
}
