package server

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-gateway-service/client/service"
	"go-gateway-service/config"
	"go-gateway-service/controller"
	_ "go-gateway-service/docs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func (server *Server) InitializeRoutes() {
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	authCC, err := grpc.Dial(config.Config.AccountClientUrl, transportOption)

	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	//defer authCC.Close()
	//
	//c := auth.NewAuthServiceClient(authCC)
	//
	//response, err := c.Login(context.Background(), &auth.LoginRequest{Username: "Hello From Client!", Password: "s"})
	//if err != nil {
	//	log.Fatalf("Error when calling Login: %s", err)
	//}
	//log.Printf("Response from server: %s", response.Response)

	authServiceClient := service.NewAuthClient(authCC)
	authController := controller.NewAuthController(authServiceClient)

	auth := server.router.Group("/api/v2")
	auth.Use()
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
	}

	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
