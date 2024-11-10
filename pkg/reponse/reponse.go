package reponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json"data"`
}


// success response

func SuccessReponse(c *gin.Context, code int, data interface{}){
	c.JSON(http.StatusOK, ReponseData{
		Code: code,
		Message: msg[code],
		Data: data,
	})
}

func ErrorResponse(c *gin.Context, code int){
	c.JSON(http.StatusOK, ReponseData{
		Code: code,
		Message: msg[code],
		Data: nil,
	})
}