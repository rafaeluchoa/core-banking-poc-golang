package uc

import (
	"fmt"
	"nk/account/internal/domain"
	"nk/account/internal/repo"
)

type AccountListUc struct {
	accountRepo *repo.AccountRepo
}

func NewAccountListUc(accountRepo *repo.AccountRepo) *AccountListUc {
	return &AccountListUc{
		accountRepo: accountRepo,
	}
}

func (s *AccountListUc) List(customerId string) (*[]domain.Account, error) {
	accounts, err := s.accountRepo.ListAllByCustomerId(customerId)
	if err != nil {
		return nil, fmt.Errorf("create account: %v", err)
	}

	return accounts, err
}
