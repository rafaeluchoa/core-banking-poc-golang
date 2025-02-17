package uc

import (
	"fmt"
	"nk/account/internal/domain"
	"nk/account/internal/repo"
)

type AccountCreateUc struct {
	accountRepo *repo.AccountRepo
}

func NewAccountCreateUc(accountRepo *repo.AccountRepo) *AccountCreateUc {
	return &AccountCreateUc{
		accountRepo: accountRepo,
	}
}

func (s *AccountCreateUc) Create(customerId string) (*domain.Account, error) {
	account := &domain.Account{
		Id:         repo.UUID(),
		CustomerId: customerId,
	}

	err := s.accountRepo.Create(account)
	if err != nil {
		return nil, fmt.Errorf("create account: %v", err)
	}

	return account, err
}
