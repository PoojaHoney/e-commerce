package main

import (
	"fmt"
	"log"
	"net"

	"e-commerce/products_service/pkg/config"
	"e-commerce/products_service/pkg/db"
	pb "e-commerce/products_service/pkg/pb"
	services "e-commerce/products_service/pkg/services"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
