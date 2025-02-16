package app

import (
	"nk/account/api"
	"nk/account/internal/bootstrap"
)

const (
	CONFIG = "config"
)

func Run(path string) *bootstrap.Launcher {
	launcher := bootstrap.NewLauncher()

	apiApp := bootstrap.NewApiApp(
		bootstrap.Load[bootstrap.ApiConfig](path, CONFIG, "api"),
	)

	apiApp.AddController(api.NewAccountCtr())
	launcher.Run(apiApp)

	return launcher
}

func Start() {
	Run(".").Wait()
}
