package app

import (
	"database/sql"
	"nk/account/internal/api"
	"nk/account/pkg/boot"
)

const (
	CONFIG = "config"
)

func Start() {
	Run(".").Wait()
}

func Run(path string) *boot.Launcher {

	s := &server{
		path: path,
		c:    boot.NewContext(),
		l:    boot.NewLauncher(),
	}

	s.registerDb()
	s.registerMigration()

	s.registerBus()

	s.registerAccount()

	s.registerApi()

	return s.l
}

type server struct {
	path string
	c    *boot.Context
	l    *boot.Launcher
}

func (s *server) registerDb() {
	db := boot.NewDbApp(
		boot.Load[boot.DbConfig](s.path, CONFIG, "db"),
	)
	s.l.Run(db)

	boot.Register(s.c, func(c *boot.Context) *sql.DB {
		return db.GetDb()
	})
}

func (s *server) registerMigration() {
	db := boot.Get[sql.DB](s.c)

	migration := boot.NewMigrationApp(
		boot.Load[boot.MigrationConfig](s.path, CONFIG, "migration"),
		db,
		s.path,
	)

	s.l.Run(migration)
}

func (s *server) registerApi() {
	apiApp := boot.NewApiApp(
		boot.Load[boot.ApiConfig](s.path, CONFIG, "api"),
	)

	apiApp.AddController(boot.Get[api.AccountCtr](s.c))

	s.l.Run(apiApp)
}

func (s *server) registerBus() {
	bus := boot.NewEventBus(
		boot.Load[boot.KafkaConfig](s.path, CONFIG, "bus"),
	)

	boot.Register(s.c, func(ctx *boot.Context) *boot.EventBus {
		return bus
	})
}
