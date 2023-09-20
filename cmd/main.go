package main

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"

	"github.com/gabefgonc/PetShopMGMT/internal/app/infra/config"
	"github.com/gabefgonc/PetShopMGMT/internal/app/infra/database"
	"github.com/gabefgonc/PetShopMGMT/internal/app/infra/httpapi"
	"github.com/gabefgonc/PetShopMGMT/internal/app/infra/httpapi/controllers"
	"github.com/gabefgonc/PetShopMGMT/internal/app/infra/httpapi/routers"
)

func main() {
	app := fx.New(
		config.Module,
		httpapi.Module,
		controllers.Module,
		routers.Module,
		database.Module,
		fx.Invoke(func(*fasthttp.Server) {}),
	)
	app.Run()
}
