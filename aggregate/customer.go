// package aggregate holds our aggregate that combines many entities into a
// full object
package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/munhwas1140/ddd-go/entity"
	"github.com/munhwas1140/ddd-go/valueobject"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have valid name")
)

type Customer struct {
	// Person is the root entity of the customer
	// which means person.ID is the main identifier for the customer
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

// NewCustomer is a factory to create customer aggregate
// it will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}
