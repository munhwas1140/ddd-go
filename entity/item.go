package entity

import "github.com/google/uuid"

// Item is ans entity that represents a item in all domains
type Item struct {
	// ID an the identifier of the entity
	ID          uuid.UUID
	Name        string
	Description string
}
