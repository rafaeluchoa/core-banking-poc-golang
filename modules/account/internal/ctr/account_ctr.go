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
	app.Get(api.APIAccount, func(c *fiber.Ctx) error {
		var req api.AccountListReq
		err := c.QueryParser(&req)

		if err != nil {
			return err
		}

		return c.JSON(s.List(&req))
	})

	app.Post(api.APIAccount, func(c *fiber.Ctx) error {
		var req api.AccountCreateReq
		err := c.BodyParser(&req)

		if err != nil {
			return err
		}

		return c.JSON(s.Create(&req))
	})
}

func toAccount(d *domain.Account) api.Account {
	return api.Account{
		ID:         d.ID,
		CustomerID: d.CustomerID,
	}
}

func toListAccount(l []*domain.Account) []api.Account {
	list := make([]api.Account, len(l))
	for i, d := range l {
		list[i] = toAccount(d)
	}

	return list
}
