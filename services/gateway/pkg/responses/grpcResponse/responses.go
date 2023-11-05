package grpcResponse

import (
	"github.com/capstone-project-bunker/backend/services/gateway/pkg/baseProto"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Body any
	BaseResponse baseProto.Response
}

type GenericResponse interface {
	GetBaseResponse() baseProto.Response
}


// parsing response in runtime to generic response struct
// in development...
// func ParseResponse(resp any) (*Response, error){
// 	vPtr := reflect.ValueOf(resp)
// 	v := reflect.Indirect(vPtr)
// 	fmt.Println("\n\n\nYOOOASD")
// 	if v.FieldByName("BaseResponse").Kind() != reflect.Ptr {
// 		return nil, fmt.Errorf("response must contain 'BaseResponse' field")
// 	}
	
// 	baseRespPtr := v.FieldByName("BaseResponse")
// 	baseResp := reflect.Indirect(baseRespPtr)
// 	fmt.Println("BASE RESPPPP, ", baseResp)
// 	status := baseResp.FieldByName("Status")
// 	errorMessage := baseResp.FieldByName("Error")
	
// 	var response Response
// 	vResponse := reflect.ValueOf(&response)
// 	vResponse.Elem().FieldByName("BaseResponse").FieldByName("Status").SetInt(status.Int())
// 	vResponse.Elem().FieldByName("BaseResponse").FieldByName("Error").SetString(errorMessage.String())
// 	fmt.Println("STATUS: ", response)
// 	return &response, nil
// }

func (r *Response) AbortWithStatusJSONError(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(int(r.BaseResponse.Status), r.Body)
}
