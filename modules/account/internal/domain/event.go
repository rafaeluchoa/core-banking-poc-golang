package domain

type Event struct {
	ID        string `json:"id"`
	EventType string `json:"event_type"`
	EntityID  string `json:"entity_id"`
}
