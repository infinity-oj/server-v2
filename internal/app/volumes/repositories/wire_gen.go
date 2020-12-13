// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package repositories

import (
	"github.com/google/wire"
	"github.com/infinity-oj/server-v2/internal/pkg/config"
	"github.com/infinity-oj/server-v2/internal/pkg/database"
	"github.com/infinity-oj/server-v2/internal/pkg/files"
	"github.com/infinity-oj/server-v2/internal/pkg/log"
)

// Injectors from wire.go:

func CreateFileRepository(f string) (Repository, error) {
	viper, err := config.New(f)
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
	filesOptions, err := files.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	fileManager, err := files.New(filesOptions)
	if err != nil {
		return nil, err
	}
	repository := NewFileManager(logger, fileManager)
	return repository, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, files.ProviderSet, ProviderSet)
