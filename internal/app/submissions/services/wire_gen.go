// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package services

import (
	"github.com/google/wire"
	repositories3 "github.com/infinity-oj/server-v2/internal/app/judgements/repositories"
	repositories2 "github.com/infinity-oj/server-v2/internal/app/problems/repositories"
	"github.com/infinity-oj/server-v2/internal/app/submissions/repositories"
	"github.com/infinity-oj/server-v2/internal/pkg/config"
	"github.com/infinity-oj/server-v2/internal/pkg/database"
	"github.com/infinity-oj/server-v2/internal/pkg/log"
)

// Injectors from wire.go:

func CreateSubmissionsService(cf string, sto repositories.Repository, sto2 repositories2.Repository, sto3 repositories3.Repository) (SubmissionsService, error) {
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
	submissionsService := NewSubmissionService(logger, sto, sto2, sto3)
	return submissionsService, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, ProviderSet)
