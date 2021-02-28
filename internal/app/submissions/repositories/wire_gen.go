// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package repositories

import (
	"github.com/google/wire"
	"github.com/infinity-oj/server-v2/internal/pkg/configs"
	"github.com/infinity-oj/server-v2/internal/pkg/database"
	"github.com/infinity-oj/server-v2/internal/pkg/log"
)

// Injectors from wire.go:

func CreateDetailRepository(f string) (Repository, error) {
	viper, err := configs.New(f)
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
	repository := NewMysqlSubmissionsRepository(logger, db)
	return repository, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, configs.ProviderSet, database.ProviderSet, ProviderSet)
