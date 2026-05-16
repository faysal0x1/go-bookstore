package bootstrap

import (
	"github.com/faysal0x1/go-bookstore/pkg/config"
	"github.com/faysal0x1/go-bookstore/pkg/controllers"
	"github.com/faysal0x1/go-bookstore/pkg/database"
	"github.com/faysal0x1/go-bookstore/pkg/middleware"
	"github.com/faysal0x1/go-bookstore/pkg/repository"
	"github.com/faysal0x1/go-bookstore/pkg/services"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

var AllModules = fx.Options(
	config.Module,
	database.Module,
	repository.Module,
	services.Module,
	controllers.Module,
	middleware.Module,
	fx.Provide(mux.NewRouter),
)
