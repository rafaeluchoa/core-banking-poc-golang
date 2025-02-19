package uc

import (
	"fmt"
	"nk/account/internal/boot"
	"nk/account/internal/domain"
	"nk/account/internal/repo"
)

const (
	TOPIC_ACCOUNT_STATUS_CHANGED = "AccountStatusChanged"
)

type AccountCreateUc struct {
	accountRepo *repo.AccountRepo
	producer    *boot.EventProducer
}

func NewAccountCreateUc(
	accountRepo *repo.AccountRepo,
	eventBus *boot.EventBus,
) *AccountCreateUc {
	return &AccountCreateUc{
		accountRepo: accountRepo,
		producer:    eventBus.NewProducer(TOPIC_ACCOUNT_STATUS_CHANGED),
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

	s.producer.Pub(account.Id)

	return account, err
}
