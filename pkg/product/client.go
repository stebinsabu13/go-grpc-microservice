package product

import (
	"fmt"

	"github.com/stebin13/go-grpc-microservice/pkg/config"
	"github.com/stebin13/go-grpc-microservice/pkg/product/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServCli(c *config.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductSuvUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewProductServiceClient(cc)
}
