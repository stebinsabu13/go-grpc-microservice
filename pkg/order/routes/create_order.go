package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stebin13/go-grpc-microservice/pkg/order/pb"
)

type CreateOrderReq struct {
	ProductId int64 `json:"productid"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(c *gin.Context, ordercli pb.OrderServiceClient) {
	var body CreateOrderReq
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	userid, ok := c.Get("userId")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "unauthorized access")
		return
	}
	res, err := ordercli.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    userid.(int64),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusCreated, &res)
}
