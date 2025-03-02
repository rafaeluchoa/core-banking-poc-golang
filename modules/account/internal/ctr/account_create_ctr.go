package ctr

import (
	"nk/account/api"
)

// CreateAccount List accounts by customerId
// @Summary Create a new account to a customer
// @Description Create a new account to a customer
// @Tags accounts
// @Accept json
// @Produce json
// @Param request body api.AccountCreateReq true "Request"
// @Success 200 {object} api.AccountCreateRes
// @Router /account [post]
func (s *AccountCtr) Create(req *api.AccountCreateReq) *api.AccountCreateRes {
	account, err := s.createUc.Create(req.CustomerID)
	if err != nil {
		return &api.AccountCreateRes{
			Response: api.Response{
				Code:    "0001",
				Message: err.Error(),
			},
		}
	}

	return &api.AccountCreateRes{
		AccountID: account.ID,
	}
}
