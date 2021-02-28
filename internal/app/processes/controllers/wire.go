// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/infinity-oj/server-v2/internal/app/processes/repositories"
	"github.com/infinity-oj/server-v2/internal/app/processes/services"
	"github.com/infinity-oj/server-v2/internal/pkg/configs"
	"github.com/infinity-oj/server-v2/internal/pkg/database"
	"github.com/infinity-oj/server-v2/internal/pkg/log"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	configs.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	repositories.ProviderSet,

	ProviderSet,
)

func CreateProblemController(cf string) (Controller, error) {
	panic(wire.Build(providerSet))
}
