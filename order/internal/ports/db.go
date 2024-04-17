package ports

import "github.com/bruceneco/go-ms-grpc/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
