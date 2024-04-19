package adapters

import (
	"github.com/bruceneco/go-ms-grpc/order/internal/adapters/db"
	"github.com/bruceneco/go-ms-grpc/order/internal/adapters/grpc"
	"github.com/bruceneco/go-ms-grpc/order/internal/adapters/payment"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	grpc.NewAdapter,
	db.NewAdapter,
	payment.NewAdapter,
)
