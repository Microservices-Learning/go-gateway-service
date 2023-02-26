package main

import (
	"go-gateway-service/config"
	"go-gateway-service/server"
	"log"
)

// @title       Go Microservices REST API
// @version     1.0
// @description This is a documents for Go Microservices REST API.

// @contact.name  Developer: Hoang Ngo
// @contact.email hoanggg2110@gmail.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host                       localhost:8000
// @BasePath                   /
// @schemes                    https
// @securityDefinitions.apikey Jwt
// @in                         header
// @name                       Authorization
func main() {
	start()
}

func start() {
	err := config.Config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	goServer := server.NewServer()
	goServer.Initialize()
	goServer.InitializeRoutes()
	goServer.Run(config.Config.Port)
	//var conn *grpc.ClientConn
	//conn, err := grpc.Dial(":6565", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("did not connect: %s", err)
	//}
	//defer conn.Close()
	//c := auth.NewAuthServiceClient(conn)
	//response, err := c.Login(context.Background(), &auth.LoginRequest{Username: "foo", Password: "s"})
	//if err != nil {
	//	log.Fatalf("Error when calling SayHello: %s", err)
	//}
	//log.Printf("Response from server: %s", response.Response)
}
