package respond

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int	`json:"code"`
	Message string `json:"message"`
	Data 	interface{} `json:"data"`
	Error  string      `json:"error"`
}

func Success (message string, data interface{},  c *gin.Context){
	c.JSON(http.StatusOK, Response{Code: 200,Message: message,})
}

func Failed(message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: 200,Message: message,})
}