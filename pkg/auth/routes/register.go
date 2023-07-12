package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stebin13/go-grpc-microservice/pkg/auth/pb"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context, authcli pb.AuthServiceClient) {
	var body RegisterRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	res, err := authcli.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(int(http.StatusOK), &res)
}
