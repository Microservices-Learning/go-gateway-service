package common

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go-gateway-service/common/constants"
	"go-gateway-service/model"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"strings"
)

const (
	UnauthorizedForbidden = "You do not have permission to access this feature"
)

func GetMetadataFromContext(ctx *gin.Context) metadata.MD {
	md, exists := ctx.Get("md")
	var meta metadata.MD
	if exists {
		meta = md.(metadata.MD)
	}
	return meta
}

func GetPrincipalFromContext(ctx *gin.Context) *model.GrpcPrincipal {
	md := GetMetadataFromContext(ctx)
	bearerToken := md.Get("Authorization")[0]
	jwtToken := strings.Replace(bearerToken, constants.GrpcJwtHeaderPrefix, "", 1)
	return ParseToken(jwtToken)
}

func GenerateJwtToken(method jwt.SigningMethod, claims jwt.Claims, secret string) (string, error) {
	token, err := jwt.NewWithClaims(method, claims).SignedString([]byte(secret))

	if err != nil {
		log.Println("Failed when generating gRPC token")
	}

	return token, err
}

func ParseToken(tokenString string) *model.GrpcPrincipal {
	if tokenString == "" || tokenString == "null" {
		log.Println("No token string provided")
		return nil
	}

	token, _ := jwt.ParseWithClaims(tokenString, &model.GrpcPrincipal{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.GrpcJwtSecret), nil
	})

	claims, _ := token.Claims.(*model.GrpcPrincipal)

	return claims
}

func HasRole(role string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		principle := GetPrincipalFromContext(ctx)
		for _, userRole := range principle.Roles {
			if userRole == role {
				ctx.Next()
				return
			}
		}
		abort(ctx)
		return
	}
}

func HasAnyRole(roles []string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		principle := GetPrincipalFromContext(ctx)
		for _, role := range principle.Roles {
			if ContainsString(roles, role) {
				ctx.Next()
				return
			}
		}
		abort(ctx)
		return
	}
}

func HasAllRoles(roles []string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		principle := GetPrincipalFromContext(ctx)
		for _, role := range roles {
			if ContainsString(principle.Roles, role) == false {
				abort(ctx)
				return
			}
		}
		ctx.Next()
		return
	}
}

func HasPermission(permission string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		principle := GetPrincipalFromContext(ctx)
		for _, userRole := range principle.Permissions {
			if userRole == permission {
				ctx.Next()
				return
			}
		}
		abort(ctx)
		return
	}
}

func HasAnyPermission(permissions []string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		principle := GetPrincipalFromContext(ctx)
		for _, permission := range principle.Permissions {
			if ContainsString(permissions, permission) {
				ctx.Next()
				return
			}
		}
		abort(ctx)
		return
	}
}

func HasAllPermissions(permissions []string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		principle := GetPrincipalFromContext(ctx)
		for _, permission := range permissions {
			if ContainsString(principle.Permissions, permission) == false {
				abort(ctx)
				return
			}
		}
		ctx.Next()
		return
	}
}

func HasAnyPermissionOrRole(permissions []string, roles []string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		principle := GetPrincipalFromContext(ctx)
		for _, permission := range principle.Permissions {
			if ContainsString(permissions, permission) {
				ctx.Next()
				return
			}
		}

		for _, role := range principle.Roles {
			if ContainsString(roles, role) {
				ctx.Next()
				return
			}
		}

		abort(ctx)
		return
	}
}

func IsPublicUser(roleCodes []string) bool {
	return ContainsString(roleCodes, "PUBLIC_USERS")
}

func abort(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusForbidden, model.ErrorResponse{
		ErrorCode:        http.StatusText(http.StatusForbidden),
		ErrorDescription: UnauthorizedForbidden,
	})
}
