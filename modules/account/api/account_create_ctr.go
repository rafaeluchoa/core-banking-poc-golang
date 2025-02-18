package api

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
