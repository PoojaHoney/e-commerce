package main

import (
	"log"

	"e-commerce/api_gateway/pkg/auth"
	"e-commerce/api_gateway/pkg/config"
	"e-commerce/api_gateway/pkg/order"
	"e-commerce/api_gateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
