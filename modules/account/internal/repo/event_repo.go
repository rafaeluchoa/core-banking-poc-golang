package repo

import (
	"database/sql"
	"nk/account/internal/domain"

	"github.com/Masterminds/squirrel"
)

type EventRepo struct {
	Repository[domain.Event]
}

func NewEventRepo(db *sql.DB) *EventRepo {
	return &EventRepo{
		Repository: Repository[domain.Event]{
			db: db,
			p:  squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
			factory: func() *domain.Event {
				return &domain.Event{}
			},
			fields: func(instance *domain.Event) []any {
				return []any{
					&instance.ID,
					&instance.EventType,
					&instance.EntityID,
				}
			},
		},
	}
}

func (s *EventRepo) Create(event *domain.Event) error {
	return s.exec(s.p.Insert("event").
		Columns("id", "event_type", "entity_id").
		Values(event.ID, event.EventType, event.EntityID),
	)
}
