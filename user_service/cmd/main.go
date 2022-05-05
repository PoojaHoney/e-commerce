package main

import (
	"fmt"
	"log"
	"net"

	"e-commerce/user_service/pkg/config"
	"e-commerce/user_service/pkg/db"
	"e-commerce/user_service/pkg/pb"
	"e-commerce/user_service/pkg/services"
	"e-commerce/user_service/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "user_service",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
