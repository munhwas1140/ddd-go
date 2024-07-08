// Package entities holds al the entities that are shared across subdomains
package entity

import "github.com/google/uuid"

// Person is ans entity that represents a person in all domains
type Person struct {
	// ID an the identifier of the entity
	ID   uuid.UUID
	Name string
	Age  int
}
