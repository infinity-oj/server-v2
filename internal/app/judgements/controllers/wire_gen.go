// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/infinity-oj/server-v2/internal/app/judgements/repositories"
	"github.com/infinity-oj/server-v2/internal/app/judgements/services"
	repositories2 "github.com/infinity-oj/server-v2/internal/app/processes/repositories"
	repositories3 "github.com/infinity-oj/server-v2/internal/app/submissions/repositories"
	"github.com/infinity-oj/server-v2/internal/pkg/config"
	"github.com/infinity-oj/server-v2/internal/pkg/database"
	"github.com/infinity-oj/server-v2/internal/pkg/log"
)

// Injectors from wire.go:

func CreateJudgementsController(cf string, sto repositories.Repository, sto2 repositories2.Repository, sto3 repositories3.Repository) (Controller, error) {
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
	judgementsService := services.NewJudgementsService(logger, sto, sto2, sto3)
	controller := New(logger, judgementsService)
	return controller, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, services.ProviderSet, ProviderSet)