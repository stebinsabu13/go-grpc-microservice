package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/stebin13/go-grpc-microservice/pkg/auth/routes"
	"github.com/stebin13/go-grpc-microservice/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)

	return svc
}

func (s *ServiceClient) Register(c *gin.Context) {
	routes.Register(c, s.Client)
}

func (s *ServiceClient) Login(c *gin.Context) {
	routes.Login(c, s.Client)
}
