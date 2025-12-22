package main

import (
	"go-kit-2026/internal/app/bootstrap"
	"go-kit-2026/internal/app/config"
	"go-kit-2026/internal/app/routes"
	"log"
)

func main() {
	cfg := config.Load()
	app := bootstrap.NewApplication(cfg)
	services := bootstrap.InitServices(app.DB, cfg)
	routes.RegisterV1(
		app.App,
		services.Auth,
		services.Authorization,
		cfg.JWT.Secret,
	)
	for _, r := range app.App.Routes() {
		log.Printf("%s %s", r.Method, r.Path)
	}

	log.Printf("APP %s running on port %s", cfg.App.Name, cfg.App.Port)

	if err := app.App.Run(":" + cfg.App.Port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
