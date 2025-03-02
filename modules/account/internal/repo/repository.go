package repo

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type Repository[T any] struct {
	db      *sql.DB
	p       squirrel.StatementBuilderType
	factory func() *T
	fields  func(instance *T) []any
}

func (s *Repository[T]) exec(builder squirrel.InsertBuilder) error {
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

func (s *Repository[T]) row(builder squirrel.SelectBuilder) (*T, error) {
	cmd, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	instance := s.factory()
	fields := s.fields(instance)

	row := s.db.QueryRow(cmd, args...)
	if err := row.Scan(fields...); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("record not found")
		}

		return nil, fmt.Errorf("failed to scan record: %v", err)
	}

	return instance, nil
}

func (s *Repository[T]) rows(builder squirrel.SelectBuilder) ([]*T, error) {
	cmd, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	rows, err := s.db.Query(cmd, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var slice []*T

	for rows.Next() {
		instance := s.factory()
		fields := s.fields(instance)

		if err := rows.Scan(fields...); err != nil {
			return nil, fmt.Errorf("failed to scan record: %v", err)
		}

		slice = append(slice, instance)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in rows iteration: %v", err)
	}

	return slice, err
}
