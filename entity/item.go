package entity

import (
	"github.com/google/uuid"
)

// Item --> Item is an entity that represents a person in all domains
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
