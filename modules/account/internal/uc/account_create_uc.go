package uc

import (
	"fmt"
	"log"
	"nk/account/internal/domain"
	"nk/account/internal/repo"
	"nk/account/pkg/boot"
)

const (
	TOPIC_ACCOUNT_STATUS_CHANGED = "AccountStatusChanged"
)

type AccountCreateUc struct {
	accountRepo *repo.AccountRepo
	eventRepo   *repo.EventRepo
	producer    *boot.EventProducer
}

func NewAccountCreateUc(
	accountRepo *repo.AccountRepo,
	eventRepo *repo.EventRepo,
	eventBus *boot.EventBus,
) *AccountCreateUc {
	uc := &AccountCreateUc{
		accountRepo: accountRepo,
		eventRepo:   eventRepo,
		producer:    eventBus.NewProducer(TOPIC_ACCOUNT_STATUS_CHANGED),
	}

	eventBus.NewConsumer(TOPIC_ACCOUNT_STATUS_CHANGED).
		On(uc.accountStatusChangedHandler)

	return uc
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

	// TODO: test
	s.eventRepo.Create(&domain.Event{
		Id:        repo.UUID(),
		EventType: TOPIC_ACCOUNT_STATUS_CHANGED,
		EntityId:  account.Id,
	})
	// s.producer.Pub(account.Id)

	return account, err
}

func (s *AccountCreateUc) accountStatusChangedHandler(accountId string, err error) {
	if err != nil {
		log.Printf("Error on receive %v", err)
		return
	}
	log.Println("Account Created ", accountId)
}
