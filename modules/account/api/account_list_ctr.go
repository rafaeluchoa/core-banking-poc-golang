package api

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
