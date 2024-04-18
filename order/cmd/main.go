package main

import (
	"github.com/bruceneco/go-ms-grpc/order/config"
	"github.com/bruceneco/go-ms-grpc/order/internal/adapters"
	"github.com/bruceneco/go-ms-grpc/order/internal/adapters/grpc"
	"github.com/bruceneco/go-ms-grpc/order/internal/application/core"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		config.Module,
		adapters.Module,
		core.Module,
		fx.Invoke(func(server *grpc.Adapter) {
			server.Run()
		}),
	)

	app.Run()
}
