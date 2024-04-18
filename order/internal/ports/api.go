package ports

import "github.com/bruceneco/go-ms-grpc/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order *domain.Order) (*domain.Order, error)
}
