package app

import (
	"database/sql"
	"nk/account/api"
	"nk/account/internal/bootstrap"
)

const (
	CONFIG = "config"
)

func Start() {
	Run(".").Wait()
}

func Run(path string) *bootstrap.Launcher {
	l := bootstrap.NewLauncher()

	db := runDb(path, l)
	runMigration(path, l, db.GetDb())
	runApi(path, l)

	return l
}

func runDb(path string, l *bootstrap.Launcher) *bootstrap.DbApp {
	db := bootstrap.NewDbApp(
		bootstrap.Load[bootstrap.DbConfig](path, CONFIG, "db"),
	)
	l.Run(db)
	return db
}

func runMigration(path string, l *bootstrap.Launcher, db *sql.DB) {
	migration := bootstrap.NewMigrationApp(
		bootstrap.Load[bootstrap.MigrationConfig](path, CONFIG, "migration"),
		db,
	)
	l.Run(migration)
}

func runApi(path string, l *bootstrap.Launcher) {
	apiApp := bootstrap.NewApiApp(
		bootstrap.Load[bootstrap.ApiConfig](path, CONFIG, "api"),
	)

	apiApp.AddController(api.NewAccountCtr())

	l.Run(apiApp)
}
