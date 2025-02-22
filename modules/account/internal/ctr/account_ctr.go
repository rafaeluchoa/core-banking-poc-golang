package ctr

import (
	"nk/account/api"
	"nk/account/internal/domain"
	"nk/account/internal/uc"

	"github.com/gofiber/fiber/v2"
)

type AccountCtr struct {
	createUc *uc.AccountCreateUc
	listUc   *uc.AccountListUc
}

func NewAccountCtr(
	createUc *uc.AccountCreateUc,
	listUc *uc.AccountListUc,
) *AccountCtr {
	return &AccountCtr{
		createUc: createUc,
		listUc:   listUc,
	}
}

func (s *AccountCtr) AddRoutes(app *fiber.App) {
	app.Get(api.API_ACCOUNT, func(c *fiber.Ctx) error {
		var req api.AccountListReq
		c.QueryParser(&req)
		return c.JSON(s.List(&req))
	})

	app.Post(api.API_ACCOUNT, func(c *fiber.Ctx) error {
		var req api.AccountCreateReq
		c.BodyParser(&req)
		return c.JSON(s.Create(&req))
	})
}

func toAccount(d domain.Account) api.Account {
	return api.Account{
		Id:         d.Id,
		CustomerId: d.CustomerId,
	}
}

func toListAccount(l []domain.Account) []api.Account {
	list := make([]api.Account, len(l))
	for i, d := range l {
		list[i] = toAccount(d)
	}
	return list
}
