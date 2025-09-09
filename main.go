package main

import (
	"fmt"
	"log"

	"github.com/dennyaris/portal-news/internal/app"
	"github.com/dennyaris/portal-news/internal/config"
)

func main() {
	cfg := config.Load()
	application := app.New(cfg)
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("news portal (GORM) listening on %s (env=%s)\n", addr, cfg.AppEnv)
	if err := application.Fiber.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
