package app

import (
	"database/sql"
	"nk/account/internal/api"
	"nk/account/internal/repo"
	"nk/account/internal/uc"
	"nk/account/pkg/boot"
)

func (s *server) registerAccount() {
	s.registerAccountRepos()
	s.registerAccountUcs()
	s.registerAccountCtrs()
}

func (s *server) registerAccountRepos() {
	boot.Register(s.c, func(c *boot.Context) *repo.AccountRepo {
		db := boot.Get[sql.DB](c)
		return repo.NewAccountRepo(db)
	})
}

func (s *server) registerAccountUcs() {
	boot.Register(s.c, func(c *boot.Context) *uc.AccountCreateUc {
		repo := boot.Get[repo.AccountRepo](c)
		eventBus := boot.Get[boot.EventBus](c)
		return uc.NewAccountCreateUc(repo, eventBus)
	})

	boot.Register(s.c, func(c *boot.Context) *uc.AccountListUc {
		repo := boot.Get[repo.AccountRepo](c)
		return uc.NewAccountListUc(repo)
	})
}

func (s *server) registerAccountCtrs() {
	boot.Register(s.c, func(c *boot.Context) *api.AccountCtr {

		createUc := boot.Get[uc.AccountCreateUc](c)
		listUc := boot.Get[uc.AccountListUc](c)
		return api.NewAccountCtr(createUc, listUc)
	})
}
