package controller

import (
	"github.com/gin-gonic/gin"
	"go-gateway-service/client/gen-proto/auth"
	"go-gateway-service/client/service"
	"go-gateway-service/model"
	"net/http"
)

type AuthController struct {
	authClient *service.AuthClient
}

func NewAuthController(accountClient *service.AuthClient) AuthController {
	return AuthController{
		authClient: accountClient,
	}
}

// Login
// @Summary     Login
// @Description Get Access Token and Refresh Token
// @Tags        Auth
// @Accept      json
// @Param       user    body     auth.LoginRequest true "Login Information"
// @Success     200     {object} auth.LoginResponse_Data
// @Failure     400,500 {object} model.ErrorResponse
// @Router      /api/v2/login [POST]
func (controller *AuthController) Login(ctx *gin.Context) {
	req := new(auth.LoginRequest)

	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if res, err := controller.authClient.Login(req); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	} else {
		if res.GetSuccess() {
			ctx.JSON(http.StatusOK, res.GetData())
		} else {
			ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
		}
	}
}

// Register
// @Summary     Register
// @Description Register new User
// @Tags        Auth
// @Accept      json
// @Param       user    body     auth.RegisterRequest true "Register Information"
// @Success     200     {object} auth.RegisterResponse
// @Failure     400,500 {object} model.ErrorResponse
// @Router      /api/v2/register [POST]
func (controller *AuthController) Register(ctx *gin.Context) {
	req := new(auth.RegisterRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := controller.authClient.Register(req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	} else {
		if res.GetStatus() == http.StatusCreated {
			ctx.JSON(http.StatusOK, res.GetStatus())
		} else {
			ctx.JSON(http.StatusBadRequest, res.GetError())
		}
	}
}
