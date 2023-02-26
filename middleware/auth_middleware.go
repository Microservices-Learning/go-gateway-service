package middleware

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	grpc "go-gateway-service/client/gen-proto/common"
	"go-gateway-service/common"
	"go-gateway-service/common/constants"
	"go-gateway-service/model"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	UnauthorizedTokenMissing = "Token is missing"
	UnauthorizedTokenInvalid = "The access token is not valid"
)

type AdditionalClaims struct {
	Permissions []string `json:"permissions"`
	RoleIds     []string `json:"role_ids"`
}

func (middleware *Middleware) AuthMiddleware(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")

	if authorization == "" {
		abort(ctx, UnauthorizedTokenMissing)
		return
	}

	tokenString := strings.Split(authorization, " ")[1]

	if isInvalidToken(tokenString) {
		abort(ctx, UnauthorizedTokenInvalid)
		return
	}

	principal := common.ParseToken(tokenString)

	md := make(model.CacheableMetadata)
	if cached, err := middleware.redisClient.Get(constants.RedisPrincipalPrefix+principal.ExternalId, ctx); err != nil {
		authResp, authErr := middleware.obtainUserPermissions(principal)
		if authErr != nil {
			abort(ctx, authErr.Error())
			return
		}

		principal.Permissions = authResp.Permissions
		principal.RoleIds = authResp.RoleIds

		grpcToken, genErr := common.GenerateJwtToken(jwt.SigningMethodHS256, principal, constants.GrpcJwtSecret)
		if genErr != nil {
			abort(ctx, genErr.Error())
			return
		}

		md.Put("Authorization", constants.GrpcJwtHeaderPrefix+grpcToken)

		cacheErr := middleware.cachePrincipal(principal, md, ctx)
		if cacheErr != nil {
			abort(ctx, cacheErr.Error())
			return
		}
	} else {
		if cacheErr := json.Unmarshal([]byte(cached), &md); cacheErr != nil {
			log.Println("Failed when unmarshalling metadata", cacheErr)
			abort(ctx, "Unmarshalling error "+cacheErr.Error())
		}
	}

	ctx.Set("md", md.AsGrpcMetadata())
	ctx.Next()
}
