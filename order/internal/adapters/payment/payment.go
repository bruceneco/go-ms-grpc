package payment

import (
	"context"
	"github.com/bruceneco/go-ms-grpc/order/config"
	"github.com/bruceneco/go-ms-grpc/order/internal/application/core/domain"
	"github.com/bruceneco/go-ms-grpc/order/internal/ports"
	paymentgrpc "github.com/bruceneco/go-ms-grpc/proto/go/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment paymentgrpc.PaymentClient
}

func (a *Adapter) Charge(order *domain.Order) error {
	_, err := a.payment.Create(context.Background(), &paymentgrpc.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}

func NewAdapter(cfg *config.Config) (ports.PaymentPort, error) {
	paymentServiceURL := cfg.GetString(config.EnvPaymentServiceURL)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(paymentServiceURL, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := paymentgrpc.NewPaymentClient(conn)
	return &Adapter{payment: client}, nil
}
