package domain

type Event struct {
	Id        string `json:"id"`
	EventType string `json:"event_type"`
	EntityId  string `json:"entity_id"`
}
