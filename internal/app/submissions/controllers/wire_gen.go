// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package controllers

import (
	"github.com/google/wire"
	repositories2 "github.com/infinity-oj/server-v2/internal/app/problems/repositories"
	"github.com/infinity-oj/server-v2/internal/app/submissions/repositories"
	"github.com/infinity-oj/server-v2/internal/app/submissions/services"
	"github.com/infinity-oj/server-v2/internal/pkg/config"
	"github.com/infinity-oj/server-v2/internal/pkg/database"
	"github.com/infinity-oj/server-v2/internal/pkg/log"
)

// Injectors from wire.go:

func CreateSubmissionController(cf string) (Controller, error) {
	viper, err := config.New(cf)
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
	databaseOptions, err := database.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	db, err := database.New(databaseOptions)
	if err != nil {
		return nil, err
	}
	repository := repositories.NewMysqlSubmissionsRepository(logger, db)
	repositoriesRepository := repositories2.New(logger, db)
	submissionsService := services.NewSubmissionService(logger, repository, repositoriesRepository)
	controller := New(logger, submissionsService)
	return controller, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, services.ProviderSet, repositories.ProviderSet, repositories2.ProviderSet, ProviderSet)
