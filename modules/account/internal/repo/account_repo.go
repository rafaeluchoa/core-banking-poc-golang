package repo

import (
	"database/sql"
	"fmt"
	"nk/account/internal/domain"

	"github.com/Masterminds/squirrel"
)

type AccountRepo struct {
	db *sql.DB
	p  squirrel.StatementBuilderType
}

func NewAccountRepo(db *sql.DB) *AccountRepo {
	return &AccountRepo{
		db: db,
		p:  squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (s *AccountRepo) Create(account *domain.Account) error {
	builder := s.p.Insert("account").
		Columns("id", "customer_id").
		Values(account.Id, account.CustomerId)

	cmd, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %v", err)
	}

	_, err = s.db.Exec(cmd, args...)
	if err != nil {
		return fmt.Errorf("failed to insert account: %v\n%s\n%v", err, cmd, args)
	}

	return nil
}

func (s *AccountRepo) GetById(id string) (*domain.Account, error) {
	builder := s.p.Select("id", "customer_id").
		From("account").
		Where(squirrel.Eq{"id": id})

	cmd, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	row := s.db.QueryRow(cmd, args...)
	account := &domain.Account{}
	if err := row.Scan(&account.Id, &account.CustomerId); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("account not found")
		}
		return nil, fmt.Errorf("failed to scan account: %v", err)
	}

	return account, nil
}

func (s *AccountRepo) ListAllByCustomerId(customerId string) (*[]domain.Account, error) {
	builder := s.p.Select("id", "customer_id").
		From("account").
		Where(squirrel.Eq{"customer_id": customerId})

	cmd, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	rows, err := s.db.Query(cmd, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var accounts []domain.Account
	for rows.Next() {
		var account domain.Account
		if err := rows.Scan(&account.Id, &account.CustomerId); err != nil {
			return nil, fmt.Errorf("failed to scan account: %v", err)
		}
		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows iteration: %v", err)
	}

	return &accounts, nil
}
