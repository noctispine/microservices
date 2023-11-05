package routes

import (
	"log"
	"net/http"

	"github.com/capstone-project-bunker/backend/services/gateway/pkg/responses"
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/users/pb"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context, client pb.UserServiceClient)  {
	res, err := client.GetAll(c, &pb.GetAllRequest{})
	
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	if res.BaseResponse.Error != "" {
		responses.AbortWithStatusJSONErrorMessage(c, res.BaseResponse.Status, res.BaseResponse.Error)
		return
	}

	c.JSON(http.StatusOK, res.Users)
}

func GetById(c *gin.Context, client pb.UserServiceClient) {
	idString := c.Params.ByName("id")
	
	if idString == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := client.GetById(c, &pb.GetByIdRequest{
		Id: idString,
	})

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if res.BaseResponse.Error != "" {
		responses.AbortWithStatusJSONErrorMessage(c, res.BaseResponse.Status, res.BaseResponse.Error)
		return
	}

	c.JSON(int(res.BaseResponse.Status), res.User)
}


func DeleteById(c *gin.Context, client pb.UserServiceClient) {
	idString := c.Params.ByName("id")

	if idString != "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := client.DeleteById(c, &pb.DeleteByIdRequest{
		Id: idString,
	})

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if res.BaseResponse.Error != "" {
		responses.AbortWithStatusJSONErrorMessage(c, res.BaseResponse.Status, res.BaseResponse.Error)
		return
	}

	c.Status(int(res.BaseResponse.Status))

}