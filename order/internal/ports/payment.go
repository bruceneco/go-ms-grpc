package ports

import "github.com/bruceneco/go-ms-grpc/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
