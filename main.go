package main

import (
	"github.com/VReference/api/app"
	"github.com/VReference/api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}