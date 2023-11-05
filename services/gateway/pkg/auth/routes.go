package auth

import (
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/auth/routes"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
    svc := &ServiceClient{
        Client: InitServiceClient(c),
    }

    routes := r.Group("/auth")
    {
        routes.POST("/login", svc.Login)
        routes.POST("/register", svc.Register)
    }

    return svc
}

func (svc *ServiceClient) Login(c *gin.Context) {
    routes.Login(c, svc.Client)
}

func (svc *ServiceClient) Register(c *gin.Context) {
    routes.Register(c, svc.Client)
}