package main

import (
	"github.com/behrouz-rfa/mongo-specification/pkg/database/factory"
	"github.com/behrouz-rfa/mongo-specification/pkg/infrastructure/database"
	"project-structure/internal/config"
	"project-structure/internal/module/user/application/service"
	"project-structure/internal/module/user/presentation/api"
	"project-structure/pkg/common"
	"project-structure/pkg/logger"

	factoryInternal "project-structure/internal/factory"
)

func main() {
	//db config

	// initialize logger
	logger.Init()
	lg := logger.General.Component("main")
	cfg := config.LoadConfig()
	dbController := DbController(cfg.Mongo.ToPgConfig())

	repoFactory := factoryInternal.NewRepoFactory()

	userService := service.NewUser(repoFactory, dbController)

	// create a new API server
	apiServer := api.NewServer(&api.ServerConfig{
		Port:    cfg.Port,
		GinMode: cfg.GinMode,
		Services: &common.Services{
			User: userService,
		},
	})

	// start the API server
	if err := apiServer.Start(); err != nil {
		lg.WithError(err).Fatal("failed to start API server")
	}
}

func DbController(config factory.MongoConfig) database.DatabaseController {
	controller := factory.NewDatabaseController(factory.Mongo,
		[]database.DocumentBase{}, []database.DocumentBase{},
		config,
	)
	err := controller.Generate()
	if err != nil {
		panic(err)
	}
	err = controller.Init()
	if err != nil {
		panic(err)
	}

	return controller
}
