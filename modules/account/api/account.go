package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

const (
	API_ACCOUNT = "/account"
)

type Account struct {
	AccountId  string
	CustomerId string
	Balance    float32
}

type AccountCtr struct {
}

func NewAccountCtr() *AccountCtr {
	return &AccountCtr{}
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
	AccountId string `json:"accountId"`
}

func (s *AccountCtr) Create(req *AccountCreateReq) *AccountCreateRes {
	log.Println(req.CustomerId)
	return &AccountCreateRes{
		AccountId: "123",
	}
}

type AccountListReq struct {
	CustomerId string `json:"customerId"`
}

type AccountListRes struct {
	Accounts []*Account `json:"accounts"`
}

func (s *AccountCtr) List(req *AccountListReq) []*Account {
	log.Println(req.CustomerId)
	return []*Account{
		{
			AccountId:  "123",
			CustomerId: req.CustomerId,
		},
		{
			AccountId:  "124",
			CustomerId: req.CustomerId,
		},
	}
}
