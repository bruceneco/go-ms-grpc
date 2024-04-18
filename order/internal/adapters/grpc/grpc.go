package grpc

import (
	"context"
	"fmt"
	"github.com/bruceneco/go-ms-grpc/order/config"
	"github.com/bruceneco/go-ms-grpc/order/internal/application/core/domain"
	"github.com/bruceneco/go-ms-grpc/order/internal/ports"
	ordergrpc "github.com/bruceneco/go-ms-grpc/proto/go/order"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

type Adapter struct {
	api  ports.APIPort
	port uint16
	cfg  *config.Config
	ordergrpc.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, cfg *config.Config, lc fx.Lifecycle) *Adapter {
	port, err := strconv.Atoi(cfg.GetString(config.EnvAppPort))
	if err != nil {
		panic(fmt.Errorf("port must be an integer: %v\n", err))
	}

	adapter := &Adapter{api: api, port: uint16(port), cfg: cfg}
	return adapter
}

func (a *Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", a.port, err)
	}
	grpcServer := grpc.NewServer()
	ordergrpc.RegisterOrderServer(grpcServer, a)
	if a.cfg.GetString(config.EnvMode) == "development" {
		reflection.Register(grpcServer)
	}
	log.Printf("starting grpc server on port %v\n", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (a *Adapter) Create(ctx context.Context, request *ordergrpc.CreateOrderRequest) (*ordergrpc.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, orderItem := range request.Items {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.Name,
			Quantity:    orderItem.Quantity,
			UnitPrice:   orderItem.UnitPrice,
		})
	}
	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}
	return &ordergrpc.CreateOrderResponse{OrderId: result.ID}, nil
}
