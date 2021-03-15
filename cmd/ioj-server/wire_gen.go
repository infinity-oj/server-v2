// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/infinity-oj/server-v2/internal/app/accounts"
	"github.com/infinity-oj/server-v2/internal/app/judgements"
	"github.com/infinity-oj/server-v2/internal/app/problems"
	"github.com/infinity-oj/server-v2/internal/app/processes"
	"github.com/infinity-oj/server-v2/internal/app/server"
	"github.com/infinity-oj/server-v2/internal/app/submissions"
	"github.com/infinity-oj/server-v2/internal/app/volumes"
	"github.com/infinity-oj/server-v2/internal/app/volumes/controllers"
	"github.com/infinity-oj/server-v2/internal/app/volumes/repositories"
	"github.com/infinity-oj/server-v2/internal/app/volumes/services"
	"github.com/infinity-oj/server-v2/internal/app/volumes/storages"
	"github.com/infinity-oj/server-v2/internal/pkg/configs"
	"github.com/infinity-oj/server-v2/internal/pkg/database"
	"github.com/infinity-oj/server-v2/internal/pkg/files"
	"github.com/infinity-oj/server-v2/internal/pkg/http"
	"github.com/infinity-oj/server-v2/internal/pkg/jaeger"
	"github.com/infinity-oj/server-v2/internal/pkg/log"
)

// Injectors from wire.go:

func CreateApp(cf string) (*server.Application, error) {
	viper, err := configs.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	serverOptions, err := server.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	databaseOptions, err := database.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	db, err := database.New(databaseOptions)
	if err != nil {
		return nil, err
	}
	repository := accounts.NewRepository(logger, db)
	service := accounts.NewService(logger, repository)
	controller := accounts.NewController(logger, service)
	initAccountGroupFn := accounts.CreateInitControllersFn(controller)
	judgementsRepository := judgements.NewRepository(logger, db)
	problemsRepository := problems.NewRepository(logger, db)
	processesRepository := processes.NewRepository(logger, db)
	submissionsRepository := submissions.NewRepository(logger, db)
	judgementsService := judgements.NewService(logger, judgementsRepository, problemsRepository, processesRepository, submissionsRepository)
	judgementsController := judgements.NewController(logger, judgementsService)
	initJudgementGroupFn := judgements.CreateInitControllersFn(judgementsController)
	submissionsService := submissions.NewService(logger, submissionsRepository, problemsRepository)
	submissionsController := submissions.NewController(logger, submissionsService)
	initSubmissionGroupFn := submissions.CreateInitControllersFn(submissionsController)
	problemsService := problems.NewService(logger, problemsRepository)
	problemsController := problems.NewController(logger, problemsService)
	initProblemGroupFn := problems.CreateInitControllersFn(problemsController)
	filesOptions, err := files.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	fileManager, err := files.New(filesOptions)
	if err != nil {
		return nil, err
	}
	storage := storages.NewFileManager(logger, fileManager)
	repositoriesRepository := repositories.NewRepository(logger, db)
	servicesService := services.NewVolumeService(logger, storage, repositoriesRepository)
	controllersController := controllers.New(logger, servicesService)
	initVolumeGroupFn := volumes.CreateInitControllersFn(controllersController)
	processesService := processes.NewService(logger, processesRepository)
	processesController := processes.NewController(logger, processesService)
	initProcessGroupFn := processes.CreateInitControllersFn(processesController)
	initControllers := server.CreateInitControllersFn(initAccountGroupFn, initJudgementGroupFn, initSubmissionGroupFn, initProblemGroupFn, initVolumeGroupFn, initProcessGroupFn)
	configuration, err := jaeger.NewConfiguration(viper, logger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	engine := http.NewRouter(httpOptions, logger, initControllers, tracer)
	httpServer, err := http.New(httpOptions, logger, engine)
	if err != nil {
		return nil, err
	}
	application, err := server.NewApp(serverOptions, logger, httpServer)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, configs.ProviderSet, http.ProviderSet, database.ProviderSet, jaeger.ProviderSet, files.ProviderSet, server.ProviderSet, accounts.ProviderSet, problems.ProviderSet, submissions.ProviderSet, judgements.ProviderSet, processes.ProviderSet, volumes.ProviderSet)
