package ctr

import "nk/account/api"

// GetAccount List accounts by customerId
// @Summary Return account list
// @Description Return account from a customerId
// @Tags accounts
// @Accept json
// @Produce json
// @Param request query api.AccountListReq true "Request"
// @Success 200 {object} api.AccountListRes
// @Router /account [get]
func (s *AccountCtr) List(req *api.AccountListReq) *api.AccountListRes {
	accounts, err := s.listUc.List(req.CustomerId)
	if err != nil {
		return &api.AccountListRes{
			Response: api.Response{
				Code:    "0002",
				Message: err.Error(),
			},
		}
	}

	return &api.AccountListRes{
		Accounts: toListAccount(accounts),
	}
}
