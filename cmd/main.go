package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/stebin13/go-grpc-microservice/pkg/auth"
	"github.com/stebin13/go-grpc-microservice/pkg/config"
	"github.com/stebin13/go-grpc-microservice/pkg/order"
	"github.com/stebin13/go-grpc-microservice/pkg/product"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	r := gin.Default()
	authsrv := auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, authsrv)
	order.RegisterRoutes(r, &c, authsrv)
	r.Run(c.Port)
}
