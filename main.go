package main

import (
	"userModel/app/controller"
	"userModel/app"
)

func main() {

	// Initializing application
	app := app.App{}
	app.Init()

	user := controller.Users{}
	user.Init(app.HTTPServer)

	// Start application
	app.Run()
}
