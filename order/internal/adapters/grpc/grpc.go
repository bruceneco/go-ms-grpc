package grpc

import (
	"context"
	"fmt"
	"github.com/bruceneco/go-ms-grpc/order/internal/application/core/domain"
	"github.com/bruceneco/go-ms-grpc/order/internal/ports"
	ordergrpc "github.com/bruceneco/go-ms-grpc/proto/go/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

type Adapter struct {
	api  ports.APIPort
	port uint16
	ordergrpc.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port uint16) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a *Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", a.port, err)
	}
	grpcServer := grpc.NewServer()
	ordergrpc.RegisterOrderServer(grpcServer, a)
	// TODO: use viper for env loading
	if os.Getenv("ENV") == "development" {
		reflection.Register(grpcServer)
	}
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
