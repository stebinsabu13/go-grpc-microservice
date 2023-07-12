package product

import (
	"github.com/gin-gonic/gin"
	"github.com/stebin13/go-grpc-microservice/pkg/auth"
	"github.com/stebin13/go-grpc-microservice/pkg/config"
	"github.com/stebin13/go-grpc-microservice/pkg/product/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authsrv *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authsrv)
	svc := &ServiceClient{
		Client: InitServCli(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
}

func (s *ServiceClient) CreateProduct(c *gin.Context) {
	routes.CreateProduct(c, s.Client)
}

func (s *ServiceClient) FindOne(c *gin.Context) {
	routes.FindOne(c, s.Client)
}
