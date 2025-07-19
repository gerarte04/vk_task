package main

import (
	"fmt"
	"log"
	_ "marketplace/docs"
	apiHttp "marketplace/internal/api/http"
	"marketplace/internal/config"
	repo "marketplace/internal/repository/postgres"
	"marketplace/internal/usecases/service"
	pkgConfig "marketplace/pkg/config"
	"marketplace/pkg/database/postgres"
	"marketplace/pkg/http/handlers"
	"marketplace/pkg/http/server"

	"github.com/go-chi/chi/v5"
)

// @title Marketplace API
// @version 1.0

// @host localhost:8080
// @BasePath /api/v1
func main() {
	appFlags := pkgConfig.ParseFlags()
    var cfg config.Config
    pkgConfig.MustLoadConfig(appFlags.ConfigPath, &cfg)

	log.Printf("Marketplace server is starting")

	pool, err := postgres.NewPostgresPool(cfg.PostgresCfg)
	if err != nil {
		log.Fatalf("Failed to connect PostgreSQL: %s", err.Error())
	} else {
		log.Printf("Connected to PostgreSQL successfully")
	}

	adRepo := repo.NewAdRepo(pool, cfg.SvcCfg)
	userRepo := repo.NewUserRepo(pool)

	adService := service.NewAdService(adRepo)
	authService, err := service.NewAuthService(userRepo, cfg.JwtCfg, cfg.CryptCfg)
	if err != nil {
		log.Fatalf("Failed to create AuthService: %s", err.Error())
	}

	adHandler, err := apiHttp.NewAdHandler(adService, cfg.PathCfg, cfg.SvcCfg, cfg.JwtCfg.PublicKeyPEM)
	if err != nil {
		log.Fatalf("Failed to create AdHandler: %s", err.Error())
	}

	authHandler := apiHttp.NewAuthHandler(authService, cfg.PathCfg, cfg.SvcCfg)

	log.Printf("All services were created successfully")

	r := chi.NewRouter()
	handlers.RouteHandlers(r, cfg.PathCfg.ApiPath,
		handlers.WithLogger(),
		handlers.WithRecovery(),
		handlers.WithSwagger(),
		adHandler.WithAdHandlers(),
		authHandler.WithAuthHandlers(),
	)

	addr := fmt.Sprintf("%s:%s", cfg.HttpCfg.Host, cfg.HttpCfg.Port)
	log.Printf("Starting HTTP server at %s...", addr)

	if err := server.CreateServer(addr, r); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
