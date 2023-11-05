package users

import (
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/auth"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/config"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/users/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
    svc := &ServiceClient{
        Client: InitServiceClient(c),
    }

	a := auth.InitAuthMiddleware(authSvc)

    routes := r.Group("/users")
    {
        routes.Use(a.ValidateToken) 
        routes.GET("/", svc.GetAll)
        routes.DELETE("/:id", svc.DeleteById)
    }


    // routes.POST("/register", svc.Register)

    return svc
}

func (svc *ServiceClient) GetAll(c *gin.Context) {
    routes.GetAll(c, svc.Client)
}

func (svc *ServiceClient) DeleteById(c *gin.Context) {
    routes.DeleteById(c, svc.Client)
}

// func (svc *ServiceClient) Register(c *gin.Context) {
//     routes.Register(c, svc.Client)
// }