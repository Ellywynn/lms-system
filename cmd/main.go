package main

import (
	"log"

	"github.com/ellywynn/lms-system/internal/app"
)

func main() {
	cfg := app.NewAppConfig()
	srv := app.NewApp(cfg)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
