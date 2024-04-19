package api

import (
	"github.com/bruceneco/go-ms-grpc/order/internal/application/core/domain"
	"github.com/bruceneco/go-ms-grpc/order/internal/ports"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) ports.APIPort {
	return &Application{
		db,
		payment,
	}
}

func (a *Application) PlaceOrder(order *domain.Order) (*domain.Order, error) {
	err := a.db.Save(order)
	if err != nil {
		return nil, err
	}
	paymentErr := a.payment.Charge(order)
	if paymentErr != nil {
		return nil, paymentErr
	}
	return order, nil
}
