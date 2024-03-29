package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT")

        if c.Request.Method == "OPTIONS" {
            log.Println(c.Request)
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func AddCorelationID() gin.HandlerFunc {
	return func(c *gin.Context){
        c.Request.Header.Add("X-Correlation-Id", uuid.NewString())
		c.Next()
	}
}


