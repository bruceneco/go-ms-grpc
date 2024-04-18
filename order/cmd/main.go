package main

import (
	"github.com/bruceneco/go-ms-grpc/order/config"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(config.Module)

	app.Run()
}
