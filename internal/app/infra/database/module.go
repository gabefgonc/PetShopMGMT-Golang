package database

import (
	"go.uber.org/fx"

	"github.com/gabefgonc/PetShopMGMT/internal/app/infra/database/repository"
)

var Module = fx.Options(
	repository.Module,
	fx.Provide(
		NewPostgresDatabase,
	),
)
