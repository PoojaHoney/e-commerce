package order

import (
	"github.com/gin-gonic/gin"
	"e-commerce/api_gateway/pkg/auth"
	"e-commerce/api_gateway/pkg/config"
	"e-commerce/api_gateway/pkg/order/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
