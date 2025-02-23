package repo

import (
	"database/sql"
	"nk/account/internal/domain"

	"github.com/Masterminds/squirrel"
)

type AccountRepo struct {
	Repository[domain.Account]
}

func NewAccountRepo(db *sql.DB) *AccountRepo {
	return &AccountRepo{
		Repository: Repository[domain.Account]{
			db: db,
			p:  squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
			factory: func() *domain.Account {
				return &domain.Account{}
			},
			fields: func(instance *domain.Account) []any {
				return []any{
					&instance.Id,
					&instance.CustomerId,
				}
			},
		},
	}
}

func (s *AccountRepo) Create(account *domain.Account) error {
	return s.exec(s.p.Insert("account").
		Columns("id", "customer_id").
		Values(account.Id, account.CustomerId),
	)
}

func (s *AccountRepo) GetById(id string) (*domain.Account, error) {
	return s.row(s.p.Select("id", "customer_id").
		From("account").
		Where(squirrel.Eq{"id": id}),
	)
}

func (s *AccountRepo) ListAllByCustomerId(customerId string) ([]*domain.Account, error) {
	return s.rows(s.p.Select("id", "customer_id").
		From("account").
		Where(squirrel.Eq{"customer_id": customerId}),
	)
}
