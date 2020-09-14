package app

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type App struct {
	HTTPServer *HTTPServer
}

func (app *App) Init() {

	// Initializing configuration
	if os.Getenv("GO_ENV") == "production" {
		viper.SetConfigName("config")
		viper.AddConfigPath("./config")
	} else {
		viper.SetConfigName("config.local")
		viper.AddConfigPath("./config")
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	app.HTTPServer = &HTTPServer{}
	app.HTTPServer.Init(app)
}

func (app *App) Run() {
	app.HTTPServer.Start()
}
