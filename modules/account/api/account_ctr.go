package api

import (
	"nk/account/internal/domain"
	"nk/account/internal/uc"

	"github.com/gofiber/fiber/v2"
)

const (
	API_ACCOUNT = "/account"
)

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

type AccountCreateReq struct {
	CustomerId string `json:"customerId"`
}

type AccountCreateRes struct {
	Response
	AccountId string `json:"accountId"`
}

func (s *AccountCtr) Create(req *AccountCreateReq) *AccountCreateRes {
	account, err := s.createUc.Create(req.CustomerId)
	if err != nil {
		return &AccountCreateRes{
			Response: Response{
				Code:    "0001",
				Message: err.Error(),
			},
		}
	}

	return &AccountCreateRes{
		AccountId: account.Id,
	}
}

type AccountListReq struct {
	CustomerId string `json:"customerId"`
}

type AccountListRes struct {
	Response
	Accounts []Account `json:"accounts"`
}

func (s *AccountCtr) List(req *AccountListReq) *AccountListRes {
	accounts, err := s.listUc.List(req.CustomerId)
	if err != nil {
		return &AccountListRes{
			Response: Response{
				Code:    "0002",
				Message: err.Error(),
			},
		}
	}

	return &AccountListRes{
		Accounts: toListAccount(*accounts),
	}
}
