package service

import (
	"context"
	"go-gateway-service/client/gen-proto/auth"
	"go-gateway-service/common/constants"
	"google.golang.org/grpc"
	"time"
)

/*
 * gRPC Account Client
 */

type AuthClient struct {
	authClient auth.AuthServiceClient
}

func NewAuthClient(cc *grpc.ClientConn) *AuthClient {
	accountClient := auth.NewAuthServiceClient(cc)
	return &AuthClient{authClient: accountClient}
}

func (service *AuthClient) Login(req *auth.LoginRequest) (res *auth.LoginResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	defer cancel()
	res, err = service.authClient.Login(ctx, req)
	return
}

func (service *AuthClient) Register(req *auth.RegisterRequest) (res *auth.RegisterResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	defer cancel()
	res, err = service.authClient.Register(ctx, req)
	return
}

func (service *AuthClient) Validate(req *auth.ValidateRequest) (res *auth.ValidateResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	defer cancel()
	res, err = service.authClient.Validate(ctx, req)
	return
}
