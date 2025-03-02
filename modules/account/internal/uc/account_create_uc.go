package uc

import (
	"fmt"
	"log"
	"nk/account/internal/domain"
	"nk/account/internal/repo"
	"nk/account/pkg/boot"
)

const (
	TopicAccountStatusChanged = "AccountStatusChanged"
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
		producer:    eventBus.NewProducer(TopicAccountStatusChanged),
	}

	eventBus.NewConsumer(TopicAccountStatusChanged).
		On(uc.accountStatusChangedHandler)

	return uc
}

func (s *AccountCreateUc) Create(customerID string) (*domain.Account, error) {
	account := &domain.Account{
		ID:         repo.UUID(),
		CustomerID: customerID,
	}

	err := s.accountRepo.Create(account)
	if err != nil {
		return nil, fmt.Errorf("create account: %v", err)
	}

	// TODO: test, modified to CDC, testing for while
	err = s.eventRepo.Create(&domain.Event{
		ID:        repo.UUID(),
		EventType: TopicAccountStatusChanged,
		EntityID:  account.ID,
	})
	if err != nil {
		return nil, err
	}
	// s.producer.Pub(account.Id)

	return account, err
}

func (s *AccountCreateUc) accountStatusChangedHandler(event *domain.Event, err error) {
	if err != nil {
		log.Printf("Error on receive %v", err)
		return
	}
	log.Printf("Account Created %v", event.EntityID)
}
