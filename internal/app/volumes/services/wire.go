// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/infinity-oj/server-v2/internal/app/volumes/repositories"
	"github.com/infinity-oj/server-v2/internal/app/volumes/storages"
	"github.com/infinity-oj/server-v2/internal/pkg/configs"
	"github.com/infinity-oj/server-v2/internal/pkg/database"
	"github.com/infinity-oj/server-v2/internal/pkg/log"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	configs.ProviderSet,
	database.ProviderSet,
	ProviderSet,
)

func CreateVolumesService(cf string, sto storages.Storage, repository repositories.Repository) (Service, error) {
	panic(wire.Build(testProviderSet))
}
