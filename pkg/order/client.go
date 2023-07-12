package order

import (
	"fmt"

	"github.com/stebin13/go-grpc-microservice/pkg/config"
	"github.com/stebin13/go-grpc-microservice/pkg/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServCli(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSuvUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}
