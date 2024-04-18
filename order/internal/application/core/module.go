package core

import (
	"github.com/bruceneco/go-ms-grpc/order/internal/application/core/api"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	api.NewApplication,
)
