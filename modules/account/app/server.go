package app

import (
	"database/sql"
	"nk/account/api"
	"nk/account/internal/boot"
	"nk/account/internal/repo"
	"nk/account/internal/uc"
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

func (s *server) registerAccount() {
	boot.Register(s.c, func(c *boot.Context) *repo.AccountRepo {
		db := boot.Get[sql.DB](c)
		return repo.NewAccountRepo(db)
	})

	boot.Register(s.c, func(c *boot.Context) *uc.AccountCreateUc {
		repo := boot.Get[repo.AccountRepo](c)
		return uc.NewAccountCreateUc(repo)
	})

	boot.Register(s.c, func(c *boot.Context) *uc.AccountListUc {
		repo := boot.Get[repo.AccountRepo](c)
		return uc.NewAccountListUc(repo)
	})

	boot.Register(s.c, func(c *boot.Context) *api.AccountCtr {
		createUc := boot.Get[uc.AccountCreateUc](c)
		listUc := boot.Get[uc.AccountListUc](c)
		return api.NewAccountCtr(createUc, listUc)
	})
}
