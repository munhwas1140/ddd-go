package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/munhwas1140/ddd-go/aggregate"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy beverage", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	peenuts, err := aggregate.NewProduct("Peanuts", "Snacks", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := aggregate.NewProduct("Wine", "nasty drink", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{
		beer, peenuts, wine,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("munhwas1140")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	oreder := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), oreder)
	if err != nil {
		t.Error(err)
	}

}
