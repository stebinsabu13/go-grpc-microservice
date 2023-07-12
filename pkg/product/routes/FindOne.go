package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stebin13/go-grpc-microservice/pkg/product/pb"
)

func FindOne(c *gin.Context, prodcli pb.ProductServiceClient) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	res, err1 := prodcli.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, err1)
		return
	}
	c.JSON(http.StatusOK, &res)
}
