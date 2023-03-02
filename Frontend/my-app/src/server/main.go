package main

import (
	"github.com/server/api/app"
	"github.com/server/api/config"
	"os"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(os.Getenv("PORT"))
}