package app

import (
	"database/sql"
	"nk/account/internal/ctr"
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

	boot.Register(s.c, func(c *boot.Context) *repo.EventRepo {
		db := boot.Get[sql.DB](c)
		return repo.NewEventRepo(db)
	})
}

func (s *server) registerAccountUcs() {
	boot.Register(s.c, func(c *boot.Context) *uc.AccountCreateUc {
		accountRepo := boot.Get[repo.AccountRepo](c)
		eventRepo := boot.Get[repo.EventRepo](c)
		eventBus := boot.Get[boot.EventBus](c)
		return uc.NewAccountCreateUc(accountRepo, eventRepo, eventBus)
	})

	boot.Register(s.c, func(c *boot.Context) *uc.AccountListUc {
		repo := boot.Get[repo.AccountRepo](c)
		return uc.NewAccountListUc(repo)
	})
}

func (s *server) registerAccountCtrs() {
	boot.Register(s.c, func(c *boot.Context) *ctr.AccountCtr {

		createUc := boot.Get[uc.AccountCreateUc](c)
		listUc := boot.Get[uc.AccountListUc](c)
		return ctr.NewAccountCtr(createUc, listUc)
	})
}
