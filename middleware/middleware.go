package middleware

import (
	"emarket-gateway-service/client"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	redisClient      *client.RedisClient
	permissionClient *client.PermissionClient
	leaseClient      *client.LeaseClient
}

func NewMiddleware(
	redisClient *client.RedisClient,
	permissionClient *client.PermissionClient,
	leaseClient *client.LeaseClient,
) Middleware {
	return Middleware{
		redisClient:      redisClient,
		permissionClient: permissionClient,
		leaseClient:      leaseClient,
	}
}

func (middleware *Middleware) GinMiddleware(ctx *gin.Context) {
	ctx.Next()
}
