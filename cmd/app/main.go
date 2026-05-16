package main

// @title Go Bookstore API
// @version 1.0
// @description This is a professional bookstore management server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9010
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/faysal0x1/go-bookstore/pkg/bootstrap"
	"github.com/faysal0x1/go-bookstore/pkg/config"
	"github.com/faysal0x1/go-bookstore/pkg/database"
	"github.com/faysal0x1/go-bookstore/pkg/database/seeders"
	"github.com/faysal0x1/go-bookstore/pkg/middleware"
	"github.com/faysal0x1/go-bookstore/pkg/routes"
	"github.com/faysal0x1/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	// 1. Parse Command Line Flags
	migrate := flag.String("migrate", "", "Run migrations: 'fresh' drops all tables first")
	seed := flag.Bool("seed", false, "Run seeders after migrations")
	flag.Parse()

	// 2. Start Application
	fx.New(
		bootstrap.AllModules,
		fx.Invoke(
			utils.InitLogger,
			func(db *gorm.DB, shutdown fx.Shutdowner) {
				if *migrate == "fresh" {
					database.DropTables(db)
				}
				database.Migrate(db)
				if *seed {
					seeders.RunDatabaseSeeder(db)
				}
				if *migrate != "" || *seed {
					log.Println("CLI operation completed. Exiting.")
					os.Exit(0)
				}
			},
			setupGlobalMiddleware,
			routes.RegisterRoutes,
			startServer,
		),
	).Run()
}

func setupGlobalMiddleware(r *mux.Router) {
	r.Use(middleware.Logger)
}

func startServer(lc fx.Lifecycle, r *mux.Router, cfg *config.Config) {
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Printf("Server starting on port %s", cfg.ServerPort)
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping server...")
			return srv.Shutdown(ctx)
		},
	})
}
