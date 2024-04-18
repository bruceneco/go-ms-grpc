package api

import (
	"github.com/bruceneco/go-ms-grpc/order/internal/application/core/domain"
	"github.com/bruceneco/go-ms-grpc/order/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db,
	}
}

func (a *Application) PlaceOrder(order *domain.Order) (*domain.Order, error) {
	err := a.db.Save(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
