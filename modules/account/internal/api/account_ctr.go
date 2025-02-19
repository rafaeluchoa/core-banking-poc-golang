package api

import (
	"nk/account/internal/domain"
	"nk/account/internal/uc"

	"github.com/gofiber/fiber/v2"
)

const (
	API_ACCOUNT = "/account"
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
	app.Get(API_ACCOUNT, func(c *fiber.Ctx) error {
		var req AccountListReq
		c.QueryParser(&req)
		return c.JSON(s.List(&req))
	})

	app.Post(API_ACCOUNT, func(c *fiber.Ctx) error {
		var req AccountCreateReq
		c.BodyParser(&req)
		return c.JSON(s.Create(&req))
	})
}

type Account struct {
	Id         string
	CustomerId string
}

func toAccount(d domain.Account) Account {
	return Account{
		Id:         d.Id,
		CustomerId: d.CustomerId,
	}
}

func toListAccount(l []domain.Account) []Account {
	list := make([]Account, len(l))
	for i, d := range l {
		list[i] = toAccount(d)
	}
	return list
}
