package order

import (
	"github.com/gin-gonic/gin"
	"github.com/stebin13/go-grpc-microservice/pkg/auth"
	"github.com/stebin13/go-grpc-microservice/pkg/config"
	"github.com/stebin13/go-grpc-microservice/pkg/order/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authsrv *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authsrv)
	svc := &ServiceClient{
		Client: InitServCli(c),
	}

	route := r.Group("/order")
	route.Use(a.AuthRequired)
	route.POST("/", svc.CreateOrder)
}

func (s *ServiceClient) CreateOrder(c *gin.Context) {
	routes.CreateOrder(c, s.Client)
}
