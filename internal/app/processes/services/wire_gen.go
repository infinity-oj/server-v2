// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package services

import (
	"github.com/google/wire"
	"github.com/infinity-oj/server-v2/internal/app/processes/repositories"
	"github.com/infinity-oj/server-v2/internal/pkg/configs"
	"github.com/infinity-oj/server-v2/internal/pkg/database"
	"github.com/infinity-oj/server-v2/internal/pkg/log"
)

// Injectors from wire.go:

func CreateUsersService(cf string, sto repositories.Repository) (ProcessesService, error) {
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
	processesService := NewProcessService(logger, sto)
	return processesService, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, configs.ProviderSet, database.ProviderSet, ProviderSet)
