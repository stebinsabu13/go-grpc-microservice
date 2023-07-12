package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stebin13/go-grpc-microservice/pkg/product/pb"
)

type CreateProductReq struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

func CreateProduct(c *gin.Context, prodcli pb.ProductServiceClient) {
	var body CreateProductReq
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	res, err := prodcli.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, &res)
}
