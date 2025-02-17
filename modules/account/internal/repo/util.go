package repo

import (
	"github.com/google/uuid"
)

func UUID() string {
	return uuid.Must(uuid.NewV7()).String()
}
