package main

import (
	"userModel/app"
)

func main() {

	// Initializing application
	app := app.App{}
	app.Init()

	// Start application
	app.Run()
}
